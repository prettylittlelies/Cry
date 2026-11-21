package commands

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/prettylittlelies/cry/snowflake"
)

type BucketType int

const (
	BucketDefault BucketType = iota
	BucketUser
	BucketGuild
	BucketChannel
	BucketMember
	BucketCategory
	BucketRole
)

func (b BucketType) String() string {
	return [...]string{"default", "user", "guild", "channel", "member", "category", "role"}[b]
}

type BucketSource interface {
	AuthorID() snowflake.ID
	GuildID() snowflake.ID
	ChannelID() snowflake.ID
	CategoryID() snowflake.ID
	TopRoleID() snowflake.ID
	Private() bool
}

type bucketKey struct {
	a, b snowflake.ID
}

func (b BucketType) Key(src BucketSource) any {
	switch b {
	case BucketUser:
		return src.AuthorID()
	case BucketGuild:
		if src.GuildID() != 0 {
			return src.GuildID()
		}
		return src.AuthorID()
	case BucketChannel:
		return src.ChannelID()
	case BucketMember:
		return bucketKey{src.GuildID(), src.AuthorID()}
	case BucketCategory:
		if c := src.CategoryID(); c != 0 {
			return c
		}
		return src.ChannelID()
	case BucketRole:
		if src.Private() {
			return src.ChannelID()
		}
		return src.TopRoleID()
	}
	return nil
}

type Cooldown struct {
	Rate   int
	Per    time.Duration
	window time.Time
	tokens int
	last   time.Time
}

func NewCooldown(rate int, per time.Duration) *Cooldown {
	return &Cooldown{Rate: rate, Per: per, tokens: rate}
}

func (c *Cooldown) Tokens(now time.Time) int {
	tokens := max(c.tokens, 0)
	if now.After(c.window.Add(c.Per)) {
		tokens = c.Rate
	}
	return tokens
}

func (c *Cooldown) RetryAfter(now time.Time) time.Duration {
	if c.Tokens(now) == 0 {
		return c.Per - now.Sub(c.window)
	}
	return 0
}

func (c *Cooldown) UpdateRateLimit(now time.Time, tokens int) time.Duration {
	c.last = now
	c.tokens = c.Tokens(now)
	if c.tokens == c.Rate {
		c.window = now
	}
	c.tokens -= tokens
	if c.tokens < 0 {
		return c.Per - now.Sub(c.window)
	}
	return 0
}

func (c *Cooldown) Reset() {
	c.tokens = c.Rate
	c.last = time.Time{}
}

func (c *Cooldown) Copy() *Cooldown {
	return NewCooldown(c.Rate, c.Per)
}

type CooldownMapping struct {
	mu       sync.Mutex
	cooldown *Cooldown
	factory  func(BucketSource) *Cooldown
	Type     BucketType
	cache    map[any]*Cooldown
}

func NewCooldownMapping(c *Cooldown, t BucketType) *CooldownMapping {
	return &CooldownMapping{cooldown: c, Type: t, cache: map[any]*Cooldown{}}
}

func NewDynamicCooldownMapping(factory func(BucketSource) *Cooldown, t BucketType) *CooldownMapping {
	return &CooldownMapping{factory: factory, Type: t, cache: map[any]*Cooldown{}}
}

func (m *CooldownMapping) Valid() bool {
	return m.cooldown != nil || m.factory != nil
}

func (m *CooldownMapping) create(src BucketSource) *Cooldown {
	if m.factory != nil {
		return m.factory(src)
	}
	return m.cooldown.Copy()
}

func (m *CooldownMapping) evict(now time.Time) {
	for k, v := range m.cache {
		if now.After(v.last.Add(v.Per)) {
			delete(m.cache, k)
		}
	}
}

func (m *CooldownMapping) bucket(src BucketSource, now time.Time) *Cooldown {
	if m.Type == BucketDefault && m.factory == nil {
		return m.cooldown
	}
	m.evict(now)
	key := m.Type.Key(src)
	b, ok := m.cache[key]
	if !ok {
		if b = m.create(src); b != nil {
			m.cache[key] = b
		}
	}
	return b
}

func (m *CooldownMapping) Bucket(src BucketSource, now time.Time) *Cooldown {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.bucket(src, now)
}

func (m *CooldownMapping) UpdateRateLimit(src BucketSource, now time.Time, tokens int) (*Cooldown, time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	b := m.bucket(src, now)
	if b == nil {
		return nil, 0
	}
	return b, b.UpdateRateLimit(now, tokens)
}

type semaphore struct {
	value   int
	waiters []chan struct{}
}

type MaxConcurrency struct {
	Number  int
	Per     BucketType
	Wait    bool
	mu      sync.Mutex
	mapping map[any]*semaphore
}

func NewMaxConcurrency(number int, per BucketType, wait bool) *MaxConcurrency {
	if number <= 0 {
		panic("commands: max_concurrency 'number' cannot be less than 1")
	}
	return &MaxConcurrency{Number: number, Per: per, Wait: wait, mapping: map[any]*semaphore{}}
}

func (m *MaxConcurrency) Copy() *MaxConcurrency {
	return NewMaxConcurrency(m.Number, m.Per, m.Wait)
}

func (m *MaxConcurrency) Acquire(ctx context.Context, src BucketSource) error {
	key := m.Per.Key(src)
	for {
		m.mu.Lock()
		sem, ok := m.mapping[key]
		if !ok {
			sem = &semaphore{value: m.Number}
			m.mapping[key] = sem
		}
		if sem.value > 0 {
			sem.value--
			m.mu.Unlock()
			return nil
		}
		if !m.Wait {
			m.mu.Unlock()
			return &MaxConcurrencyReached{newKind(commandErr, "Too many people are using this command. It can only be used %s concurrently.", concurrencyLabel(m.Number, m.Per)), m.Number, m.Per}
		}
		wake := make(chan struct{})
		sem.waiters = append(sem.waiters, wake)
		m.mu.Unlock()
		select {
		case <-wake:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (m *MaxConcurrency) Release(src BucketSource) {
	key := m.Per.Key(src)
	m.mu.Lock()
	defer m.mu.Unlock()
	sem, ok := m.mapping[key]
	if !ok {
		return
	}
	sem.value++
	if len(sem.waiters) > 0 {
		close(sem.waiters[0])
		sem.waiters = sem.waiters[1:]
	}
	if sem.value >= m.Number && len(sem.waiters) == 0 {
		delete(m.mapping, key)
	}
}

func concurrencyLabel(n int, per BucketType) string {
	suffix := "per " + per.String()
	if per == BucketDefault {
		suffix = "globally"
	}
	plural := "times"
	if n == 1 {
		plural = "time"
	}
	return strconv.Itoa(n) + " " + plural + " " + suffix
}

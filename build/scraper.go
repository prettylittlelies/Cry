package build

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/prettylittlelies/cry/internal/fetch"
)

const browserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/152.0.0.0 Safari/537.36"

type Source interface {
	Fetch(ctx context.Context) (*Snapshot, error)
}

type Option func(*Scraper)

func WithDoer(d fetch.Doer) Option {
	return func(s *Scraper) { s.client.Doer = d }
}

func WithEndpoints(e Endpoints) Option {
	return func(s *Scraper) { s.endpoints = e }
}

func WithChannels(c ...Channel) Option {
	return func(s *Scraper) { s.channels = c }
}

type Scraper struct {
	client    fetch.Client
	endpoints Endpoints
	channels  []Channel
}

var _ Source = (*Scraper)(nil)

func NewScraper(opts ...Option) *Scraper {
	s := &Scraper{
		client:    fetch.New(&http.Client{Timeout: 30 * time.Second}, http.Header{"User-Agent": {browserAgent}}),
		endpoints: DefaultEndpoints(),
		channels:  Channels(),
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

func (s *Scraper) Fetch(ctx context.Context) (*Snapshot, error) {
	index := s.electronIndex(ctx)
	snap := &Snapshot{FetchedAt: time.Now().UTC(), Releases: make(map[Channel]Release, len(s.channels))}
	results := make([]Release, len(s.channels))
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		errs []error
	)
	fail := func(source string, c Channel, err error) {
		mu.Lock()
		errs = append(errs, &SourceError{Source: source, Channel: c, Err: err})
		mu.Unlock()
	}
	for i, c := range s.channels {
		results[i].Channel = c
		wg.Go(func() {
			n, err := s.number(ctx, c)
			if err != nil {
				fail("web", c, err)
			}
			results[i].Number = n
		})
		wg.Go(func() {
			w, err := s.windows(ctx, c)
			if err != nil {
				fail("windows", c, err)
			}
			results[i].Windows = w
		})
		wg.Go(func() {
			m, err := s.macOS(ctx, c, index)
			if err != nil {
				fail("macos", c, err)
			}
			results[i].MacOS = m
		})
	}
	wg.Go(func() {
		m, err := s.mobile(ctx)
		if err != nil {
			fail("mobile", "", err)
		}
		snap.Mobile = m
	})
	wg.Wait()
	for _, r := range results {
		snap.Releases[r.Channel] = r
	}
	return snap, errors.Join(errs...)
}

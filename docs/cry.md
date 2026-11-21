# cry

```go
import "github.com/prettylittlelies/cry"
```

## Index

- [func On](#func-on)
- [func ReactionEmoji](#func-reactionemoji)
- [func WaitFor](#func-waitfor)
- [type Channel](#type-channel)
- [type Client](#type-client)
- [type Guild](#type-guild)
- [type HistoryOptions](#type-historyoptions)
- [type Member](#type-member)
- [type Message](#type-message)
- [type Option](#type-option)
- [type RawEvent](#type-rawevent)
- [type User](#type-user)

## Constants

```go
const DefaultMessageCache = 1000
```

## Variables

```go
var ErrUnknownChannel = errors.New("cry: release not present in build snapshot")
```

## Functions

### func On

```go
func On[T any](c *Client, fn func(context.Context, *T)) (remove func())
```

### func ReactionEmoji

```go
func ReactionEmoji(emoji string) string
```

### func WaitFor

```go
func WaitFor[T any](ctx context.Context, c *Client, check func(*T) bool) (*T, error)
```

## Types

### type Channel

```go
type Channel struct {
	ID snowflake.ID
	// contains filtered or unexported fields
}
```

#### func (Channel) BulkDelete

```go
func (ch Channel) BulkDelete(ctx context.Context, ids ...snowflake.ID) error
```

#### func (Channel) Cached

```go
func (ch Channel) Cached() (discord.Channel, bool)
```

#### func (Channel) Delete

```go
func (ch Channel) Delete(ctx context.Context) error
```

#### func (Channel) Edit

```go
func (ch Channel) Edit(ctx context.Context, params discord.ModifyChannelParams) (*discord.Channel, error)
```

#### func (Channel) Fetch

```go
func (ch Channel) Fetch(ctx context.Context) (*discord.Channel, error)
```

#### func (Channel) History

```go
func (ch Channel) History(ctx context.Context, opts HistoryOptions) iter.Seq2[discord.Message, error]
```

#### func (Channel) Message

```go
func (ch Channel) Message(id snowflake.ID) Message
```

#### func (Channel) Pins

```go
func (ch Channel) Pins(ctx context.Context) ([]discord.MessagePin, error)
```

#### func (Channel) Send

```go
func (ch Channel) Send(ctx context.Context, content string) (*discord.Message, error)
```

#### func (Channel) SendMessage

```go
func (ch Channel) SendMessage(ctx context.Context, params discord.CreateMessageParams) (*discord.Message, error)
```

#### func (Channel) Typing

```go
func (ch Channel) Typing(ctx context.Context) error
```

### type Client

```go
type Client struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(ctx context.Context, token string, opts ...Option) (*Client, error)
```

#### func (*Client) API

```go
func (c *Client) API() *api.API
```

#### func (*Client) Channel

```go
func (c *Client) Channel(id snowflake.ID) Channel
```

#### func (*Client) Dispatch

```go
func (c *Client) Dispatch(ctx context.Context, v any)
```

#### func (*Client) Gateway

```go
func (c *Client) Gateway() *gateway.Gateway
```

#### func (*Client) Guild

```go
func (c *Client) Guild(id snowflake.ID) Guild
```

#### func (*Client) Member

```go
func (c *Client) Member(guild, user snowflake.ID) Member
```

#### func (*Client) Message

```go
func (c *Client) Message(channel, id snowflake.ID) Message
```

#### func (*Client) REST

```go
func (c *Client) REST() *rest.Client
```

#### func (*Client) Run

```go
func (c *Client) Run(ctx context.Context) error
```

#### func (*Client) Session

```go
func (c *Client) Session() *properties.Session
```

#### func (*Client) State

```go
func (c *Client) State() *state.State
```

#### func (*Client) User

```go
func (c *Client) User(id snowflake.ID) User
```

### type Guild

```go
type Guild struct {
	ID snowflake.ID
	// contains filtered or unexported fields
}
```

#### func (Guild) Cached

```go
func (g Guild) Cached() (discord.GatewayGuild, bool)
```

#### func (Guild) Fetch

```go
func (g Guild) Fetch(ctx context.Context) (*discord.Guild, error)
```

#### func (Guild) FetchMembers

```go
func (g Guild) FetchMembers(ctx context.Context, users ...snowflake.ID) ([]discord.GuildMember, error)
```

#### func (Guild) Leave

```go
func (g Guild) Leave(ctx context.Context) error
```

#### func (Guild) Member

```go
func (g Guild) Member(user snowflake.ID) Member
```

#### func (Guild) Members

```go
func (g Guild) Members() []discord.GuildMember
```

#### func (Guild) QueryMembers

```go
func (g Guild) QueryMembers(ctx context.Context, query string, limit int, presences bool) ([]discord.GuildMember, error)
```

#### func (Guild) Roles

```go
func (g Guild) Roles(ctx context.Context) ([]discord.Role, error)
```

#### func (Guild) SetNick

```go
func (g Guild) SetNick(ctx context.Context, nick string) error
```

### type HistoryOptions

```go
type HistoryOptions struct {
	Limit  int
	Before snowflake.ID
	After  snowflake.ID
	Around snowflake.ID
}
```

### type Member

```go
type Member struct {
	GuildID snowflake.ID
	UserID  snowflake.ID
	// contains filtered or unexported fields
}
```

#### func (Member) AddRoles

```go
func (m Member) AddRoles(ctx context.Context, roles ...snowflake.ID) error
```

#### func (Member) Ban

```go
func (m Member) Ban(ctx context.Context, deleteMessages time.Duration) error
```

#### func (Member) Cached

```go
func (m Member) Cached() (discord.GuildMember, bool)
```

#### func (Member) Edit

```go
func (m Member) Edit(ctx context.Context, params discord.ModifyGuildMemberParams) error
```

#### func (Member) Fetch

```go
func (m Member) Fetch(ctx context.Context) (*discord.GuildMember, error)
```

#### func (Member) Guild

```go
func (m Member) Guild() Guild
```

#### func (Member) Kick

```go
func (m Member) Kick(ctx context.Context) error
```

#### func (Member) RemoveRoles

```go
func (m Member) RemoveRoles(ctx context.Context, roles ...snowflake.ID) error
```

#### func (Member) Timeout

```go
func (m Member) Timeout(ctx context.Context, until time.Time) error
```

#### func (Member) Unban

```go
func (m Member) Unban(ctx context.Context) error
```

#### func (Member) User

```go
func (m Member) User() User
```

### type Message

```go
type Message struct {
	ChannelID snowflake.ID
	ID        snowflake.ID
	// contains filtered or unexported fields
}
```

#### func (Message) Cached

```go
func (m Message) Cached() (discord.MessageCreateEvent, bool)
```

#### func (Message) Channel

```go
func (m Message) Channel() Channel
```

#### func (Message) ClearReaction

```go
func (m Message) ClearReaction(ctx context.Context, emoji string) error
```

#### func (Message) ClearReactions

```go
func (m Message) ClearReactions(ctx context.Context) error
```

#### func (Message) Delete

```go
func (m Message) Delete(ctx context.Context) error
```

#### func (Message) Edit

```go
func (m Message) Edit(ctx context.Context, content string) (*discord.Message, error)
```

#### func (Message) EditMessage

```go
func (m Message) EditMessage(ctx context.Context, params discord.EditMessageParams) (*discord.Message, error)
```

#### func (Message) Fetch

```go
func (m Message) Fetch(ctx context.Context) (discord.Message, error)
```

#### func (Message) Pin

```go
func (m Message) Pin(ctx context.Context) error
```

#### func (Message) React

```go
func (m Message) React(ctx context.Context, emoji string) error
```

#### func (Message) RemoveReaction

```go
func (m Message) RemoveReaction(ctx context.Context, emoji string, user snowflake.ID) error
```

#### func (Message) Reply

```go
func (m Message) Reply(ctx context.Context, content string) (*discord.Message, error)
```

#### func (Message) ReplyMessage

```go
func (m Message) ReplyMessage(ctx context.Context, params discord.CreateMessageParams) (*discord.Message, error)
```

#### func (Message) SuperReact

```go
func (m Message) SuperReact(ctx context.Context, emoji string) error
```

#### func (Message) Unpin

```go
func (m Message) Unpin(ctx context.Context) error
```

#### func (Message) Unreact

```go
func (m Message) Unreact(ctx context.Context, emoji string) error
```

### type Option

```go
type Option func(*config)
```

#### func WithAPIBase

```go
func WithAPIBase(url string) Option
```

#### func WithCaptcha

```go
func WithCaptcha(h rest.CaptchaHandler) Option
```

#### func WithChannel

```go
func WithChannel(ch build.Channel) Option
```

#### func WithFastConnect

```go
func WithFastConnect(enabled bool) Option
```

#### func WithGatewayEndpoint

```go
func WithGatewayEndpoint(url string) Option
```

#### func WithHost

```go
func WithHost(h properties.Host) Option
```

#### func WithLogger

```go
func WithLogger(l *slog.Logger) Option
```

#### func WithMessageCache

```go
func WithMessageCache(n int) Option
```

#### func WithRootCAs

```go
func WithRootCAs(pool *x509.CertPool) Option
```

#### func WithSnapshot

```go
func WithSnapshot(s *build.Snapshot) Option
```

### type RawEvent

```go
type RawEvent struct {
	Type string
	Seq  int64
	Data json.RawMessage
}
```

### type User

```go
type User struct {
	ID snowflake.ID
	// contains filtered or unexported fields
}
```

#### func (User) Block

```go
func (u User) Block(ctx context.Context) error
```

#### func (User) Cached

```go
func (u User) Cached() (discord.PartialUser, bool)
```

#### func (User) DM

```go
func (u User) DM(ctx context.Context) (Channel, error)
```

#### func (User) Fetch

```go
func (u User) Fetch(ctx context.Context) (*discord.User, error)
```

#### func (User) Friend

```go
func (u User) Friend(ctx context.Context) error
```

#### func (User) Note

```go
func (u User) Note(ctx context.Context) (string, error)
```

#### func (User) Profile

```go
func (u User) Profile(ctx context.Context, query discord.GetUserProfileQuery) (*discord.GetUserProfileResponse, error)
```

#### func (User) Relationship

```go
func (u User) Relationship() (discord.Relationship, bool)
```

#### func (User) RemoveRelationship

```go
func (u User) RemoveRelationship(ctx context.Context) error
```

#### func (User) Send

```go
func (u User) Send(ctx context.Context, content string) (*discord.Message, error)
```

#### func (User) SetNote

```go
func (u User) SetNote(ctx context.Context, note string) error
```


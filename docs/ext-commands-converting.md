# converting

```go
import "github.com/prettylittlelies/cry/ext/commands/converting"
```

## Index

- [func Object](#func-object)
- [func ParseTimestamp](#func-parsetimestamp)
- [type CleanOptions](#type-cleanoptions)
- [type Colour](#type-colour)
- [type PartialEmoji](#type-partialemoji)
- [type PartialMessage](#type-partialmessage)
- [type ResolvedInvite](#type-resolvedinvite)
- [type Scope](#type-scope)
- [type Thread](#type-thread)

## Variables

```go
var (
	ErrNotFound           = errors.New("converting: not found")
	ErrChannelNotFound    = errors.New("converting: channel not found")
	ErrChannelNotReadable = errors.New("converting: channel not readable")
)
```

## Functions

### func Object

```go
func Object(argument string) (snowflake.ID, bool)
```

### func ParseTimestamp

```go
func ParseTimestamp(argument string) (time.Time, bool)
```

## Types

### type CleanOptions

```go
type CleanOptions struct {
	FixChannelMentions bool
	UseNicknames       bool
	EscapeMarkdown     bool
	RemoveMarkdown     bool
}
```

#### func DefaultCleanOptions

```go
func DefaultCleanOptions() CleanOptions
```

### type Colour

```go
type Colour uint32
```

#### func ParseColour

```go
func ParseColour(argument string) (Colour, bool)
```

#### func RGB

```go
func RGB(r, g, b uint8) Colour
```

#### func (Colour) B

```go
func (c Colour) B() uint8
```

#### func (Colour) G

```go
func (c Colour) G() uint8
```

#### func (Colour) R

```go
func (c Colour) R() uint8
```

#### func (Colour) String

```go
func (c Colour) String() string
```

### type PartialEmoji

```go
type PartialEmoji struct {
	ID       snowflake.ID
	Name     string
	Animated bool
}
```

#### func ParsePartialEmoji

```go
func ParsePartialEmoji(argument string) (PartialEmoji, bool)
```

### type PartialMessage

```go
type PartialMessage struct {
	ID      snowflake.ID
	Channel discord.Channel
}
```

### type ResolvedInvite

```go
type ResolvedInvite struct {
	Code    string
	EventID string
}
```

#### func ResolveInvite

```go
func ResolveInvite(argument string) (ResolvedInvite, bool)
```

### type Scope

```go
type Scope struct {
	Context context.Context
	Client  *cry.Client
	GuildID snowflake.ID
}
```

#### func (Scope) Channel

```go
func (s Scope) Channel(argument string) (discord.Channel, bool)
```

#### func (Scope) ChannelOfType

```go
func (s Scope) ChannelOfType(argument string, types ...discord.ChannelType) (discord.Channel, bool)
```

#### func (Scope) Clean

```go
func (s Scope) Clean(opts CleanOptions, mentions []discord.PartialUser, argument string) string
```

#### func (Scope) Emoji

```go
func (s Scope) Emoji(argument string) (discord.Emoji, bool)
```

#### func (Scope) Guild

```go
func (s Scope) Guild(argument string) (discord.GatewayGuild, bool)
```

#### func (Scope) Invite

```go
func (s Scope) Invite(argument string) (discord.Invite, bool)
```

#### func (Scope) Member

```go
func (s Scope) Member(argument string) (discord.GuildMember, bool)
```

#### func (Scope) Message

```go
func (s Scope) Message(argument string, current snowflake.ID) (discord.Message, error)
```

#### func (Scope) PartialMessage

```go
func (s Scope) PartialMessage(argument string, current snowflake.ID) (PartialMessage, error)
```

#### func (Scope) Role

```go
func (s Scope) Role(argument string) (discord.Role, bool)
```

#### func (Scope) ScheduledEvent

```go
func (s Scope) ScheduledEvent(argument string) (discord.GuildScheduledEvent, bool)
```

#### func (Scope) Sticker

```go
func (s Scope) Sticker(argument string) (discord.Sticker, bool)
```

#### func (Scope) Thread

```go
func (s Scope) Thread(argument string) (Thread, bool)
```

#### func (Scope) User

```go
func (s Scope) User(argument string) (discord.PartialUser, bool)
```

### type Thread

```go
type Thread struct{ discord.Channel }
```


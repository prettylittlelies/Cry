# state

```go
import "github.com/prettylittlelies/cry/state"
```

## Index

- [type State](#type-state)

## Types

### type State

```go
type State struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(messageLimit int) *State
```

#### func (*State) AllGuilds

```go
func (s *State) AllGuilds() []discord.GatewayGuild
```

#### func (*State) Apply

```go
func (s *State) Apply(event any)
```

#### func (*State) Channel

```go
func (s *State) Channel(id snowflake.ID) (discord.Channel, bool)
```

#### func (*State) ChannelPermissions

```go
func (s *State) ChannelPermissions(channel, user snowflake.ID) discord.Permissions
```

#### func (*State) DirectMessage

```go
func (s *State) DirectMessage(user snowflake.ID) (snowflake.ID, bool)
```

#### func (*State) Guild

```go
func (s *State) Guild(id snowflake.ID) (discord.GatewayGuild, bool)
```

#### func (*State) GuildChannels

```go
func (s *State) GuildChannels(guild snowflake.ID) []discord.Channel
```

#### func (*State) GuildPermissions

```go
func (s *State) GuildPermissions(guild, user snowflake.ID) discord.Permissions
```

#### func (*State) Guilds

```go
func (s *State) Guilds() []snowflake.ID
```

#### func (*State) Me

```go
func (s *State) Me() discord.User
```

#### func (*State) Member

```go
func (s *State) Member(guild, user snowflake.ID) (discord.GuildMember, bool)
```

#### func (*State) Members

```go
func (s *State) Members(guild snowflake.ID) []discord.GuildMember
```

#### func (*State) Message

```go
func (s *State) Message(id snowflake.ID) (discord.MessageCreateEvent, bool)
```

#### func (*State) Presence

```go
func (s *State) Presence(user snowflake.ID) (discord.PresenceUpdateEvent, bool)
```

#### func (*State) Relationship

```go
func (s *State) Relationship(user snowflake.ID) (discord.Relationship, bool)
```

#### func (*State) SessionID

```go
func (s *State) SessionID() string
```

#### func (*State) User

```go
func (s *State) User(id snowflake.ID) (discord.PartialUser, bool)
```

#### func (*State) UserByName

```go
func (s *State) UserByName(name string) (discord.PartialUser, bool)
```

#### func (*State) VoiceState

```go
func (s *State) VoiceState(guild, user snowflake.ID) (discord.VoiceState, bool)
```

#### func (*State) VoiceStates

```go
func (s *State) VoiceStates(guild snowflake.ID) []discord.VoiceState
```


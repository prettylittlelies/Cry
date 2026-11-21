# gateway

```go
import "github.com/prettylittlelies/cry/gateway"
```

## Index

- [type Capabilities](#type-capabilities)
- [type ClientState](#type-clientstate)
- [type CloseCode](#type-closecode)
- [type Config](#type-config)
- [type Event](#type-event)
- [type Gateway](#type-gateway)
- [type GuildSubscription](#type-guildsubscription)
- [type MemberRequest](#type-memberrequest)
- [type Opcode](#type-opcode)
- [type Presence](#type-presence)
- [type QoS](#type-qos)
- [type QoSReason](#type-qosreason)
- [type StreamCreate](#type-streamcreate)

## Constants

```go
const (
	DefaultEndpoint = "wss://gateway.discord.gg"
	APIVersion      = 9
)
```

```go
const ColdStart = "COLD_START"
```

```go
const QoSVersion = 31
```

## Variables

```go
var ErrAuthenticationFailed = errors.New("gateway: authentication failed")
```

## Types

### type Capabilities

```go
type Capabilities uint64
```

```go
const DefaultCapabilities Capabilities = 1767421
```

### type ClientState

```go
type ClientState struct {
	GuildVersions map[string]int64 `json:"guild_versions"`
}
```

### type CloseCode

```go
type CloseCode int
```

```go
const (
	CloseUnknownError         CloseCode = 4000
	CloseUnknownOpcode        CloseCode = 4001
	CloseDecodeError          CloseCode = 4002
	CloseNotAuthenticated     CloseCode = 4003
	CloseAuthenticationFailed CloseCode = 4004
	CloseAlreadyAuthenticated CloseCode = 4005
	CloseSessionNoLongerValid CloseCode = 4006
	CloseInvalidSeq           CloseCode = 4007
	CloseRateLimited          CloseCode = 4008
	CloseSessionTimedOut      CloseCode = 4009
	CloseInvalidShard         CloseCode = 4010
	CloseShardingRequired     CloseCode = 4011
	CloseInvalidAPIVersion    CloseCode = 4012
	CloseInvalidIntents       CloseCode = 4013
	CloseDisallowedIntents    CloseCode = 4014
)
```

### type Config

```go
type Config struct {
	Token          string
	Session        *properties.Session
	Browser        websocket.Browser
	Hello          fingerprint.ClientHello
	Jar            http.CookieJar
	Endpoint       string
	Encoding       codec.Encoding
	Compression    codec.Compression
	Capabilities   Capabilities
	Presence       Presence
	InstallationID string
	FastConnect    bool
	Handler        func(Event)
	Logger         *slog.Logger
}
```

### type Event

```go
type Event struct {
	Type string
	Seq  int64
	Data json.RawMessage
}
```

### type Gateway

```go
type Gateway struct {
	// contains filtered or unexported fields
}
```

#### func New

```go
func New(cfg Config) *Gateway
```

#### func (*Gateway) Browser

```go
func (g *Gateway) Browser() websocket.Browser
```

#### func (*Gateway) CallConnect

```go
func (g *Gateway) CallConnect(ctx context.Context, channel snowflake.ID) error
```

#### func (*Gateway) Dialer

```go
func (g *Gateway) Dialer() *fingerprint.Dialer
```

#### func (*Gateway) Jar

```go
func (g *Gateway) Jar() http.CookieJar
```

#### func (*Gateway) Presence

```go
func (g *Gateway) Presence() Presence
```

#### func (*Gateway) RequestChannelInfo

```go
func (g *Gateway) RequestChannelInfo(ctx context.Context, r discord.RequestChannelInfo) error
```

#### func (*Gateway) RequestChannelMemberCount

```go
func (g *Gateway) RequestChannelMemberCount(ctx context.Context, guild, channel snowflake.ID) error
```

#### func (*Gateway) RequestChannelStatuses

```go
func (g *Gateway) RequestChannelStatuses(ctx context.Context, guild snowflake.ID) error
```

#### func (*Gateway) RequestLastMessages

```go
func (g *Gateway) RequestLastMessages(ctx context.Context, r discord.RequestLastMessages) error
```

#### func (*Gateway) RequestMembers

```go
func (g *Gateway) RequestMembers(ctx context.Context, r MemberRequest) error
```

#### func (*Gateway) RequestSoundboardSounds

```go
func (g *Gateway) RequestSoundboardSounds(ctx context.Context, guilds ...snowflake.ID) error
```

#### func (*Gateway) Run

```go
func (g *Gateway) Run(ctx context.Context) error
```

#### func (*Gateway) SearchRecentMembers

```go
func (g *Gateway) SearchRecentMembers(ctx context.Context, r discord.SearchRecentMembers) error
```

#### func (*Gateway) Send

```go
func (g *Gateway) Send(ctx context.Context, op Opcode, d any) error
```

#### func (*Gateway) SetActivity

```go
func (g *Gateway) SetActivity(q QoS)
```

#### func (*Gateway) SetPresence

```go
func (g *Gateway) SetPresence(ctx context.Context, p Presence) error
```

#### func (*Gateway) StreamCreate

```go
func (g *Gateway) StreamCreate(ctx context.Context, s StreamCreate) error
```

#### func (*Gateway) StreamDelete

```go
func (g *Gateway) StreamDelete(ctx context.Context, key string) error
```

#### func (*Gateway) StreamPing

```go
func (g *Gateway) StreamPing(ctx context.Context, key string) error
```

#### func (*Gateway) StreamSetPaused

```go
func (g *Gateway) StreamSetPaused(ctx context.Context, key string, paused bool) error
```

#### func (*Gateway) StreamWatch

```go
func (g *Gateway) StreamWatch(ctx context.Context, key string) error
```

#### func (*Gateway) SubscribeGuilds

```go
func (g *Gateway) SubscribeGuilds(ctx context.Context, subscriptions map[snowflake.ID]GuildSubscription) error
```

#### func (*Gateway) UpdateVoiceState

```go
func (g *Gateway) UpdateVoiceState(ctx context.Context, v discord.UpdateVoiceState) error
```

### type GuildSubscription

```go
type GuildSubscription struct {
	Typing            *bool                     `json:"typing,omitempty"`
	Threads           *bool                     `json:"threads,omitempty"`
	Activities        *bool                     `json:"activities,omitempty"`
	MemberUpdates     *bool                     `json:"member_updates,omitempty"`
	Members           []snowflake.ID            `json:"members,omitempty"`
	Channels          map[snowflake.ID][][2]int `json:"channels,omitempty"`
	ThreadMemberLists []snowflake.ID            `json:"thread_member_lists,omitempty"`
}
```

### type MemberRequest

```go
type MemberRequest struct {
	GuildIDs  []snowflake.ID `json:"guild_id"`
	Query     *string        `json:"query"`
	Limit     *int           `json:"limit"`
	Presences bool           `json:"presences"`
	UserIDs   []snowflake.ID `json:"user_ids"`
	Nonce     string         `json:"nonce,omitempty"`
}
```

### type Opcode

```go
type Opcode int
```

```go
const (
	OpDispatch                        Opcode = 0
	OpHeartbeat                       Opcode = 1
	OpIdentify                        Opcode = 2
	OpPresenceUpdate                  Opcode = 3
	OpVoiceStateUpdate                Opcode = 4
	OpVoiceServerPing                 Opcode = 5
	OpResume                          Opcode = 6
	OpReconnect                       Opcode = 7
	OpRequestGuildMembers             Opcode = 8
	OpInvalidSession                  Opcode = 9
	OpHello                           Opcode = 10
	OpHeartbeatAck                    Opcode = 11
	OpCallConnect                     Opcode = 13
	OpGuildSubscriptions              Opcode = 14
	OpLobbyConnect                    Opcode = 15
	OpLobbyDisconnect                 Opcode = 16
	OpLobbyVoiceStates                Opcode = 17
	OpStreamCreate                    Opcode = 18
	OpStreamDelete                    Opcode = 19
	OpStreamWatch                     Opcode = 20
	OpStreamPing                      Opcode = 21
	OpStreamSetPaused                 Opcode = 22
	OpRequestGuildApplicationCommands Opcode = 24
	OpEmbeddedActivityLaunch          Opcode = 25
	OpEmbeddedActivityClose           Opcode = 26
	OpEmbeddedActivityUpdate          Opcode = 27
	OpRequestForumUnreads             Opcode = 28
	OpRemoteCommand                   Opcode = 29
	OpGetDeletedEntityIDsNotMatching  Opcode = 30
	OpRequestSoundboardSounds         Opcode = 31
	OpSpeedTestCreate                 Opcode = 32
	OpSpeedTestDelete                 Opcode = 33
	OpRequestLastMessages             Opcode = 34
	OpSearchRecentMembers             Opcode = 35
	OpRequestChannelStatuses          Opcode = 36
	OpGuildSubscriptionsBulk          Opcode = 37
	OpGuildChannelsResync             Opcode = 38
	OpRequestChannelMemberCount       Opcode = 39
	OpQoSHeartbeat                    Opcode = 40
	OpUpdateTimeSpentSessionID        Opcode = 41
	OpRequestChannelInfo              Opcode = 43
)
```

### type Presence

```go
type Presence struct {
	Status     discord.StatusType `json:"status"`
	Since      int64              `json:"since"`
	Activities []discord.Activity `json:"activities"`
	AFK        bool               `json:"afk"`
}
```

#### func DefaultPresence

```go
func DefaultPresence() Presence
```

### type QoS

```go
type QoS struct {
	Active  bool        `json:"active"`
	Ver     int         `json:"ver"`
	Reasons []QoSReason `json:"reasons"`
}
```

#### func NewQoS

```go
func NewQoS(reasons ...QoSReason) QoS
```

### type QoSReason

```go
type QoSReason string
```

```go
const (
	Foregrounded QoSReason = "foregrounded"
	RTCConnected QoSReason = "rtc_connected"
)
```

### type StreamCreate

```go
type StreamCreate struct {
	Type            discord.StreamType `json:"type"`
	GuildID         *snowflake.ID      `json:"guild_id"`
	ChannelID       snowflake.ID       `json:"channel_id"`
	PreferredRegion string             `json:"preferred_region,omitempty"`
}
```


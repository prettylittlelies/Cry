package gateway

type Opcode int

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

type CloseCode int

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

const DefaultCapabilities Capabilities = 1767421

type Capabilities uint64

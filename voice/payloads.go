package voice

import (
	"encoding/json"

	"github.com/prettylittlelies/cry/snowflake"
)

type Opcode int

const (
	OpIdentify           Opcode = 0
	OpSelectProtocol     Opcode = 1
	OpReady              Opcode = 2
	OpHeartbeat          Opcode = 3
	OpSessionDescription Opcode = 4
	OpSpeaking           Opcode = 5
	OpHeartbeatAck       Opcode = 6
	OpResume             Opcode = 7
	OpHello              Opcode = 8
	OpResumed            Opcode = 9
	OpClientsConnect     Opcode = 11
	OpVideo              Opcode = 12
	OpClientDisconnect   Opcode = 13
	OpSessionUpdate      Opcode = 14
	OpClientFlags        Opcode = 18
	OpClientPlatform     Opcode = 20
)

const GatewayVersion = 8

type SpeakingFlags int

const (
	Voice      SpeakingFlags = 1 << 0
	Soundshare SpeakingFlags = 1 << 1
	Priority   SpeakingFlags = 1 << 2
)

type packet struct {
	Op  Opcode          `json:"op"`
	D   json.RawMessage `json:"d"`
	Seq *int            `json:"seq,omitempty"`
}

type outbound struct {
	Op Opcode `json:"op"`
	D  any    `json:"d"`
}

type identify struct {
	ServerID        snowflake.ID `json:"server_id"`
	ChannelID       snowflake.ID `json:"channel_id"`
	UserID          snowflake.ID `json:"user_id"`
	SessionID       string       `json:"session_id"`
	Token           string       `json:"token"`
	Video           bool         `json:"video"`
	Streams         []streamInfo `json:"streams,omitempty"`
	MaxDAVEProtocol int          `json:"max_dave_protocol_version"`
}

type resolution struct {
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type streamInfo struct {
	Type          string      `json:"type"`
	RID           string      `json:"rid"`
	Quality       int         `json:"quality"`
	Active        bool        `json:"active"`
	MaxBitrate    int         `json:"max_bitrate,omitempty"`
	MaxFramerate  int         `json:"max_framerate,omitempty"`
	MaxResolution *resolution `json:"max_resolution,omitempty"`
	SSRC          uint32      `json:"ssrc,omitempty"`
	RTXSSRC       uint32      `json:"rtx_ssrc,omitempty"`
}

type videoState struct {
	AudioSSRC uint32       `json:"audio_ssrc"`
	VideoSSRC uint32       `json:"video_ssrc"`
	RTXSSRC   uint32       `json:"rtx_ssrc"`
	Streams   []streamInfo `json:"streams"`
	UserID    snowflake.ID `json:"user_id,omitempty"`
}

type resume struct {
	ServerID  snowflake.ID `json:"server_id"`
	ChannelID snowflake.ID `json:"channel_id"`
	SessionID string       `json:"session_id"`
	Token     string       `json:"token"`
	SeqAck    int          `json:"seq_ack"`
}

type hello struct {
	Version  int     `json:"v"`
	Interval float64 `json:"heartbeat_interval"`
}

type heartbeat struct {
	T      int64 `json:"t"`
	SeqAck int   `json:"seq_ack"`
}

type ready struct {
	SSRC    uint32       `json:"ssrc"`
	IP      string       `json:"ip"`
	Port    int          `json:"port"`
	Modes   []string     `json:"modes"`
	Streams []streamInfo `json:"streams"`
}

type protocolData struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Mode    string `json:"mode"`
}

type codec struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	Priority       int    `json:"priority"`
	PayloadType    int    `json:"payload_type"`
	RTXPayloadType int    `json:"rtx_payload_type,omitempty"`
	Encode         *bool  `json:"encode,omitempty"`
	Decode         *bool  `json:"decode,omitempty"`
}

type selectProtocol struct {
	Protocol        string       `json:"protocol"`
	Data            protocolData `json:"data"`
	Codecs          []codec      `json:"codecs"`
	RTCConnectionID string       `json:"rtc_connection_id"`
}

type sessionDescription struct {
	AudioCodec     string `json:"audio_codec"`
	VideoCodec     string `json:"video_codec"`
	MediaSessionID string `json:"media_session_id"`
	Mode           string `json:"mode"`
	SecretKey      []int  `json:"secret_key"`
	DAVEProtocol   int    `json:"dave_protocol_version"`
}

type speaking struct {
	Speaking SpeakingFlags `json:"speaking"`
	Delay    int           `json:"delay"`
	SSRC     uint32        `json:"ssrc"`
	UserID   snowflake.ID  `json:"user_id,omitempty"`
}

type clientsConnect struct {
	UserIDs []snowflake.ID `json:"user_ids"`
}

type clientDisconnect struct {
	UserID snowflake.ID `json:"user_id"`
}

type daveTransition struct {
	ProtocolVersion int    `json:"protocol_version"`
	TransitionID    uint16 `json:"transition_id"`
	Epoch           uint64 `json:"epoch,omitempty"`
}

type transitionReady struct {
	TransitionID uint16 `json:"transition_id"`
}

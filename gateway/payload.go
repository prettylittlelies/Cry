package gateway

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type inbound struct {
	Op Opcode          `json:"op"`
	D  json.RawMessage `json:"d"`
	S  *int64          `json:"s"`
	T  string          `json:"t"`
}

type outbound struct {
	Op Opcode `json:"op"`
	D  any    `json:"d"`
}

type Event struct {
	Type string
	Seq  int64
	Data json.RawMessage
}

type hello struct {
	HeartbeatInterval int64 `json:"heartbeat_interval"`
}

type ready struct {
	SessionID        string `json:"session_id"`
	ResumeGatewayURL string `json:"resume_gateway_url"`
	QoSToken         string `json:"qos_token"`
}

type heartbeat struct {
	Seq int64 `json:"seq"`
	QoS *QoS  `json:"qos"`
}

type resume struct {
	Token     string `json:"token"`
	SessionID string `json:"session_id"`
	Seq       int64  `json:"seq"`
}

type timeSpentSession struct {
	InitializationTimestamp int64  `json:"initialization_timestamp"`
	SessionID               string `json:"session_id"`
	ClientLaunchID          string `json:"client_launch_id"`
}

const timeSpentIdle = 30 * time.Minute

type timeSpent struct {
	id       string
	created  time.Time
	lastUsed time.Time
	sent     string
}

func (t *timeSpent) touch(now time.Time) {
	if t.id == "" || now.Sub(t.lastUsed) >= timeSpentIdle {
		t.id, t.created = uuid.NewString(), now
	}
	t.lastUsed = now
}

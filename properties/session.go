package properties

import (
	"encoding/base64"
	"encoding/json"
	"sync"
	"time"

	"github.com/google/uuid"
)

const HeartbeatRotation = 30 * time.Minute

type Session struct {
	mu        sync.Mutex
	props     Desktop
	rotatedAt time.Time
	encoded   string
	now       func() time.Time
}

func NewSession(d *Desktop) *Session {
	s := &Session{props: *d, now: time.Now}
	s.props.ClientAppState = Focused
	s.rotate()
	return s
}

func (s *Session) rotate() {
	s.props.ClientHeartbeatSessionID = uuid.NewString()
	s.rotatedAt = s.now()
	s.encoded = ""
}

func (s *Session) current() *Desktop {
	if s.now().Sub(s.rotatedAt) >= HeartbeatRotation {
		s.rotate()
	}
	return &s.props
}

func (s *Session) Properties() Desktop {
	s.mu.Lock()
	defer s.mu.Unlock()
	return *s.current()
}

func (s *Session) Encoded() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := s.current()
	if s.encoded == "" {
		raw, _ := json.Marshal(p)
		s.encoded = base64.StdEncoding.EncodeToString(raw)
	}
	return s.encoded
}

func (s *Session) SetAppState(state AppState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.props.ClientAppState != state {
		s.props.ClientAppState = state
		s.encoded = ""
	}
}

func (s *Session) UserAgent() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.props.BrowserUserAgent
}

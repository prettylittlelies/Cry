package voice

import (
	"encoding/json"
	"sync"

	"github.com/prettylittlelies/cry/mls"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/voice/dave"
)

const (
	opPrepareTransition    Opcode = Opcode(dave.OpPrepareTransition)
	opExecuteTransition    Opcode = Opcode(dave.OpExecuteTransition)
	opTransitionReady      Opcode = Opcode(dave.OpTransitionReady)
	opPrepareEpoch         Opcode = Opcode(dave.OpPrepareEpoch)
	opInvalidCommitWelcome Opcode = Opcode(dave.OpInvalidCommitWelcome)
)

type e2ee struct {
	conn    *Conn
	session *dave.Session

	mu          sync.Mutex
	version     int
	external    *mls.ExternalSender
	users       map[snowflake.ID]bool
	transitions map[uint16]int
	encryptor   dave.Encryptor
	decryptors  map[snowflake.ID]*dave.Decryptor
}

func newE2EE(c *Conn, s *dave.Session) *e2ee {
	return &e2ee{conn: c, session: s, users: map[snowflake.ID]bool{}, transitions: map[uint16]int{}, decryptors: map[snowflake.ID]*dave.Decryptor{}}
}

func (e *e2ee) recognized(user snowflake.ID) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.users[user]
}

func (e *e2ee) connected(users ...snowflake.ID) {
	e.mu.Lock()
	for _, u := range users {
		e.users[u] = true
	}
	e.mu.Unlock()
}

func (e *e2ee) disconnected(user snowflake.ID) {
	e.mu.Lock()
	delete(e.users, user)
	delete(e.decryptors, user)
	e.mu.Unlock()
}

func (e *e2ee) active() bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.version > 0 && e.encryptor.Active()
}

func (e *e2ee) start(version int) error {
	if version == 0 {
		return nil
	}
	return e.reset(version)
}

func (e *e2ee) reset(version int) error {
	e.session.Reset(version, e.conn.group())
	e.mu.Lock()
	external := e.external
	e.mu.Unlock()
	if external != nil {
		if err := e.session.SetExternalSender(external); err != nil {
			return err
		}
	}
	kp, err := e.session.KeyPackage()
	if err != nil {
		return err
	}
	return e.conn.ws.Load().sendBinary(uint8(dave.OpKeyPackage), kp)
}

func (e *e2ee) onJSON(op Opcode, d json.RawMessage) error {
	var t daveTransition
	switch op {
	case opPrepareTransition, opExecuteTransition, opPrepareEpoch:
		if err := json.Unmarshal(d, &t); err != nil {
			return err
		}
	default:
		return nil
	}
	switch op {
	case opPrepareTransition:
		e.mu.Lock()
		e.transitions[t.TransitionID] = t.ProtocolVersion
		e.mu.Unlock()
		if t.TransitionID == 0 {
			e.execute(0)
			return nil
		}
		return e.ready(t.TransitionID)
	case opExecuteTransition:
		e.execute(t.TransitionID)
	case opPrepareEpoch:
		e.mu.Lock()
		e.transitions[t.TransitionID] = t.ProtocolVersion
		e.mu.Unlock()
		if t.Epoch == 1 {
			return e.reset(t.ProtocolVersion)
		}
	}
	return nil
}

func (e *e2ee) ready(transition uint16) error {
	return e.conn.ws.Load().send(opTransitionReady, transitionReady{TransitionID: transition})
}

func (e *e2ee) onBinary(op dave.Opcode, payload []byte) error {
	switch op {
	case dave.OpExternalSender:
		ext, err := dave.ParseExternalSender(payload)
		if err != nil {
			return err
		}
		e.mu.Lock()
		e.external = ext
		e.mu.Unlock()
		return e.session.SetExternalSender(ext)
	case dave.OpProposals:
		p, err := dave.ParseProposals(payload)
		if err != nil {
			return err
		}
		out, err := e.session.ProcessProposals(p, e.recognized)
		if err != nil || out == nil {
			return nil
		}
		return e.conn.ws.Load().sendBinary(uint8(dave.OpCommitWelcome), out)
	case dave.OpAnnounceCommit:
		transition, commit, err := dave.ParseTransition(payload)
		if err != nil {
			return err
		}
		return e.settle(transition, e.session.ProcessCommit(commit))
	case dave.OpWelcome:
		transition, welcome, err := dave.ParseTransition(payload)
		if err != nil {
			return err
		}
		return e.settle(transition, e.session.ProcessWelcome(welcome, e.recognized))
	}
	return nil
}

func (e *e2ee) settle(transition uint16, err error) error {
	if err != nil {
		if sendErr := e.conn.ws.Load().send(opInvalidCommitWelcome, transitionReady{TransitionID: transition}); sendErr != nil {
			return sendErr
		}
		e.mu.Lock()
		version := e.version
		e.mu.Unlock()
		if version == 0 {
			version = dave.ProtocolVersion
		}
		return e.reset(version)
	}
	e.refreshDecryptors()
	e.mu.Lock()
	e.transitions[transition] = dave.ProtocolVersion
	e.mu.Unlock()
	if transition == 0 {
		e.execute(0)
		return nil
	}
	return e.ready(transition)
}

func (e *e2ee) refreshDecryptors() {
	for _, user := range e.session.Members() {
		if user == e.conn.self {
			continue
		}
		r, err := e.session.Ratchet(user)
		if err != nil {
			continue
		}
		e.mu.Lock()
		d := e.decryptors[user]
		if d == nil {
			d = new(dave.Decryptor)
			e.decryptors[user] = d
		}
		e.mu.Unlock()
		d.SetRatchet(r)
	}
}

func (e *e2ee) execute(transition uint16) {
	e.mu.Lock()
	version, ok := e.transitions[transition]
	delete(e.transitions, transition)
	if ok {
		e.version = version
	}
	e.mu.Unlock()
	if !ok {
		return
	}
	if version == 0 {
		e.encryptor.SetRatchet(nil)
		return
	}
	if r, err := e.session.Ratchet(e.conn.self); err == nil {
		e.encryptor.SetRatchet(r)
	}
}

func (e *e2ee) protect(frame []byte) ([]byte, error) {
	e.mu.Lock()
	version := e.version
	e.mu.Unlock()
	if version == 0 || !e.encryptor.Active() {
		return frame, nil
	}
	return e.encryptor.Encrypt(frame)
}

func (e *e2ee) protectRanges(frame []byte, clear []dave.Range) ([]byte, error) {
	e.mu.Lock()
	version := e.version
	e.mu.Unlock()
	if version == 0 || !e.encryptor.Active() {
		return frame, nil
	}
	return e.encryptor.EncryptRanges(frame, clear)
}

func (e *e2ee) unprotect(user snowflake.ID, frame []byte) ([]byte, error) {
	e.mu.Lock()
	d := e.decryptors[user]
	e.mu.Unlock()
	if d == nil {
		return frame, nil
	}
	out, err := d.Decrypt(frame)
	if err == dave.ErrNotProtocolFrame {
		return frame, nil
	}
	return out, err
}

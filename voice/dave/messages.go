package dave

import (
	"errors"
	"fmt"

	"github.com/prettylittlelies/cry/mls"
)

type Opcode uint8

const (
	OpPrepareTransition    Opcode = 21
	OpExecuteTransition    Opcode = 22
	OpTransitionReady      Opcode = 23
	OpPrepareEpoch         Opcode = 24
	OpExternalSender       Opcode = 25
	OpKeyPackage           Opcode = 26
	OpProposals            Opcode = 27
	OpCommitWelcome        Opcode = 28
	OpAnnounceCommit       Opcode = 29
	OpWelcome              Opcode = 30
	OpInvalidCommitWelcome Opcode = 31
)

type ProposalsOperation uint8

const (
	ProposalsAppend ProposalsOperation = 0
	ProposalsRevoke ProposalsOperation = 1
)

var ErrMalformed = errors.New("dave: malformed binary message")

type Proposals struct {
	Operation ProposalsOperation
	Messages  []*mls.Message
	Refs      [][]byte
}

func ParseProposals(payload []byte) (*Proposals, error) {
	if len(payload) < 1 {
		return nil, ErrMalformed
	}
	p := &Proposals{Operation: ProposalsOperation(payload[0])}
	r := mls.NewReader(payload[1:])
	switch p.Operation {
	case ProposalsAppend:
		r.Vector(func(r *mls.Reader) {
			m := new(mls.Message)
			r.Decode(m)
			p.Messages = append(p.Messages, m)
		})
	case ProposalsRevoke:
		r.Vector(func(r *mls.Reader) { p.Refs = append(p.Refs, r.Opaque()) })
	default:
		return nil, fmt.Errorf("%w: proposals operation %d", ErrMalformed, p.Operation)
	}
	if r.Err() != nil || !r.Empty() {
		return nil, fmt.Errorf("%w: proposals: %v", ErrMalformed, r.Err())
	}
	return p, nil
}

func ParseExternalSender(payload []byte) (*mls.ExternalSender, error) {
	var e mls.ExternalSender
	if err := mls.Unmarshal(payload, &e); err != nil {
		return nil, fmt.Errorf("%w: external sender: %w", ErrMalformed, err)
	}
	return &e, nil
}

func ParseTransition(payload []byte) (uint16, []byte, error) {
	if len(payload) < 2 {
		return 0, nil, ErrMalformed
	}
	return uint16(payload[0])<<8 | uint16(payload[1]), payload[2:], nil
}

func EncodeKeyPackage(kp *mls.KeyPackage) []byte {
	return append([]byte{byte(OpKeyPackage)}, mls.Marshal(&mls.Message{WireFormat: mls.WireKeyPackage, KeyPackage: kp})...)
}

func EncodeCommitWelcome(commit *mls.Message, welcome *mls.Welcome) []byte {
	var w mls.Writer
	w.U8(uint8(OpCommitWelcome))
	w.Encode(commit)
	if welcome != nil {
		w.Encode(welcome)
	}
	return w.Bytes()
}

func EncodeProposals(p *Proposals) []byte {
	var w mls.Writer
	w.U8(uint8(p.Operation))
	w.Vector(func(w *mls.Writer) {
		for _, m := range p.Messages {
			w.Encode(m)
		}
		for _, ref := range p.Refs {
			w.Opaque(ref)
		}
	})
	return w.Bytes()
}

func ParseCommitWelcome(payload []byte) (*mls.Message, *mls.Welcome, error) {
	r := mls.NewReader(payload)
	commit := new(mls.Message)
	r.Decode(commit)
	var welcome *mls.Welcome
	if r.Err() == nil && !r.Empty() {
		welcome = new(mls.Welcome)
		r.Decode(welcome)
	}
	if r.Err() != nil || !r.Empty() {
		return nil, nil, fmt.Errorf("%w: commit welcome: %v", ErrMalformed, r.Err())
	}
	return commit, welcome, nil
}

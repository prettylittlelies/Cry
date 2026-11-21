package mls

import "fmt"

type PSKStore map[string][]byte

func (s *CipherSuite) PSKSecret(ids []PreSharedKeyID, values [][]byte) []byte {
	secret := make([]byte, s.SecretSize())
	for i, id := range ids {
		var w Writer
		id.encode(&w)
		w.U16(uint16(i))
		w.U16(uint16(len(ids)))
		input := s.ExpandWithLabel(s.Extract(nil, values[i]), "derived psk", w.buf, s.SecretSize())
		secret = s.Extract(input, secret)
	}
	return secret
}

const resumptionHistory = 16

type pskSource struct {
	external   PSKStore
	groupID    []byte
	resumption map[uint64][]byte
}

func (src pskSource) resolve(s *CipherSuite, ids []PreSharedKeyID) ([]byte, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	values := make([][]byte, len(ids))
	for i, id := range ids {
		var v []byte
		var ok bool
		switch id.Type {
		case PSKExternal:
			v, ok = src.external[string(id.ID)]
		case PSKResumption:
			if string(id.GroupID) == string(src.groupID) {
				v, ok = src.resumption[id.Epoch]
			}
		}
		if !ok {
			return nil, fmt.Errorf("%w: unknown psk %x epoch %d", ErrInvalidProposal, id.ID, id.Epoch)
		}
		values[i] = v
	}
	return s.PSKSecret(ids, values), nil
}

func (g *Group) pskSource() pskSource {
	return pskSource{external: g.psks, groupID: g.context.GroupID, resumption: g.resumption}
}

func (g *Group) rememberResumption() {
	next := make(map[uint64][]byte, resumptionHistory)
	for epoch, secret := range g.resumption {
		if epoch+resumptionHistory > g.context.Epoch {
			next[epoch] = secret
		}
	}
	next[g.context.Epoch] = g.secrets.Resumption
	g.resumption = next
}

func pskIDs(proposals []CachedProposal) []PreSharedKeyID {
	var out []PreSharedKeyID
	for _, p := range proposals {
		if p.Proposal.Type == ProposalPSK {
			out = append(out, *p.Proposal.PSK)
		}
	}
	return out
}

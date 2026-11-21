package mls

import (
	"bytes"
	"crypto/hmac"
	"fmt"
)

func (g *Group) senderKey(sender Sender) ([]byte, error) {
	switch sender.Type {
	case SenderMember:
		leaf := g.tree.Leaf(LeafIndex(sender.Index))
		if leaf == nil {
			return nil, fmt.Errorf("%w: leaf %d", ErrUnknownSender, sender.Index)
		}
		return leaf.SignatureKey, nil
	case SenderExternal:
		senders, err := g.ExternalSenders()
		if err != nil {
			return nil, err
		}
		if int(sender.Index) >= len(senders) {
			return nil, fmt.Errorf("%w: external sender %d", ErrUnknownSender, sender.Index)
		}
		return senders[sender.Index].SignatureKey, nil
	}
	return nil, fmt.Errorf("%w: sender type %d", ErrUnsupportedFeature, sender.Type)
}

func (g *Group) Verify(m *Message) (*AuthenticatedContent, error) {
	if m.WireFormat != WirePublicMessage {
		return nil, fmt.Errorf("%w: wire format %d", ErrUnsupportedFeature, m.WireFormat)
	}
	pm := m.Public
	if !bytes.Equal(pm.Content.GroupID, g.context.GroupID) {
		return nil, ErrWrongGroup
	}
	if pm.Content.Epoch != g.context.Epoch {
		return nil, fmt.Errorf("%w: %d, at %d", ErrWrongEpoch, pm.Content.Epoch, g.context.Epoch)
	}
	if pm.Content.Type == ContentApplication {
		return nil, fmt.Errorf("%w: public application message", ErrUnsupportedFeature)
	}
	if pm.Content.Sender.Type == SenderExternal && pm.Content.Type != ContentProposal {
		return nil, fmt.Errorf("%w: external sender may only propose", ErrInvalidProposal)
	}
	ac := pm.Authenticated()
	key, err := g.senderKey(pm.Content.Sender)
	if err != nil {
		return nil, err
	}
	if err := ac.Verify(g.suite, key, &g.context); err != nil {
		return nil, err
	}
	if pm.Content.Sender.Type == SenderMember {
		want := g.suite.MAC(g.secrets.Membership, membershipInput(&ac, &g.context))
		if !hmac.Equal(want, pm.MembershipTag) {
			return nil, ErrMembershipTag
		}
	}
	return &ac, nil
}

func (g *Group) frame(content FramedContent) (*AuthenticatedContent, error) {
	ac := &AuthenticatedContent{WireFormat: WirePublicMessage, Content: content}
	if err := ac.Sign(g.suite, g.signer, &g.context); err != nil {
		return nil, err
	}
	return ac, nil
}

func (g *Group) publish(ac *AuthenticatedContent) *Message {
	pm := &PublicMessage{Content: ac.Content, Auth: ac.Auth}
	if ac.Content.Sender.Type == SenderMember {
		pm.MembershipTag = g.suite.MAC(g.secrets.Membership, membershipInput(ac, &g.context))
	}
	return &Message{WireFormat: WirePublicMessage, Public: pm}
}

func (g *Group) CacheProposal(ac *AuthenticatedContent) ([]byte, error) {
	if ac.Content.Type != ContentProposal {
		return nil, fmt.Errorf("%w: content type %d", ErrInvalidProposal, ac.Content.Type)
	}
	ref := ac.ProposalRef(g.suite)
	for _, p := range g.proposals {
		if bytes.Equal(p.Ref, ref) {
			return nil, fmt.Errorf("%w: duplicate reference", ErrInvalidProposal)
		}
	}
	if err := g.validateProposal(ac.Content.Proposal, ac.Content.Sender); err != nil {
		return nil, err
	}
	g.proposals = append(g.proposals, CachedProposal{Ref: ref, Proposal: *ac.Content.Proposal, Sender: ac.Content.Sender})
	return ref, nil
}

func (g *Group) RevokeProposal(ref []byte) bool {
	for i, p := range g.proposals {
		if bytes.Equal(p.Ref, ref) {
			g.proposals = append(g.proposals[:i:i], g.proposals[i+1:]...)
			return true
		}
	}
	return false
}

func (g *Group) validateProposal(p *Proposal, sender Sender) error {
	switch p.Type {
	case ProposalAdd:
		if err := p.Add.Verify(g.suite); err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidProposal, err)
		}
	case ProposalRemove:
		if g.tree.Leaf(p.Remove) == nil {
			return fmt.Errorf("%w: remove of blank leaf %d", ErrInvalidProposal, p.Remove)
		}
	case ProposalUpdate:
		if sender.Type != SenderMember {
			return fmt.Errorf("%w: update from non-member", ErrInvalidProposal)
		}
		if p.Update.Source != SourceUpdate {
			return fmt.Errorf("%w: update leaf source %d", ErrInvalidProposal, p.Update.Source)
		}
		if err := p.Update.Verify(g.suite, g.context.GroupID, LeafIndex(sender.Index)); err != nil {
			return fmt.Errorf("%w: %w", ErrInvalidProposal, err)
		}
	case ProposalPSK:
		if _, err := g.pskSource().resolve(g.suite, []PreSharedKeyID{*p.PSK}); err != nil {
			return err
		}
	case ProposalGroupContextExtensions:
	default:
		return fmt.Errorf("%w: proposal type %d", ErrUnsupportedFeature, p.Type)
	}
	return nil
}

func ExternalProposal(s *CipherSuite, signer *SignaturePrivateKey, senderIndex uint32, groupID []byte, epoch uint64, p Proposal) (*Message, error) {
	ac := &AuthenticatedContent{WireFormat: WirePublicMessage, Content: FramedContent{
		GroupID:  groupID,
		Epoch:    epoch,
		Sender:   Sender{Type: SenderExternal, Index: senderIndex},
		Type:     ContentProposal,
		Proposal: &p,
	}}
	if err := ac.Sign(s, signer, nil); err != nil {
		return nil, err
	}
	return &Message{WireFormat: WirePublicMessage, Public: &PublicMessage{Content: ac.Content, Auth: ac.Auth}}, nil
}

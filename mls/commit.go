package mls

import (
	"bytes"
	"crypto/hmac"
	"fmt"
	"slices"
)

type CommitOptions struct {
	InlineTree bool
	ForcePath  bool
}

type CommitResult struct {
	Commit  *Message
	Welcome *Welcome
	Next    *Group
}

type resolvedProposals struct {
	proposals []CachedProposal
	joiners   []LeafIndex
	added     []*KeyPackage
}

func (g *Group) resolve(commit *Commit, committer Sender) ([]CachedProposal, error) {
	out := make([]CachedProposal, 0, len(commit.Proposals))
	for _, por := range commit.Proposals {
		if por.Type == ProposalValue {
			if err := g.validateProposal(por.Proposal, committer); err != nil {
				return nil, err
			}
			out = append(out, CachedProposal{Proposal: *por.Proposal, Sender: committer})
			continue
		}
		i := slices.IndexFunc(g.proposals, func(p CachedProposal) bool { return bytes.Equal(p.Ref, por.Ref) })
		if i < 0 {
			return nil, ErrUnknownProposal
		}
		out = append(out, g.proposals[i])
	}
	return out, nil
}

func pathRequired(proposals []CachedProposal) bool {
	if len(proposals) == 0 {
		return true
	}
	return slices.ContainsFunc(proposals, func(p CachedProposal) bool {
		switch p.Proposal.Type {
		case ProposalAdd, ProposalPSK, ProposalReInit:
			return false
		}
		return true
	})
}

func (g *Group) applyProposals(proposals []CachedProposal, committer LeafIndex) (*resolvedProposals, error) {
	res := &resolvedProposals{proposals: proposals}
	removed := map[LeafIndex]bool{}
	for _, kind := range []ProposalType{ProposalGroupContextExtensions, ProposalUpdate, ProposalRemove, ProposalAdd} {
		for _, p := range proposals {
			if p.Proposal.Type != kind {
				continue
			}
			switch kind {
			case ProposalGroupContextExtensions:
				g.context.Extensions = slices.Clone(p.Proposal.Extensions)
			case ProposalUpdate:
				sender := LeafIndex(p.Sender.Index)
				if sender == committer {
					return nil, fmt.Errorf("%w: committer updates itself by proposal", ErrInvalidCommit)
				}
				g.tree.UpdateLeaf(sender, p.Proposal.Update.Clone())
			case ProposalRemove:
				if p.Proposal.Remove == committer && p.Sender.Type == SenderMember {
					return nil, fmt.Errorf("%w: committer removes itself", ErrInvalidCommit)
				}
				if removed[p.Proposal.Remove] || g.tree.Leaf(p.Proposal.Remove) == nil {
					return nil, fmt.Errorf("%w: remove of leaf %d", ErrInvalidCommit, p.Proposal.Remove)
				}
				removed[p.Proposal.Remove] = true
				g.tree.RemoveLeaf(p.Proposal.Remove)
			case ProposalAdd:
				kp := p.Proposal.Add
				for _, leaf := range g.tree.Members {
					if bytes.Equal(leaf.SignatureKey, kp.LeafNode.SignatureKey) || bytes.Equal(leaf.EncryptionKey, kp.LeafNode.EncryptionKey) {
						return nil, fmt.Errorf("%w: duplicate key in add", ErrInvalidCommit)
					}
				}
				res.joiners = append(res.joiners, g.tree.AddLeaf(kp.LeafNode.Clone()))
				res.added = append(res.added, kp)
			}
		}
	}
	g.own.prune(g.tree)
	return res, nil
}

func (g *Group) provisional(next *Group) GroupContextFunc {
	return func(treeHash []byte) *GroupContext {
		ctx := next.context
		ctx.Epoch = g.context.Epoch + 1
		ctx.TreeHash = treeHash
		ctx.ConfirmedTranscriptHash = g.context.ConfirmedTranscriptHash
		return &ctx
	}
}

func (g *Group) advance(next *Group, ac *AuthenticatedContent, commitSecret, psk []byte) []byte {
	s := g.suite
	next.context.Epoch = g.context.Epoch + 1
	next.context.TreeHash = next.tree.TreeHash(s)
	next.context.ConfirmedTranscriptHash = s.ConfirmedTranscriptHash(g.interim, ac)
	joiner := s.JoinerSecret(g.secrets.Init, commitSecret, &next.context)
	next.secrets = s.EpochSecrets(joiner, psk, &next.context)
	next.proposals = nil
	next.rememberResumption()
	return joiner
}

func (g *Group) Commit(opts CommitOptions) (*CommitResult, error) {
	s := g.suite
	next := g.Clone()
	self := Sender{Type: SenderMember, Index: uint32(g.own.Leaf)}
	refs := make([]ProposalOrRef, len(g.proposals))
	for i, p := range g.proposals {
		refs[i] = ProposalOrRef{Type: ProposalReference, Ref: p.Ref}
	}
	applied, err := next.applyProposals(g.proposals, g.own.Leaf)
	if err != nil {
		return nil, err
	}
	commit := &Commit{Proposals: refs}
	commitSecret := make([]byte, s.SecretSize())
	var pathSecrets map[NodeIndex][]byte
	if opts.ForcePath || pathRequired(g.proposals) {
		leaf := next.tree.Leaf(g.own.Leaf).Clone()
		res, err := next.tree.Encap(s, next.own, leaf, g.signer, g.context.GroupID, g.provisional(next), applied.joiners)
		if err != nil {
			return nil, err
		}
		next.own = res.Secrets
		commit.Path, commitSecret, pathSecrets = &res.Path, res.CommitSecret, res.PathSecrets
	}
	ac, err := g.frame(FramedContent{GroupID: g.context.GroupID, Epoch: g.context.Epoch, Sender: self, Type: ContentCommit, Commit: commit})
	if err != nil {
		return nil, err
	}
	psks := pskIDs(g.proposals)
	psk, err := g.pskSource().resolve(s, psks)
	if err != nil {
		return nil, err
	}
	joiner := g.advance(next, ac, commitSecret, psk)
	ac.Auth.ConfirmationTag = s.MAC(next.secrets.Confirmation, next.context.ConfirmedTranscriptHash)
	next.interim = s.InterimTranscriptHash(next.context.ConfirmedTranscriptHash, ac.Auth.ConfirmationTag)
	result := &CommitResult{Commit: g.publish(ac), Next: next}
	if len(applied.joiners) > 0 {
		if result.Welcome, err = next.welcome(joiner, psk, psks, ac.Auth.ConfirmationTag, applied, pathSecrets, opts.InlineTree); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (g *Group) HandleCommit(ac *AuthenticatedContent) (*Group, error) {
	s := g.suite
	if ac.Content.Type != ContentCommit || ac.Content.Sender.Type != SenderMember {
		return nil, fmt.Errorf("%w: not a member commit", ErrInvalidCommit)
	}
	committer := LeafIndex(ac.Content.Sender.Index)
	if committer == g.own.Leaf {
		return nil, fmt.Errorf("%w: own commit must be adopted, not processed", ErrInvalidCommit)
	}
	commit := ac.Content.Commit
	proposals, err := g.resolve(commit, ac.Content.Sender)
	if err != nil {
		return nil, err
	}
	next := g.Clone()
	applied, err := next.applyProposals(proposals, committer)
	if err != nil {
		return nil, err
	}
	if slices.ContainsFunc(applied.proposals, func(p CachedProposal) bool {
		return p.Proposal.Type == ProposalRemove && p.Proposal.Remove == g.own.Leaf
	}) {
		return nil, fmt.Errorf("%w: removed from group", ErrInvalidCommit)
	}
	commitSecret := make([]byte, s.SecretSize())
	switch {
	case commit.Path != nil:
		if commitSecret, err = next.tree.Decap(s, next.own, committer, commit.Path, g.context.GroupID, g.provisional(next), applied.joiners); err != nil {
			return nil, err
		}
	case pathRequired(proposals):
		return nil, fmt.Errorf("%w: missing required path", ErrInvalidCommit)
	}
	psk, err := g.pskSource().resolve(s, pskIDs(proposals))
	if err != nil {
		return nil, err
	}
	g.advance(next, ac, commitSecret, psk)
	want := s.MAC(next.secrets.Confirmation, next.context.ConfirmedTranscriptHash)
	if !hmac.Equal(want, ac.Auth.ConfirmationTag) {
		return nil, ErrConfirmationTag
	}
	next.interim = s.InterimTranscriptHash(next.context.ConfirmedTranscriptHash, ac.Auth.ConfirmationTag)
	return next, nil
}

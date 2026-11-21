package mls

import (
	"bytes"
	"crypto/hmac"
	"fmt"
	"slices"
)

func (g *Group) welcome(joiner, psk []byte, psks []PreSharedKeyID, confirmationTag []byte, applied *resolvedProposals, pathSecrets map[NodeIndex][]byte, inlineTree bool) (*Welcome, error) {
	s := g.suite
	info := GroupInfo{GroupContext: g.context, ConfirmationTag: confirmationTag, Signer: g.own.Leaf}
	if inlineTree {
		info.Extensions = Extensions{{Type: ExtensionRatchetTree, Data: Marshal(g.tree)}}
	}
	if err := info.Sign(s, g.signer); err != nil {
		return nil, err
	}
	key, nonce := s.welcomeKeyNonce(s.WelcomeSecret(joiner, psk))
	encrypted, err := s.Seal(key, nonce, nil, Marshal(&info))
	if err != nil {
		return nil, err
	}
	w := &Welcome{CipherSuite: s.ID, EncryptedGroupInfo: encrypted}
	for i, leaf := range applied.joiners {
		secrets := GroupSecrets{JoinerSecret: joiner, PSKs: psks}
		if pathSecrets != nil {
			secrets.PathSecret = pathSecrets[CommonAncestor(leaf.Node(), g.own.Leaf.Node())]
		}
		kp := applied.added[i]
		ct, err := s.EncryptWithLabel(kp.InitKey, "Welcome", encrypted, Marshal(&secrets))
		if err != nil {
			return nil, err
		}
		w.Secrets = append(w.Secrets, EncryptedGroupSecrets{NewMember: kp.Ref(s), Secrets: ct})
	}
	return w, nil
}

type JoinOptions struct {
	Tree *RatchetTree
	PSKs PSKStore
}

func OpenWelcome(s *CipherSuite, w *Welcome, kp *KeyPackageSecrets, psks PSKStore) (*GroupSecrets, *GroupInfo, []byte, error) {
	if w.CipherSuite != s.ID {
		return nil, nil, nil, fmt.Errorf("%w: suite %d", ErrInvalidWelcome, w.CipherSuite)
	}
	ref := kp.Package.Ref(s)
	i := slices.IndexFunc(w.Secrets, func(e EncryptedGroupSecrets) bool { return bytes.Equal(e.NewMember, ref) })
	if i < 0 {
		return nil, nil, nil, fmt.Errorf("%w: no secrets for key package", ErrInvalidWelcome)
	}
	raw, err := s.DecryptWithLabel(kp.Init, "Welcome", w.EncryptedGroupInfo, w.Secrets[i].Secrets)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%w: %w", ErrInvalidWelcome, err)
	}
	var secrets GroupSecrets
	if err := Unmarshal(raw, &secrets); err != nil {
		return nil, nil, nil, fmt.Errorf("%w: %w", ErrInvalidWelcome, err)
	}
	psk, err := pskSource{external: psks}.resolve(s, secrets.PSKs)
	if err != nil {
		return nil, nil, nil, err
	}
	key, nonce := s.welcomeKeyNonce(s.WelcomeSecret(secrets.JoinerSecret, psk))
	if raw, err = s.Open(key, nonce, nil, w.EncryptedGroupInfo); err != nil {
		return nil, nil, nil, fmt.Errorf("%w: %w", ErrInvalidWelcome, err)
	}
	var info GroupInfo
	if err := Unmarshal(raw, &info); err != nil {
		return nil, nil, nil, fmt.Errorf("%w: %w", ErrInvalidWelcome, err)
	}
	if info.GroupContext.CipherSuite != s.ID || info.GroupContext.Version != MLS10 {
		return nil, nil, nil, fmt.Errorf("%w: group context parameters", ErrInvalidWelcome)
	}
	return &secrets, &info, psk, nil
}

func Join(s *CipherSuite, w *Welcome, kp *KeyPackageSecrets, opts JoinOptions) (*Group, error) {
	secrets, info, psk, err := OpenWelcome(s, w, kp, opts.PSKs)
	if err != nil {
		return nil, err
	}
	tree := opts.Tree
	if inline, ok := info.Extensions.Find(ExtensionRatchetTree); ok {
		tree = new(RatchetTree)
		if err := Unmarshal(inline, tree); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidWelcome, err)
		}
	}
	if tree == nil {
		return nil, fmt.Errorf("%w: no ratchet tree", ErrInvalidWelcome)
	}
	tree = tree.Clone()
	if !bytes.Equal(tree.TreeHash(s), info.GroupContext.TreeHash) {
		return nil, fmt.Errorf("%w: tree hash mismatch", ErrInvalidWelcome)
	}
	if err := tree.VerifyLeaves(s, info.GroupContext.GroupID); err != nil {
		return nil, err
	}
	if err := tree.VerifyParentHashes(s); err != nil {
		return nil, err
	}
	signer := tree.Leaf(info.Signer)
	if signer == nil {
		return nil, fmt.Errorf("%w: unknown signer %d", ErrInvalidWelcome, info.Signer)
	}
	if err := info.Verify(s, signer.SignatureKey); err != nil {
		return nil, err
	}
	encoded := Marshal(&kp.Package.LeafNode)
	self, ok := tree.Find(func(l *LeafNode) bool { return bytes.Equal(Marshal(l), encoded) })
	if !ok {
		return nil, fmt.Errorf("%w: own leaf not in tree", ErrInvalidWelcome)
	}
	own := NewTreeSecrets(self, kp.Encryption)
	if secrets.PathSecret != nil {
		if _, err := tree.ImportPathSecret(s, own, CommonAncestor(self.Node(), info.Signer.Node()), secrets.PathSecret); err != nil {
			return nil, err
		}
	}
	g := &Group{suite: s, context: info.GroupContext, tree: tree, own: own, signer: kp.Signer, psks: opts.PSKs}
	g.secrets = s.EpochSecrets(secrets.JoinerSecret, psk, &g.context)
	if !hmac.Equal(s.MAC(g.secrets.Confirmation, g.context.ConfirmedTranscriptHash), info.ConfirmationTag) {
		return nil, ErrConfirmationTag
	}
	g.interim = s.InterimTranscriptHash(g.context.ConfirmedTranscriptHash, info.ConfirmationTag)
	g.rememberResumption()
	return g, nil
}

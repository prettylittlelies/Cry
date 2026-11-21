# mls

```go
import "github.com/prettylittlelies/cry/mls"
```

## Index

- [func Marshal](#func-marshal)
- [func NodeWidth](#func-nodewidth)
- [func OpenWelcome](#func-openwelcome)
- [func Unmarshal](#func-unmarshal)
- [type AuthenticatedContent](#type-authenticatedcontent)
- [type CachedProposal](#type-cachedproposal)
- [type Capabilities](#type-capabilities)
- [type CipherSuite](#type-ciphersuite)
- [type CipherSuiteID](#type-ciphersuiteid)
- [type Commit](#type-commit)
- [type CommitOptions](#type-commitoptions)
- [type CommitResult](#type-commitresult)
- [type ContentType](#type-contenttype)
- [type Credential](#type-credential)
- [type CredentialType](#type-credentialtype)
- [type Decodable](#type-decodable)
- [type Encodable](#type-encodable)
- [type EncryptedGroupSecrets](#type-encryptedgroupsecrets)
- [type EpochSecrets](#type-epochsecrets)
- [type Extension](#type-extension)
- [type ExtensionType](#type-extensiontype)
- [type Extensions](#type-extensions)
- [type ExternalSender](#type-externalsender)
- [type ExternalSenders](#type-externalsenders)
- [type FramedContent](#type-framedcontent)
- [type FramedContentAuthData](#type-framedcontentauthdata)
- [type Group](#type-group)
- [type GroupContext](#type-groupcontext)
- [type GroupContextFunc](#type-groupcontextfunc)
- [type GroupInfo](#type-groupinfo)
- [type GroupSecrets](#type-groupsecrets)
- [type HPKECiphertext](#type-hpkeciphertext)
- [type JoinOptions](#type-joinoptions)
- [type KeyPackage](#type-keypackage)
- [type KeyPackageSecrets](#type-keypackagesecrets)
- [type LeafIndex](#type-leafindex)
- [type LeafNode](#type-leafnode)
- [type LeafNodeSource](#type-leafnodesource)
- [type LeafOptions](#type-leafoptions)
- [type Lifetime](#type-lifetime)
- [type Message](#type-message)
- [type Node](#type-node)
- [type NodeIndex](#type-nodeindex)
- [type NodeType](#type-nodetype)
- [type PSKStore](#type-pskstore)
- [type PSKType](#type-psktype)
- [type ParentNode](#type-parentnode)
- [type PathResult](#type-pathresult)
- [type PreSharedKeyID](#type-presharedkeyid)
- [type PrivateMessage](#type-privatemessage)
- [type Proposal](#type-proposal)
- [type ProposalOrRef](#type-proposalorref)
- [type ProposalOrRefType](#type-proposalorreftype)
- [type ProposalType](#type-proposaltype)
- [type ProtocolVersion](#type-protocolversion)
- [type PublicMessage](#type-publicmessage)
- [type RatchetTree](#type-ratchettree)
- [type ReInit](#type-reinit)
- [type Reader](#type-reader)
- [type Sender](#type-sender)
- [type SenderType](#type-sendertype)
- [type SignaturePrivateKey](#type-signatureprivatekey)
- [type TreeSecrets](#type-treesecrets)
- [type UpdatePath](#type-updatepath)
- [type UpdatePathNode](#type-updatepathnode)
- [type Welcome](#type-welcome)
- [type WireFormat](#type-wireformat)
- [type Writer](#type-writer)

## Variables

```go
var (
	ErrTruncated     = errors.New("mls: truncated input")
	ErrTrailingData  = errors.New("mls: trailing data")
	ErrVarintTooLong = errors.New("mls: varint exceeds 30 bits")
	ErrVarintPadding = errors.New("mls: non-minimal varint")
)
```

```go
var (
	ErrWrongGroup         = errors.New("mls: message for another group")
	ErrWrongEpoch         = errors.New("mls: message for another epoch")
	ErrUnknownSender      = errors.New("mls: unknown sender")
	ErrUnknownProposal    = errors.New("mls: commit references unknown proposal")
	ErrInvalidProposal    = errors.New("mls: invalid proposal")
	ErrInvalidCommit      = errors.New("mls: invalid commit")
	ErrInvalidWelcome     = errors.New("mls: invalid welcome")
	ErrMembershipTag      = errors.New("mls: invalid membership tag")
	ErrConfirmationTag    = errors.New("mls: invalid confirmation tag")
	ErrUnsupportedFeature = errors.New("mls: unsupported feature")
)
```

```go
var ErrInvalidSignature = errors.New("mls: invalid signature")
```

```go
var ErrInvalidTree = errors.New("mls: invalid ratchet tree")
```

```go
var ErrPathMismatch = errors.New("mls: update path inconsistent with tree")
```

```go
var ErrUnsupportedSuite = errors.New("mls: unsupported cipher suite")
```

## Functions

### func Marshal

```go
func Marshal(v encoder) []byte
```

### func NodeWidth

```go
func NodeWidth(leaves uint32) uint32
```

### func OpenWelcome

```go
func OpenWelcome(s *CipherSuite, w *Welcome, kp *KeyPackageSecrets, psks PSKStore) (*GroupSecrets, *GroupInfo, []byte, error)
```

### func Unmarshal

```go
func Unmarshal(b []byte, v decoder) error
```

## Types

### type AuthenticatedContent

```go
type AuthenticatedContent struct {
	WireFormat WireFormat
	Content    FramedContent
	Auth       FramedContentAuthData
}
```

#### func (*AuthenticatedContent) ProposalRef

```go
func (a *AuthenticatedContent) ProposalRef(s *CipherSuite) []byte
```

#### func (*AuthenticatedContent) Sign

```go
func (a *AuthenticatedContent) Sign(s *CipherSuite, key *SignaturePrivateKey, ctx *GroupContext) error
```

#### func (*AuthenticatedContent) Verify

```go
func (a *AuthenticatedContent) Verify(s *CipherSuite, public []byte, ctx *GroupContext) error
```

### type CachedProposal

```go
type CachedProposal struct {
	Ref      []byte
	Proposal Proposal
	Sender   Sender
}
```

### type Capabilities

```go
type Capabilities struct {
	Versions     []ProtocolVersion
	CipherSuites []CipherSuiteID
	Extensions   []ExtensionType
	Proposals    []ProposalType
	Credentials  []CredentialType
}
```

#### func (*Capabilities) SupportsExtension

```go
func (c *Capabilities) SupportsExtension(t ExtensionType) bool
```

#### func (*Capabilities) SupportsProposal

```go
func (c *Capabilities) SupportsProposal(t ProposalType) bool
```

### type CipherSuite

```go
type CipherSuite struct {
	ID CipherSuiteID
	// contains filtered or unexported fields
}
```

#### func Suite

```go
func Suite(id CipherSuiteID) (*CipherSuite, error)
```

#### func (*CipherSuite) ConfirmedTranscriptHash

```go
func (s *CipherSuite) ConfirmedTranscriptHash(interim []byte, content *AuthenticatedContent) []byte
```

#### func (*CipherSuite) DecryptWithLabel

```go
func (s *CipherSuite) DecryptWithLabel(private hpke.PrivateKey, label string, context []byte, ct HPKECiphertext) ([]byte, error)
```

#### func (*CipherSuite) DeriveKeyPair

```go
func (s *CipherSuite) DeriveKeyPair(secret []byte) (hpke.PrivateKey, error)
```

#### func (*CipherSuite) DeriveSecret

```go
func (s *CipherSuite) DeriveSecret(secret []byte, label string) []byte
```

#### func (*CipherSuite) DeriveTreeSecret

```go
func (s *CipherSuite) DeriveTreeSecret(secret []byte, label string, generation uint32, length int) []byte
```

#### func (*CipherSuite) EncryptWithLabel

```go
func (s *CipherSuite) EncryptWithLabel(public []byte, label string, context, plaintext []byte) (HPKECiphertext, error)
```

#### func (*CipherSuite) EpochSecrets

```go
func (s *CipherSuite) EpochSecrets(joiner, psk []byte, ctx *GroupContext) EpochSecrets
```

#### func (*CipherSuite) Expand

```go
func (s *CipherSuite) Expand(prk, info []byte, length int) []byte
```

#### func (*CipherSuite) ExpandWithLabel

```go
func (s *CipherSuite) ExpandWithLabel(secret []byte, label string, context []byte, length int) []byte
```

#### func (*CipherSuite) Export

```go
func (s *CipherSuite) Export(exporter []byte, label string, context []byte, length int) []byte
```

#### func (*CipherSuite) ExternalKey

```go
func (s *CipherSuite) ExternalKey(external []byte) (hpke.PrivateKey, error)
```

#### func (*CipherSuite) Extract

```go
func (s *CipherSuite) Extract(salt, ikm []byte) []byte
```

#### func (*CipherSuite) GenerateHPKEKey

```go
func (s *CipherSuite) GenerateHPKEKey() (hpke.PrivateKey, error)
```

#### func (*CipherSuite) GenerateSignatureKey

```go
func (s *CipherSuite) GenerateSignatureKey() (*SignaturePrivateKey, error)
```

#### func (*CipherSuite) Hash

```go
func (s *CipherSuite) Hash(data []byte) []byte
```

#### func (*CipherSuite) InterimTranscriptHash

```go
func (s *CipherSuite) InterimTranscriptHash(confirmed, confirmationTag []byte) []byte
```

#### func (*CipherSuite) JoinerSecret

```go
func (s *CipherSuite) JoinerSecret(initSecret, commitSecret []byte, ctx *GroupContext) []byte
```

#### func (*CipherSuite) KeySize

```go
func (s *CipherSuite) KeySize() int
```

#### func (*CipherSuite) MAC

```go
func (s *CipherSuite) MAC(key, data []byte) []byte
```

#### func (*CipherSuite) NonceSize

```go
func (s *CipherSuite) NonceSize() int
```

#### func (*CipherSuite) Open

```go
func (s *CipherSuite) Open(key, nonce, aad, ciphertext []byte) ([]byte, error)
```

#### func (*CipherSuite) PSKSecret

```go
func (s *CipherSuite) PSKSecret(ids []PreSharedKeyID, values [][]byte) []byte
```

#### func (*CipherSuite) ParseHPKEPrivateKey

```go
func (s *CipherSuite) ParseHPKEPrivateKey(raw []byte) (hpke.PrivateKey, error)
```

#### func (*CipherSuite) ParseSignaturePrivateKey

```go
func (s *CipherSuite) ParseSignaturePrivateKey(raw []byte) (*SignaturePrivateKey, error)
```

#### func (*CipherSuite) RandomSecret

```go
func (s *CipherSuite) RandomSecret() []byte
```

#### func (*CipherSuite) RefHash

```go
func (s *CipherSuite) RefHash(label string, value []byte) []byte
```

#### func (*CipherSuite) Seal

```go
func (s *CipherSuite) Seal(key, nonce, aad, plaintext []byte) ([]byte, error)
```

#### func (*CipherSuite) SecretSize

```go
func (s *CipherSuite) SecretSize() int
```

#### func (*CipherSuite) SignWithLabel

```go
func (s *CipherSuite) SignWithLabel(key *SignaturePrivateKey, label string, content []byte) ([]byte, error)
```

#### func (*CipherSuite) VerifyWithLabel

```go
func (s *CipherSuite) VerifyWithLabel(public []byte, label string, content, signature []byte) bool
```

#### func (*CipherSuite) WelcomeSecret

```go
func (s *CipherSuite) WelcomeSecret(joiner, psk []byte) []byte
```

### type CipherSuiteID

```go
type CipherSuiteID uint16
```

```go
const P256AES128GCMSHA256P256 CipherSuiteID = 2
```

### type Commit

```go
type Commit struct {
	Proposals []ProposalOrRef
	Path      *UpdatePath
}
```

### type CommitOptions

```go
type CommitOptions struct {
	InlineTree bool
	ForcePath  bool
}
```

### type CommitResult

```go
type CommitResult struct {
	Commit  *Message
	Welcome *Welcome
	Next    *Group
}
```

### type ContentType

```go
type ContentType uint8
```

```go
const (
	ContentApplication ContentType = 1
	ContentProposal    ContentType = 2
	ContentCommit      ContentType = 3
)
```

### type Credential

```go
type Credential struct {
	Type         CredentialType
	Identity     []byte
	Certificates [][]byte
}
```

#### func BasicCredential

```go
func BasicCredential(identity []byte) Credential
```

#### func (*Credential) Equal

```go
func (c *Credential) Equal(o Credential) bool
```

### type CredentialType

```go
type CredentialType uint16
```

```go
const (
	CredentialBasic CredentialType = 1
	CredentialX509  CredentialType = 2
)
```

### type Decodable

```go
type Decodable = decoder
```

### type Encodable

```go
type Encodable = encoder
```

### type EncryptedGroupSecrets

```go
type EncryptedGroupSecrets struct {
	NewMember []byte
	Secrets   HPKECiphertext
}
```

### type EpochSecrets

```go
type EpochSecrets struct {
	SenderData    []byte
	Encryption    []byte
	Exporter      []byte
	External      []byte
	Confirmation  []byte
	Membership    []byte
	Resumption    []byte
	Authenticator []byte
	Init          []byte
}
```

### type Extension

```go
type Extension struct {
	Type ExtensionType
	Data []byte
}
```

### type ExtensionType

```go
type ExtensionType uint16
```

```go
const (
	ExtensionApplicationID        ExtensionType = 1
	ExtensionRatchetTree          ExtensionType = 2
	ExtensionRequiredCapabilities ExtensionType = 3
	ExtensionExternalPub          ExtensionType = 4
	ExtensionExternalSenders      ExtensionType = 5
)
```

### type Extensions

```go
type Extensions []Extension
```

#### func (Extensions) Find

```go
func (es Extensions) Find(t ExtensionType) ([]byte, bool)
```

#### func (Extensions) With

```go
func (es Extensions) With(t ExtensionType, data []byte) Extensions
```

### type ExternalSender

```go
type ExternalSender struct {
	SignatureKey []byte
	Credential   Credential
}
```

### type ExternalSenders

```go
type ExternalSenders []ExternalSender
```

#### func (ExternalSenders) Extension

```go
func (es ExternalSenders) Extension() Extension
```

### type FramedContent

```go
type FramedContent struct {
	GroupID           []byte
	Epoch             uint64
	Sender            Sender
	AuthenticatedData []byte
	Type              ContentType
	Application       []byte
	Proposal          *Proposal
	Commit            *Commit
}
```

### type FramedContentAuthData

```go
type FramedContentAuthData struct {
	Signature       []byte
	ConfirmationTag []byte
}
```

### type Group

```go
type Group struct {
	// contains filtered or unexported fields
}
```

#### func Join

```go
func Join(s *CipherSuite, w *Welcome, kp *KeyPackageSecrets, opts JoinOptions) (*Group, error)
```

#### func NewGroup

```go
func NewGroup(s *CipherSuite, groupID []byte, kp *KeyPackageSecrets, extensions Extensions) (*Group, error)
```

#### func (*Group) AddPSK

```go
func (g *Group) AddPSK(id, secret []byte)
```

#### func (*Group) CacheProposal

```go
func (g *Group) CacheProposal(ac *AuthenticatedContent) ([]byte, error)
```

#### func (*Group) Clone

```go
func (g *Group) Clone() *Group
```

#### func (*Group) Commit

```go
func (g *Group) Commit(opts CommitOptions) (*CommitResult, error)
```

#### func (*Group) Context

```go
func (g *Group) Context() GroupContext
```

#### func (*Group) EncryptionKey

```go
func (g *Group) EncryptionKey(n NodeIndex) (hpke.PrivateKey, bool)
```

#### func (*Group) Epoch

```go
func (g *Group) Epoch() uint64
```

#### func (*Group) EpochAuthenticator

```go
func (g *Group) EpochAuthenticator() []byte
```

#### func (*Group) Export

```go
func (g *Group) Export(label string, context []byte, length int) []byte
```

#### func (*Group) ExternalSenders

```go
func (g *Group) ExternalSenders() (ExternalSenders, error)
```

#### func (*Group) GroupID

```go
func (g *Group) GroupID() []byte
```

#### func (*Group) HandleCommit

```go
func (g *Group) HandleCommit(ac *AuthenticatedContent) (*Group, error)
```

#### func (*Group) Proposals

```go
func (g *Group) Proposals() []CachedProposal
```

#### func (*Group) RevokeProposal

```go
func (g *Group) RevokeProposal(ref []byte) bool
```

#### func (*Group) Self

```go
func (g *Group) Self() LeafIndex
```

#### func (*Group) Signer

```go
func (g *Group) Signer() *SignaturePrivateKey
```

#### func (*Group) Suite

```go
func (g *Group) Suite() *CipherSuite
```

#### func (*Group) Tree

```go
func (g *Group) Tree() *RatchetTree
```

#### func (*Group) Verify

```go
func (g *Group) Verify(m *Message) (*AuthenticatedContent, error)
```

### type GroupContext

```go
type GroupContext struct {
	Version                 ProtocolVersion
	CipherSuite             CipherSuiteID
	GroupID                 []byte
	Epoch                   uint64
	TreeHash                []byte
	ConfirmedTranscriptHash []byte
	Extensions              Extensions
}
```

### type GroupContextFunc

```go
type GroupContextFunc func(treeHash []byte) *GroupContext
```

### type GroupInfo

```go
type GroupInfo struct {
	GroupContext    GroupContext
	Extensions      Extensions
	ConfirmationTag []byte
	Signer          LeafIndex
	Signature       []byte
}
```

#### func (*GroupInfo) Sign

```go
func (g *GroupInfo) Sign(s *CipherSuite, key *SignaturePrivateKey) error
```

#### func (*GroupInfo) Verify

```go
func (g *GroupInfo) Verify(s *CipherSuite, public []byte) error
```

### type GroupSecrets

```go
type GroupSecrets struct {
	JoinerSecret []byte
	PathSecret   []byte
	PSKs         []PreSharedKeyID
}
```

### type HPKECiphertext

```go
type HPKECiphertext struct {
	KEMOutput  []byte
	Ciphertext []byte
}
```

### type JoinOptions

```go
type JoinOptions struct {
	Tree *RatchetTree
	PSKs PSKStore
}
```

### type KeyPackage

```go
type KeyPackage struct {
	Version     ProtocolVersion
	CipherSuite CipherSuiteID
	InitKey     []byte
	LeafNode    LeafNode
	Extensions  Extensions
	Signature   []byte
}
```

#### func (*KeyPackage) Ref

```go
func (k *KeyPackage) Ref(s *CipherSuite) []byte
```

#### func (*KeyPackage) Verify

```go
func (k *KeyPackage) Verify(s *CipherSuite) error
```

### type KeyPackageSecrets

```go
type KeyPackageSecrets struct {
	Package    KeyPackage
	Init       hpke.PrivateKey
	Encryption hpke.PrivateKey
	Signer     *SignaturePrivateKey
}
```

#### func NewKeyPackage

```go
func NewKeyPackage(s *CipherSuite, signer *SignaturePrivateKey, opts LeafOptions) (*KeyPackageSecrets, error)
```

### type LeafIndex

```go
type LeafIndex uint32
```

#### func (LeafIndex) Node

```go
func (l LeafIndex) Node() NodeIndex
```

### type LeafNode

```go
type LeafNode struct {
	EncryptionKey []byte
	SignatureKey  []byte
	Credential    Credential
	Capabilities  Capabilities
	Source        LeafNodeSource
	Lifetime      Lifetime
	ParentHash    []byte
	Extensions    Extensions
	Signature     []byte
}
```

#### func (*LeafNode) Clone

```go
func (l *LeafNode) Clone() LeafNode
```

#### func (*LeafNode) Sign

```go
func (l *LeafNode) Sign(s *CipherSuite, key *SignaturePrivateKey, groupID []byte, leaf LeafIndex) error
```

#### func (*LeafNode) Verify

```go
func (l *LeafNode) Verify(s *CipherSuite, groupID []byte, leaf LeafIndex) error
```

### type LeafNodeSource

```go
type LeafNodeSource uint8
```

```go
const (
	SourceKeyPackage LeafNodeSource = 1
	SourceUpdate     LeafNodeSource = 2
	SourceCommit     LeafNodeSource = 3
)
```

### type LeafOptions

```go
type LeafOptions struct {
	Credential   Credential
	Capabilities Capabilities
	Extensions   Extensions
}
```

### type Lifetime

```go
type Lifetime struct {
	NotBefore uint64
	NotAfter  uint64
}
```

#### func InfiniteLifetime

```go
func InfiniteLifetime() Lifetime
```

### type Message

```go
type Message struct {
	WireFormat WireFormat
	Public     *PublicMessage
	Private    *PrivateMessage
	Welcome    *Welcome
	GroupInfo  *GroupInfo
	KeyPackage *KeyPackage
}
```

#### func ExternalProposal

```go
func ExternalProposal(s *CipherSuite, signer *SignaturePrivateKey, senderIndex uint32, groupID []byte, epoch uint64, p Proposal) (*Message, error)
```

#### func (*Message) Epoch

```go
func (m *Message) Epoch() uint64
```

#### func (*Message) GroupID

```go
func (m *Message) GroupID() []byte
```

### type Node

```go
type Node struct {
	Leaf   *LeafNode
	Parent *ParentNode
}
```

#### func (*Node) EncryptionKey

```go
func (n *Node) EncryptionKey() []byte
```

### type NodeIndex

```go
type NodeIndex uint32
```

#### func CommonAncestor

```go
func CommonAncestor(x, y NodeIndex) NodeIndex
```

#### func Copath

```go
func Copath(n NodeIndex, leaves uint32) []NodeIndex
```

#### func DirectPath

```go
func DirectPath(n NodeIndex, leaves uint32) []NodeIndex
```

#### func Root

```go
func Root(leaves uint32) NodeIndex
```

#### func (NodeIndex) IsLeaf

```go
func (n NodeIndex) IsLeaf() bool
```

#### func (NodeIndex) Leaf

```go
func (n NodeIndex) Leaf() LeafIndex
```

#### func (NodeIndex) Left

```go
func (n NodeIndex) Left() NodeIndex
```

#### func (NodeIndex) Level

```go
func (n NodeIndex) Level() int
```

#### func (NodeIndex) Parent

```go
func (n NodeIndex) Parent() NodeIndex
```

#### func (NodeIndex) Right

```go
func (n NodeIndex) Right() NodeIndex
```

#### func (NodeIndex) Sibling

```go
func (n NodeIndex) Sibling() NodeIndex
```

### type NodeType

```go
type NodeType uint8
```

```go
const (
	NodeLeaf   NodeType = 1
	NodeParent NodeType = 2
)
```

### type PSKStore

```go
type PSKStore map[string][]byte
```

### type PSKType

```go
type PSKType uint8
```

```go
const (
	PSKExternal   PSKType = 1
	PSKResumption PSKType = 2
)
```

### type ParentNode

```go
type ParentNode struct {
	EncryptionKey  []byte
	ParentHash     []byte
	UnmergedLeaves []LeafIndex
}
```

### type PathResult

```go
type PathResult struct {
	Path         UpdatePath
	CommitSecret []byte
	PathSecrets  map[NodeIndex][]byte
	Secrets      *TreeSecrets
}
```

### type PreSharedKeyID

```go
type PreSharedKeyID struct {
	Type    PSKType
	ID      []byte
	Usage   uint8
	GroupID []byte
	Epoch   uint64
	Nonce   []byte
}
```

### type PrivateMessage

```go
type PrivateMessage struct {
	GroupID             []byte
	Epoch               uint64
	Type                ContentType
	AuthenticatedData   []byte
	EncryptedSenderData []byte
	Ciphertext          []byte
}
```

### type Proposal

```go
type Proposal struct {
	Type       ProposalType
	Add        *KeyPackage
	Update     *LeafNode
	Remove     LeafIndex
	PSK        *PreSharedKeyID
	ReInit     *ReInit
	KEMOutput  []byte
	Extensions Extensions
}
```

### type ProposalOrRef

```go
type ProposalOrRef struct {
	Type     ProposalOrRefType
	Proposal *Proposal
	Ref      []byte
}
```

### type ProposalOrRefType

```go
type ProposalOrRefType uint8
```

```go
const (
	ProposalValue     ProposalOrRefType = 1
	ProposalReference ProposalOrRefType = 2
)
```

### type ProposalType

```go
type ProposalType uint16
```

```go
const (
	ProposalAdd                    ProposalType = 1
	ProposalUpdate                 ProposalType = 2
	ProposalRemove                 ProposalType = 3
	ProposalPSK                    ProposalType = 4
	ProposalReInit                 ProposalType = 5
	ProposalExternalInit           ProposalType = 6
	ProposalGroupContextExtensions ProposalType = 7
)
```

### type ProtocolVersion

```go
type ProtocolVersion uint16
```

```go
const MLS10 ProtocolVersion = 1
```

### type PublicMessage

```go
type PublicMessage struct {
	Content       FramedContent
	Auth          FramedContentAuthData
	MembershipTag []byte
}
```

#### func (*PublicMessage) Authenticated

```go
func (m *PublicMessage) Authenticated() AuthenticatedContent
```

### type RatchetTree

```go
type RatchetTree struct {
	// contains filtered or unexported fields
}
```

#### func NewRatchetTree

```go
func NewRatchetTree(first LeafNode) *RatchetTree
```

#### func (*RatchetTree) AddLeaf

```go
func (t *RatchetTree) AddLeaf(leaf LeafNode) LeafIndex
```

#### func (*RatchetTree) Blank

```go
func (t *RatchetTree) Blank(i NodeIndex) bool
```

#### func (*RatchetTree) Clone

```go
func (t *RatchetTree) Clone() *RatchetTree
```

#### func (*RatchetTree) Decap

```go
func (t *RatchetTree) Decap(s *CipherSuite, own *TreeSecrets, sender LeafIndex, path *UpdatePath, groupID []byte, context GroupContextFunc, exclude []LeafIndex) ([]byte, error)
```

#### func (*RatchetTree) Encap

```go
func (t *RatchetTree) Encap(s *CipherSuite, own *TreeSecrets, leaf LeafNode, signer *SignaturePrivateKey, groupID []byte, context GroupContextFunc, exclude []LeafIndex) (*PathResult, error)
```

#### func (*RatchetTree) FilteredDirectPath

```go
func (t *RatchetTree) FilteredDirectPath(l LeafIndex) []NodeIndex
```

#### func (*RatchetTree) Find

```go
func (t *RatchetTree) Find(match func(*LeafNode) bool) (LeafIndex, bool)
```

#### func (*RatchetTree) ImportPathSecret

```go
func (t *RatchetTree) ImportPathSecret(s *CipherSuite, own *TreeSecrets, from NodeIndex, secret []byte) ([]byte, error)
```

#### func (*RatchetTree) Leaf

```go
func (t *RatchetTree) Leaf(l LeafIndex) *LeafNode
```

#### func (*RatchetTree) Leaves

```go
func (t *RatchetTree) Leaves() uint32
```

#### func (*RatchetTree) Members

```go
func (t *RatchetTree) Members(yield func(LeafIndex, *LeafNode) bool)
```

#### func (*RatchetTree) NodeHash

```go
func (t *RatchetTree) NodeHash(s *CipherSuite, i NodeIndex) []byte
```

#### func (*RatchetTree) Parent

```go
func (t *RatchetTree) Parent(i NodeIndex) *ParentNode
```

#### func (*RatchetTree) ParentHash

```go
func (t *RatchetTree) ParentHash(s *CipherSuite, parent NodeIndex, sibling NodeIndex) []byte
```

#### func (*RatchetTree) RemoveLeaf

```go
func (t *RatchetTree) RemoveLeaf(l LeafIndex)
```

#### func (*RatchetTree) Resolution

```go
func (t *RatchetTree) Resolution(i NodeIndex) []NodeIndex
```

#### func (*RatchetTree) TreeHash

```go
func (t *RatchetTree) TreeHash(s *CipherSuite) []byte
```

#### func (*RatchetTree) UpdateLeaf

```go
func (t *RatchetTree) UpdateLeaf(l LeafIndex, leaf LeafNode)
```

#### func (*RatchetTree) VerifyLeaves

```go
func (t *RatchetTree) VerifyLeaves(s *CipherSuite, groupID []byte) error
```

#### func (*RatchetTree) VerifyParentHashes

```go
func (t *RatchetTree) VerifyParentHashes(s *CipherSuite) error
```

### type ReInit

```go
type ReInit struct {
	GroupID     []byte
	Version     ProtocolVersion
	CipherSuite CipherSuiteID
	Extensions  Extensions
}
```

### type Reader

```go
type Reader struct {
	// contains filtered or unexported fields
}
```

#### func NewReader

```go
func NewReader(b []byte) *Reader
```

#### func (*Reader) Bool

```go
func (r *Reader) Bool() bool
```

#### func (*Reader) Decode

```go
func (r *Reader) Decode(v decoder)
```

#### func (*Reader) Empty

```go
func (r *Reader) Empty() bool
```

#### func (*Reader) Err

```go
func (r *Reader) Err() error
```

#### func (*Reader) Fail

```go
func (r *Reader) Fail(err error)
```

#### func (*Reader) Opaque

```go
func (r *Reader) Opaque() []byte
```

#### func (*Reader) Optional

```go
func (r *Reader) Optional(body func(*Reader)) bool
```

#### func (*Reader) Raw

```go
func (r *Reader) Raw(n int) []byte
```

#### func (*Reader) Remaining

```go
func (r *Reader) Remaining() []byte
```

#### func (*Reader) U16

```go
func (r *Reader) U16() uint16
```

#### func (*Reader) U32

```go
func (r *Reader) U32() uint32
```

#### func (*Reader) U64

```go
func (r *Reader) U64() uint64
```

#### func (*Reader) U8

```go
func (r *Reader) U8() uint8
```

#### func (*Reader) Varint

```go
func (r *Reader) Varint() int
```

#### func (*Reader) Vector

```go
func (r *Reader) Vector(body func(*Reader))
```

### type Sender

```go
type Sender struct {
	Type  SenderType
	Index uint32
}
```

### type SenderType

```go
type SenderType uint8
```

```go
const (
	SenderMember            SenderType = 1
	SenderExternal          SenderType = 2
	SenderNewMemberProposal SenderType = 3
	SenderNewMemberCommit   SenderType = 4
)
```

### type SignaturePrivateKey

```go
type SignaturePrivateKey struct {
	// contains filtered or unexported fields
}
```

#### func (*SignaturePrivateKey) Bytes

```go
func (k *SignaturePrivateKey) Bytes() []byte
```

#### func (*SignaturePrivateKey) Public

```go
func (k *SignaturePrivateKey) Public() []byte
```

### type TreeSecrets

```go
type TreeSecrets struct {
	Leaf LeafIndex
	// contains filtered or unexported fields
}
```

#### func NewTreeSecrets

```go
func NewTreeSecrets(leaf LeafIndex, encryption hpke.PrivateKey) *TreeSecrets
```

#### func (*TreeSecrets) Clone

```go
func (ts *TreeSecrets) Clone() *TreeSecrets
```

#### func (*TreeSecrets) Key

```go
func (ts *TreeSecrets) Key(t *RatchetTree, n NodeIndex) (hpke.PrivateKey, bool)
```

#### func (*TreeSecrets) Set

```go
func (ts *TreeSecrets) Set(n NodeIndex, key hpke.PrivateKey)
```

### type UpdatePath

```go
type UpdatePath struct {
	LeafNode LeafNode
	Nodes    []UpdatePathNode
}
```

### type UpdatePathNode

```go
type UpdatePathNode struct {
	EncryptionKey       []byte
	EncryptedPathSecret []HPKECiphertext
}
```

### type Welcome

```go
type Welcome struct {
	CipherSuite        CipherSuiteID
	Secrets            []EncryptedGroupSecrets
	EncryptedGroupInfo []byte
}
```

### type WireFormat

```go
type WireFormat uint16
```

```go
const (
	WirePublicMessage  WireFormat = 1
	WirePrivateMessage WireFormat = 2
	WireWelcome        WireFormat = 3
	WireGroupInfo      WireFormat = 4
	WireKeyPackage     WireFormat = 5
)
```

### type Writer

```go
type Writer struct {
	// contains filtered or unexported fields
}
```

#### func (*Writer) Bool

```go
func (w *Writer) Bool(v bool)
```

#### func (*Writer) Bytes

```go
func (w *Writer) Bytes() []byte
```

#### func (*Writer) Encode

```go
func (w *Writer) Encode(v encoder)
```

#### func (*Writer) Opaque

```go
func (w *Writer) Opaque(b []byte)
```

#### func (*Writer) Optional

```go
func (w *Writer) Optional(present bool, body func(*Writer))
```

#### func (*Writer) Raw

```go
func (w *Writer) Raw(b []byte)
```

#### func (*Writer) U16

```go
func (w *Writer) U16(v uint16)
```

#### func (*Writer) U32

```go
func (w *Writer) U32(v uint32)
```

#### func (*Writer) U64

```go
func (w *Writer) U64(v uint64)
```

#### func (*Writer) U8

```go
func (w *Writer) U8(v uint8)
```

#### func (*Writer) Varint

```go
func (w *Writer) Varint(n int)
```

#### func (*Writer) Vector

```go
func (w *Writer) Vector(body func(*Writer))
```


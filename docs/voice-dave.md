# dave

```go
import "github.com/prettylittlelies/cry/voice/dave"
```

## Index

- [func EncodeCommitWelcome](#func-encodecommitwelcome)
- [func EncodeProposals](#func-encodeproposals)
- [func IsSilence](#func-issilence)
- [func ParseCommitWelcome](#func-parsecommitwelcome)
- [func ParseExternalSender](#func-parseexternalsender)
- [func ParseTransition](#func-parsetransition)
- [func UserOf](#func-userof)
- [type Decryptor](#type-decryptor)
- [type Encryptor](#type-encryptor)
- [type KeyRatchet](#type-keyratchet)
- [type Opcode](#type-opcode)
- [type Proposals](#type-proposals)
- [type ProposalsOperation](#type-proposalsoperation)
- [type Range](#type-range)
- [type Recognizer](#type-recognizer)
- [type Session](#type-session)

## Constants

```go
const (
	ProtocolVersion = 1
)
```

## Variables

```go
var (
	ErrNotProtocolFrame = errors.New("dave: not a protocol frame")
	ErrNoKey            = errors.New("dave: no key for sender")
)
```

```go
var (
	ErrNoGroup         = errors.New("dave: no mls group")
	ErrUnexpectedUser  = errors.New("dave: unexpected user in proposal")
	ErrDuplicateUser   = errors.New("dave: duplicate user credential in group")
	ErrExternalSender  = errors.New("dave: group external sender mismatch")
	ErrForbiddenChange = errors.New("dave: committer changed its identity")
)
```

```go
var ErrAuthentication = errors.New("dave: frame authentication failed")
```

```go
var ErrErasedGeneration = errors.New("dave: key generation erased")
```

```go
var ErrLEB128 = errors.New("dave: malformed uleb128")
```

```go
var ErrMalformed = errors.New("dave: malformed binary message")
```

## Functions

### func EncodeCommitWelcome

```go
func EncodeCommitWelcome(commit *mls.Message, welcome *mls.Welcome) []byte
```

### func EncodeProposals

```go
func EncodeProposals(p *Proposals) []byte
```

### func IsSilence

```go
func IsSilence(frame []byte) bool
```

### func ParseCommitWelcome

```go
func ParseCommitWelcome(payload []byte) (*mls.Message, *mls.Welcome, error)
```

### func ParseExternalSender

```go
func ParseExternalSender(payload []byte) (*mls.ExternalSender, error)
```

### func ParseTransition

```go
func ParseTransition(payload []byte) (uint16, []byte, error)
```

### func UserOf

```go
func UserOf(c mls.Credential) (snowflake.ID, bool)
```

## Types

### type Decryptor

```go
type Decryptor struct {
	// contains filtered or unexported fields
}
```

#### func (*Decryptor) Decrypt

```go
func (d *Decryptor) Decrypt(frame []byte) ([]byte, error)
```

#### func (*Decryptor) SetRatchet

```go
func (d *Decryptor) SetRatchet(r *KeyRatchet)
```

### type Encryptor

```go
type Encryptor struct {
	// contains filtered or unexported fields
}
```

#### func (*Encryptor) Active

```go
func (e *Encryptor) Active() bool
```

#### func (*Encryptor) Encrypt

```go
func (e *Encryptor) Encrypt(frame []byte) ([]byte, error)
```

#### func (*Encryptor) EncryptRanges

```go
func (e *Encryptor) EncryptRanges(frame []byte, clear []Range) ([]byte, error)
```

#### func (*Encryptor) SetRatchet

```go
func (e *Encryptor) SetRatchet(r *KeyRatchet)
```

### type KeyRatchet

```go
type KeyRatchet struct {
	// contains filtered or unexported fields
}
```

#### func NewKeyRatchet

```go
func NewKeyRatchet(suite *mls.CipherSuite, base []byte) *KeyRatchet
```

#### func (*KeyRatchet) Erase

```go
func (r *KeyRatchet) Erase(generation uint32)
```

#### func (*KeyRatchet) Key

```go
func (r *KeyRatchet) Key(generation uint32) ([]byte, error)
```

### type Opcode

```go
type Opcode uint8
```

```go
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
```

### type Proposals

```go
type Proposals struct {
	Operation ProposalsOperation
	Messages  []*mls.Message
	Refs      [][]byte
}
```

#### func ParseProposals

```go
func ParseProposals(payload []byte) (*Proposals, error)
```

### type ProposalsOperation

```go
type ProposalsOperation uint8
```

```go
const (
	ProposalsAppend ProposalsOperation = 0
	ProposalsRevoke ProposalsOperation = 1
)
```

### type Range

```go
type Range struct{ Offset, Size int }
```

#### func VP8Ranges

```go
func VP8Ranges(frame []byte) []Range
```

### type Recognizer

```go
type Recognizer func(user snowflake.ID) bool
```

### type Session

```go
type Session struct {
	// contains filtered or unexported fields
}
```

#### func NewSession

```go
func NewSession(self snowflake.ID, signer *mls.SignaturePrivateKey) (*Session, error)
```

#### func (*Session) Epoch

```go
func (s *Session) Epoch() uint64
```

#### func (*Session) EpochAuthenticator

```go
func (s *Session) EpochAuthenticator() []byte
```

#### func (*Session) Established

```go
func (s *Session) Established() bool
```

#### func (*Session) Group

```go
func (s *Session) Group() *mls.Group
```

#### func (*Session) KeyPackage

```go
func (s *Session) KeyPackage() ([]byte, error)
```

#### func (*Session) Members

```go
func (s *Session) Members() []snowflake.ID
```

#### func (*Session) ProcessCommit

```go
func (s *Session) ProcessCommit(raw []byte) error
```

#### func (*Session) ProcessProposals

```go
func (s *Session) ProcessProposals(p *Proposals, recognized Recognizer) ([]byte, error)
```

#### func (*Session) ProcessWelcome

```go
func (s *Session) ProcessWelcome(raw []byte, recognized Recognizer) error
```

#### func (*Session) Ratchet

```go
func (s *Session) Ratchet(user snowflake.ID) (*KeyRatchet, error)
```

#### func (*Session) Reset

```go
func (s *Session) Reset(version int, group snowflake.ID)
```

#### func (*Session) SetExternalSender

```go
func (s *Session) SetExternalSender(e *mls.ExternalSender) error
```

#### func (*Session) SignaturePublicKey

```go
func (s *Session) SignaturePublicKey() []byte
```

#### func (*Session) Suite

```go
func (s *Session) Suite() *mls.CipherSuite
```


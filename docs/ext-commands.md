# commands

```go
import "github.com/prettylittlelies/cry/ext/commands"
```

## Index

- [func ConvertBool](#func-convertbool)
- [func IsFlagConverter](#func-isflagconverter)
- [func MaxWidth](#func-maxwidth)
- [func RegisterConverter](#func-registerconverter)
- [func RegisterExtension](#func-registerextension)
- [func StringWidth](#func-stringwidth)
- [type BadArgument](#type-badargument)
- [type BadBoolArgument](#type-badboolargument)
- [type BadColourArgument](#type-badcolourargument)
- [type BadFlagArgument](#type-badflagargument)
- [type BadInviteArgument](#type-badinviteargument)
- [type BadLiteralArgument](#type-badliteralargument)
- [type BadTimestampArgument](#type-badtimestampargument)
- [type BadUnionArgument](#type-badunionargument)
- [type Bot](#type-bot)
- [type BotCheckOncer](#type-botcheckoncer)
- [type BotChecker](#type-botchecker)
- [type BotMissingAnyRole](#type-botmissinganyrole)
- [type BotMissingPermissions](#type-botmissingpermissions)
- [type BotMissingRole](#type-botmissingrole)
- [type BotOption](#type-botoption)
- [type BucketSource](#type-bucketsource)
- [type BucketType](#type-buckettype)
- [type ChannelNotReadable](#type-channelnotreadable)
- [type Check](#type-check)
- [type CheckAnyFailure](#type-checkanyfailure)
- [type CheckFailure](#type-checkfailure)
- [type CleanContent](#type-cleancontent)
- [type Cog](#type-cog)
- [type CogAfterInvoker](#type-cogafterinvoker)
- [type CogBeforeInvoker](#type-cogbeforeinvoker)
- [type CogChecker](#type-cogchecker)
- [type CogDescriber](#type-cogdescriber)
- [type CogErrorHandler](#type-cogerrorhandler)
- [type CogListener](#type-coglistener)
- [type CogLoader](#type-cogloader)
- [type CogUnloader](#type-cogunloader)
- [type Command](#type-command)
- [type CommandCompletionEvent](#type-commandcompletionevent)
- [type CommandError](#type-commanderror)
- [type CommandErrorEvent](#type-commanderrorevent)
- [type CommandEvent](#type-commandevent)
- [type CommandInvokeError](#type-commandinvokeerror)
- [type CommandNotFound](#type-commandnotfound)
- [type CommandOnCooldown](#type-commandoncooldown)
- [type CommandRegistrationError](#type-commandregistrationerror)
- [type Context](#type-context)
- [type ConversionError](#type-conversionerror)
- [type Cooldown](#type-cooldown)
- [type CooldownMapping](#type-cooldownmapping)
- [type DMHelp](#type-dmhelp)
- [type DefaultHelp](#type-defaulthelp)
- [type DisabledCommand](#type-disabledcommand)
- [type ErrorHandler](#type-errorhandler)
- [type ExpectedClosingQuoteError](#type-expectedclosingquoteerror)
- [type Extension](#type-extension)
- [type ExtensionAlreadyLoaded](#type-extensionalreadyloaded)
- [type ExtensionError](#type-extensionerror)
- [type ExtensionFailed](#type-extensionfailed)
- [type ExtensionNotFound](#type-extensionnotfound)
- [type ExtensionNotLoaded](#type-extensionnotloaded)
- [type Flag](#type-flag)
- [type FlagConverter](#type-flagconverter)
- [type FlagError](#type-flagerror)
- [type HelpBase](#type-helpbase)
- [type HelpCommand](#type-helpcommand)
- [type HelpErrorHandler](#type-helperrorhandler)
- [type Hook](#type-hook)
- [type InvalidEndOfQuotedStringError](#type-invalidendofquotedstringerror)
- [type MaxConcurrency](#type-maxconcurrency)
- [type MaxConcurrencyReached](#type-maxconcurrencyreached)
- [type MinimalHelp](#type-minimalhelp)
- [type MissingAnyRole](#type-missinganyrole)
- [type MissingFlagArgument](#type-missingflagargument)
- [type MissingPermissions](#type-missingpermissions)
- [type MissingRequiredArgument](#type-missingrequiredargument)
- [type MissingRequiredAttachment](#type-missingrequiredattachment)
- [type MissingRequiredFlag](#type-missingrequiredflag)
- [type MissingRole](#type-missingrole)
- [type Mode](#type-mode)
- [type NSFWChannelRequired](#type-nsfwchannelrequired)
- [type NoEntryPointError](#type-noentrypointerror)
- [type NoPrivateMessage](#type-noprivatemessage)
- [type NotFoundError](#type-notfounderror)
- [type NotOwner](#type-notowner)
- [type ObjectNotFound](#type-objectnotfound)
- [type Option](#type-option)
- [type PaginatedHelp](#type-paginatedhelp)
- [type Paginator](#type-paginator)
- [type ParamKind](#type-paramkind)
- [type Parameter](#type-parameter)
- [type PartialEmojiConversionFailure](#type-partialemojiconversionfailure)
- [type Prefix](#type-prefix)
- [type PrivateMessageOnly](#type-privatemessageonly)
- [type RangeError](#type-rangeerror)
- [type RoleRef](#type-roleref)
- [type StringView](#type-stringview)
- [type TooManyArguments](#type-toomanyarguments)
- [type TooManyFlags](#type-toomanyflags)
- [type UnexpectedQuoteError](#type-unexpectedquoteerror)
- [type UserInputError](#type-userinputerror)
- [type VerifyChecks](#type-verifychecks)

## Variables

```go
var (
	ErrCommand         = errors.New("command error")
	ErrUserInput       = errors.New("user input error")
	ErrBadArgument     = errors.New("bad argument")
	ErrCheckFailure    = errors.New("check failure")
	ErrArgumentParsing = errors.New("argument parsing error")
	ErrFlag            = errors.New("flag error")
	ErrExtension       = errors.New("extension error")
)
```

```go
var ErrOwnersNotSet = errors.New("commands: owners aren't set")
```

## Functions

### func ConvertBool

```go
func ConvertBool(argument string) (bool, error)
```

### func IsFlagConverter

```go
func IsFlagConverter(t reflect.Type) bool
```

### func MaxWidth

```go
func MaxWidth(names []string) int
```

### func RegisterConverter

```go
func RegisterConverter[T any](fn func(ctx *Context, argument string) (T, error))
```

### func RegisterExtension

```go
func RegisterExtension(name string, ext Extension)
```

### func StringWidth

```go
func StringWidth(s string) int
```

## Types

### type BadArgument

```go
type BadArgument struct {
	// contains filtered or unexported fields
}
```

#### func NewBadArgument

```go
func NewBadArgument(format string, args ...any) *BadArgument
```

#### func (*BadArgument) Error

```go
func (k *BadArgument) Error() string
```

#### func (*BadArgument) Is

```go
func (k *BadArgument) Is(target error) bool
```

### type BadBoolArgument

```go
type BadBoolArgument struct {
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*BadBoolArgument) Error

```go
func (k *BadBoolArgument) Error() string
```

#### func (*BadBoolArgument) Is

```go
func (k *BadBoolArgument) Is(target error) bool
```

### type BadColourArgument

```go
type BadColourArgument struct {
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*BadColourArgument) Error

```go
func (k *BadColourArgument) Error() string
```

#### func (*BadColourArgument) Is

```go
func (k *BadColourArgument) Is(target error) bool
```

### type BadFlagArgument

```go
type BadFlagArgument struct {
	Flag     *Flag
	Argument string
	Original error
	// contains filtered or unexported fields
}
```

#### func (*BadFlagArgument) Error

```go
func (k *BadFlagArgument) Error() string
```

#### func (*BadFlagArgument) Is

```go
func (k *BadFlagArgument) Is(target error) bool
```

#### func (*BadFlagArgument) Unwrap

```go
func (e *BadFlagArgument) Unwrap() error
```

### type BadInviteArgument

```go
type BadInviteArgument struct {
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*BadInviteArgument) Error

```go
func (k *BadInviteArgument) Error() string
```

#### func (*BadInviteArgument) Is

```go
func (k *BadInviteArgument) Is(target error) bool
```

### type BadLiteralArgument

```go
type BadLiteralArgument struct {
	Param    *Parameter
	Literals []string
	Errors   []error
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*BadLiteralArgument) Error

```go
func (k *BadLiteralArgument) Error() string
```

#### func (*BadLiteralArgument) Is

```go
func (k *BadLiteralArgument) Is(target error) bool
```

### type BadTimestampArgument

```go
type BadTimestampArgument struct {
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*BadTimestampArgument) Error

```go
func (k *BadTimestampArgument) Error() string
```

#### func (*BadTimestampArgument) Is

```go
func (k *BadTimestampArgument) Is(target error) bool
```

### type BadUnionArgument

```go
type BadUnionArgument struct {
	Param      *Parameter
	Converters []string
	Errors     []error
	// contains filtered or unexported fields
}
```

#### func (*BadUnionArgument) Error

```go
func (k *BadUnionArgument) Error() string
```

#### func (*BadUnionArgument) Is

```go
func (k *BadUnionArgument) Is(target error) bool
```

### type Bot

```go
type Bot struct {
	Description string
	// contains filtered or unexported fields
}
```

#### func NewBot

```go
func NewBot(c *cry.Client, opts ...BotOption) *Bot
```

#### func (*Bot) AddCheck

```go
func (b *Bot) AddCheck(check Check, once bool) (remove func())
```

#### func (*Bot) AddCog

```go
func (b *Bot) AddCog(ctx context.Context, cog Cog, override bool) error
```

#### func (*Bot) AddCommand

```go
func (b *Bot) AddCommand(c *Command) error
```

#### func (*Bot) AfterInvoke

```go
func (b *Bot) AfterInvoke(h Hook)
```

#### func (*Bot) BeforeInvoke

```go
func (b *Bot) BeforeInvoke(h Hook)
```

#### func (*Bot) Client

```go
func (b *Bot) Client() *cry.Client
```

#### func (*Bot) Close

```go
func (b *Bot) Close(ctx context.Context)
```

#### func (*Bot) Cogs

```go
func (b *Bot) Cogs() map[string]Cog
```

#### func (*Bot) Commands

```go
func (g *Bot) Commands() []*Command
```

#### func (*Bot) Extensions

```go
func (b *Bot) Extensions() []string
```

#### func (*Bot) Get

```go
func (g *Bot) Get(name string) *Command
```

#### func (*Bot) GetCog

```go
func (b *Bot) GetCog(name string) Cog
```

#### func (*Bot) GetCommand

```go
func (b *Bot) GetCommand(name string) *Command
```

#### func (*Bot) GetContext

```go
func (b *Bot) GetContext(ctx context.Context, m *discord.MessageCreateEvent) *Context
```

#### func (*Bot) GetPrefix

```go
func (b *Bot) GetPrefix(m *discord.MessageCreateEvent) []string
```

#### func (*Bot) HelpCommand

```go
func (b *Bot) HelpCommand() HelpCommand
```

#### func (*Bot) Invoke

```go
func (b *Bot) Invoke(c *Context)
```

#### func (*Bot) IsOwner

```go
func (b *Bot) IsOwner(user snowflake.ID) (bool, error)
```

#### func (*Bot) LoadExtension

```go
func (b *Bot) LoadExtension(ctx context.Context, name string) error
```

#### func (*Bot) ProcessCommands

```go
func (b *Bot) ProcessCommands(ctx context.Context, m *discord.MessageCreateEvent)
```

#### func (*Bot) ReloadExtension

```go
func (b *Bot) ReloadExtension(ctx context.Context, name string) error
```

#### func (*Bot) RemoveCog

```go
func (b *Bot) RemoveCog(ctx context.Context, name string) Cog
```

#### func (*Bot) RemoveCommand

```go
func (b *Bot) RemoveCommand(name string) *Command
```

#### func (*Bot) SetHelpCommand

```go
func (b *Bot) SetHelpCommand(h HelpCommand) error
```

#### func (*Bot) UnloadExtension

```go
func (b *Bot) UnloadExtension(ctx context.Context, name string) error
```

#### func (*Bot) WalkCommands

```go
func (g *Bot) WalkCommands() iter.Seq[*Command]
```

### type BotCheckOncer

```go
type BotCheckOncer interface {
	BotCheckOnce(ctx *Context) (bool, error)
}
```

### type BotChecker

```go
type BotChecker interface {
	BotCheck(ctx *Context) (bool, error)
}
```

### type BotMissingAnyRole

```go
type BotMissingAnyRole struct {
	MissingRoles []string
	// contains filtered or unexported fields
}
```

#### func (*BotMissingAnyRole) Error

```go
func (k *BotMissingAnyRole) Error() string
```

#### func (*BotMissingAnyRole) Is

```go
func (k *BotMissingAnyRole) Is(target error) bool
```

### type BotMissingPermissions

```go
type BotMissingPermissions struct {
	MissingPermissions []string
	// contains filtered or unexported fields
}
```

#### func (*BotMissingPermissions) Error

```go
func (k *BotMissingPermissions) Error() string
```

#### func (*BotMissingPermissions) Is

```go
func (k *BotMissingPermissions) Is(target error) bool
```

### type BotMissingRole

```go
type BotMissingRole struct {
	MissingRole string
	// contains filtered or unexported fields
}
```

#### func (*BotMissingRole) Error

```go
func (k *BotMissingRole) Error() string
```

#### func (*BotMissingRole) Is

```go
func (k *BotMissingRole) Is(target error) bool
```

### type BotOption

```go
type BotOption func(*Bot)
```

#### func WithCaseInsensitive

```go
func WithCaseInsensitive() BotOption
```

#### func WithDescription

```go
func WithDescription(d string) BotOption
```

#### func WithHelpCommand

```go
func WithHelpCommand(h HelpCommand) BotOption
```

#### func WithLogger

```go
func WithLogger(l *slog.Logger) BotOption
```

#### func WithMode

```go
func WithMode(m Mode) BotOption
```

#### func WithOwnerIDs

```go
func WithOwnerIDs(ids ...snowflake.ID) BotOption
```

#### func WithPrefix

```go
func WithPrefix(p Prefix) BotOption
```

#### func WithStripAfterPrefix

```go
func WithStripAfterPrefix() BotOption
```

### type BucketSource

```go
type BucketSource interface {
	AuthorID() snowflake.ID
	GuildID() snowflake.ID
	ChannelID() snowflake.ID
	CategoryID() snowflake.ID
	TopRoleID() snowflake.ID
	Private() bool
}
```

### type BucketType

```go
type BucketType int
```

```go
const (
	BucketDefault BucketType = iota
	BucketUser
	BucketGuild
	BucketChannel
	BucketMember
	BucketCategory
	BucketRole
)
```

#### func (BucketType) Key

```go
func (b BucketType) Key(src BucketSource) any
```

#### func (BucketType) String

```go
func (b BucketType) String() string
```

### type ChannelNotReadable

```go
type ChannelNotReadable struct {
	Channel string
	// contains filtered or unexported fields
}
```

#### func (*ChannelNotReadable) Error

```go
func (k *ChannelNotReadable) Error() string
```

#### func (*ChannelNotReadable) Is

```go
func (k *ChannelNotReadable) Is(target error) bool
```

### type Check

```go
type Check func(ctx *Context) (bool, error)
```

#### func BotHasAnyRole

```go
func BotHasAnyRole(r ...RoleRef) Check
```

#### func BotHasGuildPermissions

```go
func BotHasGuildPermissions(p ...discord.Permission) Check
```

#### func BotHasPermissions

```go
func BotHasPermissions(p ...discord.Permission) Check
```

#### func BotHasRole

```go
func BotHasRole(r RoleRef) Check
```

#### func CheckAny

```go
func CheckAny(checks ...Check) Check
```

#### func DMOnly

```go
func DMOnly() Check
```

#### func GuildOnly

```go
func GuildOnly() Check
```

#### func HasAnyRole

```go
func HasAnyRole(r ...RoleRef) Check
```

#### func HasGuildPermissions

```go
func HasGuildPermissions(p ...discord.Permission) Check
```

#### func HasPermissions

```go
func HasPermissions(p ...discord.Permission) Check
```

#### func HasRole

```go
func HasRole(r RoleRef) Check
```

#### func IsNSFW

```go
func IsNSFW() Check
```

#### func IsOwner

```go
func IsOwner() Check
```

### type CheckAnyFailure

```go
type CheckAnyFailure struct {
	Errors []error
	// contains filtered or unexported fields
}
```

#### func (*CheckAnyFailure) Error

```go
func (k *CheckAnyFailure) Error() string
```

#### func (*CheckAnyFailure) Is

```go
func (k *CheckAnyFailure) Is(target error) bool
```

### type CheckFailure

```go
type CheckFailure struct {
	// contains filtered or unexported fields
}
```

#### func NewCheckFailure

```go
func NewCheckFailure(message string) *CheckFailure
```

#### func (*CheckFailure) Error

```go
func (k *CheckFailure) Error() string
```

#### func (*CheckFailure) Is

```go
func (k *CheckFailure) Is(target error) bool
```

### type CleanContent

```go
type CleanContent string
```

### type Cog

```go
type Cog interface {
	CogName() string
	CogCommands() []*Command
}
```

### type CogAfterInvoker

```go
type CogAfterInvoker interface {
	CogAfterInvoke(ctx *Context) error
}
```

### type CogBeforeInvoker

```go
type CogBeforeInvoker interface {
	CogBeforeInvoke(ctx *Context) error
}
```

### type CogChecker

```go
type CogChecker interface {
	CogCheck(ctx *Context) (bool, error)
}
```

### type CogDescriber

```go
type CogDescriber interface {
	CogDescription() string
}
```

### type CogErrorHandler

```go
type CogErrorHandler interface {
	CogCommandError(ctx *Context, err error)
}
```

### type CogListener

```go
type CogListener interface {
	CogListeners(b *Bot) []func()
}
```

### type CogLoader

```go
type CogLoader interface {
	CogLoad(ctx context.Context, b *Bot) error
}
```

### type CogUnloader

```go
type CogUnloader interface {
	CogUnload(ctx context.Context, b *Bot)
}
```

### type Command

```go
type Command struct {
	Name                 string
	Aliases              []string
	Help                 string
	Brief                string
	Usage                string
	Description          string
	Enabled              bool
	Hidden               bool
	RestIsRaw            bool
	IgnoreExtra          bool
	RequireVarPositional bool
	CooldownAfterParsing bool
	InvokeWithoutCommand bool
	Extras               map[string]any
	Checks               []Check
	Params               []*Parameter
	Parent               *Command
	Cog                  Cog
	OnError              ErrorHandler
	// contains filtered or unexported fields
}
```

#### func New

```go
func New[P any](name string, fn func(ctx *Context, params P) error, opts ...Option) *Command
```

#### func (*Command) AddCommand

```go
func (c *Command) AddCommand(sub *Command) error
```

#### func (*Command) CanRun

```go
func (c *Command) CanRun(ctx *Context) (bool, error)
```

#### func (*Command) Commands

```go
func (g *Command) Commands() []*Command
```

#### func (*Command) CooldownRetryAfter

```go
func (c *Command) CooldownRetryAfter(ctx *Context) time.Duration
```

#### func (*Command) Get

```go
func (g *Command) Get(name string) *Command
```

#### func (*Command) GetCommand

```go
func (g *Command) GetCommand(name string) *Command
```

#### func (*Command) IsGroup

```go
func (c *Command) IsGroup() bool
```

#### func (*Command) IsOnCooldown

```go
func (c *Command) IsOnCooldown(ctx *Context) bool
```

#### func (*Command) Parents

```go
func (c *Command) Parents() []*Command
```

#### func (*Command) QualifiedName

```go
func (c *Command) QualifiedName() string
```

#### func (*Command) ResetCooldown

```go
func (c *Command) ResetCooldown(ctx *Context)
```

#### func (*Command) Root

```go
func (c *Command) Root() *Command
```

#### func (*Command) ShortDoc

```go
func (c *Command) ShortDoc() string
```

#### func (*Command) Signature

```go
func (c *Command) Signature() string
```

#### func (*Command) WalkCommands

```go
func (g *Command) WalkCommands() iter.Seq[*Command]
```

### type CommandCompletionEvent

```go
type CommandCompletionEvent struct{ Context *Context }
```

### type CommandError

```go
type CommandError struct {
	// contains filtered or unexported fields
}
```

#### func NewCommandError

```go
func NewCommandError(message string) *CommandError
```

#### func (*CommandError) Error

```go
func (k *CommandError) Error() string
```

#### func (*CommandError) Is

```go
func (k *CommandError) Is(target error) bool
```

### type CommandErrorEvent

```go
type CommandErrorEvent struct {
	Context *Context
	Err     error
}
```

### type CommandEvent

```go
type CommandEvent struct{ Context *Context }
```

### type CommandInvokeError

```go
type CommandInvokeError struct {
	Original error
	// contains filtered or unexported fields
}
```

#### func (*CommandInvokeError) Error

```go
func (k *CommandInvokeError) Error() string
```

#### func (*CommandInvokeError) Is

```go
func (k *CommandInvokeError) Is(target error) bool
```

#### func (*CommandInvokeError) Unwrap

```go
func (e *CommandInvokeError) Unwrap() error
```

### type CommandNotFound

```go
type CommandNotFound struct {
	// contains filtered or unexported fields
}
```

#### func (*CommandNotFound) Error

```go
func (k *CommandNotFound) Error() string
```

#### func (*CommandNotFound) Is

```go
func (k *CommandNotFound) Is(target error) bool
```

### type CommandOnCooldown

```go
type CommandOnCooldown struct {
	Cooldown   *Cooldown
	RetryAfter time.Duration
	Type       BucketType
	// contains filtered or unexported fields
}
```

#### func (*CommandOnCooldown) Error

```go
func (k *CommandOnCooldown) Error() string
```

#### func (*CommandOnCooldown) Is

```go
func (k *CommandOnCooldown) Is(target error) bool
```

### type CommandRegistrationError

```go
type CommandRegistrationError struct {
	Name          string
	AliasConflict bool
}
```

#### func (*CommandRegistrationError) Error

```go
func (e *CommandRegistrationError) Error() string
```

### type Context

```go
type Context struct {
	context.Context
	Message           *discord.MessageCreateEvent
	Bot               *Bot
	View              *StringView
	Prefix            string
	HasPrefix         bool
	Command           *Command
	InvokedWith       string
	InvokedParents    []string
	InvokedSubcommand *Command
	SubcommandPassed  string
	CommandFailed     bool
	CurrentParameter  *Parameter
	CurrentArgument   string
	// contains filtered or unexported fields
}
```

#### func (*Context) Author

```go
func (c *Context) Author() discord.PartialUser
```

#### func (*Context) AuthorID

```go
func (c *Context) AuthorID() snowflake.ID
```

#### func (*Context) CategoryID

```go
func (c *Context) CategoryID() snowflake.ID
```

#### func (*Context) Channel

```go
func (c *Context) Channel() cry.Channel
```

#### func (*Context) ChannelID

```go
func (c *Context) ChannelID() snowflake.ID
```

#### func (*Context) CleanPrefix

```go
func (c *Context) CleanPrefix() string
```

#### func (*Context) Cog

```go
func (c *Context) Cog() Cog
```

#### func (*Context) GuildID

```go
func (c *Context) GuildID() snowflake.ID
```

#### func (*Context) MessageHandle

```go
func (c *Context) MessageHandle() cry.Message
```

#### func (*Context) Private

```go
func (c *Context) Private() bool
```

#### func (*Context) Reinvoke

```go
func (c *Context) Reinvoke(callHooks, restart bool) error
```

#### func (*Context) Reply

```go
func (c *Context) Reply(content string) (*discord.Message, error)
```

#### func (*Context) Scope

```go
func (c *Context) Scope() converting.Scope
```

#### func (*Context) Send

```go
func (c *Context) Send(content string) (*discord.Message, error)
```

#### func (*Context) SendAuthor

```go
func (c *Context) SendAuthor(content string) (*discord.Message, error)
```

#### func (*Context) SendMessage

```go
func (c *Context) SendMessage(params discord.CreateMessageParams) (*discord.Message, error)
```

#### func (*Context) TopRoleID

```go
func (c *Context) TopRoleID() snowflake.ID
```

#### func (*Context) Valid

```go
func (c *Context) Valid() bool
```

### type ConversionError

```go
type ConversionError struct {
	Converter string
	Original  error
	// contains filtered or unexported fields
}
```

#### func (*ConversionError) Error

```go
func (k *ConversionError) Error() string
```

#### func (*ConversionError) Is

```go
func (k *ConversionError) Is(target error) bool
```

#### func (*ConversionError) Unwrap

```go
func (e *ConversionError) Unwrap() error
```

### type Cooldown

```go
type Cooldown struct {
	Rate int
	Per  time.Duration
	// contains filtered or unexported fields
}
```

#### func NewCooldown

```go
func NewCooldown(rate int, per time.Duration) *Cooldown
```

#### func (*Cooldown) Copy

```go
func (c *Cooldown) Copy() *Cooldown
```

#### func (*Cooldown) Reset

```go
func (c *Cooldown) Reset()
```

#### func (*Cooldown) RetryAfter

```go
func (c *Cooldown) RetryAfter(now time.Time) time.Duration
```

#### func (*Cooldown) Tokens

```go
func (c *Cooldown) Tokens(now time.Time) int
```

#### func (*Cooldown) UpdateRateLimit

```go
func (c *Cooldown) UpdateRateLimit(now time.Time, tokens int) time.Duration
```

### type CooldownMapping

```go
type CooldownMapping struct {
	Type BucketType
	// contains filtered or unexported fields
}
```

#### func NewCooldownMapping

```go
func NewCooldownMapping(c *Cooldown, t BucketType) *CooldownMapping
```

#### func NewDynamicCooldownMapping

```go
func NewDynamicCooldownMapping(factory func(BucketSource) *Cooldown, t BucketType) *CooldownMapping
```

#### func (*CooldownMapping) Bucket

```go
func (m *CooldownMapping) Bucket(src BucketSource, now time.Time) *Cooldown
```

#### func (*CooldownMapping) UpdateRateLimit

```go
func (m *CooldownMapping) UpdateRateLimit(src BucketSource, now time.Time, tokens int) (*Cooldown, time.Duration)
```

#### func (*CooldownMapping) Valid

```go
func (m *CooldownMapping) Valid() bool
```

### type DMHelp

```go
type DMHelp int
```

```go
const (
	DMNever DMHelp = iota
	DMAlways
	DMOverThreshold
)
```

### type DefaultHelp

```go
type DefaultHelp struct {
	PaginatedHelp
	Width                      int
	Indent                     int
	ArgumentsHeading           string
	DefaultArgumentDescription string
	ShowParameterDescriptions  bool
}
```

#### func NewDefaultHelp

```go
func NewDefaultHelp() *DefaultHelp
```

#### func (*DefaultHelp) AddCommandArguments

```go
func (h *DefaultHelp) AddCommandArguments(command *Command) error
```

#### func (*DefaultHelp) AddCommandFormatting

```go
func (h *DefaultHelp) AddCommandFormatting(command *Command) error
```

#### func (*DefaultHelp) AddIndentedCommands

```go
func (h *DefaultHelp) AddIndentedCommands(commands []*Command, heading string, maxSize int) error
```

#### func (*DefaultHelp) CommandSignature

```go
func (h *DefaultHelp) CommandSignature(command *Command) string
```

#### func (*DefaultHelp) EndingNote

```go
func (h *DefaultHelp) EndingNote() string
```

#### func (*DefaultHelp) SendBotHelp

```go
func (h *DefaultHelp) SendBotHelp(map[Cog][]*Command) error
```

#### func (*DefaultHelp) SendCogHelp

```go
func (h *DefaultHelp) SendCogHelp(cog Cog) error
```

#### func (*DefaultHelp) SendCommandHelp

```go
func (h *DefaultHelp) SendCommandHelp(command *Command) error
```

#### func (*DefaultHelp) SendGroupHelp

```go
func (h *DefaultHelp) SendGroupHelp(group *Command) error
```

#### func (*DefaultHelp) ShortenText

```go
func (h *DefaultHelp) ShortenText(text string) string
```

### type DisabledCommand

```go
type DisabledCommand struct {
	// contains filtered or unexported fields
}
```

#### func (*DisabledCommand) Error

```go
func (k *DisabledCommand) Error() string
```

#### func (*DisabledCommand) Is

```go
func (k *DisabledCommand) Is(target error) bool
```

### type ErrorHandler

```go
type ErrorHandler func(ctx *Context, err error)
```

### type ExpectedClosingQuoteError

```go
type ExpectedClosingQuoteError struct {
	CloseQuote rune
	// contains filtered or unexported fields
}
```

#### func (*ExpectedClosingQuoteError) Error

```go
func (k *ExpectedClosingQuoteError) Error() string
```

#### func (*ExpectedClosingQuoteError) Is

```go
func (k *ExpectedClosingQuoteError) Is(target error) bool
```

### type Extension

```go
type Extension struct {
	Setup    func(ctx context.Context, b *Bot) error
	Teardown func(ctx context.Context, b *Bot)
}
```

### type ExtensionAlreadyLoaded

```go
type ExtensionAlreadyLoaded struct{ ExtensionError }
```

#### func (*ExtensionAlreadyLoaded) Error

```go
func (k *ExtensionAlreadyLoaded) Error() string
```

#### func (*ExtensionAlreadyLoaded) Is

```go
func (k *ExtensionAlreadyLoaded) Is(target error) bool
```

### type ExtensionError

```go
type ExtensionError struct {
	Name string
	// contains filtered or unexported fields
}
```

#### func (*ExtensionError) Error

```go
func (k *ExtensionError) Error() string
```

#### func (*ExtensionError) Is

```go
func (k *ExtensionError) Is(target error) bool
```

### type ExtensionFailed

```go
type ExtensionFailed struct {
	ExtensionError
	Original error
}
```

#### func (*ExtensionFailed) Error

```go
func (k *ExtensionFailed) Error() string
```

#### func (*ExtensionFailed) Is

```go
func (k *ExtensionFailed) Is(target error) bool
```

#### func (*ExtensionFailed) Unwrap

```go
func (e *ExtensionFailed) Unwrap() error
```

### type ExtensionNotFound

```go
type ExtensionNotFound struct{ ExtensionError }
```

#### func (*ExtensionNotFound) Error

```go
func (k *ExtensionNotFound) Error() string
```

#### func (*ExtensionNotFound) Is

```go
func (k *ExtensionNotFound) Is(target error) bool
```

### type ExtensionNotLoaded

```go
type ExtensionNotLoaded struct{ ExtensionError }
```

#### func (*ExtensionNotLoaded) Error

```go
func (k *ExtensionNotLoaded) Error() string
```

#### func (*ExtensionNotLoaded) Is

```go
func (k *ExtensionNotLoaded) Is(target error) bool
```

### type Flag

```go
type Flag struct {
	Name        string
	Aliases     []string
	Attribute   string
	Type        reflect.Type
	Default     string
	HasDefault  bool
	MaxArgs     int
	Override    bool
	Description string
	Positional  bool
	CastToDict  bool
	// contains filtered or unexported fields
}
```

#### func FlagsOf

```go
func FlagsOf(t reflect.Type) ([]*Flag, error)
```

#### func (*Flag) Required

```go
func (f *Flag) Required() bool
```

### type FlagConverter

```go
type FlagConverter struct{}
```

### type FlagError

```go
type FlagError struct {
	// contains filtered or unexported fields
}
```

#### func (*FlagError) Error

```go
func (k *FlagError) Error() string
```

#### func (*FlagError) Is

```go
func (k *FlagError) Is(target error) bool
```

### type HelpBase

```go
type HelpBase struct {
	ShowHidden     bool
	VerifyChecks   VerifyChecks
	CommandName    string
	CommandOptions []Option
	Cog            Cog
	Context        *Context
	// contains filtered or unexported fields
}
```

#### func (*HelpBase) BotMapping

```go
func (h *HelpBase) BotMapping() map[Cog][]*Command
```

#### func (*HelpBase) CogCommands

```go
func (h *HelpBase) CogCommands(cog Cog) []*Command
```

#### func (*HelpBase) Command

```go
func (h *HelpBase) Command() *Command
```

#### func (*HelpBase) CommandNotFound

```go
func (h *HelpBase) CommandNotFound(name string) string
```

#### func (*HelpBase) CommandSignature

```go
func (h *HelpBase) CommandSignature(command *Command) string
```

#### func (*HelpBase) FilterCommands

```go
func (h *HelpBase) FilterCommands(commands []*Command, sort bool, key func(*Command) string) []*Command
```

#### func (*HelpBase) InvokedWith

```go
func (h *HelpBase) InvokedWith() string
```

#### func (*HelpBase) PrepareHelpCommand

```go
func (h *HelpBase) PrepareHelpCommand(*Context, string) error
```

#### func (*HelpBase) RemoveMentions

```go
func (h *HelpBase) RemoveMentions(s string) string
```

#### func (*HelpBase) SendBotHelp

```go
func (h *HelpBase) SendBotHelp(map[Cog][]*Command) error
```

#### func (*HelpBase) SendCogHelp

```go
func (h *HelpBase) SendCogHelp(Cog) error
```

#### func (*HelpBase) SendCommandHelp

```go
func (h *HelpBase) SendCommandHelp(*Command) error
```

#### func (*HelpBase) SendErrorMessage

```go
func (h *HelpBase) SendErrorMessage(message string) error
```

#### func (*HelpBase) SendGroupHelp

```go
func (h *HelpBase) SendGroupHelp(*Command) error
```

#### func (*HelpBase) SetCog

```go
func (h *HelpBase) SetCog(cog Cog)
```

#### func (*HelpBase) SubcommandNotFound

```go
func (h *HelpBase) SubcommandNotFound(command *Command, name string) string
```

### type HelpCommand

```go
type HelpCommand interface {
	PrepareHelpCommand(ctx *Context, command string) error
	SendBotHelp(mapping map[Cog][]*Command) error
	SendCogHelp(cog Cog) error
	SendGroupHelp(group *Command) error
	SendCommandHelp(command *Command) error
	SendErrorMessage(message string) error
	CommandNotFound(name string) string
	SubcommandNotFound(command *Command, name string) string
	// contains filtered or unexported methods
}
```

### type HelpErrorHandler

```go
type HelpErrorHandler interface {
	OnHelpCommandError(ctx *Context, err error)
}
```

### type Hook

```go
type Hook func(ctx *Context) error
```

### type InvalidEndOfQuotedStringError

```go
type InvalidEndOfQuotedStringError struct {
	Char rune
	// contains filtered or unexported fields
}
```

#### func (*InvalidEndOfQuotedStringError) Error

```go
func (k *InvalidEndOfQuotedStringError) Error() string
```

#### func (*InvalidEndOfQuotedStringError) Is

```go
func (k *InvalidEndOfQuotedStringError) Is(target error) bool
```

### type MaxConcurrency

```go
type MaxConcurrency struct {
	Number int
	Per    BucketType
	Wait   bool
	// contains filtered or unexported fields
}
```

#### func NewMaxConcurrency

```go
func NewMaxConcurrency(number int, per BucketType, wait bool) *MaxConcurrency
```

#### func (*MaxConcurrency) Acquire

```go
func (m *MaxConcurrency) Acquire(ctx context.Context, src BucketSource) error
```

#### func (*MaxConcurrency) Copy

```go
func (m *MaxConcurrency) Copy() *MaxConcurrency
```

#### func (*MaxConcurrency) Release

```go
func (m *MaxConcurrency) Release(src BucketSource)
```

### type MaxConcurrencyReached

```go
type MaxConcurrencyReached struct {
	Number int
	Per    BucketType
	// contains filtered or unexported fields
}
```

#### func (*MaxConcurrencyReached) Error

```go
func (k *MaxConcurrencyReached) Error() string
```

#### func (*MaxConcurrencyReached) Is

```go
func (k *MaxConcurrencyReached) Is(target error) bool
```

### type MinimalHelp

```go
type MinimalHelp struct {
	PaginatedHelp
	AliasesHeading string
}
```

#### func NewMinimalHelp

```go
func NewMinimalHelp() *MinimalHelp
```

#### func (*MinimalHelp) AddBotCommandsFormatting

```go
func (h *MinimalHelp) AddBotCommandsFormatting(commands []*Command, heading string) error
```

#### func (*MinimalHelp) AddCommandFormatting

```go
func (h *MinimalHelp) AddCommandFormatting(command *Command) error
```

#### func (*MinimalHelp) AddSubcommandFormatting

```go
func (h *MinimalHelp) AddSubcommandFormatting(command *Command) error
```

#### func (*MinimalHelp) CommandSignature

```go
func (h *MinimalHelp) CommandSignature(command *Command) string
```

#### func (*MinimalHelp) EndingNote

```go
func (h *MinimalHelp) EndingNote() string
```

#### func (*MinimalHelp) OpeningNote

```go
func (h *MinimalHelp) OpeningNote() string
```

#### func (*MinimalHelp) SendBotHelp

```go
func (h *MinimalHelp) SendBotHelp(map[Cog][]*Command) error
```

#### func (*MinimalHelp) SendCogHelp

```go
func (h *MinimalHelp) SendCogHelp(cog Cog) error
```

#### func (*MinimalHelp) SendCommandHelp

```go
func (h *MinimalHelp) SendCommandHelp(command *Command) error
```

#### func (*MinimalHelp) SendGroupHelp

```go
func (h *MinimalHelp) SendGroupHelp(group *Command) error
```

### type MissingAnyRole

```go
type MissingAnyRole struct {
	MissingRoles []string
	// contains filtered or unexported fields
}
```

#### func (*MissingAnyRole) Error

```go
func (k *MissingAnyRole) Error() string
```

#### func (*MissingAnyRole) Is

```go
func (k *MissingAnyRole) Is(target error) bool
```

### type MissingFlagArgument

```go
type MissingFlagArgument struct {
	Flag *Flag
	// contains filtered or unexported fields
}
```

#### func (*MissingFlagArgument) Error

```go
func (k *MissingFlagArgument) Error() string
```

#### func (*MissingFlagArgument) Is

```go
func (k *MissingFlagArgument) Is(target error) bool
```

### type MissingPermissions

```go
type MissingPermissions struct {
	MissingPermissions []string
	// contains filtered or unexported fields
}
```

#### func (*MissingPermissions) Error

```go
func (k *MissingPermissions) Error() string
```

#### func (*MissingPermissions) Is

```go
func (k *MissingPermissions) Is(target error) bool
```

### type MissingRequiredArgument

```go
type MissingRequiredArgument struct {
	Param *Parameter
	// contains filtered or unexported fields
}
```

#### func (*MissingRequiredArgument) Error

```go
func (k *MissingRequiredArgument) Error() string
```

#### func (*MissingRequiredArgument) Is

```go
func (k *MissingRequiredArgument) Is(target error) bool
```

### type MissingRequiredAttachment

```go
type MissingRequiredAttachment struct {
	Param *Parameter
	// contains filtered or unexported fields
}
```

#### func (*MissingRequiredAttachment) Error

```go
func (k *MissingRequiredAttachment) Error() string
```

#### func (*MissingRequiredAttachment) Is

```go
func (k *MissingRequiredAttachment) Is(target error) bool
```

### type MissingRequiredFlag

```go
type MissingRequiredFlag struct {
	Flag *Flag
	// contains filtered or unexported fields
}
```

#### func (*MissingRequiredFlag) Error

```go
func (k *MissingRequiredFlag) Error() string
```

#### func (*MissingRequiredFlag) Is

```go
func (k *MissingRequiredFlag) Is(target error) bool
```

### type MissingRole

```go
type MissingRole struct {
	MissingRole string
	// contains filtered or unexported fields
}
```

#### func (*MissingRole) Error

```go
func (k *MissingRole) Error() string
```

#### func (*MissingRole) Is

```go
func (k *MissingRole) Is(target error) bool
```

### type Mode

```go
type Mode int
```

```go
const (
	Default Mode = iota
	SelfBot
	UserBot
)
```

### type NSFWChannelRequired

```go
type NSFWChannelRequired struct {
	Channel string
	// contains filtered or unexported fields
}
```

#### func (*NSFWChannelRequired) Error

```go
func (k *NSFWChannelRequired) Error() string
```

#### func (*NSFWChannelRequired) Is

```go
func (k *NSFWChannelRequired) Is(target error) bool
```

### type NoEntryPointError

```go
type NoEntryPointError struct{ ExtensionError }
```

#### func (*NoEntryPointError) Error

```go
func (k *NoEntryPointError) Error() string
```

#### func (*NoEntryPointError) Is

```go
func (k *NoEntryPointError) Is(target error) bool
```

### type NoPrivateMessage

```go
type NoPrivateMessage struct {
	// contains filtered or unexported fields
}
```

#### func (*NoPrivateMessage) Error

```go
func (k *NoPrivateMessage) Error() string
```

#### func (*NoPrivateMessage) Is

```go
func (k *NoPrivateMessage) Is(target error) bool
```

### type NotFoundError

```go
type NotFoundError struct {
	Kind     string
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*NotFoundError) Error

```go
func (k *NotFoundError) Error() string
```

#### func (*NotFoundError) Is

```go
func (k *NotFoundError) Is(target error) bool
```

### type NotOwner

```go
type NotOwner struct {
	// contains filtered or unexported fields
}
```

#### func (*NotOwner) Error

```go
func (k *NotOwner) Error() string
```

#### func (*NotOwner) Is

```go
func (k *NotOwner) Is(target error) bool
```

### type ObjectNotFound

```go
type ObjectNotFound struct {
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*ObjectNotFound) Error

```go
func (k *ObjectNotFound) Error() string
```

#### func (*ObjectNotFound) Is

```go
func (k *ObjectNotFound) Is(target error) bool
```

### type Option

```go
type Option func(*Command)
```

#### func AfterInvoke

```go
func AfterInvoke(h Hook) Option
```

#### func Aliases

```go
func Aliases(a ...string) Option
```

#### func BeforeInvoke

```go
func BeforeInvoke(h Hook) Option
```

#### func Brief

```go
func Brief(s string) Option
```

#### func CaseInsensitive

```go
func CaseInsensitive() Option
```

#### func Checks

```go
func Checks(checks ...Check) Option
```

#### func CooldownAfterParsing

```go
func CooldownAfterParsing() Option
```

#### func Description

```go
func Description(s string) Option
```

#### func Disabled

```go
func Disabled() Option
```

#### func Extras

```go
func Extras(m map[string]any) Option
```

#### func Help

```go
func Help(s string) Option
```

#### func Hidden

```go
func Hidden() Option
```

#### func InvokeWithoutCommand

```go
func InvokeWithoutCommand() Option
```

#### func OnError

```go
func OnError(h ErrorHandler) Option
```

#### func RequireVarPositional

```go
func RequireVarPositional() Option
```

#### func RestIsRaw

```go
func RestIsRaw() Option
```

#### func StrictArguments

```go
func StrictArguments() Option
```

#### func Subcommands

```go
func Subcommands(subs ...*Command) Option
```

#### func Usage

```go
func Usage(s string) Option
```

#### func WithCooldown

```go
func WithCooldown(rate int, per time.Duration, t BucketType) Option
```

#### func WithDynamicCooldown

```go
func WithDynamicCooldown(factory func(BucketSource) *Cooldown, t BucketType) Option
```

#### func WithMaxConcurrency

```go
func WithMaxConcurrency(number int, per BucketType, wait bool) Option
```

### type PaginatedHelp

```go
type PaginatedHelp struct {
	HelpBase
	SortCommands    bool
	DMHelp          DMHelp
	DMHelpThreshold int
	NoCategory      string
	CommandsHeading string
	Paginator       Paginator
}
```

#### func (*PaginatedHelp) PrepareHelpCommand

```go
func (h *PaginatedHelp) PrepareHelpCommand(*Context, string) error
```

#### func (*PaginatedHelp) Send

```go
func (h *PaginatedHelp) Send(content string) error
```

#### func (*PaginatedHelp) SendErrorMessage

```go
func (h *PaginatedHelp) SendErrorMessage(message string) error
```

#### func (*PaginatedHelp) SendPages

```go
func (h *PaginatedHelp) SendPages() error
```

### type Paginator

```go
type Paginator struct {
	Prefix  string
	Suffix  string
	MaxSize int
	Linesep string
	// contains filtered or unexported fields
}
```

#### func NewPaginator

```go
func NewPaginator(prefix, suffix string) Paginator
```

#### func (*Paginator) AddLine

```go
func (p *Paginator) AddLine(line string) error
```

#### func (*Paginator) AddParagraph

```go
func (p *Paginator) AddParagraph(line string) error
```

#### func (*Paginator) Clear

```go
func (p *Paginator) Clear()
```

#### func (*Paginator) ClosePage

```go
func (p *Paginator) ClosePage()
```

#### func (*Paginator) Len

```go
func (p *Paginator) Len() int
```

#### func (*Paginator) Pages

```go
func (p *Paginator) Pages() []string
```

### type ParamKind

```go
type ParamKind int
```

```go
const (
	Positional ParamKind = iota
	KeywordOnly
	VarPositional
)
```

### type Parameter

```go
type Parameter struct {
	Name             string
	DisplayedName    string
	Description      string
	DisplayedDefault string
	Kind             ParamKind
	Type             reflect.Type
	Default          string
	HasDefault       bool
	Greedy           bool
	Literals         []string
	Min, Max         any
	Tag              reflect.StructTag
	// contains filtered or unexported fields
}
```

#### func (*Parameter) DisplayDefault

```go
func (p *Parameter) DisplayDefault() (string, bool)
```

#### func (*Parameter) Optional

```go
func (p *Parameter) Optional() bool
```

#### func (*Parameter) Required

```go
func (p *Parameter) Required() bool
```

### type PartialEmojiConversionFailure

```go
type PartialEmojiConversionFailure struct {
	Argument string
	// contains filtered or unexported fields
}
```

#### func (*PartialEmojiConversionFailure) Error

```go
func (k *PartialEmojiConversionFailure) Error() string
```

#### func (*PartialEmojiConversionFailure) Is

```go
func (k *PartialEmojiConversionFailure) Is(target error) bool
```

### type Prefix

```go
type Prefix func(b *Bot, m *discord.MessageCreateEvent) []string
```

#### func Prefixes

```go
func Prefixes(p ...string) Prefix
```

#### func WhenMentioned

```go
func WhenMentioned() Prefix
```

#### func WhenMentionedOr

```go
func WhenMentionedOr(p ...string) Prefix
```

### type PrivateMessageOnly

```go
type PrivateMessageOnly struct {
	// contains filtered or unexported fields
}
```

#### func (*PrivateMessageOnly) Error

```go
func (k *PrivateMessageOnly) Error() string
```

#### func (*PrivateMessageOnly) Is

```go
func (k *PrivateMessageOnly) Is(target error) bool
```

### type RangeError

```go
type RangeError struct {
	Value            any
	Minimum, Maximum any
	// contains filtered or unexported fields
}
```

#### func (*RangeError) Error

```go
func (k *RangeError) Error() string
```

#### func (*RangeError) Is

```go
func (k *RangeError) Is(target error) bool
```

### type RoleRef

```go
type RoleRef struct {
	ID   snowflake.ID
	Name string
}
```

#### func RoleID

```go
func RoleID(id snowflake.ID) RoleRef
```

#### func RoleName

```go
func RoleName(name string) RoleRef
```

#### func (RoleRef) String

```go
func (r RoleRef) String() string
```

### type StringView

```go
type StringView struct {
	Index    int
	Previous int
	// contains filtered or unexported fields
}
```

#### func NewStringView

```go
func NewStringView(s string) *StringView
```

#### func (*StringView) Current

```go
func (v *StringView) Current() (rune, bool)
```

#### func (*StringView) EOF

```go
func (v *StringView) EOF() bool
```

#### func (*StringView) GetQuotedWord

```go
func (v *StringView) GetQuotedWord() (string, bool, error)
```

#### func (*StringView) GetWord

```go
func (v *StringView) GetWord() string
```

#### func (*StringView) Read

```go
func (v *StringView) Read(n int) string
```

#### func (*StringView) ReadRest

```go
func (v *StringView) ReadRest() string
```

#### func (*StringView) SkipString

```go
func (v *StringView) SkipString(s string) bool
```

#### func (*StringView) SkipWS

```go
func (v *StringView) SkipWS() bool
```

#### func (*StringView) Undo

```go
func (v *StringView) Undo()
```

### type TooManyArguments

```go
type TooManyArguments struct {
	// contains filtered or unexported fields
}
```

#### func (*TooManyArguments) Error

```go
func (k *TooManyArguments) Error() string
```

#### func (*TooManyArguments) Is

```go
func (k *TooManyArguments) Is(target error) bool
```

### type TooManyFlags

```go
type TooManyFlags struct {
	Flag   *Flag
	Values []string
	// contains filtered or unexported fields
}
```

#### func (*TooManyFlags) Error

```go
func (k *TooManyFlags) Error() string
```

#### func (*TooManyFlags) Is

```go
func (k *TooManyFlags) Is(target error) bool
```

### type UnexpectedQuoteError

```go
type UnexpectedQuoteError struct {
	Quote rune
	// contains filtered or unexported fields
}
```

#### func (*UnexpectedQuoteError) Error

```go
func (k *UnexpectedQuoteError) Error() string
```

#### func (*UnexpectedQuoteError) Is

```go
func (k *UnexpectedQuoteError) Is(target error) bool
```

### type UserInputError

```go
type UserInputError struct {
	// contains filtered or unexported fields
}
```

#### func (*UserInputError) Error

```go
func (k *UserInputError) Error() string
```

#### func (*UserInputError) Is

```go
func (k *UserInputError) Is(target error) bool
```

### type VerifyChecks

```go
type VerifyChecks int
```

```go
const (
	VerifyAlways VerifyChecks = iota
	VerifyNever
	VerifyInGuild
)
```


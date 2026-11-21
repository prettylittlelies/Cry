package commands

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

var (
	ErrCommand         = errors.New("command error")
	ErrUserInput       = errors.New("user input error")
	ErrBadArgument     = errors.New("bad argument")
	ErrCheckFailure    = errors.New("check failure")
	ErrArgumentParsing = errors.New("argument parsing error")
	ErrFlag            = errors.New("flag error")
	ErrExtension       = errors.New("extension error")
)

var (
	userInput  = []error{ErrCommand, ErrUserInput}
	badArg     = []error{ErrCommand, ErrUserInput, ErrBadArgument}
	checkFail  = []error{ErrCommand, ErrCheckFailure}
	parsing    = []error{ErrCommand, ErrUserInput, ErrArgumentParsing}
	flagErr    = []error{ErrCommand, ErrUserInput, ErrBadArgument, ErrFlag}
	commandErr = []error{ErrCommand}
	extension  = []error{ErrExtension}
)

type kinded struct {
	msg   string
	kinds []error
}

func (k *kinded) Error() string {
	return k.msg
}

func (k *kinded) Is(target error) bool {
	return slices.Contains(k.kinds, target)
}

func newKind(kinds []error, format string, args ...any) kinded {
	return kinded{msg: fmt.Sprintf(format, args...), kinds: kinds}
}

type CommandError struct{ kinded }

func NewCommandError(message string) *CommandError {
	return &CommandError{newKind(commandErr, "%s", message)}
}

type ConversionError struct {
	kinded
	Converter string
	Original  error
}

func (e *ConversionError) Unwrap() error { return e.Original }

type UserInputError struct{ kinded }

type CommandNotFound struct{ kinded }

type MissingRequiredArgument struct {
	kinded
	Param *Parameter
}

type MissingRequiredAttachment struct {
	kinded
	Param *Parameter
}

type TooManyArguments struct{ kinded }

type BadArgument struct{ kinded }

func NewBadArgument(format string, args ...any) *BadArgument {
	return &BadArgument{newKind(badArg, format, args...)}
}

type CheckFailure struct{ kinded }

func NewCheckFailure(message string) *CheckFailure {
	return &CheckFailure{newKind(checkFail, "%s", message)}
}

type CheckAnyFailure struct {
	kinded
	Errors []error
}

type PrivateMessageOnly struct{ kinded }

type NoPrivateMessage struct{ kinded }

type NotOwner struct{ kinded }

type NotFoundError struct {
	kinded
	Kind     string
	Argument string
}

func notFound(kind, argument string) *NotFoundError {
	return &NotFoundError{kinded: newKind(badArg, "%s %q not found.", kind, argument), Kind: kind, Argument: argument}
}

type ObjectNotFound struct {
	kinded
	Argument string
}

type ChannelNotReadable struct {
	kinded
	Channel string
}

type BadColourArgument struct {
	kinded
	Argument string
}

type BadInviteArgument struct {
	kinded
	Argument string
}

type PartialEmojiConversionFailure struct {
	kinded
	Argument string
}

type BadBoolArgument struct {
	kinded
	Argument string
}

type BadTimestampArgument struct {
	kinded
	Argument string
}

type RangeError struct {
	kinded
	Value            any
	Minimum, Maximum any
}

type DisabledCommand struct{ kinded }

type CommandInvokeError struct {
	kinded
	Original error
}

func (e *CommandInvokeError) Unwrap() error { return e.Original }

type CommandOnCooldown struct {
	kinded
	Cooldown   *Cooldown
	RetryAfter time.Duration
	Type       BucketType
}

type MaxConcurrencyReached struct {
	kinded
	Number int
	Per    BucketType
}

type MissingRole struct {
	kinded
	MissingRole string
}

type BotMissingRole struct {
	kinded
	MissingRole string
}

type MissingAnyRole struct {
	kinded
	MissingRoles []string
}

type BotMissingAnyRole struct {
	kinded
	MissingRoles []string
}

type NSFWChannelRequired struct {
	kinded
	Channel string
}

type MissingPermissions struct {
	kinded
	MissingPermissions []string
}

type BotMissingPermissions struct {
	kinded
	MissingPermissions []string
}

type BadUnionArgument struct {
	kinded
	Param      *Parameter
	Converters []string
	Errors     []error
}

type BadLiteralArgument struct {
	kinded
	Param    *Parameter
	Literals []string
	Errors   []error
	Argument string
}

type UnexpectedQuoteError struct {
	kinded
	Quote rune
}

type InvalidEndOfQuotedStringError struct {
	kinded
	Char rune
}

type ExpectedClosingQuoteError struct {
	kinded
	CloseQuote rune
}

type ExtensionError struct {
	kinded
	Name string
}

type ExtensionAlreadyLoaded struct{ ExtensionError }

type ExtensionNotLoaded struct{ ExtensionError }

type NoEntryPointError struct{ ExtensionError }

type ExtensionNotFound struct{ ExtensionError }

type ExtensionFailed struct {
	ExtensionError
	Original error
}

func (e *ExtensionFailed) Unwrap() error { return e.Original }

type CommandRegistrationError struct {
	Name          string
	AliasConflict bool
}

func (e *CommandRegistrationError) Error() string {
	kind := "command"
	if e.AliasConflict {
		kind = "alias"
	}
	return fmt.Sprintf("The %s %s is already an existing command or alias.", kind, e.Name)
}

type FlagError struct{ kinded }

type TooManyFlags struct {
	kinded
	Flag   *Flag
	Values []string
}

type BadFlagArgument struct {
	kinded
	Flag     *Flag
	Argument string
	Original error
}

func (e *BadFlagArgument) Unwrap() error { return e.Original }

type MissingRequiredFlag struct {
	kinded
	Flag *Flag
}

type MissingFlagArgument struct {
	kinded
	Flag *Flag
}

func displayed(p *Parameter) string {
	if p.DisplayedName != "" {
		return p.DisplayedName
	}
	return p.Name
}

func humanJoin(items []string, conj string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return items[0]
	case 2:
		return items[0] + " " + conj + " " + items[1]
	}
	return strings.Join(items[:len(items)-1], ", ") + ", " + conj + " " + items[len(items)-1]
}

func errMissingArgument(p *Parameter) *MissingRequiredArgument {
	return &MissingRequiredArgument{newKind(userInput, "%s is a required argument that is missing.", displayed(p)), p}
}

func errMissingAttachment(p *Parameter) *MissingRequiredAttachment {
	return &MissingRequiredAttachment{newKind(userInput, "%s is a required argument that is missing an attachment.", displayed(p)), p}
}

func errTooManyArguments(name string) *TooManyArguments {
	return &TooManyArguments{newKind(userInput, "Too many arguments passed to %s", name)}
}

func errCommandNotFound(invoked string) *CommandNotFound {
	return &CommandNotFound{newKind(commandErr, "Command %q is not found", invoked)}
}

func errDisabled(name string) *DisabledCommand {
	return &DisabledCommand{newKind(commandErr, "%s command is disabled", name)}
}

func errInvoke(err error) *CommandInvokeError {
	return &CommandInvokeError{newKind(commandErr, "Command raised an exception: %T: %v", err, err), err}
}

func errNoPrivateMessage() *NoPrivateMessage {
	return &NoPrivateMessage{newKind(checkFail, "This command cannot be used in private messages.")}
}

func errPrivateMessageOnly() *PrivateMessageOnly {
	return &PrivateMessageOnly{newKind(checkFail, "This command can only be used in private messages.")}
}

func errNotOwner() *NotOwner {
	return &NotOwner{newKind(checkFail, "You do not own this bot.")}
}

func errCheckAny(errs []error) *CheckAnyFailure {
	return &CheckAnyFailure{newKind(checkFail, "You do not have permission to run this command."), errs}
}

func errBadBool(argument string) *BadBoolArgument {
	return &BadBoolArgument{newKind(badArg, "%s is not a recognised boolean option", argument), argument}
}

func errObjectNotFound(argument string) *ObjectNotFound {
	return &ObjectNotFound{newKind(badArg, "%q does not follow a valid ID or mention format.", argument), argument}
}

func errBadColour(argument string) *BadColourArgument {
	return &BadColourArgument{newKind(badArg, "Colour %q is invalid.", argument), argument}
}

func errMissingPermissions(missing []string, bot bool) error {
	msg := "You are missing " + humanJoin(missing, "and") + " permission(s) to run this command."
	if bot {
		return &BotMissingPermissions{newKind(checkFail, "Bot requires %s permission(s) to run this command.", humanJoin(missing, "and")), missing}
	}
	return &MissingPermissions{newKind(checkFail, "%s", msg), missing}
}

func errMissingRole(role string, bot bool) error {
	if bot {
		return &BotMissingRole{newKind(checkFail, "Bot requires the role %s to run this command.", role), role}
	}
	return &MissingRole{newKind(checkFail, "Role %s is required to run this command.", role), role}
}

func errMissingAnyRole(roles []string, bot bool) error {
	if bot {
		return &BotMissingAnyRole{newKind(checkFail, "Bot is missing at least one of the required roles: %s", humanJoin(roles, "or")), roles}
	}
	return &MissingAnyRole{newKind(checkFail, "You are missing at least one of the required roles: %s", humanJoin(roles, "or")), roles}
}

func errNSFW(channel string) *NSFWChannelRequired {
	return &NSFWChannelRequired{newKind(checkFail, "Channel '%s' needs to be NSFW for this command to work.", channel), channel}
}

func errUnexpectedQuote(q rune) *UnexpectedQuoteError {
	return &UnexpectedQuoteError{newKind(parsing, "Unexpected quote mark, %q, in non-quoted string", q), q}
}

func errInvalidEndOfQuoted(c rune) *InvalidEndOfQuotedStringError {
	return &InvalidEndOfQuotedStringError{newKind(parsing, "Expected space after closing quotation but received %q", c), c}
}

func errExpectedClosingQuote(q rune) *ExpectedClosingQuoteError {
	return &ExpectedClosingQuoteError{newKind(parsing, "Expected closing %c.", q), q}
}

func errExtension(kind string, name string) ExtensionError {
	return ExtensionError{newKind(extension, kind, name), name}
}

func errPartialEmoji(argument string) *PartialEmojiConversionFailure {
	return &PartialEmojiConversionFailure{newKind(badArg, "Couldn't convert %q to PartialEmoji.", argument), argument}
}

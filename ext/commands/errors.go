package commands

import (
	"errors"
	"fmt"
)

var (
	errEndOfInput      = errors.New("commands: end of input")
	ErrCommandNotFound = errors.New("commands: command not found")
	ErrCheckFailure    = errors.New("commands: check failed")
	ErrGuildOnly       = fmt.Errorf("%w: guild only", ErrCheckFailure)
	ErrDMOnly          = fmt.Errorf("%w: direct messages only", ErrCheckFailure)
	ErrNotAllowed      = fmt.Errorf("%w: user not allowed", ErrCheckFailure)
)

type UnclosedQuoteError struct {
	Quote rune
}

func (e *UnclosedQuoteError) Error() string {
	return fmt.Sprintf("commands: expected closing %q", e.Quote)
}

type UnexpectedQuoteError struct {
	Quote rune
}

func (e *UnexpectedQuoteError) Error() string {
	return fmt.Sprintf("commands: expected space after closing %q", e.Quote)
}

type MissingArgumentError struct {
	Name string
}

func (e *MissingArgumentError) Error() string {
	return fmt.Sprintf("commands: missing argument %q", e.Name)
}

type BadArgumentError struct {
	Name  string
	Value string
	Err   error
}

func (e *BadArgumentError) Error() string {
	return fmt.Sprintf("commands: bad argument %q=%q: %v", e.Name, e.Value, e.Err)
}

func (e *BadArgumentError) Unwrap() error {
	return e.Err
}

type CommandError struct {
	Command string
	Err     error
}

func (e *CommandError) Error() string {
	return fmt.Sprintf("commands: %s: %v", e.Command, e.Err)
}

func (e *CommandError) Unwrap() error {
	return e.Err
}

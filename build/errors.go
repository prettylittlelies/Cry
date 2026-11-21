package build

import "fmt"

type ParseError struct {
	Kind  string
	Input string
	Err   error
}

func (e *ParseError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("build: invalid %s %q", e.Kind, e.Input)
	}
	return fmt.Sprintf("build: invalid %s %q: %v", e.Kind, e.Input, e.Err)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

type ExtractError struct {
	Source string
	Field  string
}

func (e *ExtractError) Error() string {
	return fmt.Sprintf("build: %s not found in %s", e.Field, e.Source)
}

type SourceError struct {
	Source  string
	Channel Channel
	Err     error
}

func (e *SourceError) Error() string {
	if e.Channel == "" {
		return fmt.Sprintf("build: %s: %v", e.Source, e.Err)
	}
	return fmt.Sprintf("build: %s (%s): %v", e.Source, e.Channel, e.Err)
}

func (e *SourceError) Unwrap() error {
	return e.Err
}

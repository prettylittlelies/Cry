package commands

import "strings"

type Handler func(ctx *Context, args *Args) error

type Command struct {
	Name        string
	Aliases     []string
	Help        string
	Usage       string
	Checks      []Check
	Run         Handler
	Subcommands []*Command
	parent      *Command
}

func (c *Command) QualifiedName() string {
	if c.parent == nil {
		return c.Name
	}
	return c.parent.QualifiedName() + " " + c.Name
}

func (c *Command) names() []string {
	return append([]string{c.Name}, c.Aliases...)
}

func (c *Command) sub(name string, fold bool) *Command {
	for _, s := range c.Subcommands {
		for _, n := range s.names() {
			if n == name || fold && strings.EqualFold(n, name) {
				return s
			}
		}
	}
	return nil
}

func (c *Command) link() {
	for _, s := range c.Subcommands {
		s.parent = c
		s.link()
	}
}

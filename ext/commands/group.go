package commands

import (
	"iter"
	"slices"
	"strings"
)

type group struct {
	all  map[string]*Command
	fold bool
}

func newGroup(fold bool) group {
	return group{all: map[string]*Command{}, fold: fold}
}

func (g *group) setFold(fold bool) {
	old := g.all
	g.fold, g.all = fold, make(map[string]*Command, len(old))
	for k, c := range old {
		g.all[g.key(k)] = c
	}
}

func (g *group) key(name string) string {
	if g.fold {
		return strings.ToLower(name)
	}
	return name
}

func (g *group) add(c *Command) error {
	if _, ok := g.all[g.key(c.Name)]; ok {
		return &CommandRegistrationError{Name: c.Name}
	}
	g.all[g.key(c.Name)] = c
	for _, alias := range c.Aliases {
		if _, ok := g.all[g.key(alias)]; ok {
			g.remove(c.Name)
			return &CommandRegistrationError{Name: alias, AliasConflict: true}
		}
		g.all[g.key(alias)] = c
	}
	return nil
}

func (g *group) remove(name string) *Command {
	c, ok := g.all[g.key(name)]
	if !ok {
		return nil
	}
	delete(g.all, g.key(name))
	if name != c.Name && g.key(name) != g.key(c.Name) {
		return c
	}
	for _, alias := range c.Aliases {
		if g.all[g.key(alias)] == c {
			delete(g.all, g.key(alias))
		}
	}
	return c
}

func (g *group) Get(name string) *Command {
	return g.all[g.key(name)]
}

func (g *group) GetCommand(name string) *Command {
	if !strings.Contains(name, " ") {
		return g.Get(name)
	}
	names := strings.Fields(name)
	cmd := g.Get(names[0])
	for _, n := range names[1:] {
		if cmd == nil {
			return nil
		}
		cmd = cmd.Get(n)
	}
	return cmd
}

func (g *group) Commands() []*Command {
	var out []*Command
	for _, c := range g.all {
		if !slices.Contains(out, c) {
			out = append(out, c)
		}
	}
	slices.SortFunc(out, func(a, b *Command) int { return strings.Compare(a.Name, b.Name) })
	return out
}

func (g *group) WalkCommands() iter.Seq[*Command] {
	return func(yield func(*Command) bool) {
		for _, c := range g.Commands() {
			if !yield(c) {
				return
			}
			for sub := range c.WalkCommands() {
				if !yield(sub) {
					return
				}
			}
		}
	}
}

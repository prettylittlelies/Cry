package commands

import (
	"slices"
	"strings"
)

type PaginatedHelp struct {
	HelpBase
	SortCommands    bool
	DMHelp          DMHelp
	DMHelpThreshold int
	NoCategory      string
	CommandsHeading string
	Paginator       Paginator
}

func newPaginatedHelp(paginator Paginator, commandsHeading string) PaginatedHelp {
	return PaginatedHelp{
		SortCommands:    true,
		DMHelpThreshold: 1000,
		NoCategory:      "No Category",
		CommandsHeading: commandsHeading,
		Paginator:       paginator,
	}
}

func (h *PaginatedHelp) PrepareHelpCommand(*Context, string) error {
	h.Paginator.Clear()
	return nil
}

func (h *PaginatedHelp) toAuthor() bool {
	return h.DMHelp == DMAlways || h.DMHelp == DMOverThreshold && h.Paginator.Len() > h.DMHelpThreshold
}

func (h *PaginatedHelp) Send(content string) error {
	var err error
	if h.toAuthor() {
		_, err = h.Context.SendAuthor(content)
	} else {
		_, err = h.Context.Send(content)
	}
	return err
}

func (h *PaginatedHelp) SendErrorMessage(message string) error {
	return h.Send(message)
}

func (h *PaginatedHelp) SendPages() error {
	for _, page := range h.Paginator.Pages() {
		if err := h.Send(page); err != nil {
			return err
		}
	}
	return nil
}

func (h *PaginatedHelp) addHelpText(text string) error {
	if text == "" {
		return nil
	}
	if h.Paginator.AddParagraph(text) == nil {
		return nil
	}
	for _, line := range strings.Split(text, "\n") {
		if err := h.Paginator.AddLine(line); err != nil {
			return err
		}
	}
	return h.Paginator.AddLine("")
}

func (h *PaginatedHelp) addNote(note string) error {
	if note == "" {
		return nil
	}
	if err := h.Paginator.AddLine(""); err != nil {
		return err
	}
	return h.Paginator.AddLine(note)
}

func (h *PaginatedHelp) addParagraph(text string) error {
	if text == "" {
		return nil
	}
	return h.Paginator.AddParagraph(text)
}

func (h *PaginatedHelp) categories(commands []*Command, suffix string) ([]string, map[string][]*Command) {
	category := func(c *Command) string {
		if c.Cog != nil {
			return c.Cog.CogName() + suffix
		}
		return "​" + h.NoCategory + suffix
	}
	filtered := h.FilterCommands(commands, true, category)
	var order []string
	groups := map[string][]*Command{}
	for _, c := range filtered {
		key := category(c)
		if _, seen := groups[key]; !seen {
			order = append(order, key)
		}
		groups[key] = append(groups[key], c)
	}
	if h.SortCommands {
		for _, cmds := range groups {
			sortByName(cmds)
		}
	}
	return order, groups
}

func sortByName(commands []*Command) {
	slices.SortStableFunc(commands, func(a, b *Command) int { return strings.Compare(a.Name, b.Name) })
}

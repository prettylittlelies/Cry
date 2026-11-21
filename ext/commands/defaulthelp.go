package commands

import (
	"strings"
)

type DefaultHelp struct {
	PaginatedHelp
	Width                      int
	Indent                     int
	ArgumentsHeading           string
	DefaultArgumentDescription string
	ShowParameterDescriptions  bool
}

func NewDefaultHelp() *DefaultHelp {
	return &DefaultHelp{
		PaginatedHelp:              newPaginatedHelp(NewPaginator("```", "```"), "Commands:"),
		Width:                      80,
		Indent:                     2,
		ArgumentsHeading:           "Arguments:",
		DefaultArgumentDescription: "No description given",
		ShowParameterDescriptions:  true,
	}
}

func (h *DefaultHelp) ShortenText(text string) string {
	if r := []rune(text); len(r) > h.Width {
		return strings.TrimRight(string(r[:h.Width-3]), " \t\n\r\v\f") + "..."
	}
	return text
}

func (h *DefaultHelp) EndingNote() string {
	prefix, name := h.Context.CleanPrefix(), h.InvokedWith()
	return "Type " + prefix + name + " command for more info on a command.\n" +
		"You can also type " + prefix + name + " category for more info on a category."
}

func (h *DefaultHelp) CommandSignature(command *Command) string {
	if !h.ShowParameterDescriptions {
		return h.HelpBase.CommandSignature(command)
	}
	name := command.Name
	if len(command.Aliases) > 0 {
		name = "[" + command.Name + "|" + strings.Join(command.Aliases, "|") + "]"
	}
	return h.Context.CleanPrefix() + name
}

func (h *DefaultHelp) pad(name string, maxSize int) string {
	width := maxSize - (StringWidth(name) - runes(name))
	return strings.Repeat(" ", h.Indent) + name + strings.Repeat(" ", max(0, width-runes(name)))
}

func (h *DefaultHelp) AddIndentedCommands(commands []*Command, heading string, maxSize int) error {
	if len(commands) == 0 {
		return nil
	}
	if err := h.Paginator.AddLine(heading); err != nil {
		return err
	}
	if maxSize == 0 {
		maxSize = MaxWidth(commandNames(commands))
	}
	for _, c := range commands {
		if err := h.Paginator.AddLine(h.ShortenText(h.pad(c.Name, maxSize) + " " + c.ShortDoc())); err != nil {
			return err
		}
	}
	return nil
}

func (h *DefaultHelp) AddCommandArguments(command *Command) error {
	if len(command.Params) == 0 {
		return nil
	}
	if err := h.Paginator.AddLine(h.ArgumentsHeading); err != nil {
		return err
	}
	names := make([]string, len(command.Params))
	for i, p := range command.Params {
		names[i] = p.Name
	}
	maxSize := MaxWidth(names)
	for _, p := range command.Params {
		name := p.DisplayedName
		if name == "" {
			name = p.Name
		}
		description := p.Description
		if description == "" {
			description = h.DefaultArgumentDescription
		}
		entry := h.ShortenText(h.pad(name, maxSize) + " " + description)
		if d, ok := p.DisplayDefault(); ok {
			entry += " (default: " + d + ")"
		}
		if err := h.Paginator.AddLine(entry); err != nil {
			return err
		}
	}
	return nil
}

func (h *DefaultHelp) AddCommandFormatting(command *Command) error {
	if err := h.addParagraph(command.Description); err != nil {
		return err
	}
	if err := h.Paginator.AddParagraph(h.CommandSignature(command)); err != nil {
		return err
	}
	if err := h.addHelpText(command.Help); err != nil {
		return err
	}
	if h.ShowParameterDescriptions {
		return h.AddCommandArguments(command)
	}
	return nil
}

func (h *DefaultHelp) SendBotHelp(map[Cog][]*Command) error {
	if err := h.addParagraph(h.Context.Bot.Description); err != nil {
		return err
	}
	order, groups := h.categories(h.Context.Bot.Commands(), ":")
	var all []string
	for _, cmds := range groups {
		all = append(all, commandNames(cmds)...)
	}
	maxSize := MaxWidth(all)
	for _, category := range order {
		if err := h.AddIndentedCommands(groups[category], category, maxSize); err != nil {
			return err
		}
	}
	if err := h.addNote(h.EndingNote()); err != nil {
		return err
	}
	return h.SendPages()
}

func (h *DefaultHelp) SendCommandHelp(command *Command) error {
	if err := h.AddCommandFormatting(command); err != nil {
		return err
	}
	h.Paginator.ClosePage()
	return h.SendPages()
}

func (h *DefaultHelp) SendGroupHelp(group *Command) error {
	if err := h.AddCommandFormatting(group); err != nil {
		return err
	}
	filtered := h.FilterCommands(group.Commands(), h.SortCommands, nil)
	if err := h.AddIndentedCommands(filtered, h.CommandsHeading, 0); err != nil {
		return err
	}
	if len(filtered) > 0 {
		if err := h.addNote(h.EndingNote()); err != nil {
			return err
		}
	}
	return h.SendPages()
}

func (h *DefaultHelp) SendCogHelp(cog Cog) error {
	if d, ok := cog.(CogDescriber); ok {
		if err := h.addParagraph(d.CogDescription()); err != nil {
			return err
		}
	}
	filtered := h.FilterCommands(topLevel(cog), h.SortCommands, nil)
	if err := h.AddIndentedCommands(filtered, h.CommandsHeading, 0); err != nil {
		return err
	}
	if err := h.addNote(h.EndingNote()); err != nil {
		return err
	}
	return h.SendPages()
}

func topLevel(cog Cog) []*Command {
	var out []*Command
	for _, c := range cog.CogCommands() {
		if c.Parent == nil {
			out = append(out, c)
		}
	}
	return out
}

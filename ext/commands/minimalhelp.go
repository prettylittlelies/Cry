package commands

import "strings"

type MinimalHelp struct {
	PaginatedHelp
	AliasesHeading string
}

func NewMinimalHelp() *MinimalHelp {
	return &MinimalHelp{
		PaginatedHelp:  newPaginatedHelp(NewPaginator("", ""), "Commands"),
		AliasesHeading: "Aliases:",
	}
}

func (h *MinimalHelp) OpeningNote() string {
	prefix, name := h.Context.CleanPrefix(), h.InvokedWith()
	return "Use `" + prefix + name + " [command]` for more info on a command.\n" +
		"You can also use `" + prefix + name + " [category]` for more info on a category."
}

func (h *MinimalHelp) EndingNote() string { return "" }

func (h *MinimalHelp) CommandSignature(command *Command) string {
	return h.Context.CleanPrefix() + command.QualifiedName() + " " + command.Signature()
}

func (h *MinimalHelp) AddBotCommandsFormatting(commands []*Command, heading string) error {
	if len(commands) == 0 {
		return nil
	}
	if err := h.Paginator.AddLine("__**" + heading + "**__"); err != nil {
		return err
	}
	return h.Paginator.AddLine(strings.Join(commandNames(commands), " "))
}

func (h *MinimalHelp) AddSubcommandFormatting(command *Command) error {
	line := h.Context.CleanPrefix() + command.QualifiedName()
	if doc := command.ShortDoc(); doc != "" {
		line += " – " + doc
	}
	return h.Paginator.AddLine(line)
}

func (h *MinimalHelp) AddCommandFormatting(command *Command) error {
	if err := h.addParagraph(command.Description); err != nil {
		return err
	}
	signature := h.CommandSignature(command)
	if len(command.Aliases) > 0 {
		if err := h.Paginator.AddLine(signature); err != nil {
			return err
		}
		if err := h.Paginator.AddParagraph("**" + h.AliasesHeading + "** " + strings.Join(command.Aliases, ", ")); err != nil {
			return err
		}
	} else if err := h.Paginator.AddParagraph(signature); err != nil {
		return err
	}
	return h.addHelpText(command.Help)
}

func (h *MinimalHelp) addSubcommands(heading string, commands []*Command) error {
	if err := h.Paginator.AddLine(heading); err != nil {
		return err
	}
	for _, c := range commands {
		if err := h.AddSubcommandFormatting(c); err != nil {
			return err
		}
	}
	return h.addNote(h.EndingNote())
}

func (h *MinimalHelp) SendBotHelp(map[Cog][]*Command) error {
	if err := h.addParagraph(h.Context.Bot.Description); err != nil {
		return err
	}
	if err := h.addParagraph(h.OpeningNote()); err != nil {
		return err
	}
	order, groups := h.categories(h.Context.Bot.Commands(), "")
	for _, category := range order {
		if err := h.AddBotCommandsFormatting(groups[category], category); err != nil {
			return err
		}
	}
	if err := h.addNote(h.EndingNote()); err != nil {
		return err
	}
	return h.SendPages()
}

func (h *MinimalHelp) SendCogHelp(cog Cog) error {
	if err := h.addParagraph(h.Context.Bot.Description); err != nil {
		return err
	}
	if err := h.addParagraph(h.OpeningNote()); err != nil {
		return err
	}
	if d, ok := cog.(CogDescriber); ok {
		if err := h.addParagraph(d.CogDescription()); err != nil {
			return err
		}
	}
	if filtered := h.FilterCommands(topLevel(cog), h.SortCommands, nil); len(filtered) > 0 {
		if err := h.addSubcommands("**"+cog.CogName()+" "+h.CommandsHeading+"**", filtered); err != nil {
			return err
		}
	}
	return h.SendPages()
}

func (h *MinimalHelp) SendGroupHelp(group *Command) error {
	if err := h.AddCommandFormatting(group); err != nil {
		return err
	}
	if filtered := h.FilterCommands(group.Commands(), h.SortCommands, nil); len(filtered) > 0 {
		if err := h.addParagraph(h.OpeningNote()); err != nil {
			return err
		}
		if err := h.addSubcommands("**"+h.CommandsHeading+"**", filtered); err != nil {
			return err
		}
	}
	return h.SendPages()
}

func (h *MinimalHelp) SendCommandHelp(command *Command) error {
	if err := h.AddCommandFormatting(command); err != nil {
		return err
	}
	h.Paginator.ClosePage()
	return h.SendPages()
}

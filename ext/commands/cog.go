package commands

type Cog interface {
	Name() string
	Commands() []*Command
}

type Loader interface {
	Load(b *Bot) error
}

type Unloader interface {
	Unload(b *Bot)
}

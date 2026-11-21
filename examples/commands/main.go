package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/commands"
)

type echoParams struct {
	Times int    `default:"1" min:"1" max:"5" description:"how many times to repeat"`
	Text  string `commands:"rest" description:"what to say"`
}

type whoisParams struct {
	Member discord.GuildMember `description:"a mention, id or name"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c, err := cry.New(ctx, os.Getenv("CRY_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot := commands.NewBot(c, commands.WithPrefix(commands.Prefixes(".")), commands.WithMode(commands.SelfBot))

	bot.AddCommand(commands.New("ping", func(ctx *commands.Context, _ struct{}) error {
		_, err := ctx.Reply("pong")
		return err
	}, commands.Help("Checks that the bot is alive.")))

	bot.AddCommand(commands.New("echo", func(ctx *commands.Context, p echoParams) error {
		for range p.Times {
			if _, err := ctx.Send(p.Text); err != nil {
				return err
			}
		}
		return nil
	}, commands.Help("Repeats your text.")))

	bot.AddCommand(commands.New("whois", func(ctx *commands.Context, p whoisParams) error {
		_, err := ctx.Send(fmt.Sprintf("%s joined at %s", p.Member.User.Username, p.Member.JoinedAt.Format("2006-01-02")))
		return err
	}, commands.Help("Shows when a member joined.")))

	if err := c.Run(ctx); err != nil && ctx.Err() == nil {
		log.Fatal(err)
	}
}

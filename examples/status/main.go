package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/ext/tasks"
	"github.com/prettylittlelies/cry/gateway"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c, err := cry.New(ctx, os.Getenv("CRY_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	clock := tasks.Every(time.Minute, func(ctx context.Context) error {
		state := fmt.Sprintf("it is %s", time.Now().Format("15:04"))
		return c.Gateway().SetPresence(ctx, gateway.Presence{
			Status:     discord.StatusTypeOnline,
			Activities: []discord.Activity{{Name: "Custom Status", Type: discord.ActivityTypeCustom, State: &state}},
		})
	})

	cry.On(c, func(ctx context.Context, e *discord.ReadyEvent) {
		log.Printf("ready as %s", e.User.Username)
		clock.Start(ctx)
	})

	if err := c.Run(ctx); err != nil && ctx.Err() == nil {
		log.Fatal(err)
	}
}

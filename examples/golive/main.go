package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
	"github.com/prettylittlelies/cry/voice"
)

func id(raw string) snowflake.ID {
	n, err := strconv.ParseUint(raw, 10, 64)
	if err != nil {
		log.Fatalf("bad id %q", raw)
	}
	return snowflake.ID(n)
}

func main() {
	guild := flag.String("guild", "", "guild id")
	channel := flag.String("channel", "", "voice channel id")
	video := flag.String("ivf", "", "vp8 ivf file to stream")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	c, err := cry.New(ctx, os.Getenv("CRY_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	ready := make(chan struct{})
	var once sync.Once
	cry.On(c, func(context.Context, *discord.ReadyEvent) { once.Do(func() { close(ready) }) })
	go c.Run(ctx)
	<-ready

	conn, err := voice.Join(ctx, c, id(*guild), id(*channel), voice.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	stream, err := conn.StartStream(ctx, voice.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close(context.Background())

	f, err := os.Open(*video)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := stream.PlayIVF(ctx, f, voice.VideoOptions{}); err != nil {
		log.Fatal(err)
	}
}

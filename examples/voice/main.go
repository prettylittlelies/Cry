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
	audio := flag.String("ogg", "", "ogg/opus file to play")
	camera := flag.String("ivf", "", "vp8 ivf file to show as camera")
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

	opts := voice.Options{OnOpus: func(user snowflake.ID, frame []byte) {}}
	if *camera != "" {
		opts.Video = voice.Camera
	}
	conn, err := voice.Join(ctx, c, id(*guild), id(*channel), opts)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	log.Printf("joined, end-to-end encrypted: %v", conn.Encrypted())

	if *audio != "" {
		f, err := os.Open(*audio)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := conn.Play(ctx, voice.NewOggOpus(f)); err != nil {
			log.Fatal(err)
		}
	}
	if *camera != "" {
		f, err := os.Open(*camera)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if err := conn.PlayIVF(ctx, f, voice.VideoOptions{}); err != nil {
			log.Fatal(err)
		}
	}
}

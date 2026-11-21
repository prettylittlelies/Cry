# cry

A Go library for Discord user accounts, modelled on [discord.py-self](https://github.com/dolfies/discord.py-self).
It speaks the same gateway, REST and voice protocols as the desktop client, including Discord's
end-to-end encrypted voice and video (DAVE).

> Automating a user account is against Discord's Terms of Service and can get the account banned.
> Use it on accounts you own and are prepared to lose.

```
go get github.com/prettylittlelies/cry
```

Requires Go 1.27 or newer.

## What's in it

| Package | Purpose |
| --- | --- |
| [`cry`](docs/cry.md) | The client: login, events, cache, and handles for channels, messages, guilds, members and users |
| [`discord`](docs/discord.md) | Typed models, events and enums, generated from [discord-userdoccers](https://github.com/discord-userdoccers/discord-userdoccers) |
| [`api`](docs/api.md) | One typed method per REST endpoint, also generated |
| [`gateway`](docs/gateway.md) | Gateway connection, resume, presence, member requests, guild subscriptions, voice state |
| [`state`](docs/state.md) | Cache of guilds, channels, members, messages, presences, relationships and voice states |
| [`ext/commands`](docs/ext-commands.md) | Command framework ported from discord.py's `ext.commands`: converters, flags, checks, cooldowns, cogs, help |
| [`ext/tasks`](docs/ext-tasks.md) | Repeating background tasks, like discord.py's `ext.tasks` |
| [`voice`](docs/voice.md) | Voice and video: audio playback and receiving, camera video, Go Live screen share |
| [`voice/dave`](docs/voice-dave.md) | DAVE end-to-end encryption for voice and video |
| [`mls`](docs/mls.md) | Messaging Layer Security (RFC 9420), the group key agreement DAVE is built on |

The full API reference is generated into [`docs/`](docs/README.md).

## Quick start

Set `CRY_TOKEN` to the account's token, then:

```go
c, err := cry.New(ctx, os.Getenv("CRY_TOKEN"))
if err != nil {
	log.Fatal(err)
}

cry.On(c, func(ctx context.Context, e *discord.MessageCreateEvent) {
	if e.Content == "!ping" && e.Author.ID == c.State().Me().ID {
		c.Message(e.ChannelID, e.ID).Reply(ctx, "pong")
	}
})

log.Fatal(c.Run(ctx))
```

`cry.On` takes any event type from the `discord` package. `cry.WaitFor` waits for the next event that matches a check.

## Working with objects

Handles wrap the REST API the way discord.py's models do. They don't fetch anything until you call a method.

```go
ch := c.Channel(channelID)
ch.Send(ctx, "hello")
for m, err := range ch.History(ctx, cry.HistoryOptions{Limit: 200}) { ... }

msg := c.Message(channelID, messageID)
msg.React(ctx, "👋")
msg.Edit(ctx, "edited")

c.Member(guildID, userID).AddRoles(ctx, roleID)
members, _ := c.Guild(guildID).QueryMembers(ctx, "ali", 10, false)
c.User(userID).Send(ctx, "a direct message")
```

Every endpoint is also available directly through `c.API()`.

## Status and presence

```go
state := "listening to something"
c.Gateway().SetPresence(ctx, gateway.Presence{
	Status:     discord.StatusTypeOnline,
	Activities: []discord.Activity{{Name: "Custom Status", Type: discord.ActivityTypeCustom, State: &state}},
})
```

See [`examples/status`](examples/status/main.go) for a status that updates every minute.

## Commands

```go
bot := commands.NewBot(c, commands.WithPrefix(commands.Prefixes(".")), commands.WithMode(commands.SelfBot))

type echoParams struct {
	Times int    `default:"1" min:"1" max:"5"`
	Text  string `commands:"rest"`
}

bot.AddCommand(commands.New("echo", func(ctx *commands.Context, p echoParams) error {
	_, err := ctx.Send(strings.Repeat(p.Text+" ", p.Times))
	return err
}))
```

A command's parameters are the fields of a struct. Field types choose the converter (members, users, channels,
roles, messages, emojis, colours, timestamps and more), and tags give defaults, ranges, rest and greedy arguments.
`help` is built in. See [`examples/commands`](examples/commands/main.go).

## Voice and video

```go
conn, err := voice.Join(ctx, c, guildID, channelID, voice.Options{
	OnOpus: func(user snowflake.ID, frame []byte) { /* incoming audio */ },
})
defer conn.Close(ctx)

f, _ := os.Open("song.ogg")
conn.Play(ctx, voice.NewOggOpus(f))
```

Calls are end-to-end encrypted with DAVE whenever Discord requires it, and `conn.Encrypted()` reports when it's active.
For camera video, join with `voice.Options{Video: voice.Camera}` and call `conn.PlayIVF`. For Go Live, call
`conn.StartStream` and play into the returned connection. See [`examples/voice`](examples/voice/main.go) and
[`examples/golive`](examples/golive/main.go).

The library sends already-encoded media. Prepare files with ffmpeg:

```
ffmpeg -i in.mp3 -c:a libopus -b:a 96k -ar 48000 -ac 2 out.ogg
ffmpeg -i in.mp4 -an -c:v libvpx -b:v 1500k -deadline realtime -g 30 out.ivf
```

`voice.NewOggWriter` saves received audio as `.ogg`.

## Testing

```
go test ./tests/...
```

- **MLS** is checked against the official [MLS interop test vectors](https://github.com/mlswg/mls-implementations).
- **DAVE** can be checked against Discord's own [libdave](https://github.com/discord/libdave):
  build the harness with [`tools/daveinterop/build.sh`](tools/daveinterop/build.sh), then run the tests with
  `CRY_LIBDAVE_HARNESS` pointing at it.
- **Live tests** run against real Discord and only build with the `live` tag. They read the token from
  `CRY_TOKEN` or `tests/live/.token`:

  ```
  go test -tags live ./tests/live -run Voice -v
  ```

  `CRY_VOICE_GUILD` and `CRY_VOICE_CHANNEL` pick the channel. `CRY_VOICE_OGG`, `CRY_VOICE_CAMERA`,
  `CRY_VOICE_STREAM` and `CRY_VOICE_RECORD` turn on audio, camera, Go Live and recording.

## Regenerating code and docs

```
go run ./tools/gendiscord <discord-userdoccers>/pages discord api
go run ./tools/gendocs . docs
```

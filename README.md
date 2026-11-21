[![cry](.github/assets/banner.svg)](https://github.com/prettylittlelies/Cry)

<p align="center"><b>A Go library for Discord user accounts, modelled on discord.py-self.</b></p>

<p align="center">
  <a href="https://github.com/prettylittlelies/Cry"><img src="https://img.shields.io/github/languages/code-size/prettylittlelies/Cry?style=for-the-badge&labelColor=0a0a0a&color=cfff53&label=code%20size" alt="Code size"></a>
  <a href="https://github.com/prettylittlelies/Cry/stargazers"><img src="https://img.shields.io/github/stars/prettylittlelies/Cry?style=for-the-badge&labelColor=0a0a0a&color=cfff53&logo=github&logoColor=white" alt="Stars"></a>
  <br>
  <a href="https://go.dev"><img src="https://img.shields.io/badge/Go_1.27-000000?style=for-the-badge&logo=go&logoColor=00add8" alt="Go"></a>
  <a href="https://discord.com"><img src="https://img.shields.io/badge/Discord-000000?style=for-the-badge&logo=discord&logoColor=5865f2" alt="Discord"></a>
</p>

<p align="center">
  <a href="#features">Features</a> · <a href="#install">Install</a> · <a href="#usage">Usage</a> · <a href="#voice-and-video">Voice and video</a> · <a href="docs/README.md">API reference</a> · <a href="#tests">Tests</a>
</p>

---

Automating a user account is against Discord's Terms of Service and can get it banned. Use it on accounts you own and can afford to lose.

## Features

- **Gateway, REST and cache.** One typed method per REST endpoint and typed models for every event, generated from [discord-userdoccers](https://github.com/discord-userdoccers/discord-userdoccers).
- **Handles.** `c.Channel(id).Send`, `c.Message(ch, id).Reply`, `c.Member(g, u).AddRoles`, the way discord.py's models work.
- **Commands.** A port of discord.py's `ext.commands`: converters, flags, checks, cooldowns, cogs and help.
- **Tasks.** Repeating background jobs, like `ext.tasks`.
- **Voice.** Play and record audio, with Discord's end-to-end encryption (DAVE) handled for you.
- **Video.** Camera video and Go Live screen share from VP8 files.
- **MLS.** A full RFC 9420 implementation, checked against the official interop vectors and Discord's libdave.

## Install

Requires Go 1.27 or newer.

```sh
go get github.com/prettylittlelies/cry
```

Set `CRY_TOKEN` to the account's token before running anything below.

## Usage

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

### Handles

```go
c.Channel(channelID).Send(ctx, "hello")
for m, err := range c.Channel(channelID).History(ctx, cry.HistoryOptions{Limit: 200}) { ... }
c.Message(channelID, messageID).React(ctx, "👋")
c.Member(guildID, userID).AddRoles(ctx, roleID)
c.User(userID).Send(ctx, "a direct message")
```

Every endpoint is also on `c.API()`.

### Status

```go
state := "listening to something"
c.Gateway().SetPresence(ctx, gateway.Presence{
	Status:     discord.StatusTypeOnline,
	Activities: []discord.Activity{{Name: "Custom Status", Type: discord.ActivityTypeCustom, State: &state}},
})
```

### Commands

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

A command's parameters are the fields of a struct. The field type picks the converter, and tags set defaults, ranges, and rest or greedy arguments.

## Voice and video

```go
conn, err := voice.Join(ctx, c, guildID, channelID, voice.Options{})
defer conn.Close(ctx)

f, _ := os.Open("song.ogg")
conn.Play(ctx, voice.NewOggOpus(f))
```

| Want | Do |
| --- | --- |
| Play audio | `conn.Play` with `voice.NewOggOpus` |
| Hear others | `voice.Options{OnOpus: ...}`, save with `voice.NewOggWriter` |
| Camera | Join with `voice.Options{Video: voice.Camera}`, then `conn.PlayIVF` |
| Go Live | `conn.StartStream`, then `PlayIVF` on the returned connection |
| Check encryption | `conn.Encrypted()` and `conn.EpochAuthenticator()` |

Media goes in already encoded. Convert with ffmpeg:

```sh
ffmpeg -i in.mp3 -c:a libopus -b:a 96k -ar 48000 -ac 2 out.ogg
ffmpeg -i in.mp4 -an -c:v libvpx -b:v 1500k -deadline realtime -g 30 out.ivf
```

Full programs are in [`examples/`](examples).

## Tests

```sh
go test ./tests/...
```

Live tests run against real Discord and read the token from `CRY_TOKEN` or `tests/live/.token`:

```sh
go test -tags live ./tests/live -run Voice -v
```

`CRY_VOICE_GUILD` and `CRY_VOICE_CHANNEL` pick the channel. `CRY_VOICE_OGG`, `CRY_VOICE_CAMERA`, `CRY_VOICE_STREAM` and `CRY_VOICE_RECORD` turn on audio, camera, Go Live and recording.

To check DAVE against Discord's own [libdave](https://github.com/discord/libdave), build the harness with [`tools/daveinterop/build.sh`](tools/daveinterop/build.sh) and set `CRY_LIBDAVE_HARNESS` to it.

Regenerate the models and docs with:

```sh
go run ./tools/gendiscord <discord-userdoccers>/pages discord api
go run ./tools/gendocs . docs
```

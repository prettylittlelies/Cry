package voice

import (
	"context"
	"fmt"

	"github.com/prettylittlelies/cry"
	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/gateway"
)

func (v *Conn) StreamKey() string {
	if v.guild != 0 {
		return fmt.Sprintf("guild:%d:%d:%d", v.guild, v.channel, v.self)
	}
	return fmt.Sprintf("call:%d:%d", v.channel, v.self)
}

func (v *Conn) StartStream(ctx context.Context, opts Options) (*Conn, error) {
	if v.stream != "" {
		return nil, fmt.Errorf("voice: cannot stream from a stream connection")
	}
	key := v.StreamKey()
	created := make(chan *discord.StreamCreateEvent, 1)
	servers := make(chan *discord.StreamServerUpdateEvent, 1)
	deleted := make(chan *discord.StreamDeleteEvent, 1)
	defer cry.On(v.client, func(_ context.Context, e *discord.StreamCreateEvent) {
		if e.StreamKey == key {
			trySend(created, e)
		}
	})()
	defer cry.On(v.client, func(_ context.Context, e *discord.StreamServerUpdateEvent) {
		if e.StreamKey == key {
			trySend(servers, e)
		}
	})()
	defer cry.On(v.client, func(_ context.Context, e *discord.StreamDeleteEvent) {
		if e.StreamKey == key {
			trySend(deleted, e)
		}
	})()

	req := gateway.StreamCreate{Type: discord.StreamTypeCall, ChannelID: v.channel}
	if v.guild != 0 {
		req.Type, req.GuildID = discord.StreamTypeGuild, &v.guild
	}
	if len(v.opts.PreferredRegions) > 0 {
		req.PreferredRegion = v.opts.PreferredRegions[0]
	}
	if err := v.client.Gateway().StreamCreate(ctx, req); err != nil {
		return nil, err
	}
	var create *discord.StreamCreateEvent
	var server *discord.StreamServerUpdateEvent
	for create == nil || server == nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case d := <-deleted:
			return nil, fmt.Errorf("voice: stream refused: %s", d.Reason)
		case create = <-created:
		case server = <-servers:
		}
	}
	if server.Endpoint == nil {
		return nil, fmt.Errorf("voice: stream server has no endpoint yet")
	}
	if opts.Video == "" {
		opts.Video = Screen
	}
	opts.Signer = v.opts.Signer
	opts.DisableDAVE = v.opts.DisableDAVE
	s, err := newConn(v.client, v.guild, v.channel, opts)
	if err != nil {
		return nil, err
	}
	s.stream = key
	s.identServer, s.identChannel = create.RTCServerID, create.RTCChannelID
	s.daveGroup = create.RTCServerID - 1
	s.sessionID, s.token, s.endpoint = v.sessionID, server.Token, *server.Endpoint
	if err := s.start(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (v *Conn) stopStream(ctx context.Context) error {
	return v.client.Gateway().StreamDelete(ctx, v.stream)
}

func trySend[T any](ch chan T, v T) {
	select {
	case ch <- v:
	default:
	}
}

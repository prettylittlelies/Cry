package gateway

import (
	"context"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

type StreamCreate struct {
	Type            discord.StreamType `json:"type"`
	GuildID         *snowflake.ID      `json:"guild_id"`
	ChannelID       snowflake.ID       `json:"channel_id"`
	PreferredRegion string             `json:"preferred_region,omitempty"`
}

type streamKey struct {
	StreamKey string `json:"stream_key"`
}

type streamPaused struct {
	StreamKey string `json:"stream_key"`
	Paused    bool   `json:"paused"`
}

func (g *Gateway) UpdateVoiceState(ctx context.Context, v discord.UpdateVoiceState) error {
	if v.ChannelID != nil && len(v.PreferredRegions) > 0 && v.PreferredRegion == nil {
		v.PreferredRegion = &v.PreferredRegions[0]
	}
	return g.Send(ctx, OpVoiceStateUpdate, v)
}

func (g *Gateway) CallConnect(ctx context.Context, channel snowflake.ID) error {
	return g.Send(ctx, OpCallConnect, discord.RequestCallConnect{ChannelID: channel})
}

func (g *Gateway) StreamCreate(ctx context.Context, s StreamCreate) error {
	return g.Send(ctx, OpStreamCreate, s)
}

func (g *Gateway) StreamWatch(ctx context.Context, key string) error {
	return g.Send(ctx, OpStreamWatch, streamKey{key})
}

func (g *Gateway) StreamPing(ctx context.Context, key string) error {
	return g.Send(ctx, OpStreamPing, streamKey{key})
}

func (g *Gateway) StreamDelete(ctx context.Context, key string) error {
	return g.Send(ctx, OpStreamDelete, streamKey{key})
}

func (g *Gateway) StreamSetPaused(ctx context.Context, key string, paused bool) error {
	return g.Send(ctx, OpStreamSetPaused, streamPaused{key, paused})
}

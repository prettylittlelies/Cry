package api

import (
	"context"

	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/snowflake"
)

func (a *API) DeleteOwnReaction(ctx context.Context, channelID, messageID snowflake.ID, emoji string, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/{message_id}/reactions/{emoji}/@me", rest.P("channel_id", channelID), rest.P("message_id", messageID), rest.P("emoji", emoji))}
	return none(ctx, a, req, opts)
}

func (a *API) DeleteUserReaction(ctx context.Context, channelID, messageID snowflake.ID, emoji string, userID snowflake.ID, opts ...Option) error {
	req := &rest.Request{Route: rest.NewRoute("DELETE", "/channels/{channel_id}/messages/{message_id}/reactions/{emoji}/{user_id}", rest.P("channel_id", channelID), rest.P("message_id", messageID), rest.P("emoji", emoji), rest.P("user_id", userID))}
	return none(ctx, a, req, opts)
}

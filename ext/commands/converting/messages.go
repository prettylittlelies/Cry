package converting

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/rest"
	"github.com/prettylittlelies/cry/snowflake"
)

var (
	ErrNotFound           = errors.New("converting: not found")
	ErrChannelNotFound    = errors.New("converting: channel not found")
	ErrChannelNotReadable = errors.New("converting: channel not readable")
)

var (
	messageIDPattern   = regexp.MustCompile(`^(?:([0-9]{15,20})-)?([0-9]{15,20})$`)
	messageLinkPattern = regexp.MustCompile(`^https?://(?:(ptb|canary|www)\.)?discord(?:app)?\.com/channels/([0-9]{15,20}|@me)/([0-9]{15,20})/([0-9]{15,20})/?$`)
)

type PartialMessage struct {
	ID      snowflake.ID
	Channel discord.Channel
}

type messageRef struct {
	guild, channel, message snowflake.ID
	explicitGuild           bool
}

func (s Scope) parseMessageRef(argument string, current snowflake.ID) (messageRef, bool) {
	ref := messageRef{guild: s.GuildID, channel: current}
	if m := messageIDPattern.FindStringSubmatch(argument); m != nil {
		if m[1] != "" {
			ref.channel, _ = parseID(m[1])
		}
		ref.message, _ = parseID(m[2])
		return ref, true
	}
	m := messageLinkPattern.FindStringSubmatch(argument)
	if m == nil {
		return ref, false
	}
	ref.explicitGuild, ref.guild = true, 0
	if m[2] != "@me" {
		ref.guild, _ = parseID(m[2])
	}
	ref.channel, _ = parseID(m[3])
	ref.message, _ = parseID(m[4])
	return ref, true
}

func (s Scope) messageChannel(ref messageRef) (discord.Channel, bool) {
	st := s.Client.State()
	if ref.explicitGuild && ref.guild != 0 {
		if _, ok := st.Guild(ref.guild); !ok {
			return discord.Channel{}, false
		}
	}
	c, ok := st.Channel(ref.channel)
	if !ok || ref.explicitGuild && c.GuildID != ref.guild {
		return discord.Channel{}, false
	}
	return c, c.Type.Messageable()
}

func (s Scope) PartialMessage(argument string, current snowflake.ID) (PartialMessage, error) {
	ref, ok := s.parseMessageRef(argument, current)
	if !ok {
		return PartialMessage{}, ErrNotFound
	}
	c, ok := s.messageChannel(ref)
	if !ok {
		return PartialMessage{}, ErrChannelNotFound
	}
	return PartialMessage{ID: ref.message, Channel: c}, nil
}

func (s Scope) Message(argument string, current snowflake.ID) (discord.Message, error) {
	ref, ok := s.parseMessageRef(argument, current)
	if !ok {
		return discord.Message{}, ErrNotFound
	}
	if cached, ok := s.Client.State().Message(ref.message); ok {
		return cached.AsMessage(), nil
	}
	if _, ok := s.messageChannel(ref); !ok {
		return discord.Message{}, ErrChannelNotFound
	}
	raw, err := s.Client.API().GetMessage(s.Context, ref.channel, ref.message)
	var httpErr *rest.HTTPError
	switch {
	case errors.As(err, &httpErr) && httpErr.Status == http.StatusForbidden:
		return discord.Message{}, ErrChannelNotReadable
	case err != nil:
		return discord.Message{}, ErrNotFound
	}
	var msg discord.Message
	if err := json.Unmarshal(raw, &msg); err != nil {
		return discord.Message{}, err
	}
	return msg, nil
}

package discord

import (
	"time"

	"github.com/prettylittlelies/cry/snowflake"
)

type RelationshipType int

const (
	RelationshipNone RelationshipType = iota
	RelationshipFriend
	RelationshipBlocked
	RelationshipIncoming
	RelationshipOutgoing
	RelationshipImplicit
)

type Relationship struct {
	ID            snowflake.ID     `json:"id"`
	Type          RelationshipType `json:"type"`
	UserID        snowflake.ID     `json:"user_id,omitempty"`
	User          *User            `json:"user,omitempty"`
	Nickname      *string          `json:"nickname"`
	Since         *time.Time       `json:"since,omitempty"`
	IsSpamRequest bool             `json:"is_spam_request,omitempty"`
	UserIgnored   bool             `json:"user_ignored,omitempty"`
}

func (r *Relationship) Target() snowflake.ID {
	if r.User != nil {
		return r.User.ID
	}
	if r.UserID != 0 {
		return r.UserID
	}
	return r.ID
}

type ReadState struct {
	ID               snowflake.ID  `json:"id"`
	LastMessageID    *snowflake.ID `json:"last_message_id"`
	MentionCount     int           `json:"mention_count"`
	LastPinTimestamp *time.Time    `json:"last_pin_timestamp,omitempty"`
	Flags            int           `json:"flags,omitempty"`
	ReadStateType    int           `json:"read_state_type,omitempty"`
	LastViewed       *int          `json:"last_viewed,omitempty"`
	BadgeCount       int           `json:"badge_count,omitempty"`
}

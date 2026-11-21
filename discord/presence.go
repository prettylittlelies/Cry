package discord

import "github.com/prettylittlelies/cry/snowflake"

type Status string

const (
	StatusOnline    Status = "online"
	StatusIdle      Status = "idle"
	StatusDND       Status = "dnd"
	StatusInvisible Status = "invisible"
	StatusOffline   Status = "offline"
)

type ActivityType int

const (
	ActivityPlaying ActivityType = iota
	ActivityStreaming
	ActivityListening
	ActivityWatching
	ActivityCustom
	ActivityCompeting
	ActivityHang
)

type ActivityTimestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type ActivityParty struct {
	ID   string `json:"id,omitempty"`
	Size []int  `json:"size,omitempty"`
}

type ActivityEmoji struct {
	Name     string        `json:"name"`
	ID       *snowflake.ID `json:"id,omitempty"`
	Animated bool          `json:"animated,omitempty"`
}

type Activity struct {
	Name          string              `json:"name"`
	Type          ActivityType        `json:"type"`
	URL           *string             `json:"url,omitempty"`
	CreatedAt     int64               `json:"created_at,omitempty"`
	Timestamps    *ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID *snowflake.ID       `json:"application_id,omitempty"`
	Details       *string             `json:"details,omitempty"`
	State         *string             `json:"state,omitempty"`
	Emoji         *ActivityEmoji      `json:"emoji,omitempty"`
	SessionID     *string             `json:"session_id,omitempty"`
	Party         *ActivityParty      `json:"party,omitempty"`
	Assets        *ActivityAssets     `json:"assets,omitempty"`
	SyncID        *string             `json:"sync_id,omitempty"`
	Flags         int                 `json:"flags,omitempty"`
	Buttons       []string            `json:"buttons,omitempty"`
	Platform      *string             `json:"platform,omitempty"`
}

type ClientStatus struct {
	Desktop  Status `json:"desktop,omitempty"`
	Mobile   Status `json:"mobile,omitempty"`
	Web      Status `json:"web,omitempty"`
	Embedded Status `json:"embedded,omitempty"`
	VR       Status `json:"vr,omitempty"`
}

type Presence struct {
	User         *User         `json:"user,omitempty"`
	UserID       snowflake.ID  `json:"user_id,omitempty"`
	GuildID      *snowflake.ID `json:"guild_id,omitempty"`
	Status       Status        `json:"status"`
	Activities   []Activity    `json:"activities"`
	ClientStatus ClientStatus  `json:"client_status"`
}

func (p *Presence) ID() snowflake.ID {
	if p.User != nil {
		return p.User.ID
	}
	return p.UserID
}

type ClientInfo struct {
	Client  string `json:"client"`
	OS      string `json:"os"`
	Version int    `json:"version"`
}

type Session struct {
	SessionID  string     `json:"session_id"`
	Status     Status     `json:"status"`
	Active     bool       `json:"active,omitempty"`
	Activities []Activity `json:"activities"`
	ClientInfo ClientInfo `json:"client_info"`
}

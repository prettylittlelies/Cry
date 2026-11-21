package discord

import (
	"encoding/json"
	"strconv"

	"github.com/prettylittlelies/cry/snowflake"
)

type Permissions uint64

const (
	PermCreateInstantInvite Permissions = 1 << iota
	PermKickMembers
	PermBanMembers
	PermAdministrator
	PermManageChannels
	PermManageGuild
	PermAddReactions
	PermViewAuditLog
	PermPrioritySpeaker
	PermStream
	PermViewChannel
	PermSendMessages
	PermSendTTSMessages
	PermManageMessages
	PermEmbedLinks
	PermAttachFiles
	PermReadMessageHistory
	PermMentionEveryone
	PermUseExternalEmojis
	PermViewGuildInsights
	PermConnect
	PermSpeak
	PermMuteMembers
	PermDeafenMembers
	PermMoveMembers
	PermUseVAD
	PermChangeNickname
	PermManageNicknames
	PermManageRoles
	PermManageWebhooks
	PermManageGuildExpressions
	PermUseApplicationCommands
	PermRequestToSpeak
	PermManageEvents
	PermManageThreads
	PermCreatePublicThreads
	PermCreatePrivateThreads
	PermUseExternalStickers
	PermSendMessagesInThreads
	PermUseEmbeddedActivities
	PermModerateMembers
)

const PermAll Permissions = 1<<64 - 1

func (p Permissions) Has(want Permissions) bool {
	return p&PermAdministrator != 0 || p&want == want
}

func (p Permissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(p), 10))
}

func (p *Permissions) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		var n uint64
		if err := json.Unmarshal(b, &n); err != nil {
			return err
		}
		*p = Permissions(n)
		return nil
	}
	n, err := strconv.ParseUint(s, 10, 64)
	*p = Permissions(n)
	return err
}

type OverwriteType int

const (
	OverwriteRole OverwriteType = iota
	OverwriteMember
)

type Overwrite struct {
	ID    snowflake.ID  `json:"id"`
	Type  OverwriteType `json:"type"`
	Allow Permissions   `json:"allow"`
	Deny  Permissions   `json:"deny"`
}

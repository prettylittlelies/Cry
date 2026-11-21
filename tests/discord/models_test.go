package discord_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/prettylittlelies/cry/discord"
)

const readyFixture = `{
	"v": 9,
	"session_id": "abc",
	"user": {"id": 53908232506183680, "username": "cry", "discriminator": "0", "global_name": "Cry", "avatar": null, "flags": 1, "mfa_enabled": true, "verified": true, "email": null, "premium_type": 2, "bio": "", "banner": null, "accent_color": null},
	"users": [{"id": "81384788765712384", "username": "friend", "discriminator": "0", "avatar": null, "public_flags": 64}],
	"guilds": [{
		"id": 81384788765712384,
		"joined_at": "2021-03-01T12:00:00.000000+00:00",
		"large": false,
		"member_count": 42,
		"properties": {"id": "81384788765712384", "name": "cry", "icon": null, "owner_id": "53908232506183680", "premium_tier": 2, "features": ["COMMUNITY"]},
		"channels": [{"id": "1", "type": 0, "name": "general", "position": 0}],
		"roles": [{"id": 81384788765712384, "name": "@everyone", "permissions": "1071698660929", "position": 0, "color": 0, "hoist": false, "managed": false, "mentionable": false, "flags": 0}]
	}],
	"merged_members": [[{"user_id": 53908232506183680, "roles": [], "joined_at": "2021-03-01T12:00:00+00:00", "nick": "me", "avatar": null, "flags": 0, "deaf": false, "mute": false}]],
	"private_channels": [{"id": 7, "type": 1, "recipient_ids": [81384788765712384], "last_message_id": 9}],
	"relationships": [{"id": "81384788765712384", "type": 1, "user_id": 81384788765712384, "nickname": null, "since": "2022-01-01T00:00:00+00:00"}]
}`

func TestReadyDecodes(t *testing.T) {
	var r discord.Ready
	if err := json.Unmarshal([]byte(readyFixture), &r); err != nil {
		t.Fatal(err)
	}
	if r.User.ID != 53908232506183680 || r.User.DisplayName() != "Cry" || r.User.PremiumType != discord.PremiumTypeTier2 {
		t.Fatalf("user %+v", r.User)
	}
	if r.User.Flags&discord.UserFlagsStaff == 0 {
		t.Fatalf("user flags %v", r.User.Flags)
	}
	if len(r.Users) != 1 || r.Users[0].Mention() != "<@81384788765712384>" {
		t.Fatalf("users %+v", r.Users)
	}
	g := r.Guilds[0]
	if g.ID != 81384788765712384 || g.Properties.Name != "cry" || g.MemberCount != 42 || g.JoinedAt.Year() != 2021 {
		t.Fatalf("guild %+v", g)
	}
	if g.Channels[0].Type != discord.ChannelTypeGuildText || g.Channels[0].Mention() != "<#1>" {
		t.Fatalf("channel %+v", g.Channels[0])
	}
	everyone := g.Roles[0]
	if !everyone.Permissions.Has(discord.PermissionSendMessages) || everyone.Permissions.Has(discord.PermissionBanMembers) {
		t.Fatalf("permissions %v", everyone.Permissions)
	}
	if len(r.MergedMembers) != 1 || r.MergedMembers[0][0].Nick == nil || *r.MergedMembers[0][0].Nick != "me" {
		t.Fatalf("merged members %+v", r.MergedMembers)
	}
	if r.PrivateChannels[0].Type != discord.ChannelTypeDM || r.Relationships[0].Type != discord.RelationshipTypeFriend {
		t.Fatalf("dm=%+v rel=%+v", r.PrivateChannels[0], r.Relationships[0])
	}
}

const profileFixture = `{
	"user": {"id": "5", "username": "cry", "discriminator": "0", "avatar": null},
	"user_profile": {"bio": "hi", "pronouns": "they/them", "accent_color": 16711680, "theme_colors": [1, 2]},
	"badges": [{"id": "staff", "description": "Discord Staff", "icon": "abc"}],
	"guild_badges": [],
	"mutual_friends": [{"id": "6", "username": "pal", "discriminator": "0", "avatar": null}],
	"connected_accounts": [{"type": "spotify", "id": "x", "name": "cry", "verified": true}],
	"premium_type": 2,
	"premium_since": "2020-01-01T00:00:00+00:00",
	"premium_guild_since": null
}`

func TestProfileDecodes(t *testing.T) {
	var p discord.GetUserProfileResponse
	if err := json.Unmarshal([]byte(profileFixture), &p); err != nil {
		t.Fatal(err)
	}
	if p.UserProfile.Pronouns != "they/them" || p.UserProfile.Bio != "hi" || p.UserProfile.AccentColor == nil || *p.UserProfile.AccentColor != 16711680 {
		t.Fatalf("profile %+v", p.UserProfile)
	}
	if len(p.Badges) != 1 || p.Badges[0].ID != "staff" || len(p.MutualFriends) != 1 || len(p.ConnectedAccounts) != 1 {
		t.Fatalf("badges=%+v friends=%+v accounts=%+v", p.Badges, p.MutualFriends, p.ConnectedAccounts)
	}
	if p.PremiumType == nil || *p.PremiumType != discord.PremiumTypeTier2 || p.PremiumSince == nil || p.PremiumGuildSince != nil {
		t.Fatalf("premium %+v %v %v", p.PremiumType, p.PremiumSince, p.PremiumGuildSince)
	}
}

func TestPermissionsJSON(t *testing.T) {
	var p discord.Permissions
	if err := json.Unmarshal([]byte(`"8"`), &p); err != nil || !p.Has(discord.PermissionBanMembers, discord.PermissionManageGuild) {
		t.Fatalf("administrator implies everything: %v %v", p, err)
	}
	q := discord.Permissions(0).With(discord.PermissionSendMessages, discord.PermissionViewChannel).Without(discord.PermissionViewChannel)
	out, _ := json.Marshal(q)
	if string(out) != `"2048"` {
		t.Fatalf("marshal %s", out)
	}
}

func TestImageData(t *testing.T) {
	img := discord.NewImageData([]byte("\x89PNG\r\n\x1a\n0000"))
	if !strings.HasPrefix(string(img), "data:image/png;base64,") {
		t.Fatalf("image data %q", img)
	}
}

func TestJumpURL(t *testing.T) {
	m := discord.Message{ID: 3, ChannelID: 2}
	if m.JumpURL(1) != "https://discord.com/channels/1/2/3" || m.JumpURL(0) != "https://discord.com/channels/@me/2/3" {
		t.Fatal(m.JumpURL(1), m.JumpURL(0))
	}
}

package state

import (
	"slices"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/snowflake"
)

const everything = discord.Permissions(1<<64 - 1)

func (s *State) GuildPermissions(guild, user snowflake.ID) discord.Permissions {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.guildPermissions(guild, user)
}

func (s *State) guildPermissions(guild, user snowflake.ID) discord.Permissions {
	g, ok := s.guilds[guild]
	if !ok {
		return 0
	}
	if g.Properties.OwnerID == user {
		return everything
	}
	member, ok := s.members[guild][user]
	var perms discord.Permissions
	for _, r := range g.Roles {
		if r.ID == guild || ok && slices.Contains(member.Roles, r.ID) {
			perms |= r.Permissions
		}
	}
	if perms.Has(discord.PermissionAdministrator) {
		return everything
	}
	return perms
}

func (s *State) ChannelPermissions(channel, user snowflake.ID) discord.Permissions {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ch, ok := s.channels[channel]
	if !ok || ch.GuildID == 0 {
		return everything
	}
	base := s.guildPermissions(ch.GuildID, user)
	if base == everything {
		return everything
	}
	member := s.members[ch.GuildID][user]
	var roleAllow, roleDeny discord.Permissions
	var memberOverwrite *discord.PermissionOverwrite
	for i := range ch.PermissionOverwrites {
		o := &ch.PermissionOverwrites[i]
		switch {
		case o.ID == ch.GuildID:
			base = (base &^ o.Deny) | o.Allow
		case o.ID == user:
			memberOverwrite = o
		case member != nil && slices.Contains(member.Roles, o.ID):
			roleAllow |= o.Allow
			roleDeny |= o.Deny
		}
	}
	base = (base &^ roleDeny) | roleAllow
	if memberOverwrite != nil {
		base = (base &^ memberOverwrite.Deny) | memberOverwrite.Allow
	}
	return base
}

package state

import (
	"slices"

	"github.com/prettylittlelies/cry/discord"
)

func sameEmoji(a, b discord.Emoji) bool {
	if a.ID != nil || b.ID != nil {
		return a.ID != nil && b.ID != nil && *a.ID == *b.ID
	}
	return a.Name == b.Name
}

func addReaction(reactions []discord.Reaction, emoji discord.Emoji, burst, mine bool) []discord.Reaction {
	i := slices.IndexFunc(reactions, func(r discord.Reaction) bool { return sameEmoji(r.Emoji, emoji) })
	if i < 0 {
		reactions = append(reactions, discord.Reaction{Emoji: emoji})
		i = len(reactions) - 1
	}
	r := &reactions[i]
	r.Count++
	if burst {
		r.CountDetails.Burst++
		r.MeBurst = r.MeBurst || mine
	} else {
		r.CountDetails.Normal++
		r.Me = r.Me || mine
	}
	return reactions
}

func removeReaction(reactions []discord.Reaction, emoji discord.Emoji, burst, mine bool) []discord.Reaction {
	i := slices.IndexFunc(reactions, func(r discord.Reaction) bool { return sameEmoji(r.Emoji, emoji) })
	if i < 0 {
		return reactions
	}
	r := &reactions[i]
	r.Count = max(0, r.Count-1)
	if burst {
		r.CountDetails.Burst = max(0, r.CountDetails.Burst-1)
		r.MeBurst = r.MeBurst && !mine
	} else {
		r.CountDetails.Normal = max(0, r.CountDetails.Normal-1)
		r.Me = r.Me && !mine
	}
	if r.Count == 0 {
		return slices.Delete(reactions, i, i+1)
	}
	return reactions
}

package discord_test

import (
	"testing"

	"github.com/prettylittlelies/cry/discord"
)

func TestMarkdownMatchesDiscordPy(t *testing.T) {
	for _, tc := range []struct{ in, escaped, asNeeded, removed string }{
		{"**bold** and *it*", "\\*\\*bold\\*\\* and \\*it\\*", "\\*\\*bold\\*\\* and \\*it*", "bold and it"},
		{"a_b_c", "a\\_b\\_c", "a\\_b_c", "abc"},
		{"`code` here", "\\`code\\` here", "\\`code` here", "code here"},
		{"> quote\n# head\n - item", "\\> quote\n\\# head\n\\ - item", "\\> quote\n\\# head\n\\ - item", "quote\n head\n item"},
		{"[x](http://y) http://a.b/c_d", "\\[x](http://y) http://a.b/c_d", "\\[x](http://y) http://a.b/c_d", " http://a.b/c_d"},
		{"lone * star", "lone \\* star", "lone * star", "lone  star"},
		{"\\ back *x*", "\\\\ back \\*x\\*", "\\\\ back \\*x*", " back x"},
		{"||spoil||", "\\|\\|spoil\\|\\|", "\\|\\|spoil||", "spoil"},
		{"~~s~~ __u__", "\\~\\~s\\~\\~ \\_\\_u\\_\\_", "\\~\\~s~~ \\_\\_u__", "s u"},
		{">>> block\n### h3", "\\>>> block\n\\### h3", "\\>>> block\n\\### h3", "block\n h3"},
		{"<https://x.y/a_b> _z_", "<https://x.y/a_b> \\_z\\_", "<https://x.y/a\\_b> \\_z_", "<https://x.y/a_b> z"},
	} {
		if got := discord.EscapeMarkdown(tc.in, false, true); got != tc.escaped {
			t.Errorf("escape %q: %q want %q", tc.in, got, tc.escaped)
		}
		if got := discord.EscapeMarkdown(tc.in, true, true); got != tc.asNeeded {
			t.Errorf("as needed %q: %q want %q", tc.in, got, tc.asNeeded)
		}
		if got := discord.RemoveMarkdown(tc.in, true); got != tc.removed {
			t.Errorf("remove %q: %q want %q", tc.in, got, tc.removed)
		}
	}
	in := "@everyone @here <@!123456789012345678> <@&123456789012345678>"
	want := "@\u200beveryone @\u200bhere <@\u200b!123456789012345678> <@\u200b&123456789012345678>"
	if got := discord.EscapeMentions(in); got != want {
		t.Errorf("mentions %q", got)
	}
}

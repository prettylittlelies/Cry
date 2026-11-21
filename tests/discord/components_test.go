package discord_test

import (
	"encoding/json"
	"testing"

	"github.com/prettylittlelies/cry/discord"
)

const componentsFixture = `[
	{"type": 1, "id": 1, "components": [
		{"type": 2, "id": 2, "style": 1, "label": "Accept", "custom_id": "accept"},
		{"type": 3, "id": 3, "custom_id": "pick", "options": [{"label": "A", "value": "a"}]}
	]},
	{"type": 17, "id": 4, "components": [
		{"type": 10, "id": 5, "content": "hello"},
		{"type": 9, "id": 6, "components": [{"type": 10, "id": 7, "content": "inside"}],
			"accessory": {"type": 11, "id": 8, "media": {"url": "https://cdn.discordapp.com/x.png"}}}
	]},
	{"type": 99, "id": 9, "future": true}
]`

func TestComponentsDecodeByType(t *testing.T) {
	var cs discord.Components
	if err := json.Unmarshal([]byte(componentsFixture), &cs); err != nil {
		t.Fatal(err)
	}
	if len(cs) != 3 {
		t.Fatalf("got %d components", len(cs))
	}
	row, ok := cs[0].(*discord.ActionRow)
	if !ok || row.ComponentType() != discord.ComponentTypeActionRow {
		t.Fatalf("first is %T", cs[0])
	}
	if b, ok := row.Components[0].(*discord.Button); !ok || b.CustomID != "accept" {
		t.Fatalf("button %#v", row.Components[0])
	}
	if s, ok := row.Components[1].(*discord.StringSelect); !ok || s.Options[0].Value != "a" {
		t.Fatalf("select %#v", row.Components[1])
	}
	container, ok := cs[1].(*discord.Container)
	if !ok {
		t.Fatalf("second is %T", cs[1])
	}
	if text, ok := container.Components[0].(*discord.TextDisplay); !ok || text.Content != "hello" {
		t.Fatalf("text %#v", container.Components[0])
	}
	section, ok := container.Components[1].(*discord.Section)
	if !ok {
		t.Fatalf("section is %T", container.Components[1])
	}
	if _, ok := section.Accessory.Component.(*discord.Thumbnail); !ok {
		t.Fatalf("accessory is %T", section.Accessory.Component)
	}
	unknown, ok := cs[2].(*discord.UnknownComponent)
	if !ok || unknown.ComponentType() != 99 {
		t.Fatalf("third is %T", cs[2])
	}
}

func TestComponentsRoundTrip(t *testing.T) {
	var cs discord.Components
	if err := json.Unmarshal([]byte(componentsFixture), &cs); err != nil {
		t.Fatal(err)
	}
	out, err := json.Marshal(cs)
	if err != nil {
		t.Fatal(err)
	}
	var again discord.Components
	if err := json.Unmarshal(out, &again); err != nil {
		t.Fatal(err)
	}
	if len(again) != 3 {
		t.Fatalf("round trip lost components: %s", out)
	}
	if string(mustMarshal(t, again[2])) != `{"type":99,"id":9,"future":true}` {
		t.Fatalf("unknown component not preserved: %s", mustMarshal(t, again[2]))
	}
	if b, ok := again[0].(*discord.ActionRow).Components[0].(*discord.Button); !ok || b.Label != "Accept" {
		t.Fatalf("button after round trip %#v", again[0])
	}
}

func mustMarshal(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

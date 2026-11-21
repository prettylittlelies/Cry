package internal_test

import (
	"testing"

	"github.com/prettylittlelies/cry/discord"
	"github.com/prettylittlelies/cry/internal/lenient"
)

func TestLenientCoercesETFShapes(t *testing.T) {
	raw := `{"user":{"id":"222222222222222222","username":"x","discriminator":"0","avatar":null},"status":"online","client_status":{},
		"activities":[{"id":"a","name":"Spotify","type":2,"created_at":"1700000000000","timestamps":{"start":1700000000000,"end":1700000300000}}]}`
	var e discord.PresenceUpdateEvent
	if err := lenient.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	a := e.Activities[0]
	if a.CreatedAt != 1700000000000 || e.User.ID != 222222222222222222 {
		t.Fatalf("decoded %+v", a)
	}
	var c discord.ChannelUpdateEvent
	if err := lenient.Unmarshal([]byte(`{"id":"333333333333333333","type":0,"version":1700000000000}`), &c); err != nil {
		t.Fatal(err)
	}
}

func TestLenientDropsShapeMismatches(t *testing.T) {
	type inner struct{ A int }
	type outer struct {
		Keep  string           `json:"keep"`
		Inner inner            `json:"inner"`
		List  []inner          `json:"list"`
		Byid  map[string]inner `json:"byid"`
	}
	var o outer
	err := lenient.Unmarshal([]byte(`{"keep":"yes","inner":[1,2],"list":[{"A":1},[3]],"byid":{"x":[4],"y":{"A":5}}}`), &o)
	if err != nil {
		t.Fatal(err)
	}
	if o.Keep != "yes" || o.Inner.A != 0 || len(o.List) != 2 || o.List[0].A != 1 || o.Byid["y"].A != 5 {
		t.Fatalf("%+v", o)
	}
}

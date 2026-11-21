package ext_test

import (
	"errors"
	"maps"
	"reflect"
	"slices"
	"testing"

	"github.com/prettylittlelies/cry/ext/commands"
)

type banFlags struct {
	commands.FlagConverter
	Member string
	Reason string   `default:"no reason"`
	Days   int      `default:"1" aliases:"d"`
	Tags   []string `flag:"tag" default:""`
	Limit  *int
}

type banCommand struct {
	Flags banFlags `commands:"rest"`
}

func TestFlagConverter(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	var got banFlags
	h.bot.AddCommand(commands.New("ban", func(_ *commands.Context, p banCommand) error { got = p.Flags; return nil }))

	h.run(other, 10, "!ban member: alice reason: spamming a lot d: 7 tag: a tag: b")
	if got.Member != "alice" || got.Reason != "spamming a lot" || got.Days != 7 || !slices.Equal(got.Tags, []string{"a", "b"}) || got.Limit != nil {
		t.Fatalf("got %+v", got)
	}
	h.run(other, 10, "!ban member: bob limit: 3")
	if got.Member != "bob" || got.Reason != "no reason" || got.Days != 1 || got.Tags != nil || got.Limit == nil || *got.Limit != 3 {
		t.Fatalf("defaults %+v", got)
	}

	for input, want := range map[string]any{
		"!ban reason: x":           &commands.MissingRequiredFlag{},
		"!ban member: reason: x":   &commands.MissingFlagArgument{},
		"!ban member: a days: x":   &commands.BadFlagArgument{},
		"!ban member: a member: b": &commands.TooManyFlags{},
		"!ban member: a days:":     &commands.MissingFlagArgument{},
	} {
		h.run(other, 10, input)
		err := h.err(t)
		target := reflect.New(reflect.TypeOf(want))
		if !errors.As(err, target.Interface()) || !errors.Is(err, commands.ErrFlag) {
			t.Errorf("%q: %v", input, err)
		}
	}
	h.run(other, 10, "!ban member: a member: b")
	if err := h.err(t); err.Error() != "Too many flag values, expected 1 but received 2." {
		t.Errorf("message %q", err)
	}
}

type searchFlags struct {
	commands.FlagConverter `prefix:"--" delimiter:" " case_insensitive:"true"`
	Query                  string `positional:"true"`
	Limit                  int    `default:"10"`
	Point                  [2]int
	Pairs                  map[string]int `flag:"set"`
	Last                   string         `default:"" max_args:"1" override:"true"`
}

type searchCommand struct {
	Flags searchFlags `commands:"rest"`
}

func TestFlagConverterOptions(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	var got searchFlags
	h.bot.AddCommand(commands.New("search", func(_ *commands.Context, p searchCommand) error { got = p.Flags; return nil }))
	h.run(other, 10, `!search cute cats --LIMIT 5 --point 3 4 --set a 1 --set b 2 --last x --last y`)
	if got.Query != "cute cats" || got.Limit != 5 || got.Point != [2]int{3, 4} || got.Last != "y" ||
		!maps.Equal(got.Pairs, map[string]int{"a": 1, "b": 2}) {
		t.Fatalf("got %+v", got)
	}
	h.run(other, 10, `!search q --point 1`)
	var missing *commands.MissingFlagArgument
	if err := h.err(t); !errors.As(err, &missing) || missing.Flag.Name != "point" {
		t.Fatalf("tuple: %v", err)
	}
}

type optionalFlags struct {
	commands.FlagConverter
	Silent bool `default:"false"`
}

func TestFlagConverterDefaultsAndValidation(t *testing.T) {
	h := newBot(t, commands.WithMode(commands.UserBot))
	called := false
	h.bot.AddCommand(commands.New("quiet", func(_ *commands.Context, p struct {
		Flags optionalFlags `commands:"rest"`
	}) error {
		called = !p.Flags.Silent
		return nil
	}))
	h.run(other, 10, "!quiet")
	if !called {
		t.Fatal("constructible flags must default when the argument is missing")
	}

	type conflicting struct {
		commands.FlagConverter
		A string `aliases:"b"`
		B string
	}
	if _, err := commands.FlagsOf(reflect.TypeFor[conflicting]()); err == nil {
		t.Fatal("alias conflict must fail")
	}
	type spaced struct {
		commands.FlagConverter
		A string `flag:"a b"`
	}
	if _, err := commands.FlagsOf(reflect.TypeFor[spaced]()); err == nil {
		t.Fatal("spaces must fail")
	}
	defer func() {
		if recover() == nil {
			t.Fatal("invalid flag definitions must panic at command construction")
		}
	}()
	commands.New("bad", func(*commands.Context, struct {
		Flags spaced `commands:"rest"`
	}) error {
		return nil
	})
}

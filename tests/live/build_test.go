//go:build live

package live_test

import (
	"context"
	"testing"
	"time"

	"github.com/prettylittlelies/cry/build"
)

func TestBuildSnapshot(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	snap, err := build.NewScraper().Fetch(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range build.Channels() {
		r, _ := snap.Release(c)
		t.Logf("%s: build=%s win=%s native=%s mac=%s electron=%s chrome=%s", c, r.Number, r.Windows.Host, r.Windows.Native, r.MacOS.Host, r.MacOS.Electron, r.MacOS.Chrome)
	}
	t.Logf("android=%s (min %d) ios=%s (min os %s)", snap.Mobile.Android, snap.Mobile.AndroidMinimum, snap.Mobile.IOS, snap.Mobile.IOSMinimumOS)
}

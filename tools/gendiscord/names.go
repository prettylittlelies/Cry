package main

import (
	"github.com/prettylittlelies/cry/tools/internal/naming"
	"strings"
)

type namer struct {
	taken map[string]string
}

var docPriority = []string{"gateway-events", "user", "guild", "channel", "message", "presence", "relationships", "emoji", "sticker", "application", "user-settings"}

func priority(doc string) int {
	for i, d := range docPriority {
		if d == doc {
			return i
		}
	}
	return len(docPriority)
}

func (n *namer) assign(doc, title string, count map[string]int) string {
	base := naming.GoName(title)
	if _, used := n.taken[base]; count[base] > 1 && used {
		base = naming.GoName(doc) + base
	}
	for i := 2; ; i++ {
		if _, used := n.taken[base]; !used {
			n.taken[base] = title
			return base
		}
		base = naming.GoName(title) + naming.GoName(doc)
		if i > 2 {
			base += naming.GoName(strings.Repeat("x", i))
		}
	}
}

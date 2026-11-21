package commands

import (
	"context"
	"sync"
)

type Extension struct {
	Setup    func(ctx context.Context, b *Bot) error
	Teardown func(ctx context.Context, b *Bot)
}

var (
	extensionsMu sync.RWMutex
	registry     = map[string]Extension{}
)

func RegisterExtension(name string, ext Extension) {
	extensionsMu.Lock()
	registry[name] = ext
	extensionsMu.Unlock()
}

func lookupExtension(name string) (Extension, bool) {
	extensionsMu.RLock()
	defer extensionsMu.RUnlock()
	e, ok := registry[name]
	return e, ok
}

func (b *Bot) Extensions() []string {
	b.mu.RLock()
	defer b.mu.RUnlock()
	out := make([]string, 0, len(b.extensions))
	for name := range b.extensions {
		out = append(out, name)
	}
	return out
}

func (b *Bot) LoadExtension(ctx context.Context, name string) error {
	b.mu.Lock()
	if _, loaded := b.extensions[name]; loaded {
		b.mu.Unlock()
		return &ExtensionAlreadyLoaded{errExtension("Extension %q is already loaded.", name)}
	}
	b.mu.Unlock()
	ext, ok := lookupExtension(name)
	if !ok {
		return &ExtensionNotFound{errExtension("Extension %q could not be loaded.", name)}
	}
	if ext.Setup == nil {
		return &NoEntryPointError{errExtension("Extension %q has no 'setup' function.", name)}
	}
	rec := &extensionRecord{ext: ext}
	b.mu.Lock()
	b.loading = rec
	b.mu.Unlock()
	err := ext.Setup(ctx, b)
	b.mu.Lock()
	b.loading = nil
	b.mu.Unlock()
	if err != nil {
		b.removeReferences(ctx, rec)
		return &ExtensionFailed{ExtensionError: errExtension("Extension %q raised an error: "+err.Error(), name), Original: err}
	}
	b.mu.Lock()
	b.extensions[name] = rec
	b.mu.Unlock()
	return nil
}

func (b *Bot) removeReferences(ctx context.Context, rec *extensionRecord) {
	for _, cog := range rec.cogs {
		b.RemoveCog(ctx, cog)
	}
	for _, cmd := range rec.commands {
		b.RemoveCommand(cmd)
	}
}

func (b *Bot) UnloadExtension(ctx context.Context, name string) error {
	b.mu.Lock()
	rec, ok := b.extensions[name]
	delete(b.extensions, name)
	b.mu.Unlock()
	if !ok {
		return &ExtensionNotLoaded{errExtension("Extension %q has not been loaded.", name)}
	}
	b.removeReferences(ctx, rec)
	if rec.ext.Teardown != nil {
		rec.ext.Teardown(ctx, b)
	}
	return nil
}

func (b *Bot) ReloadExtension(ctx context.Context, name string) error {
	if err := b.UnloadExtension(ctx, name); err != nil {
		return err
	}
	return b.LoadExtension(ctx, name)
}

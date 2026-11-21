package commands

import (
	"reflect"

	"github.com/prettylittlelies/cry/discord"
)

var attachmentType = reflect.TypeFor[discord.Attachment]()

func (p *Parameter) isAttachment() bool {
	return p.converter() == attachmentType
}

func (c *Context) nextAttachment() (discord.Attachment, bool) {
	if c.attachments >= len(c.Message.Attachments) {
		return discord.Attachment{}, false
	}
	a := c.Message.Attachments[c.attachments]
	c.attachments++
	return a, true
}

func (c *Context) restAttachments() []discord.Attachment {
	rest := c.Message.Attachments[min(c.attachments, len(c.Message.Attachments)):]
	c.attachments = len(c.Message.Attachments)
	return append([]discord.Attachment(nil), rest...)
}

func transformAttachment(ctx *Context, p *Parameter) (reflect.Value, error) {
	switch {
	case p.Greedy:
		return reflect.ValueOf(ctx.restAttachments()), nil
	case p.Optional():
		a, ok := ctx.nextAttachment()
		if !ok {
			return reflect.Zero(p.Type), nil
		}
		return reflect.ValueOf(&a), nil
	}
	a, ok := ctx.nextAttachment()
	if !ok {
		return reflect.Value{}, errMissingAttachment(p)
	}
	return reflect.ValueOf(a), nil
}

func attachmentSignature(p *Parameter, name string) string {
	switch {
	case p.Optional():
		return "[" + name + " (upload a file)]"
	case p.Greedy:
		return "[" + name + " (upload files)]..."
	}
	return "<" + name + " (upload a file)>"
}

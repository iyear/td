// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ClickChatSponsoredMessageRequest represents TL type `clickChatSponsoredMessage#39ef7a17`.
type ClickChatSponsoredMessageRequest struct {
	// Chat identifier of the sponsored message
	ChatID int64
	// Identifier of the sponsored message
	MessageID int64
	// Pass true if the media was clicked in the sponsored message
	IsMediaClick bool
	// Pass true if the user expanded the video from the sponsored message fullscreen before
	// the click
	FromFullscreen bool
}

// ClickChatSponsoredMessageRequestTypeID is TL type id of ClickChatSponsoredMessageRequest.
const ClickChatSponsoredMessageRequestTypeID = 0x39ef7a17

// Ensuring interfaces in compile-time for ClickChatSponsoredMessageRequest.
var (
	_ bin.Encoder     = &ClickChatSponsoredMessageRequest{}
	_ bin.Decoder     = &ClickChatSponsoredMessageRequest{}
	_ bin.BareEncoder = &ClickChatSponsoredMessageRequest{}
	_ bin.BareDecoder = &ClickChatSponsoredMessageRequest{}
)

func (c *ClickChatSponsoredMessageRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.MessageID == 0) {
		return false
	}
	if !(c.IsMediaClick == false) {
		return false
	}
	if !(c.FromFullscreen == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClickChatSponsoredMessageRequest) String() string {
	if c == nil {
		return "ClickChatSponsoredMessageRequest(nil)"
	}
	type Alias ClickChatSponsoredMessageRequest
	return fmt.Sprintf("ClickChatSponsoredMessageRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ClickChatSponsoredMessageRequest) TypeID() uint32 {
	return ClickChatSponsoredMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ClickChatSponsoredMessageRequest) TypeName() string {
	return "clickChatSponsoredMessage"
}

// TypeInfo returns info about TL type.
func (c *ClickChatSponsoredMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "clickChatSponsoredMessage",
		ID:   ClickChatSponsoredMessageRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "IsMediaClick",
			SchemaName: "is_media_click",
		},
		{
			Name:       "FromFullscreen",
			SchemaName: "from_fullscreen",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ClickChatSponsoredMessageRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clickChatSponsoredMessage#39ef7a17 as nil")
	}
	b.PutID(ClickChatSponsoredMessageRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ClickChatSponsoredMessageRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clickChatSponsoredMessage#39ef7a17 as nil")
	}
	b.PutInt53(c.ChatID)
	b.PutInt53(c.MessageID)
	b.PutBool(c.IsMediaClick)
	b.PutBool(c.FromFullscreen)
	return nil
}

// Decode implements bin.Decoder.
func (c *ClickChatSponsoredMessageRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clickChatSponsoredMessage#39ef7a17 to nil")
	}
	if err := b.ConsumeID(ClickChatSponsoredMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ClickChatSponsoredMessageRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clickChatSponsoredMessage#39ef7a17 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field message_id: %w", err)
		}
		c.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field is_media_click: %w", err)
		}
		c.IsMediaClick = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field from_fullscreen: %w", err)
		}
		c.FromFullscreen = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ClickChatSponsoredMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode clickChatSponsoredMessage#39ef7a17 as nil")
	}
	b.ObjStart()
	b.PutID("clickChatSponsoredMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(c.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(c.MessageID)
	b.Comma()
	b.FieldStart("is_media_click")
	b.PutBool(c.IsMediaClick)
	b.Comma()
	b.FieldStart("from_fullscreen")
	b.PutBool(c.FromFullscreen)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ClickChatSponsoredMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode clickChatSponsoredMessage#39ef7a17 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("clickChatSponsoredMessage"); err != nil {
				return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field chat_id: %w", err)
			}
			c.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field message_id: %w", err)
			}
			c.MessageID = value
		case "is_media_click":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field is_media_click: %w", err)
			}
			c.IsMediaClick = value
		case "from_fullscreen":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode clickChatSponsoredMessage#39ef7a17: field from_fullscreen: %w", err)
			}
			c.FromFullscreen = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (c *ClickChatSponsoredMessageRequest) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}

// GetMessageID returns value of MessageID field.
func (c *ClickChatSponsoredMessageRequest) GetMessageID() (value int64) {
	if c == nil {
		return
	}
	return c.MessageID
}

// GetIsMediaClick returns value of IsMediaClick field.
func (c *ClickChatSponsoredMessageRequest) GetIsMediaClick() (value bool) {
	if c == nil {
		return
	}
	return c.IsMediaClick
}

// GetFromFullscreen returns value of FromFullscreen field.
func (c *ClickChatSponsoredMessageRequest) GetFromFullscreen() (value bool) {
	if c == nil {
		return
	}
	return c.FromFullscreen
}

// ClickChatSponsoredMessage invokes method clickChatSponsoredMessage#39ef7a17 returning error if any.
func (c *Client) ClickChatSponsoredMessage(ctx context.Context, request *ClickChatSponsoredMessageRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

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

// GetChatMessageCalendarRequest represents TL type `getChatMessageCalendar#ec8f2114`.
type GetChatMessageCalendarRequest struct {
	// Identifier of the chat in which to return information about messages
	ChatID int64
	// Filter for message content. Filters searchMessagesFilterEmpty,
	// searchMessagesFilterMention and searchMessagesFilterUnreadMention are unsupported in
	// this function
	Filter SearchMessagesFilterClass
	// The message identifier from which to return information about messages; use 0 to get
	// results from the last message
	FromMessageID int64
}

// GetChatMessageCalendarRequestTypeID is TL type id of GetChatMessageCalendarRequest.
const GetChatMessageCalendarRequestTypeID = 0xec8f2114

// Ensuring interfaces in compile-time for GetChatMessageCalendarRequest.
var (
	_ bin.Encoder     = &GetChatMessageCalendarRequest{}
	_ bin.Decoder     = &GetChatMessageCalendarRequest{}
	_ bin.BareEncoder = &GetChatMessageCalendarRequest{}
	_ bin.BareDecoder = &GetChatMessageCalendarRequest{}
)

func (g *GetChatMessageCalendarRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.FromMessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatMessageCalendarRequest) String() string {
	if g == nil {
		return "GetChatMessageCalendarRequest(nil)"
	}
	type Alias GetChatMessageCalendarRequest
	return fmt.Sprintf("GetChatMessageCalendarRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatMessageCalendarRequest) TypeID() uint32 {
	return GetChatMessageCalendarRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatMessageCalendarRequest) TypeName() string {
	return "getChatMessageCalendar"
}

// TypeInfo returns info about TL type.
func (g *GetChatMessageCalendarRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatMessageCalendar",
		ID:   GetChatMessageCalendarRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "FromMessageID",
			SchemaName: "from_message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatMessageCalendarRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageCalendar#ec8f2114 as nil")
	}
	b.PutID(GetChatMessageCalendarRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatMessageCalendarRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageCalendar#ec8f2114 as nil")
	}
	b.PutInt53(g.ChatID)
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessageCalendar#ec8f2114: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessageCalendar#ec8f2114: field filter: %w", err)
	}
	b.PutInt53(g.FromMessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatMessageCalendarRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageCalendar#ec8f2114 to nil")
	}
	if err := b.ConsumeID(GetChatMessageCalendarRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatMessageCalendarRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageCalendar#ec8f2114 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: field from_message_id: %w", err)
		}
		g.FromMessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatMessageCalendarRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessageCalendar#ec8f2114 as nil")
	}
	b.ObjStart()
	b.PutID("getChatMessageCalendar")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.FieldStart("filter")
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessageCalendar#ec8f2114: field filter is nil")
	}
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessageCalendar#ec8f2114: field filter: %w", err)
	}
	b.FieldStart("from_message_id")
	b.PutInt53(g.FromMessageID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatMessageCalendarRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessageCalendar#ec8f2114 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatMessageCalendar"); err != nil {
				return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: field chat_id: %w", err)
			}
			g.ChatID = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: field filter: %w", err)
			}
			g.Filter = value
		case "from_message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessageCalendar#ec8f2114: field from_message_id: %w", err)
			}
			g.FromMessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatMessageCalendarRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetFilter returns value of Filter field.
func (g *GetChatMessageCalendarRequest) GetFilter() (value SearchMessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetFromMessageID returns value of FromMessageID field.
func (g *GetChatMessageCalendarRequest) GetFromMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.FromMessageID
}

// GetChatMessageCalendar invokes method getChatMessageCalendar#ec8f2114 returning error if any.
func (c *Client) GetChatMessageCalendar(ctx context.Context, request *GetChatMessageCalendarRequest) (*MessageCalendar, error) {
	var result MessageCalendar

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

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

// GetCustomEmojiStickersRequest represents TL type `getCustomEmojiStickers#a5156885`.
type GetCustomEmojiStickersRequest struct {
	// Identifiers of custom emoji stickers. At most 200 custom emoji stickers can be
	// received simultaneously
	CustomEmojiIDs []int64
}

// GetCustomEmojiStickersRequestTypeID is TL type id of GetCustomEmojiStickersRequest.
const GetCustomEmojiStickersRequestTypeID = 0xa5156885

// Ensuring interfaces in compile-time for GetCustomEmojiStickersRequest.
var (
	_ bin.Encoder     = &GetCustomEmojiStickersRequest{}
	_ bin.Decoder     = &GetCustomEmojiStickersRequest{}
	_ bin.BareEncoder = &GetCustomEmojiStickersRequest{}
	_ bin.BareDecoder = &GetCustomEmojiStickersRequest{}
)

func (g *GetCustomEmojiStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.CustomEmojiIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCustomEmojiStickersRequest) String() string {
	if g == nil {
		return "GetCustomEmojiStickersRequest(nil)"
	}
	type Alias GetCustomEmojiStickersRequest
	return fmt.Sprintf("GetCustomEmojiStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCustomEmojiStickersRequest) TypeID() uint32 {
	return GetCustomEmojiStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCustomEmojiStickersRequest) TypeName() string {
	return "getCustomEmojiStickers"
}

// TypeInfo returns info about TL type.
func (g *GetCustomEmojiStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCustomEmojiStickers",
		ID:   GetCustomEmojiStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CustomEmojiIDs",
			SchemaName: "custom_emoji_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCustomEmojiStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCustomEmojiStickers#a5156885 as nil")
	}
	b.PutID(GetCustomEmojiStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCustomEmojiStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCustomEmojiStickers#a5156885 as nil")
	}
	b.PutInt(len(g.CustomEmojiIDs))
	for _, v := range g.CustomEmojiIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCustomEmojiStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCustomEmojiStickers#a5156885 to nil")
	}
	if err := b.ConsumeID(GetCustomEmojiStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCustomEmojiStickers#a5156885: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCustomEmojiStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCustomEmojiStickers#a5156885 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode getCustomEmojiStickers#a5156885: field custom_emoji_ids: %w", err)
		}

		if headerLen > 0 {
			g.CustomEmojiIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getCustomEmojiStickers#a5156885: field custom_emoji_ids: %w", err)
			}
			g.CustomEmojiIDs = append(g.CustomEmojiIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetCustomEmojiStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCustomEmojiStickers#a5156885 as nil")
	}
	b.ObjStart()
	b.PutID("getCustomEmojiStickers")
	b.Comma()
	b.FieldStart("custom_emoji_ids")
	b.ArrStart()
	for _, v := range g.CustomEmojiIDs {
		b.PutLong(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetCustomEmojiStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getCustomEmojiStickers#a5156885 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getCustomEmojiStickers"); err != nil {
				return fmt.Errorf("unable to decode getCustomEmojiStickers#a5156885: %w", err)
			}
		case "custom_emoji_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Long()
				if err != nil {
					return fmt.Errorf("unable to decode getCustomEmojiStickers#a5156885: field custom_emoji_ids: %w", err)
				}
				g.CustomEmojiIDs = append(g.CustomEmojiIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode getCustomEmojiStickers#a5156885: field custom_emoji_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomEmojiIDs returns value of CustomEmojiIDs field.
func (g *GetCustomEmojiStickersRequest) GetCustomEmojiIDs() (value []int64) {
	if g == nil {
		return
	}
	return g.CustomEmojiIDs
}

// GetCustomEmojiStickers invokes method getCustomEmojiStickers#a5156885 returning error if any.
func (c *Client) GetCustomEmojiStickers(ctx context.Context, customemojiids []int64) (*Stickers, error) {
	var result Stickers

	request := &GetCustomEmojiStickersRequest{
		CustomEmojiIDs: customemojiids,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

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

// RemoveStickerFromSetRequest represents TL type `removeStickerFromSet#61e1eea4`.
type RemoveStickerFromSetRequest struct {
	// Sticker
	Sticker InputFileClass
}

// RemoveStickerFromSetRequestTypeID is TL type id of RemoveStickerFromSetRequest.
const RemoveStickerFromSetRequestTypeID = 0x61e1eea4

// Ensuring interfaces in compile-time for RemoveStickerFromSetRequest.
var (
	_ bin.Encoder     = &RemoveStickerFromSetRequest{}
	_ bin.Decoder     = &RemoveStickerFromSetRequest{}
	_ bin.BareEncoder = &RemoveStickerFromSetRequest{}
	_ bin.BareDecoder = &RemoveStickerFromSetRequest{}
)

func (r *RemoveStickerFromSetRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveStickerFromSetRequest) String() string {
	if r == nil {
		return "RemoveStickerFromSetRequest(nil)"
	}
	type Alias RemoveStickerFromSetRequest
	return fmt.Sprintf("RemoveStickerFromSetRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveStickerFromSetRequest) TypeID() uint32 {
	return RemoveStickerFromSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveStickerFromSetRequest) TypeName() string {
	return "removeStickerFromSet"
}

// TypeInfo returns info about TL type.
func (r *RemoveStickerFromSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeStickerFromSet",
		ID:   RemoveStickerFromSetRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveStickerFromSetRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeStickerFromSet#61e1eea4 as nil")
	}
	b.PutID(RemoveStickerFromSetRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveStickerFromSetRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeStickerFromSet#61e1eea4 as nil")
	}
	if r.Sticker == nil {
		return fmt.Errorf("unable to encode removeStickerFromSet#61e1eea4: field sticker is nil")
	}
	if err := r.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode removeStickerFromSet#61e1eea4: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveStickerFromSetRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeStickerFromSet#61e1eea4 to nil")
	}
	if err := b.ConsumeID(RemoveStickerFromSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeStickerFromSet#61e1eea4: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveStickerFromSetRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeStickerFromSet#61e1eea4 to nil")
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode removeStickerFromSet#61e1eea4: field sticker: %w", err)
		}
		r.Sticker = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveStickerFromSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeStickerFromSet#61e1eea4 as nil")
	}
	b.ObjStart()
	b.PutID("removeStickerFromSet")
	b.FieldStart("sticker")
	if r.Sticker == nil {
		return fmt.Errorf("unable to encode removeStickerFromSet#61e1eea4: field sticker is nil")
	}
	if err := r.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode removeStickerFromSet#61e1eea4: field sticker: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveStickerFromSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeStickerFromSet#61e1eea4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeStickerFromSet"); err != nil {
				return fmt.Errorf("unable to decode removeStickerFromSet#61e1eea4: %w", err)
			}
		case "sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode removeStickerFromSet#61e1eea4: field sticker: %w", err)
			}
			r.Sticker = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSticker returns value of Sticker field.
func (r *RemoveStickerFromSetRequest) GetSticker() (value InputFileClass) {
	if r == nil {
		return
	}
	return r.Sticker
}

// RemoveStickerFromSet invokes method removeStickerFromSet#61e1eea4 returning error if any.
func (c *Client) RemoveStickerFromSet(ctx context.Context, sticker InputFileClass) error {
	var ok Ok

	request := &RemoveStickerFromSetRequest{
		Sticker: sticker,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

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

// GetWebPagePreviewRequest represents TL type `getWebPagePreview#222e062c`.
type GetWebPagePreviewRequest struct {
	// Message text with formatting
	Text FormattedText
}

// GetWebPagePreviewRequestTypeID is TL type id of GetWebPagePreviewRequest.
const GetWebPagePreviewRequestTypeID = 0x222e062c

// Ensuring interfaces in compile-time for GetWebPagePreviewRequest.
var (
	_ bin.Encoder     = &GetWebPagePreviewRequest{}
	_ bin.Decoder     = &GetWebPagePreviewRequest{}
	_ bin.BareEncoder = &GetWebPagePreviewRequest{}
	_ bin.BareDecoder = &GetWebPagePreviewRequest{}
)

func (g *GetWebPagePreviewRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Text.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetWebPagePreviewRequest) String() string {
	if g == nil {
		return "GetWebPagePreviewRequest(nil)"
	}
	type Alias GetWebPagePreviewRequest
	return fmt.Sprintf("GetWebPagePreviewRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetWebPagePreviewRequest) TypeID() uint32 {
	return GetWebPagePreviewRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetWebPagePreviewRequest) TypeName() string {
	return "getWebPagePreview"
}

// TypeInfo returns info about TL type.
func (g *GetWebPagePreviewRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getWebPagePreview",
		ID:   GetWebPagePreviewRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetWebPagePreviewRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getWebPagePreview#222e062c as nil")
	}
	b.PutID(GetWebPagePreviewRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetWebPagePreviewRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getWebPagePreview#222e062c as nil")
	}
	if err := g.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getWebPagePreview#222e062c: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetWebPagePreviewRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getWebPagePreview#222e062c to nil")
	}
	if err := b.ConsumeID(GetWebPagePreviewRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getWebPagePreview#222e062c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetWebPagePreviewRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getWebPagePreview#222e062c to nil")
	}
	{
		if err := g.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getWebPagePreview#222e062c: field text: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetWebPagePreviewRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getWebPagePreview#222e062c as nil")
	}
	b.ObjStart()
	b.PutID("getWebPagePreview")
	b.FieldStart("text")
	if err := g.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getWebPagePreview#222e062c: field text: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetWebPagePreviewRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getWebPagePreview#222e062c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getWebPagePreview"); err != nil {
				return fmt.Errorf("unable to decode getWebPagePreview#222e062c: %w", err)
			}
		case "text":
			if err := g.Text.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getWebPagePreview#222e062c: field text: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (g *GetWebPagePreviewRequest) GetText() (value FormattedText) {
	if g == nil {
		return
	}
	return g.Text
}

// GetWebPagePreview invokes method getWebPagePreview#222e062c returning error if any.
func (c *Client) GetWebPagePreview(ctx context.Context, text FormattedText) (*WebPage, error) {
	var result WebPage

	request := &GetWebPagePreviewRequest{
		Text: text,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

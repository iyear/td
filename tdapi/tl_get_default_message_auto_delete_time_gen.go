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

// GetDefaultMessageAutoDeleteTimeRequest represents TL type `getDefaultMessageAutoDeleteTime#e520759a`.
type GetDefaultMessageAutoDeleteTimeRequest struct {
}

// GetDefaultMessageAutoDeleteTimeRequestTypeID is TL type id of GetDefaultMessageAutoDeleteTimeRequest.
const GetDefaultMessageAutoDeleteTimeRequestTypeID = 0xe520759a

// Ensuring interfaces in compile-time for GetDefaultMessageAutoDeleteTimeRequest.
var (
	_ bin.Encoder     = &GetDefaultMessageAutoDeleteTimeRequest{}
	_ bin.Decoder     = &GetDefaultMessageAutoDeleteTimeRequest{}
	_ bin.BareEncoder = &GetDefaultMessageAutoDeleteTimeRequest{}
	_ bin.BareDecoder = &GetDefaultMessageAutoDeleteTimeRequest{}
)

func (g *GetDefaultMessageAutoDeleteTimeRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetDefaultMessageAutoDeleteTimeRequest) String() string {
	if g == nil {
		return "GetDefaultMessageAutoDeleteTimeRequest(nil)"
	}
	type Alias GetDefaultMessageAutoDeleteTimeRequest
	return fmt.Sprintf("GetDefaultMessageAutoDeleteTimeRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetDefaultMessageAutoDeleteTimeRequest) TypeID() uint32 {
	return GetDefaultMessageAutoDeleteTimeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetDefaultMessageAutoDeleteTimeRequest) TypeName() string {
	return "getDefaultMessageAutoDeleteTime"
}

// TypeInfo returns info about TL type.
func (g *GetDefaultMessageAutoDeleteTimeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getDefaultMessageAutoDeleteTime",
		ID:   GetDefaultMessageAutoDeleteTimeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetDefaultMessageAutoDeleteTimeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultMessageAutoDeleteTime#e520759a as nil")
	}
	b.PutID(GetDefaultMessageAutoDeleteTimeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetDefaultMessageAutoDeleteTimeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultMessageAutoDeleteTime#e520759a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetDefaultMessageAutoDeleteTimeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultMessageAutoDeleteTime#e520759a to nil")
	}
	if err := b.ConsumeID(GetDefaultMessageAutoDeleteTimeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getDefaultMessageAutoDeleteTime#e520759a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetDefaultMessageAutoDeleteTimeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultMessageAutoDeleteTime#e520759a to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetDefaultMessageAutoDeleteTimeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultMessageAutoDeleteTime#e520759a as nil")
	}
	b.ObjStart()
	b.PutID("getDefaultMessageAutoDeleteTime")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetDefaultMessageAutoDeleteTimeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultMessageAutoDeleteTime#e520759a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getDefaultMessageAutoDeleteTime"); err != nil {
				return fmt.Errorf("unable to decode getDefaultMessageAutoDeleteTime#e520759a: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDefaultMessageAutoDeleteTime invokes method getDefaultMessageAutoDeleteTime#e520759a returning error if any.
func (c *Client) GetDefaultMessageAutoDeleteTime(ctx context.Context) (*MessageAutoDeleteTime, error) {
	var result MessageAutoDeleteTime

	request := &GetDefaultMessageAutoDeleteTimeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
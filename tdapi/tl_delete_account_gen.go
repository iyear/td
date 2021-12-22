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

// DeleteAccountRequest represents TL type `deleteAccount#b84ad084`.
type DeleteAccountRequest struct {
	// The reason why the account was deleted; optional
	Reason string
}

// DeleteAccountRequestTypeID is TL type id of DeleteAccountRequest.
const DeleteAccountRequestTypeID = 0xb84ad084

// Ensuring interfaces in compile-time for DeleteAccountRequest.
var (
	_ bin.Encoder     = &DeleteAccountRequest{}
	_ bin.Decoder     = &DeleteAccountRequest{}
	_ bin.BareEncoder = &DeleteAccountRequest{}
	_ bin.BareDecoder = &DeleteAccountRequest{}
)

func (d *DeleteAccountRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Reason == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteAccountRequest) String() string {
	if d == nil {
		return "DeleteAccountRequest(nil)"
	}
	type Alias DeleteAccountRequest
	return fmt.Sprintf("DeleteAccountRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteAccountRequest) TypeID() uint32 {
	return DeleteAccountRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteAccountRequest) TypeName() string {
	return "deleteAccount"
}

// TypeInfo returns info about TL type.
func (d *DeleteAccountRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteAccount",
		ID:   DeleteAccountRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reason",
			SchemaName: "reason",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteAccountRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteAccount#b84ad084 as nil")
	}
	b.PutID(DeleteAccountRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteAccountRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteAccount#b84ad084 as nil")
	}
	b.PutString(d.Reason)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteAccountRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteAccount#b84ad084 to nil")
	}
	if err := b.ConsumeID(DeleteAccountRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteAccount#b84ad084: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteAccountRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteAccount#b84ad084 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode deleteAccount#b84ad084: field reason: %w", err)
		}
		d.Reason = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteAccountRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteAccount#b84ad084 as nil")
	}
	b.ObjStart()
	b.PutID("deleteAccount")
	b.FieldStart("reason")
	b.PutString(d.Reason)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteAccountRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteAccount#b84ad084 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteAccount"); err != nil {
				return fmt.Errorf("unable to decode deleteAccount#b84ad084: %w", err)
			}
		case "reason":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode deleteAccount#b84ad084: field reason: %w", err)
			}
			d.Reason = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReason returns value of Reason field.
func (d *DeleteAccountRequest) GetReason() (value string) {
	if d == nil {
		return
	}
	return d.Reason
}

// DeleteAccount invokes method deleteAccount#b84ad084 returning error if any.
func (c *Client) DeleteAccount(ctx context.Context, reason string) error {
	var ok Ok

	request := &DeleteAccountRequest{
		Reason: reason,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

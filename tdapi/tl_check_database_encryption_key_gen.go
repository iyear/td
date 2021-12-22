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

// CheckDatabaseEncryptionKeyRequest represents TL type `checkDatabaseEncryptionKey#3cb92f9b`.
type CheckDatabaseEncryptionKeyRequest struct {
	// Encryption key to check or set up
	EncryptionKey []byte
}

// CheckDatabaseEncryptionKeyRequestTypeID is TL type id of CheckDatabaseEncryptionKeyRequest.
const CheckDatabaseEncryptionKeyRequestTypeID = 0x3cb92f9b

// Ensuring interfaces in compile-time for CheckDatabaseEncryptionKeyRequest.
var (
	_ bin.Encoder     = &CheckDatabaseEncryptionKeyRequest{}
	_ bin.Decoder     = &CheckDatabaseEncryptionKeyRequest{}
	_ bin.BareEncoder = &CheckDatabaseEncryptionKeyRequest{}
	_ bin.BareDecoder = &CheckDatabaseEncryptionKeyRequest{}
)

func (c *CheckDatabaseEncryptionKeyRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.EncryptionKey == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckDatabaseEncryptionKeyRequest) String() string {
	if c == nil {
		return "CheckDatabaseEncryptionKeyRequest(nil)"
	}
	type Alias CheckDatabaseEncryptionKeyRequest
	return fmt.Sprintf("CheckDatabaseEncryptionKeyRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckDatabaseEncryptionKeyRequest) TypeID() uint32 {
	return CheckDatabaseEncryptionKeyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckDatabaseEncryptionKeyRequest) TypeName() string {
	return "checkDatabaseEncryptionKey"
}

// TypeInfo returns info about TL type.
func (c *CheckDatabaseEncryptionKeyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkDatabaseEncryptionKey",
		ID:   CheckDatabaseEncryptionKeyRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EncryptionKey",
			SchemaName: "encryption_key",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckDatabaseEncryptionKeyRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkDatabaseEncryptionKey#3cb92f9b as nil")
	}
	b.PutID(CheckDatabaseEncryptionKeyRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckDatabaseEncryptionKeyRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkDatabaseEncryptionKey#3cb92f9b as nil")
	}
	b.PutBytes(c.EncryptionKey)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckDatabaseEncryptionKeyRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkDatabaseEncryptionKey#3cb92f9b to nil")
	}
	if err := b.ConsumeID(CheckDatabaseEncryptionKeyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkDatabaseEncryptionKey#3cb92f9b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckDatabaseEncryptionKeyRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkDatabaseEncryptionKey#3cb92f9b to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode checkDatabaseEncryptionKey#3cb92f9b: field encryption_key: %w", err)
		}
		c.EncryptionKey = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckDatabaseEncryptionKeyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkDatabaseEncryptionKey#3cb92f9b as nil")
	}
	b.ObjStart()
	b.PutID("checkDatabaseEncryptionKey")
	b.FieldStart("encryption_key")
	b.PutBytes(c.EncryptionKey)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckDatabaseEncryptionKeyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkDatabaseEncryptionKey#3cb92f9b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkDatabaseEncryptionKey"); err != nil {
				return fmt.Errorf("unable to decode checkDatabaseEncryptionKey#3cb92f9b: %w", err)
			}
		case "encryption_key":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode checkDatabaseEncryptionKey#3cb92f9b: field encryption_key: %w", err)
			}
			c.EncryptionKey = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEncryptionKey returns value of EncryptionKey field.
func (c *CheckDatabaseEncryptionKeyRequest) GetEncryptionKey() (value []byte) {
	if c == nil {
		return
	}
	return c.EncryptionKey
}

// CheckDatabaseEncryptionKey invokes method checkDatabaseEncryptionKey#3cb92f9b returning error if any.
func (c *Client) CheckDatabaseEncryptionKey(ctx context.Context, encryptionkey []byte) error {
	var ok Ok

	request := &CheckDatabaseEncryptionKeyRequest{
		EncryptionKey: encryptionkey,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

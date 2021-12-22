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

// CheckPasswordRecoveryCodeRequest represents TL type `checkPasswordRecoveryCode#f4081e18`.
type CheckPasswordRecoveryCodeRequest struct {
	// Recovery code to check
	RecoveryCode string
}

// CheckPasswordRecoveryCodeRequestTypeID is TL type id of CheckPasswordRecoveryCodeRequest.
const CheckPasswordRecoveryCodeRequestTypeID = 0xf4081e18

// Ensuring interfaces in compile-time for CheckPasswordRecoveryCodeRequest.
var (
	_ bin.Encoder     = &CheckPasswordRecoveryCodeRequest{}
	_ bin.Decoder     = &CheckPasswordRecoveryCodeRequest{}
	_ bin.BareEncoder = &CheckPasswordRecoveryCodeRequest{}
	_ bin.BareDecoder = &CheckPasswordRecoveryCodeRequest{}
)

func (c *CheckPasswordRecoveryCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.RecoveryCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckPasswordRecoveryCodeRequest) String() string {
	if c == nil {
		return "CheckPasswordRecoveryCodeRequest(nil)"
	}
	type Alias CheckPasswordRecoveryCodeRequest
	return fmt.Sprintf("CheckPasswordRecoveryCodeRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckPasswordRecoveryCodeRequest) TypeID() uint32 {
	return CheckPasswordRecoveryCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckPasswordRecoveryCodeRequest) TypeName() string {
	return "checkPasswordRecoveryCode"
}

// TypeInfo returns info about TL type.
func (c *CheckPasswordRecoveryCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkPasswordRecoveryCode",
		ID:   CheckPasswordRecoveryCodeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RecoveryCode",
			SchemaName: "recovery_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckPasswordRecoveryCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkPasswordRecoveryCode#f4081e18 as nil")
	}
	b.PutID(CheckPasswordRecoveryCodeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckPasswordRecoveryCodeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkPasswordRecoveryCode#f4081e18 as nil")
	}
	b.PutString(c.RecoveryCode)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckPasswordRecoveryCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkPasswordRecoveryCode#f4081e18 to nil")
	}
	if err := b.ConsumeID(CheckPasswordRecoveryCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkPasswordRecoveryCode#f4081e18: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckPasswordRecoveryCodeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkPasswordRecoveryCode#f4081e18 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkPasswordRecoveryCode#f4081e18: field recovery_code: %w", err)
		}
		c.RecoveryCode = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckPasswordRecoveryCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkPasswordRecoveryCode#f4081e18 as nil")
	}
	b.ObjStart()
	b.PutID("checkPasswordRecoveryCode")
	b.FieldStart("recovery_code")
	b.PutString(c.RecoveryCode)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckPasswordRecoveryCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkPasswordRecoveryCode#f4081e18 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkPasswordRecoveryCode"); err != nil {
				return fmt.Errorf("unable to decode checkPasswordRecoveryCode#f4081e18: %w", err)
			}
		case "recovery_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkPasswordRecoveryCode#f4081e18: field recovery_code: %w", err)
			}
			c.RecoveryCode = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRecoveryCode returns value of RecoveryCode field.
func (c *CheckPasswordRecoveryCodeRequest) GetRecoveryCode() (value string) {
	if c == nil {
		return
	}
	return c.RecoveryCode
}

// CheckPasswordRecoveryCode invokes method checkPasswordRecoveryCode#f4081e18 returning error if any.
func (c *Client) CheckPasswordRecoveryCode(ctx context.Context, recoverycode string) error {
	var ok Ok

	request := &CheckPasswordRecoveryCodeRequest{
		RecoveryCode: recoverycode,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

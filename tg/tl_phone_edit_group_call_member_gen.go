// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// PhoneEditGroupCallMemberRequest represents TL type `phone.editGroupCallMember#a5e76cd8`.
//
// See https://core.telegram.org/method/phone.editGroupCallMember for reference.
type PhoneEditGroupCallMemberRequest struct {
	// Flags field of PhoneEditGroupCallMemberRequest.
	Flags bin.Fields `tl:"flags"`
	// Muted field of PhoneEditGroupCallMemberRequest.
	Muted bool `tl:"muted"`
	// Call field of PhoneEditGroupCallMemberRequest.
	Call InputGroupCall `tl:"call"`
	// UserID field of PhoneEditGroupCallMemberRequest.
	UserID InputUserClass `tl:"user_id"`
	// Volume field of PhoneEditGroupCallMemberRequest.
	//
	// Use SetVolume and GetVolume helpers.
	Volume int `tl:"volume"`
}

// PhoneEditGroupCallMemberRequestTypeID is TL type id of PhoneEditGroupCallMemberRequest.
const PhoneEditGroupCallMemberRequestTypeID = 0xa5e76cd8

func (e *PhoneEditGroupCallMemberRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Muted == false) {
		return false
	}
	if !(e.Call.Zero()) {
		return false
	}
	if !(e.UserID == nil) {
		return false
	}
	if !(e.Volume == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *PhoneEditGroupCallMemberRequest) String() string {
	if e == nil {
		return "PhoneEditGroupCallMemberRequest(nil)"
	}
	type Alias PhoneEditGroupCallMemberRequest
	return fmt.Sprintf("PhoneEditGroupCallMemberRequest%+v", Alias(*e))
}

// FillFrom fills PhoneEditGroupCallMemberRequest from given interface.
func (e *PhoneEditGroupCallMemberRequest) FillFrom(from interface {
	GetMuted() (value bool)
	GetCall() (value InputGroupCall)
	GetUserID() (value InputUserClass)
	GetVolume() (value int, ok bool)
}) {
	e.Muted = from.GetMuted()
	e.Call = from.GetCall()
	e.UserID = from.GetUserID()
	if val, ok := from.GetVolume(); ok {
		e.Volume = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *PhoneEditGroupCallMemberRequest) TypeID() uint32 {
	return PhoneEditGroupCallMemberRequestTypeID
}

// TypeName returns name of type in TL schema.
func (e *PhoneEditGroupCallMemberRequest) TypeName() string {
	return "phone.editGroupCallMember"
}

// Encode implements bin.Encoder.
func (e *PhoneEditGroupCallMemberRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.editGroupCallMember#a5e76cd8 as nil")
	}
	b.PutID(PhoneEditGroupCallMemberRequestTypeID)
	if !(e.Muted == false) {
		e.Flags.Set(0)
	}
	if !(e.Volume == 0) {
		e.Flags.Set(1)
	}
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#a5e76cd8: field flags: %w", err)
	}
	if err := e.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#a5e76cd8: field call: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#a5e76cd8: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#a5e76cd8: field user_id: %w", err)
	}
	if e.Flags.Has(1) {
		b.PutInt(e.Volume)
	}
	return nil
}

// SetMuted sets value of Muted conditional field.
func (e *PhoneEditGroupCallMemberRequest) SetMuted(value bool) {
	if value {
		e.Flags.Set(0)
		e.Muted = true
	} else {
		e.Flags.Unset(0)
		e.Muted = false
	}
}

// GetMuted returns value of Muted conditional field.
func (e *PhoneEditGroupCallMemberRequest) GetMuted() (value bool) {
	return e.Flags.Has(0)
}

// GetCall returns value of Call field.
func (e *PhoneEditGroupCallMemberRequest) GetCall() (value InputGroupCall) {
	return e.Call
}

// GetUserID returns value of UserID field.
func (e *PhoneEditGroupCallMemberRequest) GetUserID() (value InputUserClass) {
	return e.UserID
}

// SetVolume sets value of Volume conditional field.
func (e *PhoneEditGroupCallMemberRequest) SetVolume(value int) {
	e.Flags.Set(1)
	e.Volume = value
}

// GetVolume returns value of Volume conditional field and
// boolean which is true if field was set.
func (e *PhoneEditGroupCallMemberRequest) GetVolume() (value int, ok bool) {
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.Volume, true
}

// Decode implements bin.Decoder.
func (e *PhoneEditGroupCallMemberRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.editGroupCallMember#a5e76cd8 to nil")
	}
	if err := b.ConsumeID(PhoneEditGroupCallMemberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.editGroupCallMember#a5e76cd8: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#a5e76cd8: field flags: %w", err)
		}
	}
	e.Muted = e.Flags.Has(0)
	{
		if err := e.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#a5e76cd8: field call: %w", err)
		}
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#a5e76cd8: field user_id: %w", err)
		}
		e.UserID = value
	}
	if e.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#a5e76cd8: field volume: %w", err)
		}
		e.Volume = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneEditGroupCallMemberRequest.
var (
	_ bin.Encoder = &PhoneEditGroupCallMemberRequest{}
	_ bin.Decoder = &PhoneEditGroupCallMemberRequest{}
)

// PhoneEditGroupCallMember invokes method phone.editGroupCallMember#a5e76cd8 returning error if any.
//
// See https://core.telegram.org/method/phone.editGroupCallMember for reference.
func (c *Client) PhoneEditGroupCallMember(ctx context.Context, request *PhoneEditGroupCallMemberRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

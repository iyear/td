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

// HelpSupport represents TL type `help.support#17c6b5f6`.
// Info on support user.
//
// See https://core.telegram.org/constructor/help.support for reference.
type HelpSupport struct {
	// Phone number
	PhoneNumber string `tl:"phone_number"`
	// User
	User UserClass `tl:"user"`
}

// HelpSupportTypeID is TL type id of HelpSupport.
const HelpSupportTypeID = 0x17c6b5f6

func (s *HelpSupport) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.User == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *HelpSupport) String() string {
	if s == nil {
		return "HelpSupport(nil)"
	}
	type Alias HelpSupport
	return fmt.Sprintf("HelpSupport%+v", Alias(*s))
}

// FillFrom fills HelpSupport from given interface.
func (s *HelpSupport) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetUser() (value UserClass)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.User = from.GetUser()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *HelpSupport) TypeID() uint32 {
	return HelpSupportTypeID
}

// TypeName returns name of type in TL schema.
func (s *HelpSupport) TypeName() string {
	return "help.support"
}

// Encode implements bin.Encoder.
func (s *HelpSupport) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode help.support#17c6b5f6 as nil")
	}
	b.PutID(HelpSupportTypeID)
	b.PutString(s.PhoneNumber)
	if s.User == nil {
		return fmt.Errorf("unable to encode help.support#17c6b5f6: field user is nil")
	}
	if err := s.User.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.support#17c6b5f6: field user: %w", err)
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *HelpSupport) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetUser returns value of User field.
func (s *HelpSupport) GetUser() (value UserClass) {
	return s.User
}

// GetUserAsNotEmpty returns mapped value of User field.
func (s *HelpSupport) GetUserAsNotEmpty() (*User, bool) {
	return s.User.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (s *HelpSupport) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode help.support#17c6b5f6 to nil")
	}
	if err := b.ConsumeID(HelpSupportTypeID); err != nil {
		return fmt.Errorf("unable to decode help.support#17c6b5f6: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.support#17c6b5f6: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.support#17c6b5f6: field user: %w", err)
		}
		s.User = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpSupport.
var (
	_ bin.Encoder = &HelpSupport{}
	_ bin.Decoder = &HelpSupport{}
)

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

// AccountSaveThemeRequest represents TL type `account.saveTheme#f257106c`.
// Save a theme
//
// See https://core.telegram.org/method/account.saveTheme for reference.
type AccountSaveThemeRequest struct {
	// Theme to save
	Theme InputThemeClass `tl:"theme"`
	// Unsave
	Unsave bool `tl:"unsave"`
}

// AccountSaveThemeRequestTypeID is TL type id of AccountSaveThemeRequest.
const AccountSaveThemeRequestTypeID = 0xf257106c

func (s *AccountSaveThemeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Theme == nil) {
		return false
	}
	if !(s.Unsave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSaveThemeRequest) String() string {
	if s == nil {
		return "AccountSaveThemeRequest(nil)"
	}
	type Alias AccountSaveThemeRequest
	return fmt.Sprintf("AccountSaveThemeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSaveThemeRequest from given interface.
func (s *AccountSaveThemeRequest) FillFrom(from interface {
	GetTheme() (value InputThemeClass)
	GetUnsave() (value bool)
}) {
	s.Theme = from.GetTheme()
	s.Unsave = from.GetUnsave()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSaveThemeRequest) TypeID() uint32 {
	return AccountSaveThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (s *AccountSaveThemeRequest) TypeName() string {
	return "account.saveTheme"
}

// Encode implements bin.Encoder.
func (s *AccountSaveThemeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveTheme#f257106c as nil")
	}
	b.PutID(AccountSaveThemeRequestTypeID)
	if s.Theme == nil {
		return fmt.Errorf("unable to encode account.saveTheme#f257106c: field theme is nil")
	}
	if err := s.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveTheme#f257106c: field theme: %w", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// GetTheme returns value of Theme field.
func (s *AccountSaveThemeRequest) GetTheme() (value InputThemeClass) {
	return s.Theme
}

// GetUnsave returns value of Unsave field.
func (s *AccountSaveThemeRequest) GetUnsave() (value bool) {
	return s.Unsave
}

// Decode implements bin.Decoder.
func (s *AccountSaveThemeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveTheme#f257106c to nil")
	}
	if err := b.ConsumeID(AccountSaveThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveTheme#f257106c: %w", err)
	}
	{
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveTheme#f257106c: field theme: %w", err)
		}
		s.Theme = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveTheme#f257106c: field unsave: %w", err)
		}
		s.Unsave = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveThemeRequest.
var (
	_ bin.Encoder = &AccountSaveThemeRequest{}
	_ bin.Decoder = &AccountSaveThemeRequest{}
)

// AccountSaveTheme invokes method account.saveTheme#f257106c returning error if any.
// Save a theme
//
// See https://core.telegram.org/method/account.saveTheme for reference.
func (c *Client) AccountSaveTheme(ctx context.Context, request *AccountSaveThemeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}

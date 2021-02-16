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

// AccountGetPrivacyRequest represents TL type `account.getPrivacy#dadbc950`.
// Get privacy settings of current account
//
// See https://core.telegram.org/method/account.getPrivacy for reference.
type AccountGetPrivacyRequest struct {
	// Peer category whose privacy settings should be fetched
	Key InputPrivacyKeyClass
}

// AccountGetPrivacyRequestTypeID is TL type id of AccountGetPrivacyRequest.
const AccountGetPrivacyRequestTypeID = 0xdadbc950

func (g *AccountGetPrivacyRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Key == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetPrivacyRequest) String() string {
	if g == nil {
		return "AccountGetPrivacyRequest(nil)"
	}
	type Alias AccountGetPrivacyRequest
	return fmt.Sprintf("AccountGetPrivacyRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetPrivacyRequest) TypeID() uint32 {
	return AccountGetPrivacyRequestTypeID
}

// Encode implements bin.Encoder.
func (g *AccountGetPrivacyRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getPrivacy#dadbc950 as nil")
	}
	b.PutID(AccountGetPrivacyRequestTypeID)
	if g.Key == nil {
		return fmt.Errorf("unable to encode account.getPrivacy#dadbc950: field key is nil")
	}
	if err := g.Key.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getPrivacy#dadbc950: field key: %w", err)
	}
	return nil
}

// GetKey returns value of Key field.
func (g *AccountGetPrivacyRequest) GetKey() (value InputPrivacyKeyClass) {
	return g.Key
}

// Decode implements bin.Decoder.
func (g *AccountGetPrivacyRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getPrivacy#dadbc950 to nil")
	}
	if err := b.ConsumeID(AccountGetPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getPrivacy#dadbc950: %w", err)
	}
	{
		value, err := DecodeInputPrivacyKey(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getPrivacy#dadbc950: field key: %w", err)
		}
		g.Key = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetPrivacyRequest.
var (
	_ bin.Encoder = &AccountGetPrivacyRequest{}
	_ bin.Decoder = &AccountGetPrivacyRequest{}
)

// AccountGetPrivacy invokes method account.getPrivacy#dadbc950 returning error if any.
// Get privacy settings of current account
//
// Possible errors:
//  400 PRIVACY_KEY_INVALID: The privacy key is invalid
//
// See https://core.telegram.org/method/account.getPrivacy for reference.
func (c *Client) AccountGetPrivacy(ctx context.Context, key InputPrivacyKeyClass) (*AccountPrivacyRules, error) {
	var result AccountPrivacyRules

	request := &AccountGetPrivacyRequest{
		Key: key,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

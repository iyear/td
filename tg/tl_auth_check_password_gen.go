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

// AuthCheckPasswordRequest represents TL type `auth.checkPassword#d18b4d16`.
// Try logging to an account protected by a 2FA password¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// See https://core.telegram.org/method/auth.checkPassword for reference.
type AuthCheckPasswordRequest struct {
	// The account's password (see SRP¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	Password InputCheckPasswordSRPClass
}

// AuthCheckPasswordRequestTypeID is TL type id of AuthCheckPasswordRequest.
const AuthCheckPasswordRequestTypeID = 0xd18b4d16

func (c *AuthCheckPasswordRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCheckPasswordRequest) String() string {
	if c == nil {
		return "AuthCheckPasswordRequest(nil)"
	}
	type Alias AuthCheckPasswordRequest
	return fmt.Sprintf("AuthCheckPasswordRequest%+v", Alias(*c))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AuthCheckPasswordRequest) TypeID() uint32 {
	return AuthCheckPasswordRequestTypeID
}

// Encode implements bin.Encoder.
func (c *AuthCheckPasswordRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.checkPassword#d18b4d16 as nil")
	}
	b.PutID(AuthCheckPasswordRequestTypeID)
	if c.Password == nil {
		return fmt.Errorf("unable to encode auth.checkPassword#d18b4d16: field password is nil")
	}
	if err := c.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.checkPassword#d18b4d16: field password: %w", err)
	}
	return nil
}

// GetPassword returns value of Password field.
func (c *AuthCheckPasswordRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	return c.Password
}

// Decode implements bin.Decoder.
func (c *AuthCheckPasswordRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.checkPassword#d18b4d16 to nil")
	}
	if err := b.ConsumeID(AuthCheckPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.checkPassword#d18b4d16: %w", err)
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.checkPassword#d18b4d16: field password: %w", err)
		}
		c.Password = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthCheckPasswordRequest.
var (
	_ bin.Encoder = &AuthCheckPasswordRequest{}
	_ bin.Decoder = &AuthCheckPasswordRequest{}
)

// AuthCheckPassword invokes method auth.checkPassword#d18b4d16 returning error if any.
// Try logging to an account protected by a 2FA password¹.
//
// Links:
//  1) https://core.telegram.org/api/srp
//
// Possible errors:
//  400 PASSWORD_HASH_INVALID: The provided password isn't valid
//  400 SRP_ID_INVALID: Invalid SRP ID provided
//  400 SRP_PASSWORD_CHANGED: Password has changed
//
// See https://core.telegram.org/method/auth.checkPassword for reference.
func (c *Client) AuthCheckPassword(ctx context.Context, password InputCheckPasswordSRPClass) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	request := &AuthCheckPasswordRequest{
		Password: password,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}

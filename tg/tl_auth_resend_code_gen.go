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

// AuthResendCodeRequest represents TL type `auth.resendCode#3ef1a9bf`.
// Resend the login code via another medium, the phone code type is determined by the return value of the previous auth.sendCode/auth.resendCode: see login¹ for more info.
//
// Links:
//  1) https://core.telegram.org/api/auth
//
// See https://core.telegram.org/method/auth.resendCode for reference.
type AuthResendCodeRequest struct {
	// The phone number
	PhoneNumber string
	// The phone code hash obtained from auth.sendCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.sendCode
	PhoneCodeHash string
}

// AuthResendCodeRequestTypeID is TL type id of AuthResendCodeRequest.
const AuthResendCodeRequestTypeID = 0x3ef1a9bf

func (r *AuthResendCodeRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.PhoneNumber == "") {
		return false
	}
	if !(r.PhoneCodeHash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthResendCodeRequest) String() string {
	if r == nil {
		return "AuthResendCodeRequest(nil)"
	}
	type Alias AuthResendCodeRequest
	return fmt.Sprintf("AuthResendCodeRequest%+v", Alias(*r))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *AuthResendCodeRequest) TypeID() uint32 {
	return AuthResendCodeRequestTypeID
}

// Encode implements bin.Encoder.
func (r *AuthResendCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resendCode#3ef1a9bf as nil")
	}
	b.PutID(AuthResendCodeRequestTypeID)
	b.PutString(r.PhoneNumber)
	b.PutString(r.PhoneCodeHash)
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (r *AuthResendCodeRequest) GetPhoneNumber() (value string) {
	return r.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (r *AuthResendCodeRequest) GetPhoneCodeHash() (value string) {
	return r.PhoneCodeHash
}

// Decode implements bin.Decoder.
func (r *AuthResendCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resendCode#3ef1a9bf to nil")
	}
	if err := b.ConsumeID(AuthResendCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.resendCode#3ef1a9bf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#3ef1a9bf: field phone_number: %w", err)
		}
		r.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resendCode#3ef1a9bf: field phone_code_hash: %w", err)
		}
		r.PhoneCodeHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthResendCodeRequest.
var (
	_ bin.Encoder = &AuthResendCodeRequest{}
	_ bin.Decoder = &AuthResendCodeRequest{}
)

// AuthResendCode invokes method auth.resendCode#3ef1a9bf returning error if any.
// Resend the login code via another medium, the phone code type is determined by the return value of the previous auth.sendCode/auth.resendCode: see login¹ for more info.
//
// Links:
//  1) https://core.telegram.org/api/auth
//
// Possible errors:
//  400 PHONE_CODE_EXPIRED: The phone code you provided has expired, this may happen if it was sent to any chat on telegram (if the code is sent through a telegram chat (not the official account) to avoid it append or prepend to the code some chars)
//  400 PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
//  400 PHONE_NUMBER_INVALID: The phone number is invalid
//
// See https://core.telegram.org/method/auth.resendCode for reference.
func (c *Client) AuthResendCode(ctx context.Context, request *AuthResendCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

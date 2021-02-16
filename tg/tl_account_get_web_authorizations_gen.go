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

// AccountGetWebAuthorizationsRequest represents TL type `account.getWebAuthorizations#182e6d6f`.
// Get web login widget¹ authorizations
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.getWebAuthorizations for reference.
type AccountGetWebAuthorizationsRequest struct {
}

// AccountGetWebAuthorizationsRequestTypeID is TL type id of AccountGetWebAuthorizationsRequest.
const AccountGetWebAuthorizationsRequestTypeID = 0x182e6d6f

func (g *AccountGetWebAuthorizationsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetWebAuthorizationsRequest) String() string {
	if g == nil {
		return "AccountGetWebAuthorizationsRequest(nil)"
	}
	type Alias AccountGetWebAuthorizationsRequest
	return fmt.Sprintf("AccountGetWebAuthorizationsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetWebAuthorizationsRequest) TypeID() uint32 {
	return AccountGetWebAuthorizationsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *AccountGetWebAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWebAuthorizations#182e6d6f as nil")
	}
	b.PutID(AccountGetWebAuthorizationsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetWebAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWebAuthorizations#182e6d6f to nil")
	}
	if err := b.ConsumeID(AccountGetWebAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWebAuthorizations#182e6d6f: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetWebAuthorizationsRequest.
var (
	_ bin.Encoder = &AccountGetWebAuthorizationsRequest{}
	_ bin.Decoder = &AccountGetWebAuthorizationsRequest{}
)

// AccountGetWebAuthorizations invokes method account.getWebAuthorizations#182e6d6f returning error if any.
// Get web login widget¹ authorizations
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.getWebAuthorizations for reference.
func (c *Client) AccountGetWebAuthorizations(ctx context.Context) (*AccountWebAuthorizations, error) {
	var result AccountWebAuthorizations

	request := &AccountGetWebAuthorizationsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

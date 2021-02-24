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

// AccountResetWallPapersRequest represents TL type `account.resetWallPapers#bb3b9804`.
// Delete installed wallpapers
//
// See https://core.telegram.org/method/account.resetWallPapers for reference.
type AccountResetWallPapersRequest struct {
}

// AccountResetWallPapersRequestTypeID is TL type id of AccountResetWallPapersRequest.
const AccountResetWallPapersRequestTypeID = 0xbb3b9804

func (r *AccountResetWallPapersRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetWallPapersRequest) String() string {
	if r == nil {
		return "AccountResetWallPapersRequest(nil)"
	}
	type Alias AccountResetWallPapersRequest
	return fmt.Sprintf("AccountResetWallPapersRequest%+v", Alias(*r))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *AccountResetWallPapersRequest) TypeID() uint32 {
	return AccountResetWallPapersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (r *AccountResetWallPapersRequest) TypeName() string {
	return "account.resetWallPapers"
}

// Encode implements bin.Encoder.
func (r *AccountResetWallPapersRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWallPapers#bb3b9804 as nil")
	}
	b.PutID(AccountResetWallPapersRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetWallPapersRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWallPapers#bb3b9804 to nil")
	}
	if err := b.ConsumeID(AccountResetWallPapersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWallPapers#bb3b9804: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetWallPapersRequest.
var (
	_ bin.Encoder = &AccountResetWallPapersRequest{}
	_ bin.Decoder = &AccountResetWallPapersRequest{}
)

// AccountResetWallPapers invokes method account.resetWallPapers#bb3b9804 returning error if any.
// Delete installed wallpapers
//
// See https://core.telegram.org/method/account.resetWallPapers for reference.
func (c *Client) AccountResetWallPapers(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountResetWallPapersRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}

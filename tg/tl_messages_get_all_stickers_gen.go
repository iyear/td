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

// MessagesGetAllStickersRequest represents TL type `messages.getAllStickers#1c9618b1`.
// Get all installed stickers
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
type MessagesGetAllStickersRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
}

// MessagesGetAllStickersRequestTypeID is TL type id of MessagesGetAllStickersRequest.
const MessagesGetAllStickersRequestTypeID = 0x1c9618b1

func (g *MessagesGetAllStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAllStickersRequest) String() string {
	if g == nil {
		return "MessagesGetAllStickersRequest(nil)"
	}
	type Alias MessagesGetAllStickersRequest
	return fmt.Sprintf("MessagesGetAllStickersRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetAllStickersRequest) TypeID() uint32 {
	return MessagesGetAllStickersRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetAllStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllStickers#1c9618b1 as nil")
	}
	b.PutID(MessagesGetAllStickersRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetAllStickersRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllStickers#1c9618b1 to nil")
	}
	if err := b.ConsumeID(MessagesGetAllStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAllStickers#1c9618b1: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getAllStickers#1c9618b1: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllStickersRequest.
var (
	_ bin.Encoder = &MessagesGetAllStickersRequest{}
	_ bin.Decoder = &MessagesGetAllStickersRequest{}
)

// MessagesGetAllStickers invokes method messages.getAllStickers#1c9618b1 returning error if any.
// Get all installed stickers
//
// See https://core.telegram.org/method/messages.getAllStickers for reference.
func (c *Client) MessagesGetAllStickers(ctx context.Context, hash int) (MessagesAllStickersClass, error) {
	var result MessagesAllStickersBox

	request := &MessagesGetAllStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AllStickers, nil
}

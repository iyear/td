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

// MessagesGetAllDraftsRequest represents TL type `messages.getAllDrafts#6a3f8d65`.
// Save get all message drafts¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.getAllDrafts for reference.
type MessagesGetAllDraftsRequest struct {
}

// MessagesGetAllDraftsRequestTypeID is TL type id of MessagesGetAllDraftsRequest.
const MessagesGetAllDraftsRequestTypeID = 0x6a3f8d65

func (g *MessagesGetAllDraftsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetAllDraftsRequest) String() string {
	if g == nil {
		return "MessagesGetAllDraftsRequest(nil)"
	}
	type Alias MessagesGetAllDraftsRequest
	return fmt.Sprintf("MessagesGetAllDraftsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetAllDraftsRequest) TypeID() uint32 {
	return MessagesGetAllDraftsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *MessagesGetAllDraftsRequest) TypeName() string {
	return "messages.getAllDrafts"
}

// Encode implements bin.Encoder.
func (g *MessagesGetAllDraftsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getAllDrafts#6a3f8d65 as nil")
	}
	b.PutID(MessagesGetAllDraftsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetAllDraftsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getAllDrafts#6a3f8d65 to nil")
	}
	if err := b.ConsumeID(MessagesGetAllDraftsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getAllDrafts#6a3f8d65: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetAllDraftsRequest.
var (
	_ bin.Encoder = &MessagesGetAllDraftsRequest{}
	_ bin.Decoder = &MessagesGetAllDraftsRequest{}
)

// MessagesGetAllDrafts invokes method messages.getAllDrafts#6a3f8d65 returning error if any.
// Save get all message drafts¹.
//
// Links:
//  1) https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.getAllDrafts for reference.
func (c *Client) MessagesGetAllDrafts(ctx context.Context) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesGetAllDraftsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

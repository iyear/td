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

// MessagesReadMessageContentsRequest represents TL type `messages.readMessageContents#36a73f77`.
// Notifies the sender about the recipient having listened a voice message or watched a video.
//
// See https://core.telegram.org/method/messages.readMessageContents for reference.
type MessagesReadMessageContentsRequest struct {
	// Message ID list
	ID []int `tl:"id"`
}

// MessagesReadMessageContentsRequestTypeID is TL type id of MessagesReadMessageContentsRequest.
const MessagesReadMessageContentsRequestTypeID = 0x36a73f77

func (r *MessagesReadMessageContentsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReadMessageContentsRequest) String() string {
	if r == nil {
		return "MessagesReadMessageContentsRequest(nil)"
	}
	type Alias MessagesReadMessageContentsRequest
	return fmt.Sprintf("MessagesReadMessageContentsRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReadMessageContentsRequest from given interface.
func (r *MessagesReadMessageContentsRequest) FillFrom(from interface {
	GetID() (value []int)
}) {
	r.ID = from.GetID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *MessagesReadMessageContentsRequest) TypeID() uint32 {
	return MessagesReadMessageContentsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (r *MessagesReadMessageContentsRequest) TypeName() string {
	return "messages.readMessageContents"
}

// Encode implements bin.Encoder.
func (r *MessagesReadMessageContentsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.readMessageContents#36a73f77 as nil")
	}
	b.PutID(MessagesReadMessageContentsRequestTypeID)
	b.PutVectorHeader(len(r.ID))
	for _, v := range r.ID {
		b.PutInt(v)
	}
	return nil
}

// GetID returns value of ID field.
func (r *MessagesReadMessageContentsRequest) GetID() (value []int) {
	return r.ID
}

// Decode implements bin.Decoder.
func (r *MessagesReadMessageContentsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.readMessageContents#36a73f77 to nil")
	}
	if err := b.ConsumeID(MessagesReadMessageContentsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.readMessageContents#36a73f77: field id: %w", err)
			}
			r.ID = append(r.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReadMessageContentsRequest.
var (
	_ bin.Encoder = &MessagesReadMessageContentsRequest{}
	_ bin.Decoder = &MessagesReadMessageContentsRequest{}
)

// MessagesReadMessageContents invokes method messages.readMessageContents#36a73f77 returning error if any.
// Notifies the sender about the recipient having listened a voice message or watched a video.
//
// See https://core.telegram.org/method/messages.readMessageContents for reference.
func (c *Client) MessagesReadMessageContents(ctx context.Context, id []int) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages

	request := &MessagesReadMessageContentsRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

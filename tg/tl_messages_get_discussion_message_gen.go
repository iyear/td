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

// MessagesGetDiscussionMessageRequest represents TL type `messages.getDiscussionMessage#446972fd`.
// Get discussion message¹ from the associated discussion group² of a channel to show it on top of the comment section, without actually joining the group
//
// Links:
//  1) https://core.telegram.org/api/threads
//  2) https://core.telegram.org/api/discussion
//
// See https://core.telegram.org/method/messages.getDiscussionMessage for reference.
type MessagesGetDiscussionMessageRequest struct {
	// Channel ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Peer InputPeerClass `tl:"peer"`
	// Message ID
	MsgID int `tl:"msg_id"`
}

// MessagesGetDiscussionMessageRequestTypeID is TL type id of MessagesGetDiscussionMessageRequest.
const MessagesGetDiscussionMessageRequestTypeID = 0x446972fd

func (g *MessagesGetDiscussionMessageRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDiscussionMessageRequest) String() string {
	if g == nil {
		return "MessagesGetDiscussionMessageRequest(nil)"
	}
	type Alias MessagesGetDiscussionMessageRequest
	return fmt.Sprintf("MessagesGetDiscussionMessageRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetDiscussionMessageRequest from given interface.
func (g *MessagesGetDiscussionMessageRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
}) {
	g.Peer = from.GetPeer()
	g.MsgID = from.GetMsgID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetDiscussionMessageRequest) TypeID() uint32 {
	return MessagesGetDiscussionMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *MessagesGetDiscussionMessageRequest) TypeName() string {
	return "messages.getDiscussionMessage"
}

// Encode implements bin.Encoder.
func (g *MessagesGetDiscussionMessageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDiscussionMessage#446972fd as nil")
	}
	b.PutID(MessagesGetDiscussionMessageRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getDiscussionMessage#446972fd: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDiscussionMessage#446972fd: field peer: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetDiscussionMessageRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetMsgID returns value of MsgID field.
func (g *MessagesGetDiscussionMessageRequest) GetMsgID() (value int) {
	return g.MsgID
}

// Decode implements bin.Decoder.
func (g *MessagesGetDiscussionMessageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDiscussionMessage#446972fd to nil")
	}
	if err := b.ConsumeID(MessagesGetDiscussionMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDiscussionMessage#446972fd: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDiscussionMessageRequest.
var (
	_ bin.Encoder = &MessagesGetDiscussionMessageRequest{}
	_ bin.Decoder = &MessagesGetDiscussionMessageRequest{}
)

// MessagesGetDiscussionMessage invokes method messages.getDiscussionMessage#446972fd returning error if any.
// Get discussion message¹ from the associated discussion group² of a channel to show it on top of the comment section, without actually joining the group
//
// Links:
//  1) https://core.telegram.org/api/threads
//  2) https://core.telegram.org/api/discussion
//
// See https://core.telegram.org/method/messages.getDiscussionMessage for reference.
// Can be used by bots.
func (c *Client) MessagesGetDiscussionMessage(ctx context.Context, request *MessagesGetDiscussionMessageRequest) (*MessagesDiscussionMessage, error) {
	var result MessagesDiscussionMessage

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

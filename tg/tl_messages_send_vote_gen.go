// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessagesSendVoteRequest represents TL type `messages.sendVote#10ea6184`.
// Vote in a poll¹
//
// Links:
//  1) https://core.telegram.org/constructor/poll
//
// See https://core.telegram.org/method/messages.sendVote for reference.
type MessagesSendVoteRequest struct {
	// The chat where the poll was sent
	Peer InputPeerClass
	// The message ID of the poll
	MsgID int
	// The options that were chosen
	Options [][]byte
}

// MessagesSendVoteRequestTypeID is TL type id of MessagesSendVoteRequest.
const MessagesSendVoteRequestTypeID = 0x10ea6184

// Ensuring interfaces in compile-time for MessagesSendVoteRequest.
var (
	_ bin.Encoder     = &MessagesSendVoteRequest{}
	_ bin.Decoder     = &MessagesSendVoteRequest{}
	_ bin.BareEncoder = &MessagesSendVoteRequest{}
	_ bin.BareDecoder = &MessagesSendVoteRequest{}
)

func (s *MessagesSendVoteRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.Options == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendVoteRequest) String() string {
	if s == nil {
		return "MessagesSendVoteRequest(nil)"
	}
	type Alias MessagesSendVoteRequest
	return fmt.Sprintf("MessagesSendVoteRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSendVoteRequest from given interface.
func (s *MessagesSendVoteRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetOptions() (value [][]byte)
}) {
	s.Peer = from.GetPeer()
	s.MsgID = from.GetMsgID()
	s.Options = from.GetOptions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSendVoteRequest) TypeID() uint32 {
	return MessagesSendVoteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSendVoteRequest) TypeName() string {
	return "messages.sendVote"
}

// TypeInfo returns info about TL type.
func (s *MessagesSendVoteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sendVote",
		ID:   MessagesSendVoteRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSendVoteRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendVote#10ea6184 as nil")
	}
	b.PutID(MessagesSendVoteRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSendVoteRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendVote#10ea6184 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendVote#10ea6184: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendVote#10ea6184: field peer: %w", err)
	}
	b.PutInt(s.MsgID)
	b.PutVectorHeader(len(s.Options))
	for _, v := range s.Options {
		b.PutBytes(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSendVoteRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendVote#10ea6184 to nil")
	}
	if err := b.ConsumeID(MessagesSendVoteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendVote#10ea6184: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSendVoteRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendVote#10ea6184 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field options: %w", err)
		}

		if headerLen > 0 {
			s.Options = make([][]byte, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode messages.sendVote#10ea6184: field options: %w", err)
			}
			s.Options = append(s.Options, value)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendVoteRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetMsgID returns value of MsgID field.
func (s *MessagesSendVoteRequest) GetMsgID() (value int) {
	if s == nil {
		return
	}
	return s.MsgID
}

// GetOptions returns value of Options field.
func (s *MessagesSendVoteRequest) GetOptions() (value [][]byte) {
	if s == nil {
		return
	}
	return s.Options
}

// MessagesSendVote invokes method messages.sendVote#10ea6184 returning error if any.
// Vote in a poll¹
//
// Links:
//  1) https://core.telegram.org/constructor/poll
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 MESSAGE_ID_INVALID: The provided message id is invalid.
//  400 MESSAGE_POLL_CLOSED: Poll closed.
//  400 OPTIONS_TOO_MUCH: Too many options provided.
//  400 OPTION_INVALID: Invalid option selected.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//  400 REVOTE_NOT_ALLOWED: You cannot change your vote.
//
// See https://core.telegram.org/method/messages.sendVote for reference.
func (c *Client) MessagesSendVote(ctx context.Context, request *MessagesSendVoteRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

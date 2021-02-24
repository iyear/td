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

// ChannelsLeaveChannelRequest represents TL type `channels.leaveChannel#f836aa95`.
// Leave a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.leaveChannel for reference.
type ChannelsLeaveChannelRequest struct {
	// Channel/supergroup¹ to leave
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass `tl:"channel"`
}

// ChannelsLeaveChannelRequestTypeID is TL type id of ChannelsLeaveChannelRequest.
const ChannelsLeaveChannelRequestTypeID = 0xf836aa95

func (l *ChannelsLeaveChannelRequest) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *ChannelsLeaveChannelRequest) String() string {
	if l == nil {
		return "ChannelsLeaveChannelRequest(nil)"
	}
	type Alias ChannelsLeaveChannelRequest
	return fmt.Sprintf("ChannelsLeaveChannelRequest%+v", Alias(*l))
}

// FillFrom fills ChannelsLeaveChannelRequest from given interface.
func (l *ChannelsLeaveChannelRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	l.Channel = from.GetChannel()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *ChannelsLeaveChannelRequest) TypeID() uint32 {
	return ChannelsLeaveChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (l *ChannelsLeaveChannelRequest) TypeName() string {
	return "channels.leaveChannel"
}

// Encode implements bin.Encoder.
func (l *ChannelsLeaveChannelRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode channels.leaveChannel#f836aa95 as nil")
	}
	b.PutID(ChannelsLeaveChannelRequestTypeID)
	if l.Channel == nil {
		return fmt.Errorf("unable to encode channels.leaveChannel#f836aa95: field channel is nil")
	}
	if err := l.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.leaveChannel#f836aa95: field channel: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (l *ChannelsLeaveChannelRequest) GetChannel() (value InputChannelClass) {
	return l.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (l *ChannelsLeaveChannelRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return l.Channel.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (l *ChannelsLeaveChannelRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode channels.leaveChannel#f836aa95 to nil")
	}
	if err := b.ConsumeID(ChannelsLeaveChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.leaveChannel#f836aa95: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.leaveChannel#f836aa95: field channel: %w", err)
		}
		l.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsLeaveChannelRequest.
var (
	_ bin.Encoder = &ChannelsLeaveChannelRequest{}
	_ bin.Decoder = &ChannelsLeaveChannelRequest{}
)

// ChannelsLeaveChannel invokes method channels.leaveChannel#f836aa95 returning error if any.
// Leave a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  403 CHANNEL_PUBLIC_GROUP_NA: channel/supergroup not available
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 USER_CREATOR: You can't leave this channel, because you're its creator
//  400 USER_NOT_PARTICIPANT: You're not a member of this supergroup/channel
//
// See https://core.telegram.org/method/channels.leaveChannel for reference.
// Can be used by bots.
func (c *Client) ChannelsLeaveChannel(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsLeaveChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

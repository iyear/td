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

// ChannelsTogglePreHistoryHiddenRequest represents TL type `channels.togglePreHistoryHidden#eabbb94c`.
// Hide/unhide message history for new channel/supergroup users
//
// See https://core.telegram.org/method/channels.togglePreHistoryHidden for reference.
type ChannelsTogglePreHistoryHiddenRequest struct {
	// Channel/supergroup
	Channel InputChannelClass `tl:"channel"`
	// Hide/unhide
	Enabled bool `tl:"enabled"`
}

// ChannelsTogglePreHistoryHiddenRequestTypeID is TL type id of ChannelsTogglePreHistoryHiddenRequest.
const ChannelsTogglePreHistoryHiddenRequestTypeID = 0xeabbb94c

func (t *ChannelsTogglePreHistoryHiddenRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsTogglePreHistoryHiddenRequest) String() string {
	if t == nil {
		return "ChannelsTogglePreHistoryHiddenRequest(nil)"
	}
	type Alias ChannelsTogglePreHistoryHiddenRequest
	return fmt.Sprintf("ChannelsTogglePreHistoryHiddenRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsTogglePreHistoryHiddenRequest from given interface.
func (t *ChannelsTogglePreHistoryHiddenRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEnabled() (value bool)
}) {
	t.Channel = from.GetChannel()
	t.Enabled = from.GetEnabled()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *ChannelsTogglePreHistoryHiddenRequest) TypeID() uint32 {
	return ChannelsTogglePreHistoryHiddenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (t *ChannelsTogglePreHistoryHiddenRequest) TypeName() string {
	return "channels.togglePreHistoryHidden"
}

// Encode implements bin.Encoder.
func (t *ChannelsTogglePreHistoryHiddenRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.togglePreHistoryHidden#eabbb94c as nil")
	}
	b.PutID(ChannelsTogglePreHistoryHiddenRequestTypeID)
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.togglePreHistoryHidden#eabbb94c: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.togglePreHistoryHidden#eabbb94c: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsTogglePreHistoryHiddenRequest) GetChannel() (value InputChannelClass) {
	return t.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsTogglePreHistoryHiddenRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsTogglePreHistoryHiddenRequest) GetEnabled() (value bool) {
	return t.Enabled
}

// Decode implements bin.Decoder.
func (t *ChannelsTogglePreHistoryHiddenRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.togglePreHistoryHidden#eabbb94c to nil")
	}
	if err := b.ConsumeID(ChannelsTogglePreHistoryHiddenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.togglePreHistoryHidden#eabbb94c: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.togglePreHistoryHidden#eabbb94c: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.togglePreHistoryHidden#eabbb94c: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsTogglePreHistoryHiddenRequest.
var (
	_ bin.Encoder = &ChannelsTogglePreHistoryHiddenRequest{}
	_ bin.Decoder = &ChannelsTogglePreHistoryHiddenRequest{}
)

// ChannelsTogglePreHistoryHidden invokes method channels.togglePreHistoryHidden#eabbb94c returning error if any.
// Hide/unhide message history for new channel/supergroup users
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 CHAT_LINK_EXISTS: The chat is public, you can't hide the history to new users
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//
// See https://core.telegram.org/method/channels.togglePreHistoryHidden for reference.
func (c *Client) ChannelsTogglePreHistoryHidden(ctx context.Context, request *ChannelsTogglePreHistoryHiddenRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

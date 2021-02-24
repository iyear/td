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

// ChannelsToggleSlowModeRequest represents TL type `channels.toggleSlowMode#edd49ef0`.
// Toggle supergroup slow mode: if enabled, users will only be able to send one message every seconds seconds
//
// See https://core.telegram.org/method/channels.toggleSlowMode for reference.
type ChannelsToggleSlowModeRequest struct {
	// The supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass `tl:"channel"`
	// Users will only be able to send one message every seconds seconds, 0 to disable the limitation
	Seconds int `tl:"seconds"`
}

// ChannelsToggleSlowModeRequestTypeID is TL type id of ChannelsToggleSlowModeRequest.
const ChannelsToggleSlowModeRequestTypeID = 0xedd49ef0

func (t *ChannelsToggleSlowModeRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Seconds == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsToggleSlowModeRequest) String() string {
	if t == nil {
		return "ChannelsToggleSlowModeRequest(nil)"
	}
	type Alias ChannelsToggleSlowModeRequest
	return fmt.Sprintf("ChannelsToggleSlowModeRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsToggleSlowModeRequest from given interface.
func (t *ChannelsToggleSlowModeRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetSeconds() (value int)
}) {
	t.Channel = from.GetChannel()
	t.Seconds = from.GetSeconds()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *ChannelsToggleSlowModeRequest) TypeID() uint32 {
	return ChannelsToggleSlowModeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (t *ChannelsToggleSlowModeRequest) TypeName() string {
	return "channels.toggleSlowMode"
}

// Encode implements bin.Encoder.
func (t *ChannelsToggleSlowModeRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.toggleSlowMode#edd49ef0 as nil")
	}
	b.PutID(ChannelsToggleSlowModeRequestTypeID)
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.toggleSlowMode#edd49ef0: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.toggleSlowMode#edd49ef0: field channel: %w", err)
	}
	b.PutInt(t.Seconds)
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsToggleSlowModeRequest) GetChannel() (value InputChannelClass) {
	return t.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsToggleSlowModeRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// GetSeconds returns value of Seconds field.
func (t *ChannelsToggleSlowModeRequest) GetSeconds() (value int) {
	return t.Seconds
}

// Decode implements bin.Decoder.
func (t *ChannelsToggleSlowModeRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.toggleSlowMode#edd49ef0 to nil")
	}
	if err := b.ConsumeID(ChannelsToggleSlowModeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.toggleSlowMode#edd49ef0: field seconds: %w", err)
		}
		t.Seconds = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsToggleSlowModeRequest.
var (
	_ bin.Encoder = &ChannelsToggleSlowModeRequest{}
	_ bin.Decoder = &ChannelsToggleSlowModeRequest{}
)

// ChannelsToggleSlowMode invokes method channels.toggleSlowMode#edd49ef0 returning error if any.
// Toggle supergroup slow mode: if enabled, users will only be able to send one message every seconds seconds
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  400 INPUT_METHOD_INVALID_1192227_X: Invalid method
//  400 INPUT_METHOD_INVALID_1604042050_X: Invalid method
//  400 SECONDS_INVALID: Invalid duration provided
//
// See https://core.telegram.org/method/channels.toggleSlowMode for reference.
func (c *Client) ChannelsToggleSlowMode(ctx context.Context, request *ChannelsToggleSlowModeRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

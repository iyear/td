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

// ChannelsConvertToGigagroupRequest represents TL type `channels.convertToGigagroup#b290c69`.
//
// See https://core.telegram.org/method/channels.convertToGigagroup for reference.
type ChannelsConvertToGigagroupRequest struct {
	// Channel field of ChannelsConvertToGigagroupRequest.
	Channel InputChannelClass `tl:"channel"`
}

// ChannelsConvertToGigagroupRequestTypeID is TL type id of ChannelsConvertToGigagroupRequest.
const ChannelsConvertToGigagroupRequestTypeID = 0xb290c69

func (c *ChannelsConvertToGigagroupRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsConvertToGigagroupRequest) String() string {
	if c == nil {
		return "ChannelsConvertToGigagroupRequest(nil)"
	}
	type Alias ChannelsConvertToGigagroupRequest
	return fmt.Sprintf("ChannelsConvertToGigagroupRequest%+v", Alias(*c))
}

// FillFrom fills ChannelsConvertToGigagroupRequest from given interface.
func (c *ChannelsConvertToGigagroupRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	c.Channel = from.GetChannel()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelsConvertToGigagroupRequest) TypeID() uint32 {
	return ChannelsConvertToGigagroupRequestTypeID
}

// TypeName returns name of type in TL schema.
func (c *ChannelsConvertToGigagroupRequest) TypeName() string {
	return "channels.convertToGigagroup"
}

// Encode implements bin.Encoder.
func (c *ChannelsConvertToGigagroupRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.convertToGigagroup#b290c69 as nil")
	}
	b.PutID(ChannelsConvertToGigagroupRequestTypeID)
	if c.Channel == nil {
		return fmt.Errorf("unable to encode channels.convertToGigagroup#b290c69: field channel is nil")
	}
	if err := c.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.convertToGigagroup#b290c69: field channel: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (c *ChannelsConvertToGigagroupRequest) GetChannel() (value InputChannelClass) {
	return c.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (c *ChannelsConvertToGigagroupRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return c.Channel.AsNotEmpty()
}

// Decode implements bin.Decoder.
func (c *ChannelsConvertToGigagroupRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.convertToGigagroup#b290c69 to nil")
	}
	if err := b.ConsumeID(ChannelsConvertToGigagroupRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.convertToGigagroup#b290c69: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.convertToGigagroup#b290c69: field channel: %w", err)
		}
		c.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsConvertToGigagroupRequest.
var (
	_ bin.Encoder = &ChannelsConvertToGigagroupRequest{}
	_ bin.Decoder = &ChannelsConvertToGigagroupRequest{}
)

// ChannelsConvertToGigagroup invokes method channels.convertToGigagroup#b290c69 returning error if any.
//
// See https://core.telegram.org/method/channels.convertToGigagroup for reference.
func (c *Client) ChannelsConvertToGigagroup(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsConvertToGigagroupRequest{
		Channel: channel,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

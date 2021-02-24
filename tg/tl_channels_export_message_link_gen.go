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

// ChannelsExportMessageLinkRequest represents TL type `channels.exportMessageLink#e63fadeb`.
// Get link and embed info of a message in a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.exportMessageLink for reference.
type ChannelsExportMessageLinkRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether to include other grouped media (for albums)
	Grouped bool `tl:"grouped"`
	// Whether to also include a thread ID, if available, inside of the link
	Thread bool `tl:"thread"`
	// Channel
	Channel InputChannelClass `tl:"channel"`
	// Message ID
	ID int `tl:"id"`
}

// ChannelsExportMessageLinkRequestTypeID is TL type id of ChannelsExportMessageLinkRequest.
const ChannelsExportMessageLinkRequestTypeID = 0xe63fadeb

func (e *ChannelsExportMessageLinkRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Grouped == false) {
		return false
	}
	if !(e.Thread == false) {
		return false
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsExportMessageLinkRequest) String() string {
	if e == nil {
		return "ChannelsExportMessageLinkRequest(nil)"
	}
	type Alias ChannelsExportMessageLinkRequest
	return fmt.Sprintf("ChannelsExportMessageLinkRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsExportMessageLinkRequest from given interface.
func (e *ChannelsExportMessageLinkRequest) FillFrom(from interface {
	GetGrouped() (value bool)
	GetThread() (value bool)
	GetChannel() (value InputChannelClass)
	GetID() (value int)
}) {
	e.Grouped = from.GetGrouped()
	e.Thread = from.GetThread()
	e.Channel = from.GetChannel()
	e.ID = from.GetID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *ChannelsExportMessageLinkRequest) TypeID() uint32 {
	return ChannelsExportMessageLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (e *ChannelsExportMessageLinkRequest) TypeName() string {
	return "channels.exportMessageLink"
}

// Encode implements bin.Encoder.
func (e *ChannelsExportMessageLinkRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.exportMessageLink#e63fadeb as nil")
	}
	b.PutID(ChannelsExportMessageLinkRequestTypeID)
	if !(e.Grouped == false) {
		e.Flags.Set(0)
	}
	if !(e.Thread == false) {
		e.Flags.Set(1)
	}
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field flags: %w", err)
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field channel: %w", err)
	}
	b.PutInt(e.ID)
	return nil
}

// SetGrouped sets value of Grouped conditional field.
func (e *ChannelsExportMessageLinkRequest) SetGrouped(value bool) {
	if value {
		e.Flags.Set(0)
		e.Grouped = true
	} else {
		e.Flags.Unset(0)
		e.Grouped = false
	}
}

// GetGrouped returns value of Grouped conditional field.
func (e *ChannelsExportMessageLinkRequest) GetGrouped() (value bool) {
	return e.Flags.Has(0)
}

// SetThread sets value of Thread conditional field.
func (e *ChannelsExportMessageLinkRequest) SetThread(value bool) {
	if value {
		e.Flags.Set(1)
		e.Thread = true
	} else {
		e.Flags.Unset(1)
		e.Thread = false
	}
}

// GetThread returns value of Thread conditional field.
func (e *ChannelsExportMessageLinkRequest) GetThread() (value bool) {
	return e.Flags.Has(1)
}

// GetChannel returns value of Channel field.
func (e *ChannelsExportMessageLinkRequest) GetChannel() (value InputChannelClass) {
	return e.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsExportMessageLinkRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// GetID returns value of ID field.
func (e *ChannelsExportMessageLinkRequest) GetID() (value int) {
	return e.ID
}

// Decode implements bin.Decoder.
func (e *ChannelsExportMessageLinkRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.exportMessageLink#e63fadeb to nil")
	}
	if err := b.ConsumeID(ChannelsExportMessageLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field flags: %w", err)
		}
	}
	e.Grouped = e.Flags.Has(0)
	e.Thread = e.Flags.Has(1)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsExportMessageLinkRequest.
var (
	_ bin.Encoder = &ChannelsExportMessageLinkRequest{}
	_ bin.Decoder = &ChannelsExportMessageLinkRequest{}
)

// ChannelsExportMessageLink invokes method channels.exportMessageLink#e63fadeb returning error if any.
// Get link and embed info of a message in a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.exportMessageLink for reference.
func (c *Client) ChannelsExportMessageLink(ctx context.Context, request *ChannelsExportMessageLinkRequest) (*ExportedMessageLink, error) {
	var result ExportedMessageLink

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

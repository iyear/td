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

// MessagesAffectedMessages represents TL type `messages.affectedMessages#84d19185`.
// Events affected by operation
//
// See https://core.telegram.org/constructor/messages.affectedMessages for reference.
type MessagesAffectedMessages struct {
	// Event count after generation¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// Number of events that were generated¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	PtsCount int
}

// MessagesAffectedMessagesTypeID is TL type id of MessagesAffectedMessages.
const MessagesAffectedMessagesTypeID = 0x84d19185

func (a *MessagesAffectedMessages) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Pts == 0) {
		return false
	}
	if !(a.PtsCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAffectedMessages) String() string {
	if a == nil {
		return "MessagesAffectedMessages(nil)"
	}
	type Alias MessagesAffectedMessages
	return fmt.Sprintf("MessagesAffectedMessages%+v", Alias(*a))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (a *MessagesAffectedMessages) TypeID() uint32 {
	return MessagesAffectedMessagesTypeID
}

// Encode implements bin.Encoder.
func (a *MessagesAffectedMessages) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedMessages#84d19185 as nil")
	}
	b.PutID(MessagesAffectedMessagesTypeID)
	b.PutInt(a.Pts)
	b.PutInt(a.PtsCount)
	return nil
}

// GetPts returns value of Pts field.
func (a *MessagesAffectedMessages) GetPts() (value int) {
	return a.Pts
}

// GetPtsCount returns value of PtsCount field.
func (a *MessagesAffectedMessages) GetPtsCount() (value int) {
	return a.PtsCount
}

// Decode implements bin.Decoder.
func (a *MessagesAffectedMessages) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedMessages#84d19185 to nil")
	}
	if err := b.ConsumeID(MessagesAffectedMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.affectedMessages#84d19185: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedMessages#84d19185: field pts: %w", err)
		}
		a.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedMessages#84d19185: field pts_count: %w", err)
		}
		a.PtsCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAffectedMessages.
var (
	_ bin.Encoder = &MessagesAffectedMessages{}
	_ bin.Decoder = &MessagesAffectedMessages{}
)

// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// Long represents TL type `long#22076cba`.
//
// See https://core.telegram.org/constructor/long for reference.
type Long struct {
}

// LongTypeID is TL type id of Long.
const LongTypeID = 0x22076cba

func (l *Long) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *Long) String() string {
	if l == nil {
		return "Long(nil)"
	}
	type Alias Long
	return fmt.Sprintf("Long%+v", Alias(*l))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (l *Long) TypeID() uint32 {
	return LongTypeID
}

// Encode implements bin.Encoder.
func (l *Long) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode long#22076cba as nil")
	}
	b.PutID(LongTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (l *Long) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode long#22076cba to nil")
	}
	if err := b.ConsumeID(LongTypeID); err != nil {
		return fmt.Errorf("unable to decode long#22076cba: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Long.
var (
	_ bin.Encoder = &Long{}
	_ bin.Decoder = &Long{}
)

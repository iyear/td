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

// MessagesStickerSetInstallResultSuccess represents TL type `messages.stickerSetInstallResultSuccess#38641628`.
// The stickerset was installed successfully
//
// See https://core.telegram.org/constructor/messages.stickerSetInstallResultSuccess for reference.
type MessagesStickerSetInstallResultSuccess struct {
}

// MessagesStickerSetInstallResultSuccessTypeID is TL type id of MessagesStickerSetInstallResultSuccess.
const MessagesStickerSetInstallResultSuccessTypeID = 0x38641628

func (s *MessagesStickerSetInstallResultSuccess) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickerSetInstallResultSuccess) String() string {
	if s == nil {
		return "MessagesStickerSetInstallResultSuccess(nil)"
	}
	type Alias MessagesStickerSetInstallResultSuccess
	return fmt.Sprintf("MessagesStickerSetInstallResultSuccess%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesStickerSetInstallResultSuccess) TypeID() uint32 {
	return MessagesStickerSetInstallResultSuccessTypeID
}

// TypeName returns name of type in TL schema.
func (s *MessagesStickerSetInstallResultSuccess) TypeName() string {
	return "messages.stickerSetInstallResultSuccess"
}

// Encode implements bin.Encoder.
func (s *MessagesStickerSetInstallResultSuccess) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSetInstallResultSuccess#38641628 as nil")
	}
	b.PutID(MessagesStickerSetInstallResultSuccessTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSetInstallResultSuccess) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSetInstallResultSuccess#38641628 to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetInstallResultSuccessTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSetInstallResultSuccess#38641628: %w", err)
	}
	return nil
}

// construct implements constructor of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultSuccess) construct() MessagesStickerSetInstallResultClass {
	return &s
}

// Ensuring interfaces in compile-time for MessagesStickerSetInstallResultSuccess.
var (
	_ bin.Encoder = &MessagesStickerSetInstallResultSuccess{}
	_ bin.Decoder = &MessagesStickerSetInstallResultSuccess{}

	_ MessagesStickerSetInstallResultClass = &MessagesStickerSetInstallResultSuccess{}
)

// MessagesStickerSetInstallResultArchive represents TL type `messages.stickerSetInstallResultArchive#35e410a8`.
// The stickerset was installed, but since there are too many stickersets some were archived
//
// See https://core.telegram.org/constructor/messages.stickerSetInstallResultArchive for reference.
type MessagesStickerSetInstallResultArchive struct {
	// Archived stickersets
	Sets []StickerSetCoveredClass `tl:"sets"`
}

// MessagesStickerSetInstallResultArchiveTypeID is TL type id of MessagesStickerSetInstallResultArchive.
const MessagesStickerSetInstallResultArchiveTypeID = 0x35e410a8

func (s *MessagesStickerSetInstallResultArchive) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickerSetInstallResultArchive) String() string {
	if s == nil {
		return "MessagesStickerSetInstallResultArchive(nil)"
	}
	type Alias MessagesStickerSetInstallResultArchive
	return fmt.Sprintf("MessagesStickerSetInstallResultArchive%+v", Alias(*s))
}

// FillFrom fills MessagesStickerSetInstallResultArchive from given interface.
func (s *MessagesStickerSetInstallResultArchive) FillFrom(from interface {
	GetSets() (value []StickerSetCoveredClass)
}) {
	s.Sets = from.GetSets()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesStickerSetInstallResultArchive) TypeID() uint32 {
	return MessagesStickerSetInstallResultArchiveTypeID
}

// TypeName returns name of type in TL schema.
func (s *MessagesStickerSetInstallResultArchive) TypeName() string {
	return "messages.stickerSetInstallResultArchive"
}

// Encode implements bin.Encoder.
func (s *MessagesStickerSetInstallResultArchive) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSetInstallResultArchive#35e410a8 as nil")
	}
	b.PutID(MessagesStickerSetInstallResultArchiveTypeID)
	b.PutVectorHeader(len(s.Sets))
	for idx, v := range s.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.stickerSetInstallResultArchive#35e410a8: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSetInstallResultArchive#35e410a8: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetSets returns value of Sets field.
func (s *MessagesStickerSetInstallResultArchive) GetSets() (value []StickerSetCoveredClass) {
	return s.Sets
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassSlice helper.
func (s *MessagesStickerSetInstallResultArchive) MapSets() (value StickerSetCoveredClassSlice) {
	return StickerSetCoveredClassSlice(s.Sets)
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSetInstallResultArchive) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSetInstallResultArchive#35e410a8 to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetInstallResultArchiveTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSetInstallResultArchive#35e410a8: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSetInstallResultArchive#35e410a8: field sets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.stickerSetInstallResultArchive#35e410a8: field sets: %w", err)
			}
			s.Sets = append(s.Sets, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultArchive) construct() MessagesStickerSetInstallResultClass {
	return &s
}

// Ensuring interfaces in compile-time for MessagesStickerSetInstallResultArchive.
var (
	_ bin.Encoder = &MessagesStickerSetInstallResultArchive{}
	_ bin.Decoder = &MessagesStickerSetInstallResultArchive{}

	_ MessagesStickerSetInstallResultClass = &MessagesStickerSetInstallResultArchive{}
)

// MessagesStickerSetInstallResultClass represents messages.StickerSetInstallResult generic type.
//
// See https://core.telegram.org/type/messages.StickerSetInstallResult for reference.
//
// Example:
//  g, err := tg.DecodeMessagesStickerSetInstallResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesStickerSetInstallResultSuccess: // messages.stickerSetInstallResultSuccess#38641628
//  case *tg.MessagesStickerSetInstallResultArchive: // messages.stickerSetInstallResultArchive#35e410a8
//  default: panic(v)
//  }
type MessagesStickerSetInstallResultClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesStickerSetInstallResultClass

	// TypeID returns type id in TL schema.
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessagesStickerSetInstallResult implements binary de-serialization for MessagesStickerSetInstallResultClass.
func DecodeMessagesStickerSetInstallResult(buf *bin.Buffer) (MessagesStickerSetInstallResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesStickerSetInstallResultSuccessTypeID:
		// Decoding messages.stickerSetInstallResultSuccess#38641628.
		v := MessagesStickerSetInstallResultSuccess{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickerSetInstallResultClass: %w", err)
		}
		return &v, nil
	case MessagesStickerSetInstallResultArchiveTypeID:
		// Decoding messages.stickerSetInstallResultArchive#35e410a8.
		v := MessagesStickerSetInstallResultArchive{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesStickerSetInstallResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesStickerSetInstallResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesStickerSetInstallResult boxes the MessagesStickerSetInstallResultClass providing a helper.
type MessagesStickerSetInstallResultBox struct {
	StickerSetInstallResult MessagesStickerSetInstallResultClass
}

// Decode implements bin.Decoder for MessagesStickerSetInstallResultBox.
func (b *MessagesStickerSetInstallResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesStickerSetInstallResultBox to nil")
	}
	v, err := DecodeMessagesStickerSetInstallResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerSetInstallResult = v
	return nil
}

// Encode implements bin.Encode for MessagesStickerSetInstallResultBox.
func (b *MessagesStickerSetInstallResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StickerSetInstallResult == nil {
		return fmt.Errorf("unable to encode MessagesStickerSetInstallResultClass as nil")
	}
	return b.StickerSetInstallResult.Encode(buf)
}

// MessagesStickerSetInstallResultClassSlice is adapter for slice of MessagesStickerSetInstallResultClass.
type MessagesStickerSetInstallResultClassSlice []MessagesStickerSetInstallResultClass

// First returns first element of slice (if exists).
func (s MessagesStickerSetInstallResultClassSlice) First() (v MessagesStickerSetInstallResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesStickerSetInstallResultClassSlice) Last() (v MessagesStickerSetInstallResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesStickerSetInstallResultClassSlice) PopFirst() (v MessagesStickerSetInstallResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesStickerSetInstallResultClassSlice) Pop() (v MessagesStickerSetInstallResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

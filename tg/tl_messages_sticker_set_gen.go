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

// MessagesStickerSet represents TL type `messages.stickerSet#b60a24a6`.
// Stickerset and stickers inside it
//
// See https://core.telegram.org/constructor/messages.stickerSet for reference.
type MessagesStickerSet struct {
	// The stickerset
	Set StickerSet `tl:"set"`
	// Emoji info for stickers
	Packs []StickerPack `tl:"packs"`
	// Stickers in stickerset
	Documents []DocumentClass `tl:"documents"`
}

// MessagesStickerSetTypeID is TL type id of MessagesStickerSet.
const MessagesStickerSetTypeID = 0xb60a24a6

func (s *MessagesStickerSet) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Set.Zero()) {
		return false
	}
	if !(s.Packs == nil) {
		return false
	}
	if !(s.Documents == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesStickerSet) String() string {
	if s == nil {
		return "MessagesStickerSet(nil)"
	}
	type Alias MessagesStickerSet
	return fmt.Sprintf("MessagesStickerSet%+v", Alias(*s))
}

// FillFrom fills MessagesStickerSet from given interface.
func (s *MessagesStickerSet) FillFrom(from interface {
	GetSet() (value StickerSet)
	GetPacks() (value []StickerPack)
	GetDocuments() (value []DocumentClass)
}) {
	s.Set = from.GetSet()
	s.Packs = from.GetPacks()
	s.Documents = from.GetDocuments()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesStickerSet) TypeID() uint32 {
	return MessagesStickerSetTypeID
}

// TypeName returns name of type in TL schema.
func (s *MessagesStickerSet) TypeName() string {
	return "messages.stickerSet"
}

// Encode implements bin.Encoder.
func (s *MessagesStickerSet) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.stickerSet#b60a24a6 as nil")
	}
	b.PutID(MessagesStickerSetTypeID)
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.stickerSet#b60a24a6: field set: %w", err)
	}
	b.PutVectorHeader(len(s.Packs))
	for idx, v := range s.Packs {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSet#b60a24a6: field packs element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Documents))
	for idx, v := range s.Documents {
		if v == nil {
			return fmt.Errorf("unable to encode messages.stickerSet#b60a24a6: field documents element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.stickerSet#b60a24a6: field documents element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetSet returns value of Set field.
func (s *MessagesStickerSet) GetSet() (value StickerSet) {
	return s.Set
}

// GetPacks returns value of Packs field.
func (s *MessagesStickerSet) GetPacks() (value []StickerPack) {
	return s.Packs
}

// GetDocuments returns value of Documents field.
func (s *MessagesStickerSet) GetDocuments() (value []DocumentClass) {
	return s.Documents
}

// MapDocuments returns field Documents wrapped in DocumentClassSlice helper.
func (s *MessagesStickerSet) MapDocuments() (value DocumentClassSlice) {
	return DocumentClassSlice(s.Documents)
}

// Decode implements bin.Decoder.
func (s *MessagesStickerSet) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.stickerSet#b60a24a6 to nil")
	}
	if err := b.ConsumeID(MessagesStickerSetTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.stickerSet#b60a24a6: %w", err)
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#b60a24a6: field set: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#b60a24a6: field packs: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerPack
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.stickerSet#b60a24a6: field packs: %w", err)
			}
			s.Packs = append(s.Packs, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.stickerSet#b60a24a6: field documents: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.stickerSet#b60a24a6: field documents: %w", err)
			}
			s.Documents = append(s.Documents, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesStickerSet.
var (
	_ bin.Encoder = &MessagesStickerSet{}
	_ bin.Decoder = &MessagesStickerSet{}
)

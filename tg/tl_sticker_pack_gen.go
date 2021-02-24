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

// StickerPack represents TL type `stickerPack#12b299d4`.
// A stickerpack is a group of stickers associated to the same emoji.
// It is not a sticker pack the way it is usually intended, you may be looking for a StickerSet¹.
//
// Links:
//  1) https://core.telegram.org/type/StickerSet
//
// See https://core.telegram.org/constructor/stickerPack for reference.
type StickerPack struct {
	// Emoji
	Emoticon string `tl:"emoticon"`
	// Stickers
	Documents []int64 `tl:"documents"`
}

// StickerPackTypeID is TL type id of StickerPack.
const StickerPackTypeID = 0x12b299d4

func (s *StickerPack) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Emoticon == "") {
		return false
	}
	if !(s.Documents == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerPack) String() string {
	if s == nil {
		return "StickerPack(nil)"
	}
	type Alias StickerPack
	return fmt.Sprintf("StickerPack%+v", Alias(*s))
}

// FillFrom fills StickerPack from given interface.
func (s *StickerPack) FillFrom(from interface {
	GetEmoticon() (value string)
	GetDocuments() (value []int64)
}) {
	s.Emoticon = from.GetEmoticon()
	s.Documents = from.GetDocuments()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *StickerPack) TypeID() uint32 {
	return StickerPackTypeID
}

// TypeName returns name of type in TL schema.
func (s *StickerPack) TypeName() string {
	return "stickerPack"
}

// Encode implements bin.Encoder.
func (s *StickerPack) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerPack#12b299d4 as nil")
	}
	b.PutID(StickerPackTypeID)
	b.PutString(s.Emoticon)
	b.PutVectorHeader(len(s.Documents))
	for _, v := range s.Documents {
		b.PutLong(v)
	}
	return nil
}

// GetEmoticon returns value of Emoticon field.
func (s *StickerPack) GetEmoticon() (value string) {
	return s.Emoticon
}

// GetDocuments returns value of Documents field.
func (s *StickerPack) GetDocuments() (value []int64) {
	return s.Documents
}

// Decode implements bin.Decoder.
func (s *StickerPack) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerPack#12b299d4 to nil")
	}
	if err := b.ConsumeID(StickerPackTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerPack#12b299d4: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerPack#12b299d4: field emoticon: %w", err)
		}
		s.Emoticon = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickerPack#12b299d4: field documents: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode stickerPack#12b299d4: field documents: %w", err)
			}
			s.Documents = append(s.Documents, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StickerPack.
var (
	_ bin.Encoder = &StickerPack{}
	_ bin.Decoder = &StickerPack{}
)

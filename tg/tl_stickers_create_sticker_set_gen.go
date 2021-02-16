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

// StickersCreateStickerSetRequest represents TL type `stickers.createStickerSet#f1036780`.
// Create a stickerset, bots only.
//
// See https://core.telegram.org/method/stickers.createStickerSet for reference.
type StickersCreateStickerSetRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a mask stickerset
	Masks bool
	// Whether this is an animated stickerset
	Animated bool
	// Stickerset owner
	UserID InputUserClass
	// Stickerset name, 1-64 chars
	Title string
	// Sticker set name. Can contain only English letters, digits and underscores. Must end with "by" ( is case insensitive); 1-64 characters
	ShortName string
	// Thumbnail
	//
	// Use SetThumb and GetThumb helpers.
	Thumb InputDocumentClass
	// Stickers
	Stickers []InputStickerSetItem
}

// StickersCreateStickerSetRequestTypeID is TL type id of StickersCreateStickerSetRequest.
const StickersCreateStickerSetRequestTypeID = 0xf1036780

func (c *StickersCreateStickerSetRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Masks == false) {
		return false
	}
	if !(c.Animated == false) {
		return false
	}
	if !(c.UserID == nil) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.ShortName == "") {
		return false
	}
	if !(c.Thumb == nil) {
		return false
	}
	if !(c.Stickers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *StickersCreateStickerSetRequest) String() string {
	if c == nil {
		return "StickersCreateStickerSetRequest(nil)"
	}
	type Alias StickersCreateStickerSetRequest
	return fmt.Sprintf("StickersCreateStickerSetRequest%+v", Alias(*c))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *StickersCreateStickerSetRequest) TypeID() uint32 {
	return StickersCreateStickerSetRequestTypeID
}

// Encode implements bin.Encoder.
func (c *StickersCreateStickerSetRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.createStickerSet#f1036780 as nil")
	}
	b.PutID(StickersCreateStickerSetRequestTypeID)
	if !(c.Masks == false) {
		c.Flags.Set(0)
	}
	if !(c.Animated == false) {
		c.Flags.Set(1)
	}
	if !(c.Thumb == nil) {
		c.Flags.Set(2)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.createStickerSet#f1036780: field flags: %w", err)
	}
	if c.UserID == nil {
		return fmt.Errorf("unable to encode stickers.createStickerSet#f1036780: field user_id is nil")
	}
	if err := c.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.createStickerSet#f1036780: field user_id: %w", err)
	}
	b.PutString(c.Title)
	b.PutString(c.ShortName)
	if c.Flags.Has(2) {
		if c.Thumb == nil {
			return fmt.Errorf("unable to encode stickers.createStickerSet#f1036780: field thumb is nil")
		}
		if err := c.Thumb.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickers.createStickerSet#f1036780: field thumb: %w", err)
		}
	}
	b.PutVectorHeader(len(c.Stickers))
	for idx, v := range c.Stickers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickers.createStickerSet#f1036780: field stickers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetMasks sets value of Masks conditional field.
func (c *StickersCreateStickerSetRequest) SetMasks(value bool) {
	if value {
		c.Flags.Set(0)
		c.Masks = true
	} else {
		c.Flags.Unset(0)
		c.Masks = false
	}
}

// GetMasks returns value of Masks conditional field.
func (c *StickersCreateStickerSetRequest) GetMasks() (value bool) {
	return c.Flags.Has(0)
}

// SetAnimated sets value of Animated conditional field.
func (c *StickersCreateStickerSetRequest) SetAnimated(value bool) {
	if value {
		c.Flags.Set(1)
		c.Animated = true
	} else {
		c.Flags.Unset(1)
		c.Animated = false
	}
}

// GetAnimated returns value of Animated conditional field.
func (c *StickersCreateStickerSetRequest) GetAnimated() (value bool) {
	return c.Flags.Has(1)
}

// GetUserID returns value of UserID field.
func (c *StickersCreateStickerSetRequest) GetUserID() (value InputUserClass) {
	return c.UserID
}

// GetTitle returns value of Title field.
func (c *StickersCreateStickerSetRequest) GetTitle() (value string) {
	return c.Title
}

// GetShortName returns value of ShortName field.
func (c *StickersCreateStickerSetRequest) GetShortName() (value string) {
	return c.ShortName
}

// SetThumb sets value of Thumb conditional field.
func (c *StickersCreateStickerSetRequest) SetThumb(value InputDocumentClass) {
	c.Flags.Set(2)
	c.Thumb = value
}

// GetThumb returns value of Thumb conditional field and
// boolean which is true if field was set.
func (c *StickersCreateStickerSetRequest) GetThumb() (value InputDocumentClass, ok bool) {
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.Thumb, true
}

// GetStickers returns value of Stickers field.
func (c *StickersCreateStickerSetRequest) GetStickers() (value []InputStickerSetItem) {
	return c.Stickers
}

// Decode implements bin.Decoder.
func (c *StickersCreateStickerSetRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.createStickerSet#f1036780 to nil")
	}
	if err := b.ConsumeID(StickersCreateStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field flags: %w", err)
		}
	}
	c.Masks = c.Flags.Has(0)
	c.Animated = c.Flags.Has(1)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field title: %w", err)
		}
		c.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field short_name: %w", err)
		}
		c.ShortName = value
	}
	if c.Flags.Has(2) {
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field thumb: %w", err)
		}
		c.Thumb = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field stickers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value InputStickerSetItem
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode stickers.createStickerSet#f1036780: field stickers: %w", err)
			}
			c.Stickers = append(c.Stickers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersCreateStickerSetRequest.
var (
	_ bin.Encoder = &StickersCreateStickerSetRequest{}
	_ bin.Decoder = &StickersCreateStickerSetRequest{}
)

// StickersCreateStickerSet invokes method stickers.createStickerSet#f1036780 returning error if any.
// Create a stickerset, bots only.
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot
//  400 PACK_SHORT_NAME_INVALID: Short pack name invalid
//  400 PACK_SHORT_NAME_OCCUPIED: A stickerpack with this name already exists
//  400 PACK_TITLE_INVALID: The stickerpack title is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 SHORTNAME_OCCUPY_FAILED: An internal error occurred
//  400 STICKERS_EMPTY: No sticker provided
//  400 STICKER_EMOJI_INVALID: Sticker emoji invalid
//  400 STICKER_FILE_INVALID: Sticker file invalid
//  400 STICKER_PNG_DIMENSIONS: Sticker png dimensions invalid
//  400 STICKER_PNG_NOPNG: One of the specified stickers is not a valid PNG file
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/stickers.createStickerSet for reference.
// Can be used by bots.
func (c *Client) StickersCreateStickerSet(ctx context.Context, request *StickersCreateStickerSetRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

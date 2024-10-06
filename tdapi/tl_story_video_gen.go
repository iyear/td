// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// StoryVideo represents TL type `storyVideo#562b0a45`.
type StoryVideo struct {
	// Duration of the video, in seconds
	Duration float64
	// Video width
	Width int32
	// Video height
	Height int32
	// True, if stickers were added to the video. The list of corresponding sticker sets can
	// be received using getAttachedStickerSets
	HasStickers bool
	// True, if the video has no sound
	IsAnimation bool
	// Video minithumbnail; may be null
	Minithumbnail Minithumbnail
	// Video thumbnail in JPEG or MPEG4 format; may be null
	Thumbnail Thumbnail
	// Size of file prefix, which is expected to be preloaded, in bytes
	PreloadPrefixSize int32
	// Timestamp of the frame used as video thumbnail
	CoverFrameTimestamp float64
	// File containing the video
	Video File
}

// StoryVideoTypeID is TL type id of StoryVideo.
const StoryVideoTypeID = 0x562b0a45

// Ensuring interfaces in compile-time for StoryVideo.
var (
	_ bin.Encoder     = &StoryVideo{}
	_ bin.Decoder     = &StoryVideo{}
	_ bin.BareEncoder = &StoryVideo{}
	_ bin.BareDecoder = &StoryVideo{}
)

func (s *StoryVideo) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Duration == 0) {
		return false
	}
	if !(s.Width == 0) {
		return false
	}
	if !(s.Height == 0) {
		return false
	}
	if !(s.HasStickers == false) {
		return false
	}
	if !(s.IsAnimation == false) {
		return false
	}
	if !(s.Minithumbnail.Zero()) {
		return false
	}
	if !(s.Thumbnail.Zero()) {
		return false
	}
	if !(s.PreloadPrefixSize == 0) {
		return false
	}
	if !(s.CoverFrameTimestamp == 0) {
		return false
	}
	if !(s.Video.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryVideo) String() string {
	if s == nil {
		return "StoryVideo(nil)"
	}
	type Alias StoryVideo
	return fmt.Sprintf("StoryVideo%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryVideo) TypeID() uint32 {
	return StoryVideoTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryVideo) TypeName() string {
	return "storyVideo"
}

// TypeInfo returns info about TL type.
func (s *StoryVideo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyVideo",
		ID:   StoryVideoTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Width",
			SchemaName: "width",
		},
		{
			Name:       "Height",
			SchemaName: "height",
		},
		{
			Name:       "HasStickers",
			SchemaName: "has_stickers",
		},
		{
			Name:       "IsAnimation",
			SchemaName: "is_animation",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "PreloadPrefixSize",
			SchemaName: "preload_prefix_size",
		},
		{
			Name:       "CoverFrameTimestamp",
			SchemaName: "cover_frame_timestamp",
		},
		{
			Name:       "Video",
			SchemaName: "video",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryVideo) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyVideo#562b0a45 as nil")
	}
	b.PutID(StoryVideoTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryVideo) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyVideo#562b0a45 as nil")
	}
	b.PutDouble(s.Duration)
	b.PutInt32(s.Width)
	b.PutInt32(s.Height)
	b.PutBool(s.HasStickers)
	b.PutBool(s.IsAnimation)
	if err := s.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyVideo#562b0a45: field minithumbnail: %w", err)
	}
	if err := s.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyVideo#562b0a45: field thumbnail: %w", err)
	}
	b.PutInt32(s.PreloadPrefixSize)
	b.PutDouble(s.CoverFrameTimestamp)
	if err := s.Video.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyVideo#562b0a45: field video: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryVideo) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyVideo#562b0a45 to nil")
	}
	if err := b.ConsumeID(StoryVideoTypeID); err != nil {
		return fmt.Errorf("unable to decode storyVideo#562b0a45: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryVideo) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyVideo#562b0a45 to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field duration: %w", err)
		}
		s.Duration = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field width: %w", err)
		}
		s.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field height: %w", err)
		}
		s.Height = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field has_stickers: %w", err)
		}
		s.HasStickers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field is_animation: %w", err)
		}
		s.IsAnimation = value
	}
	{
		if err := s.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field minithumbnail: %w", err)
		}
	}
	{
		if err := s.Thumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field thumbnail: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field preload_prefix_size: %w", err)
		}
		s.PreloadPrefixSize = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field cover_frame_timestamp: %w", err)
		}
		s.CoverFrameTimestamp = value
	}
	{
		if err := s.Video.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyVideo#562b0a45: field video: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryVideo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyVideo#562b0a45 as nil")
	}
	b.ObjStart()
	b.PutID("storyVideo")
	b.Comma()
	b.FieldStart("duration")
	b.PutDouble(s.Duration)
	b.Comma()
	b.FieldStart("width")
	b.PutInt32(s.Width)
	b.Comma()
	b.FieldStart("height")
	b.PutInt32(s.Height)
	b.Comma()
	b.FieldStart("has_stickers")
	b.PutBool(s.HasStickers)
	b.Comma()
	b.FieldStart("is_animation")
	b.PutBool(s.IsAnimation)
	b.Comma()
	b.FieldStart("minithumbnail")
	if err := s.Minithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyVideo#562b0a45: field minithumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("thumbnail")
	if err := s.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyVideo#562b0a45: field thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("preload_prefix_size")
	b.PutInt32(s.PreloadPrefixSize)
	b.Comma()
	b.FieldStart("cover_frame_timestamp")
	b.PutDouble(s.CoverFrameTimestamp)
	b.Comma()
	b.FieldStart("video")
	if err := s.Video.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyVideo#562b0a45: field video: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryVideo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyVideo#562b0a45 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyVideo"); err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: %w", err)
			}
		case "duration":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field duration: %w", err)
			}
			s.Duration = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field width: %w", err)
			}
			s.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field height: %w", err)
			}
			s.Height = value
		case "has_stickers":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field has_stickers: %w", err)
			}
			s.HasStickers = value
		case "is_animation":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field is_animation: %w", err)
			}
			s.IsAnimation = value
		case "minithumbnail":
			if err := s.Minithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field minithumbnail: %w", err)
			}
		case "thumbnail":
			if err := s.Thumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field thumbnail: %w", err)
			}
		case "preload_prefix_size":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field preload_prefix_size: %w", err)
			}
			s.PreloadPrefixSize = value
		case "cover_frame_timestamp":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field cover_frame_timestamp: %w", err)
			}
			s.CoverFrameTimestamp = value
		case "video":
			if err := s.Video.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyVideo#562b0a45: field video: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDuration returns value of Duration field.
func (s *StoryVideo) GetDuration() (value float64) {
	if s == nil {
		return
	}
	return s.Duration
}

// GetWidth returns value of Width field.
func (s *StoryVideo) GetWidth() (value int32) {
	if s == nil {
		return
	}
	return s.Width
}

// GetHeight returns value of Height field.
func (s *StoryVideo) GetHeight() (value int32) {
	if s == nil {
		return
	}
	return s.Height
}

// GetHasStickers returns value of HasStickers field.
func (s *StoryVideo) GetHasStickers() (value bool) {
	if s == nil {
		return
	}
	return s.HasStickers
}

// GetIsAnimation returns value of IsAnimation field.
func (s *StoryVideo) GetIsAnimation() (value bool) {
	if s == nil {
		return
	}
	return s.IsAnimation
}

// GetMinithumbnail returns value of Minithumbnail field.
func (s *StoryVideo) GetMinithumbnail() (value Minithumbnail) {
	if s == nil {
		return
	}
	return s.Minithumbnail
}

// GetThumbnail returns value of Thumbnail field.
func (s *StoryVideo) GetThumbnail() (value Thumbnail) {
	if s == nil {
		return
	}
	return s.Thumbnail
}

// GetPreloadPrefixSize returns value of PreloadPrefixSize field.
func (s *StoryVideo) GetPreloadPrefixSize() (value int32) {
	if s == nil {
		return
	}
	return s.PreloadPrefixSize
}

// GetCoverFrameTimestamp returns value of CoverFrameTimestamp field.
func (s *StoryVideo) GetCoverFrameTimestamp() (value float64) {
	if s == nil {
		return
	}
	return s.CoverFrameTimestamp
}

// GetVideo returns value of Video field.
func (s *StoryVideo) GetVideo() (value File) {
	if s == nil {
		return
	}
	return s.Video
}

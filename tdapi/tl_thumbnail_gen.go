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

// Thumbnail represents TL type `thumbnail#4a1ae06b`.
type Thumbnail struct {
	// Thumbnail format
	Format ThumbnailFormatClass
	// Thumbnail width
	Width int32
	// Thumbnail height
	Height int32
	// The thumbnail
	File File
}

// ThumbnailTypeID is TL type id of Thumbnail.
const ThumbnailTypeID = 0x4a1ae06b

// Ensuring interfaces in compile-time for Thumbnail.
var (
	_ bin.Encoder     = &Thumbnail{}
	_ bin.Decoder     = &Thumbnail{}
	_ bin.BareEncoder = &Thumbnail{}
	_ bin.BareDecoder = &Thumbnail{}
)

func (t *Thumbnail) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Format == nil) {
		return false
	}
	if !(t.Width == 0) {
		return false
	}
	if !(t.Height == 0) {
		return false
	}
	if !(t.File.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Thumbnail) String() string {
	if t == nil {
		return "Thumbnail(nil)"
	}
	type Alias Thumbnail
	return fmt.Sprintf("Thumbnail%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Thumbnail) TypeID() uint32 {
	return ThumbnailTypeID
}

// TypeName returns name of type in TL schema.
func (*Thumbnail) TypeName() string {
	return "thumbnail"
}

// TypeInfo returns info about TL type.
func (t *Thumbnail) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnail",
		ID:   ThumbnailTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Format",
			SchemaName: "format",
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
			Name:       "File",
			SchemaName: "file",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *Thumbnail) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnail#4a1ae06b as nil")
	}
	b.PutID(ThumbnailTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *Thumbnail) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnail#4a1ae06b as nil")
	}
	if t.Format == nil {
		return fmt.Errorf("unable to encode thumbnail#4a1ae06b: field format is nil")
	}
	if err := t.Format.Encode(b); err != nil {
		return fmt.Errorf("unable to encode thumbnail#4a1ae06b: field format: %w", err)
	}
	b.PutInt32(t.Width)
	b.PutInt32(t.Height)
	if err := t.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode thumbnail#4a1ae06b: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *Thumbnail) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnail#4a1ae06b to nil")
	}
	if err := b.ConsumeID(ThumbnailTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnail#4a1ae06b: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *Thumbnail) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnail#4a1ae06b to nil")
	}
	{
		value, err := DecodeThumbnailFormat(b)
		if err != nil {
			return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field format: %w", err)
		}
		t.Format = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field width: %w", err)
		}
		t.Width = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field height: %w", err)
		}
		t.Height = value
	}
	{
		if err := t.File.Decode(b); err != nil {
			return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field file: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *Thumbnail) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnail#4a1ae06b as nil")
	}
	b.ObjStart()
	b.PutID("thumbnail")
	b.FieldStart("format")
	if t.Format == nil {
		return fmt.Errorf("unable to encode thumbnail#4a1ae06b: field format is nil")
	}
	if err := t.Format.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode thumbnail#4a1ae06b: field format: %w", err)
	}
	b.FieldStart("width")
	b.PutInt32(t.Width)
	b.FieldStart("height")
	b.PutInt32(t.Height)
	b.FieldStart("file")
	if err := t.File.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode thumbnail#4a1ae06b: field file: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *Thumbnail) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnail#4a1ae06b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnail"); err != nil {
				return fmt.Errorf("unable to decode thumbnail#4a1ae06b: %w", err)
			}
		case "format":
			value, err := DecodeTDLibJSONThumbnailFormat(b)
			if err != nil {
				return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field format: %w", err)
			}
			t.Format = value
		case "width":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field width: %w", err)
			}
			t.Width = value
		case "height":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field height: %w", err)
			}
			t.Height = value
		case "file":
			if err := t.File.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode thumbnail#4a1ae06b: field file: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFormat returns value of Format field.
func (t *Thumbnail) GetFormat() (value ThumbnailFormatClass) {
	if t == nil {
		return
	}
	return t.Format
}

// GetWidth returns value of Width field.
func (t *Thumbnail) GetWidth() (value int32) {
	if t == nil {
		return
	}
	return t.Width
}

// GetHeight returns value of Height field.
func (t *Thumbnail) GetHeight() (value int32) {
	if t == nil {
		return
	}
	return t.Height
}

// GetFile returns value of File field.
func (t *Thumbnail) GetFile() (value File) {
	if t == nil {
		return
	}
	return t.File
}

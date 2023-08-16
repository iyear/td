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

// StoryAreaTypeLocation represents TL type `storyAreaTypeLocation#9ebe1186`.
type StoryAreaTypeLocation struct {
	// The location
	Location Location
}

// StoryAreaTypeLocationTypeID is TL type id of StoryAreaTypeLocation.
const StoryAreaTypeLocationTypeID = 0x9ebe1186

// construct implements constructor of StoryAreaTypeClass.
func (s StoryAreaTypeLocation) construct() StoryAreaTypeClass { return &s }

// Ensuring interfaces in compile-time for StoryAreaTypeLocation.
var (
	_ bin.Encoder     = &StoryAreaTypeLocation{}
	_ bin.Decoder     = &StoryAreaTypeLocation{}
	_ bin.BareEncoder = &StoryAreaTypeLocation{}
	_ bin.BareDecoder = &StoryAreaTypeLocation{}

	_ StoryAreaTypeClass = &StoryAreaTypeLocation{}
)

func (s *StoryAreaTypeLocation) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Location.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryAreaTypeLocation) String() string {
	if s == nil {
		return "StoryAreaTypeLocation(nil)"
	}
	type Alias StoryAreaTypeLocation
	return fmt.Sprintf("StoryAreaTypeLocation%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryAreaTypeLocation) TypeID() uint32 {
	return StoryAreaTypeLocationTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryAreaTypeLocation) TypeName() string {
	return "storyAreaTypeLocation"
}

// TypeInfo returns info about TL type.
func (s *StoryAreaTypeLocation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyAreaTypeLocation",
		ID:   StoryAreaTypeLocationTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryAreaTypeLocation) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyAreaTypeLocation#9ebe1186 as nil")
	}
	b.PutID(StoryAreaTypeLocationTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryAreaTypeLocation) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyAreaTypeLocation#9ebe1186 as nil")
	}
	if err := s.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyAreaTypeLocation#9ebe1186: field location: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryAreaTypeLocation) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyAreaTypeLocation#9ebe1186 to nil")
	}
	if err := b.ConsumeID(StoryAreaTypeLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode storyAreaTypeLocation#9ebe1186: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryAreaTypeLocation) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyAreaTypeLocation#9ebe1186 to nil")
	}
	{
		if err := s.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyAreaTypeLocation#9ebe1186: field location: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryAreaTypeLocation) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyAreaTypeLocation#9ebe1186 as nil")
	}
	b.ObjStart()
	b.PutID("storyAreaTypeLocation")
	b.Comma()
	b.FieldStart("location")
	if err := s.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyAreaTypeLocation#9ebe1186: field location: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryAreaTypeLocation) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyAreaTypeLocation#9ebe1186 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyAreaTypeLocation"); err != nil {
				return fmt.Errorf("unable to decode storyAreaTypeLocation#9ebe1186: %w", err)
			}
		case "location":
			if err := s.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyAreaTypeLocation#9ebe1186: field location: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (s *StoryAreaTypeLocation) GetLocation() (value Location) {
	if s == nil {
		return
	}
	return s.Location
}

// StoryAreaTypeVenue represents TL type `storyAreaTypeVenue#18ae4d06`.
type StoryAreaTypeVenue struct {
	// Information about the venue
	Venue Venue
}

// StoryAreaTypeVenueTypeID is TL type id of StoryAreaTypeVenue.
const StoryAreaTypeVenueTypeID = 0x18ae4d06

// construct implements constructor of StoryAreaTypeClass.
func (s StoryAreaTypeVenue) construct() StoryAreaTypeClass { return &s }

// Ensuring interfaces in compile-time for StoryAreaTypeVenue.
var (
	_ bin.Encoder     = &StoryAreaTypeVenue{}
	_ bin.Decoder     = &StoryAreaTypeVenue{}
	_ bin.BareEncoder = &StoryAreaTypeVenue{}
	_ bin.BareDecoder = &StoryAreaTypeVenue{}

	_ StoryAreaTypeClass = &StoryAreaTypeVenue{}
)

func (s *StoryAreaTypeVenue) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Venue.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryAreaTypeVenue) String() string {
	if s == nil {
		return "StoryAreaTypeVenue(nil)"
	}
	type Alias StoryAreaTypeVenue
	return fmt.Sprintf("StoryAreaTypeVenue%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryAreaTypeVenue) TypeID() uint32 {
	return StoryAreaTypeVenueTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryAreaTypeVenue) TypeName() string {
	return "storyAreaTypeVenue"
}

// TypeInfo returns info about TL type.
func (s *StoryAreaTypeVenue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyAreaTypeVenue",
		ID:   StoryAreaTypeVenueTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Venue",
			SchemaName: "venue",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryAreaTypeVenue) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyAreaTypeVenue#18ae4d06 as nil")
	}
	b.PutID(StoryAreaTypeVenueTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryAreaTypeVenue) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyAreaTypeVenue#18ae4d06 as nil")
	}
	if err := s.Venue.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyAreaTypeVenue#18ae4d06: field venue: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryAreaTypeVenue) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyAreaTypeVenue#18ae4d06 to nil")
	}
	if err := b.ConsumeID(StoryAreaTypeVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode storyAreaTypeVenue#18ae4d06: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryAreaTypeVenue) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyAreaTypeVenue#18ae4d06 to nil")
	}
	{
		if err := s.Venue.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyAreaTypeVenue#18ae4d06: field venue: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryAreaTypeVenue) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyAreaTypeVenue#18ae4d06 as nil")
	}
	b.ObjStart()
	b.PutID("storyAreaTypeVenue")
	b.Comma()
	b.FieldStart("venue")
	if err := s.Venue.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode storyAreaTypeVenue#18ae4d06: field venue: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryAreaTypeVenue) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyAreaTypeVenue#18ae4d06 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyAreaTypeVenue"); err != nil {
				return fmt.Errorf("unable to decode storyAreaTypeVenue#18ae4d06: %w", err)
			}
		case "venue":
			if err := s.Venue.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode storyAreaTypeVenue#18ae4d06: field venue: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVenue returns value of Venue field.
func (s *StoryAreaTypeVenue) GetVenue() (value Venue) {
	if s == nil {
		return
	}
	return s.Venue
}

// StoryAreaTypeClassName is schema name of StoryAreaTypeClass.
const StoryAreaTypeClassName = "StoryAreaType"

// StoryAreaTypeClass represents StoryAreaType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStoryAreaType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StoryAreaTypeLocation: // storyAreaTypeLocation#9ebe1186
//	case *tdapi.StoryAreaTypeVenue: // storyAreaTypeVenue#18ae4d06
//	default: panic(v)
//	}
type StoryAreaTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryAreaTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeStoryAreaType implements binary de-serialization for StoryAreaTypeClass.
func DecodeStoryAreaType(buf *bin.Buffer) (StoryAreaTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryAreaTypeLocationTypeID:
		// Decoding storyAreaTypeLocation#9ebe1186.
		v := StoryAreaTypeLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case StoryAreaTypeVenueTypeID:
		// Decoding storyAreaTypeVenue#18ae4d06.
		v := StoryAreaTypeVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryAreaTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryAreaTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStoryAreaType implements binary de-serialization for StoryAreaTypeClass.
func DecodeTDLibJSONStoryAreaType(buf tdjson.Decoder) (StoryAreaTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "storyAreaTypeLocation":
		// Decoding storyAreaTypeLocation#9ebe1186.
		v := StoryAreaTypeLocation{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case "storyAreaTypeVenue":
		// Decoding storyAreaTypeVenue#18ae4d06.
		v := StoryAreaTypeVenue{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryAreaTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryAreaTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StoryAreaType boxes the StoryAreaTypeClass providing a helper.
type StoryAreaTypeBox struct {
	StoryAreaType StoryAreaTypeClass
}

// Decode implements bin.Decoder for StoryAreaTypeBox.
func (b *StoryAreaTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryAreaTypeBox to nil")
	}
	v, err := DecodeStoryAreaType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryAreaType = v
	return nil
}

// Encode implements bin.Encode for StoryAreaTypeBox.
func (b *StoryAreaTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryAreaType == nil {
		return fmt.Errorf("unable to encode StoryAreaTypeClass as nil")
	}
	return b.StoryAreaType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StoryAreaTypeBox.
func (b *StoryAreaTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryAreaTypeBox to nil")
	}
	v, err := DecodeTDLibJSONStoryAreaType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryAreaType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StoryAreaTypeBox.
func (b *StoryAreaTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StoryAreaType == nil {
		return fmt.Errorf("unable to encode StoryAreaTypeClass as nil")
	}
	return b.StoryAreaType.EncodeTDLibJSON(buf)
}

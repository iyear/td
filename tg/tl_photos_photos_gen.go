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

// PhotosPhotos represents TL type `photos.photos#8dca6aa5`.
// Full list of photos with auxiliary data.
//
// See https://core.telegram.org/constructor/photos.photos for reference.
type PhotosPhotos struct {
	// List of photos
	Photos []PhotoClass
	// List of mentioned users
	Users []UserClass
}

// PhotosPhotosTypeID is TL type id of PhotosPhotos.
const PhotosPhotosTypeID = 0x8dca6aa5

func (p *PhotosPhotos) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Photos == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotosPhotos) String() string {
	if p == nil {
		return "PhotosPhotos(nil)"
	}
	type Alias PhotosPhotos
	return fmt.Sprintf("PhotosPhotos%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhotosPhotos) TypeID() uint32 {
	return PhotosPhotosTypeID
}

// Encode implements bin.Encoder.
func (p *PhotosPhotos) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photos#8dca6aa5 as nil")
	}
	b.PutID(PhotosPhotosTypeID)
	b.PutVectorHeader(len(p.Photos))
	for idx, v := range p.Photos {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field photos element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field photos element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPhotos returns value of Photos field.
func (p *PhotosPhotos) GetPhotos() (value []PhotoClass) {
	return p.Photos
}

// GetUsers returns value of Users field.
func (p *PhotosPhotos) GetUsers() (value []UserClass) {
	return p.Users
}

// Decode implements bin.Decoder.
func (p *PhotosPhotos) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photos#8dca6aa5 to nil")
	}
	if err := b.ConsumeID(PhotosPhotosTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.photos#8dca6aa5: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field photos: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field photos: %w", err)
			}
			p.Photos = append(p.Photos, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// construct implements constructor of PhotosPhotosClass.
func (p PhotosPhotos) construct() PhotosPhotosClass { return &p }

// Ensuring interfaces in compile-time for PhotosPhotos.
var (
	_ bin.Encoder = &PhotosPhotos{}
	_ bin.Decoder = &PhotosPhotos{}

	_ PhotosPhotosClass = &PhotosPhotos{}
)

// PhotosPhotosSlice represents TL type `photos.photosSlice#15051f54`.
// Incomplete list of photos with auxiliary data.
//
// See https://core.telegram.org/constructor/photos.photosSlice for reference.
type PhotosPhotosSlice struct {
	// Total number of photos
	Count int
	// List of photos
	Photos []PhotoClass
	// List of mentioned users
	Users []UserClass
}

// PhotosPhotosSliceTypeID is TL type id of PhotosPhotosSlice.
const PhotosPhotosSliceTypeID = 0x15051f54

func (p *PhotosPhotosSlice) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Count == 0) {
		return false
	}
	if !(p.Photos == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotosPhotosSlice) String() string {
	if p == nil {
		return "PhotosPhotosSlice(nil)"
	}
	type Alias PhotosPhotosSlice
	return fmt.Sprintf("PhotosPhotosSlice%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhotosPhotosSlice) TypeID() uint32 {
	return PhotosPhotosSliceTypeID
}

// Encode implements bin.Encoder.
func (p *PhotosPhotosSlice) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photosSlice#15051f54 as nil")
	}
	b.PutID(PhotosPhotosSliceTypeID)
	b.PutInt(p.Count)
	b.PutVectorHeader(len(p.Photos))
	for idx, v := range p.Photos {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field photos element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field photos element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (p *PhotosPhotosSlice) GetCount() (value int) {
	return p.Count
}

// GetPhotos returns value of Photos field.
func (p *PhotosPhotosSlice) GetPhotos() (value []PhotoClass) {
	return p.Photos
}

// GetUsers returns value of Users field.
func (p *PhotosPhotosSlice) GetUsers() (value []UserClass) {
	return p.Users
}

// Decode implements bin.Decoder.
func (p *PhotosPhotosSlice) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photosSlice#15051f54 to nil")
	}
	if err := b.ConsumeID(PhotosPhotosSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.photosSlice#15051f54: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field count: %w", err)
		}
		p.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field photos: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field photos: %w", err)
			}
			p.Photos = append(p.Photos, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// construct implements constructor of PhotosPhotosClass.
func (p PhotosPhotosSlice) construct() PhotosPhotosClass { return &p }

// Ensuring interfaces in compile-time for PhotosPhotosSlice.
var (
	_ bin.Encoder = &PhotosPhotosSlice{}
	_ bin.Decoder = &PhotosPhotosSlice{}

	_ PhotosPhotosClass = &PhotosPhotosSlice{}
)

// PhotosPhotosClass represents photos.Photos generic type.
//
// See https://core.telegram.org/type/photos.Photos for reference.
//
// Example:
//  g, err := tg.DecodePhotosPhotos(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PhotosPhotos: // photos.photos#8dca6aa5
//  case *tg.PhotosPhotosSlice: // photos.photosSlice#15051f54
//  default: panic(v)
//  }
type PhotosPhotosClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhotosPhotosClass

	// List of photos
	GetPhotos() (value []PhotoClass)
	// List of mentioned users
	GetUsers() (value []UserClass)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodePhotosPhotos implements binary de-serialization for PhotosPhotosClass.
func DecodePhotosPhotos(buf *bin.Buffer) (PhotosPhotosClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotosPhotosTypeID:
		// Decoding photos.photos#8dca6aa5.
		v := PhotosPhotos{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotosPhotosClass: %w", err)
		}
		return &v, nil
	case PhotosPhotosSliceTypeID:
		// Decoding photos.photosSlice#15051f54.
		v := PhotosPhotosSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotosPhotosClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotosPhotosClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhotosPhotos boxes the PhotosPhotosClass providing a helper.
type PhotosPhotosBox struct {
	Photos PhotosPhotosClass
}

// Decode implements bin.Decoder for PhotosPhotosBox.
func (b *PhotosPhotosBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotosPhotosBox to nil")
	}
	v, err := DecodePhotosPhotos(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Photos = v
	return nil
}

// Encode implements bin.Encode for PhotosPhotosBox.
func (b *PhotosPhotosBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Photos == nil {
		return fmt.Errorf("unable to encode PhotosPhotosClass as nil")
	}
	return b.Photos.Encode(buf)
}

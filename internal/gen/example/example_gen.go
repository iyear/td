// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"fmt"

	"github.com/ernado/td/internal/bin"
	"golang.org/x/xerrors"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}

// Int32 represents TL type int32#5cb934fa.
type Int32 struct {
}

// Int32TypeID is TL type id of Int32.
const Int32TypeID = 0x5cb934fa

// Encode implements bin.Encoder.
func (i Int32) Encode(b *bin.Buffer) {
	b.PutID(Int32TypeID)
}

// Decode implements bin.Decoder.
func (i *Int32) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(Int32TypeID); err != nil {
		return fmt.Errorf("unable to decode int32#5cb934fa: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Int32.
var (
	_ bin.Encoder = Int32{}
	_ bin.Decoder = &Int32{}
)

// String represents TL type string#b5286e24.
type String struct {
}

// StringTypeID is TL type id of String.
const StringTypeID = 0xb5286e24

// Encode implements bin.Encoder.
func (s String) Encode(b *bin.Buffer) {
	b.PutID(StringTypeID)
}

// Decode implements bin.Decoder.
func (s *String) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(StringTypeID); err != nil {
		return fmt.Errorf("unable to decode string#b5286e24: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for String.
var (
	_ bin.Encoder = String{}
	_ bin.Decoder = &String{}
)

// BoolFalse represents TL type boolFalse#bc799737.
type BoolFalse struct {
}

// BoolFalseTypeID is TL type id of BoolFalse.
const BoolFalseTypeID = 0xbc799737

// Encode implements bin.Encoder.
func (b BoolFalse) Encode(buf *bin.Buffer) {
	buf.PutID(BoolFalseTypeID)
}

// Decode implements bin.Decoder.
func (b *BoolFalse) Decode(buf *bin.Buffer) error {
	if err := buf.ConsumeID(BoolFalseTypeID); err != nil {
		return fmt.Errorf("unable to decode boolFalse#bc799737: %w", err)
	}
	return nil
}

// construct implements constructor of Bool.
func (b BoolFalse) construct() Bool { return &b }

// Ensuring interfaces in compile-time for BoolFalse.
var (
	_ bin.Encoder = BoolFalse{}
	_ bin.Decoder = &BoolFalse{}

	_ Bool = &BoolFalse{}
)

// BoolTrue represents TL type boolTrue#997275b5.
type BoolTrue struct {
}

// BoolTrueTypeID is TL type id of BoolTrue.
const BoolTrueTypeID = 0x997275b5

// Encode implements bin.Encoder.
func (b BoolTrue) Encode(buf *bin.Buffer) {
	buf.PutID(BoolTrueTypeID)
}

// Decode implements bin.Decoder.
func (b *BoolTrue) Decode(buf *bin.Buffer) error {
	if err := buf.ConsumeID(BoolTrueTypeID); err != nil {
		return fmt.Errorf("unable to decode boolTrue#997275b5: %w", err)
	}
	return nil
}

// construct implements constructor of Bool.
func (b BoolTrue) construct() Bool { return &b }

// Ensuring interfaces in compile-time for BoolTrue.
var (
	_ bin.Encoder = BoolTrue{}
	_ bin.Decoder = &BoolTrue{}

	_ Bool = &BoolTrue{}
)

// An object of this type can be returned on every function call, in case of an error
type Error struct {
	// Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
	Code int32
	// Error message; subject to future changes
	Message string
	// Temporary field of Error.
	Temporary bool
}

// ErrorTypeID is TL type id of Error.
const ErrorTypeID = 0x14feebbc

// Encode implements bin.Encoder.
func (e Error) Encode(b *bin.Buffer) {
	b.PutID(ErrorTypeID)
	b.PutInt32(e.Code)
	b.PutString(e.Message)
	b.PutBool(e.Temporary)
}

// Decode implements bin.Decoder.
func (e *Error) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(ErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode error#14feebbc: %w", err)
	}
	{
		v, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode error#14feebbc: field code: %w", err)
		}
		e.Code = v
	}
	{
		v, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode error#14feebbc: field message: %w", err)
		}
		e.Message = v
	}
	{
		v, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode error#14feebbc: field temporary: %w", err)
		}
		e.Temporary = v
	}
	return nil
}

// Ensuring interfaces in compile-time for Error.
var (
	_ bin.Encoder = Error{}
	_ bin.Decoder = &Error{}
)

// can be returned by functions as result.
type Ok struct {
}

// OkTypeID is TL type id of Ok.
const OkTypeID = 0xd4edbe69

// Encode implements bin.Encoder.
func (o Ok) Encode(b *bin.Buffer) {
	b.PutID(OkTypeID)
}

// Decode implements bin.Decoder.
func (o *Ok) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(OkTypeID); err != nil {
		return fmt.Errorf("unable to decode ok#d4edbe69: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Ok.
var (
	_ bin.Encoder = Ok{}
	_ bin.Decoder = &Ok{}
)

// Message represents TL type message#ec200d96.
type Message struct {
	// Err field of Message.
	Err Error
}

// MessageTypeID is TL type id of Message.
const MessageTypeID = 0xec200d96

// Encode implements bin.Encoder.
func (m Message) Encode(b *bin.Buffer) {
	b.PutID(MessageTypeID)
	m.Err.Encode(b)
}

// Decode implements bin.Decoder.
func (m *Message) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(MessageTypeID); err != nil {
		return fmt.Errorf("unable to decode message#ec200d96: %w", err)
	}
	{
		if err := m.Err.Decode(b); err != nil {
			return fmt.Errorf("unable to decode message#ec200d96: field err: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for Message.
var (
	_ bin.Encoder = Message{}
	_ bin.Decoder = &Message{}
)

// SMS represents TL type sms#ed8bebfe.
type SMS struct {
	// Text field of SMS.
	Text string
}

// SMSTypeID is TL type id of SMS.
const SMSTypeID = 0xed8bebfe

// Encode implements bin.Encoder.
func (s SMS) Encode(b *bin.Buffer) {
	b.PutID(SMSTypeID)
	b.PutString(s.Text)
}

// Decode implements bin.Decoder.
func (s *SMS) Decode(b *bin.Buffer) error {
	if err := b.ConsumeID(SMSTypeID); err != nil {
		return fmt.Errorf("unable to decode sms#ed8bebfe: %w", err)
	}
	{
		v, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sms#ed8bebfe: field text: %w", err)
		}
		s.Text = v
	}
	return nil
}

// Ensuring interfaces in compile-time for SMS.
var (
	_ bin.Encoder = SMS{}
	_ bin.Decoder = &SMS{}
)

// Bool represents Bool generic type.
//
// Example:
//  g, err := DecodeBool(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *BoolFalse: // boolFalse#bc799737
//  case *BoolTrue: // boolTrue#997275b5
//  default: panic(v)
//  }
type Bool interface {
	bin.Encoder
	bin.Decoder
	construct() Bool
}

// DecodeBool implements binary de-serialization for Bool.
func DecodeBool(buf *bin.Buffer) (Bool, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {

	case 0xbc799737:
		v := BoolFalse{}
		if err := v.Decode(buf); err != nil {
			return nil, xerrors.Errorf("unable to decode Bool: %w", err)
		}
		return &v, nil

	case 0x997275b5:
		v := BoolTrue{}
		if err := v.Decode(buf); err != nil {
			return nil, xerrors.Errorf("unable to decode Bool: %w", err)
		}
		return &v, nil

	default:
		return nil, xerrors.Errorf("unable to decode Bool: %w", bin.NewUnexpectedID(id))
	}
}

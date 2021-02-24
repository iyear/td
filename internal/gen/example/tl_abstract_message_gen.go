// Code generated by gotdgen, DO NOT EDIT.

package td

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

// BigMessage represents TL type `bigMessage#7490dcc5`.
//
// See https://localhost:80/doc/constructor/bigMessage for reference.
type BigMessage struct {
	// ID field of BigMessage.
	ID int32 `tl:"id"`
	// Count field of BigMessage.
	Count int32 `tl:"count"`
	// TargetId field of BigMessage.
	TargetId int32 `tl:"targetId"`
	// Escape field of BigMessage.
	Escape bool `tl:"escape"`
	// Summary field of BigMessage.
	Summary bool `tl:"summary"`
}

// BigMessageTypeID is TL type id of BigMessage.
const BigMessageTypeID = 0x7490dcc5

func (b *BigMessage) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ID == 0) {
		return false
	}
	if !(b.Count == 0) {
		return false
	}
	if !(b.TargetId == 0) {
		return false
	}
	if !(b.Escape == false) {
		return false
	}
	if !(b.Summary == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BigMessage) String() string {
	if b == nil {
		return "BigMessage(nil)"
	}
	type Alias BigMessage
	return fmt.Sprintf("BigMessage%+v", Alias(*b))
}

// FillFrom fills BigMessage from given interface.
func (b *BigMessage) FillFrom(from interface {
	GetID() (value int32)
	GetCount() (value int32)
	GetTargetId() (value int32)
	GetEscape() (value bool)
	GetSummary() (value bool)
}) {
	b.ID = from.GetID()
	b.Count = from.GetCount()
	b.TargetId = from.GetTargetId()
	b.Escape = from.GetEscape()
	b.Summary = from.GetSummary()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (b *BigMessage) TypeID() uint32 {
	return BigMessageTypeID
}

// TypeName returns name of type in TL schema.
func (b *BigMessage) TypeName() string {
	return "bigMessage"
}

// Encode implements bin.Encoder.
func (b *BigMessage) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bigMessage#7490dcc5 as nil")
	}
	buf.PutID(BigMessageTypeID)
	buf.PutInt32(b.ID)
	buf.PutInt32(b.Count)
	buf.PutInt32(b.TargetId)
	buf.PutBool(b.Escape)
	buf.PutBool(b.Summary)
	return nil
}

// GetID returns value of ID field.
func (b *BigMessage) GetID() (value int32) {
	return b.ID
}

// GetCount returns value of Count field.
func (b *BigMessage) GetCount() (value int32) {
	return b.Count
}

// GetTargetId returns value of TargetId field.
func (b *BigMessage) GetTargetId() (value int32) {
	return b.TargetId
}

// GetEscape returns value of Escape field.
func (b *BigMessage) GetEscape() (value bool) {
	return b.Escape
}

// GetSummary returns value of Summary field.
func (b *BigMessage) GetSummary() (value bool) {
	return b.Summary
}

// Decode implements bin.Decoder.
func (b *BigMessage) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bigMessage#7490dcc5 to nil")
	}
	if err := buf.ConsumeID(BigMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode bigMessage#7490dcc5: %w", err)
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode bigMessage#7490dcc5: field id: %w", err)
		}
		b.ID = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode bigMessage#7490dcc5: field count: %w", err)
		}
		b.Count = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode bigMessage#7490dcc5: field targetId: %w", err)
		}
		b.TargetId = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode bigMessage#7490dcc5: field escape: %w", err)
		}
		b.Escape = value
	}
	{
		value, err := buf.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode bigMessage#7490dcc5: field summary: %w", err)
		}
		b.Summary = value
	}
	return nil
}

// construct implements constructor of AbstractMessageClass.
func (b BigMessage) construct() AbstractMessageClass { return &b }

// Ensuring interfaces in compile-time for BigMessage.
var (
	_ bin.Encoder = &BigMessage{}
	_ bin.Decoder = &BigMessage{}

	_ AbstractMessageClass = &BigMessage{}
)

// NoMessage represents TL type `noMessage#ee6324c4`.
//
// See https://localhost:80/doc/constructor/noMessage for reference.
type NoMessage struct {
}

// NoMessageTypeID is TL type id of NoMessage.
const NoMessageTypeID = 0xee6324c4

func (n *NoMessage) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NoMessage) String() string {
	if n == nil {
		return "NoMessage(nil)"
	}
	type Alias NoMessage
	return fmt.Sprintf("NoMessage%+v", Alias(*n))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *NoMessage) TypeID() uint32 {
	return NoMessageTypeID
}

// TypeName returns name of type in TL schema.
func (n *NoMessage) TypeName() string {
	return "noMessage"
}

// Encode implements bin.Encoder.
func (n *NoMessage) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode noMessage#ee6324c4 as nil")
	}
	b.PutID(NoMessageTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *NoMessage) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode noMessage#ee6324c4 to nil")
	}
	if err := b.ConsumeID(NoMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode noMessage#ee6324c4: %w", err)
	}
	return nil
}

// construct implements constructor of AbstractMessageClass.
func (n NoMessage) construct() AbstractMessageClass { return &n }

// Ensuring interfaces in compile-time for NoMessage.
var (
	_ bin.Encoder = &NoMessage{}
	_ bin.Decoder = &NoMessage{}

	_ AbstractMessageClass = &NoMessage{}
)

// TargetsMessage represents TL type `targetsMessage#cc6136f1`.
//
// See https://localhost:80/doc/constructor/targetsMessage for reference.
type TargetsMessage struct {
	// Targets field of TargetsMessage.
	Targets []int32 `tl:"targets"`
}

// TargetsMessageTypeID is TL type id of TargetsMessage.
const TargetsMessageTypeID = 0xcc6136f1

func (t *TargetsMessage) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Targets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TargetsMessage) String() string {
	if t == nil {
		return "TargetsMessage(nil)"
	}
	type Alias TargetsMessage
	return fmt.Sprintf("TargetsMessage%+v", Alias(*t))
}

// FillFrom fills TargetsMessage from given interface.
func (t *TargetsMessage) FillFrom(from interface {
	GetTargets() (value []int32)
}) {
	t.Targets = from.GetTargets()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TargetsMessage) TypeID() uint32 {
	return TargetsMessageTypeID
}

// TypeName returns name of type in TL schema.
func (t *TargetsMessage) TypeName() string {
	return "targetsMessage"
}

// Encode implements bin.Encoder.
func (t *TargetsMessage) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode targetsMessage#cc6136f1 as nil")
	}
	b.PutID(TargetsMessageTypeID)
	b.PutVectorHeader(len(t.Targets))
	for _, v := range t.Targets {
		b.PutInt32(v)
	}
	return nil
}

// GetTargets returns value of Targets field.
func (t *TargetsMessage) GetTargets() (value []int32) {
	return t.Targets
}

// Decode implements bin.Decoder.
func (t *TargetsMessage) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode targetsMessage#cc6136f1 to nil")
	}
	if err := b.ConsumeID(TargetsMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode targetsMessage#cc6136f1: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode targetsMessage#cc6136f1: field targets: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode targetsMessage#cc6136f1: field targets: %w", err)
			}
			t.Targets = append(t.Targets, value)
		}
	}
	return nil
}

// construct implements constructor of AbstractMessageClass.
func (t TargetsMessage) construct() AbstractMessageClass { return &t }

// Ensuring interfaces in compile-time for TargetsMessage.
var (
	_ bin.Encoder = &TargetsMessage{}
	_ bin.Decoder = &TargetsMessage{}

	_ AbstractMessageClass = &TargetsMessage{}
)

// FieldsMessage represents TL type `fieldsMessage#947225b5`.
//
// See https://localhost:80/doc/constructor/fieldsMessage for reference.
type FieldsMessage struct {
	// Flags field of FieldsMessage.
	Flags bin.Fields `tl:"flags"`
	// Escape field of FieldsMessage.
	//
	// Use SetEscape and GetEscape helpers.
	Escape bool `tl:"escape"`
	// TTLSeconds field of FieldsMessage.
	//
	// Use SetTTLSeconds and GetTTLSeconds helpers.
	TTLSeconds int `tl:"ttl_seconds"`
}

// FieldsMessageTypeID is TL type id of FieldsMessage.
const FieldsMessageTypeID = 0x947225b5

func (f *FieldsMessage) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Flags.Zero()) {
		return false
	}
	if !(f.Escape == false) {
		return false
	}
	if !(f.TTLSeconds == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FieldsMessage) String() string {
	if f == nil {
		return "FieldsMessage(nil)"
	}
	type Alias FieldsMessage
	return fmt.Sprintf("FieldsMessage%+v", Alias(*f))
}

// FillFrom fills FieldsMessage from given interface.
func (f *FieldsMessage) FillFrom(from interface {
	GetEscape() (value bool, ok bool)
	GetTTLSeconds() (value int, ok bool)
}) {
	if val, ok := from.GetEscape(); ok {
		f.Escape = val
	}

	if val, ok := from.GetTTLSeconds(); ok {
		f.TTLSeconds = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *FieldsMessage) TypeID() uint32 {
	return FieldsMessageTypeID
}

// TypeName returns name of type in TL schema.
func (f *FieldsMessage) TypeName() string {
	return "fieldsMessage"
}

// Encode implements bin.Encoder.
func (f *FieldsMessage) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fieldsMessage#947225b5 as nil")
	}
	b.PutID(FieldsMessageTypeID)
	if !(f.Escape == false) {
		f.Flags.Set(0)
	}
	if !(f.TTLSeconds == 0) {
		f.Flags.Set(1)
	}
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode fieldsMessage#947225b5: field flags: %w", err)
	}
	if f.Flags.Has(0) {
		b.PutBool(f.Escape)
	}
	if f.Flags.Has(1) {
		b.PutInt(f.TTLSeconds)
	}
	return nil
}

// SetEscape sets value of Escape conditional field.
func (f *FieldsMessage) SetEscape(value bool) {
	f.Flags.Set(0)
	f.Escape = value
}

// GetEscape returns value of Escape conditional field and
// boolean which is true if field was set.
func (f *FieldsMessage) GetEscape() (value bool, ok bool) {
	if !f.Flags.Has(0) {
		return value, false
	}
	return f.Escape, true
}

// SetTTLSeconds sets value of TTLSeconds conditional field.
func (f *FieldsMessage) SetTTLSeconds(value int) {
	f.Flags.Set(1)
	f.TTLSeconds = value
}

// GetTTLSeconds returns value of TTLSeconds conditional field and
// boolean which is true if field was set.
func (f *FieldsMessage) GetTTLSeconds() (value int, ok bool) {
	if !f.Flags.Has(1) {
		return value, false
	}
	return f.TTLSeconds, true
}

// Decode implements bin.Decoder.
func (f *FieldsMessage) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fieldsMessage#947225b5 to nil")
	}
	if err := b.ConsumeID(FieldsMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode fieldsMessage#947225b5: %w", err)
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode fieldsMessage#947225b5: field flags: %w", err)
		}
	}
	if f.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode fieldsMessage#947225b5: field escape: %w", err)
		}
		f.Escape = value
	}
	if f.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode fieldsMessage#947225b5: field ttl_seconds: %w", err)
		}
		f.TTLSeconds = value
	}
	return nil
}

// construct implements constructor of AbstractMessageClass.
func (f FieldsMessage) construct() AbstractMessageClass { return &f }

// Ensuring interfaces in compile-time for FieldsMessage.
var (
	_ bin.Encoder = &FieldsMessage{}
	_ bin.Decoder = &FieldsMessage{}

	_ AbstractMessageClass = &FieldsMessage{}
)

// BytesMessage represents TL type `bytesMessage#f990a67d`.
//
// See https://localhost:80/doc/constructor/bytesMessage for reference.
type BytesMessage struct {
	// Data field of BytesMessage.
	Data []byte `tl:"data"`
}

// BytesMessageTypeID is TL type id of BytesMessage.
const BytesMessageTypeID = 0xf990a67d

func (b *BytesMessage) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Data == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BytesMessage) String() string {
	if b == nil {
		return "BytesMessage(nil)"
	}
	type Alias BytesMessage
	return fmt.Sprintf("BytesMessage%+v", Alias(*b))
}

// FillFrom fills BytesMessage from given interface.
func (b *BytesMessage) FillFrom(from interface {
	GetData() (value []byte)
}) {
	b.Data = from.GetData()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (b *BytesMessage) TypeID() uint32 {
	return BytesMessageTypeID
}

// TypeName returns name of type in TL schema.
func (b *BytesMessage) TypeName() string {
	return "bytesMessage"
}

// Encode implements bin.Encoder.
func (b *BytesMessage) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytesMessage#f990a67d as nil")
	}
	buf.PutID(BytesMessageTypeID)
	buf.PutBytes(b.Data)
	return nil
}

// GetData returns value of Data field.
func (b *BytesMessage) GetData() (value []byte) {
	return b.Data
}

// Decode implements bin.Decoder.
func (b *BytesMessage) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytesMessage#f990a67d to nil")
	}
	if err := buf.ConsumeID(BytesMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode bytesMessage#f990a67d: %w", err)
	}
	{
		value, err := buf.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode bytesMessage#f990a67d: field data: %w", err)
		}
		b.Data = value
	}
	return nil
}

// construct implements constructor of AbstractMessageClass.
func (b BytesMessage) construct() AbstractMessageClass { return &b }

// Ensuring interfaces in compile-time for BytesMessage.
var (
	_ bin.Encoder = &BytesMessage{}
	_ bin.Decoder = &BytesMessage{}

	_ AbstractMessageClass = &BytesMessage{}
)

// AbstractMessageClass represents AbstractMessage generic type.
//
// See https://localhost:80/doc/type/AbstractMessage for reference.
//
// Example:
//  g, err := td.DecodeAbstractMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *td.BigMessage: // bigMessage#7490dcc5
//  case *td.NoMessage: // noMessage#ee6324c4
//  case *td.TargetsMessage: // targetsMessage#cc6136f1
//  case *td.FieldsMessage: // fieldsMessage#947225b5
//  case *td.BytesMessage: // bytesMessage#f990a67d
//  default: panic(v)
//  }
type AbstractMessageClass interface {
	bin.Encoder
	bin.Decoder
	construct() AbstractMessageClass

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

// DecodeAbstractMessage implements binary de-serialization for AbstractMessageClass.
func DecodeAbstractMessage(buf *bin.Buffer) (AbstractMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BigMessageTypeID:
		// Decoding bigMessage#7490dcc5.
		v := BigMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AbstractMessageClass: %w", err)
		}
		return &v, nil
	case NoMessageTypeID:
		// Decoding noMessage#ee6324c4.
		v := NoMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AbstractMessageClass: %w", err)
		}
		return &v, nil
	case TargetsMessageTypeID:
		// Decoding targetsMessage#cc6136f1.
		v := TargetsMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AbstractMessageClass: %w", err)
		}
		return &v, nil
	case FieldsMessageTypeID:
		// Decoding fieldsMessage#947225b5.
		v := FieldsMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AbstractMessageClass: %w", err)
		}
		return &v, nil
	case BytesMessageTypeID:
		// Decoding bytesMessage#f990a67d.
		v := BytesMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AbstractMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AbstractMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// AbstractMessage boxes the AbstractMessageClass providing a helper.
type AbstractMessageBox struct {
	AbstractMessage AbstractMessageClass
}

// Decode implements bin.Decoder for AbstractMessageBox.
func (b *AbstractMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AbstractMessageBox to nil")
	}
	v, err := DecodeAbstractMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AbstractMessage = v
	return nil
}

// Encode implements bin.Encode for AbstractMessageBox.
func (b *AbstractMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AbstractMessage == nil {
		return fmt.Errorf("unable to encode AbstractMessageClass as nil")
	}
	return b.AbstractMessage.Encode(buf)
}

// AbstractMessageClassSlice is adapter for slice of AbstractMessageClass.
type AbstractMessageClassSlice []AbstractMessageClass

// First returns first element of slice (if exists).
func (s AbstractMessageClassSlice) First() (v AbstractMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s AbstractMessageClassSlice) Last() (v AbstractMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *AbstractMessageClassSlice) PopFirst() (v AbstractMessageClass, ok bool) {
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
func (s *AbstractMessageClassSlice) Pop() (v AbstractMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

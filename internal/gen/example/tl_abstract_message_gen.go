// Code generated by gotdgen, DO NOT EDIT.

package td

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

// BigMessage represents TL type `bigMessage#7490dcc5`.
//
// See https://localhost:80/doc/constructor/bigMessage for reference.
type BigMessage struct {
	// ID field of BigMessage.
	ID int32
	// Count field of BigMessage.
	Count int32
	// TargetID field of BigMessage.
	TargetID int32
	// Escape field of BigMessage.
	Escape bool
	// Summary field of BigMessage.
	Summary bool
}

// BigMessageTypeID is TL type id of BigMessage.
const BigMessageTypeID = 0x7490dcc5

// construct implements constructor of AbstractMessageClass.
func (b BigMessage) construct() AbstractMessageClass { return &b }

// Ensuring interfaces in compile-time for BigMessage.
var (
	_ bin.Encoder     = &BigMessage{}
	_ bin.Decoder     = &BigMessage{}
	_ bin.BareEncoder = &BigMessage{}
	_ bin.BareDecoder = &BigMessage{}

	_ AbstractMessageClass = &BigMessage{}
)

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
	if !(b.TargetID == 0) {
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

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BigMessage) TypeID() uint32 {
	return BigMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*BigMessage) TypeName() string {
	return "bigMessage"
}

// TypeInfo returns info about TL type.
func (b *BigMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bigMessage",
		ID:   BigMessageTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "TargetID",
			SchemaName: "targetId",
		},
		{
			Name:       "Escape",
			SchemaName: "escape",
		},
		{
			Name:       "Summary",
			SchemaName: "summary",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BigMessage) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bigMessage#7490dcc5 as nil")
	}
	buf.PutID(BigMessageTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BigMessage) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bigMessage#7490dcc5 as nil")
	}
	buf.PutInt32(b.ID)
	buf.PutInt32(b.Count)
	buf.PutInt32(b.TargetID)
	buf.PutBool(b.Escape)
	buf.PutBool(b.Summary)
	return nil
}

// Decode implements bin.Decoder.
func (b *BigMessage) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bigMessage#7490dcc5 to nil")
	}
	if err := buf.ConsumeID(BigMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode bigMessage#7490dcc5: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BigMessage) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bigMessage#7490dcc5 to nil")
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
		b.TargetID = value
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

// GetID returns value of ID field.
func (b *BigMessage) GetID() (value int32) {
	if b == nil {
		return
	}
	return b.ID
}

// GetCount returns value of Count field.
func (b *BigMessage) GetCount() (value int32) {
	if b == nil {
		return
	}
	return b.Count
}

// GetTargetID returns value of TargetID field.
func (b *BigMessage) GetTargetID() (value int32) {
	if b == nil {
		return
	}
	return b.TargetID
}

// GetEscape returns value of Escape field.
func (b *BigMessage) GetEscape() (value bool) {
	if b == nil {
		return
	}
	return b.Escape
}

// GetSummary returns value of Summary field.
func (b *BigMessage) GetSummary() (value bool) {
	if b == nil {
		return
	}
	return b.Summary
}

// NoMessage represents TL type `noMessage#ee6324c4`.
//
// See https://localhost:80/doc/constructor/noMessage for reference.
type NoMessage struct {
}

// NoMessageTypeID is TL type id of NoMessage.
const NoMessageTypeID = 0xee6324c4

// construct implements constructor of AbstractMessageClass.
func (n NoMessage) construct() AbstractMessageClass { return &n }

// Ensuring interfaces in compile-time for NoMessage.
var (
	_ bin.Encoder     = &NoMessage{}
	_ bin.Decoder     = &NoMessage{}
	_ bin.BareEncoder = &NoMessage{}
	_ bin.BareDecoder = &NoMessage{}

	_ AbstractMessageClass = &NoMessage{}
)

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

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NoMessage) TypeID() uint32 {
	return NoMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*NoMessage) TypeName() string {
	return "noMessage"
}

// TypeInfo returns info about TL type.
func (n *NoMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "noMessage",
		ID:   NoMessageTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NoMessage) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode noMessage#ee6324c4 as nil")
	}
	b.PutID(NoMessageTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NoMessage) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode noMessage#ee6324c4 as nil")
	}
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
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NoMessage) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode noMessage#ee6324c4 to nil")
	}
	return nil
}

// TargetsMessage represents TL type `targetsMessage#cc6136f1`.
//
// See https://localhost:80/doc/constructor/targetsMessage for reference.
type TargetsMessage struct {
	// Targets field of TargetsMessage.
	Targets []int32
}

// TargetsMessageTypeID is TL type id of TargetsMessage.
const TargetsMessageTypeID = 0xcc6136f1

// construct implements constructor of AbstractMessageClass.
func (t TargetsMessage) construct() AbstractMessageClass { return &t }

// Ensuring interfaces in compile-time for TargetsMessage.
var (
	_ bin.Encoder     = &TargetsMessage{}
	_ bin.Decoder     = &TargetsMessage{}
	_ bin.BareEncoder = &TargetsMessage{}
	_ bin.BareDecoder = &TargetsMessage{}

	_ AbstractMessageClass = &TargetsMessage{}
)

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

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TargetsMessage) TypeID() uint32 {
	return TargetsMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*TargetsMessage) TypeName() string {
	return "targetsMessage"
}

// TypeInfo returns info about TL type.
func (t *TargetsMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "targetsMessage",
		ID:   TargetsMessageTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Targets",
			SchemaName: "targets",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TargetsMessage) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode targetsMessage#cc6136f1 as nil")
	}
	b.PutID(TargetsMessageTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TargetsMessage) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode targetsMessage#cc6136f1 as nil")
	}
	b.PutInt(len(t.Targets))
	for _, v := range t.Targets {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TargetsMessage) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode targetsMessage#cc6136f1 to nil")
	}
	if err := b.ConsumeID(TargetsMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode targetsMessage#cc6136f1: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TargetsMessage) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode targetsMessage#cc6136f1 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode targetsMessage#cc6136f1: field targets: %w", err)
		}

		if headerLen > 0 {
			t.Targets = make([]int32, 0, headerLen%bin.PreallocateLimit)
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

// GetTargets returns value of Targets field.
func (t *TargetsMessage) GetTargets() (value []int32) {
	if t == nil {
		return
	}
	return t.Targets
}

// FieldsMessage represents TL type `fieldsMessage#947225b5`.
//
// See https://localhost:80/doc/constructor/fieldsMessage for reference.
type FieldsMessage struct {
	// Flags field of FieldsMessage.
	Flags bin.Fields
	// Escape field of FieldsMessage.
	//
	// Use SetEscape and GetEscape helpers.
	Escape bool
	// TTLSeconds field of FieldsMessage.
	//
	// Use SetTTLSeconds and GetTTLSeconds helpers.
	TTLSeconds int
}

// FieldsMessageTypeID is TL type id of FieldsMessage.
const FieldsMessageTypeID = 0x947225b5

// construct implements constructor of AbstractMessageClass.
func (f FieldsMessage) construct() AbstractMessageClass { return &f }

// Ensuring interfaces in compile-time for FieldsMessage.
var (
	_ bin.Encoder     = &FieldsMessage{}
	_ bin.Decoder     = &FieldsMessage{}
	_ bin.BareEncoder = &FieldsMessage{}
	_ bin.BareDecoder = &FieldsMessage{}

	_ AbstractMessageClass = &FieldsMessage{}
)

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

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FieldsMessage) TypeID() uint32 {
	return FieldsMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*FieldsMessage) TypeName() string {
	return "fieldsMessage"
}

// TypeInfo returns info about TL type.
func (f *FieldsMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "fieldsMessage",
		ID:   FieldsMessageTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Escape",
			SchemaName: "escape",
			Null:       !f.Flags.Has(0),
		},
		{
			Name:       "TTLSeconds",
			SchemaName: "ttl_seconds",
			Null:       !f.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (f *FieldsMessage) SetFlags() {
	if !(f.Escape == false) {
		f.Flags.Set(0)
	}
	if !(f.TTLSeconds == 0) {
		f.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (f *FieldsMessage) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fieldsMessage#947225b5 as nil")
	}
	b.PutID(FieldsMessageTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FieldsMessage) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode fieldsMessage#947225b5 as nil")
	}
	f.SetFlags()
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

// Decode implements bin.Decoder.
func (f *FieldsMessage) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fieldsMessage#947225b5 to nil")
	}
	if err := b.ConsumeID(FieldsMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode fieldsMessage#947225b5: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FieldsMessage) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode fieldsMessage#947225b5 to nil")
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

// SetEscape sets value of Escape conditional field.
func (f *FieldsMessage) SetEscape(value bool) {
	f.Flags.Set(0)
	f.Escape = value
}

// GetEscape returns value of Escape conditional field and
// boolean which is true if field was set.
func (f *FieldsMessage) GetEscape() (value bool, ok bool) {
	if f == nil {
		return
	}
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
	if f == nil {
		return
	}
	if !f.Flags.Has(1) {
		return value, false
	}
	return f.TTLSeconds, true
}

// BytesMessage represents TL type `bytesMessage#f990a67d`.
//
// See https://localhost:80/doc/constructor/bytesMessage for reference.
type BytesMessage struct {
	// Data field of BytesMessage.
	Data []byte
}

// BytesMessageTypeID is TL type id of BytesMessage.
const BytesMessageTypeID = 0xf990a67d

// construct implements constructor of AbstractMessageClass.
func (b BytesMessage) construct() AbstractMessageClass { return &b }

// Ensuring interfaces in compile-time for BytesMessage.
var (
	_ bin.Encoder     = &BytesMessage{}
	_ bin.Decoder     = &BytesMessage{}
	_ bin.BareEncoder = &BytesMessage{}
	_ bin.BareDecoder = &BytesMessage{}

	_ AbstractMessageClass = &BytesMessage{}
)

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

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BytesMessage) TypeID() uint32 {
	return BytesMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*BytesMessage) TypeName() string {
	return "bytesMessage"
}

// TypeInfo returns info about TL type.
func (b *BytesMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bytesMessage",
		ID:   BytesMessageTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Data",
			SchemaName: "data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BytesMessage) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytesMessage#f990a67d as nil")
	}
	buf.PutID(BytesMessageTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BytesMessage) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytesMessage#f990a67d as nil")
	}
	buf.PutBytes(b.Data)
	return nil
}

// Decode implements bin.Decoder.
func (b *BytesMessage) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytesMessage#f990a67d to nil")
	}
	if err := buf.ConsumeID(BytesMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode bytesMessage#f990a67d: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BytesMessage) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytesMessage#f990a67d to nil")
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

// GetData returns value of Data field.
func (b *BytesMessage) GetData() (value []byte) {
	if b == nil {
		return
	}
	return b.Data
}

// AbstractMessageClassName is schema name of AbstractMessageClass.
const AbstractMessageClassName = "AbstractMessage"

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
	bin.BareEncoder
	bin.BareDecoder
	construct() AbstractMessageClass

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

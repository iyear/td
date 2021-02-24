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

// HelpPromoDataEmpty represents TL type `help.promoDataEmpty#98f6ac75`.
// No PSA/MTProxy info is available
//
// See https://core.telegram.org/constructor/help.promoDataEmpty for reference.
type HelpPromoDataEmpty struct {
	// Re-fetch PSA/MTProxy info after the specified number of seconds
	Expires int `tl:"expires"`
}

// HelpPromoDataEmptyTypeID is TL type id of HelpPromoDataEmpty.
const HelpPromoDataEmptyTypeID = 0x98f6ac75

func (p *HelpPromoDataEmpty) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Expires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPromoDataEmpty) String() string {
	if p == nil {
		return "HelpPromoDataEmpty(nil)"
	}
	type Alias HelpPromoDataEmpty
	return fmt.Sprintf("HelpPromoDataEmpty%+v", Alias(*p))
}

// FillFrom fills HelpPromoDataEmpty from given interface.
func (p *HelpPromoDataEmpty) FillFrom(from interface {
	GetExpires() (value int)
}) {
	p.Expires = from.GetExpires()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *HelpPromoDataEmpty) TypeID() uint32 {
	return HelpPromoDataEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (p *HelpPromoDataEmpty) TypeName() string {
	return "help.promoDataEmpty"
}

// Encode implements bin.Encoder.
func (p *HelpPromoDataEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.promoDataEmpty#98f6ac75 as nil")
	}
	b.PutID(HelpPromoDataEmptyTypeID)
	b.PutInt(p.Expires)
	return nil
}

// GetExpires returns value of Expires field.
func (p *HelpPromoDataEmpty) GetExpires() (value int) {
	return p.Expires
}

// Decode implements bin.Decoder.
func (p *HelpPromoDataEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.promoDataEmpty#98f6ac75 to nil")
	}
	if err := b.ConsumeID(HelpPromoDataEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.promoDataEmpty#98f6ac75: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoDataEmpty#98f6ac75: field expires: %w", err)
		}
		p.Expires = value
	}
	return nil
}

// construct implements constructor of HelpPromoDataClass.
func (p HelpPromoDataEmpty) construct() HelpPromoDataClass { return &p }

// Ensuring interfaces in compile-time for HelpPromoDataEmpty.
var (
	_ bin.Encoder = &HelpPromoDataEmpty{}
	_ bin.Decoder = &HelpPromoDataEmpty{}

	_ HelpPromoDataClass = &HelpPromoDataEmpty{}
)

// HelpPromoData represents TL type `help.promoData#8c39793f`.
// MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/constructor/help.promoData for reference.
type HelpPromoData struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// MTProxy-related channel
	Proxy bool `tl:"proxy"`
	// Expiry of PSA/MTProxy info
	Expires int `tl:"expires"`
	// MTProxy/PSA peer
	Peer PeerClass `tl:"peer"`
	// Chat info
	Chats []ChatClass `tl:"chats"`
	// User info
	Users []UserClass `tl:"users"`
	// PSA type
	//
	// Use SetPsaType and GetPsaType helpers.
	PsaType string `tl:"psa_type"`
	// PSA message
	//
	// Use SetPsaMessage and GetPsaMessage helpers.
	PsaMessage string `tl:"psa_message"`
}

// HelpPromoDataTypeID is TL type id of HelpPromoData.
const HelpPromoDataTypeID = 0x8c39793f

func (p *HelpPromoData) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Proxy == false) {
		return false
	}
	if !(p.Expires == 0) {
		return false
	}
	if !(p.Peer == nil) {
		return false
	}
	if !(p.Chats == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}
	if !(p.PsaType == "") {
		return false
	}
	if !(p.PsaMessage == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *HelpPromoData) String() string {
	if p == nil {
		return "HelpPromoData(nil)"
	}
	type Alias HelpPromoData
	return fmt.Sprintf("HelpPromoData%+v", Alias(*p))
}

// FillFrom fills HelpPromoData from given interface.
func (p *HelpPromoData) FillFrom(from interface {
	GetProxy() (value bool)
	GetExpires() (value int)
	GetPeer() (value PeerClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
	GetPsaType() (value string, ok bool)
	GetPsaMessage() (value string, ok bool)
}) {
	p.Proxy = from.GetProxy()
	p.Expires = from.GetExpires()
	p.Peer = from.GetPeer()
	p.Chats = from.GetChats()
	p.Users = from.GetUsers()
	if val, ok := from.GetPsaType(); ok {
		p.PsaType = val
	}

	if val, ok := from.GetPsaMessage(); ok {
		p.PsaMessage = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *HelpPromoData) TypeID() uint32 {
	return HelpPromoDataTypeID
}

// TypeName returns name of type in TL schema.
func (p *HelpPromoData) TypeName() string {
	return "help.promoData"
}

// Encode implements bin.Encoder.
func (p *HelpPromoData) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.promoData#8c39793f as nil")
	}
	b.PutID(HelpPromoDataTypeID)
	if !(p.Proxy == false) {
		p.Flags.Set(0)
	}
	if !(p.PsaType == "") {
		p.Flags.Set(1)
	}
	if !(p.PsaMessage == "") {
		p.Flags.Set(2)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.promoData#8c39793f: field flags: %w", err)
	}
	b.PutInt(p.Expires)
	if p.Peer == nil {
		return fmt.Errorf("unable to encode help.promoData#8c39793f: field peer is nil")
	}
	if err := p.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.promoData#8c39793f: field peer: %w", err)
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field users element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(1) {
		b.PutString(p.PsaType)
	}
	if p.Flags.Has(2) {
		b.PutString(p.PsaMessage)
	}
	return nil
}

// SetProxy sets value of Proxy conditional field.
func (p *HelpPromoData) SetProxy(value bool) {
	if value {
		p.Flags.Set(0)
		p.Proxy = true
	} else {
		p.Flags.Unset(0)
		p.Proxy = false
	}
}

// GetProxy returns value of Proxy conditional field.
func (p *HelpPromoData) GetProxy() (value bool) {
	return p.Flags.Has(0)
}

// GetExpires returns value of Expires field.
func (p *HelpPromoData) GetExpires() (value int) {
	return p.Expires
}

// GetPeer returns value of Peer field.
func (p *HelpPromoData) GetPeer() (value PeerClass) {
	return p.Peer
}

// GetChats returns value of Chats field.
func (p *HelpPromoData) GetChats() (value []ChatClass) {
	return p.Chats
}

// MapChats returns field Chats wrapped in ChatClassSlice helper.
func (p *HelpPromoData) MapChats() (value ChatClassSlice) {
	return ChatClassSlice(p.Chats)
}

// GetUsers returns value of Users field.
func (p *HelpPromoData) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (p *HelpPromoData) MapUsers() (value UserClassSlice) {
	return UserClassSlice(p.Users)
}

// SetPsaType sets value of PsaType conditional field.
func (p *HelpPromoData) SetPsaType(value string) {
	p.Flags.Set(1)
	p.PsaType = value
}

// GetPsaType returns value of PsaType conditional field and
// boolean which is true if field was set.
func (p *HelpPromoData) GetPsaType() (value string, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.PsaType, true
}

// SetPsaMessage sets value of PsaMessage conditional field.
func (p *HelpPromoData) SetPsaMessage(value string) {
	p.Flags.Set(2)
	p.PsaMessage = value
}

// GetPsaMessage returns value of PsaMessage conditional field and
// boolean which is true if field was set.
func (p *HelpPromoData) GetPsaMessage() (value string, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.PsaMessage, true
}

// Decode implements bin.Decoder.
func (p *HelpPromoData) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.promoData#8c39793f to nil")
	}
	if err := b.ConsumeID(HelpPromoDataTypeID); err != nil {
		return fmt.Errorf("unable to decode help.promoData#8c39793f: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field flags: %w", err)
		}
	}
	p.Proxy = p.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field expires: %w", err)
		}
		p.Expires = value
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field peer: %w", err)
		}
		p.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.promoData#8c39793f: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.promoData#8c39793f: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	if p.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field psa_type: %w", err)
		}
		p.PsaType = value
	}
	if p.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field psa_message: %w", err)
		}
		p.PsaMessage = value
	}
	return nil
}

// construct implements constructor of HelpPromoDataClass.
func (p HelpPromoData) construct() HelpPromoDataClass { return &p }

// Ensuring interfaces in compile-time for HelpPromoData.
var (
	_ bin.Encoder = &HelpPromoData{}
	_ bin.Decoder = &HelpPromoData{}

	_ HelpPromoDataClass = &HelpPromoData{}
)

// HelpPromoDataClass represents help.PromoData generic type.
//
// See https://core.telegram.org/type/help.PromoData for reference.
//
// Example:
//  g, err := tg.DecodeHelpPromoData(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpPromoDataEmpty: // help.promoDataEmpty#98f6ac75
//  case *tg.HelpPromoData: // help.promoData#8c39793f
//  default: panic(v)
//  }
type HelpPromoDataClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpPromoDataClass

	// TypeID returns type id in TL schema.
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Re-fetch PSA/MTProxy info after the specified number of seconds
	GetExpires() (value int)
	// AsNotEmpty tries to map HelpPromoDataClass to HelpPromoData.
	AsNotEmpty() (*HelpPromoData, bool)
}

// AsNotEmpty tries to map HelpPromoDataEmpty to HelpPromoData.
func (p *HelpPromoDataEmpty) AsNotEmpty() (*HelpPromoData, bool) {
	return nil, false
}

// AsNotEmpty tries to map HelpPromoData to HelpPromoData.
func (p *HelpPromoData) AsNotEmpty() (*HelpPromoData, bool) {
	return p, true
}

// DecodeHelpPromoData implements binary de-serialization for HelpPromoDataClass.
func DecodeHelpPromoData(buf *bin.Buffer) (HelpPromoDataClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpPromoDataEmptyTypeID:
		// Decoding help.promoDataEmpty#98f6ac75.
		v := HelpPromoDataEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPromoDataClass: %w", err)
		}
		return &v, nil
	case HelpPromoDataTypeID:
		// Decoding help.promoData#8c39793f.
		v := HelpPromoData{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPromoDataClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpPromoDataClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpPromoData boxes the HelpPromoDataClass providing a helper.
type HelpPromoDataBox struct {
	PromoData HelpPromoDataClass
}

// Decode implements bin.Decoder for HelpPromoDataBox.
func (b *HelpPromoDataBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpPromoDataBox to nil")
	}
	v, err := DecodeHelpPromoData(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PromoData = v
	return nil
}

// Encode implements bin.Encode for HelpPromoDataBox.
func (b *HelpPromoDataBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PromoData == nil {
		return fmt.Errorf("unable to encode HelpPromoDataClass as nil")
	}
	return b.PromoData.Encode(buf)
}

// HelpPromoDataClassSlice is adapter for slice of HelpPromoDataClass.
type HelpPromoDataClassSlice []HelpPromoDataClass

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s HelpPromoDataClassSlice) AppendOnlyNotEmpty(to []*HelpPromoData) []*HelpPromoData {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s HelpPromoDataClassSlice) AsNotEmpty() (to []*HelpPromoData) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s HelpPromoDataClassSlice) FirstAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s HelpPromoDataClassSlice) LastAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *HelpPromoDataClassSlice) PopFirstAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *HelpPromoDataClassSlice) PopAsNotEmpty() (v *HelpPromoData, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// First returns first element of slice (if exists).
func (s HelpPromoDataClassSlice) First() (v HelpPromoDataClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpPromoDataClassSlice) Last() (v HelpPromoDataClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpPromoDataClassSlice) PopFirst() (v HelpPromoDataClass, ok bool) {
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
func (s *HelpPromoDataClassSlice) Pop() (v HelpPromoDataClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

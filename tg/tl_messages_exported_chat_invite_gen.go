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

// MessagesExportedChatInvite represents TL type `messages.exportedChatInvite#1871be50`.
//
// See https://core.telegram.org/constructor/messages.exportedChatInvite for reference.
type MessagesExportedChatInvite struct {
	// Invite field of MessagesExportedChatInvite.
	Invite ChatInviteExported `tl:"invite"`
	// Users field of MessagesExportedChatInvite.
	Users []UserClass `tl:"users"`
}

// MessagesExportedChatInviteTypeID is TL type id of MessagesExportedChatInvite.
const MessagesExportedChatInviteTypeID = 0x1871be50

func (e *MessagesExportedChatInvite) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Invite.Zero()) {
		return false
	}
	if !(e.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesExportedChatInvite) String() string {
	if e == nil {
		return "MessagesExportedChatInvite(nil)"
	}
	type Alias MessagesExportedChatInvite
	return fmt.Sprintf("MessagesExportedChatInvite%+v", Alias(*e))
}

// FillFrom fills MessagesExportedChatInvite from given interface.
func (e *MessagesExportedChatInvite) FillFrom(from interface {
	GetInvite() (value ChatInviteExported)
	GetUsers() (value []UserClass)
}) {
	e.Invite = from.GetInvite()
	e.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *MessagesExportedChatInvite) TypeID() uint32 {
	return MessagesExportedChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (e *MessagesExportedChatInvite) TypeName() string {
	return "messages.exportedChatInvite"
}

// Encode implements bin.Encoder.
func (e *MessagesExportedChatInvite) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportedChatInvite#1871be50 as nil")
	}
	b.PutID(MessagesExportedChatInviteTypeID)
	if err := e.Invite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportedChatInvite#1871be50: field invite: %w", err)
	}
	b.PutVectorHeader(len(e.Users))
	for idx, v := range e.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.exportedChatInvite#1871be50: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.exportedChatInvite#1871be50: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetInvite returns value of Invite field.
func (e *MessagesExportedChatInvite) GetInvite() (value ChatInviteExported) {
	return e.Invite
}

// GetUsers returns value of Users field.
func (e *MessagesExportedChatInvite) GetUsers() (value []UserClass) {
	return e.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (e *MessagesExportedChatInvite) MapUsers() (value UserClassSlice) {
	return UserClassSlice(e.Users)
}

// Decode implements bin.Decoder.
func (e *MessagesExportedChatInvite) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportedChatInvite#1871be50 to nil")
	}
	if err := b.ConsumeID(MessagesExportedChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportedChatInvite#1871be50: %w", err)
	}
	{
		if err := e.Invite.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInvite#1871be50: field invite: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInvite#1871be50: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.exportedChatInvite#1871be50: field users: %w", err)
			}
			e.Users = append(e.Users, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesExportedChatInviteClass.
func (e MessagesExportedChatInvite) construct() MessagesExportedChatInviteClass { return &e }

// Ensuring interfaces in compile-time for MessagesExportedChatInvite.
var (
	_ bin.Encoder = &MessagesExportedChatInvite{}
	_ bin.Decoder = &MessagesExportedChatInvite{}

	_ MessagesExportedChatInviteClass = &MessagesExportedChatInvite{}
)

// MessagesExportedChatInviteReplaced represents TL type `messages.exportedChatInviteReplaced#222600ef`.
//
// See https://core.telegram.org/constructor/messages.exportedChatInviteReplaced for reference.
type MessagesExportedChatInviteReplaced struct {
	// Invite field of MessagesExportedChatInviteReplaced.
	Invite ChatInviteExported `tl:"invite"`
	// NewInvite field of MessagesExportedChatInviteReplaced.
	NewInvite ChatInviteExported `tl:"new_invite"`
	// Users field of MessagesExportedChatInviteReplaced.
	Users []UserClass `tl:"users"`
}

// MessagesExportedChatInviteReplacedTypeID is TL type id of MessagesExportedChatInviteReplaced.
const MessagesExportedChatInviteReplacedTypeID = 0x222600ef

func (e *MessagesExportedChatInviteReplaced) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Invite.Zero()) {
		return false
	}
	if !(e.NewInvite.Zero()) {
		return false
	}
	if !(e.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesExportedChatInviteReplaced) String() string {
	if e == nil {
		return "MessagesExportedChatInviteReplaced(nil)"
	}
	type Alias MessagesExportedChatInviteReplaced
	return fmt.Sprintf("MessagesExportedChatInviteReplaced%+v", Alias(*e))
}

// FillFrom fills MessagesExportedChatInviteReplaced from given interface.
func (e *MessagesExportedChatInviteReplaced) FillFrom(from interface {
	GetInvite() (value ChatInviteExported)
	GetNewInvite() (value ChatInviteExported)
	GetUsers() (value []UserClass)
}) {
	e.Invite = from.GetInvite()
	e.NewInvite = from.GetNewInvite()
	e.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *MessagesExportedChatInviteReplaced) TypeID() uint32 {
	return MessagesExportedChatInviteReplacedTypeID
}

// TypeName returns name of type in TL schema.
func (e *MessagesExportedChatInviteReplaced) TypeName() string {
	return "messages.exportedChatInviteReplaced"
}

// Encode implements bin.Encoder.
func (e *MessagesExportedChatInviteReplaced) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportedChatInviteReplaced#222600ef as nil")
	}
	b.PutID(MessagesExportedChatInviteReplacedTypeID)
	if err := e.Invite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportedChatInviteReplaced#222600ef: field invite: %w", err)
	}
	if err := e.NewInvite.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportedChatInviteReplaced#222600ef: field new_invite: %w", err)
	}
	b.PutVectorHeader(len(e.Users))
	for idx, v := range e.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.exportedChatInviteReplaced#222600ef: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.exportedChatInviteReplaced#222600ef: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetInvite returns value of Invite field.
func (e *MessagesExportedChatInviteReplaced) GetInvite() (value ChatInviteExported) {
	return e.Invite
}

// GetNewInvite returns value of NewInvite field.
func (e *MessagesExportedChatInviteReplaced) GetNewInvite() (value ChatInviteExported) {
	return e.NewInvite
}

// GetUsers returns value of Users field.
func (e *MessagesExportedChatInviteReplaced) GetUsers() (value []UserClass) {
	return e.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (e *MessagesExportedChatInviteReplaced) MapUsers() (value UserClassSlice) {
	return UserClassSlice(e.Users)
}

// Decode implements bin.Decoder.
func (e *MessagesExportedChatInviteReplaced) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportedChatInviteReplaced#222600ef to nil")
	}
	if err := b.ConsumeID(MessagesExportedChatInviteReplacedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportedChatInviteReplaced#222600ef: %w", err)
	}
	{
		if err := e.Invite.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInviteReplaced#222600ef: field invite: %w", err)
		}
	}
	{
		if err := e.NewInvite.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInviteReplaced#222600ef: field new_invite: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInviteReplaced#222600ef: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.exportedChatInviteReplaced#222600ef: field users: %w", err)
			}
			e.Users = append(e.Users, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesExportedChatInviteClass.
func (e MessagesExportedChatInviteReplaced) construct() MessagesExportedChatInviteClass { return &e }

// Ensuring interfaces in compile-time for MessagesExportedChatInviteReplaced.
var (
	_ bin.Encoder = &MessagesExportedChatInviteReplaced{}
	_ bin.Decoder = &MessagesExportedChatInviteReplaced{}

	_ MessagesExportedChatInviteClass = &MessagesExportedChatInviteReplaced{}
)

// MessagesExportedChatInviteClass represents messages.ExportedChatInvite generic type.
//
// See https://core.telegram.org/type/messages.ExportedChatInvite for reference.
//
// Example:
//  g, err := tg.DecodeMessagesExportedChatInvite(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesExportedChatInvite: // messages.exportedChatInvite#1871be50
//  case *tg.MessagesExportedChatInviteReplaced: // messages.exportedChatInviteReplaced#222600ef
//  default: panic(v)
//  }
type MessagesExportedChatInviteClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessagesExportedChatInviteClass

	// TypeID returns type id in TL schema.
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Invite field of MessagesExportedChatInvite.
	GetInvite() (value ChatInviteExported)
	// Users field of MessagesExportedChatInvite.
	GetUsers() (value []UserClass)
	// Users field of MessagesExportedChatInvite.
	MapUsers() (value UserClassSlice)
}

// DecodeMessagesExportedChatInvite implements binary de-serialization for MessagesExportedChatInviteClass.
func DecodeMessagesExportedChatInvite(buf *bin.Buffer) (MessagesExportedChatInviteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesExportedChatInviteTypeID:
		// Decoding messages.exportedChatInvite#1871be50.
		v := MessagesExportedChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesExportedChatInviteClass: %w", err)
		}
		return &v, nil
	case MessagesExportedChatInviteReplacedTypeID:
		// Decoding messages.exportedChatInviteReplaced#222600ef.
		v := MessagesExportedChatInviteReplaced{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesExportedChatInviteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesExportedChatInviteClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesExportedChatInvite boxes the MessagesExportedChatInviteClass providing a helper.
type MessagesExportedChatInviteBox struct {
	ExportedChatInvite MessagesExportedChatInviteClass
}

// Decode implements bin.Decoder for MessagesExportedChatInviteBox.
func (b *MessagesExportedChatInviteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesExportedChatInviteBox to nil")
	}
	v, err := DecodeMessagesExportedChatInvite(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ExportedChatInvite = v
	return nil
}

// Encode implements bin.Encode for MessagesExportedChatInviteBox.
func (b *MessagesExportedChatInviteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ExportedChatInvite == nil {
		return fmt.Errorf("unable to encode MessagesExportedChatInviteClass as nil")
	}
	return b.ExportedChatInvite.Encode(buf)
}

// MessagesExportedChatInviteClassSlice is adapter for slice of MessagesExportedChatInviteClass.
type MessagesExportedChatInviteClassSlice []MessagesExportedChatInviteClass

// First returns first element of slice (if exists).
func (s MessagesExportedChatInviteClassSlice) First() (v MessagesExportedChatInviteClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesExportedChatInviteClassSlice) Last() (v MessagesExportedChatInviteClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesExportedChatInviteClassSlice) PopFirst() (v MessagesExportedChatInviteClass, ok bool) {
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
func (s *MessagesExportedChatInviteClassSlice) Pop() (v MessagesExportedChatInviteClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

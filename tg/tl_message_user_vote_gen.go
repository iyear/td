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

// MessageUserVote represents TL type `messageUserVote#a28e5559`.
// How a user voted in a poll
//
// See https://core.telegram.org/constructor/messageUserVote for reference.
type MessageUserVote struct {
	// User ID
	UserID int
	// The option chosen by the user
	Option []byte
	// When did the user cast the vote
	Date int
}

// MessageUserVoteTypeID is TL type id of MessageUserVote.
const MessageUserVoteTypeID = 0xa28e5559

func (m *MessageUserVote) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.Option == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageUserVote) String() string {
	if m == nil {
		return "MessageUserVote(nil)"
	}
	type Alias MessageUserVote
	return fmt.Sprintf("MessageUserVote%+v", Alias(*m))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageUserVote) TypeID() uint32 {
	return MessageUserVoteTypeID
}

// Encode implements bin.Encoder.
func (m *MessageUserVote) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVote#a28e5559 as nil")
	}
	b.PutID(MessageUserVoteTypeID)
	b.PutInt(m.UserID)
	b.PutBytes(m.Option)
	b.PutInt(m.Date)
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageUserVote) GetUserID() (value int) {
	return m.UserID
}

// GetOption returns value of Option field.
func (m *MessageUserVote) GetOption() (value []byte) {
	return m.Option
}

// GetDate returns value of Date field.
func (m *MessageUserVote) GetDate() (value int) {
	return m.Date
}

// Decode implements bin.Decoder.
func (m *MessageUserVote) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVote#a28e5559 to nil")
	}
	if err := b.ConsumeID(MessageUserVoteTypeID); err != nil {
		return fmt.Errorf("unable to decode messageUserVote#a28e5559: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVote#a28e5559: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVote#a28e5559: field option: %w", err)
		}
		m.Option = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVote#a28e5559: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// construct implements constructor of MessageUserVoteClass.
func (m MessageUserVote) construct() MessageUserVoteClass { return &m }

// Ensuring interfaces in compile-time for MessageUserVote.
var (
	_ bin.Encoder = &MessageUserVote{}
	_ bin.Decoder = &MessageUserVote{}

	_ MessageUserVoteClass = &MessageUserVote{}
)

// MessageUserVoteInputOption represents TL type `messageUserVoteInputOption#36377430`.
// How a user voted in a poll (reduced constructor, returned if an option was provided to messages.getPollVotes¹)
//
// Links:
//  1) https://core.telegram.org/method/messages.getPollVotes
//
// See https://core.telegram.org/constructor/messageUserVoteInputOption for reference.
type MessageUserVoteInputOption struct {
	// The user that voted for the queried option
	UserID int
	// When did the user cast the vote
	Date int
}

// MessageUserVoteInputOptionTypeID is TL type id of MessageUserVoteInputOption.
const MessageUserVoteInputOptionTypeID = 0x36377430

func (m *MessageUserVoteInputOption) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageUserVoteInputOption) String() string {
	if m == nil {
		return "MessageUserVoteInputOption(nil)"
	}
	type Alias MessageUserVoteInputOption
	return fmt.Sprintf("MessageUserVoteInputOption%+v", Alias(*m))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageUserVoteInputOption) TypeID() uint32 {
	return MessageUserVoteInputOptionTypeID
}

// Encode implements bin.Encoder.
func (m *MessageUserVoteInputOption) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVoteInputOption#36377430 as nil")
	}
	b.PutID(MessageUserVoteInputOptionTypeID)
	b.PutInt(m.UserID)
	b.PutInt(m.Date)
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageUserVoteInputOption) GetUserID() (value int) {
	return m.UserID
}

// GetDate returns value of Date field.
func (m *MessageUserVoteInputOption) GetDate() (value int) {
	return m.Date
}

// Decode implements bin.Decoder.
func (m *MessageUserVoteInputOption) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVoteInputOption#36377430 to nil")
	}
	if err := b.ConsumeID(MessageUserVoteInputOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode messageUserVoteInputOption#36377430: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteInputOption#36377430: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteInputOption#36377430: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// construct implements constructor of MessageUserVoteClass.
func (m MessageUserVoteInputOption) construct() MessageUserVoteClass { return &m }

// Ensuring interfaces in compile-time for MessageUserVoteInputOption.
var (
	_ bin.Encoder = &MessageUserVoteInputOption{}
	_ bin.Decoder = &MessageUserVoteInputOption{}

	_ MessageUserVoteClass = &MessageUserVoteInputOption{}
)

// MessageUserVoteMultiple represents TL type `messageUserVoteMultiple#e8fe0de`.
// How a user voted in a multiple-choice poll
//
// See https://core.telegram.org/constructor/messageUserVoteMultiple for reference.
type MessageUserVoteMultiple struct {
	// User ID
	UserID int
	// Options chosen by the user
	Options [][]byte
	// When did the user cast their votes
	Date int
}

// MessageUserVoteMultipleTypeID is TL type id of MessageUserVoteMultiple.
const MessageUserVoteMultipleTypeID = 0xe8fe0de

func (m *MessageUserVoteMultiple) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.UserID == 0) {
		return false
	}
	if !(m.Options == nil) {
		return false
	}
	if !(m.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageUserVoteMultiple) String() string {
	if m == nil {
		return "MessageUserVoteMultiple(nil)"
	}
	type Alias MessageUserVoteMultiple
	return fmt.Sprintf("MessageUserVoteMultiple%+v", Alias(*m))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *MessageUserVoteMultiple) TypeID() uint32 {
	return MessageUserVoteMultipleTypeID
}

// Encode implements bin.Encoder.
func (m *MessageUserVoteMultiple) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageUserVoteMultiple#e8fe0de as nil")
	}
	b.PutID(MessageUserVoteMultipleTypeID)
	b.PutInt(m.UserID)
	b.PutVectorHeader(len(m.Options))
	for _, v := range m.Options {
		b.PutBytes(v)
	}
	b.PutInt(m.Date)
	return nil
}

// GetUserID returns value of UserID field.
func (m *MessageUserVoteMultiple) GetUserID() (value int) {
	return m.UserID
}

// GetOptions returns value of Options field.
func (m *MessageUserVoteMultiple) GetOptions() (value [][]byte) {
	return m.Options
}

// GetDate returns value of Date field.
func (m *MessageUserVoteMultiple) GetDate() (value int) {
	return m.Date
}

// Decode implements bin.Decoder.
func (m *MessageUserVoteMultiple) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageUserVoteMultiple#e8fe0de to nil")
	}
	if err := b.ConsumeID(MessageUserVoteMultipleTypeID); err != nil {
		return fmt.Errorf("unable to decode messageUserVoteMultiple#e8fe0de: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteMultiple#e8fe0de: field user_id: %w", err)
		}
		m.UserID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteMultiple#e8fe0de: field options: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode messageUserVoteMultiple#e8fe0de: field options: %w", err)
			}
			m.Options = append(m.Options, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageUserVoteMultiple#e8fe0de: field date: %w", err)
		}
		m.Date = value
	}
	return nil
}

// construct implements constructor of MessageUserVoteClass.
func (m MessageUserVoteMultiple) construct() MessageUserVoteClass { return &m }

// Ensuring interfaces in compile-time for MessageUserVoteMultiple.
var (
	_ bin.Encoder = &MessageUserVoteMultiple{}
	_ bin.Decoder = &MessageUserVoteMultiple{}

	_ MessageUserVoteClass = &MessageUserVoteMultiple{}
)

// MessageUserVoteClass represents MessageUserVote generic type.
//
// See https://core.telegram.org/type/MessageUserVote for reference.
//
// Example:
//  g, err := tg.DecodeMessageUserVote(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessageUserVote: // messageUserVote#a28e5559
//  case *tg.MessageUserVoteInputOption: // messageUserVoteInputOption#36377430
//  case *tg.MessageUserVoteMultiple: // messageUserVoteMultiple#e8fe0de
//  default: panic(v)
//  }
type MessageUserVoteClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessageUserVoteClass

	// User ID
	GetUserID() (value int)
	// When did the user cast the vote
	GetDate() (value int)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeMessageUserVote implements binary de-serialization for MessageUserVoteClass.
func DecodeMessageUserVote(buf *bin.Buffer) (MessageUserVoteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageUserVoteTypeID:
		// Decoding messageUserVote#a28e5559.
		v := MessageUserVote{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", err)
		}
		return &v, nil
	case MessageUserVoteInputOptionTypeID:
		// Decoding messageUserVoteInputOption#36377430.
		v := MessageUserVoteInputOption{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", err)
		}
		return &v, nil
	case MessageUserVoteMultipleTypeID:
		// Decoding messageUserVoteMultiple#e8fe0de.
		v := MessageUserVoteMultiple{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageUserVoteClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessageUserVote boxes the MessageUserVoteClass providing a helper.
type MessageUserVoteBox struct {
	MessageUserVote MessageUserVoteClass
}

// Decode implements bin.Decoder for MessageUserVoteBox.
func (b *MessageUserVoteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageUserVoteBox to nil")
	}
	v, err := DecodeMessageUserVote(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageUserVote = v
	return nil
}

// Encode implements bin.Encode for MessageUserVoteBox.
func (b *MessageUserVoteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageUserVote == nil {
		return fmt.Errorf("unable to encode MessageUserVoteClass as nil")
	}
	return b.MessageUserVote.Encode(buf)
}

// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// NewSessionCreated represents TL type `new_session_created#9ec20908`.
type NewSessionCreated struct {
	// FirstMsgID field of NewSessionCreated.
	FirstMsgID int64 `tl:"first_msg_id"`
	// UniqueID field of NewSessionCreated.
	UniqueID int64 `tl:"unique_id"`
	// ServerSalt field of NewSessionCreated.
	ServerSalt int64 `tl:"server_salt"`
}

// NewSessionCreatedTypeID is TL type id of NewSessionCreated.
const NewSessionCreatedTypeID = 0x9ec20908

func (n *NewSessionCreated) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.FirstMsgID == 0) {
		return false
	}
	if !(n.UniqueID == 0) {
		return false
	}
	if !(n.ServerSalt == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NewSessionCreated) String() string {
	if n == nil {
		return "NewSessionCreated(nil)"
	}
	type Alias NewSessionCreated
	return fmt.Sprintf("NewSessionCreated%+v", Alias(*n))
}

// FillFrom fills NewSessionCreated from given interface.
func (n *NewSessionCreated) FillFrom(from interface {
	GetFirstMsgID() (value int64)
	GetUniqueID() (value int64)
	GetServerSalt() (value int64)
}) {
	n.FirstMsgID = from.GetFirstMsgID()
	n.UniqueID = from.GetUniqueID()
	n.ServerSalt = from.GetServerSalt()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *NewSessionCreated) TypeID() uint32 {
	return NewSessionCreatedTypeID
}

// TypeName returns name of type in TL schema.
func (n *NewSessionCreated) TypeName() string {
	return "new_session_created"
}

// Encode implements bin.Encoder.
func (n *NewSessionCreated) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode new_session_created#9ec20908 as nil")
	}
	b.PutID(NewSessionCreatedTypeID)
	b.PutLong(n.FirstMsgID)
	b.PutLong(n.UniqueID)
	b.PutLong(n.ServerSalt)
	return nil
}

// GetFirstMsgID returns value of FirstMsgID field.
func (n *NewSessionCreated) GetFirstMsgID() (value int64) {
	return n.FirstMsgID
}

// GetUniqueID returns value of UniqueID field.
func (n *NewSessionCreated) GetUniqueID() (value int64) {
	return n.UniqueID
}

// GetServerSalt returns value of ServerSalt field.
func (n *NewSessionCreated) GetServerSalt() (value int64) {
	return n.ServerSalt
}

// Decode implements bin.Decoder.
func (n *NewSessionCreated) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode new_session_created#9ec20908 to nil")
	}
	if err := b.ConsumeID(NewSessionCreatedTypeID); err != nil {
		return fmt.Errorf("unable to decode new_session_created#9ec20908: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode new_session_created#9ec20908: field first_msg_id: %w", err)
		}
		n.FirstMsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode new_session_created#9ec20908: field unique_id: %w", err)
		}
		n.UniqueID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode new_session_created#9ec20908: field server_salt: %w", err)
		}
		n.ServerSalt = value
	}
	return nil
}

// Ensuring interfaces in compile-time for NewSessionCreated.
var (
	_ bin.Encoder = &NewSessionCreated{}
	_ bin.Decoder = &NewSessionCreated{}
)

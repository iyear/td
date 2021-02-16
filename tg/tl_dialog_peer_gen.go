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

// DialogPeer represents TL type `dialogPeer#e56dbf05`.
// Peer
//
// See https://core.telegram.org/constructor/dialogPeer for reference.
type DialogPeer struct {
	// Peer
	Peer PeerClass
}

// DialogPeerTypeID is TL type id of DialogPeer.
const DialogPeerTypeID = 0xe56dbf05

func (d *DialogPeer) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogPeer) String() string {
	if d == nil {
		return "DialogPeer(nil)"
	}
	type Alias DialogPeer
	return fmt.Sprintf("DialogPeer%+v", Alias(*d))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DialogPeer) TypeID() uint32 {
	return DialogPeerTypeID
}

// Encode implements bin.Encoder.
func (d *DialogPeer) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogPeer#e56dbf05 as nil")
	}
	b.PutID(DialogPeerTypeID)
	if d.Peer == nil {
		return fmt.Errorf("unable to encode dialogPeer#e56dbf05: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogPeer#e56dbf05: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *DialogPeer) GetPeer() (value PeerClass) {
	return d.Peer
}

// Decode implements bin.Decoder.
func (d *DialogPeer) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogPeer#e56dbf05 to nil")
	}
	if err := b.ConsumeID(DialogPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogPeer#e56dbf05: %w", err)
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode dialogPeer#e56dbf05: field peer: %w", err)
		}
		d.Peer = value
	}
	return nil
}

// construct implements constructor of DialogPeerClass.
func (d DialogPeer) construct() DialogPeerClass { return &d }

// Ensuring interfaces in compile-time for DialogPeer.
var (
	_ bin.Encoder = &DialogPeer{}
	_ bin.Decoder = &DialogPeer{}

	_ DialogPeerClass = &DialogPeer{}
)

// DialogPeerFolder represents TL type `dialogPeerFolder#514519e2`.
// Peer folder¹
//
// Links:
//  1) https://core.telegram.org/api/folders#peer-folders
//
// See https://core.telegram.org/constructor/dialogPeerFolder for reference.
type DialogPeerFolder struct {
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// DialogPeerFolderTypeID is TL type id of DialogPeerFolder.
const DialogPeerFolderTypeID = 0x514519e2

func (d *DialogPeerFolder) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogPeerFolder) String() string {
	if d == nil {
		return "DialogPeerFolder(nil)"
	}
	type Alias DialogPeerFolder
	return fmt.Sprintf("DialogPeerFolder%+v", Alias(*d))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *DialogPeerFolder) TypeID() uint32 {
	return DialogPeerFolderTypeID
}

// Encode implements bin.Encoder.
func (d *DialogPeerFolder) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogPeerFolder#514519e2 as nil")
	}
	b.PutID(DialogPeerFolderTypeID)
	b.PutInt(d.FolderID)
	return nil
}

// GetFolderID returns value of FolderID field.
func (d *DialogPeerFolder) GetFolderID() (value int) {
	return d.FolderID
}

// Decode implements bin.Decoder.
func (d *DialogPeerFolder) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogPeerFolder#514519e2 to nil")
	}
	if err := b.ConsumeID(DialogPeerFolderTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogPeerFolder#514519e2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogPeerFolder#514519e2: field folder_id: %w", err)
		}
		d.FolderID = value
	}
	return nil
}

// construct implements constructor of DialogPeerClass.
func (d DialogPeerFolder) construct() DialogPeerClass { return &d }

// Ensuring interfaces in compile-time for DialogPeerFolder.
var (
	_ bin.Encoder = &DialogPeerFolder{}
	_ bin.Decoder = &DialogPeerFolder{}

	_ DialogPeerClass = &DialogPeerFolder{}
)

// DialogPeerClass represents DialogPeer generic type.
//
// See https://core.telegram.org/type/DialogPeer for reference.
//
// Example:
//  g, err := tg.DecodeDialogPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.DialogPeer: // dialogPeer#e56dbf05
//  case *tg.DialogPeerFolder: // dialogPeerFolder#514519e2
//  default: panic(v)
//  }
type DialogPeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() DialogPeerClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeDialogPeer implements binary de-serialization for DialogPeerClass.
func DecodeDialogPeer(buf *bin.Buffer) (DialogPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DialogPeerTypeID:
		// Decoding dialogPeer#e56dbf05.
		v := DialogPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogPeerClass: %w", err)
		}
		return &v, nil
	case DialogPeerFolderTypeID:
		// Decoding dialogPeerFolder#514519e2.
		v := DialogPeerFolder{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// DialogPeer boxes the DialogPeerClass providing a helper.
type DialogPeerBox struct {
	DialogPeer DialogPeerClass
}

// Decode implements bin.Decoder for DialogPeerBox.
func (b *DialogPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DialogPeerBox to nil")
	}
	v, err := DecodeDialogPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DialogPeer = v
	return nil
}

// Encode implements bin.Encode for DialogPeerBox.
func (b *DialogPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DialogPeer == nil {
		return fmt.Errorf("unable to encode DialogPeerClass as nil")
	}
	return b.DialogPeer.Encode(buf)
}

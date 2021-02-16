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

// PeerBlocked represents TL type `peerBlocked#e8fd8014`.
// Information about a blocked peer
//
// See https://core.telegram.org/constructor/peerBlocked for reference.
type PeerBlocked struct {
	// Peer ID
	PeerID PeerClass
	// When was the peer blocked
	Date int
}

// PeerBlockedTypeID is TL type id of PeerBlocked.
const PeerBlockedTypeID = 0xe8fd8014

func (p *PeerBlocked) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PeerID == nil) {
		return false
	}
	if !(p.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerBlocked) String() string {
	if p == nil {
		return "PeerBlocked(nil)"
	}
	type Alias PeerBlocked
	return fmt.Sprintf("PeerBlocked%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PeerBlocked) TypeID() uint32 {
	return PeerBlockedTypeID
}

// Encode implements bin.Encoder.
func (p *PeerBlocked) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerBlocked#e8fd8014 as nil")
	}
	b.PutID(PeerBlockedTypeID)
	if p.PeerID == nil {
		return fmt.Errorf("unable to encode peerBlocked#e8fd8014: field peer_id is nil")
	}
	if err := p.PeerID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode peerBlocked#e8fd8014: field peer_id: %w", err)
	}
	b.PutInt(p.Date)
	return nil
}

// GetPeerID returns value of PeerID field.
func (p *PeerBlocked) GetPeerID() (value PeerClass) {
	return p.PeerID
}

// GetDate returns value of Date field.
func (p *PeerBlocked) GetDate() (value int) {
	return p.Date
}

// Decode implements bin.Decoder.
func (p *PeerBlocked) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerBlocked#e8fd8014 to nil")
	}
	if err := b.ConsumeID(PeerBlockedTypeID); err != nil {
		return fmt.Errorf("unable to decode peerBlocked#e8fd8014: %w", err)
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode peerBlocked#e8fd8014: field peer_id: %w", err)
		}
		p.PeerID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerBlocked#e8fd8014: field date: %w", err)
		}
		p.Date = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PeerBlocked.
var (
	_ bin.Encoder = &PeerBlocked{}
	_ bin.Decoder = &PeerBlocked{}
)

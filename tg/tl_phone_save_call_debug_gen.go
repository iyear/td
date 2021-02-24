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

// PhoneSaveCallDebugRequest represents TL type `phone.saveCallDebug#277add7e`.
// Send phone call debug data to server
//
// See https://core.telegram.org/method/phone.saveCallDebug for reference.
type PhoneSaveCallDebugRequest struct {
	// Phone call
	Peer InputPhoneCall `tl:"peer"`
	// Debug statistics obtained from libtgvoip
	Debug DataJSON `tl:"debug"`
}

// PhoneSaveCallDebugRequestTypeID is TL type id of PhoneSaveCallDebugRequest.
const PhoneSaveCallDebugRequestTypeID = 0x277add7e

func (s *PhoneSaveCallDebugRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Debug.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneSaveCallDebugRequest) String() string {
	if s == nil {
		return "PhoneSaveCallDebugRequest(nil)"
	}
	type Alias PhoneSaveCallDebugRequest
	return fmt.Sprintf("PhoneSaveCallDebugRequest%+v", Alias(*s))
}

// FillFrom fills PhoneSaveCallDebugRequest from given interface.
func (s *PhoneSaveCallDebugRequest) FillFrom(from interface {
	GetPeer() (value InputPhoneCall)
	GetDebug() (value DataJSON)
}) {
	s.Peer = from.GetPeer()
	s.Debug = from.GetDebug()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *PhoneSaveCallDebugRequest) TypeID() uint32 {
	return PhoneSaveCallDebugRequestTypeID
}

// TypeName returns name of type in TL schema.
func (s *PhoneSaveCallDebugRequest) TypeName() string {
	return "phone.saveCallDebug"
}

// Encode implements bin.Encoder.
func (s *PhoneSaveCallDebugRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.saveCallDebug#277add7e as nil")
	}
	b.PutID(PhoneSaveCallDebugRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallDebug#277add7e: field peer: %w", err)
	}
	if err := s.Debug.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.saveCallDebug#277add7e: field debug: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *PhoneSaveCallDebugRequest) GetPeer() (value InputPhoneCall) {
	return s.Peer
}

// GetDebug returns value of Debug field.
func (s *PhoneSaveCallDebugRequest) GetDebug() (value DataJSON) {
	return s.Debug
}

// Decode implements bin.Decoder.
func (s *PhoneSaveCallDebugRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.saveCallDebug#277add7e to nil")
	}
	if err := b.ConsumeID(PhoneSaveCallDebugRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: field peer: %w", err)
		}
	}
	{
		if err := s.Debug.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.saveCallDebug#277add7e: field debug: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneSaveCallDebugRequest.
var (
	_ bin.Encoder = &PhoneSaveCallDebugRequest{}
	_ bin.Decoder = &PhoneSaveCallDebugRequest{}
)

// PhoneSaveCallDebug invokes method phone.saveCallDebug#277add7e returning error if any.
// Send phone call debug data to server
//
// Possible errors:
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//  400 DATA_JSON_INVALID: The provided JSON data is invalid
//
// See https://core.telegram.org/method/phone.saveCallDebug for reference.
func (c *Client) PhoneSaveCallDebug(ctx context.Context, request *PhoneSaveCallDebugRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}

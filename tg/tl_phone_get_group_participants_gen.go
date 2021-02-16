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

// PhoneGetGroupParticipantsRequest represents TL type `phone.getGroupParticipants#c9f1d285`.
//
// See https://core.telegram.org/method/phone.getGroupParticipants for reference.
type PhoneGetGroupParticipantsRequest struct {
	// Call field of PhoneGetGroupParticipantsRequest.
	Call InputGroupCall
	// Ids field of PhoneGetGroupParticipantsRequest.
	Ids []int
	// Sources field of PhoneGetGroupParticipantsRequest.
	Sources []int
	// Offset field of PhoneGetGroupParticipantsRequest.
	Offset string
	// Limit field of PhoneGetGroupParticipantsRequest.
	Limit int
}

// PhoneGetGroupParticipantsRequestTypeID is TL type id of PhoneGetGroupParticipantsRequest.
const PhoneGetGroupParticipantsRequestTypeID = 0xc9f1d285

func (g *PhoneGetGroupParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Call.Zero()) {
		return false
	}
	if !(g.Ids == nil) {
		return false
	}
	if !(g.Sources == nil) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *PhoneGetGroupParticipantsRequest) String() string {
	if g == nil {
		return "PhoneGetGroupParticipantsRequest(nil)"
	}
	type Alias PhoneGetGroupParticipantsRequest
	return fmt.Sprintf("PhoneGetGroupParticipantsRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *PhoneGetGroupParticipantsRequest) TypeID() uint32 {
	return PhoneGetGroupParticipantsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *PhoneGetGroupParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode phone.getGroupParticipants#c9f1d285 as nil")
	}
	b.PutID(PhoneGetGroupParticipantsRequestTypeID)
	if err := g.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.getGroupParticipants#c9f1d285: field call: %w", err)
	}
	b.PutVectorHeader(len(g.Ids))
	for _, v := range g.Ids {
		b.PutInt(v)
	}
	b.PutVectorHeader(len(g.Sources))
	for _, v := range g.Sources {
		b.PutInt(v)
	}
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// GetCall returns value of Call field.
func (g *PhoneGetGroupParticipantsRequest) GetCall() (value InputGroupCall) {
	return g.Call
}

// GetIds returns value of Ids field.
func (g *PhoneGetGroupParticipantsRequest) GetIds() (value []int) {
	return g.Ids
}

// GetSources returns value of Sources field.
func (g *PhoneGetGroupParticipantsRequest) GetSources() (value []int) {
	return g.Sources
}

// GetOffset returns value of Offset field.
func (g *PhoneGetGroupParticipantsRequest) GetOffset() (value string) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *PhoneGetGroupParticipantsRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *PhoneGetGroupParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode phone.getGroupParticipants#c9f1d285 to nil")
	}
	if err := b.ConsumeID(PhoneGetGroupParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: %w", err)
	}
	{
		if err := g.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field call: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field ids: %w", err)
			}
			g.Ids = append(g.Ids, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field sources: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field sources: %w", err)
			}
			g.Sources = append(g.Sources, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.getGroupParticipants#c9f1d285: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneGetGroupParticipantsRequest.
var (
	_ bin.Encoder = &PhoneGetGroupParticipantsRequest{}
	_ bin.Decoder = &PhoneGetGroupParticipantsRequest{}
)

// PhoneGetGroupParticipants invokes method phone.getGroupParticipants#c9f1d285 returning error if any.
//
// See https://core.telegram.org/method/phone.getGroupParticipants for reference.
func (c *Client) PhoneGetGroupParticipants(ctx context.Context, request *PhoneGetGroupParticipantsRequest) (*PhoneGroupParticipants, error) {
	var result PhoneGroupParticipants

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

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

// MessagesGetSearchCountersRequest represents TL type `messages.getSearchCounters#732eef00`.
// Get the number of results that would be found by a messages.search¹ call with the same parameters
//
// Links:
//  1) https://core.telegram.org/method/messages.search
//
// See https://core.telegram.org/method/messages.getSearchCounters for reference.
type MessagesGetSearchCountersRequest struct {
	// Peer where to search
	Peer InputPeerClass `tl:"peer"`
	// Search filters
	Filters []MessagesFilterClass `tl:"filters"`
}

// MessagesGetSearchCountersRequestTypeID is TL type id of MessagesGetSearchCountersRequest.
const MessagesGetSearchCountersRequestTypeID = 0x732eef00

func (g *MessagesGetSearchCountersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Filters == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetSearchCountersRequest) String() string {
	if g == nil {
		return "MessagesGetSearchCountersRequest(nil)"
	}
	type Alias MessagesGetSearchCountersRequest
	return fmt.Sprintf("MessagesGetSearchCountersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetSearchCountersRequest from given interface.
func (g *MessagesGetSearchCountersRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetFilters() (value []MessagesFilterClass)
}) {
	g.Peer = from.GetPeer()
	g.Filters = from.GetFilters()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetSearchCountersRequest) TypeID() uint32 {
	return MessagesGetSearchCountersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *MessagesGetSearchCountersRequest) TypeName() string {
	return "messages.getSearchCounters"
}

// Encode implements bin.Encoder.
func (g *MessagesGetSearchCountersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getSearchCounters#732eef00 as nil")
	}
	b.PutID(MessagesGetSearchCountersRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.Filters))
	for idx, v := range g.Filters {
		if v == nil {
			return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field filters element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.getSearchCounters#732eef00: field filters element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetSearchCountersRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetFilters returns value of Filters field.
func (g *MessagesGetSearchCountersRequest) GetFilters() (value []MessagesFilterClass) {
	return g.Filters
}

// MapFilters returns field Filters wrapped in MessagesFilterClassSlice helper.
func (g *MessagesGetSearchCountersRequest) MapFilters() (value MessagesFilterClassSlice) {
	return MessagesFilterClassSlice(g.Filters)
}

// Decode implements bin.Decoder.
func (g *MessagesGetSearchCountersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getSearchCounters#732eef00 to nil")
	}
	if err := b.ConsumeID(MessagesGetSearchCountersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: field filters: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.getSearchCounters#732eef00: field filters: %w", err)
			}
			g.Filters = append(g.Filters, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetSearchCountersRequest.
var (
	_ bin.Encoder = &MessagesGetSearchCountersRequest{}
	_ bin.Decoder = &MessagesGetSearchCountersRequest{}
)

// MessagesGetSearchCounters invokes method messages.getSearchCounters#732eef00 returning error if any.
// Get the number of results that would be found by a messages.search¹ call with the same parameters
//
// Links:
//  1) https://core.telegram.org/method/messages.search
//
// See https://core.telegram.org/method/messages.getSearchCounters for reference.
func (c *Client) MessagesGetSearchCounters(ctx context.Context, request *MessagesGetSearchCountersRequest) ([]MessagesSearchCounter, error) {
	var result MessagesSearchCounterVector

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}

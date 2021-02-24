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

// InvokeWithMessagesRangeRequest represents TL type `invokeWithMessagesRange#365275f2`.
// Invoke with the given message range
//
// See https://core.telegram.org/constructor/invokeWithMessagesRange for reference.
type InvokeWithMessagesRangeRequest struct {
	// Message range
	Range MessageRange `tl:"range"`
	// Query
	Query bin.Object `tl:"query"`
}

// InvokeWithMessagesRangeRequestTypeID is TL type id of InvokeWithMessagesRangeRequest.
const InvokeWithMessagesRangeRequestTypeID = 0x365275f2

func (i *InvokeWithMessagesRangeRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Range.Zero()) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeWithMessagesRangeRequest) String() string {
	if i == nil {
		return "InvokeWithMessagesRangeRequest(nil)"
	}
	type Alias InvokeWithMessagesRangeRequest
	return fmt.Sprintf("InvokeWithMessagesRangeRequest%+v", Alias(*i))
}

// FillFrom fills InvokeWithMessagesRangeRequest from given interface.
func (i *InvokeWithMessagesRangeRequest) FillFrom(from interface {
	GetRange() (value MessageRange)
	GetQuery() (value bin.Object)
}) {
	i.Range = from.GetRange()
	i.Query = from.GetQuery()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InvokeWithMessagesRangeRequest) TypeID() uint32 {
	return InvokeWithMessagesRangeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (i *InvokeWithMessagesRangeRequest) TypeName() string {
	return "invokeWithMessagesRange"
}

// Encode implements bin.Encoder.
func (i *InvokeWithMessagesRangeRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeWithMessagesRange#365275f2 as nil")
	}
	b.PutID(InvokeWithMessagesRangeRequestTypeID)
	if err := i.Range.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithMessagesRange#365275f2: field range: %w", err)
	}
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeWithMessagesRange#365275f2: field query: %w", err)
	}
	return nil
}

// GetRange returns value of Range field.
func (i *InvokeWithMessagesRangeRequest) GetRange() (value MessageRange) {
	return i.Range
}

// GetQuery returns value of Query field.
func (i *InvokeWithMessagesRangeRequest) GetQuery() (value bin.Object) {
	return i.Query
}

// Decode implements bin.Decoder.
func (i *InvokeWithMessagesRangeRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeWithMessagesRange#365275f2 to nil")
	}
	if err := b.ConsumeID(InvokeWithMessagesRangeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeWithMessagesRange#365275f2: %w", err)
	}
	{
		if err := i.Range.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithMessagesRange#365275f2: field range: %w", err)
		}
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeWithMessagesRange#365275f2: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeWithMessagesRangeRequest.
var (
	_ bin.Encoder = &InvokeWithMessagesRangeRequest{}
	_ bin.Decoder = &InvokeWithMessagesRangeRequest{}
)

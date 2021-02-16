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

// ShippingOption represents TL type `shippingOption#b6213cdf`.
// Shipping option
//
// See https://core.telegram.org/constructor/shippingOption for reference.
type ShippingOption struct {
	// Option ID
	ID string
	// Title
	Title string
	// List of price portions
	Prices []LabeledPrice
}

// ShippingOptionTypeID is TL type id of ShippingOption.
const ShippingOptionTypeID = 0xb6213cdf

func (s *ShippingOption) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == "") {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Prices == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ShippingOption) String() string {
	if s == nil {
		return "ShippingOption(nil)"
	}
	type Alias ShippingOption
	return fmt.Sprintf("ShippingOption%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *ShippingOption) TypeID() uint32 {
	return ShippingOptionTypeID
}

// Encode implements bin.Encoder.
func (s *ShippingOption) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode shippingOption#b6213cdf as nil")
	}
	b.PutID(ShippingOptionTypeID)
	b.PutString(s.ID)
	b.PutString(s.Title)
	b.PutVectorHeader(len(s.Prices))
	for idx, v := range s.Prices {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode shippingOption#b6213cdf: field prices element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (s *ShippingOption) GetID() (value string) {
	return s.ID
}

// GetTitle returns value of Title field.
func (s *ShippingOption) GetTitle() (value string) {
	return s.Title
}

// GetPrices returns value of Prices field.
func (s *ShippingOption) GetPrices() (value []LabeledPrice) {
	return s.Prices
}

// Decode implements bin.Decoder.
func (s *ShippingOption) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode shippingOption#b6213cdf to nil")
	}
	if err := b.ConsumeID(ShippingOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode shippingOption#b6213cdf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#b6213cdf: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#b6213cdf: field title: %w", err)
		}
		s.Title = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode shippingOption#b6213cdf: field prices: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LabeledPrice
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode shippingOption#b6213cdf: field prices: %w", err)
			}
			s.Prices = append(s.Prices, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ShippingOption.
var (
	_ bin.Encoder = &ShippingOption{}
	_ bin.Decoder = &ShippingOption{}
)

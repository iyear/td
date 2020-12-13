// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// TestVectorInt represents TL type `testVectorInt#df9eb113`.
//
// See https://localhost:80/doc/constructor/testVectorInt for reference.
type TestVectorInt struct {
	// Vector of numbers
	Value []int32
}

// TestVectorIntTypeID is TL type id of TestVectorInt.
const TestVectorIntTypeID = 0xdf9eb113

// Encode implements bin.Encoder.
func (t *TestVectorInt) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorInt#df9eb113 as nil")
	}
	b.PutID(TestVectorIntTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, v := range t.Value {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorInt) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorInt#df9eb113 to nil")
	}
	if err := b.ConsumeID(TestVectorIntTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorInt#df9eb113: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorInt#df9eb113: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorInt#df9eb113: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorInt.
var (
	_ bin.Encoder = &TestVectorInt{}
	_ bin.Decoder = &TestVectorInt{}
)
// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// TestDummyFunctionRequest represents TL type `test.dummyFunction#c8357709`.
//
// See https://core.telegram.org/method/test.dummyFunction for reference.
type TestDummyFunctionRequest struct {
}

// TestDummyFunctionRequestTypeID is TL type id of TestDummyFunctionRequest.
const TestDummyFunctionRequestTypeID = 0xc8357709

func (d *TestDummyFunctionRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *TestDummyFunctionRequest) String() string {
	if d == nil {
		return "TestDummyFunctionRequest(nil)"
	}
	type Alias TestDummyFunctionRequest
	return fmt.Sprintf("TestDummyFunctionRequest%+v", Alias(*d))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *TestDummyFunctionRequest) TypeID() uint32 {
	return TestDummyFunctionRequestTypeID
}

// Encode implements bin.Encoder.
func (d *TestDummyFunctionRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode test.dummyFunction#c8357709 as nil")
	}
	b.PutID(TestDummyFunctionRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *TestDummyFunctionRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode test.dummyFunction#c8357709 to nil")
	}
	if err := b.ConsumeID(TestDummyFunctionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode test.dummyFunction#c8357709: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for TestDummyFunctionRequest.
var (
	_ bin.Encoder = &TestDummyFunctionRequest{}
	_ bin.Decoder = &TestDummyFunctionRequest{}
)

// TestDummyFunction invokes method test.dummyFunction#c8357709 returning error if any.
//
// See https://core.telegram.org/method/test.dummyFunction for reference.
func (c *Client) TestDummyFunction(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &TestDummyFunctionRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}

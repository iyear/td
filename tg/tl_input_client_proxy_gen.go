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

// InputClientProxy represents TL type `inputClientProxy#75588b3f`.
// Info about an MTProxy¹ used to connect.
//
// Links:
//  1) https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation
//
// See https://core.telegram.org/constructor/inputClientProxy for reference.
type InputClientProxy struct {
	// Proxy address
	Address string `tl:"address"`
	// Proxy port
	Port int `tl:"port"`
}

// InputClientProxyTypeID is TL type id of InputClientProxy.
const InputClientProxyTypeID = 0x75588b3f

func (i *InputClientProxy) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Address == "") {
		return false
	}
	if !(i.Port == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputClientProxy) String() string {
	if i == nil {
		return "InputClientProxy(nil)"
	}
	type Alias InputClientProxy
	return fmt.Sprintf("InputClientProxy%+v", Alias(*i))
}

// FillFrom fills InputClientProxy from given interface.
func (i *InputClientProxy) FillFrom(from interface {
	GetAddress() (value string)
	GetPort() (value int)
}) {
	i.Address = from.GetAddress()
	i.Port = from.GetPort()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputClientProxy) TypeID() uint32 {
	return InputClientProxyTypeID
}

// TypeName returns name of type in TL schema.
func (i *InputClientProxy) TypeName() string {
	return "inputClientProxy"
}

// Encode implements bin.Encoder.
func (i *InputClientProxy) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputClientProxy#75588b3f as nil")
	}
	b.PutID(InputClientProxyTypeID)
	b.PutString(i.Address)
	b.PutInt(i.Port)
	return nil
}

// GetAddress returns value of Address field.
func (i *InputClientProxy) GetAddress() (value string) {
	return i.Address
}

// GetPort returns value of Port field.
func (i *InputClientProxy) GetPort() (value int) {
	return i.Port
}

// Decode implements bin.Decoder.
func (i *InputClientProxy) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputClientProxy#75588b3f to nil")
	}
	if err := b.ConsumeID(InputClientProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputClientProxy#75588b3f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputClientProxy#75588b3f: field address: %w", err)
		}
		i.Address = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputClientProxy#75588b3f: field port: %w", err)
		}
		i.Port = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputClientProxy.
var (
	_ bin.Encoder = &InputClientProxy{}
	_ bin.Decoder = &InputClientProxy{}
)

// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ChangeImportedContactsRequest represents TL type `changeImportedContacts#24885905`.
type ChangeImportedContactsRequest struct {
	// Contacts field of ChangeImportedContactsRequest.
	Contacts []Contact
}

// ChangeImportedContactsRequestTypeID is TL type id of ChangeImportedContactsRequest.
const ChangeImportedContactsRequestTypeID = 0x24885905

// Ensuring interfaces in compile-time for ChangeImportedContactsRequest.
var (
	_ bin.Encoder     = &ChangeImportedContactsRequest{}
	_ bin.Decoder     = &ChangeImportedContactsRequest{}
	_ bin.BareEncoder = &ChangeImportedContactsRequest{}
	_ bin.BareDecoder = &ChangeImportedContactsRequest{}
)

func (c *ChangeImportedContactsRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Contacts == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChangeImportedContactsRequest) String() string {
	if c == nil {
		return "ChangeImportedContactsRequest(nil)"
	}
	type Alias ChangeImportedContactsRequest
	return fmt.Sprintf("ChangeImportedContactsRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChangeImportedContactsRequest) TypeID() uint32 {
	return ChangeImportedContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChangeImportedContactsRequest) TypeName() string {
	return "changeImportedContacts"
}

// TypeInfo returns info about TL type.
func (c *ChangeImportedContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "changeImportedContacts",
		ID:   ChangeImportedContactsRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Contacts",
			SchemaName: "contacts",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChangeImportedContactsRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode changeImportedContacts#24885905 as nil")
	}
	b.PutID(ChangeImportedContactsRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChangeImportedContactsRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode changeImportedContacts#24885905 as nil")
	}
	b.PutInt(len(c.Contacts))
	for idx, v := range c.Contacts {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare changeImportedContacts#24885905: field contacts element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChangeImportedContactsRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode changeImportedContacts#24885905 to nil")
	}
	if err := b.ConsumeID(ChangeImportedContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode changeImportedContacts#24885905: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChangeImportedContactsRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode changeImportedContacts#24885905 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode changeImportedContacts#24885905: field contacts: %w", err)
		}

		if headerLen > 0 {
			c.Contacts = make([]Contact, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Contact
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare changeImportedContacts#24885905: field contacts: %w", err)
			}
			c.Contacts = append(c.Contacts, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChangeImportedContactsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode changeImportedContacts#24885905 as nil")
	}
	b.ObjStart()
	b.PutID("changeImportedContacts")
	b.FieldStart("contacts")
	b.ArrStart()
	for idx, v := range c.Contacts {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode changeImportedContacts#24885905: field contacts element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChangeImportedContactsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode changeImportedContacts#24885905 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("changeImportedContacts"); err != nil {
				return fmt.Errorf("unable to decode changeImportedContacts#24885905: %w", err)
			}
		case "contacts":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Contact
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode changeImportedContacts#24885905: field contacts: %w", err)
				}
				c.Contacts = append(c.Contacts, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode changeImportedContacts#24885905: field contacts: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetContacts returns value of Contacts field.
func (c *ChangeImportedContactsRequest) GetContacts() (value []Contact) {
	if c == nil {
		return
	}
	return c.Contacts
}

// ChangeImportedContacts invokes method changeImportedContacts#24885905 returning error if any.
func (c *Client) ChangeImportedContacts(ctx context.Context, contacts []Contact) (*ImportedContacts, error) {
	var result ImportedContacts

	request := &ChangeImportedContactsRequest{
		Contacts: contacts,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

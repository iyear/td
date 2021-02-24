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

// ContactsGetContactsRequest represents TL type `contacts.getContacts#c023849f`.
// Returns the current user's contact list.
//
// See https://core.telegram.org/method/contacts.getContacts for reference.
type ContactsGetContactsRequest struct {
	// If there already is a full contact list on the client, a hash¹ of a the list of contact IDs in ascending order may be passed in this parameter. If the contact set was not changed, (contacts.contactsNotModified)² will be returned.
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	//  2) https://core.telegram.org/constructor/contacts.contactsNotModified
	Hash int `tl:"hash"`
}

// ContactsGetContactsRequestTypeID is TL type id of ContactsGetContactsRequest.
const ContactsGetContactsRequestTypeID = 0xc023849f

func (g *ContactsGetContactsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ContactsGetContactsRequest) String() string {
	if g == nil {
		return "ContactsGetContactsRequest(nil)"
	}
	type Alias ContactsGetContactsRequest
	return fmt.Sprintf("ContactsGetContactsRequest%+v", Alias(*g))
}

// FillFrom fills ContactsGetContactsRequest from given interface.
func (g *ContactsGetContactsRequest) FillFrom(from interface {
	GetHash() (value int)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *ContactsGetContactsRequest) TypeID() uint32 {
	return ContactsGetContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (g *ContactsGetContactsRequest) TypeName() string {
	return "contacts.getContacts"
}

// Encode implements bin.Encoder.
func (g *ContactsGetContactsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getContacts#c023849f as nil")
	}
	b.PutID(ContactsGetContactsRequestTypeID)
	b.PutInt(g.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (g *ContactsGetContactsRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *ContactsGetContactsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getContacts#c023849f to nil")
	}
	if err := b.ConsumeID(ContactsGetContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getContacts#c023849f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getContacts#c023849f: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetContactsRequest.
var (
	_ bin.Encoder = &ContactsGetContactsRequest{}
	_ bin.Decoder = &ContactsGetContactsRequest{}
)

// ContactsGetContacts invokes method contacts.getContacts#c023849f returning error if any.
// Returns the current user's contact list.
//
// See https://core.telegram.org/method/contacts.getContacts for reference.
func (c *Client) ContactsGetContacts(ctx context.Context, hash int) (ContactsContactsClass, error) {
	var result ContactsContactsBox

	request := &ContactsGetContactsRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Contacts, nil
}

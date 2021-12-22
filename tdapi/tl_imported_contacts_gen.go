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

// ImportedContacts represents TL type `importedContacts#1119a03e`.
type ImportedContacts struct {
	// User identifiers of the imported contacts in the same order as they were specified in
	// the request; 0 if the contact is not yet a registered user
	UserIDs []int64
	// The number of users that imported the corresponding contact; 0 for already registered
	// users or if unavailable
	ImporterCount []int32
}

// ImportedContactsTypeID is TL type id of ImportedContacts.
const ImportedContactsTypeID = 0x1119a03e

// Ensuring interfaces in compile-time for ImportedContacts.
var (
	_ bin.Encoder     = &ImportedContacts{}
	_ bin.Decoder     = &ImportedContacts{}
	_ bin.BareEncoder = &ImportedContacts{}
	_ bin.BareDecoder = &ImportedContacts{}
)

func (i *ImportedContacts) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserIDs == nil) {
		return false
	}
	if !(i.ImporterCount == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ImportedContacts) String() string {
	if i == nil {
		return "ImportedContacts(nil)"
	}
	type Alias ImportedContacts
	return fmt.Sprintf("ImportedContacts%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ImportedContacts) TypeID() uint32 {
	return ImportedContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*ImportedContacts) TypeName() string {
	return "importedContacts"
}

// TypeInfo returns info about TL type.
func (i *ImportedContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "importedContacts",
		ID:   ImportedContactsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
		{
			Name:       "ImporterCount",
			SchemaName: "importer_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *ImportedContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContacts#1119a03e as nil")
	}
	b.PutID(ImportedContactsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *ImportedContacts) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContacts#1119a03e as nil")
	}
	b.PutInt(len(i.UserIDs))
	for _, v := range i.UserIDs {
		b.PutInt53(v)
	}
	b.PutInt(len(i.ImporterCount))
	for _, v := range i.ImporterCount {
		b.PutInt32(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ImportedContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContacts#1119a03e to nil")
	}
	if err := b.ConsumeID(ImportedContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode importedContacts#1119a03e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *ImportedContacts) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContacts#1119a03e to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode importedContacts#1119a03e: field user_ids: %w", err)
		}

		if headerLen > 0 {
			i.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode importedContacts#1119a03e: field user_ids: %w", err)
			}
			i.UserIDs = append(i.UserIDs, value)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode importedContacts#1119a03e: field importer_count: %w", err)
		}

		if headerLen > 0 {
			i.ImporterCount = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode importedContacts#1119a03e: field importer_count: %w", err)
			}
			i.ImporterCount = append(i.ImporterCount, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *ImportedContacts) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode importedContacts#1119a03e as nil")
	}
	b.ObjStart()
	b.PutID("importedContacts")
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range i.UserIDs {
		b.PutInt53(v)
	}
	b.ArrEnd()
	b.FieldStart("importer_count")
	b.ArrStart()
	for _, v := range i.ImporterCount {
		b.PutInt32(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *ImportedContacts) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode importedContacts#1119a03e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("importedContacts"); err != nil {
				return fmt.Errorf("unable to decode importedContacts#1119a03e: %w", err)
			}
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode importedContacts#1119a03e: field user_ids: %w", err)
				}
				i.UserIDs = append(i.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode importedContacts#1119a03e: field user_ids: %w", err)
			}
		case "importer_count":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode importedContacts#1119a03e: field importer_count: %w", err)
				}
				i.ImporterCount = append(i.ImporterCount, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode importedContacts#1119a03e: field importer_count: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserIDs returns value of UserIDs field.
func (i *ImportedContacts) GetUserIDs() (value []int64) {
	if i == nil {
		return
	}
	return i.UserIDs
}

// GetImporterCount returns value of ImporterCount field.
func (i *ImportedContacts) GetImporterCount() (value []int32) {
	if i == nil {
		return
	}
	return i.ImporterCount
}

// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ChatInviteImporter represents TL type `chatInviteImporter#8c5adfd9`.
// When and which user joined the chat using a chat invite
//
// See https://core.telegram.org/constructor/chatInviteImporter for reference.
type ChatInviteImporter struct {
	// Flags field of ChatInviteImporter.
	Flags bin.Fields
	// Requested field of ChatInviteImporter.
	Requested bool
	// The user
	UserID int64
	// When did the user join
	Date int
	// About field of ChatInviteImporter.
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// ApprovedBy field of ChatInviteImporter.
	//
	// Use SetApprovedBy and GetApprovedBy helpers.
	ApprovedBy int64
}

// ChatInviteImporterTypeID is TL type id of ChatInviteImporter.
const ChatInviteImporterTypeID = 0x8c5adfd9

// Ensuring interfaces in compile-time for ChatInviteImporter.
var (
	_ bin.Encoder     = &ChatInviteImporter{}
	_ bin.Decoder     = &ChatInviteImporter{}
	_ bin.BareEncoder = &ChatInviteImporter{}
	_ bin.BareDecoder = &ChatInviteImporter{}
)

func (c *ChatInviteImporter) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Requested == false) {
		return false
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.About == "") {
		return false
	}
	if !(c.ApprovedBy == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteImporter) String() string {
	if c == nil {
		return "ChatInviteImporter(nil)"
	}
	type Alias ChatInviteImporter
	return fmt.Sprintf("ChatInviteImporter%+v", Alias(*c))
}

// FillFrom fills ChatInviteImporter from given interface.
func (c *ChatInviteImporter) FillFrom(from interface {
	GetRequested() (value bool)
	GetUserID() (value int64)
	GetDate() (value int)
	GetAbout() (value string, ok bool)
	GetApprovedBy() (value int64, ok bool)
}) {
	c.Requested = from.GetRequested()
	c.UserID = from.GetUserID()
	c.Date = from.GetDate()
	if val, ok := from.GetAbout(); ok {
		c.About = val
	}

	if val, ok := from.GetApprovedBy(); ok {
		c.ApprovedBy = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteImporter) TypeID() uint32 {
	return ChatInviteImporterTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteImporter) TypeName() string {
	return "chatInviteImporter"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteImporter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteImporter",
		ID:   ChatInviteImporterTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Requested",
			SchemaName: "requested",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "ApprovedBy",
			SchemaName: "approved_by",
			Null:       !c.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChatInviteImporter) SetFlags() {
	if !(c.Requested == false) {
		c.Flags.Set(0)
	}
	if !(c.About == "") {
		c.Flags.Set(2)
	}
	if !(c.ApprovedBy == 0) {
		c.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (c *ChatInviteImporter) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteImporter#8c5adfd9 as nil")
	}
	b.PutID(ChatInviteImporterTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteImporter) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteImporter#8c5adfd9 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteImporter#8c5adfd9: field flags: %w", err)
	}
	b.PutLong(c.UserID)
	b.PutInt(c.Date)
	if c.Flags.Has(2) {
		b.PutString(c.About)
	}
	if c.Flags.Has(1) {
		b.PutLong(c.ApprovedBy)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteImporter) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteImporter#8c5adfd9 to nil")
	}
	if err := b.ConsumeID(ChatInviteImporterTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteImporter#8c5adfd9: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteImporter) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteImporter#8c5adfd9 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#8c5adfd9: field flags: %w", err)
		}
	}
	c.Requested = c.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#8c5adfd9: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#8c5adfd9: field date: %w", err)
		}
		c.Date = value
	}
	if c.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#8c5adfd9: field about: %w", err)
		}
		c.About = value
	}
	if c.Flags.Has(1) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteImporter#8c5adfd9: field approved_by: %w", err)
		}
		c.ApprovedBy = value
	}
	return nil
}

// SetRequested sets value of Requested conditional field.
func (c *ChatInviteImporter) SetRequested(value bool) {
	if value {
		c.Flags.Set(0)
		c.Requested = true
	} else {
		c.Flags.Unset(0)
		c.Requested = false
	}
}

// GetRequested returns value of Requested conditional field.
func (c *ChatInviteImporter) GetRequested() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// GetUserID returns value of UserID field.
func (c *ChatInviteImporter) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetDate returns value of Date field.
func (c *ChatInviteImporter) GetDate() (value int) {
	if c == nil {
		return
	}
	return c.Date
}

// SetAbout sets value of About conditional field.
func (c *ChatInviteImporter) SetAbout(value string) {
	c.Flags.Set(2)
	c.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (c *ChatInviteImporter) GetAbout() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(2) {
		return value, false
	}
	return c.About, true
}

// SetApprovedBy sets value of ApprovedBy conditional field.
func (c *ChatInviteImporter) SetApprovedBy(value int64) {
	c.Flags.Set(1)
	c.ApprovedBy = value
}

// GetApprovedBy returns value of ApprovedBy conditional field and
// boolean which is true if field was set.
func (c *ChatInviteImporter) GetApprovedBy() (value int64, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.ApprovedBy, true
}

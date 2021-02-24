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

// Folder represents TL type `folder#ff544e65`.
// Folder
//
// See https://core.telegram.org/constructor/folder for reference.
type Folder struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Automatically add new channels to this folder
	AutofillNewBroadcasts bool `tl:"autofill_new_broadcasts"`
	// Automatically add joined new public supergroups to this folder
	AutofillPublicGroups bool `tl:"autofill_public_groups"`
	// Automatically add new private chats to this folder
	AutofillNewCorrespondents bool `tl:"autofill_new_correspondents"`
	// Folder ID
	ID int `tl:"id"`
	// Folder title
	Title string `tl:"title"`
	// Folder picture
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo ChatPhotoClass `tl:"photo"`
}

// FolderTypeID is TL type id of Folder.
const FolderTypeID = 0xff544e65

func (f *Folder) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Flags.Zero()) {
		return false
	}
	if !(f.AutofillNewBroadcasts == false) {
		return false
	}
	if !(f.AutofillPublicGroups == false) {
		return false
	}
	if !(f.AutofillNewCorrespondents == false) {
		return false
	}
	if !(f.ID == 0) {
		return false
	}
	if !(f.Title == "") {
		return false
	}
	if !(f.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *Folder) String() string {
	if f == nil {
		return "Folder(nil)"
	}
	type Alias Folder
	return fmt.Sprintf("Folder%+v", Alias(*f))
}

// FillFrom fills Folder from given interface.
func (f *Folder) FillFrom(from interface {
	GetAutofillNewBroadcasts() (value bool)
	GetAutofillPublicGroups() (value bool)
	GetAutofillNewCorrespondents() (value bool)
	GetID() (value int)
	GetTitle() (value string)
	GetPhoto() (value ChatPhotoClass, ok bool)
}) {
	f.AutofillNewBroadcasts = from.GetAutofillNewBroadcasts()
	f.AutofillPublicGroups = from.GetAutofillPublicGroups()
	f.AutofillNewCorrespondents = from.GetAutofillNewCorrespondents()
	f.ID = from.GetID()
	f.Title = from.GetTitle()
	if val, ok := from.GetPhoto(); ok {
		f.Photo = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *Folder) TypeID() uint32 {
	return FolderTypeID
}

// TypeName returns name of type in TL schema.
func (f *Folder) TypeName() string {
	return "folder"
}

// Encode implements bin.Encoder.
func (f *Folder) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode folder#ff544e65 as nil")
	}
	b.PutID(FolderTypeID)
	if !(f.AutofillNewBroadcasts == false) {
		f.Flags.Set(0)
	}
	if !(f.AutofillPublicGroups == false) {
		f.Flags.Set(1)
	}
	if !(f.AutofillNewCorrespondents == false) {
		f.Flags.Set(2)
	}
	if !(f.Photo == nil) {
		f.Flags.Set(3)
	}
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode folder#ff544e65: field flags: %w", err)
	}
	b.PutInt(f.ID)
	b.PutString(f.Title)
	if f.Flags.Has(3) {
		if f.Photo == nil {
			return fmt.Errorf("unable to encode folder#ff544e65: field photo is nil")
		}
		if err := f.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode folder#ff544e65: field photo: %w", err)
		}
	}
	return nil
}

// SetAutofillNewBroadcasts sets value of AutofillNewBroadcasts conditional field.
func (f *Folder) SetAutofillNewBroadcasts(value bool) {
	if value {
		f.Flags.Set(0)
		f.AutofillNewBroadcasts = true
	} else {
		f.Flags.Unset(0)
		f.AutofillNewBroadcasts = false
	}
}

// GetAutofillNewBroadcasts returns value of AutofillNewBroadcasts conditional field.
func (f *Folder) GetAutofillNewBroadcasts() (value bool) {
	return f.Flags.Has(0)
}

// SetAutofillPublicGroups sets value of AutofillPublicGroups conditional field.
func (f *Folder) SetAutofillPublicGroups(value bool) {
	if value {
		f.Flags.Set(1)
		f.AutofillPublicGroups = true
	} else {
		f.Flags.Unset(1)
		f.AutofillPublicGroups = false
	}
}

// GetAutofillPublicGroups returns value of AutofillPublicGroups conditional field.
func (f *Folder) GetAutofillPublicGroups() (value bool) {
	return f.Flags.Has(1)
}

// SetAutofillNewCorrespondents sets value of AutofillNewCorrespondents conditional field.
func (f *Folder) SetAutofillNewCorrespondents(value bool) {
	if value {
		f.Flags.Set(2)
		f.AutofillNewCorrespondents = true
	} else {
		f.Flags.Unset(2)
		f.AutofillNewCorrespondents = false
	}
}

// GetAutofillNewCorrespondents returns value of AutofillNewCorrespondents conditional field.
func (f *Folder) GetAutofillNewCorrespondents() (value bool) {
	return f.Flags.Has(2)
}

// GetID returns value of ID field.
func (f *Folder) GetID() (value int) {
	return f.ID
}

// GetTitle returns value of Title field.
func (f *Folder) GetTitle() (value string) {
	return f.Title
}

// SetPhoto sets value of Photo conditional field.
func (f *Folder) SetPhoto(value ChatPhotoClass) {
	f.Flags.Set(3)
	f.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (f *Folder) GetPhoto() (value ChatPhotoClass, ok bool) {
	if !f.Flags.Has(3) {
		return value, false
	}
	return f.Photo, true
}

// GetPhotoAsNotEmpty returns mapped value of Photo conditional field and
// boolean which is true if field was set.
func (f *Folder) GetPhotoAsNotEmpty() (*ChatPhoto, bool) {
	if value, ok := f.GetPhoto(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// Decode implements bin.Decoder.
func (f *Folder) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode folder#ff544e65 to nil")
	}
	if err := b.ConsumeID(FolderTypeID); err != nil {
		return fmt.Errorf("unable to decode folder#ff544e65: %w", err)
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field flags: %w", err)
		}
	}
	f.AutofillNewBroadcasts = f.Flags.Has(0)
	f.AutofillPublicGroups = f.Flags.Has(1)
	f.AutofillNewCorrespondents = f.Flags.Has(2)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field id: %w", err)
		}
		f.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field title: %w", err)
		}
		f.Title = value
	}
	if f.Flags.Has(3) {
		value, err := DecodeChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field photo: %w", err)
		}
		f.Photo = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Folder.
var (
	_ bin.Encoder = &Folder{}
	_ bin.Decoder = &Folder{}
)

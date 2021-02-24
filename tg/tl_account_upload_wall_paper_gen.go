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

// AccountUploadWallPaperRequest represents TL type `account.uploadWallPaper#dd853661`.
// Create and upload a new wallpaper
//
// See https://core.telegram.org/method/account.uploadWallPaper for reference.
type AccountUploadWallPaperRequest struct {
	// The JPG/PNG wallpaper
	File InputFileClass `tl:"file"`
	// MIME type of uploaded wallpaper
	MimeType string `tl:"mime_type"`
	// Wallpaper settings
	Settings WallPaperSettings `tl:"settings"`
}

// AccountUploadWallPaperRequestTypeID is TL type id of AccountUploadWallPaperRequest.
const AccountUploadWallPaperRequestTypeID = 0xdd853661

func (u *AccountUploadWallPaperRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.MimeType == "") {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUploadWallPaperRequest) String() string {
	if u == nil {
		return "AccountUploadWallPaperRequest(nil)"
	}
	type Alias AccountUploadWallPaperRequest
	return fmt.Sprintf("AccountUploadWallPaperRequest%+v", Alias(*u))
}

// FillFrom fills AccountUploadWallPaperRequest from given interface.
func (u *AccountUploadWallPaperRequest) FillFrom(from interface {
	GetFile() (value InputFileClass)
	GetMimeType() (value string)
	GetSettings() (value WallPaperSettings)
}) {
	u.File = from.GetFile()
	u.MimeType = from.GetMimeType()
	u.Settings = from.GetSettings()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *AccountUploadWallPaperRequest) TypeID() uint32 {
	return AccountUploadWallPaperRequestTypeID
}

// TypeName returns name of type in TL schema.
func (u *AccountUploadWallPaperRequest) TypeName() string {
	return "account.uploadWallPaper"
}

// Encode implements bin.Encoder.
func (u *AccountUploadWallPaperRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadWallPaper#dd853661 as nil")
	}
	b.PutID(AccountUploadWallPaperRequestTypeID)
	if u.File == nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#dd853661: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#dd853661: field file: %w", err)
	}
	b.PutString(u.MimeType)
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#dd853661: field settings: %w", err)
	}
	return nil
}

// GetFile returns value of File field.
func (u *AccountUploadWallPaperRequest) GetFile() (value InputFileClass) {
	return u.File
}

// GetMimeType returns value of MimeType field.
func (u *AccountUploadWallPaperRequest) GetMimeType() (value string) {
	return u.MimeType
}

// GetSettings returns value of Settings field.
func (u *AccountUploadWallPaperRequest) GetSettings() (value WallPaperSettings) {
	return u.Settings
}

// Decode implements bin.Decoder.
func (u *AccountUploadWallPaperRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadWallPaper#dd853661 to nil")
	}
	if err := b.ConsumeID(AccountUploadWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: %w", err)
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: field file: %w", err)
		}
		u.File = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: field mime_type: %w", err)
		}
		u.MimeType = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUploadWallPaperRequest.
var (
	_ bin.Encoder = &AccountUploadWallPaperRequest{}
	_ bin.Decoder = &AccountUploadWallPaperRequest{}
)

// AccountUploadWallPaper invokes method account.uploadWallPaper#dd853661 returning error if any.
// Create and upload a new wallpaper
//
// See https://core.telegram.org/method/account.uploadWallPaper for reference.
func (c *Client) AccountUploadWallPaper(ctx context.Context, request *AccountUploadWallPaperRequest) (WallPaperClass, error) {
	var result WallPaperBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WallPaper, nil
}

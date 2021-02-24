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

// UploadSaveBigFilePartRequest represents TL type `upload.saveBigFilePart#de7b673d`.
// Saves a part of a large file (over 10Mb in size) to be later passed to one of the methods.
//
// See https://core.telegram.org/method/upload.saveBigFilePart for reference.
type UploadSaveBigFilePartRequest struct {
	// Random file id, created by the client
	FileID int64 `tl:"file_id"`
	// Part sequence number
	FilePart int `tl:"file_part"`
	// Total number of parts
	FileTotalParts int `tl:"file_total_parts"`
	// Binary data, part contents
	Bytes []byte `tl:"bytes"`
}

// UploadSaveBigFilePartRequestTypeID is TL type id of UploadSaveBigFilePartRequest.
const UploadSaveBigFilePartRequestTypeID = 0xde7b673d

func (s *UploadSaveBigFilePartRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.FileID == 0) {
		return false
	}
	if !(s.FilePart == 0) {
		return false
	}
	if !(s.FileTotalParts == 0) {
		return false
	}
	if !(s.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *UploadSaveBigFilePartRequest) String() string {
	if s == nil {
		return "UploadSaveBigFilePartRequest(nil)"
	}
	type Alias UploadSaveBigFilePartRequest
	return fmt.Sprintf("UploadSaveBigFilePartRequest%+v", Alias(*s))
}

// FillFrom fills UploadSaveBigFilePartRequest from given interface.
func (s *UploadSaveBigFilePartRequest) FillFrom(from interface {
	GetFileID() (value int64)
	GetFilePart() (value int)
	GetFileTotalParts() (value int)
	GetBytes() (value []byte)
}) {
	s.FileID = from.GetFileID()
	s.FilePart = from.GetFilePart()
	s.FileTotalParts = from.GetFileTotalParts()
	s.Bytes = from.GetBytes()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *UploadSaveBigFilePartRequest) TypeID() uint32 {
	return UploadSaveBigFilePartRequestTypeID
}

// TypeName returns name of type in TL schema.
func (s *UploadSaveBigFilePartRequest) TypeName() string {
	return "upload.saveBigFilePart"
}

// Encode implements bin.Encoder.
func (s *UploadSaveBigFilePartRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode upload.saveBigFilePart#de7b673d as nil")
	}
	b.PutID(UploadSaveBigFilePartRequestTypeID)
	b.PutLong(s.FileID)
	b.PutInt(s.FilePart)
	b.PutInt(s.FileTotalParts)
	b.PutBytes(s.Bytes)
	return nil
}

// GetFileID returns value of FileID field.
func (s *UploadSaveBigFilePartRequest) GetFileID() (value int64) {
	return s.FileID
}

// GetFilePart returns value of FilePart field.
func (s *UploadSaveBigFilePartRequest) GetFilePart() (value int) {
	return s.FilePart
}

// GetFileTotalParts returns value of FileTotalParts field.
func (s *UploadSaveBigFilePartRequest) GetFileTotalParts() (value int) {
	return s.FileTotalParts
}

// GetBytes returns value of Bytes field.
func (s *UploadSaveBigFilePartRequest) GetBytes() (value []byte) {
	return s.Bytes
}

// Decode implements bin.Decoder.
func (s *UploadSaveBigFilePartRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode upload.saveBigFilePart#de7b673d to nil")
	}
	if err := b.ConsumeID(UploadSaveBigFilePartRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field file_id: %w", err)
		}
		s.FileID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field file_part: %w", err)
		}
		s.FilePart = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field file_total_parts: %w", err)
		}
		s.FileTotalParts = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.saveBigFilePart#de7b673d: field bytes: %w", err)
		}
		s.Bytes = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UploadSaveBigFilePartRequest.
var (
	_ bin.Encoder = &UploadSaveBigFilePartRequest{}
	_ bin.Decoder = &UploadSaveBigFilePartRequest{}
)

// UploadSaveBigFilePart invokes method upload.saveBigFilePart#de7b673d returning error if any.
// Saves a part of a large file (over 10Mb in size) to be later passed to one of the methods.
//
// Possible errors:
//  400 FILE_PARTS_INVALID: The number of file parts is invalid
//  400 FILE_PART_EMPTY: The provided file part is empty
//  400 FILE_PART_INVALID: The file part number is invalid
//  400 FILE_PART_SIZE_CHANGED: Provided file part size has changed
//  400 FILE_PART_SIZE_INVALID: The provided file part size is invalid
//  400 FILE_PART_TOO_BIG: The uploaded file part is too big
//
// See https://core.telegram.org/method/upload.saveBigFilePart for reference.
// Can be used by bots.
func (c *Client) UploadSaveBigFilePart(ctx context.Context, request *UploadSaveBigFilePartRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}

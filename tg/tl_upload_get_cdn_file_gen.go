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

// UploadGetCDNFileRequest represents TL type `upload.getCdnFile#2000bcc3`.
// Download a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.getCdnFile for reference.
type UploadGetCDNFileRequest struct {
	// File token
	FileToken []byte
	// Offset of chunk to download
	Offset int
	// Length of chunk to download
	Limit int
}

// UploadGetCDNFileRequestTypeID is TL type id of UploadGetCDNFileRequest.
const UploadGetCDNFileRequestTypeID = 0x2000bcc3

// Ensuring interfaces in compile-time for UploadGetCDNFileRequest.
var (
	_ bin.Encoder     = &UploadGetCDNFileRequest{}
	_ bin.Decoder     = &UploadGetCDNFileRequest{}
	_ bin.BareEncoder = &UploadGetCDNFileRequest{}
	_ bin.BareDecoder = &UploadGetCDNFileRequest{}
)

func (g *UploadGetCDNFileRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.FileToken == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UploadGetCDNFileRequest) String() string {
	if g == nil {
		return "UploadGetCDNFileRequest(nil)"
	}
	type Alias UploadGetCDNFileRequest
	return fmt.Sprintf("UploadGetCDNFileRequest%+v", Alias(*g))
}

// FillFrom fills UploadGetCDNFileRequest from given interface.
func (g *UploadGetCDNFileRequest) FillFrom(from interface {
	GetFileToken() (value []byte)
	GetOffset() (value int)
	GetLimit() (value int)
}) {
	g.FileToken = from.GetFileToken()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UploadGetCDNFileRequest) TypeID() uint32 {
	return UploadGetCDNFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UploadGetCDNFileRequest) TypeName() string {
	return "upload.getCdnFile"
}

// TypeInfo returns info about TL type.
func (g *UploadGetCDNFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upload.getCdnFile",
		ID:   UploadGetCDNFileRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileToken",
			SchemaName: "file_token",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UploadGetCDNFileRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getCdnFile#2000bcc3 as nil")
	}
	b.PutID(UploadGetCDNFileRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UploadGetCDNFileRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode upload.getCdnFile#2000bcc3 as nil")
	}
	b.PutBytes(g.FileToken)
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *UploadGetCDNFileRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getCdnFile#2000bcc3 to nil")
	}
	if err := b.ConsumeID(UploadGetCDNFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UploadGetCDNFileRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode upload.getCdnFile#2000bcc3 to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: field file_token: %w", err)
		}
		g.FileToken = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode upload.getCdnFile#2000bcc3: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetFileToken returns value of FileToken field.
func (g *UploadGetCDNFileRequest) GetFileToken() (value []byte) {
	if g == nil {
		return
	}
	return g.FileToken
}

// GetOffset returns value of Offset field.
func (g *UploadGetCDNFileRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *UploadGetCDNFileRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// UploadGetCDNFile invokes method upload.getCdnFile#2000bcc3 returning error if any.
// Download a CDN¹ file.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/method/upload.getCdnFile for reference.
func (c *Client) UploadGetCDNFile(ctx context.Context, request *UploadGetCDNFileRequest) (UploadCDNFileClass, error) {
	var result UploadCDNFileBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.CdnFile, nil
}

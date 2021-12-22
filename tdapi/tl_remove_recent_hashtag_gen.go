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

// RemoveRecentHashtagRequest represents TL type `removeRecentHashtag#c393a0a4`.
type RemoveRecentHashtagRequest struct {
	// Hashtag to delete
	Hashtag string
}

// RemoveRecentHashtagRequestTypeID is TL type id of RemoveRecentHashtagRequest.
const RemoveRecentHashtagRequestTypeID = 0xc393a0a4

// Ensuring interfaces in compile-time for RemoveRecentHashtagRequest.
var (
	_ bin.Encoder     = &RemoveRecentHashtagRequest{}
	_ bin.Decoder     = &RemoveRecentHashtagRequest{}
	_ bin.BareEncoder = &RemoveRecentHashtagRequest{}
	_ bin.BareDecoder = &RemoveRecentHashtagRequest{}
)

func (r *RemoveRecentHashtagRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hashtag == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveRecentHashtagRequest) String() string {
	if r == nil {
		return "RemoveRecentHashtagRequest(nil)"
	}
	type Alias RemoveRecentHashtagRequest
	return fmt.Sprintf("RemoveRecentHashtagRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveRecentHashtagRequest) TypeID() uint32 {
	return RemoveRecentHashtagRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveRecentHashtagRequest) TypeName() string {
	return "removeRecentHashtag"
}

// TypeInfo returns info about TL type.
func (r *RemoveRecentHashtagRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeRecentHashtag",
		ID:   RemoveRecentHashtagRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hashtag",
			SchemaName: "hashtag",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveRecentHashtagRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentHashtag#c393a0a4 as nil")
	}
	b.PutID(RemoveRecentHashtagRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveRecentHashtagRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentHashtag#c393a0a4 as nil")
	}
	b.PutString(r.Hashtag)
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveRecentHashtagRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentHashtag#c393a0a4 to nil")
	}
	if err := b.ConsumeID(RemoveRecentHashtagRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeRecentHashtag#c393a0a4: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveRecentHashtagRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentHashtag#c393a0a4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode removeRecentHashtag#c393a0a4: field hashtag: %w", err)
		}
		r.Hashtag = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *RemoveRecentHashtagRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentHashtag#c393a0a4 as nil")
	}
	b.ObjStart()
	b.PutID("removeRecentHashtag")
	b.FieldStart("hashtag")
	b.PutString(r.Hashtag)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *RemoveRecentHashtagRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentHashtag#c393a0a4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("removeRecentHashtag"); err != nil {
				return fmt.Errorf("unable to decode removeRecentHashtag#c393a0a4: %w", err)
			}
		case "hashtag":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode removeRecentHashtag#c393a0a4: field hashtag: %w", err)
			}
			r.Hashtag = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetHashtag returns value of Hashtag field.
func (r *RemoveRecentHashtagRequest) GetHashtag() (value string) {
	if r == nil {
		return
	}
	return r.Hashtag
}

// RemoveRecentHashtag invokes method removeRecentHashtag#c393a0a4 returning error if any.
func (c *Client) RemoveRecentHashtag(ctx context.Context, hashtag string) error {
	var ok Ok

	request := &RemoveRecentHashtagRequest{
		Hashtag: hashtag,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

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

// EndGroupCallRecordingRequest represents TL type `endGroupCallRecording#fb7b6289`.
type EndGroupCallRecordingRequest struct {
	// Group call identifier
	GroupCallID int32
}

// EndGroupCallRecordingRequestTypeID is TL type id of EndGroupCallRecordingRequest.
const EndGroupCallRecordingRequestTypeID = 0xfb7b6289

// Ensuring interfaces in compile-time for EndGroupCallRecordingRequest.
var (
	_ bin.Encoder     = &EndGroupCallRecordingRequest{}
	_ bin.Decoder     = &EndGroupCallRecordingRequest{}
	_ bin.BareEncoder = &EndGroupCallRecordingRequest{}
	_ bin.BareDecoder = &EndGroupCallRecordingRequest{}
)

func (e *EndGroupCallRecordingRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.GroupCallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EndGroupCallRecordingRequest) String() string {
	if e == nil {
		return "EndGroupCallRecordingRequest(nil)"
	}
	type Alias EndGroupCallRecordingRequest
	return fmt.Sprintf("EndGroupCallRecordingRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EndGroupCallRecordingRequest) TypeID() uint32 {
	return EndGroupCallRecordingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EndGroupCallRecordingRequest) TypeName() string {
	return "endGroupCallRecording"
}

// TypeInfo returns info about TL type.
func (e *EndGroupCallRecordingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "endGroupCallRecording",
		ID:   EndGroupCallRecordingRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EndGroupCallRecordingRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCallRecording#fb7b6289 as nil")
	}
	b.PutID(EndGroupCallRecordingRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EndGroupCallRecordingRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCallRecording#fb7b6289 as nil")
	}
	b.PutInt32(e.GroupCallID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EndGroupCallRecordingRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCallRecording#fb7b6289 to nil")
	}
	if err := b.ConsumeID(EndGroupCallRecordingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode endGroupCallRecording#fb7b6289: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EndGroupCallRecordingRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCallRecording#fb7b6289 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode endGroupCallRecording#fb7b6289: field group_call_id: %w", err)
		}
		e.GroupCallID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EndGroupCallRecordingRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCallRecording#fb7b6289 as nil")
	}
	b.ObjStart()
	b.PutID("endGroupCallRecording")
	b.FieldStart("group_call_id")
	b.PutInt32(e.GroupCallID)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EndGroupCallRecordingRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCallRecording#fb7b6289 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("endGroupCallRecording"); err != nil {
				return fmt.Errorf("unable to decode endGroupCallRecording#fb7b6289: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode endGroupCallRecording#fb7b6289: field group_call_id: %w", err)
			}
			e.GroupCallID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (e *EndGroupCallRecordingRequest) GetGroupCallID() (value int32) {
	if e == nil {
		return
	}
	return e.GroupCallID
}

// EndGroupCallRecording invokes method endGroupCallRecording#fb7b6289 returning error if any.
func (c *Client) EndGroupCallRecording(ctx context.Context, groupcallid int32) error {
	var ok Ok

	request := &EndGroupCallRecordingRequest{
		GroupCallID: groupcallid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}

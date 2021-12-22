// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// MsgsStateReq represents TL type `msgs_state_req#da69fb52`.
type MsgsStateReq struct {
	// MsgIDs field of MsgsStateReq.
	MsgIDs []int64
}

// MsgsStateReqTypeID is TL type id of MsgsStateReq.
const MsgsStateReqTypeID = 0xda69fb52

// Ensuring interfaces in compile-time for MsgsStateReq.
var (
	_ bin.Encoder     = &MsgsStateReq{}
	_ bin.Decoder     = &MsgsStateReq{}
	_ bin.BareEncoder = &MsgsStateReq{}
	_ bin.BareDecoder = &MsgsStateReq{}
)

func (m *MsgsStateReq) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgsStateReq) String() string {
	if m == nil {
		return "MsgsStateReq(nil)"
	}
	type Alias MsgsStateReq
	return fmt.Sprintf("MsgsStateReq%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgsStateReq) TypeID() uint32 {
	return MsgsStateReqTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgsStateReq) TypeName() string {
	return "msgs_state_req"
}

// TypeInfo returns info about TL type.
func (m *MsgsStateReq) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msgs_state_req",
		ID:   MsgsStateReqTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgIDs",
			SchemaName: "msg_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgsStateReq) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_state_req#da69fb52 as nil")
	}
	b.PutID(MsgsStateReqTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgsStateReq) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msgs_state_req#da69fb52 as nil")
	}
	b.PutVectorHeader(len(m.MsgIDs))
	for _, v := range m.MsgIDs {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgsStateReq) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_state_req#da69fb52 to nil")
	}
	if err := b.ConsumeID(MsgsStateReqTypeID); err != nil {
		return fmt.Errorf("unable to decode msgs_state_req#da69fb52: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgsStateReq) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msgs_state_req#da69fb52 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msgs_state_req#da69fb52: field msg_ids: %w", err)
		}

		if headerLen > 0 {
			m.MsgIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msgs_state_req#da69fb52: field msg_ids: %w", err)
			}
			m.MsgIDs = append(m.MsgIDs, value)
		}
	}
	return nil
}

// GetMsgIDs returns value of MsgIDs field.
func (m *MsgsStateReq) GetMsgIDs() (value []int64) {
	if m == nil {
		return
	}
	return m.MsgIDs
}

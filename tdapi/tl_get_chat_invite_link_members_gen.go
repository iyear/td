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

// GetChatInviteLinkMembersRequest represents TL type `getChatInviteLinkMembers#c5b6199a`.
type GetChatInviteLinkMembersRequest struct {
	// Chat identifier
	ChatID int64
	// Invite link for which to return chat members
	InviteLink string
	// A chat member from which to return next chat members; pass null to get results from
	// the beginning
	OffsetMember ChatInviteLinkMember
	// The maximum number of chat members to return; up to 100
	Limit int32
}

// GetChatInviteLinkMembersRequestTypeID is TL type id of GetChatInviteLinkMembersRequest.
const GetChatInviteLinkMembersRequestTypeID = 0xc5b6199a

// Ensuring interfaces in compile-time for GetChatInviteLinkMembersRequest.
var (
	_ bin.Encoder     = &GetChatInviteLinkMembersRequest{}
	_ bin.Decoder     = &GetChatInviteLinkMembersRequest{}
	_ bin.BareEncoder = &GetChatInviteLinkMembersRequest{}
	_ bin.BareDecoder = &GetChatInviteLinkMembersRequest{}
)

func (g *GetChatInviteLinkMembersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.InviteLink == "") {
		return false
	}
	if !(g.OffsetMember.Zero()) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatInviteLinkMembersRequest) String() string {
	if g == nil {
		return "GetChatInviteLinkMembersRequest(nil)"
	}
	type Alias GetChatInviteLinkMembersRequest
	return fmt.Sprintf("GetChatInviteLinkMembersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatInviteLinkMembersRequest) TypeID() uint32 {
	return GetChatInviteLinkMembersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatInviteLinkMembersRequest) TypeName() string {
	return "getChatInviteLinkMembers"
}

// TypeInfo returns info about TL type.
func (g *GetChatInviteLinkMembersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatInviteLinkMembers",
		ID:   GetChatInviteLinkMembersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
		{
			Name:       "OffsetMember",
			SchemaName: "offset_member",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatInviteLinkMembersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkMembers#c5b6199a as nil")
	}
	b.PutID(GetChatInviteLinkMembersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatInviteLinkMembersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkMembers#c5b6199a as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutString(g.InviteLink)
	if err := g.OffsetMember.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatInviteLinkMembers#c5b6199a: field offset_member: %w", err)
	}
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatInviteLinkMembersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkMembers#c5b6199a to nil")
	}
	if err := b.ConsumeID(GetChatInviteLinkMembersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatInviteLinkMembersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkMembers#c5b6199a to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field invite_link: %w", err)
		}
		g.InviteLink = value
	}
	{
		if err := g.OffsetMember.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field offset_member: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatInviteLinkMembersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatInviteLinkMembers#c5b6199a as nil")
	}
	b.ObjStart()
	b.PutID("getChatInviteLinkMembers")
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.FieldStart("invite_link")
	b.PutString(g.InviteLink)
	b.FieldStart("offset_member")
	if err := g.OffsetMember.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatInviteLinkMembers#c5b6199a: field offset_member: %w", err)
	}
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatInviteLinkMembersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatInviteLinkMembers#c5b6199a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatInviteLinkMembers"); err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field chat_id: %w", err)
			}
			g.ChatID = value
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field invite_link: %w", err)
			}
			g.InviteLink = value
		case "offset_member":
			if err := g.OffsetMember.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field offset_member: %w", err)
			}
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatInviteLinkMembers#c5b6199a: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatInviteLinkMembersRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetInviteLink returns value of InviteLink field.
func (g *GetChatInviteLinkMembersRequest) GetInviteLink() (value string) {
	if g == nil {
		return
	}
	return g.InviteLink
}

// GetOffsetMember returns value of OffsetMember field.
func (g *GetChatInviteLinkMembersRequest) GetOffsetMember() (value ChatInviteLinkMember) {
	if g == nil {
		return
	}
	return g.OffsetMember
}

// GetLimit returns value of Limit field.
func (g *GetChatInviteLinkMembersRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatInviteLinkMembers invokes method getChatInviteLinkMembers#c5b6199a returning error if any.
func (c *Client) GetChatInviteLinkMembers(ctx context.Context, request *GetChatInviteLinkMembersRequest) (*ChatInviteLinkMembers, error) {
	var result ChatInviteLinkMembers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

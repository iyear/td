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

// SetStoryPrivacySettingsRequest represents TL type `setStoryPrivacySettings#d8e94332`.
type SetStoryPrivacySettingsRequest struct {
	// Identifier of the story
	StoryID int32
	// The new privacy settigs for the story
	PrivacySettings StoryPrivacySettingsClass
}

// SetStoryPrivacySettingsRequestTypeID is TL type id of SetStoryPrivacySettingsRequest.
const SetStoryPrivacySettingsRequestTypeID = 0xd8e94332

// Ensuring interfaces in compile-time for SetStoryPrivacySettingsRequest.
var (
	_ bin.Encoder     = &SetStoryPrivacySettingsRequest{}
	_ bin.Decoder     = &SetStoryPrivacySettingsRequest{}
	_ bin.BareEncoder = &SetStoryPrivacySettingsRequest{}
	_ bin.BareDecoder = &SetStoryPrivacySettingsRequest{}
)

func (s *SetStoryPrivacySettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.StoryID == 0) {
		return false
	}
	if !(s.PrivacySettings == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetStoryPrivacySettingsRequest) String() string {
	if s == nil {
		return "SetStoryPrivacySettingsRequest(nil)"
	}
	type Alias SetStoryPrivacySettingsRequest
	return fmt.Sprintf("SetStoryPrivacySettingsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetStoryPrivacySettingsRequest) TypeID() uint32 {
	return SetStoryPrivacySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetStoryPrivacySettingsRequest) TypeName() string {
	return "setStoryPrivacySettings"
}

// TypeInfo returns info about TL type.
func (s *SetStoryPrivacySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setStoryPrivacySettings",
		ID:   SetStoryPrivacySettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
		{
			Name:       "PrivacySettings",
			SchemaName: "privacy_settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetStoryPrivacySettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setStoryPrivacySettings#d8e94332 as nil")
	}
	b.PutID(SetStoryPrivacySettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetStoryPrivacySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setStoryPrivacySettings#d8e94332 as nil")
	}
	b.PutInt32(s.StoryID)
	if s.PrivacySettings == nil {
		return fmt.Errorf("unable to encode setStoryPrivacySettings#d8e94332: field privacy_settings is nil")
	}
	if err := s.PrivacySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setStoryPrivacySettings#d8e94332: field privacy_settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetStoryPrivacySettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setStoryPrivacySettings#d8e94332 to nil")
	}
	if err := b.ConsumeID(SetStoryPrivacySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setStoryPrivacySettings#d8e94332: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetStoryPrivacySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setStoryPrivacySettings#d8e94332 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setStoryPrivacySettings#d8e94332: field story_id: %w", err)
		}
		s.StoryID = value
	}
	{
		value, err := DecodeStoryPrivacySettings(b)
		if err != nil {
			return fmt.Errorf("unable to decode setStoryPrivacySettings#d8e94332: field privacy_settings: %w", err)
		}
		s.PrivacySettings = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetStoryPrivacySettingsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setStoryPrivacySettings#d8e94332 as nil")
	}
	b.ObjStart()
	b.PutID("setStoryPrivacySettings")
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(s.StoryID)
	b.Comma()
	b.FieldStart("privacy_settings")
	if s.PrivacySettings == nil {
		return fmt.Errorf("unable to encode setStoryPrivacySettings#d8e94332: field privacy_settings is nil")
	}
	if err := s.PrivacySettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setStoryPrivacySettings#d8e94332: field privacy_settings: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetStoryPrivacySettingsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setStoryPrivacySettings#d8e94332 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setStoryPrivacySettings"); err != nil {
				return fmt.Errorf("unable to decode setStoryPrivacySettings#d8e94332: %w", err)
			}
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setStoryPrivacySettings#d8e94332: field story_id: %w", err)
			}
			s.StoryID = value
		case "privacy_settings":
			value, err := DecodeTDLibJSONStoryPrivacySettings(b)
			if err != nil {
				return fmt.Errorf("unable to decode setStoryPrivacySettings#d8e94332: field privacy_settings: %w", err)
			}
			s.PrivacySettings = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStoryID returns value of StoryID field.
func (s *SetStoryPrivacySettingsRequest) GetStoryID() (value int32) {
	if s == nil {
		return
	}
	return s.StoryID
}

// GetPrivacySettings returns value of PrivacySettings field.
func (s *SetStoryPrivacySettingsRequest) GetPrivacySettings() (value StoryPrivacySettingsClass) {
	if s == nil {
		return
	}
	return s.PrivacySettings
}

// SetStoryPrivacySettings invokes method setStoryPrivacySettings#d8e94332 returning error if any.
func (c *Client) SetStoryPrivacySettings(ctx context.Context, request *SetStoryPrivacySettingsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
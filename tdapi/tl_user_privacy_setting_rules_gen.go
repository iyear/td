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

// UserPrivacySettingRules represents TL type `userPrivacySettingRules#425e6b37`.
type UserPrivacySettingRules struct {
	// A list of rules
	Rules []UserPrivacySettingRuleClass
}

// UserPrivacySettingRulesTypeID is TL type id of UserPrivacySettingRules.
const UserPrivacySettingRulesTypeID = 0x425e6b37

// Ensuring interfaces in compile-time for UserPrivacySettingRules.
var (
	_ bin.Encoder     = &UserPrivacySettingRules{}
	_ bin.Decoder     = &UserPrivacySettingRules{}
	_ bin.BareEncoder = &UserPrivacySettingRules{}
	_ bin.BareDecoder = &UserPrivacySettingRules{}
)

func (u *UserPrivacySettingRules) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Rules == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserPrivacySettingRules) String() string {
	if u == nil {
		return "UserPrivacySettingRules(nil)"
	}
	type Alias UserPrivacySettingRules
	return fmt.Sprintf("UserPrivacySettingRules%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserPrivacySettingRules) TypeID() uint32 {
	return UserPrivacySettingRulesTypeID
}

// TypeName returns name of type in TL schema.
func (*UserPrivacySettingRules) TypeName() string {
	return "userPrivacySettingRules"
}

// TypeInfo returns info about TL type.
func (u *UserPrivacySettingRules) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userPrivacySettingRules",
		ID:   UserPrivacySettingRulesTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Rules",
			SchemaName: "rules",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserPrivacySettingRules) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userPrivacySettingRules#425e6b37 as nil")
	}
	b.PutID(UserPrivacySettingRulesTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserPrivacySettingRules) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userPrivacySettingRules#425e6b37 as nil")
	}
	b.PutInt(len(u.Rules))
	for idx, v := range u.Rules {
		if v == nil {
			return fmt.Errorf("unable to encode userPrivacySettingRules#425e6b37: field rules element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare userPrivacySettingRules#425e6b37: field rules element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserPrivacySettingRules) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userPrivacySettingRules#425e6b37 to nil")
	}
	if err := b.ConsumeID(UserPrivacySettingRulesTypeID); err != nil {
		return fmt.Errorf("unable to decode userPrivacySettingRules#425e6b37: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserPrivacySettingRules) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userPrivacySettingRules#425e6b37 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userPrivacySettingRules#425e6b37: field rules: %w", err)
		}

		if headerLen > 0 {
			u.Rules = make([]UserPrivacySettingRuleClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUserPrivacySettingRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode userPrivacySettingRules#425e6b37: field rules: %w", err)
			}
			u.Rules = append(u.Rules, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UserPrivacySettingRules) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode userPrivacySettingRules#425e6b37 as nil")
	}
	b.ObjStart()
	b.PutID("userPrivacySettingRules")
	b.FieldStart("rules")
	b.ArrStart()
	for idx, v := range u.Rules {
		if v == nil {
			return fmt.Errorf("unable to encode userPrivacySettingRules#425e6b37: field rules element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode userPrivacySettingRules#425e6b37: field rules element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UserPrivacySettingRules) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode userPrivacySettingRules#425e6b37 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("userPrivacySettingRules"); err != nil {
				return fmt.Errorf("unable to decode userPrivacySettingRules#425e6b37: %w", err)
			}
		case "rules":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONUserPrivacySettingRule(b)
				if err != nil {
					return fmt.Errorf("unable to decode userPrivacySettingRules#425e6b37: field rules: %w", err)
				}
				u.Rules = append(u.Rules, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode userPrivacySettingRules#425e6b37: field rules: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRules returns value of Rules field.
func (u *UserPrivacySettingRules) GetRules() (value []UserPrivacySettingRuleClass) {
	if u == nil {
		return
	}
	return u.Rules
}

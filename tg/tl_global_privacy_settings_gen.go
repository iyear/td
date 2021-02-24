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

// GlobalPrivacySettings represents TL type `globalPrivacySettings#bea2f424`.
// Global privacy settings
//
// See https://core.telegram.org/constructor/globalPrivacySettings for reference.
type GlobalPrivacySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `tl:"flags"`
	// Whether to archive and mute new chats from non-contacts
	//
	// Use SetArchiveAndMuteNewNoncontactPeers and GetArchiveAndMuteNewNoncontactPeers helpers.
	ArchiveAndMuteNewNoncontactPeers bool `tl:"archive_and_mute_new_noncontact_peers"`
}

// GlobalPrivacySettingsTypeID is TL type id of GlobalPrivacySettings.
const GlobalPrivacySettingsTypeID = 0xbea2f424

func (g *GlobalPrivacySettings) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GlobalPrivacySettings) String() string {
	if g == nil {
		return "GlobalPrivacySettings(nil)"
	}
	type Alias GlobalPrivacySettings
	return fmt.Sprintf("GlobalPrivacySettings%+v", Alias(*g))
}

// FillFrom fills GlobalPrivacySettings from given interface.
func (g *GlobalPrivacySettings) FillFrom(from interface {
	GetArchiveAndMuteNewNoncontactPeers() (value bool, ok bool)
}) {
	if val, ok := from.GetArchiveAndMuteNewNoncontactPeers(); ok {
		g.ArchiveAndMuteNewNoncontactPeers = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *GlobalPrivacySettings) TypeID() uint32 {
	return GlobalPrivacySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (g *GlobalPrivacySettings) TypeName() string {
	return "globalPrivacySettings"
}

// Encode implements bin.Encoder.
func (g *GlobalPrivacySettings) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#bea2f424 as nil")
	}
	b.PutID(GlobalPrivacySettingsTypeID)
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode globalPrivacySettings#bea2f424: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutBool(g.ArchiveAndMuteNewNoncontactPeers)
	}
	return nil
}

// SetArchiveAndMuteNewNoncontactPeers sets value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) SetArchiveAndMuteNewNoncontactPeers(value bool) {
	g.Flags.Set(0)
	g.ArchiveAndMuteNewNoncontactPeers = value
}

// GetArchiveAndMuteNewNoncontactPeers returns value of ArchiveAndMuteNewNoncontactPeers conditional field and
// boolean which is true if field was set.
func (g *GlobalPrivacySettings) GetArchiveAndMuteNewNoncontactPeers() (value bool, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.ArchiveAndMuteNewNoncontactPeers, true
}

// Decode implements bin.Decoder.
func (g *GlobalPrivacySettings) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#bea2f424 to nil")
	}
	if err := b.ConsumeID(GlobalPrivacySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: field flags: %w", err)
		}
	}
	if g.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: field archive_and_mute_new_noncontact_peers: %w", err)
		}
		g.ArchiveAndMuteNewNoncontactPeers = value
	}
	return nil
}

// Ensuring interfaces in compile-time for GlobalPrivacySettings.
var (
	_ bin.Encoder = &GlobalPrivacySettings{}
	_ bin.Decoder = &GlobalPrivacySettings{}
)

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

// PeerSettings represents TL type `peerSettings#733f2961`.
// Peer settings
//
// See https://core.telegram.org/constructor/peerSettings for reference.
type PeerSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether we can still report the user for spam
	ReportSpam bool
	// Whether we can add the user as contact
	AddContact bool
	// Whether we can block the user
	BlockContact bool
	// Whether we can share the user's contact
	ShareContact bool
	// Whether a special exception for contacts is needed
	NeedContactsException bool
	// Whether we can report a geogroup is irrelevant for this location
	ReportGeo bool
	// Whether this peer was automatically archived according to privacy settings¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/globalPrivacySettings
	Autoarchived bool
	// InviteMembers field of PeerSettings.
	InviteMembers bool
	// Distance in meters between us and this peer
	//
	// Use SetGeoDistance and GetGeoDistance helpers.
	GeoDistance int
}

// PeerSettingsTypeID is TL type id of PeerSettings.
const PeerSettingsTypeID = 0x733f2961

func (p *PeerSettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.ReportSpam == false) {
		return false
	}
	if !(p.AddContact == false) {
		return false
	}
	if !(p.BlockContact == false) {
		return false
	}
	if !(p.ShareContact == false) {
		return false
	}
	if !(p.NeedContactsException == false) {
		return false
	}
	if !(p.ReportGeo == false) {
		return false
	}
	if !(p.Autoarchived == false) {
		return false
	}
	if !(p.InviteMembers == false) {
		return false
	}
	if !(p.GeoDistance == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PeerSettings) String() string {
	if p == nil {
		return "PeerSettings(nil)"
	}
	type Alias PeerSettings
	return fmt.Sprintf("PeerSettings%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PeerSettings) TypeID() uint32 {
	return PeerSettingsTypeID
}

// Encode implements bin.Encoder.
func (p *PeerSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode peerSettings#733f2961 as nil")
	}
	b.PutID(PeerSettingsTypeID)
	if !(p.ReportSpam == false) {
		p.Flags.Set(0)
	}
	if !(p.AddContact == false) {
		p.Flags.Set(1)
	}
	if !(p.BlockContact == false) {
		p.Flags.Set(2)
	}
	if !(p.ShareContact == false) {
		p.Flags.Set(3)
	}
	if !(p.NeedContactsException == false) {
		p.Flags.Set(4)
	}
	if !(p.ReportGeo == false) {
		p.Flags.Set(5)
	}
	if !(p.Autoarchived == false) {
		p.Flags.Set(7)
	}
	if !(p.InviteMembers == false) {
		p.Flags.Set(8)
	}
	if !(p.GeoDistance == 0) {
		p.Flags.Set(6)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode peerSettings#733f2961: field flags: %w", err)
	}
	if p.Flags.Has(6) {
		b.PutInt(p.GeoDistance)
	}
	return nil
}

// SetReportSpam sets value of ReportSpam conditional field.
func (p *PeerSettings) SetReportSpam(value bool) {
	if value {
		p.Flags.Set(0)
		p.ReportSpam = true
	} else {
		p.Flags.Unset(0)
		p.ReportSpam = false
	}
}

// GetReportSpam returns value of ReportSpam conditional field.
func (p *PeerSettings) GetReportSpam() (value bool) {
	return p.Flags.Has(0)
}

// SetAddContact sets value of AddContact conditional field.
func (p *PeerSettings) SetAddContact(value bool) {
	if value {
		p.Flags.Set(1)
		p.AddContact = true
	} else {
		p.Flags.Unset(1)
		p.AddContact = false
	}
}

// GetAddContact returns value of AddContact conditional field.
func (p *PeerSettings) GetAddContact() (value bool) {
	return p.Flags.Has(1)
}

// SetBlockContact sets value of BlockContact conditional field.
func (p *PeerSettings) SetBlockContact(value bool) {
	if value {
		p.Flags.Set(2)
		p.BlockContact = true
	} else {
		p.Flags.Unset(2)
		p.BlockContact = false
	}
}

// GetBlockContact returns value of BlockContact conditional field.
func (p *PeerSettings) GetBlockContact() (value bool) {
	return p.Flags.Has(2)
}

// SetShareContact sets value of ShareContact conditional field.
func (p *PeerSettings) SetShareContact(value bool) {
	if value {
		p.Flags.Set(3)
		p.ShareContact = true
	} else {
		p.Flags.Unset(3)
		p.ShareContact = false
	}
}

// GetShareContact returns value of ShareContact conditional field.
func (p *PeerSettings) GetShareContact() (value bool) {
	return p.Flags.Has(3)
}

// SetNeedContactsException sets value of NeedContactsException conditional field.
func (p *PeerSettings) SetNeedContactsException(value bool) {
	if value {
		p.Flags.Set(4)
		p.NeedContactsException = true
	} else {
		p.Flags.Unset(4)
		p.NeedContactsException = false
	}
}

// GetNeedContactsException returns value of NeedContactsException conditional field.
func (p *PeerSettings) GetNeedContactsException() (value bool) {
	return p.Flags.Has(4)
}

// SetReportGeo sets value of ReportGeo conditional field.
func (p *PeerSettings) SetReportGeo(value bool) {
	if value {
		p.Flags.Set(5)
		p.ReportGeo = true
	} else {
		p.Flags.Unset(5)
		p.ReportGeo = false
	}
}

// GetReportGeo returns value of ReportGeo conditional field.
func (p *PeerSettings) GetReportGeo() (value bool) {
	return p.Flags.Has(5)
}

// SetAutoarchived sets value of Autoarchived conditional field.
func (p *PeerSettings) SetAutoarchived(value bool) {
	if value {
		p.Flags.Set(7)
		p.Autoarchived = true
	} else {
		p.Flags.Unset(7)
		p.Autoarchived = false
	}
}

// GetAutoarchived returns value of Autoarchived conditional field.
func (p *PeerSettings) GetAutoarchived() (value bool) {
	return p.Flags.Has(7)
}

// SetInviteMembers sets value of InviteMembers conditional field.
func (p *PeerSettings) SetInviteMembers(value bool) {
	if value {
		p.Flags.Set(8)
		p.InviteMembers = true
	} else {
		p.Flags.Unset(8)
		p.InviteMembers = false
	}
}

// GetInviteMembers returns value of InviteMembers conditional field.
func (p *PeerSettings) GetInviteMembers() (value bool) {
	return p.Flags.Has(8)
}

// SetGeoDistance sets value of GeoDistance conditional field.
func (p *PeerSettings) SetGeoDistance(value int) {
	p.Flags.Set(6)
	p.GeoDistance = value
}

// GetGeoDistance returns value of GeoDistance conditional field and
// boolean which is true if field was set.
func (p *PeerSettings) GetGeoDistance() (value int, ok bool) {
	if !p.Flags.Has(6) {
		return value, false
	}
	return p.GeoDistance, true
}

// Decode implements bin.Decoder.
func (p *PeerSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode peerSettings#733f2961 to nil")
	}
	if err := b.ConsumeID(PeerSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode peerSettings#733f2961: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode peerSettings#733f2961: field flags: %w", err)
		}
	}
	p.ReportSpam = p.Flags.Has(0)
	p.AddContact = p.Flags.Has(1)
	p.BlockContact = p.Flags.Has(2)
	p.ShareContact = p.Flags.Has(3)
	p.NeedContactsException = p.Flags.Has(4)
	p.ReportGeo = p.Flags.Has(5)
	p.Autoarchived = p.Flags.Has(7)
	p.InviteMembers = p.Flags.Has(8)
	if p.Flags.Has(6) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode peerSettings#733f2961: field geo_distance: %w", err)
		}
		p.GeoDistance = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PeerSettings.
var (
	_ bin.Encoder = &PeerSettings{}
	_ bin.Decoder = &PeerSettings{}
)

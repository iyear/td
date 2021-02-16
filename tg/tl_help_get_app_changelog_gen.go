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

// HelpGetAppChangelogRequest represents TL type `help.getAppChangelog#9010ef6f`.
// Get changelog of current app.
// Typically, an updates¹ constructor will be returned, containing one or more updateServiceNotification² updates with app-specific changelogs.
//
// Links:
//  1) https://core.telegram.org/constructor/updates
//  2) https://core.telegram.org/constructor/updateServiceNotification
//
// See https://core.telegram.org/method/help.getAppChangelog for reference.
type HelpGetAppChangelogRequest struct {
	// Previous app version
	PrevAppVersion string
}

// HelpGetAppChangelogRequestTypeID is TL type id of HelpGetAppChangelogRequest.
const HelpGetAppChangelogRequestTypeID = 0x9010ef6f

func (g *HelpGetAppChangelogRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.PrevAppVersion == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetAppChangelogRequest) String() string {
	if g == nil {
		return "HelpGetAppChangelogRequest(nil)"
	}
	type Alias HelpGetAppChangelogRequest
	return fmt.Sprintf("HelpGetAppChangelogRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetAppChangelogRequest) TypeID() uint32 {
	return HelpGetAppChangelogRequestTypeID
}

// Encode implements bin.Encoder.
func (g *HelpGetAppChangelogRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppChangelog#9010ef6f as nil")
	}
	b.PutID(HelpGetAppChangelogRequestTypeID)
	b.PutString(g.PrevAppVersion)
	return nil
}

// GetPrevAppVersion returns value of PrevAppVersion field.
func (g *HelpGetAppChangelogRequest) GetPrevAppVersion() (value string) {
	return g.PrevAppVersion
}

// Decode implements bin.Decoder.
func (g *HelpGetAppChangelogRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppChangelog#9010ef6f to nil")
	}
	if err := b.ConsumeID(HelpGetAppChangelogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getAppChangelog#9010ef6f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getAppChangelog#9010ef6f: field prev_app_version: %w", err)
		}
		g.PrevAppVersion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetAppChangelogRequest.
var (
	_ bin.Encoder = &HelpGetAppChangelogRequest{}
	_ bin.Decoder = &HelpGetAppChangelogRequest{}
)

// HelpGetAppChangelog invokes method help.getAppChangelog#9010ef6f returning error if any.
// Get changelog of current app.
// Typically, an updates¹ constructor will be returned, containing one or more updateServiceNotification² updates with app-specific changelogs.
//
// Links:
//  1) https://core.telegram.org/constructor/updates
//  2) https://core.telegram.org/constructor/updateServiceNotification
//
// See https://core.telegram.org/method/help.getAppChangelog for reference.
func (c *Client) HelpGetAppChangelog(ctx context.Context, prevappversion string) (UpdatesClass, error) {
	var result UpdatesBox

	request := &HelpGetAppChangelogRequest{
		PrevAppVersion: prevappversion,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}

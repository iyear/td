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

// StatsMessageStats represents TL type `stats.messageStats#8999f295`.
// Message statistics
//
// See https://core.telegram.org/constructor/stats.messageStats for reference.
type StatsMessageStats struct {
	// Message view graph
	ViewsGraph StatsGraphClass
}

// StatsMessageStatsTypeID is TL type id of StatsMessageStats.
const StatsMessageStatsTypeID = 0x8999f295

func (m *StatsMessageStats) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ViewsGraph == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *StatsMessageStats) String() string {
	if m == nil {
		return "StatsMessageStats(nil)"
	}
	type Alias StatsMessageStats
	return fmt.Sprintf("StatsMessageStats%+v", Alias(*m))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (m *StatsMessageStats) TypeID() uint32 {
	return StatsMessageStatsTypeID
}

// Encode implements bin.Encoder.
func (m *StatsMessageStats) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode stats.messageStats#8999f295 as nil")
	}
	b.PutID(StatsMessageStatsTypeID)
	if m.ViewsGraph == nil {
		return fmt.Errorf("unable to encode stats.messageStats#8999f295: field views_graph is nil")
	}
	if err := m.ViewsGraph.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.messageStats#8999f295: field views_graph: %w", err)
	}
	return nil
}

// GetViewsGraph returns value of ViewsGraph field.
func (m *StatsMessageStats) GetViewsGraph() (value StatsGraphClass) {
	return m.ViewsGraph
}

// Decode implements bin.Decoder.
func (m *StatsMessageStats) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode stats.messageStats#8999f295 to nil")
	}
	if err := b.ConsumeID(StatsMessageStatsTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.messageStats#8999f295: %w", err)
	}
	{
		value, err := DecodeStatsGraph(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.messageStats#8999f295: field views_graph: %w", err)
		}
		m.ViewsGraph = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsMessageStats.
var (
	_ bin.Encoder = &StatsMessageStats{}
	_ bin.Decoder = &StatsMessageStats{}
)

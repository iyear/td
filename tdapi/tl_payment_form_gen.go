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

// PaymentForm represents TL type `paymentForm#572da1e6`.
type PaymentForm struct {
	// The payment form identifier
	ID int64
	// Full information of the invoice
	Invoice Invoice
	// Payment form URL
	URL string
	// User identifier of the seller bot
	SellerBotUserID int64
	// User identifier of the payment provider bot
	PaymentsProviderUserID int64
	// Information about the payment provider, if available, to support it natively without
	// the need for opening the URL; may be null
	PaymentsProvider PaymentsProviderStripe
	// Saved server-side order information; may be null
	SavedOrderInfo OrderInfo
	// Information about saved card credentials; may be null
	SavedCredentials SavedCredentials
	// True, if the user can choose to save credentials
	CanSaveCredentials bool
	// True, if the user will be able to save credentials protected by a password they set up
	NeedPassword bool
}

// PaymentFormTypeID is TL type id of PaymentForm.
const PaymentFormTypeID = 0x572da1e6

// Ensuring interfaces in compile-time for PaymentForm.
var (
	_ bin.Encoder     = &PaymentForm{}
	_ bin.Decoder     = &PaymentForm{}
	_ bin.BareEncoder = &PaymentForm{}
	_ bin.BareDecoder = &PaymentForm{}
)

func (p *PaymentForm) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.Invoice.Zero()) {
		return false
	}
	if !(p.URL == "") {
		return false
	}
	if !(p.SellerBotUserID == 0) {
		return false
	}
	if !(p.PaymentsProviderUserID == 0) {
		return false
	}
	if !(p.PaymentsProvider.Zero()) {
		return false
	}
	if !(p.SavedOrderInfo.Zero()) {
		return false
	}
	if !(p.SavedCredentials.Zero()) {
		return false
	}
	if !(p.CanSaveCredentials == false) {
		return false
	}
	if !(p.NeedPassword == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentForm) String() string {
	if p == nil {
		return "PaymentForm(nil)"
	}
	type Alias PaymentForm
	return fmt.Sprintf("PaymentForm%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentForm) TypeID() uint32 {
	return PaymentFormTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentForm) TypeName() string {
	return "paymentForm"
}

// TypeInfo returns info about TL type.
func (p *PaymentForm) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentForm",
		ID:   PaymentFormTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "SellerBotUserID",
			SchemaName: "seller_bot_user_id",
		},
		{
			Name:       "PaymentsProviderUserID",
			SchemaName: "payments_provider_user_id",
		},
		{
			Name:       "PaymentsProvider",
			SchemaName: "payments_provider",
		},
		{
			Name:       "SavedOrderInfo",
			SchemaName: "saved_order_info",
		},
		{
			Name:       "SavedCredentials",
			SchemaName: "saved_credentials",
		},
		{
			Name:       "CanSaveCredentials",
			SchemaName: "can_save_credentials",
		},
		{
			Name:       "NeedPassword",
			SchemaName: "need_password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentForm) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentForm#572da1e6 as nil")
	}
	b.PutID(PaymentFormTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentForm) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentForm#572da1e6 as nil")
	}
	b.PutLong(p.ID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field invoice: %w", err)
	}
	b.PutString(p.URL)
	b.PutInt53(p.SellerBotUserID)
	b.PutInt53(p.PaymentsProviderUserID)
	if err := p.PaymentsProvider.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field payments_provider: %w", err)
	}
	if err := p.SavedOrderInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field saved_order_info: %w", err)
	}
	if err := p.SavedCredentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field saved_credentials: %w", err)
	}
	b.PutBool(p.CanSaveCredentials)
	b.PutBool(p.NeedPassword)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentForm) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentForm#572da1e6 to nil")
	}
	if err := b.ConsumeID(PaymentFormTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentForm#572da1e6: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentForm) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentForm#572da1e6 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field id: %w", err)
		}
		p.ID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field invoice: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field url: %w", err)
		}
		p.URL = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field seller_bot_user_id: %w", err)
		}
		p.SellerBotUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field payments_provider_user_id: %w", err)
		}
		p.PaymentsProviderUserID = value
	}
	{
		if err := p.PaymentsProvider.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field payments_provider: %w", err)
		}
	}
	{
		if err := p.SavedOrderInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field saved_order_info: %w", err)
		}
	}
	{
		if err := p.SavedCredentials.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field saved_credentials: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field can_save_credentials: %w", err)
		}
		p.CanSaveCredentials = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode paymentForm#572da1e6: field need_password: %w", err)
		}
		p.NeedPassword = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentForm) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentForm#572da1e6 as nil")
	}
	b.ObjStart()
	b.PutID("paymentForm")
	b.FieldStart("id")
	b.PutLong(p.ID)
	b.FieldStart("invoice")
	if err := p.Invoice.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field invoice: %w", err)
	}
	b.FieldStart("url")
	b.PutString(p.URL)
	b.FieldStart("seller_bot_user_id")
	b.PutInt53(p.SellerBotUserID)
	b.FieldStart("payments_provider_user_id")
	b.PutInt53(p.PaymentsProviderUserID)
	b.FieldStart("payments_provider")
	if err := p.PaymentsProvider.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field payments_provider: %w", err)
	}
	b.FieldStart("saved_order_info")
	if err := p.SavedOrderInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field saved_order_info: %w", err)
	}
	b.FieldStart("saved_credentials")
	if err := p.SavedCredentials.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentForm#572da1e6: field saved_credentials: %w", err)
	}
	b.FieldStart("can_save_credentials")
	b.PutBool(p.CanSaveCredentials)
	b.FieldStart("need_password")
	b.PutBool(p.NeedPassword)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentForm) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentForm#572da1e6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentForm"); err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field id: %w", err)
			}
			p.ID = value
		case "invoice":
			if err := p.Invoice.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field invoice: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field url: %w", err)
			}
			p.URL = value
		case "seller_bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field seller_bot_user_id: %w", err)
			}
			p.SellerBotUserID = value
		case "payments_provider_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field payments_provider_user_id: %w", err)
			}
			p.PaymentsProviderUserID = value
		case "payments_provider":
			if err := p.PaymentsProvider.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field payments_provider: %w", err)
			}
		case "saved_order_info":
			if err := p.SavedOrderInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field saved_order_info: %w", err)
			}
		case "saved_credentials":
			if err := p.SavedCredentials.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field saved_credentials: %w", err)
			}
		case "can_save_credentials":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field can_save_credentials: %w", err)
			}
			p.CanSaveCredentials = value
		case "need_password":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode paymentForm#572da1e6: field need_password: %w", err)
			}
			p.NeedPassword = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (p *PaymentForm) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// GetInvoice returns value of Invoice field.
func (p *PaymentForm) GetInvoice() (value Invoice) {
	if p == nil {
		return
	}
	return p.Invoice
}

// GetURL returns value of URL field.
func (p *PaymentForm) GetURL() (value string) {
	if p == nil {
		return
	}
	return p.URL
}

// GetSellerBotUserID returns value of SellerBotUserID field.
func (p *PaymentForm) GetSellerBotUserID() (value int64) {
	if p == nil {
		return
	}
	return p.SellerBotUserID
}

// GetPaymentsProviderUserID returns value of PaymentsProviderUserID field.
func (p *PaymentForm) GetPaymentsProviderUserID() (value int64) {
	if p == nil {
		return
	}
	return p.PaymentsProviderUserID
}

// GetPaymentsProvider returns value of PaymentsProvider field.
func (p *PaymentForm) GetPaymentsProvider() (value PaymentsProviderStripe) {
	if p == nil {
		return
	}
	return p.PaymentsProvider
}

// GetSavedOrderInfo returns value of SavedOrderInfo field.
func (p *PaymentForm) GetSavedOrderInfo() (value OrderInfo) {
	if p == nil {
		return
	}
	return p.SavedOrderInfo
}

// GetSavedCredentials returns value of SavedCredentials field.
func (p *PaymentForm) GetSavedCredentials() (value SavedCredentials) {
	if p == nil {
		return
	}
	return p.SavedCredentials
}

// GetCanSaveCredentials returns value of CanSaveCredentials field.
func (p *PaymentForm) GetCanSaveCredentials() (value bool) {
	if p == nil {
		return
	}
	return p.CanSaveCredentials
}

// GetNeedPassword returns value of NeedPassword field.
func (p *PaymentForm) GetNeedPassword() (value bool) {
	if p == nil {
		return
	}
	return p.NeedPassword
}

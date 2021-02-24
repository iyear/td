// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// ReqDHParamsRequest represents TL type `req_DH_params#d712e4be`.
type ReqDHParamsRequest struct {
	// Nonce field of ReqDHParamsRequest.
	Nonce bin.Int128 `tl:"nonce"`
	// ServerNonce field of ReqDHParamsRequest.
	ServerNonce bin.Int128 `tl:"server_nonce"`
	// P field of ReqDHParamsRequest.
	P []byte `tl:"p"`
	// Q field of ReqDHParamsRequest.
	Q []byte `tl:"q"`
	// PublicKeyFingerprint field of ReqDHParamsRequest.
	PublicKeyFingerprint int64 `tl:"public_key_fingerprint"`
	// EncryptedData field of ReqDHParamsRequest.
	EncryptedData []byte `tl:"encrypted_data"`
}

// ReqDHParamsRequestTypeID is TL type id of ReqDHParamsRequest.
const ReqDHParamsRequestTypeID = 0xd712e4be

func (r *ReqDHParamsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}
	if !(r.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(r.P == nil) {
		return false
	}
	if !(r.Q == nil) {
		return false
	}
	if !(r.PublicKeyFingerprint == 0) {
		return false
	}
	if !(r.EncryptedData == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReqDHParamsRequest) String() string {
	if r == nil {
		return "ReqDHParamsRequest(nil)"
	}
	type Alias ReqDHParamsRequest
	return fmt.Sprintf("ReqDHParamsRequest%+v", Alias(*r))
}

// FillFrom fills ReqDHParamsRequest from given interface.
func (r *ReqDHParamsRequest) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetP() (value []byte)
	GetQ() (value []byte)
	GetPublicKeyFingerprint() (value int64)
	GetEncryptedData() (value []byte)
}) {
	r.Nonce = from.GetNonce()
	r.ServerNonce = from.GetServerNonce()
	r.P = from.GetP()
	r.Q = from.GetQ()
	r.PublicKeyFingerprint = from.GetPublicKeyFingerprint()
	r.EncryptedData = from.GetEncryptedData()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReqDHParamsRequest) TypeID() uint32 {
	return ReqDHParamsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (r *ReqDHParamsRequest) TypeName() string {
	return "req_DH_params"
}

// Encode implements bin.Encoder.
func (r *ReqDHParamsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_DH_params#d712e4be as nil")
	}
	b.PutID(ReqDHParamsRequestTypeID)
	b.PutInt128(r.Nonce)
	b.PutInt128(r.ServerNonce)
	b.PutBytes(r.P)
	b.PutBytes(r.Q)
	b.PutLong(r.PublicKeyFingerprint)
	b.PutBytes(r.EncryptedData)
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ReqDHParamsRequest) GetNonce() (value bin.Int128) {
	return r.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (r *ReqDHParamsRequest) GetServerNonce() (value bin.Int128) {
	return r.ServerNonce
}

// GetP returns value of P field.
func (r *ReqDHParamsRequest) GetP() (value []byte) {
	return r.P
}

// GetQ returns value of Q field.
func (r *ReqDHParamsRequest) GetQ() (value []byte) {
	return r.Q
}

// GetPublicKeyFingerprint returns value of PublicKeyFingerprint field.
func (r *ReqDHParamsRequest) GetPublicKeyFingerprint() (value int64) {
	return r.PublicKeyFingerprint
}

// GetEncryptedData returns value of EncryptedData field.
func (r *ReqDHParamsRequest) GetEncryptedData() (value []byte) {
	return r.EncryptedData
}

// Decode implements bin.Decoder.
func (r *ReqDHParamsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_DH_params#d712e4be to nil")
	}
	if err := b.ConsumeID(ReqDHParamsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode req_DH_params#d712e4be: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field nonce: %w", err)
		}
		r.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field server_nonce: %w", err)
		}
		r.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field p: %w", err)
		}
		r.P = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field q: %w", err)
		}
		r.Q = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field public_key_fingerprint: %w", err)
		}
		r.PublicKeyFingerprint = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field encrypted_data: %w", err)
		}
		r.EncryptedData = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ReqDHParamsRequest.
var (
	_ bin.Encoder = &ReqDHParamsRequest{}
	_ bin.Decoder = &ReqDHParamsRequest{}
)

// ReqDHParams invokes method req_DH_params#d712e4be returning error if any.
func (c *Client) ReqDHParams(ctx context.Context, request *ReqDHParamsRequest) (ServerDHParamsClass, error) {
	var result ServerDHParamsBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Server_DH_Params, nil
}

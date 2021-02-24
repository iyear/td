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

// SecureValueTypePersonalDetails represents TL type `secureValueTypePersonalDetails#9d2a81e3`.
// Personal details
//
// See https://core.telegram.org/constructor/secureValueTypePersonalDetails for reference.
type SecureValueTypePersonalDetails struct {
}

// SecureValueTypePersonalDetailsTypeID is TL type id of SecureValueTypePersonalDetails.
const SecureValueTypePersonalDetailsTypeID = 0x9d2a81e3

func (s *SecureValueTypePersonalDetails) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypePersonalDetails) String() string {
	if s == nil {
		return "SecureValueTypePersonalDetails(nil)"
	}
	type Alias SecureValueTypePersonalDetails
	return fmt.Sprintf("SecureValueTypePersonalDetails%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypePersonalDetails) TypeID() uint32 {
	return SecureValueTypePersonalDetailsTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypePersonalDetails) TypeName() string {
	return "secureValueTypePersonalDetails"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypePersonalDetails) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypePersonalDetails#9d2a81e3 as nil")
	}
	b.PutID(SecureValueTypePersonalDetailsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypePersonalDetails) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypePersonalDetails#9d2a81e3 to nil")
	}
	if err := b.ConsumeID(SecureValueTypePersonalDetailsTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypePersonalDetails#9d2a81e3: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypePersonalDetails) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypePersonalDetails.
var (
	_ bin.Encoder = &SecureValueTypePersonalDetails{}
	_ bin.Decoder = &SecureValueTypePersonalDetails{}

	_ SecureValueTypeClass = &SecureValueTypePersonalDetails{}
)

// SecureValueTypePassport represents TL type `secureValueTypePassport#3dac6a00`.
// Passport
//
// See https://core.telegram.org/constructor/secureValueTypePassport for reference.
type SecureValueTypePassport struct {
}

// SecureValueTypePassportTypeID is TL type id of SecureValueTypePassport.
const SecureValueTypePassportTypeID = 0x3dac6a00

func (s *SecureValueTypePassport) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypePassport) String() string {
	if s == nil {
		return "SecureValueTypePassport(nil)"
	}
	type Alias SecureValueTypePassport
	return fmt.Sprintf("SecureValueTypePassport%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypePassport) TypeID() uint32 {
	return SecureValueTypePassportTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypePassport) TypeName() string {
	return "secureValueTypePassport"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypePassport) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypePassport#3dac6a00 as nil")
	}
	b.PutID(SecureValueTypePassportTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypePassport) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypePassport#3dac6a00 to nil")
	}
	if err := b.ConsumeID(SecureValueTypePassportTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypePassport#3dac6a00: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypePassport) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypePassport.
var (
	_ bin.Encoder = &SecureValueTypePassport{}
	_ bin.Decoder = &SecureValueTypePassport{}

	_ SecureValueTypeClass = &SecureValueTypePassport{}
)

// SecureValueTypeDriverLicense represents TL type `secureValueTypeDriverLicense#6e425c4`.
// Driver's license
//
// See https://core.telegram.org/constructor/secureValueTypeDriverLicense for reference.
type SecureValueTypeDriverLicense struct {
}

// SecureValueTypeDriverLicenseTypeID is TL type id of SecureValueTypeDriverLicense.
const SecureValueTypeDriverLicenseTypeID = 0x6e425c4

func (s *SecureValueTypeDriverLicense) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeDriverLicense) String() string {
	if s == nil {
		return "SecureValueTypeDriverLicense(nil)"
	}
	type Alias SecureValueTypeDriverLicense
	return fmt.Sprintf("SecureValueTypeDriverLicense%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeDriverLicense) TypeID() uint32 {
	return SecureValueTypeDriverLicenseTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeDriverLicense) TypeName() string {
	return "secureValueTypeDriverLicense"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeDriverLicense) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeDriverLicense#6e425c4 as nil")
	}
	b.PutID(SecureValueTypeDriverLicenseTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeDriverLicense) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeDriverLicense#6e425c4 to nil")
	}
	if err := b.ConsumeID(SecureValueTypeDriverLicenseTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeDriverLicense#6e425c4: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeDriverLicense) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeDriverLicense.
var (
	_ bin.Encoder = &SecureValueTypeDriverLicense{}
	_ bin.Decoder = &SecureValueTypeDriverLicense{}

	_ SecureValueTypeClass = &SecureValueTypeDriverLicense{}
)

// SecureValueTypeIdentityCard represents TL type `secureValueTypeIdentityCard#a0d0744b`.
// Identity card
//
// See https://core.telegram.org/constructor/secureValueTypeIdentityCard for reference.
type SecureValueTypeIdentityCard struct {
}

// SecureValueTypeIdentityCardTypeID is TL type id of SecureValueTypeIdentityCard.
const SecureValueTypeIdentityCardTypeID = 0xa0d0744b

func (s *SecureValueTypeIdentityCard) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeIdentityCard) String() string {
	if s == nil {
		return "SecureValueTypeIdentityCard(nil)"
	}
	type Alias SecureValueTypeIdentityCard
	return fmt.Sprintf("SecureValueTypeIdentityCard%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeIdentityCard) TypeID() uint32 {
	return SecureValueTypeIdentityCardTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeIdentityCard) TypeName() string {
	return "secureValueTypeIdentityCard"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeIdentityCard) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeIdentityCard#a0d0744b as nil")
	}
	b.PutID(SecureValueTypeIdentityCardTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeIdentityCard) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeIdentityCard#a0d0744b to nil")
	}
	if err := b.ConsumeID(SecureValueTypeIdentityCardTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeIdentityCard#a0d0744b: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeIdentityCard) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeIdentityCard.
var (
	_ bin.Encoder = &SecureValueTypeIdentityCard{}
	_ bin.Decoder = &SecureValueTypeIdentityCard{}

	_ SecureValueTypeClass = &SecureValueTypeIdentityCard{}
)

// SecureValueTypeInternalPassport represents TL type `secureValueTypeInternalPassport#99a48f23`.
// Internal passport¹
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/secureValueTypeInternalPassport for reference.
type SecureValueTypeInternalPassport struct {
}

// SecureValueTypeInternalPassportTypeID is TL type id of SecureValueTypeInternalPassport.
const SecureValueTypeInternalPassportTypeID = 0x99a48f23

func (s *SecureValueTypeInternalPassport) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeInternalPassport) String() string {
	if s == nil {
		return "SecureValueTypeInternalPassport(nil)"
	}
	type Alias SecureValueTypeInternalPassport
	return fmt.Sprintf("SecureValueTypeInternalPassport%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeInternalPassport) TypeID() uint32 {
	return SecureValueTypeInternalPassportTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeInternalPassport) TypeName() string {
	return "secureValueTypeInternalPassport"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeInternalPassport) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeInternalPassport#99a48f23 as nil")
	}
	b.PutID(SecureValueTypeInternalPassportTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeInternalPassport) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeInternalPassport#99a48f23 to nil")
	}
	if err := b.ConsumeID(SecureValueTypeInternalPassportTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeInternalPassport#99a48f23: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeInternalPassport) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeInternalPassport.
var (
	_ bin.Encoder = &SecureValueTypeInternalPassport{}
	_ bin.Decoder = &SecureValueTypeInternalPassport{}

	_ SecureValueTypeClass = &SecureValueTypeInternalPassport{}
)

// SecureValueTypeAddress represents TL type `secureValueTypeAddress#cbe31e26`.
// Address
//
// See https://core.telegram.org/constructor/secureValueTypeAddress for reference.
type SecureValueTypeAddress struct {
}

// SecureValueTypeAddressTypeID is TL type id of SecureValueTypeAddress.
const SecureValueTypeAddressTypeID = 0xcbe31e26

func (s *SecureValueTypeAddress) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeAddress) String() string {
	if s == nil {
		return "SecureValueTypeAddress(nil)"
	}
	type Alias SecureValueTypeAddress
	return fmt.Sprintf("SecureValueTypeAddress%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeAddress) TypeID() uint32 {
	return SecureValueTypeAddressTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeAddress) TypeName() string {
	return "secureValueTypeAddress"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeAddress) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeAddress#cbe31e26 as nil")
	}
	b.PutID(SecureValueTypeAddressTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeAddress) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeAddress#cbe31e26 to nil")
	}
	if err := b.ConsumeID(SecureValueTypeAddressTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeAddress#cbe31e26: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeAddress) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeAddress.
var (
	_ bin.Encoder = &SecureValueTypeAddress{}
	_ bin.Decoder = &SecureValueTypeAddress{}

	_ SecureValueTypeClass = &SecureValueTypeAddress{}
)

// SecureValueTypeUtilityBill represents TL type `secureValueTypeUtilityBill#fc36954e`.
// Utility bill
//
// See https://core.telegram.org/constructor/secureValueTypeUtilityBill for reference.
type SecureValueTypeUtilityBill struct {
}

// SecureValueTypeUtilityBillTypeID is TL type id of SecureValueTypeUtilityBill.
const SecureValueTypeUtilityBillTypeID = 0xfc36954e

func (s *SecureValueTypeUtilityBill) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeUtilityBill) String() string {
	if s == nil {
		return "SecureValueTypeUtilityBill(nil)"
	}
	type Alias SecureValueTypeUtilityBill
	return fmt.Sprintf("SecureValueTypeUtilityBill%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeUtilityBill) TypeID() uint32 {
	return SecureValueTypeUtilityBillTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeUtilityBill) TypeName() string {
	return "secureValueTypeUtilityBill"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeUtilityBill) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeUtilityBill#fc36954e as nil")
	}
	b.PutID(SecureValueTypeUtilityBillTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeUtilityBill) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeUtilityBill#fc36954e to nil")
	}
	if err := b.ConsumeID(SecureValueTypeUtilityBillTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeUtilityBill#fc36954e: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeUtilityBill) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeUtilityBill.
var (
	_ bin.Encoder = &SecureValueTypeUtilityBill{}
	_ bin.Decoder = &SecureValueTypeUtilityBill{}

	_ SecureValueTypeClass = &SecureValueTypeUtilityBill{}
)

// SecureValueTypeBankStatement represents TL type `secureValueTypeBankStatement#89137c0d`.
// Bank statement
//
// See https://core.telegram.org/constructor/secureValueTypeBankStatement for reference.
type SecureValueTypeBankStatement struct {
}

// SecureValueTypeBankStatementTypeID is TL type id of SecureValueTypeBankStatement.
const SecureValueTypeBankStatementTypeID = 0x89137c0d

func (s *SecureValueTypeBankStatement) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeBankStatement) String() string {
	if s == nil {
		return "SecureValueTypeBankStatement(nil)"
	}
	type Alias SecureValueTypeBankStatement
	return fmt.Sprintf("SecureValueTypeBankStatement%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeBankStatement) TypeID() uint32 {
	return SecureValueTypeBankStatementTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeBankStatement) TypeName() string {
	return "secureValueTypeBankStatement"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeBankStatement) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeBankStatement#89137c0d as nil")
	}
	b.PutID(SecureValueTypeBankStatementTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeBankStatement) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeBankStatement#89137c0d to nil")
	}
	if err := b.ConsumeID(SecureValueTypeBankStatementTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeBankStatement#89137c0d: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeBankStatement) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeBankStatement.
var (
	_ bin.Encoder = &SecureValueTypeBankStatement{}
	_ bin.Decoder = &SecureValueTypeBankStatement{}

	_ SecureValueTypeClass = &SecureValueTypeBankStatement{}
)

// SecureValueTypeRentalAgreement represents TL type `secureValueTypeRentalAgreement#8b883488`.
// Rental agreement
//
// See https://core.telegram.org/constructor/secureValueTypeRentalAgreement for reference.
type SecureValueTypeRentalAgreement struct {
}

// SecureValueTypeRentalAgreementTypeID is TL type id of SecureValueTypeRentalAgreement.
const SecureValueTypeRentalAgreementTypeID = 0x8b883488

func (s *SecureValueTypeRentalAgreement) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeRentalAgreement) String() string {
	if s == nil {
		return "SecureValueTypeRentalAgreement(nil)"
	}
	type Alias SecureValueTypeRentalAgreement
	return fmt.Sprintf("SecureValueTypeRentalAgreement%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeRentalAgreement) TypeID() uint32 {
	return SecureValueTypeRentalAgreementTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeRentalAgreement) TypeName() string {
	return "secureValueTypeRentalAgreement"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeRentalAgreement) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeRentalAgreement#8b883488 as nil")
	}
	b.PutID(SecureValueTypeRentalAgreementTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeRentalAgreement) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeRentalAgreement#8b883488 to nil")
	}
	if err := b.ConsumeID(SecureValueTypeRentalAgreementTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeRentalAgreement#8b883488: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeRentalAgreement) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeRentalAgreement.
var (
	_ bin.Encoder = &SecureValueTypeRentalAgreement{}
	_ bin.Decoder = &SecureValueTypeRentalAgreement{}

	_ SecureValueTypeClass = &SecureValueTypeRentalAgreement{}
)

// SecureValueTypePassportRegistration represents TL type `secureValueTypePassportRegistration#99e3806a`.
// Internal registration passport¹
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/constructor/secureValueTypePassportRegistration for reference.
type SecureValueTypePassportRegistration struct {
}

// SecureValueTypePassportRegistrationTypeID is TL type id of SecureValueTypePassportRegistration.
const SecureValueTypePassportRegistrationTypeID = 0x99e3806a

func (s *SecureValueTypePassportRegistration) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypePassportRegistration) String() string {
	if s == nil {
		return "SecureValueTypePassportRegistration(nil)"
	}
	type Alias SecureValueTypePassportRegistration
	return fmt.Sprintf("SecureValueTypePassportRegistration%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypePassportRegistration) TypeID() uint32 {
	return SecureValueTypePassportRegistrationTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypePassportRegistration) TypeName() string {
	return "secureValueTypePassportRegistration"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypePassportRegistration) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypePassportRegistration#99e3806a as nil")
	}
	b.PutID(SecureValueTypePassportRegistrationTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypePassportRegistration) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypePassportRegistration#99e3806a to nil")
	}
	if err := b.ConsumeID(SecureValueTypePassportRegistrationTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypePassportRegistration#99e3806a: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypePassportRegistration) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypePassportRegistration.
var (
	_ bin.Encoder = &SecureValueTypePassportRegistration{}
	_ bin.Decoder = &SecureValueTypePassportRegistration{}

	_ SecureValueTypeClass = &SecureValueTypePassportRegistration{}
)

// SecureValueTypeTemporaryRegistration represents TL type `secureValueTypeTemporaryRegistration#ea02ec33`.
// Temporary registration
//
// See https://core.telegram.org/constructor/secureValueTypeTemporaryRegistration for reference.
type SecureValueTypeTemporaryRegistration struct {
}

// SecureValueTypeTemporaryRegistrationTypeID is TL type id of SecureValueTypeTemporaryRegistration.
const SecureValueTypeTemporaryRegistrationTypeID = 0xea02ec33

func (s *SecureValueTypeTemporaryRegistration) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeTemporaryRegistration) String() string {
	if s == nil {
		return "SecureValueTypeTemporaryRegistration(nil)"
	}
	type Alias SecureValueTypeTemporaryRegistration
	return fmt.Sprintf("SecureValueTypeTemporaryRegistration%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeTemporaryRegistration) TypeID() uint32 {
	return SecureValueTypeTemporaryRegistrationTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeTemporaryRegistration) TypeName() string {
	return "secureValueTypeTemporaryRegistration"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeTemporaryRegistration) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeTemporaryRegistration#ea02ec33 as nil")
	}
	b.PutID(SecureValueTypeTemporaryRegistrationTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeTemporaryRegistration) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeTemporaryRegistration#ea02ec33 to nil")
	}
	if err := b.ConsumeID(SecureValueTypeTemporaryRegistrationTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeTemporaryRegistration#ea02ec33: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeTemporaryRegistration) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeTemporaryRegistration.
var (
	_ bin.Encoder = &SecureValueTypeTemporaryRegistration{}
	_ bin.Decoder = &SecureValueTypeTemporaryRegistration{}

	_ SecureValueTypeClass = &SecureValueTypeTemporaryRegistration{}
)

// SecureValueTypePhone represents TL type `secureValueTypePhone#b320aadb`.
// Phone
//
// See https://core.telegram.org/constructor/secureValueTypePhone for reference.
type SecureValueTypePhone struct {
}

// SecureValueTypePhoneTypeID is TL type id of SecureValueTypePhone.
const SecureValueTypePhoneTypeID = 0xb320aadb

func (s *SecureValueTypePhone) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypePhone) String() string {
	if s == nil {
		return "SecureValueTypePhone(nil)"
	}
	type Alias SecureValueTypePhone
	return fmt.Sprintf("SecureValueTypePhone%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypePhone) TypeID() uint32 {
	return SecureValueTypePhoneTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypePhone) TypeName() string {
	return "secureValueTypePhone"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypePhone) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypePhone#b320aadb as nil")
	}
	b.PutID(SecureValueTypePhoneTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypePhone) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypePhone#b320aadb to nil")
	}
	if err := b.ConsumeID(SecureValueTypePhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypePhone#b320aadb: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypePhone) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypePhone.
var (
	_ bin.Encoder = &SecureValueTypePhone{}
	_ bin.Decoder = &SecureValueTypePhone{}

	_ SecureValueTypeClass = &SecureValueTypePhone{}
)

// SecureValueTypeEmail represents TL type `secureValueTypeEmail#8e3ca7ee`.
// Email
//
// See https://core.telegram.org/constructor/secureValueTypeEmail for reference.
type SecureValueTypeEmail struct {
}

// SecureValueTypeEmailTypeID is TL type id of SecureValueTypeEmail.
const SecureValueTypeEmailTypeID = 0x8e3ca7ee

func (s *SecureValueTypeEmail) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValueTypeEmail) String() string {
	if s == nil {
		return "SecureValueTypeEmail(nil)"
	}
	type Alias SecureValueTypeEmail
	return fmt.Sprintf("SecureValueTypeEmail%+v", Alias(*s))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *SecureValueTypeEmail) TypeID() uint32 {
	return SecureValueTypeEmailTypeID
}

// TypeName returns name of type in TL schema.
func (s *SecureValueTypeEmail) TypeName() string {
	return "secureValueTypeEmail"
}

// Encode implements bin.Encoder.
func (s *SecureValueTypeEmail) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValueTypeEmail#8e3ca7ee as nil")
	}
	b.PutID(SecureValueTypeEmailTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureValueTypeEmail) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValueTypeEmail#8e3ca7ee to nil")
	}
	if err := b.ConsumeID(SecureValueTypeEmailTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValueTypeEmail#8e3ca7ee: %w", err)
	}
	return nil
}

// construct implements constructor of SecureValueTypeClass.
func (s SecureValueTypeEmail) construct() SecureValueTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureValueTypeEmail.
var (
	_ bin.Encoder = &SecureValueTypeEmail{}
	_ bin.Decoder = &SecureValueTypeEmail{}

	_ SecureValueTypeClass = &SecureValueTypeEmail{}
)

// SecureValueTypeClass represents SecureValueType generic type.
//
// See https://core.telegram.org/type/SecureValueType for reference.
//
// Example:
//  g, err := tg.DecodeSecureValueType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.SecureValueTypePersonalDetails: // secureValueTypePersonalDetails#9d2a81e3
//  case *tg.SecureValueTypePassport: // secureValueTypePassport#3dac6a00
//  case *tg.SecureValueTypeDriverLicense: // secureValueTypeDriverLicense#6e425c4
//  case *tg.SecureValueTypeIdentityCard: // secureValueTypeIdentityCard#a0d0744b
//  case *tg.SecureValueTypeInternalPassport: // secureValueTypeInternalPassport#99a48f23
//  case *tg.SecureValueTypeAddress: // secureValueTypeAddress#cbe31e26
//  case *tg.SecureValueTypeUtilityBill: // secureValueTypeUtilityBill#fc36954e
//  case *tg.SecureValueTypeBankStatement: // secureValueTypeBankStatement#89137c0d
//  case *tg.SecureValueTypeRentalAgreement: // secureValueTypeRentalAgreement#8b883488
//  case *tg.SecureValueTypePassportRegistration: // secureValueTypePassportRegistration#99e3806a
//  case *tg.SecureValueTypeTemporaryRegistration: // secureValueTypeTemporaryRegistration#ea02ec33
//  case *tg.SecureValueTypePhone: // secureValueTypePhone#b320aadb
//  case *tg.SecureValueTypeEmail: // secureValueTypeEmail#8e3ca7ee
//  default: panic(v)
//  }
type SecureValueTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecureValueTypeClass

	// TypeID returns type id in TL schema.
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeSecureValueType implements binary de-serialization for SecureValueTypeClass.
func DecodeSecureValueType(buf *bin.Buffer) (SecureValueTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureValueTypePersonalDetailsTypeID:
		// Decoding secureValueTypePersonalDetails#9d2a81e3.
		v := SecureValueTypePersonalDetails{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypePassportTypeID:
		// Decoding secureValueTypePassport#3dac6a00.
		v := SecureValueTypePassport{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeDriverLicenseTypeID:
		// Decoding secureValueTypeDriverLicense#6e425c4.
		v := SecureValueTypeDriverLicense{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeIdentityCardTypeID:
		// Decoding secureValueTypeIdentityCard#a0d0744b.
		v := SecureValueTypeIdentityCard{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeInternalPassportTypeID:
		// Decoding secureValueTypeInternalPassport#99a48f23.
		v := SecureValueTypeInternalPassport{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeAddressTypeID:
		// Decoding secureValueTypeAddress#cbe31e26.
		v := SecureValueTypeAddress{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeUtilityBillTypeID:
		// Decoding secureValueTypeUtilityBill#fc36954e.
		v := SecureValueTypeUtilityBill{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeBankStatementTypeID:
		// Decoding secureValueTypeBankStatement#89137c0d.
		v := SecureValueTypeBankStatement{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeRentalAgreementTypeID:
		// Decoding secureValueTypeRentalAgreement#8b883488.
		v := SecureValueTypeRentalAgreement{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypePassportRegistrationTypeID:
		// Decoding secureValueTypePassportRegistration#99e3806a.
		v := SecureValueTypePassportRegistration{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeTemporaryRegistrationTypeID:
		// Decoding secureValueTypeTemporaryRegistration#ea02ec33.
		v := SecureValueTypeTemporaryRegistration{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypePhoneTypeID:
		// Decoding secureValueTypePhone#b320aadb.
		v := SecureValueTypePhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	case SecureValueTypeEmailTypeID:
		// Decoding secureValueTypeEmail#8e3ca7ee.
		v := SecureValueTypeEmail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureValueTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// SecureValueType boxes the SecureValueTypeClass providing a helper.
type SecureValueTypeBox struct {
	SecureValueType SecureValueTypeClass
}

// Decode implements bin.Decoder for SecureValueTypeBox.
func (b *SecureValueTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode SecureValueTypeBox to nil")
	}
	v, err := DecodeSecureValueType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.SecureValueType = v
	return nil
}

// Encode implements bin.Encode for SecureValueTypeBox.
func (b *SecureValueTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.SecureValueType == nil {
		return fmt.Errorf("unable to encode SecureValueTypeClass as nil")
	}
	return b.SecureValueType.Encode(buf)
}

// SecureValueTypeClassSlice is adapter for slice of SecureValueTypeClass.
type SecureValueTypeClassSlice []SecureValueTypeClass

// First returns first element of slice (if exists).
func (s SecureValueTypeClassSlice) First() (v SecureValueTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureValueTypeClassSlice) Last() (v SecureValueTypeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureValueTypeClassSlice) PopFirst() (v SecureValueTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureValueTypeClassSlice) Pop() (v SecureValueTypeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

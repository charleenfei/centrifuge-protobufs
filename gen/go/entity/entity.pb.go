// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/entity.proto

package entitypb

import (
	fmt "fmt"
	_ "github.com/centrifuge/precise-proofs/proofs/proto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// EntityRelationship allows other identities to access the Entity document.
type EntityRelationship struct {
	// owner id of the Relationship
	OwnerIdentity []byte `protobuf:"bytes,1,opt,name=owner_identity,json=ownerIdentity,proto3" json:"owner_identity,omitempty"`
	// identifier for the entity document whose data should be accessed via this relationship
	EntityIdentifier []byte `protobuf:"bytes,2,opt,name=entity_identifier,json=entityIdentifier,proto3" json:"entity_identifier,omitempty"`
	// identity to whom access should be granted
	TargetIdentity       []byte   `protobuf:"bytes,3,opt,name=target_identity,json=targetIdentity,proto3" json:"target_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityRelationship) Reset()         { *m = EntityRelationship{} }
func (m *EntityRelationship) String() string { return proto.CompactTextString(m) }
func (*EntityRelationship) ProtoMessage()    {}
func (*EntityRelationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{0}
}

func (m *EntityRelationship) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityRelationship.Unmarshal(m, b)
}
func (m *EntityRelationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityRelationship.Marshal(b, m, deterministic)
}
func (m *EntityRelationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityRelationship.Merge(m, src)
}
func (m *EntityRelationship) XXX_Size() int {
	return xxx_messageInfo_EntityRelationship.Size(m)
}
func (m *EntityRelationship) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityRelationship.DiscardUnknown(m)
}

var xxx_messageInfo_EntityRelationship proto.InternalMessageInfo

func (m *EntityRelationship) GetOwnerIdentity() []byte {
	if m != nil {
		return m.OwnerIdentity
	}
	return nil
}

func (m *EntityRelationship) GetEntityIdentifier() []byte {
	if m != nil {
		return m.EntityIdentifier
	}
	return nil
}

func (m *EntityRelationship) GetTargetIdentity() []byte {
	if m != nil {
		return m.TargetIdentity
	}
	return nil
}

// EntityData is the default entity schema
type Entity struct {
	Identity  []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	LegalName string `protobuf:"bytes,2,opt,name=legal_name,json=legalName,proto3" json:"legal_name,omitempty"`
	// address
	Addresses []*Address `protobuf:"bytes,3,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// tax information
	PaymentDetails []*PaymentDetail `protobuf:"bytes,4,rep,name=payment_details,json=paymentDetails,proto3" json:"payment_details,omitempty"`
	// Entity contact list
	Contacts             []*Contact `protobuf:"bytes,5,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{1}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Entity) GetLegalName() string {
	if m != nil {
		return m.LegalName
	}
	return ""
}

func (m *Entity) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Entity) GetPaymentDetails() []*PaymentDetail {
	if m != nil {
		return m.PaymentDetails
	}
	return nil
}

func (m *Entity) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

type Address struct {
	IsMain               bool     `protobuf:"varint,1,opt,name=is_main,json=isMain,proto3" json:"is_main,omitempty"`
	IsRemitTo            bool     `protobuf:"varint,2,opt,name=is_remit_to,json=isRemitTo,proto3" json:"is_remit_to,omitempty"`
	IsShipTo             bool     `protobuf:"varint,3,opt,name=is_ship_to,json=isShipTo,proto3" json:"is_ship_to,omitempty"`
	IsPayTo              bool     `protobuf:"varint,4,opt,name=is_pay_to,json=isPayTo,proto3" json:"is_pay_to,omitempty"`
	Label                string   `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Zip                  string   `protobuf:"bytes,6,opt,name=zip,proto3" json:"zip,omitempty"`
	State                string   `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,8,opt,name=country,proto3" json:"country,omitempty"`
	AddressLine1         string   `protobuf:"bytes,9,opt,name=address_line1,json=addressLine1,proto3" json:"address_line1,omitempty"`
	AddressLine2         string   `protobuf:"bytes,10,opt,name=address_line2,json=addressLine2,proto3" json:"address_line2,omitempty"`
	ContactPerson        string   `protobuf:"bytes,11,opt,name=contact_person,json=contactPerson,proto3" json:"contact_person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{2}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetIsMain() bool {
	if m != nil {
		return m.IsMain
	}
	return false
}

func (m *Address) GetIsRemitTo() bool {
	if m != nil {
		return m.IsRemitTo
	}
	return false
}

func (m *Address) GetIsShipTo() bool {
	if m != nil {
		return m.IsShipTo
	}
	return false
}

func (m *Address) GetIsPayTo() bool {
	if m != nil {
		return m.IsPayTo
	}
	return false
}

func (m *Address) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetAddressLine1() string {
	if m != nil {
		return m.AddressLine1
	}
	return ""
}

func (m *Address) GetAddressLine2() string {
	if m != nil {
		return m.AddressLine2
	}
	return ""
}

func (m *Address) GetContactPerson() string {
	if m != nil {
		return m.ContactPerson
	}
	return ""
}

type BankPaymentMethod struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Address              *Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	HolderName           string   `protobuf:"bytes,3,opt,name=holder_name,json=holderName,proto3" json:"holder_name,omitempty"`
	BankKey              string   `protobuf:"bytes,4,opt,name=bank_key,json=bankKey,proto3" json:"bank_key,omitempty"`
	BankAccountNumber    string   `protobuf:"bytes,5,opt,name=bank_account_number,json=bankAccountNumber,proto3" json:"bank_account_number,omitempty"`
	SupportedCurrency    string   `protobuf:"bytes,6,opt,name=supported_currency,json=supportedCurrency,proto3" json:"supported_currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankPaymentMethod) Reset()         { *m = BankPaymentMethod{} }
func (m *BankPaymentMethod) String() string { return proto.CompactTextString(m) }
func (*BankPaymentMethod) ProtoMessage()    {}
func (*BankPaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{3}
}

func (m *BankPaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankPaymentMethod.Unmarshal(m, b)
}
func (m *BankPaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankPaymentMethod.Marshal(b, m, deterministic)
}
func (m *BankPaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankPaymentMethod.Merge(m, src)
}
func (m *BankPaymentMethod) XXX_Size() int {
	return xxx_messageInfo_BankPaymentMethod.Size(m)
}
func (m *BankPaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_BankPaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_BankPaymentMethod proto.InternalMessageInfo

func (m *BankPaymentMethod) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *BankPaymentMethod) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *BankPaymentMethod) GetHolderName() string {
	if m != nil {
		return m.HolderName
	}
	return ""
}

func (m *BankPaymentMethod) GetBankKey() string {
	if m != nil {
		return m.BankKey
	}
	return ""
}

func (m *BankPaymentMethod) GetBankAccountNumber() string {
	if m != nil {
		return m.BankAccountNumber
	}
	return ""
}

func (m *BankPaymentMethod) GetSupportedCurrency() string {
	if m != nil {
		return m.SupportedCurrency
	}
	return ""
}

type CryptoPaymentMethod struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	ChainUri             string   `protobuf:"bytes,3,opt,name=chain_uri,json=chainUri,proto3" json:"chain_uri,omitempty"`
	SupportedCurrency    string   `protobuf:"bytes,4,opt,name=supported_currency,json=supportedCurrency,proto3" json:"supported_currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoPaymentMethod) Reset()         { *m = CryptoPaymentMethod{} }
func (m *CryptoPaymentMethod) String() string { return proto.CompactTextString(m) }
func (*CryptoPaymentMethod) ProtoMessage()    {}
func (*CryptoPaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{4}
}

func (m *CryptoPaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoPaymentMethod.Unmarshal(m, b)
}
func (m *CryptoPaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoPaymentMethod.Marshal(b, m, deterministic)
}
func (m *CryptoPaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoPaymentMethod.Merge(m, src)
}
func (m *CryptoPaymentMethod) XXX_Size() int {
	return xxx_messageInfo_CryptoPaymentMethod.Size(m)
}
func (m *CryptoPaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoPaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoPaymentMethod proto.InternalMessageInfo

func (m *CryptoPaymentMethod) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *CryptoPaymentMethod) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *CryptoPaymentMethod) GetChainUri() string {
	if m != nil {
		return m.ChainUri
	}
	return ""
}

func (m *CryptoPaymentMethod) GetSupportedCurrency() string {
	if m != nil {
		return m.SupportedCurrency
	}
	return ""
}

type OtherPayment struct {
	Identifier           []byte   `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	PayTo                string   `protobuf:"bytes,3,opt,name=pay_to,json=payTo,proto3" json:"pay_to,omitempty"`
	SupportedCurrency    string   `protobuf:"bytes,4,opt,name=supported_currency,json=supportedCurrency,proto3" json:"supported_currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtherPayment) Reset()         { *m = OtherPayment{} }
func (m *OtherPayment) String() string { return proto.CompactTextString(m) }
func (*OtherPayment) ProtoMessage()    {}
func (*OtherPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{5}
}

func (m *OtherPayment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherPayment.Unmarshal(m, b)
}
func (m *OtherPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherPayment.Marshal(b, m, deterministic)
}
func (m *OtherPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherPayment.Merge(m, src)
}
func (m *OtherPayment) XXX_Size() int {
	return xxx_messageInfo_OtherPayment.Size(m)
}
func (m *OtherPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherPayment.DiscardUnknown(m)
}

var xxx_messageInfo_OtherPayment proto.InternalMessageInfo

func (m *OtherPayment) GetIdentifier() []byte {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *OtherPayment) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OtherPayment) GetPayTo() string {
	if m != nil {
		return m.PayTo
	}
	return ""
}

func (m *OtherPayment) GetSupportedCurrency() string {
	if m != nil {
		return m.SupportedCurrency
	}
	return ""
}

type PaymentDetail struct {
	// fields for bank accounts and ethereum wallets
	Predefined bool `protobuf:"varint,1,opt,name=predefined,proto3" json:"predefined,omitempty"`
	// Types that are valid to be assigned to PaymentMethod:
	//	*PaymentDetail_BankPaymentMethod
	//	*PaymentDetail_CryptoPaymentMethod
	//	*PaymentDetail_OtherMethod
	PaymentMethod        isPaymentDetail_PaymentMethod `protobuf_oneof:"payment_method"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *PaymentDetail) Reset()         { *m = PaymentDetail{} }
func (m *PaymentDetail) String() string { return proto.CompactTextString(m) }
func (*PaymentDetail) ProtoMessage()    {}
func (*PaymentDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{6}
}

func (m *PaymentDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentDetail.Unmarshal(m, b)
}
func (m *PaymentDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentDetail.Marshal(b, m, deterministic)
}
func (m *PaymentDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentDetail.Merge(m, src)
}
func (m *PaymentDetail) XXX_Size() int {
	return xxx_messageInfo_PaymentDetail.Size(m)
}
func (m *PaymentDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentDetail.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentDetail proto.InternalMessageInfo

func (m *PaymentDetail) GetPredefined() bool {
	if m != nil {
		return m.Predefined
	}
	return false
}

type isPaymentDetail_PaymentMethod interface {
	isPaymentDetail_PaymentMethod()
}

type PaymentDetail_BankPaymentMethod struct {
	BankPaymentMethod *BankPaymentMethod `protobuf:"bytes,2,opt,name=bank_payment_method,json=bankPaymentMethod,proto3,oneof"`
}

type PaymentDetail_CryptoPaymentMethod struct {
	CryptoPaymentMethod *CryptoPaymentMethod `protobuf:"bytes,3,opt,name=crypto_payment_method,json=cryptoPaymentMethod,proto3,oneof"`
}

type PaymentDetail_OtherMethod struct {
	OtherMethod *OtherPayment `protobuf:"bytes,4,opt,name=other_method,json=otherMethod,proto3,oneof"`
}

func (*PaymentDetail_BankPaymentMethod) isPaymentDetail_PaymentMethod() {}

func (*PaymentDetail_CryptoPaymentMethod) isPaymentDetail_PaymentMethod() {}

func (*PaymentDetail_OtherMethod) isPaymentDetail_PaymentMethod() {}

func (m *PaymentDetail) GetPaymentMethod() isPaymentDetail_PaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetBankPaymentMethod() *BankPaymentMethod {
	if x, ok := m.GetPaymentMethod().(*PaymentDetail_BankPaymentMethod); ok {
		return x.BankPaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetCryptoPaymentMethod() *CryptoPaymentMethod {
	if x, ok := m.GetPaymentMethod().(*PaymentDetail_CryptoPaymentMethod); ok {
		return x.CryptoPaymentMethod
	}
	return nil
}

func (m *PaymentDetail) GetOtherMethod() *OtherPayment {
	if x, ok := m.GetPaymentMethod().(*PaymentDetail_OtherMethod); ok {
		return x.OtherMethod
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PaymentDetail) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PaymentDetail_BankPaymentMethod)(nil),
		(*PaymentDetail_CryptoPaymentMethod)(nil),
		(*PaymentDetail_OtherMethod)(nil),
	}
}

type Contact struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Fax                  string   `protobuf:"bytes,5,opt,name=fax,proto3" json:"fax,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b38ccb06a827056, []int{7}
}

func (m *Contact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contact.Unmarshal(m, b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return xxx_messageInfo_Contact.Size(m)
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contact) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Contact) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Contact) GetFax() string {
	if m != nil {
		return m.Fax
	}
	return ""
}

func init() {
	proto.RegisterType((*EntityRelationship)(nil), "entity.EntityRelationship")
	proto.RegisterType((*Entity)(nil), "entity.Entity")
	proto.RegisterType((*Address)(nil), "entity.Address")
	proto.RegisterType((*BankPaymentMethod)(nil), "entity.BankPaymentMethod")
	proto.RegisterType((*CryptoPaymentMethod)(nil), "entity.CryptoPaymentMethod")
	proto.RegisterType((*OtherPayment)(nil), "entity.OtherPayment")
	proto.RegisterType((*PaymentDetail)(nil), "entity.PaymentDetail")
	proto.RegisterType((*Contact)(nil), "entity.Contact")
}

func init() { proto.RegisterFile("entity/entity.proto", fileDescriptor_9b38ccb06a827056) }

var fileDescriptor_9b38ccb06a827056 = []byte{
	// 839 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xe4, 0x34,
	0x14, 0x6f, 0xe6, 0x6f, 0xf2, 0xa6, 0x9d, 0xb6, 0x6e, 0x2b, 0xb2, 0xbb, 0xb0, 0x94, 0x41, 0x15,
	0x8b, 0xa0, 0x5d, 0x51, 0x4e, 0x5c, 0x90, 0x76, 0x0a, 0x52, 0x57, 0xcb, 0x2e, 0x43, 0x28, 0x17,
	0x2e, 0x91, 0x27, 0x71, 0x77, 0xac, 0x26, 0xb6, 0xb1, 0x3d, 0x82, 0x70, 0xe2, 0x1b, 0x70, 0xe5,
	0xc0, 0x99, 0x33, 0x5f, 0x81, 0x03, 0x5f, 0x85, 0x6f, 0xc0, 0x1d, 0xf9, 0xd9, 0x99, 0x4e, 0xb7,
	0x5d, 0x09, 0x38, 0xc5, 0xef, 0xf7, 0x7b, 0xb6, 0x5f, 0xde, 0xef, 0x97, 0x17, 0xd8, 0x63, 0xc2,
	0x72, 0xdb, 0x3c, 0xf6, 0x8f, 0x13, 0xa5, 0xa5, 0x95, 0x64, 0xe0, 0xa3, 0xfb, 0xef, 0x29, 0xcd,
	0x0a, 0x6e, 0xd8, 0xb1, 0xd2, 0x52, 0x5e, 0x9a, 0xc7, 0xd7, 0x0f, 0x2b, 0x7d, 0xe0, 0x37, 0x4c,
	0x7e, 0x8b, 0x80, 0x7c, 0x8e, 0x7b, 0x32, 0x56, 0x51, 0xcb, 0xa5, 0x30, 0x0b, 0xae, 0xc8, 0x87,
	0x30, 0x96, 0xdf, 0x0b, 0xa6, 0x73, 0x5e, 0xfa, 0x13, 0xd3, 0xe8, 0x30, 0x7a, 0xb4, 0x39, 0xed,
	0xff, 0xfe, 0xc7, 0xdf, 0xb0, 0x9f, 0x6d, 0x21, 0xf9, 0x34, 0x70, 0xe4, 0x14, 0x76, 0xfd, 0x2a,
	0xa4, 0x5f, 0x72, 0xa6, 0xd3, 0xce, 0xf5, 0x86, 0xc3, 0x6c, 0xc7, 0xf3, 0x4f, 0x57, 0x34, 0x39,
	0x81, 0x6d, 0x4b, 0xf5, 0x4b, 0x66, 0xaf, 0xaf, 0xe8, 0xae, 0x5f, 0x31, 0xf6, 0x6c, 0x7b, 0xc7,
	0xe4, 0xaf, 0x08, 0x06, 0xbe, 0x50, 0xf2, 0x0e, 0xc4, 0x77, 0x97, 0xb5, 0x82, 0xc9, 0x5b, 0x00,
	0x15, 0x7b, 0x49, 0xab, 0x5c, 0xd0, 0x9a, 0x61, 0x29, 0x49, 0x96, 0x20, 0xf2, 0x82, 0xd6, 0x8c,
	0x1c, 0x43, 0x42, 0xcb, 0x52, 0x33, 0x63, 0x98, 0x49, 0xbb, 0x87, 0xdd, 0x47, 0xa3, 0xd3, 0xed,
	0x93, 0xd0, 0xc8, 0x27, 0x9e, 0xc8, 0xae, 0x33, 0xc8, 0xa7, 0xb0, 0xad, 0x68, 0x53, 0x33, 0x61,
	0xf3, 0x92, 0x59, 0xca, 0x2b, 0x93, 0xf6, 0x70, 0xd3, 0x41, 0xbb, 0x69, 0xe6, 0xe9, 0xcf, 0x90,
	0xcd, 0xc6, 0x6a, 0x3d, 0x34, 0xe4, 0x03, 0x88, 0x0b, 0x29, 0x2c, 0x2d, 0xac, 0x49, 0xfb, 0x37,
	0x6f, 0x3b, 0xf3, 0x78, 0xb6, 0x4a, 0x98, 0xfc, 0xd9, 0x81, 0x61, 0xa8, 0x81, 0xbc, 0x01, 0x43,
	0x6e, 0xf2, 0x9a, 0x72, 0x81, 0x2f, 0x1a, 0x67, 0x03, 0x6e, 0x9e, 0x53, 0x2e, 0xc8, 0x43, 0x18,
	0x71, 0x93, 0x6b, 0x56, 0x73, 0x9b, 0x5b, 0x89, 0x2f, 0x18, 0x67, 0x09, 0x37, 0x99, 0x43, 0x2e,
	0x24, 0x79, 0x13, 0x80, 0x9b, 0xdc, 0x49, 0xe9, 0xe8, 0x2e, 0xd2, 0x31, 0x37, 0x5f, 0x2f, 0xb8,
	0xba, 0x90, 0xe4, 0x3e, 0x24, 0xdc, 0xe4, 0x8a, 0x36, 0x8e, 0xec, 0x21, 0x39, 0xe4, 0x66, 0x46,
	0x9b, 0x0b, 0x49, 0xf6, 0xa1, 0x5f, 0xd1, 0x39, 0xab, 0xd2, 0x3e, 0x36, 0xcd, 0x07, 0x64, 0x07,
	0xba, 0x3f, 0x72, 0x95, 0x0e, 0x10, 0x73, 0x4b, 0x97, 0x67, 0x2c, 0xb5, 0x2c, 0x1d, 0xfa, 0x3c,
	0x0c, 0x48, 0x0a, 0xc3, 0x42, 0x2e, 0x85, 0xd5, 0x4d, 0x1a, 0x23, 0xde, 0x86, 0xe4, 0x5d, 0xd8,
	0x0a, 0x0d, 0xcd, 0x2b, 0x2e, 0xd8, 0x47, 0x69, 0x82, 0xfc, 0x66, 0x00, 0xbf, 0x70, 0xd8, 0xab,
	0x49, 0xa7, 0x29, 0xdc, 0x4a, 0x3a, 0x25, 0x47, 0x30, 0x0e, 0xcd, 0xca, 0x15, 0xd3, 0x46, 0x8a,
	0x74, 0x84, 0x59, 0x5b, 0x01, 0x9d, 0x21, 0x38, 0xf9, 0xa9, 0x03, 0xbb, 0x53, 0x2a, 0xae, 0x82,
	0x34, 0xcf, 0x99, 0x5d, 0xc8, 0x92, 0x1c, 0x01, 0xac, 0x79, 0x34, 0x5a, 0xf7, 0xe8, 0x1a, 0x41,
	0xde, 0x87, 0x61, 0xb8, 0x13, 0x7b, 0x7b, 0x87, 0x3d, 0x5a, 0x9e, 0xbc, 0x0d, 0xa3, 0x85, 0xac,
	0x4a, 0xa6, 0xbd, 0xd7, 0xba, 0x58, 0x0b, 0x78, 0x08, 0xcd, 0x76, 0x0f, 0xe2, 0x39, 0x15, 0x57,
	0xf9, 0x15, 0x6b, 0xb0, 0xd9, 0x49, 0x36, 0x74, 0xf1, 0x33, 0xd6, 0x90, 0x13, 0xd8, 0x43, 0x8a,
	0x16, 0xd8, 0xa6, 0x5c, 0x2c, 0xeb, 0x39, 0xd3, 0xa1, 0xf5, 0xbb, 0x8e, 0x7a, 0xe2, 0x99, 0x17,
	0x48, 0x90, 0x63, 0x20, 0x66, 0xa9, 0x94, 0xd4, 0x96, 0x95, 0x79, 0xb1, 0xd4, 0x9a, 0x89, 0xa2,
	0x09, 0xaa, 0xec, 0xae, 0x98, 0xb3, 0x40, 0x4c, 0x7e, 0x89, 0x60, 0xef, 0x4c, 0x37, 0xca, 0xca,
	0xff, 0xd5, 0x84, 0x31, 0x74, 0x82, 0xb7, 0x92, 0xac, 0x63, 0x25, 0x79, 0x00, 0x49, 0xb1, 0xa0,
	0x5c, 0xe4, 0x4b, 0xcd, 0xc3, 0x7b, 0xc6, 0x08, 0x7c, 0xa3, 0xf9, 0x6b, 0x4a, 0xeb, 0xbd, 0xae,
	0xb4, 0x9f, 0x23, 0xd8, 0xfc, 0xd2, 0x2e, 0x98, 0x0e, 0x95, 0xfd, 0xdb, 0x9a, 0x08, 0xf4, 0x6c,
	0xa3, 0xda, 0x4f, 0x1a, 0xd7, 0xe4, 0x00, 0x06, 0xc1, 0xcb, 0xbe, 0xa8, 0xbe, 0x42, 0x27, 0xff,
	0xc7, 0x8a, 0x7e, 0xed, 0xc0, 0xd6, 0x8d, 0xcf, 0x98, 0x3c, 0x04, 0x50, 0x9a, 0x95, 0xec, 0x92,
	0x0b, 0x56, 0x86, 0x0f, 0x70, 0x0d, 0x21, 0xcf, 0x82, 0x7a, 0xed, 0x6c, 0xa8, 0xb1, 0xbb, 0xc1,
	0x30, 0xf7, 0x5a, 0xc3, 0xdc, 0xf2, 0xe0, 0xf9, 0x86, 0x97, 0xf6, 0xa6, 0x26, 0x5f, 0xc1, 0x41,
	0x81, 0x52, 0xbd, 0x7a, 0x5c, 0x17, 0x8f, 0x7b, 0xb0, 0x1a, 0x18, 0xb7, 0xf5, 0x3c, 0xdf, 0xc8,
	0xf6, 0x8a, 0x3b, 0x64, 0xfe, 0x04, 0x36, 0xa5, 0x6b, 0x71, 0x7b, 0x52, 0x0f, 0x4f, 0xda, 0x6f,
	0x4f, 0x5a, 0x6f, 0xff, 0xf9, 0x46, 0x36, 0xc2, 0x5c, 0xbf, 0x75, 0xba, 0x03, 0xe3, 0x9b, 0x65,
	0x4c, 0xbe, 0x83, 0x61, 0x98, 0x55, 0x4e, 0x03, 0xb4, 0x7a, 0xe4, 0x35, 0x70, 0x6b, 0x37, 0x0e,
	0x2c, 0xb7, 0x55, 0x2b, 0x8c, 0x0f, 0x1c, 0xca, 0x6a, 0xca, 0xab, 0x56, 0x18, 0x0c, 0x1c, 0xaa,
	0x16, 0x52, 0xb0, 0xa0, 0x85, 0x0f, 0xdc, 0x88, 0xb9, 0xa4, 0x3f, 0x04, 0xef, 0xbb, 0xe5, 0xf4,
	0x08, 0xa0, 0x90, 0x75, 0x28, 0x77, 0x3a, 0xf2, 0xd3, 0x7f, 0xe6, 0x7e, 0x5b, 0xb3, 0xe8, 0xdb,
	0xd8, 0xc3, 0x6a, 0x3e, 0x1f, 0xe0, 0x9f, 0xec, 0xe3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x0a, 0xfd, 0xe2, 0x11, 0x07, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/document_common.proto

package commonpb

import (
	fmt "fmt"
	_ "github.com/centrifuge/precise-proofs/proofs/proto"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BinaryAttachment struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//mime type of attached file
	FileType string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	// in byte
	Size uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Data []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	//the md5 checksum of the original file for easier verification - optional
	Checksum             []byte   `protobuf:"bytes,5,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinaryAttachment) Reset()         { *m = BinaryAttachment{} }
func (m *BinaryAttachment) String() string { return proto.CompactTextString(m) }
func (*BinaryAttachment) ProtoMessage()    {}
func (*BinaryAttachment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6061ba33c5b40c39, []int{0}
}

func (m *BinaryAttachment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinaryAttachment.Unmarshal(m, b)
}
func (m *BinaryAttachment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinaryAttachment.Marshal(b, m, deterministic)
}
func (m *BinaryAttachment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinaryAttachment.Merge(m, src)
}
func (m *BinaryAttachment) XXX_Size() int {
	return xxx_messageInfo_BinaryAttachment.Size(m)
}
func (m *BinaryAttachment) XXX_DiscardUnknown() {
	xxx_messageInfo_BinaryAttachment.DiscardUnknown(m)
}

var xxx_messageInfo_BinaryAttachment proto.InternalMessageInfo

func (m *BinaryAttachment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BinaryAttachment) GetFileType() string {
	if m != nil {
		return m.FileType
	}
	return ""
}

func (m *BinaryAttachment) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *BinaryAttachment) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BinaryAttachment) GetChecksum() []byte {
	if m != nil {
		return m.Checksum
	}
	return nil
}

type PaymentDetails struct {
	//identifying this payment. could be a sequential number, could be a transaction hash of the crypto payment
	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DateExecuted *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date_executed,json=dateExecuted,proto3" json:"date_executed,omitempty"`
	//centrifuge id of payee
	Payee []byte `protobuf:"bytes,3,opt,name=payee,proto3" json:"payee,omitempty"`
	//centrifuge id of payer
	Payer    []byte `protobuf:"bytes,4,opt,name=payer,proto3" json:"payer,omitempty"`
	Amount   []byte `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency string `protobuf:"bytes,6,opt,name=currency,proto3" json:"currency,omitempty"`
	//payment reference (e.g. reference field on bank transfer)
	Reference             string `protobuf:"bytes,7,opt,name=reference,proto3" json:"reference,omitempty"`
	BankName              string `protobuf:"bytes,8,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	BankAddress           string `protobuf:"bytes,9,opt,name=bank_address,json=bankAddress,proto3" json:"bank_address,omitempty"`
	BankCountry           string `protobuf:"bytes,10,opt,name=bank_country,json=bankCountry,proto3" json:"bank_country,omitempty"`
	BankAccountNumber     string `protobuf:"bytes,11,opt,name=bank_account_number,json=bankAccountNumber,proto3" json:"bank_account_number,omitempty"`
	BankAccountCurrency   string `protobuf:"bytes,12,opt,name=bank_account_currency,json=bankAccountCurrency,proto3" json:"bank_account_currency,omitempty"`
	BankAccountHolderName string `protobuf:"bytes,13,opt,name=bank_account_holder_name,json=bankAccountHolderName,proto3" json:"bank_account_holder_name,omitempty"`
	BankKey               string `protobuf:"bytes,14,opt,name=bank_key,json=bankKey,proto3" json:"bank_key,omitempty"`
	//the ID of the chain to use in URI format. e.g. "ethereum://42/<tokenaddress>"
	CryptoChainUri string `protobuf:"bytes,15,opt,name=crypto_chain_uri,json=cryptoChainUri,proto3" json:"crypto_chain_uri,omitempty"`
	//the transaction in which the payment happened
	CryptoTransactionId string `protobuf:"bytes,16,opt,name=crypto_transaction_id,json=cryptoTransactionId,proto3" json:"crypto_transaction_id,omitempty"`
	//from address
	CryptoFrom string `protobuf:"bytes,17,opt,name=crypto_from,json=cryptoFrom,proto3" json:"crypto_from,omitempty"`
	//to address
	CryptoTo             string   `protobuf:"bytes,18,opt,name=crypto_to,json=cryptoTo,proto3" json:"crypto_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentDetails) Reset()         { *m = PaymentDetails{} }
func (m *PaymentDetails) String() string { return proto.CompactTextString(m) }
func (*PaymentDetails) ProtoMessage()    {}
func (*PaymentDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_6061ba33c5b40c39, []int{1}
}

func (m *PaymentDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentDetails.Unmarshal(m, b)
}
func (m *PaymentDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentDetails.Marshal(b, m, deterministic)
}
func (m *PaymentDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentDetails.Merge(m, src)
}
func (m *PaymentDetails) XXX_Size() int {
	return xxx_messageInfo_PaymentDetails.Size(m)
}
func (m *PaymentDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentDetails.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentDetails proto.InternalMessageInfo

func (m *PaymentDetails) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentDetails) GetDateExecuted() *timestamp.Timestamp {
	if m != nil {
		return m.DateExecuted
	}
	return nil
}

func (m *PaymentDetails) GetPayee() []byte {
	if m != nil {
		return m.Payee
	}
	return nil
}

func (m *PaymentDetails) GetPayer() []byte {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (m *PaymentDetails) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *PaymentDetails) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PaymentDetails) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *PaymentDetails) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *PaymentDetails) GetBankAddress() string {
	if m != nil {
		return m.BankAddress
	}
	return ""
}

func (m *PaymentDetails) GetBankCountry() string {
	if m != nil {
		return m.BankCountry
	}
	return ""
}

func (m *PaymentDetails) GetBankAccountNumber() string {
	if m != nil {
		return m.BankAccountNumber
	}
	return ""
}

func (m *PaymentDetails) GetBankAccountCurrency() string {
	if m != nil {
		return m.BankAccountCurrency
	}
	return ""
}

func (m *PaymentDetails) GetBankAccountHolderName() string {
	if m != nil {
		return m.BankAccountHolderName
	}
	return ""
}

func (m *PaymentDetails) GetBankKey() string {
	if m != nil {
		return m.BankKey
	}
	return ""
}

func (m *PaymentDetails) GetCryptoChainUri() string {
	if m != nil {
		return m.CryptoChainUri
	}
	return ""
}

func (m *PaymentDetails) GetCryptoTransactionId() string {
	if m != nil {
		return m.CryptoTransactionId
	}
	return ""
}

func (m *PaymentDetails) GetCryptoFrom() string {
	if m != nil {
		return m.CryptoFrom
	}
	return ""
}

func (m *PaymentDetails) GetCryptoTo() string {
	if m != nil {
		return m.CryptoTo
	}
	return ""
}

func init() {
	proto.RegisterType((*BinaryAttachment)(nil), "common.BinaryAttachment")
	proto.RegisterType((*PaymentDetails)(nil), "common.PaymentDetails")
}

func init() { proto.RegisterFile("common/document_common.proto", fileDescriptor_6061ba33c5b40c39) }

var fileDescriptor_6061ba33c5b40c39 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x53, 0xcf, 0x4e, 0xdb, 0x4c,
	0x10, 0x97, 0xf9, 0x48, 0x48, 0x36, 0x21, 0x1f, 0x2c, 0x45, 0xda, 0x06, 0x2a, 0x52, 0x2e, 0xcd,
	0x05, 0x47, 0xa2, 0x87, 0x1e, 0x2b, 0x12, 0x5a, 0xb5, 0xaa, 0x84, 0x90, 0x45, 0x2f, 0xbd, 0x58,
	0x9b, 0xf5, 0x84, 0xac, 0xc8, 0x7a, 0xad, 0xf5, 0x5a, 0xaa, 0xfb, 0x02, 0x7d, 0x98, 0x5e, 0xfa,
	0x0a, 0x7d, 0xa8, 0xde, 0xab, 0x9d, 0xb1, 0x09, 0x9c, 0xb2, 0xbf, 0x7f, 0x9b, 0xdf, 0x8c, 0x6d,
	0x76, 0xaa, 0xac, 0x31, 0x36, 0x9f, 0x65, 0x56, 0x55, 0x06, 0x72, 0x9f, 0x12, 0x8e, 0x0b, 0x67,
	0xbd, 0xe5, 0x5d, 0x42, 0xe3, 0xb3, 0x7b, 0x6b, 0xef, 0x37, 0x30, 0x43, 0x76, 0x59, 0xad, 0x66,
	0x5e, 0x1b, 0x28, 0xbd, 0x34, 0x05, 0x19, 0xc7, 0x6f, 0x0a, 0x07, 0x4a, 0x97, 0x70, 0x51, 0x38,
	0x6b, 0x57, 0xe5, 0x6c, 0xfb, 0xe3, 0x2d, 0x01, 0x32, 0x9e, 0xff, 0x8c, 0xd8, 0xc1, 0x5c, 0xe7,
	0xd2, 0xd5, 0x57, 0xde, 0x4b, 0xb5, 0x0e, 0xff, 0xc9, 0x39, 0xdb, 0xcd, 0xa5, 0x01, 0x11, 0x4d,
	0xa2, 0x69, 0x3f, 0xc1, 0x33, 0x3f, 0x61, 0xfd, 0x95, 0xde, 0x40, 0xea, 0xeb, 0x02, 0xc4, 0x0e,
	0x0a, 0xbd, 0x40, 0xdc, 0xd5, 0x05, 0x84, 0x40, 0xa9, 0x7f, 0x80, 0xf8, 0x6f, 0x12, 0x4d, 0x77,
	0x13, 0x3c, 0x07, 0x2e, 0x93, 0x5e, 0x8a, 0xdd, 0x49, 0x34, 0x1d, 0x26, 0x78, 0xe6, 0x63, 0xd6,
	0x53, 0x6b, 0x50, 0x0f, 0x65, 0x65, 0x44, 0x07, 0xf9, 0x47, 0x7c, 0xfe, 0xab, 0xc3, 0x46, 0xb7,
	0xb2, 0x0e, 0x05, 0xae, 0xc1, 0x4b, 0xbd, 0x29, 0xf9, 0x88, 0xed, 0xe8, 0xac, 0x69, 0xb1, 0xa3,
	0x33, 0xfe, 0x9e, 0xed, 0x67, 0xd2, 0x43, 0x0a, 0xdf, 0x41, 0x55, 0x1e, 0x32, 0xec, 0x31, 0xb8,
	0x1c, 0xc7, 0xb4, 0x8e, 0xb8, 0x5d, 0x47, 0x7c, 0xd7, 0xae, 0x23, 0x19, 0x86, 0xc0, 0x87, 0xc6,
	0xcf, 0x4f, 0x58, 0xa7, 0x90, 0x35, 0x50, 0xd1, 0xe1, 0xbc, 0xf3, 0xfb, 0xcf, 0x5f, 0xf6, 0x22,
	0x21, 0xae, 0x15, 0x1d, 0x35, 0x7e, 0x26, 0x3a, 0xfe, 0x8a, 0x75, 0xa5, 0xb1, 0x55, 0xee, 0xa9,
	0x37, 0xa9, 0x93, 0xa4, 0x21, 0x71, 0xb0, 0xca, 0x39, 0xc8, 0x55, 0x2d, 0xba, 0xb4, 0x9c, 0x16,
	0xf3, 0x53, 0xd6, 0x77, 0xb0, 0x82, 0x00, 0x40, 0xec, 0xa1, 0xb8, 0x25, 0xc2, 0x5e, 0x97, 0x32,
	0x7f, 0x48, 0x71, 0xe1, 0x3d, 0x8a, 0x06, 0xe2, 0x26, 0x2c, 0xfd, 0x35, 0x1b, 0xa2, 0x28, 0xb3,
	0xcc, 0x41, 0x59, 0x8a, 0x3e, 0xea, 0x83, 0xc0, 0x5d, 0x11, 0xf5, 0x68, 0x51, 0xa1, 0x87, 0xab,
	0x05, 0xdb, 0x5a, 0x16, 0x44, 0xf1, 0x98, 0x1d, 0xd1, 0x2d, 0x0a, 0x4d, 0x69, 0x5e, 0x99, 0x25,
	0x38, 0x31, 0x40, 0xe7, 0x21, 0x5e, 0x46, 0xca, 0x0d, 0x0a, 0xfc, 0x92, 0x1d, 0x3f, 0xf3, 0x3f,
	0x4e, 0x36, 0xc4, 0xc4, 0xd1, 0x93, 0xc4, 0xa2, 0x1d, 0xf2, 0x1d, 0x13, 0xcf, 0x32, 0x6b, 0xbb,
	0xc9, 0xc0, 0xd1, 0x54, 0xfb, 0x18, 0x3b, 0x7e, 0x12, 0xfb, 0x84, 0x2a, 0x8e, 0xf8, 0x92, 0xe1,
	0xb8, 0xe9, 0x03, 0xd4, 0x62, 0x84, 0xc6, 0xbd, 0x80, 0xbf, 0x40, 0xcd, 0xa7, 0xec, 0x40, 0xb9,
	0xba, 0xf0, 0x36, 0x55, 0x6b, 0xa9, 0xf3, 0xb4, 0x72, 0x5a, 0xfc, 0x8f, 0x96, 0x11, 0xf1, 0x8b,
	0x40, 0x7f, 0x75, 0x3a, 0x34, 0x6e, 0x9c, 0xde, 0xc9, 0xbc, 0x94, 0xca, 0x6b, 0x9b, 0xa7, 0x3a,
	0x13, 0x07, 0xd4, 0x98, 0xc4, 0xbb, 0xad, 0xf6, 0x39, 0xe3, 0x67, 0x6c, 0xd0, 0x64, 0x56, 0xce,
	0x1a, 0x71, 0x88, 0x4e, 0x46, 0xd4, 0x47, 0x67, 0x4d, 0x78, 0x32, 0xed, 0xa5, 0x56, 0xf0, 0xe6,
	0xa1, 0xd2, 0x45, 0x76, 0x7e, 0xc1, 0x98, 0xb2, 0x26, 0xa6, 0xef, 0x71, 0x7e, 0x74, 0xdd, 0x7c,
	0xae, 0x0b, 0xc4, 0xb7, 0xe1, 0x3d, 0xbc, 0x8d, 0xbe, 0xf5, 0x48, 0x2e, 0x96, 0xcb, 0x2e, 0xbe,
	0x9a, 0x6f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xcc, 0x97, 0xdf, 0xdf, 0x03, 0x00, 0x00,
}

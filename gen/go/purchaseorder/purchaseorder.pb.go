// Code generated by protoc-gen-go. DO NOT EDIT.
// source: purchaseorder/purchaseorder.proto

package purchaseorderpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import coredocument "github.com/CentrifugeInc/centrifuge-protobufs/gen/go/coredocument"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PurchaseOrderData is the default schema for a purchase order
type PurchaseOrderData struct {
	Country              string               `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Currency             string               `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               int64                `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Recipient            []byte               `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Sender               []byte               `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Comment              string               `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	DeliveryDate         *timestamp.Timestamp `protobuf:"bytes,9,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PurchaseOrderData) Reset()         { *m = PurchaseOrderData{} }
func (m *PurchaseOrderData) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderData) ProtoMessage()    {}
func (*PurchaseOrderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_dd1a03b1035fc7f6, []int{0}
}
func (m *PurchaseOrderData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderData.Unmarshal(m, b)
}
func (m *PurchaseOrderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderData.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderData.Merge(dst, src)
}
func (m *PurchaseOrderData) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderData.Size(m)
}
func (m *PurchaseOrderData) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderData.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderData proto.InternalMessageInfo

func (m *PurchaseOrderData) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *PurchaseOrderData) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *PurchaseOrderData) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *PurchaseOrderData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *PurchaseOrderData) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PurchaseOrderData) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *PurchaseOrderData) GetDeliveryDate() *timestamp.Timestamp {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

// PurhcaseOrderDataSalts keeps track of the salts used for each PurchaseOrderData field.
type PurchaseOrderDataSalts struct {
	Country              []byte   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Currency             []byte   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount               []byte   `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Recipient            []byte   `protobuf:"bytes,5,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Sender               []byte   `protobuf:"bytes,6,opt,name=sender,proto3" json:"sender,omitempty"`
	Comment              []byte   `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	DeliveryDate         []byte   `protobuf:"bytes,9,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PurchaseOrderDataSalts) Reset()         { *m = PurchaseOrderDataSalts{} }
func (m *PurchaseOrderDataSalts) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderDataSalts) ProtoMessage()    {}
func (*PurchaseOrderDataSalts) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_dd1a03b1035fc7f6, []int{1}
}
func (m *PurchaseOrderDataSalts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderDataSalts.Unmarshal(m, b)
}
func (m *PurchaseOrderDataSalts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderDataSalts.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderDataSalts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderDataSalts.Merge(dst, src)
}
func (m *PurchaseOrderDataSalts) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderDataSalts.Size(m)
}
func (m *PurchaseOrderDataSalts) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderDataSalts.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderDataSalts proto.InternalMessageInfo

func (m *PurchaseOrderDataSalts) GetCountry() []byte {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetCurrency() []byte {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetAmount() []byte {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetComment() []byte {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *PurchaseOrderDataSalts) GetDeliveryDate() []byte {
	if m != nil {
		return m.DeliveryDate
	}
	return nil
}

// PurchaseOrderDocument combines the salts, data & coredocument for a purchase order.
type PurchaseOrderDocument struct {
	CoreDocument         *coredocument.CoreDocument `protobuf:"bytes,1,opt,name=core_document,json=coreDocument,proto3" json:"core_document,omitempty"`
	Data                 *PurchaseOrderData         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Salts                *PurchaseOrderDataSalts    `protobuf:"bytes,3,opt,name=salts,proto3" json:"salts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PurchaseOrderDocument) Reset()         { *m = PurchaseOrderDocument{} }
func (m *PurchaseOrderDocument) String() string { return proto.CompactTextString(m) }
func (*PurchaseOrderDocument) ProtoMessage()    {}
func (*PurchaseOrderDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_purchaseorder_dd1a03b1035fc7f6, []int{2}
}
func (m *PurchaseOrderDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PurchaseOrderDocument.Unmarshal(m, b)
}
func (m *PurchaseOrderDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PurchaseOrderDocument.Marshal(b, m, deterministic)
}
func (dst *PurchaseOrderDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PurchaseOrderDocument.Merge(dst, src)
}
func (m *PurchaseOrderDocument) XXX_Size() int {
	return xxx_messageInfo_PurchaseOrderDocument.Size(m)
}
func (m *PurchaseOrderDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_PurchaseOrderDocument.DiscardUnknown(m)
}

var xxx_messageInfo_PurchaseOrderDocument proto.InternalMessageInfo

func (m *PurchaseOrderDocument) GetCoreDocument() *coredocument.CoreDocument {
	if m != nil {
		return m.CoreDocument
	}
	return nil
}

func (m *PurchaseOrderDocument) GetData() *PurchaseOrderData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PurchaseOrderDocument) GetSalts() *PurchaseOrderDataSalts {
	if m != nil {
		return m.Salts
	}
	return nil
}

func init() {
	proto.RegisterType((*PurchaseOrderData)(nil), "purchaseorder.PurchaseOrderData")
	proto.RegisterType((*PurchaseOrderDataSalts)(nil), "purchaseorder.PurchaseOrderDataSalts")
	proto.RegisterType((*PurchaseOrderDocument)(nil), "purchaseorder.PurchaseOrderDocument")
}

func init() {
	proto.RegisterFile("purchaseorder/purchaseorder.proto", fileDescriptor_purchaseorder_dd1a03b1035fc7f6)
}

var fileDescriptor_purchaseorder_dd1a03b1035fc7f6 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x99, 0xdb, 0x3f, 0xb7, 0x9d, 0x4e, 0xb9, 0x74, 0xe0, 0x96, 0x21, 0x08, 0x8d, 0x15,
	0x21, 0xab, 0x04, 0xaa, 0x3b, 0x17, 0x82, 0x76, 0xaf, 0x44, 0x57, 0x6e, 0xca, 0x74, 0x72, 0xac,
	0x81, 0x26, 0x13, 0x26, 0x13, 0xa1, 0xcf, 0xe8, 0x23, 0xf8, 0x16, 0x3e, 0x81, 0x64, 0x92, 0xb1,
	0x4d, 0x6b, 0x71, 0xe3, 0xf2, 0x9b, 0x73, 0xbe, 0xe1, 0x7c, 0xbf, 0x73, 0xf0, 0x69, 0x56, 0x28,
	0xf1, 0xc2, 0x73, 0x90, 0x2a, 0x02, 0x15, 0x34, 0x94, 0x9f, 0x29, 0xa9, 0x25, 0x1d, 0x36, 0x1e,
	0x9d, 0xc9, 0x4a, 0xca, 0xd5, 0x1a, 0x02, 0x53, 0x5c, 0x16, 0xcf, 0x81, 0x8e, 0x13, 0xc8, 0x35,
	0x4f, 0xb2, 0xaa, 0xdf, 0x99, 0x08, 0xa9, 0x20, 0x92, 0xa2, 0x48, 0x20, 0xd5, 0xc1, 0xae, 0xa8,
	0x1a, 0xa6, 0x1f, 0x08, 0x8f, 0xee, 0xeb, 0x3f, 0xef, 0xca, 0x3f, 0xe7, 0x5c, 0x73, 0xca, 0xf0,
	0x5f, 0x21, 0x8b, 0x54, 0xab, 0x0d, 0xfb, 0xe3, 0x22, 0xaf, 0x1f, 0x5a, 0x49, 0x1d, 0xdc, 0x13,
	0x85, 0x52, 0x90, 0x8a, 0x0d, 0x6b, 0x99, 0xd2, 0x97, 0xa6, 0x63, 0xdc, 0xe5, 0x49, 0xd9, 0xc7,
	0xda, 0x2e, 0xf2, 0x5a, 0x61, 0xad, 0xe8, 0x09, 0xee, 0x2b, 0x10, 0x71, 0x16, 0x43, 0xaa, 0x59,
	0xc7, 0x45, 0x1e, 0x09, 0xb7, 0x0f, 0xa5, 0x2b, 0x87, 0x34, 0x02, 0xc5, 0xba, 0xa6, 0x54, 0xab,
	0x6a, 0x86, 0xa4, 0x1c, 0x95, 0xf5, 0xec, 0x0c, 0x46, 0xd2, 0x6b, 0x3c, 0x8c, 0x60, 0x1d, 0xbf,
	0x82, 0xda, 0x2c, 0x22, 0xae, 0x81, 0xf5, 0x5d, 0xe4, 0x0d, 0x66, 0x8e, 0x5f, 0xd1, 0xf0, 0x2d,
	0x0d, 0xff, 0xd1, 0xd2, 0x08, 0x89, 0x35, 0xcc, 0xb9, 0x86, 0xe9, 0x3b, 0xc2, 0xe3, 0x83, 0xd0,
	0x0f, 0x7c, 0xad, 0xf3, 0xfd, 0xe4, 0xe4, 0x78, 0x72, 0x72, 0x34, 0x39, 0xf9, 0xdd, 0xe4, 0x64,
	0x9b, 0xfc, 0xec, 0xbb, 0xe4, 0x64, 0x2f, 0xdd, 0x1b, 0xc2, 0xff, 0x9b, 0xe9, 0xea, 0x95, 0x97,
	0xe0, 0xca, 0x13, 0x58, 0xd8, 0x1b, 0x60, 0xa8, 0x06, 0xd7, 0x38, 0x8c, 0x5b, 0xa9, 0xc0, 0x5a,
	0x42, 0x22, 0x76, 0x14, 0xbd, 0xc4, 0xed, 0x88, 0x6b, 0x6e, 0xd0, 0x0c, 0x66, 0xae, 0xdf, 0x3c,
	0xd1, 0x03, 0xa4, 0xa1, 0xe9, 0xa6, 0x57, 0xb8, 0x93, 0x97, 0x70, 0x0d, 0xb6, 0xc1, 0xec, 0xfc,
	0x27, 0x9b, 0xd9, 0x44, 0x58, 0x79, 0x6e, 0x46, 0x4f, 0xff, 0x1a, 0xed, 0xd9, 0x72, 0xd9, 0x35,
	0x0b, 0xbe, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x15, 0x16, 0xe4, 0xc6, 0x30, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/inventory/inventory.proto

package eopenio_emicro_template_srv_inventory

import (
	fmt "fmt"
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

type Inv struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId               int64    `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	UnitPrice            int64    `protobuf:"varint,3,opt,name=unitPrice,proto3" json:"unitPrice,omitempty"`
	Stock                int64    `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Version              int64    `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	CreatedTime          int64    `protobuf:"varint,6,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          int64    `protobuf:"varint,7,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Inv) Reset()         { *m = Inv{} }
func (m *Inv) String() string { return proto.CompactTextString(m) }
func (*Inv) ProtoMessage()    {}
func (*Inv) Descriptor() ([]byte, []int) {
	return fileDescriptor_d77ba5e53e1c0527, []int{0}
}

func (m *Inv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Inv.Unmarshal(m, b)
}
func (m *Inv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Inv.Marshal(b, m, deterministic)
}
func (m *Inv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Inv.Merge(m, src)
}
func (m *Inv) XXX_Size() int {
	return xxx_messageInfo_Inv.Size(m)
}
func (m *Inv) XXX_DiscardUnknown() {
	xxx_messageInfo_Inv.DiscardUnknown(m)
}

var xxx_messageInfo_Inv proto.InternalMessageInfo

func (m *Inv) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Inv) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Inv) GetUnitPrice() int64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *Inv) GetStock() int64 {
	if m != nil {
		return m.Stock
	}
	return 0
}

func (m *Inv) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Inv) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Inv) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

type InvHistory struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId               int64    `protobuf:"varint,2,opt,name=bookId,proto3" json:"bookId,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	State                int64    `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvHistory) Reset()         { *m = InvHistory{} }
func (m *InvHistory) String() string { return proto.CompactTextString(m) }
func (*InvHistory) ProtoMessage()    {}
func (*InvHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_d77ba5e53e1c0527, []int{1}
}

func (m *InvHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvHistory.Unmarshal(m, b)
}
func (m *InvHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvHistory.Marshal(b, m, deterministic)
}
func (m *InvHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvHistory.Merge(m, src)
}
func (m *InvHistory) XXX_Size() int {
	return xxx_messageInfo_InvHistory.Size(m)
}
func (m *InvHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_InvHistory.DiscardUnknown(m)
}

var xxx_messageInfo_InvHistory proto.InternalMessageInfo

func (m *InvHistory) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InvHistory) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *InvHistory) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *InvHistory) GetState() int64 {
	if m != nil {
		return m.State
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_d77ba5e53e1c0527, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	BookId               int64    `protobuf:"varint,1,opt,name=bookId,proto3" json:"bookId,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	HistoryId            int64    `protobuf:"varint,3,opt,name=historyId,proto3" json:"historyId,omitempty"`
	HistoryState         int32    `protobuf:"varint,4,opt,name=historyState,proto3" json:"historyState,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_d77ba5e53e1c0527, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetBookId() int64 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *Request) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetHistoryId() int64 {
	if m != nil {
		return m.HistoryId
	}
	return 0
}

func (m *Request) GetHistoryState() int32 {
	if m != nil {
		return m.HistoryState
	}
	return 0
}

type Response struct {
	Success              bool        `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Inv                  *Inv        `protobuf:"bytes,3,opt,name=inv,proto3" json:"inv,omitempty"`
	InvH                 *InvHistory `protobuf:"bytes,4,opt,name=invH,proto3" json:"invH,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_d77ba5e53e1c0527, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetInv() *Inv {
	if m != nil {
		return m.Inv
	}
	return nil
}

func (m *Response) GetInvH() *InvHistory {
	if m != nil {
		return m.InvH
	}
	return nil
}

func init() {
	proto.RegisterType((*Inv)(nil), "eopenio.emicro.template.srv.inventory.Inv")
	proto.RegisterType((*InvHistory)(nil), "eopenio.emicro.template.srv.inventory.InvHistory")
	proto.RegisterType((*Error)(nil), "eopenio.emicro.template.srv.inventory.Error")
	proto.RegisterType((*Request)(nil), "eopenio.emicro.template.srv.inventory.Request")
	proto.RegisterType((*Response)(nil), "eopenio.emicro.template.srv.inventory.Response")
}

func init() { proto.RegisterFile("proto/inventory/inventory.proto", fileDescriptor_d77ba5e53e1c0527) }

var fileDescriptor_d77ba5e53e1c0527 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0xad, 0x93, 0x38, 0xd9, 0x4c, 0x4a, 0x0f, 0x43, 0x59, 0x4c, 0x59, 0xe8, 0x62, 0x28, 0x94,
	0x52, 0xb4, 0x34, 0x7b, 0xed, 0xa9, 0x65, 0x61, 0x7d, 0x2b, 0xda, 0xfe, 0x80, 0x63, 0x4d, 0xa9,
	0x58, 0x47, 0x72, 0x25, 0x59, 0x50, 0xfa, 0x71, 0xfd, 0x99, 0x5e, 0xfb, 0x0f, 0xc5, 0x63, 0x27,
	0x4e, 0x0e, 0x05, 0x9f, 0xf6, 0xa6, 0xf7, 0x66, 0x46, 0xef, 0x3d, 0x8d, 0x0d, 0xaf, 0x1b, 0x67,
	0x83, 0xbd, 0xd1, 0x26, 0x92, 0x09, 0xd6, 0xfd, 0x1c, 0x4f, 0x82, 0x2b, 0xf8, 0x86, 0x6c, 0x43,
	0x46, 0x5b, 0x41, 0x7b, 0x5d, 0x39, 0x2b, 0x02, 0xed, 0x9b, 0xba, 0x0c, 0x24, 0xbc, 0x8b, 0xe2,
	0xd8, 0x9c, 0xff, 0x4e, 0x60, 0x5e, 0x98, 0x88, 0x2f, 0x60, 0xa6, 0x55, 0x96, 0x5c, 0x27, 0x6f,
	0xe7, 0x72, 0xa6, 0x15, 0x5e, 0xc2, 0x72, 0x67, 0xed, 0x63, 0xa1, 0xb2, 0x19, 0x73, 0x03, 0xc2,
	0x2b, 0x58, 0xb7, 0x46, 0x87, 0x2f, 0x4e, 0x57, 0x94, 0xcd, 0xb9, 0x34, 0x12, 0xf8, 0x12, 0x52,
	0x1f, 0x6c, 0xf5, 0x98, 0x2d, 0xb8, 0xd2, 0x03, 0xcc, 0x60, 0x15, 0xc9, 0x79, 0x6d, 0x4d, 0x96,
	0x32, 0x7f, 0x80, 0x78, 0x0d, 0x9b, 0xca, 0x51, 0x19, 0x48, 0x7d, 0xd5, 0x7b, 0xca, 0x96, 0x5c,
	0x3d, 0xa5, 0xba, 0x8e, 0xb6, 0x51, 0xc7, 0x8e, 0x55, 0xdf, 0x71, 0x42, 0xe5, 0x3b, 0x80, 0xc2,
	0xc4, 0x7b, 0xed, 0xbb, 0x3c, 0x93, 0x73, 0x5c, 0xc2, 0xb2, 0xf5, 0xe4, 0x0a, 0x35, 0x84, 0x18,
	0x50, 0x9f, 0xa0, 0x0c, 0x34, 0x26, 0x28, 0x03, 0xe5, 0xb7, 0x90, 0xde, 0x39, 0x67, 0x1d, 0x22,
	0x2c, 0x2a, 0xab, 0x88, 0x05, 0x52, 0xc9, 0xe7, 0xee, 0x2a, 0x45, 0xa1, 0xd4, 0x35, 0x4b, 0xac,
	0xe5, 0x80, 0xf2, 0x5f, 0xb0, 0x92, 0xf4, 0xa3, 0x25, 0x1f, 0x4e, 0x5c, 0x24, 0xff, 0x71, 0x31,
	0x3b, 0x73, 0x71, 0x05, 0xeb, 0xef, 0x7d, 0xa0, 0xa3, 0xc1, 0x91, 0xc0, 0x1c, 0x9e, 0x0f, 0xe0,
	0xe1, 0x68, 0x35, 0x95, 0x67, 0x5c, 0xfe, 0x37, 0x81, 0x0b, 0x49, 0xbe, 0xb1, 0xc6, 0x53, 0xb7,
	0x00, 0xdf, 0x56, 0x15, 0x79, 0xcf, 0xfa, 0x17, 0xf2, 0x00, 0xf1, 0x13, 0xa4, 0xd4, 0x05, 0x63,
	0xfd, 0xcd, 0xf6, 0xbd, 0x98, 0xf4, 0xd5, 0x08, 0x7e, 0x0c, 0xd9, 0x8f, 0xe2, 0x47, 0x98, 0x6b,
	0x13, 0xd9, 0xe6, 0x66, 0xfb, 0x6e, 0xe2, 0x0d, 0x85, 0x89, 0xb2, 0x1b, 0xc3, 0x3b, 0x58, 0x68,
	0x13, 0xef, 0x39, 0xc4, 0x66, 0xfb, 0x61, 0xfa, 0xf8, 0xb0, 0x71, 0xc9, 0xe3, 0xdb, 0x3f, 0x09,
	0xac, 0x8b, 0x43, 0x19, 0x35, 0x2c, 0x1e, 0xa8, 0xae, 0x51, 0x4c, 0xbc, 0x6e, 0xd8, 0xd3, 0xab,
	0x9b, 0xc9, 0xfd, 0xfd, 0xcb, 0xe6, 0xcf, 0xb0, 0x86, 0xd5, 0x67, 0x6b, 0xbe, 0x69, 0xb7, 0x7f,
	0x02, 0xb5, 0xdd, 0x92, 0x7f, 0xee, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x8f, 0xb4,
	0x45, 0xff, 0x03, 0x00, 0x00,
}

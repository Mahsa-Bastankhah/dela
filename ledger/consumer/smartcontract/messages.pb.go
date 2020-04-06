// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package smartcontract

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type InstanceProto struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ContractID           string   `protobuf:"bytes,3,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Deleted              bool     `protobuf:"varint,4,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceProto) Reset()         { *m = InstanceProto{} }
func (m *InstanceProto) String() string { return proto.CompactTextString(m) }
func (*InstanceProto) ProtoMessage()    {}
func (*InstanceProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *InstanceProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceProto.Unmarshal(m, b)
}
func (m *InstanceProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceProto.Marshal(b, m, deterministic)
}
func (m *InstanceProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceProto.Merge(m, src)
}
func (m *InstanceProto) XXX_Size() int {
	return xxx_messageInfo_InstanceProto.Size(m)
}
func (m *InstanceProto) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceProto.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceProto proto.InternalMessageInfo

func (m *InstanceProto) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *InstanceProto) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *InstanceProto) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *InstanceProto) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type SpawnTransactionProto struct {
	ContractID           string   `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Argument             *any.Any `protobuf:"bytes,2,opt,name=argument,proto3" json:"argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpawnTransactionProto) Reset()         { *m = SpawnTransactionProto{} }
func (m *SpawnTransactionProto) String() string { return proto.CompactTextString(m) }
func (*SpawnTransactionProto) ProtoMessage()    {}
func (*SpawnTransactionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *SpawnTransactionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnTransactionProto.Unmarshal(m, b)
}
func (m *SpawnTransactionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnTransactionProto.Marshal(b, m, deterministic)
}
func (m *SpawnTransactionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnTransactionProto.Merge(m, src)
}
func (m *SpawnTransactionProto) XXX_Size() int {
	return xxx_messageInfo_SpawnTransactionProto.Size(m)
}
func (m *SpawnTransactionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnTransactionProto.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnTransactionProto proto.InternalMessageInfo

func (m *SpawnTransactionProto) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *SpawnTransactionProto) GetArgument() *any.Any {
	if m != nil {
		return m.Argument
	}
	return nil
}

type InvokeTransactionProto struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Argument             *any.Any `protobuf:"bytes,2,opt,name=argument,proto3" json:"argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvokeTransactionProto) Reset()         { *m = InvokeTransactionProto{} }
func (m *InvokeTransactionProto) String() string { return proto.CompactTextString(m) }
func (*InvokeTransactionProto) ProtoMessage()    {}
func (*InvokeTransactionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *InvokeTransactionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvokeTransactionProto.Unmarshal(m, b)
}
func (m *InvokeTransactionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvokeTransactionProto.Marshal(b, m, deterministic)
}
func (m *InvokeTransactionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeTransactionProto.Merge(m, src)
}
func (m *InvokeTransactionProto) XXX_Size() int {
	return xxx_messageInfo_InvokeTransactionProto.Size(m)
}
func (m *InvokeTransactionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeTransactionProto.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeTransactionProto proto.InternalMessageInfo

func (m *InvokeTransactionProto) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *InvokeTransactionProto) GetArgument() *any.Any {
	if m != nil {
		return m.Argument
	}
	return nil
}

type DeleteTransactionProto struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTransactionProto) Reset()         { *m = DeleteTransactionProto{} }
func (m *DeleteTransactionProto) String() string { return proto.CompactTextString(m) }
func (*DeleteTransactionProto) ProtoMessage()    {}
func (*DeleteTransactionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *DeleteTransactionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTransactionProto.Unmarshal(m, b)
}
func (m *DeleteTransactionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTransactionProto.Marshal(b, m, deterministic)
}
func (m *DeleteTransactionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTransactionProto.Merge(m, src)
}
func (m *DeleteTransactionProto) XXX_Size() int {
	return xxx_messageInfo_DeleteTransactionProto.Size(m)
}
func (m *DeleteTransactionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTransactionProto.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTransactionProto proto.InternalMessageInfo

func (m *DeleteTransactionProto) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*InstanceProto)(nil), "smartcontract.InstanceProto")
	proto.RegisterType((*SpawnTransactionProto)(nil), "smartcontract.SpawnTransactionProto")
	proto.RegisterType((*InvokeTransactionProto)(nil), "smartcontract.InvokeTransactionProto")
	proto.RegisterType((*DeleteTransactionProto)(nil), "smartcontract.DeleteTransactionProto")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x59, 0xeb, 0x9f, 0x3a, 0x5a, 0x91, 0xa0, 0x65, 0xf5, 0x20, 0x21, 0xa7, 0xd0, 0x43,
	0x2a, 0xfa, 0x09, 0x84, 0x5e, 0x72, 0x93, 0xe8, 0xd1, 0xcb, 0x34, 0x1d, 0x43, 0x69, 0x32, 0x5b,
	0x76, 0x27, 0x95, 0x7c, 0x02, 0xbf, 0xb6, 0xb8, 0x4b, 0x44, 0x09, 0x28, 0xde, 0x76, 0x67, 0xde,
	0xfc, 0xde, 0x7b, 0x70, 0xd6, 0x90, 0x73, 0x58, 0x91, 0xcb, 0xb6, 0xd6, 0x88, 0x89, 0x26, 0xae,
	0x41, 0x2b, 0xa5, 0x61, 0xb1, 0x58, 0xca, 0xf5, 0x55, 0x65, 0x4c, 0x55, 0xd3, 0xdc, 0x2f, 0x97,
	0xed, 0xeb, 0x1c, 0xb9, 0x0b, 0xca, 0xe4, 0x5d, 0xc1, 0x24, 0x67, 0x27, 0xc8, 0x25, 0x3d, 0xfa,
	0xdb, 0x73, 0x18, 0x6d, 0xa8, 0xd3, 0x2a, 0x56, 0xe9, 0x69, 0xf1, 0xf9, 0x8c, 0x66, 0x70, 0xb0,
	0xc3, 0xba, 0x25, 0xbd, 0x17, 0xab, 0xf4, 0xe4, 0xee, 0x22, 0x0b, 0xb8, 0xac, 0xc7, 0x65, 0x0f,
	0xdc, 0x15, 0x41, 0x12, 0xdd, 0x00, 0xf4, 0xb6, 0xf9, 0x42, 0x8f, 0x62, 0x95, 0x1e, 0x17, 0xdf,
	0x26, 0x91, 0x86, 0xa3, 0x15, 0xd5, 0x24, 0xb4, 0xd2, 0xfb, 0xb1, 0x4a, 0xc7, 0x45, 0xff, 0x4d,
	0xd6, 0x70, 0xf9, 0xb4, 0xc5, 0x37, 0x7e, 0xb6, 0xc8, 0x0e, 0x4b, 0x59, 0x1b, 0x0e, 0x81, 0x7e,
	0x22, 0xd5, 0x00, 0x79, 0x0b, 0x63, 0xb4, 0x55, 0xdb, 0x10, 0xcb, 0xaf, 0x09, 0xbf, 0x54, 0xc9,
	0x0b, 0x4c, 0x73, 0xde, 0x99, 0x0d, 0x0d, 0xbc, 0x86, 0xe5, 0xff, 0x4f, 0x9f, 0xc1, 0x74, 0xe1,
	0x3b, 0xfd, 0x4d, 0x5f, 0x1e, 0x7a, 0xc6, 0xfd, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x38,
	0xb1, 0x3d, 0xc1, 0x01, 0x00, 0x00,
}

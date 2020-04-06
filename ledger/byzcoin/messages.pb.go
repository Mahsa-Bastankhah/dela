// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package byzcoin

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

// BlockPayload is the message that will be stored in the blocks. It is composed
// of the transactions and the footprint of the new inventory.
type BlockPayload struct {
	Transactions []*any.Any `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	// Footprint is an integrity check of the final state of the inventory after
	// applying the transactions.
	Footprint            []byte   `protobuf:"bytes,2,opt,name=footprint,proto3" json:"footprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockPayload) Reset()         { *m = BlockPayload{} }
func (m *BlockPayload) String() string { return proto.CompactTextString(m) }
func (*BlockPayload) ProtoMessage()    {}
func (*BlockPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *BlockPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockPayload.Unmarshal(m, b)
}
func (m *BlockPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockPayload.Marshal(b, m, deterministic)
}
func (m *BlockPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockPayload.Merge(m, src)
}
func (m *BlockPayload) XXX_Size() int {
	return xxx_messageInfo_BlockPayload.Size(m)
}
func (m *BlockPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BlockPayload proto.InternalMessageInfo

func (m *BlockPayload) GetTransactions() []*any.Any {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *BlockPayload) GetFootprint() []byte {
	if m != nil {
		return m.Footprint
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockPayload)(nil), "byzcoin.BlockPayload")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xaa, 0xac, 0x4a,
	0xce, 0xcf, 0xcc, 0x93, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x27, 0x95,
	0xa6, 0xe9, 0x27, 0xe6, 0x55, 0x42, 0xd4, 0x28, 0xa5, 0x71, 0xf1, 0x38, 0xe5, 0xe4, 0x27, 0x67,
	0x07, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x08, 0x59, 0x70, 0xf1, 0x94, 0x14, 0x25, 0xe6, 0x15,
	0x27, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x15, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x89, 0xe8,
	0x41, 0x4c, 0xd0, 0x83, 0x99, 0xa0, 0xe7, 0x98, 0x57, 0x19, 0x84, 0xa2, 0x52, 0x48, 0x86, 0x8b,
	0x33, 0x2d, 0x3f, 0xbf, 0xa4, 0xa0, 0x28, 0x33, 0xaf, 0x44, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27,
	0x08, 0x21, 0x90, 0xc4, 0x06, 0xd6, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x63, 0x69,
	0x19, 0xa4, 0x00, 0x00, 0x00,
}

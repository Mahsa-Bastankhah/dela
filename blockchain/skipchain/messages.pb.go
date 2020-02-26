// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package skipchain

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

type ForwardLink struct {
	From                 []byte   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Prepare              *any.Any `protobuf:"bytes,3,opt,name=prepare,proto3" json:"prepare,omitempty"`
	Commit               *any.Any `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwardLink) Reset()         { *m = ForwardLink{} }
func (m *ForwardLink) String() string { return proto.CompactTextString(m) }
func (*ForwardLink) ProtoMessage()    {}
func (*ForwardLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *ForwardLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardLink.Unmarshal(m, b)
}
func (m *ForwardLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardLink.Marshal(b, m, deterministic)
}
func (m *ForwardLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardLink.Merge(m, src)
}
func (m *ForwardLink) XXX_Size() int {
	return xxx_messageInfo_ForwardLink.Size(m)
}
func (m *ForwardLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardLink.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardLink proto.InternalMessageInfo

func (m *ForwardLink) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ForwardLink) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *ForwardLink) GetPrepare() *any.Any {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *ForwardLink) GetCommit() *any.Any {
	if m != nil {
		return m.Commit
	}
	return nil
}

type SkipBlockHeader struct {
	Height               uint64         `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	BaseHeight           uint64         `protobuf:"varint,2,opt,name=baseHeight,proto3" json:"baseHeight,omitempty"`
	MaximumHeight        uint64         `protobuf:"varint,3,opt,name=maximumHeight,proto3" json:"maximumHeight,omitempty"`
	GenesisID            []byte         `protobuf:"bytes,4,opt,name=genesisID,proto3" json:"genesisID,omitempty"`
	DataHash             []byte         `protobuf:"bytes,5,opt,name=dataHash,proto3" json:"dataHash,omitempty"`
	Backlinks            [][]byte       `protobuf:"bytes,6,rep,name=backlinks,proto3" json:"backlinks,omitempty"`
	Forwardlinks         []*ForwardLink `protobuf:"bytes,7,rep,name=forwardlinks,proto3" json:"forwardlinks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SkipBlockHeader) Reset()         { *m = SkipBlockHeader{} }
func (m *SkipBlockHeader) String() string { return proto.CompactTextString(m) }
func (*SkipBlockHeader) ProtoMessage()    {}
func (*SkipBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *SkipBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkipBlockHeader.Unmarshal(m, b)
}
func (m *SkipBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkipBlockHeader.Marshal(b, m, deterministic)
}
func (m *SkipBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkipBlockHeader.Merge(m, src)
}
func (m *SkipBlockHeader) XXX_Size() int {
	return xxx_messageInfo_SkipBlockHeader.Size(m)
}
func (m *SkipBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SkipBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SkipBlockHeader proto.InternalMessageInfo

func (m *SkipBlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SkipBlockHeader) GetBaseHeight() uint64 {
	if m != nil {
		return m.BaseHeight
	}
	return 0
}

func (m *SkipBlockHeader) GetMaximumHeight() uint64 {
	if m != nil {
		return m.MaximumHeight
	}
	return 0
}

func (m *SkipBlockHeader) GetGenesisID() []byte {
	if m != nil {
		return m.GenesisID
	}
	return nil
}

func (m *SkipBlockHeader) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *SkipBlockHeader) GetBacklinks() [][]byte {
	if m != nil {
		return m.Backlinks
	}
	return nil
}

func (m *SkipBlockHeader) GetForwardlinks() []*ForwardLink {
	if m != nil {
		return m.Forwardlinks
	}
	return nil
}

func init() {
	proto.RegisterType((*ForwardLink)(nil), "skipchain.ForwardLink")
	proto.RegisterType((*SkipBlockHeader)(nil), "skipchain.SkipBlockHeader")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x52, 0x53, 0x3b, 0xad, 0x15, 0x06, 0x29, 0xb1, 0x88, 0x84, 0xe2, 0xa1, 0x07,
	0xd9, 0x42, 0xbd, 0x79, 0x53, 0x44, 0x2a, 0x78, 0x8a, 0x4f, 0x30, 0x49, 0xb7, 0xc9, 0x92, 0x6e,
	0x36, 0xec, 0x6e, 0xd1, 0x9e, 0x7d, 0x00, 0x5f, 0x59, 0xba, 0x69, 0x63, 0x7b, 0xf1, 0xb6, 0xf3,
	0xfd, 0xdf, 0x30, 0xc3, 0x2c, 0x0c, 0x25, 0x37, 0x86, 0x72, 0x6e, 0x58, 0xad, 0x95, 0x55, 0xd8,
	0x33, 0xa5, 0xa8, 0xb3, 0x82, 0x44, 0x35, 0xbe, 0xce, 0x95, 0xca, 0xd7, 0x7c, 0xe6, 0x82, 0x74,
	0xb3, 0x9a, 0x51, 0xb5, 0x6d, 0xac, 0xc9, 0x8f, 0x07, 0xfd, 0x57, 0xa5, 0x3f, 0x49, 0x2f, 0xdf,
	0x45, 0x55, 0x22, 0x42, 0x67, 0xa5, 0x95, 0x8c, 0xbc, 0xd8, 0x9b, 0x0e, 0x12, 0xf7, 0xc6, 0x21,
	0xf8, 0x56, 0x45, 0xbe, 0x23, 0xbe, 0x55, 0xc8, 0xa0, 0x5b, 0x6b, 0x5e, 0x93, 0xe6, 0x51, 0x10,
	0x7b, 0xd3, 0xfe, 0xfc, 0x8a, 0x35, 0x03, 0xd8, 0x61, 0x00, 0x7b, 0xaa, 0xb6, 0xc9, 0x41, 0xc2,
	0x7b, 0x08, 0x33, 0x25, 0xa5, 0xb0, 0x51, 0xe7, 0x1f, 0x7d, 0xef, 0x4c, 0xbe, 0x7d, 0xb8, 0xfc,
	0x28, 0x45, 0xfd, 0xbc, 0x56, 0x59, 0xb9, 0xe0, 0xb4, 0xe4, 0x1a, 0x47, 0x10, 0x16, 0x5c, 0xe4,
	0x85, 0x75, 0x7b, 0x75, 0x92, 0x7d, 0x85, 0xb7, 0x00, 0x29, 0x19, 0xbe, 0x68, 0x32, 0xdf, 0x65,
	0x47, 0x04, 0xef, 0xe0, 0x42, 0xd2, 0x97, 0x90, 0x1b, 0xb9, 0x57, 0x02, 0xa7, 0x9c, 0x42, 0xbc,
	0x81, 0x5e, 0xce, 0x2b, 0x6e, 0x84, 0x79, 0x7b, 0x71, 0x2b, 0x0e, 0x92, 0x3f, 0x80, 0x63, 0x38,
	0x5f, 0x92, 0xa5, 0x05, 0x99, 0x22, 0x3a, 0x73, 0x61, 0x5b, 0xef, 0x3a, 0x53, 0xca, 0xca, 0xb5,
	0xa8, 0x4a, 0x13, 0x85, 0x71, 0xb0, 0xeb, 0x6c, 0x01, 0x3e, 0xc2, 0x60, 0xd5, 0x9c, 0xb6, 0x11,
	0xba, 0x71, 0x30, 0xed, 0xcf, 0x47, 0xac, 0xfd, 0x18, 0x76, 0x74, 0xf9, 0xe4, 0xc4, 0x4d, 0x43,
	0x77, 0x9b, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0xb0, 0x74, 0xc0, 0xd6, 0x01, 0x00,
	0x00,
}

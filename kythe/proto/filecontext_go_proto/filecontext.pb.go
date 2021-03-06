// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kythe/proto/filecontext.proto

package filecontext_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContextDependentVersion struct {
	Row                  []*ContextDependentVersion_Row `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ContextDependentVersion) Reset()         { *m = ContextDependentVersion{} }
func (m *ContextDependentVersion) String() string { return proto.CompactTextString(m) }
func (*ContextDependentVersion) ProtoMessage()    {}
func (*ContextDependentVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_filecontext_2ca3a2408c9dc5a2, []int{0}
}
func (m *ContextDependentVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContextDependentVersion.Unmarshal(m, b)
}
func (m *ContextDependentVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContextDependentVersion.Marshal(b, m, deterministic)
}
func (dst *ContextDependentVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContextDependentVersion.Merge(dst, src)
}
func (m *ContextDependentVersion) XXX_Size() int {
	return xxx_messageInfo_ContextDependentVersion.Size(m)
}
func (m *ContextDependentVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ContextDependentVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ContextDependentVersion proto.InternalMessageInfo

func (m *ContextDependentVersion) GetRow() []*ContextDependentVersion_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

type ContextDependentVersion_Column struct {
	Offset               int32    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	LinkedContext        string   `protobuf:"bytes,2,opt,name=linked_context,json=linkedContext,proto3" json:"linked_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContextDependentVersion_Column) Reset()         { *m = ContextDependentVersion_Column{} }
func (m *ContextDependentVersion_Column) String() string { return proto.CompactTextString(m) }
func (*ContextDependentVersion_Column) ProtoMessage()    {}
func (*ContextDependentVersion_Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_filecontext_2ca3a2408c9dc5a2, []int{0, 0}
}
func (m *ContextDependentVersion_Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContextDependentVersion_Column.Unmarshal(m, b)
}
func (m *ContextDependentVersion_Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContextDependentVersion_Column.Marshal(b, m, deterministic)
}
func (dst *ContextDependentVersion_Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContextDependentVersion_Column.Merge(dst, src)
}
func (m *ContextDependentVersion_Column) XXX_Size() int {
	return xxx_messageInfo_ContextDependentVersion_Column.Size(m)
}
func (m *ContextDependentVersion_Column) XXX_DiscardUnknown() {
	xxx_messageInfo_ContextDependentVersion_Column.DiscardUnknown(m)
}

var xxx_messageInfo_ContextDependentVersion_Column proto.InternalMessageInfo

func (m *ContextDependentVersion_Column) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ContextDependentVersion_Column) GetLinkedContext() string {
	if m != nil {
		return m.LinkedContext
	}
	return ""
}

type ContextDependentVersion_Row struct {
	SourceContext        string                            `protobuf:"bytes,1,opt,name=source_context,json=sourceContext,proto3" json:"source_context,omitempty"`
	Column               []*ContextDependentVersion_Column `protobuf:"bytes,2,rep,name=column,proto3" json:"column,omitempty"`
	AlwaysProcess        bool                              `protobuf:"varint,3,opt,name=always_process,json=alwaysProcess,proto3" json:"always_process,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ContextDependentVersion_Row) Reset()         { *m = ContextDependentVersion_Row{} }
func (m *ContextDependentVersion_Row) String() string { return proto.CompactTextString(m) }
func (*ContextDependentVersion_Row) ProtoMessage()    {}
func (*ContextDependentVersion_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_filecontext_2ca3a2408c9dc5a2, []int{0, 1}
}
func (m *ContextDependentVersion_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContextDependentVersion_Row.Unmarshal(m, b)
}
func (m *ContextDependentVersion_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContextDependentVersion_Row.Marshal(b, m, deterministic)
}
func (dst *ContextDependentVersion_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContextDependentVersion_Row.Merge(dst, src)
}
func (m *ContextDependentVersion_Row) XXX_Size() int {
	return xxx_messageInfo_ContextDependentVersion_Row.Size(m)
}
func (m *ContextDependentVersion_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ContextDependentVersion_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ContextDependentVersion_Row proto.InternalMessageInfo

func (m *ContextDependentVersion_Row) GetSourceContext() string {
	if m != nil {
		return m.SourceContext
	}
	return ""
}

func (m *ContextDependentVersion_Row) GetColumn() []*ContextDependentVersion_Column {
	if m != nil {
		return m.Column
	}
	return nil
}

func (m *ContextDependentVersion_Row) GetAlwaysProcess() bool {
	if m != nil {
		return m.AlwaysProcess
	}
	return false
}

func init() {
	proto.RegisterType((*ContextDependentVersion)(nil), "kythe.proto.ContextDependentVersion")
	proto.RegisterType((*ContextDependentVersion_Column)(nil), "kythe.proto.ContextDependentVersion.Column")
	proto.RegisterType((*ContextDependentVersion_Row)(nil), "kythe.proto.ContextDependentVersion.Row")
}

func init() {
	proto.RegisterFile("kythe/proto/filecontext.proto", fileDescriptor_filecontext_2ca3a2408c9dc5a2)
}

var fileDescriptor_filecontext_2ca3a2408c9dc5a2 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x18, 0x84, 0xc9, 0x16, 0x8b, 0x66, 0x59, 0x0f, 0x45, 0xb4, 0x2c, 0x88, 0x45, 0x10, 0x02, 0x42,
	0x16, 0xf4, 0x20, 0x78, 0xb4, 0x82, 0x57, 0xc9, 0xc1, 0x83, 0x97, 0xb2, 0xa6, 0x7f, 0x6b, 0xd9,
	0x6c, 0xfe, 0x92, 0x64, 0xad, 0xfb, 0x26, 0xbe, 0x8a, 0x6f, 0x27, 0x4d, 0x82, 0xf4, 0x22, 0x78,
	0x0a, 0xff, 0x37, 0x33, 0xcc, 0x10, 0x7a, 0xbe, 0xd9, 0xbb, 0x77, 0x58, 0xf5, 0x06, 0x1d, 0xae,
	0x9a, 0x4e, 0x81, 0x44, 0xed, 0xe0, 0xd3, 0x71, 0x4f, 0xb2, 0xb9, 0x97, 0xc3, 0x71, 0xf9, 0x3d,
	0xa3, 0x67, 0x65, 0x90, 0x1f, 0xa1, 0x07, 0x5d, 0x83, 0x76, 0x2f, 0x60, 0x6c, 0x87, 0x3a, 0xbb,
	0xa7, 0x89, 0xc1, 0x21, 0x27, 0x45, 0xc2, 0xe6, 0x37, 0x8c, 0x4f, 0x62, 0xfc, 0x8f, 0x08, 0x17,
	0x38, 0x88, 0x31, 0xb4, 0x7c, 0xa2, 0x69, 0x89, 0x6a, 0xb7, 0xd5, 0xd9, 0x29, 0x4d, 0xb1, 0x69,
	0x2c, 0xb8, 0x9c, 0x14, 0x84, 0x1d, 0x88, 0x78, 0x65, 0x57, 0xf4, 0x58, 0x75, 0x7a, 0x03, 0x75,
	0x15, 0xe7, 0xe5, 0xb3, 0x82, 0xb0, 0x23, 0xb1, 0x08, 0x34, 0x36, 0x2c, 0xbf, 0x08, 0x4d, 0x04,
	0x0e, 0xa3, 0xdd, 0xe2, 0xce, 0x48, 0xf8, 0xb5, 0x93, 0x60, 0x0f, 0x34, 0xda, 0xb3, 0x92, 0xa6,
	0xd2, 0xf7, 0xe6, 0x33, 0x3f, 0xfb, 0xfa, 0x5f, 0xb3, 0xc3, 0x54, 0x11, 0xa3, 0x63, 0xd7, 0x5a,
	0x0d, 0xeb, 0xbd, 0xad, 0x7a, 0x83, 0x12, 0xac, 0xcd, 0x93, 0x82, 0xb0, 0x43, 0xb1, 0x08, 0xf4,
	0x39, 0xc0, 0x87, 0x3b, 0x7a, 0x21, 0x71, 0xcb, 0x5b, 0xc4, 0x56, 0x01, 0xaf, 0xe1, 0xc3, 0x21,
	0x2a, 0x3b, 0x2d, 0x7c, 0x3d, 0x99, 0x7c, 0x7f, 0xd5, 0x62, 0xe5, 0xe9, 0x5b, 0xea, 0x9f, 0xdb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x32, 0x17, 0x89, 0xa9, 0x01, 0x00, 0x00,
}

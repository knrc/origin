// Code generated by protoc-gen-go.
// source: github.com/golang/protobuf/ptypes/wrappers/wrappers.proto
// DO NOT EDIT!

/*
Package wrappers is a generated protocol buffer package.

It is generated from these files:
	github.com/golang/protobuf/ptypes/wrappers/wrappers.proto

It has these top-level messages:
	DoubleValue
	FloatValue
	Int64Value
	UInt64Value
	Int32Value
	UInt32Value
	BoolValue
	StringValue
	BytesValue
*/
package wrappers

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

// Wrapper message for `double`.
//
// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue struct {
	// The double value.
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
}

func (m *DoubleValue) Reset()                    { *m = DoubleValue{} }
func (m *DoubleValue) String() string            { return proto.CompactTextString(m) }
func (*DoubleValue) ProtoMessage()               {}
func (*DoubleValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (*DoubleValue) XXX_WellKnownType() string   { return "DoubleValue" }

// Wrapper message for `float`.
//
// The JSON representation for `FloatValue` is JSON number.
type FloatValue struct {
	// The float value.
	Value float32 `protobuf:"fixed32,1,opt,name=value" json:"value,omitempty"`
}

func (m *FloatValue) Reset()                    { *m = FloatValue{} }
func (m *FloatValue) String() string            { return proto.CompactTextString(m) }
func (*FloatValue) ProtoMessage()               {}
func (*FloatValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }
func (*FloatValue) XXX_WellKnownType() string   { return "FloatValue" }

// Wrapper message for `int64`.
//
// The JSON representation for `Int64Value` is JSON string.
type Int64Value struct {
	// The int64 value.
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int64Value) Reset()                    { *m = Int64Value{} }
func (m *Int64Value) String() string            { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()               {}
func (*Int64Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }
func (*Int64Value) XXX_WellKnownType() string   { return "Int64Value" }

// Wrapper message for `uint64`.
//
// The JSON representation for `UInt64Value` is JSON string.
type UInt64Value struct {
	// The uint64 value.
	Value uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UInt64Value) Reset()                    { *m = UInt64Value{} }
func (m *UInt64Value) String() string            { return proto.CompactTextString(m) }
func (*UInt64Value) ProtoMessage()               {}
func (*UInt64Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }
func (*UInt64Value) XXX_WellKnownType() string   { return "UInt64Value" }

// Wrapper message for `int32`.
//
// The JSON representation for `Int32Value` is JSON number.
type Int32Value struct {
	// The int32 value.
	Value int32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *Int32Value) Reset()                    { *m = Int32Value{} }
func (m *Int32Value) String() string            { return proto.CompactTextString(m) }
func (*Int32Value) ProtoMessage()               {}
func (*Int32Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }
func (*Int32Value) XXX_WellKnownType() string   { return "Int32Value" }

// Wrapper message for `uint32`.
//
// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value struct {
	// The uint32 value.
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UInt32Value) Reset()                    { *m = UInt32Value{} }
func (m *UInt32Value) String() string            { return proto.CompactTextString(m) }
func (*UInt32Value) ProtoMessage()               {}
func (*UInt32Value) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }
func (*UInt32Value) XXX_WellKnownType() string   { return "UInt32Value" }

// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
type BoolValue struct {
	// The bool value.
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *BoolValue) Reset()                    { *m = BoolValue{} }
func (m *BoolValue) String() string            { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()               {}
func (*BoolValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }
func (*BoolValue) XXX_WellKnownType() string   { return "BoolValue" }

// Wrapper message for `string`.
//
// The JSON representation for `StringValue` is JSON string.
type StringValue struct {
	// The string value.
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (m *StringValue) String() string            { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }
func (*StringValue) XXX_WellKnownType() string   { return "StringValue" }

// Wrapper message for `bytes`.
//
// The JSON representation for `BytesValue` is JSON string.
type BytesValue struct {
	// The bytes value.
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *BytesValue) Reset()                    { *m = BytesValue{} }
func (m *BytesValue) String() string            { return proto.CompactTextString(m) }
func (*BytesValue) ProtoMessage()               {}
func (*BytesValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }
func (*BytesValue) XXX_WellKnownType() string   { return "BytesValue" }

func init() {
	proto.RegisterType((*DoubleValue)(nil), "google.protobuf.DoubleValue")
	proto.RegisterType((*FloatValue)(nil), "google.protobuf.FloatValue")
	proto.RegisterType((*Int64Value)(nil), "google.protobuf.Int64Value")
	proto.RegisterType((*UInt64Value)(nil), "google.protobuf.UInt64Value")
	proto.RegisterType((*Int32Value)(nil), "google.protobuf.Int32Value")
	proto.RegisterType((*UInt32Value)(nil), "google.protobuf.UInt32Value")
	proto.RegisterType((*BoolValue)(nil), "google.protobuf.BoolValue")
	proto.RegisterType((*StringValue)(nil), "google.protobuf.StringValue")
	proto.RegisterType((*BytesValue)(nil), "google.protobuf.BytesValue")
}

func init() {
	proto.RegisterFile("github.com/golang/protobuf/ptypes/wrappers/wrappers.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x2f,
	0x4a, 0x2c, 0x28, 0x48, 0x2d, 0x42, 0x30, 0xf4, 0xc0, 0x2a, 0x84, 0xf8, 0xd3, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0x60, 0xea, 0x95, 0x94, 0xb9, 0xb8, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0xc3,
	0x12, 0x73, 0x4a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x40, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xc6, 0x20, 0x08, 0x47, 0x49, 0x89, 0x8b, 0xcb, 0x2d, 0x27, 0x3f, 0xb1, 0x04, 0x8b, 0x1a, 0x26,
	0x24, 0x35, 0x9e, 0x79, 0x25, 0x66, 0x26, 0x58, 0xd4, 0x30, 0xc3, 0xd4, 0x28, 0x73, 0x71, 0x87,
	0xe2, 0x52, 0xc4, 0x82, 0x6a, 0x90, 0xb1, 0x11, 0x16, 0x35, 0xac, 0x68, 0x06, 0x61, 0x55, 0xc4,
	0x0b, 0x53, 0xa4, 0xc8, 0xc5, 0xe9, 0x94, 0x9f, 0x9f, 0x83, 0x45, 0x09, 0x07, 0x92, 0x39, 0xc1,
	0x25, 0x45, 0x99, 0x79, 0xe9, 0x58, 0x14, 0x71, 0x22, 0x39, 0xc8, 0xa9, 0xb2, 0x24, 0xb5, 0x18,
	0x8b, 0x1a, 0x1e, 0xa8, 0x1a, 0xa7, 0x7a, 0x2e, 0xe1, 0xe4, 0xfc, 0x5c, 0x3d, 0xb4, 0xd0, 0x75,
	0xe2, 0x0d, 0x87, 0x06, 0x7f, 0x00, 0x48, 0x24, 0x80, 0x31, 0x4a, 0x8b, 0xf8, 0xa8, 0x5b, 0xc0,
	0xc8, 0xf8, 0x83, 0x91, 0x71, 0x11, 0x13, 0xb3, 0x7b, 0x80, 0xd3, 0x2a, 0x26, 0x39, 0x77, 0x88,
	0xd1, 0x01, 0x50, 0xd5, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x79, 0xf9, 0xe5, 0x79, 0x21, 0x20,
	0x5d, 0x49, 0x6c, 0x60, 0x63, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xdf, 0x64, 0x4b,
	0x1c, 0x02, 0x00, 0x00,
}

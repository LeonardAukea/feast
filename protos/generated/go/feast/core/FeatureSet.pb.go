// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/FeatureSet.proto

package core // import "github.com/gojek/feast/protos/generated/go/feast/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gojek/feast/protos/generated/go/feast/types"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeatureSetSpec struct {
	// Name of the featureSet. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// FeatureSet version.
	Version int32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// List of entities contained within this featureSet.
	// This allows the feature to be used during joins between feature sets.
	// If the featureSet is ingested into a store that supports keys, this value
	// will be made a key.
	Entities []*EntitySpec `protobuf:"bytes,3,rep,name=entities,proto3" json:"entities,omitempty"`
	// List of features contained within this featureSet.
	Features []*FeatureSpec `protobuf:"bytes,4,rep,name=features,proto3" json:"features,omitempty"`
	// Features in this feature set will only be retrieved if they are found
	// after [time - max_age]. Missing or older feature values will be returned
	// as nulls and indicated to end user
	MaxAge *duration.Duration `protobuf:"bytes,5,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	// Source on which feature rows can be found
	Source               *Source  `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureSetSpec) Reset()         { *m = FeatureSetSpec{} }
func (m *FeatureSetSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSetSpec) ProtoMessage()    {}
func (*FeatureSetSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSet_ddc40084f83105c3, []int{0}
}
func (m *FeatureSetSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSetSpec.Unmarshal(m, b)
}
func (m *FeatureSetSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSetSpec.Marshal(b, m, deterministic)
}
func (dst *FeatureSetSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSetSpec.Merge(dst, src)
}
func (m *FeatureSetSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSetSpec.Size(m)
}
func (m *FeatureSetSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSetSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSetSpec proto.InternalMessageInfo

func (m *FeatureSetSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSetSpec) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FeatureSetSpec) GetEntities() []*EntitySpec {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *FeatureSetSpec) GetFeatures() []*FeatureSpec {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *FeatureSetSpec) GetMaxAge() *duration.Duration {
	if m != nil {
		return m.MaxAge
	}
	return nil
}

func (m *FeatureSetSpec) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type EntitySpec struct {
	// Name of the entity.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value type of the feature.
	ValueType            types.ValueType_Enum `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=feast.types.ValueType_Enum" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EntitySpec) Reset()         { *m = EntitySpec{} }
func (m *EntitySpec) String() string { return proto.CompactTextString(m) }
func (*EntitySpec) ProtoMessage()    {}
func (*EntitySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSet_ddc40084f83105c3, []int{1}
}
func (m *EntitySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntitySpec.Unmarshal(m, b)
}
func (m *EntitySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntitySpec.Marshal(b, m, deterministic)
}
func (dst *EntitySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntitySpec.Merge(dst, src)
}
func (m *EntitySpec) XXX_Size() int {
	return xxx_messageInfo_EntitySpec.Size(m)
}
func (m *EntitySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EntitySpec.DiscardUnknown(m)
}

var xxx_messageInfo_EntitySpec proto.InternalMessageInfo

func (m *EntitySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntitySpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_INVALID
}

type FeatureSpec struct {
	// Name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value type of the feature.
	ValueType            types.ValueType_Enum `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=feast.types.ValueType_Enum" json:"value_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FeatureSpec) Reset()         { *m = FeatureSpec{} }
func (m *FeatureSpec) String() string { return proto.CompactTextString(m) }
func (*FeatureSpec) ProtoMessage()    {}
func (*FeatureSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_FeatureSet_ddc40084f83105c3, []int{2}
}
func (m *FeatureSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSpec.Unmarshal(m, b)
}
func (m *FeatureSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSpec.Marshal(b, m, deterministic)
}
func (dst *FeatureSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSpec.Merge(dst, src)
}
func (m *FeatureSpec) XXX_Size() int {
	return xxx_messageInfo_FeatureSpec.Size(m)
}
func (m *FeatureSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSpec proto.InternalMessageInfo

func (m *FeatureSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeatureSpec) GetValueType() types.ValueType_Enum {
	if m != nil {
		return m.ValueType
	}
	return types.ValueType_INVALID
}

func init() {
	proto.RegisterType((*FeatureSetSpec)(nil), "feast.core.FeatureSetSpec")
	proto.RegisterType((*EntitySpec)(nil), "feast.core.EntitySpec")
	proto.RegisterType((*FeatureSpec)(nil), "feast.core.FeatureSpec")
}

func init() {
	proto.RegisterFile("feast/core/FeatureSet.proto", fileDescriptor_FeatureSet_ddc40084f83105c3)
}

var fileDescriptor_FeatureSet_ddc40084f83105c3 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x10, 0xc5, 0xf9, 0x70, 0x92, 0x09, 0x64, 0x41, 0x87, 0x8d, 0x77, 0x03, 0x8b, 0xc9, 0xc9, 0xec,
	0x41, 0x02, 0x87, 0x5e, 0x7a, 0x6b, 0x68, 0x7b, 0x2e, 0x4e, 0xe8, 0xa1, 0xb4, 0x04, 0xc5, 0x99,
	0xa8, 0x6e, 0x63, 0xcb, 0xd8, 0x72, 0x48, 0x7e, 0x41, 0xff, 0x76, 0xb1, 0x14, 0xd5, 0x3e, 0xf4,
	0xd8, 0x9b, 0xc4, 0x7b, 0xf3, 0x66, 0xde, 0x9b, 0x81, 0xd9, 0x1e, 0x79, 0xa9, 0x58, 0x2c, 0x0b,
	0x64, 0xf7, 0xc8, 0x55, 0x55, 0xe0, 0x0a, 0x15, 0xcd, 0x0b, 0xa9, 0x24, 0x01, 0x0d, 0xd2, 0x1a,
	0xfc, 0x3b, 0x35, 0x44, 0x75, 0xce, 0xb1, 0x64, 0x8f, 0xfc, 0x50, 0xa1, 0x21, 0x59, 0x40, 0x2b,
	0xac, 0x64, 0x55, 0xc4, 0x16, 0xf8, 0x27, 0xa4, 0x14, 0x07, 0x64, 0xfa, 0xb7, 0xad, 0xf6, 0x6c,
	0x57, 0x15, 0x5c, 0x25, 0x32, 0x33, 0xf8, 0xfc, 0xa3, 0x03, 0x93, 0xa6, 0xe5, 0x2a, 0xc7, 0x98,
	0x10, 0xe8, 0x65, 0x3c, 0x45, 0xcf, 0xf1, 0x9d, 0x60, 0x14, 0xe9, 0x37, 0xf1, 0x60, 0x70, 0xc4,
	0xa2, 0x4c, 0x64, 0xe6, 0x75, 0x7c, 0x27, 0xe8, 0x47, 0xf6, 0x4b, 0x42, 0x18, 0x62, 0xa6, 0x12,
	0x95, 0x60, 0xe9, 0x75, 0xfd, 0x6e, 0x30, 0x0e, 0x7f, 0xd3, 0x66, 0x62, 0x7a, 0x57, 0x63, 0xe7,
	0x5a, 0x37, 0xfa, 0xe2, 0x91, 0x05, 0x0c, 0xf7, 0xa6, 0x67, 0xe9, 0xf5, 0x74, 0xcd, 0xb4, 0x5d,
	0x63, 0xe7, 0xd1, 0x45, 0x96, 0x48, 0x42, 0x18, 0xa4, 0xfc, 0xb4, 0xe1, 0x02, 0xbd, 0xbe, 0xef,
	0x04, 0xe3, 0xf0, 0x0f, 0x35, 0xde, 0xa8, 0xf5, 0x46, 0x6f, 0x2f, 0xde, 0x22, 0x37, 0xe5, 0xa7,
	0x1b, 0x81, 0xe4, 0x3f, 0xb8, 0xa5, 0x4e, 0xc3, 0x73, 0x75, 0x09, 0x69, 0xb7, 0x31, 0x39, 0x45,
	0x17, 0xc6, 0xfc, 0x19, 0xa0, 0x19, 0xf6, 0xdb, 0x10, 0xae, 0x01, 0x8e, 0x75, 0xe6, 0x9b, 0x3a,
	0x7f, 0x9d, 0xc3, 0x24, 0x9c, 0x5d, 0x14, 0xf5, 0x4a, 0xa8, 0x5e, 0xc9, 0xfa, 0x9c, 0xd7, 0xbe,
	0xab, 0x34, 0x1a, 0x1d, 0xed, 0x7f, 0xfe, 0x02, 0xe3, 0x96, 0xad, 0x9f, 0x96, 0x5f, 0xae, 0xa1,
	0x75, 0x26, 0xcb, 0x5f, 0xcd, 0x46, 0x1f, 0xea, 0x6c, 0x9e, 0xae, 0x44, 0xa2, 0x5e, 0xab, 0x2d,
	0x8d, 0x65, 0xca, 0x84, 0x7c, 0xc3, 0x77, 0x66, 0xee, 0x45, 0x27, 0x57, 0x32, 0x81, 0x19, 0x16,
	0x5c, 0xe1, 0x8e, 0x09, 0xc9, 0x9a, 0x4b, 0xda, 0xba, 0x1a, 0x5f, 0x7c, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x3d, 0xbe, 0x9e, 0x0d, 0xa0, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/variable.proto

package tensorflow

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Protocol buffer representing a Variable.
type VariableDef struct {
	// Name of the variable tensor.
	VariableName string `protobuf:"bytes,1,opt,name=variable_name,json=variableName,proto3" json:"variable_name,omitempty"`
	// Name of the tensor holding the variable's initial value.
	InitialValueName string `protobuf:"bytes,6,opt,name=initial_value_name,json=initialValueName,proto3" json:"initial_value_name,omitempty"`
	// Name of the initializer op.
	InitializerName string `protobuf:"bytes,2,opt,name=initializer_name,json=initializerName,proto3" json:"initializer_name,omitempty"`
	// Name of the snapshot tensor.
	SnapshotName string `protobuf:"bytes,3,opt,name=snapshot_name,json=snapshotName,proto3" json:"snapshot_name,omitempty"`
	// Support for saving variables as slices of a larger variable.
	SaveSliceInfoDef *SaveSliceInfoDef `protobuf:"bytes,4,opt,name=save_slice_info_def,json=saveSliceInfoDef,proto3" json:"save_slice_info_def,omitempty"`
	// Whether to represent this as a ResourceVariable.
	IsResource           bool     `protobuf:"varint,5,opt,name=is_resource,json=isResource,proto3" json:"is_resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VariableDef) Reset()         { *m = VariableDef{} }
func (m *VariableDef) String() string { return proto.CompactTextString(m) }
func (*VariableDef) ProtoMessage()    {}
func (*VariableDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_908f2d03adae2778, []int{0}
}

func (m *VariableDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariableDef.Unmarshal(m, b)
}
func (m *VariableDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariableDef.Marshal(b, m, deterministic)
}
func (m *VariableDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariableDef.Merge(m, src)
}
func (m *VariableDef) XXX_Size() int {
	return xxx_messageInfo_VariableDef.Size(m)
}
func (m *VariableDef) XXX_DiscardUnknown() {
	xxx_messageInfo_VariableDef.DiscardUnknown(m)
}

var xxx_messageInfo_VariableDef proto.InternalMessageInfo

func (m *VariableDef) GetVariableName() string {
	if m != nil {
		return m.VariableName
	}
	return ""
}

func (m *VariableDef) GetInitialValueName() string {
	if m != nil {
		return m.InitialValueName
	}
	return ""
}

func (m *VariableDef) GetInitializerName() string {
	if m != nil {
		return m.InitializerName
	}
	return ""
}

func (m *VariableDef) GetSnapshotName() string {
	if m != nil {
		return m.SnapshotName
	}
	return ""
}

func (m *VariableDef) GetSaveSliceInfoDef() *SaveSliceInfoDef {
	if m != nil {
		return m.SaveSliceInfoDef
	}
	return nil
}

func (m *VariableDef) GetIsResource() bool {
	if m != nil {
		return m.IsResource
	}
	return false
}

type SaveSliceInfoDef struct {
	// Name of the full variable of which this is a slice.
	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	// Shape of the full variable.
	FullShape []int64 `protobuf:"varint,2,rep,packed,name=full_shape,json=fullShape,proto3" json:"full_shape,omitempty"`
	// Offset of this variable into the full variable.
	VarOffset []int64 `protobuf:"varint,3,rep,packed,name=var_offset,json=varOffset,proto3" json:"var_offset,omitempty"`
	// Shape of this variable.
	VarShape             []int64  `protobuf:"varint,4,rep,packed,name=var_shape,json=varShape,proto3" json:"var_shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveSliceInfoDef) Reset()         { *m = SaveSliceInfoDef{} }
func (m *SaveSliceInfoDef) String() string { return proto.CompactTextString(m) }
func (*SaveSliceInfoDef) ProtoMessage()    {}
func (*SaveSliceInfoDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_908f2d03adae2778, []int{1}
}

func (m *SaveSliceInfoDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveSliceInfoDef.Unmarshal(m, b)
}
func (m *SaveSliceInfoDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveSliceInfoDef.Marshal(b, m, deterministic)
}
func (m *SaveSliceInfoDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveSliceInfoDef.Merge(m, src)
}
func (m *SaveSliceInfoDef) XXX_Size() int {
	return xxx_messageInfo_SaveSliceInfoDef.Size(m)
}
func (m *SaveSliceInfoDef) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveSliceInfoDef.DiscardUnknown(m)
}

var xxx_messageInfo_SaveSliceInfoDef proto.InternalMessageInfo

func (m *SaveSliceInfoDef) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *SaveSliceInfoDef) GetFullShape() []int64 {
	if m != nil {
		return m.FullShape
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarOffset() []int64 {
	if m != nil {
		return m.VarOffset
	}
	return nil
}

func (m *SaveSliceInfoDef) GetVarShape() []int64 {
	if m != nil {
		return m.VarShape
	}
	return nil
}

func init() {
	proto.RegisterType((*VariableDef)(nil), "tensorflow.VariableDef")
	proto.RegisterType((*SaveSliceInfoDef)(nil), "tensorflow.SaveSliceInfoDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/variable.proto", fileDescriptor_908f2d03adae2778)
}

var fileDescriptor_908f2d03adae2778 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x40, 0x53, 0x86, 0x04, 0x2e, 0xa2, 0x64, 0xbe, 0x2c, 0x51, 0x23, 0x81, 0x97, 0x99, 0x18,
	0x96, 0xe8, 0x1f, 0x10, 0x5e, 0x8c, 0x89, 0x92, 0x91, 0xf0, 0xba, 0x5c, 0xf0, 0x56, 0x1a, 0xcb,
	0x4a, 0xda, 0x51, 0x12, 0x3f, 0xc1, 0x8f, 0xf0, 0x3b, 0x7d, 0x34, 0xed, 0x36, 0x21, 0xbc, 0x9e,
	0x73, 0xda, 0xed, 0xf6, 0x42, 0x5c, 0x50, 0x6e, 0x94, 0xe6, 0x52, 0xed, 0x93, 0x95, 0xd2, 0x94,
	0x70, 0x8d, 0x1b, 0xda, 0x2b, 0xfd, 0x99, 0x58, 0xd4, 0x02, 0x97, 0x92, 0xc6, 0x5b, 0xad, 0x0a,
	0x15, 0xc2, 0xa1, 0x1c, 0xfe, 0x34, 0xa0, 0xbb, 0xa8, 0xf4, 0x94, 0x78, 0x38, 0x82, 0x5e, 0x5d,
	0x67, 0x39, 0x6e, 0x28, 0x62, 0x03, 0x16, 0x77, 0xd2, 0xf3, 0x1a, 0xbe, 0xe2, 0x86, 0xc2, 0x07,
	0x08, 0x45, 0x2e, 0x0a, 0x81, 0x32, 0xb3, 0x28, 0x77, 0x55, 0xd9, 0xf2, 0x65, 0xbf, 0x32, 0x0b,
	0x27, 0x7c, 0x7d, 0x0f, 0x35, 0x13, 0x5f, 0xa4, 0xcb, 0xb6, 0xe1, 0xdb, 0xcb, 0x23, 0xee, 0xd3,
	0x11, 0xf4, 0x4c, 0x8e, 0x5b, 0xb3, 0x56, 0x45, 0xd9, 0x05, 0xe5, 0xd7, 0x6b, 0xe8, 0xa3, 0x17,
	0xb8, 0x32, 0x68, 0x29, 0x33, 0x52, 0xac, 0x28, 0x13, 0x39, 0x57, 0xd9, 0x3b, 0xf1, 0xa8, 0x39,
	0x60, 0x71, 0xf7, 0xf1, 0x66, 0x7c, 0x18, 0x6e, 0x3c, 0x47, 0x4b, 0x73, 0x57, 0x3d, 0xe7, 0x5c,
	0x4d, 0x89, 0xa7, 0x7d, 0x73, 0x42, 0xc2, 0x3b, 0xe8, 0x0a, 0x93, 0x69, 0x32, 0x6a, 0xa7, 0x57,
	0x14, 0x9d, 0x0d, 0x58, 0xdc, 0x4e, 0x41, 0x98, 0xb4, 0x22, 0xc3, 0x6f, 0x06, 0xfd, 0xd3, 0x7b,
	0xc2, 0x6b, 0xe8, 0xf0, 0x9d, 0x94, 0xc7, 0x2f, 0xd4, 0x76, 0xc0, 0xff, 0xdf, 0x2d, 0x80, 0x97,
	0x66, 0x8d, 0x5b, 0x37, 0x69, 0x10, 0x07, 0xa9, 0xcf, 0xe7, 0x0e, 0x38, 0x6d, 0x51, 0x67, 0x8a,
	0x73, 0x43, 0x45, 0x14, 0x94, 0xda, 0xa2, 0x7e, 0xf3, 0xc0, 0x5d, 0xed, 0x74, 0x79, 0xb8, 0xe9,
	0x6d, 0xdb, 0xa2, 0xf6, 0x67, 0x27, 0x09, 0x44, 0x4a, 0x7f, 0x1c, 0x8f, 0xf8, 0xbf, 0xe4, 0xc9,
	0x45, 0xbd, 0xc6, 0x99, 0x5b, 0xb2, 0x99, 0xb1, 0x5f, 0xc6, 0x96, 0x2d, 0xbf, 0xf1, 0xa7, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x12, 0x4f, 0x21, 0x1d, 0x02, 0x00, 0x00,
}

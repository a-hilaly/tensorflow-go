// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/config/logging_config.proto

package tensorflow_serving

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

type SamplingConfig struct {
	// Requests will be logged uniformly at random with this probability. Valid
	// range: [0, 1.0].
	SamplingRate         float64  `protobuf:"fixed64,1,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SamplingConfig) Reset()         { *m = SamplingConfig{} }
func (m *SamplingConfig) String() string { return proto.CompactTextString(m) }
func (*SamplingConfig) ProtoMessage()    {}
func (*SamplingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e467ec096fa84f9, []int{0}
}

func (m *SamplingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SamplingConfig.Unmarshal(m, b)
}
func (m *SamplingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SamplingConfig.Marshal(b, m, deterministic)
}
func (m *SamplingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SamplingConfig.Merge(m, src)
}
func (m *SamplingConfig) XXX_Size() int {
	return xxx_messageInfo_SamplingConfig.Size(m)
}
func (m *SamplingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SamplingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SamplingConfig proto.InternalMessageInfo

func (m *SamplingConfig) GetSamplingRate() float64 {
	if m != nil {
		return m.SamplingRate
	}
	return 0
}

// Configuration for logging query/responses.
type LoggingConfig struct {
	LogCollectorConfig   *LogCollectorConfig `protobuf:"bytes,1,opt,name=log_collector_config,json=logCollectorConfig,proto3" json:"log_collector_config,omitempty"`
	SamplingConfig       *SamplingConfig     `protobuf:"bytes,2,opt,name=sampling_config,json=samplingConfig,proto3" json:"sampling_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoggingConfig) Reset()         { *m = LoggingConfig{} }
func (m *LoggingConfig) String() string { return proto.CompactTextString(m) }
func (*LoggingConfig) ProtoMessage()    {}
func (*LoggingConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e467ec096fa84f9, []int{1}
}

func (m *LoggingConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoggingConfig.Unmarshal(m, b)
}
func (m *LoggingConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoggingConfig.Marshal(b, m, deterministic)
}
func (m *LoggingConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggingConfig.Merge(m, src)
}
func (m *LoggingConfig) XXX_Size() int {
	return xxx_messageInfo_LoggingConfig.Size(m)
}
func (m *LoggingConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggingConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LoggingConfig proto.InternalMessageInfo

func (m *LoggingConfig) GetLogCollectorConfig() *LogCollectorConfig {
	if m != nil {
		return m.LogCollectorConfig
	}
	return nil
}

func (m *LoggingConfig) GetSamplingConfig() *SamplingConfig {
	if m != nil {
		return m.SamplingConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*SamplingConfig)(nil), "tensorflow.serving.SamplingConfig")
	proto.RegisterType((*LoggingConfig)(nil), "tensorflow.serving.LoggingConfig")
}

func init() {
	proto.RegisterFile("tensorflow_serving/config/logging_config.proto", fileDescriptor_7e467ec096fa84f9)
}

var fileDescriptor_7e467ec096fa84f9 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc, 0x4b, 0xd7, 0x4f,
	0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0xc9, 0x4f, 0x4f, 0xcf, 0xcc, 0x4b, 0x8f, 0x87, 0x70,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0x10, 0xea, 0xf5, 0xa0, 0xea, 0xa5, 0x4c, 0xf0,
	0x9a, 0x11, 0x9f, 0x9c, 0x9f, 0x93, 0x93, 0x9a, 0x5c, 0x92, 0x5f, 0x84, 0x62, 0x92, 0x92, 0x29,
	0x17, 0x5f, 0x70, 0x62, 0x6e, 0x41, 0x4e, 0x66, 0x5e, 0xba, 0x33, 0x58, 0x5c, 0x48, 0x99, 0x8b,
	0xb7, 0x18, 0x2a, 0x12, 0x5f, 0x94, 0x58, 0x92, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x18, 0xc4,
	0x03, 0x13, 0x0c, 0x4a, 0x2c, 0x49, 0x55, 0xda, 0xc6, 0xc8, 0xc5, 0xeb, 0x03, 0x71, 0x19, 0x54,
	0x5b, 0x04, 0x97, 0x08, 0x36, 0x6b, 0xc0, 0xba, 0xb9, 0x8d, 0xd4, 0xf4, 0x30, 0x5d, 0xac, 0xe7,
	0x93, 0x9f, 0xee, 0x0c, 0x53, 0x0e, 0x31, 0x25, 0x48, 0x28, 0x07, 0x43, 0x4c, 0xc8, 0x9b, 0x8b,
	0x1f, 0xee, 0x20, 0xa8, 0xa1, 0x4c, 0x60, 0x43, 0x95, 0xb0, 0x19, 0x8a, 0xea, 0x9b, 0x20, 0xbe,
	0x62, 0x14, 0xbe, 0x13, 0xf3, 0x0f, 0x46, 0xc6, 0x24, 0x36, 0xb0, 0xdf, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x3f, 0xc0, 0xa7, 0xc5, 0x77, 0x01, 0x00, 0x00,
}

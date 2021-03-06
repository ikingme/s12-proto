// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logger.proto

package s12proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Level int32

const (
	Level_PANIC Level = 0
	Level_FATAL Level = 1
	Level_ERROR Level = 2
	Level_WARN  Level = 3
	Level_INFO  Level = 4
	Level_DEBUG Level = 5
)

var Level_name = map[int32]string{
	0: "PANIC",
	1: "FATAL",
	2: "ERROR",
	3: "WARN",
	4: "INFO",
	5: "DEBUG",
}

var Level_value = map[string]int32{
	"PANIC": 0,
	"FATAL": 1,
	"ERROR": 2,
	"WARN":  3,
	"INFO":  4,
	"DEBUG": 5,
}

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (x *Level) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Level_value, data, "Level")
	if err != nil {
		return err
	}
	*x = Level(value)
	return nil
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{0}
}

var E_Level = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Level)(nil),
	Field:         66200,
	Name:          "logger.level",
	Tag:           "varint,66200,opt,name=level,enum=logger.Level",
	Filename:      "logger.proto",
}

func init() {
	proto.RegisterEnum("logger.Level", Level_name, Level_value)
	proto.RegisterExtension(E_Level)
}

func init() { proto.RegisterFile("logger.proto", fileDescriptor_d43b7bfc6b6f7b16) }

var fileDescriptor_d43b7bfc6b6f7b16 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xc9, 0x4f, 0x4f,
	0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x14, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xa2, 0x49, 0xa5, 0x69, 0xfa, 0x29, 0xa9, 0xc5, 0xc9, 0x45,
	0x99, 0x05, 0x25, 0xf9, 0x50, 0x95, 0x5a, 0xee, 0x5c, 0xac, 0x3e, 0xa9, 0x65, 0xa9, 0x39, 0x42,
	0x9c, 0x5c, 0xac, 0x01, 0x8e, 0x7e, 0x9e, 0xce, 0x02, 0x0c, 0x20, 0xa6, 0x9b, 0x63, 0x88, 0xa3,
	0x8f, 0x00, 0x23, 0x88, 0xe9, 0x1a, 0x14, 0xe4, 0x1f, 0x24, 0xc0, 0x24, 0xc4, 0xc1, 0xc5, 0x12,
	0xee, 0x18, 0xe4, 0x27, 0xc0, 0x0c, 0x62, 0x79, 0xfa, 0xb9, 0xf9, 0x0b, 0xb0, 0x80, 0xa4, 0x5d,
	0x5c, 0x9d, 0x42, 0xdd, 0x05, 0x58, 0xad, 0x5c, 0xb8, 0x58, 0x73, 0xc0, 0x06, 0xc9, 0xea, 0x41,
	0x2c, 0xd5, 0x83, 0x59, 0xaa, 0xe7, 0x96, 0x99, 0x9a, 0x93, 0xe2, 0x5f, 0x50, 0x92, 0x99, 0x9f,
	0x57, 0x2c, 0x31, 0xa3, 0x95, 0x45, 0x81, 0x51, 0x83, 0xcf, 0x88, 0x57, 0x0f, 0xea, 0x62, 0xb0,
	0xf5, 0x41, 0x10, 0xcd, 0x4e, 0x66, 0x51, 0x26, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xc1, 0x89, 0x69, 0xa9, 0x25, 0x95, 0xce, 0xa5, 0x39, 0x25, 0xa5, 0x45, 0xa9,
	0xfa, 0xc5, 0x86, 0x46, 0xba, 0x60, 0x33, 0x11, 0xde, 0x29, 0x36, 0x34, 0x02, 0xb3, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x96, 0x26, 0xc6, 0xbb, 0xff, 0x00, 0x00, 0x00,
}

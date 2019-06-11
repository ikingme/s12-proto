// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: route_guide.proto

package s12_routeguide

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude             int32    `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            int32    `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{0}
}
func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLatitude() int32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Point) GetLongitude() int32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi                   *Point   `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rectangle) Reset()         { *m = Rectangle{} }
func (m *Rectangle) String() string { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()    {}
func (*Rectangle) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{1}
}
func (m *Rectangle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rectangle.Unmarshal(m, b)
}
func (m *Rectangle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rectangle.Marshal(b, m, deterministic)
}
func (m *Rectangle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rectangle.Merge(m, src)
}
func (m *Rectangle) XXX_Size() int {
	return xxx_messageInfo_Rectangle.Size(m)
}
func (m *Rectangle) XXX_DiscardUnknown() {
	xxx_messageInfo_Rectangle.DiscardUnknown(m)
}

var xxx_messageInfo_Rectangle proto.InternalMessageInfo

func (m *Rectangle) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Rectangle) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Feature struct {
	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The point where the feature is detected.
	Location             *Point   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{2}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (m *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(m, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The message to be sent.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteNote) Reset()         { *m = RouteNote{} }
func (m *RouteNote) String() string { return proto.CompactTextString(m) }
func (*RouteNote) ProtoMessage()    {}
func (*RouteNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{3}
}
func (m *RouteNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteNote.Unmarshal(m, b)
}
func (m *RouteNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteNote.Marshal(b, m, deterministic)
}
func (m *RouteNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteNote.Merge(m, src)
}
func (m *RouteNote) XXX_Size() int {
	return xxx_messageInfo_RouteNote.Size(m)
}
func (m *RouteNote) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteNote.DiscardUnknown(m)
}

var xxx_messageInfo_RouteNote proto.InternalMessageInfo

func (m *RouteNote) GetLocation() *Point {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RouteNote) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime          int32    `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteSummary) Reset()         { *m = RouteSummary{} }
func (m *RouteSummary) String() string { return proto.CompactTextString(m) }
func (*RouteSummary) ProtoMessage()    {}
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7d679f20da65b7b, []int{4}
}
func (m *RouteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSummary.Unmarshal(m, b)
}
func (m *RouteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSummary.Marshal(b, m, deterministic)
}
func (m *RouteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSummary.Merge(m, src)
}
func (m *RouteSummary) XXX_Size() int {
	return xxx_messageInfo_RouteSummary.Size(m)
}
func (m *RouteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSummary proto.InternalMessageInfo

func (m *RouteSummary) GetPointCount() int32 {
	if m != nil {
		return m.PointCount
	}
	return 0
}

func (m *RouteSummary) GetFeatureCount() int32 {
	if m != nil {
		return m.FeatureCount
	}
	return 0
}

func (m *RouteSummary) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *RouteSummary) GetElapsedTime() int32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "s12.routeguide.Point")
	proto.RegisterType((*Rectangle)(nil), "s12.routeguide.Rectangle")
	proto.RegisterType((*Feature)(nil), "s12.routeguide.Feature")
	proto.RegisterType((*RouteNote)(nil), "s12.routeguide.RouteNote")
	proto.RegisterType((*RouteSummary)(nil), "s12.routeguide.RouteSummary")
}

func init() { proto.RegisterFile("route_guide.proto", fileDescriptor_b7d679f20da65b7b) }

var fileDescriptor_b7d679f20da65b7b = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0xed, 0xa4, 0xad, 0xdd, 0xdc, 0x8d, 0x8a, 0x17, 0xc4, 0x18, 0x0b, 0x6a, 0x44, 0xe8, 0x53,
	0x68, 0x57, 0xf0, 0x51, 0xb0, 0x85, 0xcd, 0x8b, 0x48, 0x18, 0xfb, 0xa0, 0x4f, 0xcb, 0x98, 0x5c,
	0xb3, 0x03, 0x93, 0x4c, 0xc8, 0x4c, 0x40, 0x7f, 0x87, 0xff, 0x40, 0xf0, 0x7f, 0x4a, 0x26, 0xc9,
	0xf6, 0xc3, 0xee, 0xd2, 0xb7, 0xcc, 0xb9, 0xe7, 0x9e, 0x39, 0xf7, 0xdc, 0x0c, 0x3c, 0x69, 0x75,
	0x67, 0x69, 0x55, 0x76, 0xb2, 0xa0, 0xa4, 0x69, 0xb5, 0xd5, 0xf8, 0xc8, 0x9c, 0x2d, 0x12, 0x07,
	0x3b, 0x34, 0xfe, 0x08, 0x87, 0x99, 0x96, 0xb5, 0xc5, 0x08, 0x66, 0x4a, 0x58, 0x69, 0xbb, 0x82,
	0x42, 0xf6, 0x8a, 0x9d, 0x1c, 0xf2, 0xcd, 0x19, 0x8f, 0xc1, 0x57, 0xba, 0x2e, 0x87, 0xa2, 0xe7,
	0x8a, 0x57, 0x40, 0xfc, 0x0d, 0x7c, 0x4e, 0xb9, 0x15, 0x75, 0xa9, 0x08, 0xdf, 0x82, 0xa7, 0xb4,
	0x13, 0x98, 0x2f, 0x9e, 0x26, 0x37, 0x2f, 0x4b, 0xdc, 0x4d, 0xdc, 0x53, 0xba, 0xa7, 0xad, 0xa5,
	0x93, 0xda, 0x4e, 0x5b, 0xcb, 0x38, 0x83, 0xa3, 0x25, 0x09, 0xdb, 0xb5, 0x84, 0x08, 0x07, 0xb5,
	0xa8, 0x06, 0x6f, 0x3e, 0x77, 0xdf, 0x78, 0x06, 0x33, 0xa5, 0x73, 0x61, 0xa5, 0xae, 0x77, 0x6b,
	0x6d, 0x68, 0xf1, 0x57, 0xf0, 0x79, 0x5f, 0xfd, 0xac, 0xed, 0xcd, 0x7e, 0x76, 0xaf, 0x7e, 0x0c,
	0xe1, 0xa8, 0x22, 0x63, 0x44, 0x39, 0x04, 0xe1, 0xf3, 0xe9, 0x18, 0xff, 0x66, 0x10, 0x38, 0xe9,
	0x2f, 0x5d, 0x55, 0x89, 0xf6, 0x17, 0xbe, 0x84, 0x79, 0xd3, 0x77, 0xaf, 0x72, 0xdd, 0xd5, 0x76,
	0x0c, 0x15, 0x1c, 0x74, 0xd1, 0x23, 0xf8, 0x06, 0x1e, 0xfe, 0x18, 0xa6, 0x1b, 0x29, 0x43, 0xb4,
	0xc1, 0x08, 0x0e, 0xa4, 0x08, 0x66, 0x85, 0x34, 0x56, 0xd4, 0x39, 0x85, 0xfb, 0xc3, 0x5e, 0xa6,
	0x33, 0xbe, 0x86, 0x80, 0x94, 0x68, 0x0c, 0x15, 0x2b, 0x2b, 0x2b, 0x0a, 0x0f, 0x5c, 0x7d, 0x3e,
	0x62, 0x97, 0xb2, 0xa2, 0xc5, 0x5f, 0x0f, 0xc0, 0xb9, 0x4a, 0xfb, 0x71, 0xf0, 0x03, 0x40, 0x4a,
	0x76, 0xca, 0xf4, 0xee, 0x69, 0xa3, 0x67, 0xb7, 0xe1, 0x91, 0x1f, 0xef, 0xe1, 0x12, 0x82, 0x4f,
	0xd2, 0x4c, 0x02, 0x06, 0x9f, 0xdf, 0xa6, 0x6e, 0xfe, 0x84, 0x1d, 0x2a, 0xa7, 0x0c, 0x97, 0x30,
	0xe7, 0x94, 0xeb, 0xb6, 0x70, 0xde, 0xb6, 0x19, 0x39, 0xfe, 0x4f, 0xfd, 0x5a, 0xbe, 0xf1, 0xde,
	0x09, 0xc3, 0x74, 0x5c, 0xe7, 0xc5, 0x5a, 0xd8, 0x3b, 0xcc, 0x4c, 0x9b, 0x8e, 0xb6, 0x97, 0x7a,
	0x99, 0x53, 0x76, 0xfe, 0x1e, 0x5e, 0x48, 0x9d, 0x94, 0x6d, 0x93, 0x27, 0xf4, 0x53, 0x54, 0x8d,
	0x22, 0x73, 0x8d, 0x7e, 0xfe, 0xf8, 0x2a, 0xc3, 0xac, 0x7f, 0x47, 0x19, 0xfb, 0xe3, 0xed, 0xf3,
	0xcb, 0xf4, 0xfb, 0x03, 0xf7, 0xac, 0xde, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x89, 0xd1,
	0x93, 0x6b, 0x03, 0x00, 0x00,
}
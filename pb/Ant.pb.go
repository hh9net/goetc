// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Ant.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

//初始化
type EPAntInitRequest struct {
	Antkey               string   `protobuf:"bytes,1,opt,name=antkey,proto3" json:"antkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EPAntInitRequest) Reset()         { *m = EPAntInitRequest{} }
func (m *EPAntInitRequest) String() string { return proto.CompactTextString(m) }
func (*EPAntInitRequest) ProtoMessage()    {}
func (*EPAntInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e0e94b3ebffec, []int{0}
}

func (m *EPAntInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EPAntInitRequest.Unmarshal(m, b)
}
func (m *EPAntInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EPAntInitRequest.Marshal(b, m, deterministic)
}
func (m *EPAntInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EPAntInitRequest.Merge(m, src)
}
func (m *EPAntInitRequest) XXX_Size() int {
	return xxx_messageInfo_EPAntInitRequest.Size(m)
}
func (m *EPAntInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EPAntInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EPAntInitRequest proto.InternalMessageInfo

func (m *EPAntInitRequest) GetAntkey() string {
	if m != nil {
		return m.Antkey
	}
	return ""
}

type EPAntInitResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EPAntInitResponse) Reset()         { *m = EPAntInitResponse{} }
func (m *EPAntInitResponse) String() string { return proto.CompactTextString(m) }
func (*EPAntInitResponse) ProtoMessage()    {}
func (*EPAntInitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e0e94b3ebffec, []int{1}
}

func (m *EPAntInitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EPAntInitResponse.Unmarshal(m, b)
}
func (m *EPAntInitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EPAntInitResponse.Marshal(b, m, deterministic)
}
func (m *EPAntInitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EPAntInitResponse.Merge(m, src)
}
func (m *EPAntInitResponse) XXX_Size() int {
	return xxx_messageInfo_EPAntInitResponse.Size(m)
}
func (m *EPAntInitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EPAntInitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EPAntInitResponse proto.InternalMessageInfo

//天线状态
type EPAntStateRequest struct {
	Antkey               string   `protobuf:"bytes,1,opt,name=antkey,proto3" json:"antkey,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=State,proto3" json:"State,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EPAntStateRequest) Reset()         { *m = EPAntStateRequest{} }
func (m *EPAntStateRequest) String() string { return proto.CompactTextString(m) }
func (*EPAntStateRequest) ProtoMessage()    {}
func (*EPAntStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e0e94b3ebffec, []int{2}
}

func (m *EPAntStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EPAntStateRequest.Unmarshal(m, b)
}
func (m *EPAntStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EPAntStateRequest.Marshal(b, m, deterministic)
}
func (m *EPAntStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EPAntStateRequest.Merge(m, src)
}
func (m *EPAntStateRequest) XXX_Size() int {
	return xxx_messageInfo_EPAntStateRequest.Size(m)
}
func (m *EPAntStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EPAntStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EPAntStateRequest proto.InternalMessageInfo

func (m *EPAntStateRequest) GetAntkey() string {
	if m != nil {
		return m.Antkey
	}
	return ""
}

func (m *EPAntStateRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type EPAntStateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EPAntStateResponse) Reset()         { *m = EPAntStateResponse{} }
func (m *EPAntStateResponse) String() string { return proto.CompactTextString(m) }
func (*EPAntStateResponse) ProtoMessage()    {}
func (*EPAntStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e0e94b3ebffec, []int{3}
}

func (m *EPAntStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EPAntStateResponse.Unmarshal(m, b)
}
func (m *EPAntStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EPAntStateResponse.Marshal(b, m, deterministic)
}
func (m *EPAntStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EPAntStateResponse.Merge(m, src)
}
func (m *EPAntStateResponse) XXX_Size() int {
	return xxx_messageInfo_EPAntStateResponse.Size(m)
}
func (m *EPAntStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EPAntStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EPAntStateResponse proto.InternalMessageInfo

//电子车牌信息
type EPAntReadRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EPAntReadRequest) Reset()         { *m = EPAntReadRequest{} }
func (m *EPAntReadRequest) String() string { return proto.CompactTextString(m) }
func (*EPAntReadRequest) ProtoMessage()    {}
func (*EPAntReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e0e94b3ebffec, []int{4}
}

func (m *EPAntReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EPAntReadRequest.Unmarshal(m, b)
}
func (m *EPAntReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EPAntReadRequest.Marshal(b, m, deterministic)
}
func (m *EPAntReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EPAntReadRequest.Merge(m, src)
}
func (m *EPAntReadRequest) XXX_Size() int {
	return xxx_messageInfo_EPAntReadRequest.Size(m)
}
func (m *EPAntReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EPAntReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EPAntReadRequest proto.InternalMessageInfo

func (m *EPAntReadRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EPAntReadResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EPAntReadResponse) Reset()         { *m = EPAntReadResponse{} }
func (m *EPAntReadResponse) String() string { return proto.CompactTextString(m) }
func (*EPAntReadResponse) ProtoMessage()    {}
func (*EPAntReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_be7e0e94b3ebffec, []int{5}
}

func (m *EPAntReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EPAntReadResponse.Unmarshal(m, b)
}
func (m *EPAntReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EPAntReadResponse.Marshal(b, m, deterministic)
}
func (m *EPAntReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EPAntReadResponse.Merge(m, src)
}
func (m *EPAntReadResponse) XXX_Size() int {
	return xxx_messageInfo_EPAntReadResponse.Size(m)
}
func (m *EPAntReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EPAntReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EPAntReadResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EPAntInitRequest)(nil), "pb.EPAntInitRequest")
	proto.RegisterType((*EPAntInitResponse)(nil), "pb.EPAntInitResponse")
	proto.RegisterType((*EPAntStateRequest)(nil), "pb.EPAntStateRequest")
	proto.RegisterType((*EPAntStateResponse)(nil), "pb.EPAntStateResponse")
	proto.RegisterType((*EPAntReadRequest)(nil), "pb.EPAntReadRequest")
	proto.RegisterType((*EPAntReadResponse)(nil), "pb.EPAntReadResponse")
}

func init() { proto.RegisterFile("Ant.proto", fileDescriptor_be7e0e94b3ebffec) }

var fileDescriptor_be7e0e94b3ebffec = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x74, 0xcc, 0x2b, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe2, 0x12, 0x70, 0x0d, 0x70,
	0xcc, 0x2b, 0xf1, 0xcc, 0xcb, 0x2c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3,
	0x62, 0x4b, 0xcc, 0x2b, 0xc9, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2,
	0x94, 0x84, 0xb9, 0x04, 0x91, 0xd4, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x39, 0x42, 0x05,
	0x83, 0x4b, 0x12, 0x4b, 0x52, 0x09, 0x98, 0x20, 0x24, 0xc2, 0xc5, 0x0a, 0x56, 0x27, 0xc1, 0x04,
	0x16, 0x86, 0x70, 0x94, 0x44, 0xb8, 0x84, 0x90, 0x8d, 0x80, 0x1a, 0xac, 0x02, 0x75, 0x59, 0x50,
	0x6a, 0x62, 0x0a, 0xcc, 0x5c, 0x01, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa8, 0xa1, 0x20, 0x26, 0xdc,
	0x4d, 0x10, 0x55, 0x10, 0xad, 0x49, 0x6c, 0x60, 0xff, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x67, 0x56, 0x54, 0xce, 0xec, 0x00, 0x00, 0x00,
}

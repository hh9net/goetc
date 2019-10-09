// Code generated by protoc-gen-go. DO NOT EDIT.
// source: IO.proto

package pb

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//状态
type IOStateReport struct {
	State1               string   `protobuf:"bytes,1,opt,name=State1,proto3" json:"State1,omitempty"`
	State2               string   `protobuf:"bytes,2,opt,name=State2,proto3" json:"State2,omitempty"`
	State3               string   `protobuf:"bytes,3,opt,name=State3,proto3" json:"State3,omitempty"`
	State4               string   `protobuf:"bytes,4,opt,name=State4,proto3" json:"State4,omitempty"`
	State5               string   `protobuf:"bytes,5,opt,name=State5,proto3" json:"State5,omitempty"`
	State6               string   `protobuf:"bytes,6,opt,name=State6,proto3" json:"State6,omitempty"`
	State7               string   `protobuf:"bytes,7,opt,name=State7,proto3" json:"State7,omitempty"`
	State8               string   `protobuf:"bytes,8,opt,name=State8,proto3" json:"State8,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IOStateReport) Reset()         { *m = IOStateReport{} }
func (m *IOStateReport) String() string { return proto.CompactTextString(m) }
func (*IOStateReport) ProtoMessage()    {}
func (*IOStateReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_64294c404d2d4f4d, []int{0}
}

func (m *IOStateReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IOStateReport.Unmarshal(m, b)
}
func (m *IOStateReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IOStateReport.Marshal(b, m, deterministic)
}
func (m *IOStateReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IOStateReport.Merge(m, src)
}
func (m *IOStateReport) XXX_Size() int {
	return xxx_messageInfo_IOStateReport.Size(m)
}
func (m *IOStateReport) XXX_DiscardUnknown() {
	xxx_messageInfo_IOStateReport.DiscardUnknown(m)
}

var xxx_messageInfo_IOStateReport proto.InternalMessageInfo

func (m *IOStateReport) GetState1() string {
	if m != nil {
		return m.State1
	}
	return ""
}

func (m *IOStateReport) GetState2() string {
	if m != nil {
		return m.State2
	}
	return ""
}

func (m *IOStateReport) GetState3() string {
	if m != nil {
		return m.State3
	}
	return ""
}

func (m *IOStateReport) GetState4() string {
	if m != nil {
		return m.State4
	}
	return ""
}

func (m *IOStateReport) GetState5() string {
	if m != nil {
		return m.State5
	}
	return ""
}

func (m *IOStateReport) GetState6() string {
	if m != nil {
		return m.State6
	}
	return ""
}

func (m *IOStateReport) GetState7() string {
	if m != nil {
		return m.State7
	}
	return ""
}

func (m *IOStateReport) GetState8() string {
	if m != nil {
		return m.State8
	}
	return ""
}

//IO控制指令
type IOControlRequest struct {
	IONo                 string   `protobuf:"bytes,1,opt,name=IONo,proto3" json:"IONo,omitempty"`
	IOControl            string   `protobuf:"bytes,2,opt,name=IOControl,proto3" json:"IOControl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IOControlRequest) Reset()         { *m = IOControlRequest{} }
func (m *IOControlRequest) String() string { return proto.CompactTextString(m) }
func (*IOControlRequest) ProtoMessage()    {}
func (*IOControlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64294c404d2d4f4d, []int{1}
}

func (m *IOControlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IOControlRequest.Unmarshal(m, b)
}
func (m *IOControlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IOControlRequest.Marshal(b, m, deterministic)
}
func (m *IOControlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IOControlRequest.Merge(m, src)
}
func (m *IOControlRequest) XXX_Size() int {
	return xxx_messageInfo_IOControlRequest.Size(m)
}
func (m *IOControlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IOControlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IOControlRequest proto.InternalMessageInfo

func (m *IOControlRequest) GetIONo() string {
	if m != nil {
		return m.IONo
	}
	return ""
}

func (m *IOControlRequest) GetIOControl() string {
	if m != nil {
		return m.IOControl
	}
	return ""
}

type IOControlResponse struct {
	IONo                 string   `protobuf:"bytes,1,opt,name=IONo,proto3" json:"IONo,omitempty"`
	IOControl            string   `protobuf:"bytes,2,opt,name=IOControl,proto3" json:"IOControl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IOControlResponse) Reset()         { *m = IOControlResponse{} }
func (m *IOControlResponse) String() string { return proto.CompactTextString(m) }
func (*IOControlResponse) ProtoMessage()    {}
func (*IOControlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64294c404d2d4f4d, []int{2}
}

func (m *IOControlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IOControlResponse.Unmarshal(m, b)
}
func (m *IOControlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IOControlResponse.Marshal(b, m, deterministic)
}
func (m *IOControlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IOControlResponse.Merge(m, src)
}
func (m *IOControlResponse) XXX_Size() int {
	return xxx_messageInfo_IOControlResponse.Size(m)
}
func (m *IOControlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IOControlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IOControlResponse proto.InternalMessageInfo

func (m *IOControlResponse) GetIONo() string {
	if m != nil {
		return m.IONo
	}
	return ""
}

func (m *IOControlResponse) GetIOControl() string {
	if m != nil {
		return m.IOControl
	}
	return ""
}

func init() {
	proto.RegisterType((*IOStateReport)(nil), "pb.IOStateReport")
	proto.RegisterType((*IOControlRequest)(nil), "pb.IOControlRequest")
	proto.RegisterType((*IOControlResponse)(nil), "pb.IOControlResponse")
}

func init() { proto.RegisterFile("IO.proto", fileDescriptor_64294c404d2d4f4d) }

var fileDescriptor_64294c404d2d4f4d = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf0, 0xf4, 0xd7, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x3a, 0xcf, 0xc8, 0xc5, 0xeb, 0xe9, 0x1f,
	0x5c, 0x92, 0x58, 0x92, 0x1a, 0x94, 0x5a, 0x90, 0x5f, 0x54, 0x22, 0x24, 0xc6, 0xc5, 0x06, 0xe6,
	0x1a, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x70, 0x71, 0x23, 0x09, 0x26, 0x24,
	0x71, 0x23, 0xb8, 0xb8, 0xb1, 0x04, 0x33, 0x92, 0xb8, 0x31, 0x5c, 0xdc, 0x44, 0x82, 0x05, 0x49,
	0xdc, 0x04, 0x2e, 0x6e, 0x2a, 0xc1, 0x8a, 0x24, 0x6e, 0x0a, 0x17, 0x37, 0x93, 0x60, 0x43, 0x12,
	0x37, 0x83, 0x8b, 0x9b, 0x4b, 0xb0, 0x23, 0x89, 0x9b, 0xc3, 0xc5, 0x2d, 0x24, 0x38, 0x90, 0xc4,
	0x2d, 0x94, 0x5c, 0xb8, 0x04, 0x3c, 0xfd, 0x9d, 0xf3, 0xf3, 0x4a, 0x8a, 0xf2, 0x73, 0x82, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x3c, 0xfd, 0xfd, 0xf2, 0xa1, 0x3e, 0x02,
	0xb3, 0x85, 0x64, 0xb8, 0x38, 0xe1, 0xea, 0xa0, 0x5e, 0x42, 0x08, 0x28, 0xb9, 0x72, 0x09, 0x22,
	0x99, 0x52, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x4a, 0xba, 0x31, 0x49, 0x6c, 0xe0, 0x90, 0x36, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x7f, 0x31, 0x0e, 0x75, 0x01, 0x00, 0x00,
}

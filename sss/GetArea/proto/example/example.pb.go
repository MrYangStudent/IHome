// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetArea

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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	//	错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//	错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//	数据
	Data                 []*ResponseAddress `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() []*ResponseAddress {
	if m != nil {
		return m.Data
	}
	return nil
}

//	返回的数据类型
type ResponseAddress struct {
	Aid                  int32    `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Aname                string   `protobuf:"bytes,2,opt,name=aname,proto3" json:"aname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseAddress) Reset()         { *m = ResponseAddress{} }
func (m *ResponseAddress) String() string { return proto.CompactTextString(m) }
func (*ResponseAddress) ProtoMessage()    {}
func (*ResponseAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1, 0}
}

func (m *ResponseAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseAddress.Unmarshal(m, b)
}
func (m *ResponseAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseAddress.Marshal(b, m, deterministic)
}
func (m *ResponseAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseAddress.Merge(m, src)
}
func (m *ResponseAddress) XXX_Size() int {
	return xxx_messageInfo_ResponseAddress.Size(m)
}
func (m *ResponseAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseAddress proto.InternalMessageInfo

func (m *ResponseAddress) GetAid() int32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *ResponseAddress) GetAname() string {
	if m != nil {
		return m.Aname
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetArea.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetArea.Response")
	proto.RegisterType((*ResponseAddress)(nil), "go.micro.srv.GetArea.Response.address")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x85, 0xad, 0xf5, 0xff, 0x63, 0xaf, 0x8b, 0x84, 0x22, 0xa5, 0xa2, 0x94, 0x0e, 0xd2, 0x29,
	0x62, 0xdd, 0xdc, 0x1c, 0x44, 0x70, 0x0c, 0xf8, 0x00, 0x57, 0x73, 0x29, 0x05, 0xd3, 0xd4, 0x9b,
	0x28, 0xbe, 0x91, 0xaf, 0x29, 0x4d, 0xd3, 0x4d, 0x9d, 0x92, 0x73, 0x20, 0xe7, 0x3b, 0x27, 0x70,
	0x3e, 0xb3, 0x0b, 0xee, 0x9a, 0xbe, 0xd0, 0xce, 0x6f, 0xb4, 0x9d, 0x2a, 0xba, 0xb2, 0x1c, 0x9c,
	0xb2, 0xe3, 0x2b, 0x3b, 0xe5, 0xf9, 0x53, 0x3d, 0x52, 0xb8, 0x67, 0xc2, 0xb6, 0x00, 0xa1, 0xe9,
	0xfd, 0x83, 0x7c, 0x68, 0xbf, 0x33, 0x38, 0xd6, 0xe4, 0x67, 0x37, 0x79, 0x92, 0x25, 0xec, 0x88,
	0x79, 0x72, 0x55, 0xd6, 0x64, 0x5d, 0xa1, 0x57, 0x21, 0xcf, 0x60, 0x4f, 0xcc, 0xd6, 0x0f, 0xd5,
	0x61, 0xb4, 0x93, 0x92, 0x77, 0x70, 0x64, 0x30, 0x60, 0x95, 0x37, 0x79, 0x77, 0xd2, 0x5f, 0xa9,
	0xdf, 0x50, 0x6a, 0xcb, 0x56, 0x68, 0x0c, 0x93, 0xf7, 0x3a, 0xbe, 0xa9, 0x6f, 0x40, 0x24, 0x43,
	0x9e, 0x42, 0x8e, 0xa3, 0x89, 0xc8, 0x9d, 0x5e, 0xae, 0x4b, 0x0d, 0x9c, 0xd0, 0x52, 0xe2, 0xad,
	0xa2, 0x7f, 0x06, 0xf1, 0xb0, 0x6e, 0x93, 0x4f, 0x20, 0x52, 0xbe, 0xbc, 0xf8, 0x0b, 0x1b, 0xe7,
	0xd5, 0x97, 0xff, 0xb7, 0x6a, 0x0f, 0x5e, 0xf6, 0xf1, 0xa3, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xb1, 0xe7, 0xf1, 0xd5, 0x47, 0x01, 0x00, 0x00,
}
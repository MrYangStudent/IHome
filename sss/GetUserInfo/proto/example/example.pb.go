// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetUserInfo

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

// web ->srv
type Request struct {
	//	sessionid
	Sessionid            string   `protobuf:"bytes,1,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
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

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

//	srv ->web
type Response struct {
	//	错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//	错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//	"user_id": 1,
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	//  "name": "Panda",
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	//  "mobile": "110",
	Mobile string `protobuf:"bytes,5,opt,name=mobile,proto3" json:"mobile,omitempty"`
	//  "real_name": "熊猫",
	RealName string `protobuf:"bytes,6,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	//  "id_card": "210112244556677",
	IdCard string `protobuf:"bytes,7,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	//  "avatar_url":
	AvatarUrl            string   `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *Response) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Response) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

func (m *Response) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *Response) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetUserInfo.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xad, 0xee, 0xf6, 0xcf, 0x78, 0x1b, 0xc4, 0x0d, 0xae, 0x82, 0xf6, 0xa2, 0xa7, 0x08,
	0xfa, 0x11, 0x44, 0x64, 0x2f, 0x1e, 0x0a, 0xeb, 0xb5, 0x64, 0x37, 0xe3, 0x12, 0x68, 0x93, 0x3a,
	0x69, 0x17, 0x3f, 0xa7, 0x9f, 0x48, 0x9a, 0x54, 0xf4, 0xe2, 0x9e, 0x92, 0xf7, 0x9b, 0xf7, 0x02,
	0x79, 0x03, 0xcb, 0x8e, 0x5d, 0xef, 0xee, 0xe9, 0x53, 0xb5, 0x5d, 0x43, 0x3f, 0xa7, 0x0c, 0x14,
	0xc5, 0xce, 0xc9, 0xd6, 0x6c, 0xd9, 0x49, 0xcf, 0x7b, 0xf9, 0x42, 0xfd, 0xda, 0x13, 0xaf, 0xec,
	0xbb, 0x2b, 0x6f, 0x21, 0xab, 0xe8, 0x63, 0x20, 0xdf, 0xe3, 0x25, 0x14, 0x9e, 0xbc, 0x37, 0xce,
	0x1a, 0x2d, 0x92, 0xeb, 0xe4, 0xae, 0xa8, 0x7e, 0x41, 0xf9, 0x95, 0x40, 0x5e, 0x91, 0xef, 0x9c,
	0xf5, 0x84, 0x67, 0x30, 0x27, 0x66, 0xeb, 0x26, 0x5b, 0x14, 0x78, 0x0e, 0x29, 0x31, 0xb7, 0x7e,
	0x27, 0x8e, 0x03, 0x9e, 0x14, 0x2e, 0x20, 0x1b, 0x3c, 0x71, 0x6d, 0xb4, 0x38, 0x89, 0x83, 0x51,
	0xae, 0x34, 0x22, 0xcc, 0xac, 0x6a, 0x49, 0xcc, 0x02, 0x0d, 0xf7, 0xf1, 0x91, 0xd6, 0x6d, 0x4c,
	0x43, 0x62, 0x1e, 0xbd, 0x51, 0xe1, 0x12, 0x0a, 0x26, 0xd5, 0xd4, 0x21, 0x90, 0x86, 0x51, 0x3e,
	0x82, 0xd7, 0x31, 0xb4, 0x80, 0xcc, 0xe8, 0x7a, 0xab, 0x58, 0x8b, 0x2c, 0xa6, 0x8c, 0x7e, 0x52,
	0xac, 0xf1, 0x0a, 0x40, 0xed, 0x55, 0xaf, 0xb8, 0x1e, 0xb8, 0x11, 0x79, 0xfc, 0x54, 0x24, 0x6b,
	0x6e, 0x1e, 0x14, 0x64, 0xcf, 0xb1, 0x28, 0x7c, 0x83, 0xd3, 0x3f, 0xbd, 0xe0, 0x8d, 0xfc, 0xaf,
	0x32, 0x39, 0xf5, 0x75, 0x51, 0x1e, 0xb2, 0xc4, 0xa2, 0xca, 0xa3, 0x4d, 0x1a, 0x36, 0xf0, 0xf8,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xdf, 0x19, 0x83, 0xa0, 0x01, 0x00, 0x00,
}

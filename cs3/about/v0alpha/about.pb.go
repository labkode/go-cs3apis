// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/about/v0alpha/about.proto

package aboutv0alphapb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cernbox/go-cs3apis/cs3/rpc"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type ListMembersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMembersRequest) Reset()         { *m = ListMembersRequest{} }
func (m *ListMembersRequest) String() string { return proto.CompactTextString(m) }
func (*ListMembersRequest) ProtoMessage()    {}
func (*ListMembersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57540359ac0e6e7b, []int{0}
}

func (m *ListMembersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMembersRequest.Unmarshal(m, b)
}
func (m *ListMembersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMembersRequest.Marshal(b, m, deterministic)
}
func (m *ListMembersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMembersRequest.Merge(m, src)
}
func (m *ListMembersRequest) XXX_Size() int {
	return xxx_messageInfo_ListMembersRequest.Size(m)
}
func (m *ListMembersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMembersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMembersRequest proto.InternalMessageInfo

type ListMembersResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Members              []*Member   `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListMembersResponse) Reset()         { *m = ListMembersResponse{} }
func (m *ListMembersResponse) String() string { return proto.CompactTextString(m) }
func (*ListMembersResponse) ProtoMessage()    {}
func (*ListMembersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57540359ac0e6e7b, []int{1}
}

func (m *ListMembersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMembersResponse.Unmarshal(m, b)
}
func (m *ListMembersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMembersResponse.Marshal(b, m, deterministic)
}
func (m *ListMembersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMembersResponse.Merge(m, src)
}
func (m *ListMembersResponse) XXX_Size() int {
	return xxx_messageInfo_ListMembersResponse.Size(m)
}
func (m *ListMembersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMembersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMembersResponse proto.InternalMessageInfo

func (m *ListMembersResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListMembersResponse) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type GetDocumentationRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDocumentationRequest) Reset()         { *m = GetDocumentationRequest{} }
func (m *GetDocumentationRequest) String() string { return proto.CompactTextString(m) }
func (*GetDocumentationRequest) ProtoMessage()    {}
func (*GetDocumentationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57540359ac0e6e7b, []int{2}
}

func (m *GetDocumentationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentationRequest.Unmarshal(m, b)
}
func (m *GetDocumentationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentationRequest.Marshal(b, m, deterministic)
}
func (m *GetDocumentationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentationRequest.Merge(m, src)
}
func (m *GetDocumentationRequest) XXX_Size() int {
	return xxx_messageInfo_GetDocumentationRequest.Size(m)
}
func (m *GetDocumentationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentationRequest proto.InternalMessageInfo

type GetDocumentationResponse struct {
	Status               *rpc.Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Documentation        *Documentation `protobuf:"bytes,2,opt,name=documentation,proto3" json:"documentation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetDocumentationResponse) Reset()         { *m = GetDocumentationResponse{} }
func (m *GetDocumentationResponse) String() string { return proto.CompactTextString(m) }
func (*GetDocumentationResponse) ProtoMessage()    {}
func (*GetDocumentationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57540359ac0e6e7b, []int{3}
}

func (m *GetDocumentationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDocumentationResponse.Unmarshal(m, b)
}
func (m *GetDocumentationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDocumentationResponse.Marshal(b, m, deterministic)
}
func (m *GetDocumentationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDocumentationResponse.Merge(m, src)
}
func (m *GetDocumentationResponse) XXX_Size() int {
	return xxx_messageInfo_GetDocumentationResponse.Size(m)
}
func (m *GetDocumentationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDocumentationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDocumentationResponse proto.InternalMessageInfo

func (m *GetDocumentationResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetDocumentationResponse) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func init() {
	proto.RegisterType((*ListMembersRequest)(nil), "cs3.aboutv0alpha.ListMembersRequest")
	proto.RegisterType((*ListMembersResponse)(nil), "cs3.aboutv0alpha.ListMembersResponse")
	proto.RegisterType((*GetDocumentationRequest)(nil), "cs3.aboutv0alpha.GetDocumentationRequest")
	proto.RegisterType((*GetDocumentationResponse)(nil), "cs3.aboutv0alpha.GetDocumentationResponse")
}

func init() { proto.RegisterFile("cs3/about/v0alpha/about.proto", fileDescriptor_57540359ac0e6e7b) }

var fileDescriptor_57540359ac0e6e7b = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x3f, 0x4b, 0xf3, 0x40,
	0x1c, 0x26, 0x79, 0xa1, 0x2f, 0x5c, 0xdf, 0x57, 0xcb, 0x59, 0x34, 0x86, 0x5a, 0x6b, 0xa8, 0x58,
	0x1d, 0x92, 0x92, 0x7c, 0x82, 0xa6, 0x8a, 0x8b, 0x62, 0x69, 0x41, 0x44, 0x5c, 0x92, 0xf3, 0xa8,
	0x81, 0x26, 0x17, 0xef, 0x2e, 0x1d, 0x1c, 0x1d, 0x5d, 0x1c, 0x5c, 0x9d, 0x1c, 0xfd, 0x28, 0xae,
	0x7e, 0x05, 0x3f, 0x88, 0xe4, 0x2e, 0x85, 0xa4, 0x87, 0xa8, 0x63, 0x9e, 0x7f, 0x79, 0xee, 0xe1,
	0x07, 0xb6, 0x10, 0xf3, 0x9c, 0x20, 0x24, 0x19, 0x77, 0xe6, 0xfd, 0x60, 0x96, 0xde, 0x04, 0xf2,
	0xcb, 0x4e, 0x29, 0xe1, 0x04, 0x36, 0x10, 0xf3, 0x6c, 0x01, 0x14, 0xac, 0xb9, 0xa3, 0x1a, 0x28,
	0x66, 0x24, 0xa3, 0x08, 0x33, 0x69, 0x32, 0x9b, 0xb9, 0x84, 0xa6, 0xc8, 0x61, 0x3c, 0xe0, 0xd9,
	0x02, 0x6d, 0x4d, 0x09, 0x99, 0xce, 0xb0, 0x13, 0xa4, 0x91, 0x13, 0x24, 0x09, 0xe1, 0x01, 0x8f,
	0x48, 0x52, 0xb0, 0x56, 0x13, 0xc0, 0x93, 0x88, 0xf1, 0x53, 0x1c, 0x87, 0x98, 0xb2, 0x31, 0xbe,
	0xcd, 0x30, 0xe3, 0x16, 0x05, 0x6b, 0x15, 0x94, 0xa5, 0x24, 0x61, 0x18, 0xee, 0x81, 0x9a, 0x8c,
	0x36, 0xb4, 0x8e, 0xd6, 0xab, 0xbb, 0xab, 0x76, 0x5e, 0x93, 0xa6, 0xc8, 0x9e, 0x08, 0x78, 0x5c,
	0xd0, 0xd0, 0x05, 0x7f, 0x63, 0xe9, 0x35, 0xf4, 0xce, 0x9f, 0x5e, 0xdd, 0x35, 0xec, 0xe5, 0x07,
	0xd9, 0x32, 0x7c, 0xbc, 0x10, 0x5a, 0x9b, 0x60, 0xe3, 0x18, 0xf3, 0x43, 0x82, 0xb2, 0x18, 0x27,
	0xb2, 0xe4, 0xa2, 0xce, 0x83, 0x06, 0x0c, 0x95, 0xfb, 0x6d, 0xa9, 0x23, 0xf0, 0xff, 0xba, 0x9c,
	0x60, 0xe8, 0x42, 0xbf, 0xad, 0x56, 0xab, 0xfe, 0xa8, 0xea, 0x72, 0x9f, 0x75, 0xf0, 0x6f, 0x90,
	0xab, 0x27, 0x98, 0xce, 0x23, 0x84, 0xe1, 0x1d, 0xa8, 0x97, 0xc6, 0x82, 0x5d, 0x35, 0x4f, 0x5d,
	0xd8, 0xdc, 0xfd, 0x46, 0x25, 0x1f, 0x67, 0xb5, 0xef, 0xdf, 0x3f, 0x9e, 0x74, 0x03, 0xae, 0x57,
	0xaf, 0xc4, 0x29, 0x46, 0x83, 0x8f, 0x1a, 0x68, 0x2c, 0x2f, 0x03, 0xf7, 0xd5, 0xec, 0x2f, 0x96,
	0x35, 0x0f, 0x7e, 0x22, 0x2d, 0xba, 0x74, 0x45, 0x97, 0x36, 0x6c, 0x2d, 0x75, 0xa9, 0xcc, 0xe3,
	0x23, 0xd0, 0x44, 0x24, 0x56, 0x62, 0x7d, 0x20, 0x36, 0x1b, 0xe5, 0x47, 0x37, 0xd2, 0x2e, 0x57,
	0xca, 0x5c, 0x1a, 0xbe, 0xe8, 0xb5, 0xa1, 0x7f, 0x76, 0x31, 0xf0, 0x5f, 0xf5, 0xc6, 0x70, 0xe2,
	0xd9, 0x42, 0x7b, 0xde, 0x1f, 0xe4, 0xec, 0x9b, 0x80, 0xae, 0xca, 0x50, 0x58, 0x13, 0xc7, 0xeb,
	0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xae, 0xbd, 0x86, 0x9f, 0x46, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AboutServiceClient is the client API for AboutService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AboutServiceClient interface {
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
	GetDocumentation(ctx context.Context, in *GetDocumentationRequest, opts ...grpc.CallOption) (*GetDocumentationResponse, error)
}

type aboutServiceClient struct {
	cc *grpc.ClientConn
}

func NewAboutServiceClient(cc *grpc.ClientConn) AboutServiceClient {
	return &aboutServiceClient{cc}
}

func (c *aboutServiceClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := c.cc.Invoke(ctx, "/cs3.aboutv0alpha.AboutService/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aboutServiceClient) GetDocumentation(ctx context.Context, in *GetDocumentationRequest, opts ...grpc.CallOption) (*GetDocumentationResponse, error) {
	out := new(GetDocumentationResponse)
	err := c.cc.Invoke(ctx, "/cs3.aboutv0alpha.AboutService/GetDocumentation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AboutServiceServer is the server API for AboutService service.
type AboutServiceServer interface {
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	GetDocumentation(context.Context, *GetDocumentationRequest) (*GetDocumentationResponse, error)
}

func RegisterAboutServiceServer(s *grpc.Server, srv AboutServiceServer) {
	s.RegisterService(&_AboutService_serviceDesc, srv)
}

func _AboutService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AboutServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.aboutv0alpha.AboutService/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AboutServiceServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AboutService_GetDocumentation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AboutServiceServer).GetDocumentation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.aboutv0alpha.AboutService/GetDocumentation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AboutServiceServer).GetDocumentation(ctx, req.(*GetDocumentationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AboutService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.aboutv0alpha.AboutService",
	HandlerType: (*AboutServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMembers",
			Handler:    _AboutService_ListMembers_Handler,
		},
		{
			MethodName: "GetDocumentation",
			Handler:    _AboutService_GetDocumentation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/about/v0alpha/about.proto",
}

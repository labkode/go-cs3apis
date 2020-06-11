// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/app/provider/v1beta1/provider_api.proto

package providerv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// REQUIRED.
// View mode.
type OpenRequest_ViewMode int32

const (
	OpenRequest_VIEW_MODE_INVALID OpenRequest_ViewMode = 0
	// The file can be opened but not downloaded.
	OpenRequest_VIEW_MODE_VIEW_ONLY OpenRequest_ViewMode = 1
	// The file can be downloaded.
	OpenRequest_VIEW_MODE_READ_ONLY OpenRequest_ViewMode = 2
	// The file can be downloaded and updated.
	OpenRequest_VIEW_MODE_READ_WRITE OpenRequest_ViewMode = 3
)

var OpenRequest_ViewMode_name = map[int32]string{
	0: "VIEW_MODE_INVALID",
	1: "VIEW_MODE_VIEW_ONLY",
	2: "VIEW_MODE_READ_ONLY",
	3: "VIEW_MODE_READ_WRITE",
}

var OpenRequest_ViewMode_value = map[string]int32{
	"VIEW_MODE_INVALID":    0,
	"VIEW_MODE_VIEW_ONLY":  1,
	"VIEW_MODE_READ_ONLY":  2,
	"VIEW_MODE_READ_WRITE": 3,
}

func (x OpenRequest_ViewMode) String() string {
	return proto.EnumName(OpenRequest_ViewMode_name, int32(x))
}

func (OpenRequest_ViewMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0, 0}
}

type OpenRequest struct {
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The resource reference.
	Ref *v1beta11.Reference `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty"`
	// REQUIRED.
	// The access token this application provider will use when contacting
	// the storage provider to read and write.
	// Service implementors MUST make sure that the access token only grants
	// access to the requested resource.
	// Service implementors should use a ResourceId rather than a filename to grant access, as
	// ResourceIds MUST NOT change when a resource is renamed.
	// The access token MUST be short-lived.
	// TODO(labkode): investigate token derivation techniques.
	AccessToken          string               `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ViewMode             OpenRequest_ViewMode `protobuf:"varint,4,opt,name=view_mode,json=viewMode,proto3,enum=cs3.app.provider.v1beta1.OpenRequest_ViewMode" json:"view_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OpenRequest) Reset()         { *m = OpenRequest{} }
func (m *OpenRequest) String() string { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()    {}
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{0}
}

func (m *OpenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenRequest.Unmarshal(m, b)
}
func (m *OpenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenRequest.Marshal(b, m, deterministic)
}
func (m *OpenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenRequest.Merge(m, src)
}
func (m *OpenRequest) XXX_Size() int {
	return xxx_messageInfo_OpenRequest.Size(m)
}
func (m *OpenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OpenRequest proto.InternalMessageInfo

func (m *OpenRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenRequest) GetRef() *v1beta11.Reference {
	if m != nil {
		return m.Ref
	}
	return nil
}

func (m *OpenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *OpenRequest) GetViewMode() OpenRequest_ViewMode {
	if m != nil {
		return m.ViewMode
	}
	return OpenRequest_VIEW_MODE_INVALID
}

type OpenResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The url that user agents will render to clients.
	// Usually the rendering happens by using HTML iframes,
	// at least, Office 365, Collabora, OnlyOffice do like that.
	IframeUrl            string   `protobuf:"bytes,3,opt,name=iframe_url,json=iframeUrl,proto3" json:"iframe_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenResponse) Reset()         { *m = OpenResponse{} }
func (m *OpenResponse) String() string { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()    {}
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c007b70b037097fe, []int{1}
}

func (m *OpenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenResponse.Unmarshal(m, b)
}
func (m *OpenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenResponse.Marshal(b, m, deterministic)
}
func (m *OpenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenResponse.Merge(m, src)
}
func (m *OpenResponse) XXX_Size() int {
	return xxx_messageInfo_OpenResponse.Size(m)
}
func (m *OpenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OpenResponse proto.InternalMessageInfo

func (m *OpenResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *OpenResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *OpenResponse) GetIframeUrl() string {
	if m != nil {
		return m.IframeUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("cs3.app.provider.v1beta1.OpenRequest_ViewMode", OpenRequest_ViewMode_name, OpenRequest_ViewMode_value)
	proto.RegisterType((*OpenRequest)(nil), "cs3.app.provider.v1beta1.OpenRequest")
	proto.RegisterType((*OpenResponse)(nil), "cs3.app.provider.v1beta1.OpenResponse")
}

func init() {
	proto.RegisterFile("cs3/app/provider/v1beta1/provider_api.proto", fileDescriptor_c007b70b037097fe)
}

var fileDescriptor_c007b70b037097fe = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xb1, 0x53, 0x45, 0xcd, 0xa4, 0x82, 0x30, 0x80, 0x6a, 0xa2, 0x56, 0x0a, 0x91, 0x80,
	0x48, 0xa0, 0x89, 0x92, 0xac, 0x58, 0x3a, 0x4d, 0x16, 0x11, 0x6d, 0x63, 0x99, 0x36, 0x15, 0x28,
	0x92, 0xe5, 0x3a, 0x37, 0xc8, 0xa2, 0xc9, 0xdc, 0xce, 0xd8, 0xa9, 0x58, 0xf1, 0x0c, 0xbc, 0x02,
	0x4b, 0x1e, 0x85, 0x47, 0xe0, 0x69, 0xd0, 0xfc, 0xd8, 0x14, 0xaa, 0xaa, 0xdd, 0x8d, 0xcf, 0xf9,
	0x8e, 0xe7, 0xfe, 0x0c, 0x79, 0x93, 0xc8, 0x41, 0x37, 0x46, 0xec, 0xa2, 0xe0, 0x9b, 0x74, 0x01,
	0xa2, 0xbb, 0xe9, 0x9d, 0x43, 0x16, 0xf7, 0x4a, 0x21, 0x8a, 0x31, 0x65, 0x28, 0x78, 0xc6, 0xa9,
	0x97, 0xc8, 0x01, 0x8b, 0x11, 0x59, 0xe1, 0x31, 0x0b, 0x37, 0xf7, 0xd4, 0x6f, 0x04, 0x26, 0x65,
	0x5a, 0x66, 0x71, 0x96, 0x4b, 0x93, 0x6b, 0xbe, 0x55, 0xae, 0xcc, 0xb8, 0x88, 0x3f, 0xc3, 0xcd,
	0x8b, 0x04, 0x48, 0x9e, 0x8b, 0x04, 0x0a, 0x7a, 0x5f, 0xd1, 0xd9, 0x57, 0x04, 0x59, 0x22, 0xfa,
	0xcb, 0xd8, 0xed, 0xdf, 0x2e, 0xa9, 0x4f, 0x11, 0xd6, 0x21, 0x5c, 0xe6, 0x20, 0x33, 0xda, 0x23,
	0x55, 0x8e, 0xf1, 0x65, 0x0e, 0x9e, 0xd3, 0x72, 0x3a, 0xf5, 0xfe, 0x73, 0xa6, 0xaa, 0x34, 0x09,
	0x9b, 0x67, 0x53, 0x0d, 0x84, 0x16, 0xa4, 0xef, 0x48, 0x45, 0xc0, 0xd2, 0x73, 0x35, 0xff, 0x5a,
	0xf3, 0xb6, 0xba, 0x1b, 0x9d, 0xb1, 0x10, 0x96, 0x20, 0x60, 0x9d, 0x40, 0xa8, 0x32, 0xf4, 0x05,
	0xd9, 0x89, 0x93, 0x04, 0xa4, 0x8c, 0x32, 0xfe, 0x05, 0xd6, 0x5e, 0xa5, 0xe5, 0x74, 0x6a, 0x61,
	0xdd, 0x68, 0x27, 0x4a, 0xa2, 0xef, 0x49, 0x6d, 0x93, 0xc2, 0x55, 0xb4, 0xe2, 0x0b, 0xf0, 0xb6,
	0x5a, 0x4e, 0xe7, 0x61, 0x9f, 0xb1, 0xdb, 0x26, 0xc7, 0xae, 0xb5, 0xc2, 0x66, 0x29, 0x5c, 0x1d,
	0xf1, 0x05, 0x84, 0xdb, 0x1b, 0x7b, 0x6a, 0xaf, 0xc8, 0x76, 0xa1, 0xd2, 0x67, 0xe4, 0xf1, 0x6c,
	0x32, 0x3e, 0x8b, 0x8e, 0xa6, 0xa3, 0x71, 0x34, 0x39, 0x9e, 0xf9, 0x87, 0x93, 0x51, 0xe3, 0x01,
	0xdd, 0x25, 0x4f, 0xfe, 0xca, 0xfa, 0x34, 0x3d, 0x3e, 0xfc, 0xd8, 0x70, 0xfe, 0x35, 0xc2, 0xb1,
	0x3f, 0x32, 0x86, 0x4b, 0x3d, 0xf2, 0xf4, 0x3f, 0xe3, 0x2c, 0x9c, 0x9c, 0x8c, 0x1b, 0x95, 0xf6,
	0x77, 0x87, 0xec, 0x98, 0x8a, 0x24, 0xf2, 0xb5, 0x04, 0xda, 0x25, 0x55, 0xb3, 0x4a, 0x3b, 0xdd,
	0x5d, 0xdd, 0x89, 0xc0, 0xa4, 0x6c, 0xe0, 0x83, 0xb6, 0x43, 0x8b, 0x5d, 0x5b, 0x87, 0x7b, 0xdf,
	0x75, 0xec, 0x13, 0x92, 0x2e, 0x45, 0xbc, 0x82, 0x28, 0x17, 0x17, 0x76, 0xa2, 0x35, 0xa3, 0x9c,
	0x8a, 0x8b, 0xfe, 0x82, 0xd4, 0x03, 0x3b, 0x35, 0x3f, 0x98, 0xd0, 0x53, 0xb2, 0xa5, 0x2a, 0xa4,
	0x2f, 0xef, 0x35, 0xd3, 0xe6, 0xab, 0xbb, 0x30, 0xd3, 0xe8, 0xf0, 0x1b, 0xd9, 0x4b, 0xf8, 0xea,
	0x56, 0x78, 0xd8, 0x28, 0x6b, 0xc0, 0x34, 0x50, 0x0f, 0x31, 0x70, 0x3e, 0x3d, 0x2a, 0x28, 0x0b,
	0xfd, 0x70, 0x2b, 0x07, 0x7e, 0xf0, 0xd3, 0xf5, 0x0e, 0xe4, 0x80, 0xf9, 0x88, 0xac, 0xc8, 0xb0,
	0x59, 0x6f, 0xa8, 0x80, 0x5f, 0xda, 0x9a, 0xfb, 0x88, 0xf3, 0xc2, 0x9a, 0x5b, 0xeb, 0xbc, 0xaa,
	0x9f, 0xf7, 0xe0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0xb5, 0x08, 0x7b, 0x92, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProviderAPIClient is the client API for ProviderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProviderAPIClient interface {
	// Returns the iframe url
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
}

type providerAPIClient struct {
	cc *grpc.ClientConn
}

func NewProviderAPIClient(cc *grpc.ClientConn) ProviderAPIClient {
	return &providerAPIClient{cc}
}

func (c *providerAPIClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, "/cs3.app.provider.v1beta1.ProviderAPI/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderAPIServer is the server API for ProviderAPI service.
type ProviderAPIServer interface {
	// Returns the iframe url
	// MUST return CODE_NOT_FOUND if the resource does not exist.
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
}

// UnimplementedProviderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProviderAPIServer struct {
}

func (*UnimplementedProviderAPIServer) Open(ctx context.Context, req *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}

func RegisterProviderAPIServer(s *grpc.Server, srv ProviderAPIServer) {
	s.RegisterService(&_ProviderAPI_serviceDesc, srv)
}

func _ProviderAPI_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderAPIServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.app.provider.v1beta1.ProviderAPI/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderAPIServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProviderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.app.provider.v1beta1.ProviderAPI",
	HandlerType: (*ProviderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Open",
			Handler:    _ProviderAPI_Open_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/app/provider/v1beta1/provider_api.proto",
}

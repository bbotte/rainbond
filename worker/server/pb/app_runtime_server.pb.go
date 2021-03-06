// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app_runtime_server.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type StatusRequest struct {
	ServiceIds           string   `protobuf:"bytes,1,opt,name=service_ids,json=serviceIds,proto3" json:"service_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetServiceIds() string {
	if m != nil {
		return m.ServiceIds
	}
	return ""
}

type StatusMessage struct {
	Status               map[string]string `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusMessage) Reset()         { *m = StatusMessage{} }
func (m *StatusMessage) String() string { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()    {}
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{1}
}

func (m *StatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMessage.Unmarshal(m, b)
}
func (m *StatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMessage.Marshal(b, m, deterministic)
}
func (m *StatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMessage.Merge(m, src)
}
func (m *StatusMessage) XXX_Size() int {
	return xxx_messageInfo_StatusMessage.Size(m)
}
func (m *StatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMessage proto.InternalMessageInfo

func (m *StatusMessage) GetStatus() map[string]string {
	if m != nil {
		return m.Status
	}
	return nil
}

type DiskMessage struct {
	Disks                map[string]float64 `protobuf:"bytes,1,rep,name=disks,proto3" json:"disks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DiskMessage) Reset()         { *m = DiskMessage{} }
func (m *DiskMessage) String() string { return proto.CompactTextString(m) }
func (*DiskMessage) ProtoMessage()    {}
func (*DiskMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f94cf1a886c479d6, []int{2}
}

func (m *DiskMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiskMessage.Unmarshal(m, b)
}
func (m *DiskMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiskMessage.Marshal(b, m, deterministic)
}
func (m *DiskMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiskMessage.Merge(m, src)
}
func (m *DiskMessage) XXX_Size() int {
	return xxx_messageInfo_DiskMessage.Size(m)
}
func (m *DiskMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DiskMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DiskMessage proto.InternalMessageInfo

func (m *DiskMessage) GetDisks() map[string]float64 {
	if m != nil {
		return m.Disks
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "pb.StatusRequest")
	proto.RegisterType((*StatusMessage)(nil), "pb.StatusMessage")
	proto.RegisterMapType((map[string]string)(nil), "pb.StatusMessage.StatusEntry")
	proto.RegisterType((*DiskMessage)(nil), "pb.DiskMessage")
	proto.RegisterMapType((map[string]float64)(nil), "pb.DiskMessage.DisksEntry")
}

func init() { proto.RegisterFile("app_runtime_server.proto", fileDescriptor_f94cf1a886c479d6) }

var fileDescriptor_f94cf1a886c479d6 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0x29, 0x2d, 0xfc, 0x27, 0x7f, 0x5f, 0xba, 0x78, 0x08, 0x01, 0xb1, 0xec, 0xa9,
	0xa7, 0x50, 0x22, 0x4a, 0xf5, 0x56, 0x50, 0xc4, 0x83, 0x97, 0xf4, 0x03, 0x84, 0xa4, 0x19, 0x64,
	0x89, 0xa6, 0xeb, 0xce, 0xa6, 0x10, 0xf0, 0xe2, 0x37, 0x97, 0x64, 0x37, 0x34, 0xbe, 0x80, 0xb7,
	0x99, 0x67, 0x9f, 0xe1, 0x37, 0xf3, 0x2c, 0x04, 0x99, 0x52, 0xa9, 0xae, 0x2b, 0x23, 0x5f, 0x31,
	0x25, 0xd4, 0x7b, 0xd4, 0x91, 0xd2, 0x3b, 0xb3, 0xe3, 0x9e, 0xca, 0xc5, 0x12, 0x8e, 0x36, 0x26,
	0x33, 0x35, 0x25, 0xf8, 0x56, 0x23, 0x19, 0x7e, 0x01, 0x7e, 0x6b, 0x92, 0x5b, 0x4c, 0x65, 0x41,
	0x01, 0x9b, 0xb3, 0xc5, 0xbf, 0x04, 0x9c, 0xf4, 0x58, 0x90, 0xf8, 0x60, 0xfd, 0xc8, 0x13, 0x12,
	0x65, 0xcf, 0xc8, 0xaf, 0x60, 0x4a, 0x9d, 0x10, 0xb0, 0xf9, 0x78, 0xe1, 0xc7, 0xe7, 0x91, 0xca,
	0xa3, 0x2f, 0x16, 0xd7, 0xdd, 0x57, 0x46, 0x37, 0x89, 0x33, 0x87, 0x37, 0xe0, 0x0f, 0x64, 0x7e,
	0x0a, 0xe3, 0x12, 0x1b, 0x07, 0x6c, 0x4b, 0x7e, 0x06, 0x93, 0x7d, 0xf6, 0x52, 0x63, 0xe0, 0x75,
	0x9a, 0x6d, 0x6e, 0xbd, 0x15, 0x13, 0x0d, 0xf8, 0x77, 0x92, 0xca, 0x7e, 0x81, 0x25, 0x4c, 0x0a,
	0x49, 0x65, 0xcf, 0x0f, 0x5b, 0xfe, 0xe0, 0xbd, 0xab, 0x1d, 0xdc, 0x1a, 0xc3, 0x15, 0xc0, 0x41,
	0xfc, 0x0b, 0xcd, 0x06, 0xe8, 0xf8, 0x1d, 0x8e, 0xd7, 0x4a, 0x25, 0x36, 0xcf, 0x4d, 0x53, 0x6d,
	0xf9, 0x35, 0xfc, 0x7f, 0x40, 0xb3, 0x56, 0xca, 0x5e, 0xc3, 0x67, 0x87, 0xf3, 0x5d, 0xa8, 0xe1,
	0xec, 0x47, 0x22, 0x62, 0xc4, 0x63, 0x00, 0x3b, 0xd7, 0x6e, 0xf2, 0xdb, 0xd4, 0xc9, 0xb7, 0x3b,
	0xc4, 0x28, 0x9f, 0x76, 0x3f, 0x77, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xeb, 0x58, 0x41,
	0xd5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppRuntimeSyncClient is the client API for AppRuntimeSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppRuntimeSyncClient interface {
	GetAppStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusMessage, error)
	GetAppDisk(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiskMessage, error)
}

type appRuntimeSyncClient struct {
	cc *grpc.ClientConn
}

func NewAppRuntimeSyncClient(cc *grpc.ClientConn) AppRuntimeSyncClient {
	return &appRuntimeSyncClient{cc}
}

func (c *appRuntimeSyncClient) GetAppStatus(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusMessage, error) {
	out := new(StatusMessage)
	err := c.cc.Invoke(ctx, "/pb.AppRuntimeSync/GetAppStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appRuntimeSyncClient) GetAppDisk(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiskMessage, error) {
	out := new(DiskMessage)
	err := c.cc.Invoke(ctx, "/pb.AppRuntimeSync/GetAppDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppRuntimeSyncServer is the server API for AppRuntimeSync service.
type AppRuntimeSyncServer interface {
	GetAppStatus(context.Context, *StatusRequest) (*StatusMessage, error)
	GetAppDisk(context.Context, *StatusRequest) (*DiskMessage, error)
}

func RegisterAppRuntimeSyncServer(s *grpc.Server, srv AppRuntimeSyncServer) {
	s.RegisterService(&_AppRuntimeSync_serviceDesc, srv)
}

func _AppRuntimeSync_GetAppStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRuntimeSyncServer).GetAppStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AppRuntimeSync/GetAppStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRuntimeSyncServer).GetAppStatus(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppRuntimeSync_GetAppDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppRuntimeSyncServer).GetAppDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.AppRuntimeSync/GetAppDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppRuntimeSyncServer).GetAppDisk(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppRuntimeSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AppRuntimeSync",
	HandlerType: (*AppRuntimeSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppStatus",
			Handler:    _AppRuntimeSync_GetAppStatus_Handler,
		},
		{
			MethodName: "GetAppDisk",
			Handler:    _AppRuntimeSync_GetAppDisk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app_runtime_server.proto",
}

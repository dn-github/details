// Code generated by protoc-gen-go. DO NOT EDIT.
// source: details.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Book struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_c457112169d4fb2c, []int{0}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Detail struct {
	Price                int64    `protobuf:"varint,1,opt,name=price,proto3" json:"price,omitempty"`
	Genre                string   `protobuf:"bytes,2,opt,name=genre,proto3" json:"genre,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Detail) Reset()         { *m = Detail{} }
func (m *Detail) String() string { return proto.CompactTextString(m) }
func (*Detail) ProtoMessage()    {}
func (*Detail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c457112169d4fb2c, []int{1}
}

func (m *Detail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Detail.Unmarshal(m, b)
}
func (m *Detail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Detail.Marshal(b, m, deterministic)
}
func (m *Detail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Detail.Merge(m, src)
}
func (m *Detail) XXX_Size() int {
	return xxx_messageInfo_Detail.Size(m)
}
func (m *Detail) XXX_DiscardUnknown() {
	xxx_messageInfo_Detail.DiscardUnknown(m)
}

var xxx_messageInfo_Detail proto.InternalMessageInfo

func (m *Detail) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Detail) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func init() {
	proto.RegisterType((*Book)(nil), "pb.Book")
	proto.RegisterType((*Detail)(nil), "pb.Detail")
}

func init() { proto.RegisterFile("details.proto", fileDescriptor_c457112169d4fb2c) }

var fileDescriptor_c457112169d4fb2c = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x49, 0x2d, 0x49,
	0xcc, 0xcc, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe2,
	0x62, 0x71, 0xca, 0xcf, 0xcf, 0x16, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x4c, 0xb8, 0xd8, 0x5c, 0xc0, 0x1a, 0x84, 0x44, 0xb8,
	0x58, 0x0b, 0x8a, 0x32, 0x93, 0x21, 0xd2, 0xcc, 0x41, 0x10, 0x0e, 0x48, 0x34, 0x3d, 0x35, 0xaf,
	0x28, 0x55, 0x82, 0x09, 0xac, 0x09, 0xc2, 0x31, 0x32, 0xe2, 0xe2, 0x85, 0xe8, 0x0a, 0x4e, 0x2d,
	0x2a, 0x03, 0x29, 0x53, 0xe4, 0x62, 0x87, 0xda, 0x2b, 0xc4, 0xa1, 0x57, 0x90, 0xa4, 0x07, 0xb2,
	0x4f, 0x8a, 0x0b, 0xc4, 0x82, 0xa8, 0x53, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xc8, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x32, 0x10, 0xc7, 0xa1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DetailServiceClient is the client API for DetailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetailServiceClient interface {
	Details(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Detail, error)
}

type detailServiceClient struct {
	cc *grpc.ClientConn
}

func NewDetailServiceClient(cc *grpc.ClientConn) DetailServiceClient {
	return &detailServiceClient{cc}
}

func (c *detailServiceClient) Details(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Detail, error) {
	out := new(Detail)
	err := c.cc.Invoke(ctx, "/pb.DetailService/details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetailServiceServer is the server API for DetailService service.
type DetailServiceServer interface {
	Details(context.Context, *Book) (*Detail, error)
}

// UnimplementedDetailServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDetailServiceServer struct {
}

func (*UnimplementedDetailServiceServer) Details(ctx context.Context, req *Book) (*Detail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}

func RegisterDetailServiceServer(s *grpc.Server, srv DetailServiceServer) {
	s.RegisterService(&_DetailService_serviceDesc, srv)
}

func _DetailService_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetailServiceServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DetailService/Details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetailServiceServer).Details(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetailService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DetailService",
	HandlerType: (*DetailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "details",
			Handler:    _DetailService_Details_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "details.proto",
}




package testdata

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3

type EchoRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return m.Size()
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return m.Size()
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SayHelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}
func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return m.Size()
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SayHelloResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}
func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(m, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return m.Size()
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type TestAnyRequest struct {
	AnyAnimal *types.Any `protobuf:"bytes,1,opt,name=any_animal,json=anyAnimal,proto3" json:"any_animal,omitempty"`
}

func (m *TestAnyRequest) Reset()         { *m = TestAnyRequest{} }
func (m *TestAnyRequest) String() string { return proto.CompactTextString(m) }
func (*TestAnyRequest) ProtoMessage()    {}
func (*TestAnyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}
func (m *TestAnyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestAnyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestAnyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestAnyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAnyRequest.Merge(m, src)
}
func (m *TestAnyRequest) XXX_Size() int {
	return m.Size()
}
func (m *TestAnyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAnyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestAnyRequest proto.InternalMessageInfo

func (m *TestAnyRequest) GetAnyAnimal() *types.Any {
	if m != nil {
		return m.AnyAnimal
	}
	return nil
}

type TestAnyResponse struct {
	HasAnimal *HasAnimal `protobuf:"bytes,1,opt,name=has_animal,json=hasAnimal,proto3" json:"has_animal,omitempty"`
}

func (m *TestAnyResponse) Reset()         { *m = TestAnyResponse{} }
func (m *TestAnyResponse) String() string { return proto.CompactTextString(m) }
func (*TestAnyResponse) ProtoMessage()    {}
func (*TestAnyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{5}
}
func (m *TestAnyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestAnyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestAnyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestAnyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestAnyResponse.Merge(m, src)
}
func (m *TestAnyResponse) XXX_Size() int {
	return m.Size()
}
func (m *TestAnyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestAnyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestAnyResponse proto.InternalMessageInfo

func (m *TestAnyResponse) GetHasAnimal() *HasAnimal {
	if m != nil {
		return m.HasAnimal
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "testdata.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "testdata.EchoResponse")
	proto.RegisterType((*SayHelloRequest)(nil), "testdata.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "testdata.SayHelloResponse")
	proto.RegisterType((*TestAnyRequest)(nil), "testdata.TestAnyRequest")
	proto.RegisterType((*TestAnyResponse)(nil), "testdata.TestAnyResponse")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x59, 0x82, 0x02, 0x0f, 0x03, 0xa6, 0xfe, 0x08, 0xf4, 0xb0, 0x98, 0x25, 0x46, 0x2e,
	0x76, 0x09, 0xc4, 0xab, 0x09, 0x26, 0x24, 0x5c, 0x45, 0x4f, 0x5e, 0x4c, 0x81, 0xba, 0x2d, 0x6e,
	0x2d, 0xd0, 0xee, 0xb0, 0xff, 0xc2, 0x7f, 0xc9, 0x9b, 0x47, 0x8e, 0x1e, 0x0d, 0xfc, 0x23, 0x66,
	0x5b, 0xbb, 0x09, 0x21, 0x9e, 0xda, 0xd7, 0x7e, 0xde, 0xe7, 0xe5, 0x7d, 0xa1, 0xb9, 0x8c, 0xd9,
	0x2a, 0x21, 0x8b, 0x95, 0x50, 0x02, 0xd5, 0x15, 0x93, 0x6a, 0x4e, 0x15, 0xc5, 0x5d, 0x4f, 0x08,
	0x2f, 0x64, 0x6e, 0xf6, 0x3e, 0x8d, 0xdf, 0x5c, 0xca, 0x35, 0x84, 0x5b, 0x06, 0xca, 0x6b, 0xe7,
	0x06, 0x9a, 0xa3, 0x99, 0x2f, 0x26, 0x6c, 0x19, 0x33, 0xa9, 0x50, 0x07, 0x6a, 0x11, 0x93, 0x92,
	0x7a, 0xac, 0x63, 0x5d, 0x59, 0xbd, 0xc6, 0xc4, 0x94, 0x4e, 0x0f, 0x4e, 0x72, 0x50, 0x2e, 0x04,
	0x97, 0xec, 0x1f, 0xf2, 0x1a, 0xda, 0x4f, 0x34, 0x19, 0xb3, 0x30, 0x2c, 0xb4, 0x08, 0xaa, 0x9c,
	0x46, 0x86, 0xcc, 0xee, 0x0e, 0x81, 0xd3, 0x12, 0xd3, 0x52, 0x0c, 0x75, 0x6f, 0xc5, 0x98, 0x0a,
	0xb8, 0xa7, 0xd9, 0xa2, 0x76, 0x46, 0xd0, 0x7a, 0x66, 0x52, 0x0d, 0x79, 0x62, 0xac, 0x03, 0x00,
	0xca, 0x93, 0x57, 0xca, 0x83, 0x88, 0x86, 0x19, 0xdf, 0xec, 0x9f, 0x93, 0x7c, 0x77, 0x62, 0x76,
	0x27, 0x69, 0x43, 0x83, 0xf2, 0x64, 0x98, 0x61, 0xce, 0x08, 0xda, 0x85, 0x46, 0x4f, 0xed, 0x03,
	0xf8, 0x54, 0xee, 0x7a, 0xce, 0x48, 0x11, 0xd4, 0x98, 0xca, 0xbc, 0x77, 0xd2, 0xf0, 0xcd, 0xb5,
	0xff, 0x69, 0xc1, 0xd1, 0x63, 0x1a, 0x3e, 0xba, 0x83, 0x6a, 0x1a, 0x0c, 0xba, 0x28, 0x3b, 0xfe,
	0x24, 0x8a, 0x2f, 0xf7, 0x9f, 0xf5, 0xd0, 0x21, 0xd4, 0xcd, 0xfa, 0xa8, 0x5b, 0x32, 0x7b, 0xc9,
	0x61, 0x7c, 0xe8, 0x4b, 0x2b, 0xee, 0xa1, 0xa6, 0x57, 0x41, 0x9d, 0x12, 0xdb, 0x0d, 0x09, 0x77,
	0x0f, 0xfc, 0xe4, 0xfd, 0x0f, 0xe3, 0xaf, 0x8d, 0x6d, 0xad, 0x37, 0xb6, 0xf5, 0xb3, 0xb1, 0xad,
	0x8f, 0xad, 0x5d, 0x59, 0x6f, 0xed, 0xca, 0xf7, 0xd6, 0xae, 0xbc, 0x10, 0x2f, 0x50, 0x7e, 0x3c,
	0x25, 0x33, 0x11, 0xb9, 0x33, 0x21, 0x23, 0x21, 0xf5, 0x71, 0x2b, 0xe7, 0xef, 0x6e, 0x2a, 0x8c,
	0x55, 0x10, 0xba, 0xc6, 0x3c, 0x3d, 0xce, 0xd2, 0x1e, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe8,
	0xb4, 0x42, 0x4e, 0x90, 0x02, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	TestAny(ctx context.Context, in *TestAnyRequest, opts ...grpc.CallOption) (*TestAnyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/testdata.Query/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/testdata.Query/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TestAny(ctx context.Context, in *TestAnyRequest, opts ...grpc.CallOption) (*TestAnyResponse, error) {
	out := new(TestAnyResponse)
	err := c.cc.Invoke(ctx, "/testdata.Query/TestAny", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	TestAny(context.Context, *TestAnyRequest) (*TestAnyResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedQueryServer) SayHello(ctx context.Context, req *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedQueryServer) TestAny(ctx context.Context, req *TestAnyRequest) (*TestAnyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAny not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Query/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Query/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TestAny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestAnyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TestAny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testdata.Query/TestAny",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TestAny(ctx, req.(*TestAnyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testdata.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Query_Echo_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _Query_SayHello_Handler,
		},
		{
			MethodName: "TestAny",
			Handler:    _Query_TestAny_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}

func (m *EchoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EchoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EchoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EchoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SayHelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayHelloRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayHelloRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SayHelloResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SayHelloResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SayHelloResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Greeting) > 0 {
		i -= len(m.Greeting)
		copy(dAtA[i:], m.Greeting)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Greeting)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestAnyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestAnyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestAnyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AnyAnimal != nil {
		{
			size, err := m.AnyAnimal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TestAnyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestAnyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestAnyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasAnimal != nil {
		{
			size, err := m.HasAnimal.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EchoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *EchoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *SayHelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *SayHelloResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Greeting)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *TestAnyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AnyAnimal != nil {
		l = m.AnyAnimal.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *TestAnyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HasAnimal != nil {
		l = m.HasAnimal.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EchoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EchoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EchoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EchoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SayHelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SayHelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayHelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SayHelloResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SayHelloResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SayHelloResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Greeting", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Greeting = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestAnyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestAnyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestAnyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnyAnimal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AnyAnimal == nil {
				m.AnyAnimal = &types.Any{}
			}
			if err := m.AnyAnimal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestAnyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestAnyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestAnyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasAnimal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HasAnimal == nil {
				m.HasAnimal = &HasAnimal{}
			}
			if err := m.HasAnimal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)




package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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


type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo


type QueryParamsResponse struct {

	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}



type QueryBaseFeeRequest struct {
}

func (m *QueryBaseFeeRequest) Reset()         { *m = QueryBaseFeeRequest{} }
func (m *QueryBaseFeeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBaseFeeRequest) ProtoMessage()    {}
func (*QueryBaseFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{2}
}
func (m *QueryBaseFeeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBaseFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBaseFeeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBaseFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBaseFeeRequest.Merge(m, src)
}
func (m *QueryBaseFeeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBaseFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBaseFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBaseFeeRequest proto.InternalMessageInfo


type QueryBaseFeeResponse struct {
	BaseFee *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=base_fee,json=baseFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"base_fee,omitempty"`
}

func (m *QueryBaseFeeResponse) Reset()         { *m = QueryBaseFeeResponse{} }
func (m *QueryBaseFeeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBaseFeeResponse) ProtoMessage()    {}
func (*QueryBaseFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{3}
}
func (m *QueryBaseFeeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBaseFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBaseFeeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBaseFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBaseFeeResponse.Merge(m, src)
}
func (m *QueryBaseFeeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBaseFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBaseFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBaseFeeResponse proto.InternalMessageInfo



type QueryBlockGasRequest struct {
}

func (m *QueryBlockGasRequest) Reset()         { *m = QueryBlockGasRequest{} }
func (m *QueryBlockGasRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasRequest) ProtoMessage()    {}
func (*QueryBlockGasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{4}
}
func (m *QueryBlockGasRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasRequest.Merge(m, src)
}
func (m *QueryBlockGasRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasRequest proto.InternalMessageInfo


type QueryBlockGasResponse struct {
	Gas int64 `protobuf:"varint,1,opt,name=gas,proto3" json:"gas,omitempty"`
}

func (m *QueryBlockGasResponse) Reset()         { *m = QueryBlockGasResponse{} }
func (m *QueryBlockGasResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBlockGasResponse) ProtoMessage()    {}
func (*QueryBlockGasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71a07c1ffd85fde2, []int{5}
}
func (m *QueryBlockGasResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockGasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockGasResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockGasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockGasResponse.Merge(m, src)
}
func (m *QueryBlockGasResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockGasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockGasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockGasResponse proto.InternalMessageInfo

func (m *QueryBlockGasResponse) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "ethermint.feemarket.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "ethermint.feemarket.v1.QueryParamsResponse")
	proto.RegisterType((*QueryBaseFeeRequest)(nil), "ethermint.feemarket.v1.QueryBaseFeeRequest")
	proto.RegisterType((*QueryBaseFeeResponse)(nil), "ethermint.feemarket.v1.QueryBaseFeeResponse")
	proto.RegisterType((*QueryBlockGasRequest)(nil), "ethermint.feemarket.v1.QueryBlockGasRequest")
	proto.RegisterType((*QueryBlockGasResponse)(nil), "ethermint.feemarket.v1.QueryBlockGasResponse")
}

func init() {
	proto.RegisterFile("ethermint/feemarket/v1/query.proto", fileDescriptor_71a07c1ffd85fde2)
}

var fileDescriptor_71a07c1ffd85fde2 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x8e, 0x29, 0x74, 0xc3, 0x5c, 0x90, 0xe9, 0x26, 0x14, 0x21, 0xaf, 0x04, 0xa9, 0xda, 0x06,
	0xb3, 0xb5, 0x71, 0xe5, 0x14, 0x89, 0x5f, 0x37, 0x08, 0x37, 0x24, 0x34, 0x39, 0xe5, 0x2d, 0x8d,
	0xba, 0xc4, 0x59, 0xec, 0x56, 0xf4, 0xca, 0x8d, 0x0b, 0x42, 0xf0, 0xd7, 0xf0, 0x1f, 0xf4, 0x58,
	0x89, 0x0b, 0xe2, 0x50, 0xa1, 0x96, 0x3f, 0x04, 0xc5, 0x71, 0x4a, 0x03, 0x04, 0x7a, 0x8a, 0xf5,
	0xf2, 0xbd, 0xef, 0xfb, 0xde, 0xf7, 0x6c, 0xec, 0x81, 0x1e, 0x40, 0x9e, 0xc4, 0xa9, 0xe6, 0x67,
	0x00, 0x89, 0xc8, 0x87, 0xa0, 0xf9, 0xf8, 0x98, 0x5f, 0x8c, 0x20, 0x9f, 0xb0, 0x2c, 0x97, 0x5a,
	0x92, 0xdd, 0x15, 0x86, 0xad, 0x30, 0x6c, 0x7c, 0xec, 0x76, 0x22, 0x19, 0x49, 0x03, 0xe1, 0xc5,
	0xa9, 0x44, 0xbb, 0xb7, 0x22, 0x29, 0xa3, 0x73, 0xe0, 0x22, 0x8b, 0xb9, 0x48, 0x53, 0xa9, 0x85,
	0x8e, 0x65, 0xaa, 0xec, 0xdf, 0x5e, 0x83, 0xde, 0x2f, 0x62, 0x83, 0xf3, 0x3a, 0x98, 0x3c, 0x2f,
	0x2c, 0x3c, 0x13, 0xb9, 0x48, 0x54, 0x00, 0x17, 0x23, 0x50, 0xda, 0x7b, 0x81, 0x6f, 0xd4, 0xaa,
	0x2a, 0x93, 0xa9, 0x02, 0xf2, 0x00, 0xb7, 0x33, 0x53, 0xb9, 0x89, 0xba, 0x68, 0xff, 0xda, 0x09,
	0x65, 0x7f, 0x77, 0xcc, 0xca, 0x3e, 0xff, 0xf2, 0x74, 0xbe, 0xe7, 0x04, 0xb6, 0xc7, 0xdb, 0xb1,
	0xa4, 0xbe, 0x50, 0xf0, 0x08, 0xa0, 0xd2, 0x7a, 0x85, 0x3b, 0xf5, 0xb2, 0x15, 0x7b, 0x88, 0xb7,
	0x43, 0xa1, 0xe0, 0xf4, 0x0c, 0xc0, 0xc8, 0x5d, 0xf5, 0x0f, 0xbf, 0xcd, 0xf7, 0x7a, 0x51, 0xac,
	0x07, 0xa3, 0x90, 0xf5, 0x65, 0xc2, 0xfb, 0x52, 0x25, 0x52, 0xd9, 0xcf, 0x91, 0x7a, 0x3d, 0xe4,
	0x7a, 0x92, 0x81, 0x62, 0x4f, 0x53, 0x1d, 0x6c, 0x85, 0x25, 0x9d, 0xb7, 0x5b, 0xd1, 0x9f, 0xcb,
	0xfe, 0xf0, 0xb1, 0x58, 0x8d, 0x78, 0x80, 0x77, 0x7e, 0xab, 0x5b, 0xdd, 0xeb, 0xb8, 0x15, 0x89,
	0x72, 0xc2, 0x56, 0x50, 0x1c, 0x4f, 0x3e, 0xb7, 0xf0, 0x15, 0x83, 0x25, 0xef, 0x10, 0x6e, 0x97,
	0xb3, 0x91, 0xc3, 0xa6, 0xd9, 0xff, 0x8c, 0xd3, 0xbd, 0xbb, 0x11, 0xb6, 0xd4, 0xf7, 0x7a, 0x6f,
	0xbf, 0xfc, 0xf8, 0x74, 0xa9, 0x4b, 0x28, 0x6f, 0x58, 0x61, 0x19, 0x27, 0x79, 0x8f, 0xf0, 0x96,
	0xcd, 0x8c, 0xfc, 0x5b, 0xa0, 0x1e, 0xb8, 0x7b, 0x6f, 0x33, 0xb0, 0xb5, 0xb3, 0x6f, 0xec, 0x78,
	0xa4, 0xdb, 0x64, 0xa7, 0x5a, 0x12, 0xf9, 0x88, 0xf0, 0x76, 0x95, 0x26, 0xf9, 0x8f, 0x48, 0x7d,
	0x19, 0xee, 0xd1, 0x86, 0x68, 0xeb, 0xe9, 0xc0, 0x78, 0xba, 0x43, 0x6e, 0x37, 0x7a, 0x2a, 0x3a,
	0x4e, 0x23, 0xa1, 0xfc, 0x27, 0xd3, 0x05, 0x45, 0xb3, 0x05, 0x45, 0xdf, 0x17, 0x14, 0x7d, 0x58,
	0x52, 0x67, 0xb6, 0xa4, 0xce, 0xd7, 0x25, 0x75, 0x5e, 0xb2, 0xb5, 0x9b, 0xa4, 0x07, 0x22, 0x57,
	0xb1, 0x5a, 0xa3, 0x7b, 0xb3, 0x46, 0x68, 0x6e, 0x55, 0xd8, 0x36, 0x0f, 0xe6, 0xfe, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0x45, 0x36, 0x66, 0xca, 0x03, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)

	BaseFee(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error)

	BlockGas(ctx context.Context, in *QueryBlockGasRequest, opts ...grpc.CallOption) (*QueryBlockGasResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ethermint.feemarket.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BaseFee(ctx context.Context, in *QueryBaseFeeRequest, opts ...grpc.CallOption) (*QueryBaseFeeResponse, error) {
	out := new(QueryBaseFeeResponse)
	err := c.cc.Invoke(ctx, "/ethermint.feemarket.v1.Query/BaseFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BlockGas(ctx context.Context, in *QueryBlockGasRequest, opts ...grpc.CallOption) (*QueryBlockGasResponse, error) {
	out := new(QueryBlockGasResponse)
	err := c.cc.Invoke(ctx, "/ethermint.feemarket.v1.Query/BlockGas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)

	BaseFee(context.Context, *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error)

	BlockGas(context.Context, *QueryBlockGasRequest) (*QueryBlockGasResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) BaseFee(ctx context.Context, req *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseFee not implemented")
}
func (*UnimplementedQueryServer) BlockGas(ctx context.Context, req *QueryBlockGasRequest) (*QueryBlockGasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockGas not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.feemarket.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BaseFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBaseFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.feemarket.v1.Query/BaseFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseFee(ctx, req.(*QueryBaseFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BlockGas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockGasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockGas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethermint.feemarket.v1.Query/BlockGas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockGas(ctx, req.(*QueryBlockGasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethermint.feemarket.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "BaseFee",
			Handler:    _Query_BaseFee_Handler,
		},
		{
			MethodName: "BlockGas",
			Handler:    _Query_BlockGas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ethermint/feemarket/v1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryBaseFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBaseFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBaseFeeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBaseFeeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBaseFeeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBaseFeeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseFee != nil {
		{
			size := m.BaseFee.Size()
			i -= size
			if _, err := m.BaseFee.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBlockGasRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryBlockGasResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockGasResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockGasResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Gas != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Gas))
		i--
		dAtA[i] = 0x8
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryBaseFeeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBaseFeeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseFee != nil {
		l = m.BaseFee.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBlockGasRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryBlockGasResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Gas != 0 {
		n += 1 + sovQuery(uint64(m.Gas))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryBaseFeeRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBaseFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBaseFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryBaseFeeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBaseFeeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBaseFeeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
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
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.BaseFee = &v
			if err := m.BaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryBlockGasRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBlockGasRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryBlockGasResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBlockGasResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockGasResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			m.Gas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Gas |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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

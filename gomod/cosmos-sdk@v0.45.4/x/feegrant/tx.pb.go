


package feegrant

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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



type MsgGrantAllowance struct {

	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`

	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`

	Allowance *types.Any `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *MsgGrantAllowance) Reset()         { *m = MsgGrantAllowance{} }
func (m *MsgGrantAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgGrantAllowance) ProtoMessage()    {}
func (*MsgGrantAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{0}
}
func (m *MsgGrantAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantAllowance.Merge(m, src)
}
func (m *MsgGrantAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantAllowance proto.InternalMessageInfo

func (m *MsgGrantAllowance) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgGrantAllowance) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *MsgGrantAllowance) GetAllowance() *types.Any {
	if m != nil {
		return m.Allowance
	}
	return nil
}


type MsgGrantAllowanceResponse struct {
}

func (m *MsgGrantAllowanceResponse) Reset()         { *m = MsgGrantAllowanceResponse{} }
func (m *MsgGrantAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGrantAllowanceResponse) ProtoMessage()    {}
func (*MsgGrantAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{1}
}
func (m *MsgGrantAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantAllowanceResponse.Merge(m, src)
}
func (m *MsgGrantAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantAllowanceResponse proto.InternalMessageInfo


type MsgRevokeAllowance struct {

	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`

	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
}

func (m *MsgRevokeAllowance) Reset()         { *m = MsgRevokeAllowance{} }
func (m *MsgRevokeAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeAllowance) ProtoMessage()    {}
func (*MsgRevokeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{2}
}
func (m *MsgRevokeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeAllowance.Merge(m, src)
}
func (m *MsgRevokeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeAllowance proto.InternalMessageInfo

func (m *MsgRevokeAllowance) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgRevokeAllowance) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}


type MsgRevokeAllowanceResponse struct {
}

func (m *MsgRevokeAllowanceResponse) Reset()         { *m = MsgRevokeAllowanceResponse{} }
func (m *MsgRevokeAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeAllowanceResponse) ProtoMessage()    {}
func (*MsgRevokeAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{3}
}
func (m *MsgRevokeAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeAllowanceResponse.Merge(m, src)
}
func (m *MsgRevokeAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeAllowanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGrantAllowance)(nil), "cosmos.feegrant.v1beta1.MsgGrantAllowance")
	proto.RegisterType((*MsgGrantAllowanceResponse)(nil), "cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse")
	proto.RegisterType((*MsgRevokeAllowance)(nil), "cosmos.feegrant.v1beta1.MsgRevokeAllowance")
	proto.RegisterType((*MsgRevokeAllowanceResponse)(nil), "cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse")
}

func init() { proto.RegisterFile("cosmos/feegrant/v1beta1/tx.proto", fileDescriptor_dd44ad7946dad783) }

var fileDescriptor_dd44ad7946dad783 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4b, 0x4d, 0x4d, 0x2f, 0x4a, 0xcc, 0x2b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xa8,
	0xd0, 0x83, 0xa9, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07,
	0xb1, 0x20, 0xca, 0xa5, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2,
	0x34, 0xfd, 0xc4, 0xbc, 0x4a, 0x98, 0x14, 0xc4, 0xa4, 0x78, 0x88, 0x1e, 0xa8, 0xb1, 0x60, 0x8e,
	0x52, 0x1f, 0x23, 0x97, 0xa0, 0x6f, 0x71, 0xba, 0x3b, 0xc8, 0x02, 0xc7, 0x9c, 0x9c, 0xfc, 0xf2,
	0xc4, 0xbc, 0xe4, 0x54, 0x21, 0x09, 0x2e, 0x76, 0xb0, 0x95, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x30, 0x2e, 0x42, 0x26, 0x55, 0x82, 0x09, 0x59, 0x26, 0x55, 0xc8, 0x95, 0x8b,
	0x33, 0x11, 0x66, 0x80, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x88, 0x1e, 0xc4, 0x4d, 0x7a,
	0x30, 0x37, 0xe9, 0x39, 0xe6, 0x55, 0x3a, 0x09, 0x9e, 0xda, 0xa2, 0xcb, 0xeb, 0x96, 0x9a, 0x0a,
	0xb7, 0xce, 0x33, 0x08, 0xa1, 0x53, 0x49, 0x9a, 0x4b, 0x12, 0xc3, 0x3d, 0x41, 0xa9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0x1e, 0x5c, 0x42, 0xbe, 0xc5, 0xe9, 0x41, 0xa9, 0x65, 0xf9, 0xd9,
	0xa9, 0x14, 0xb9, 0x56, 0x49, 0x86, 0x4b, 0x0a, 0xd3, 0x24, 0x98, 0x3d, 0x46, 0x6f, 0x18, 0xb9,
	0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x0a, 0xb8, 0xf8, 0xd0, 0x42, 0x46, 0x4b, 0x0f, 0x47, 0xac, 0xe8,
	0x61, 0xb8, 0x5a, 0xca, 0x88, 0x78, 0xb5, 0x30, 0x9b, 0x85, 0x8a, 0xb9, 0xf8, 0xd1, 0xbd, 0xa7,
	0x8d, 0xcf, 0x18, 0x34, 0xc5, 0x52, 0xc6, 0x24, 0x28, 0x86, 0x59, 0xea, 0xe4, 0x78, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xea, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a,
	0xc9, 0xf9, 0xb9, 0xd0, 0x74, 0x03, 0xa5, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x2b, 0xe0, 0xa9, 0x37,
	0x89, 0x0d, 0x1c, 0xc5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0x47, 0xc8, 0xf2, 0xd7,
	0x02, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {


	GrantAllowance(ctx context.Context, in *MsgGrantAllowance, opts ...grpc.CallOption) (*MsgGrantAllowanceResponse, error)


	RevokeAllowance(ctx context.Context, in *MsgRevokeAllowance, opts ...grpc.CallOption) (*MsgRevokeAllowanceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GrantAllowance(ctx context.Context, in *MsgGrantAllowance, opts ...grpc.CallOption) (*MsgGrantAllowanceResponse, error) {
	out := new(MsgGrantAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Msg/GrantAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeAllowance(ctx context.Context, in *MsgRevokeAllowance, opts ...grpc.CallOption) (*MsgRevokeAllowanceResponse, error) {
	out := new(MsgRevokeAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Msg/RevokeAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {


	GrantAllowance(context.Context, *MsgGrantAllowance) (*MsgGrantAllowanceResponse, error)


	RevokeAllowance(context.Context, *MsgRevokeAllowance) (*MsgRevokeAllowanceResponse, error)
}


type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GrantAllowance(ctx context.Context, req *MsgGrantAllowance) (*MsgGrantAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantAllowance not implemented")
}
func (*UnimplementedMsgServer) RevokeAllowance(ctx context.Context, req *MsgRevokeAllowance) (*MsgRevokeAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAllowance not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GrantAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGrantAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GrantAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Msg/GrantAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GrantAllowance(ctx, req.(*MsgGrantAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevokeAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Msg/RevokeAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeAllowance(ctx, req.(*MsgRevokeAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.feegrant.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrantAllowance",
			Handler:    _Msg_GrantAllowance_Handler,
		},
		{
			MethodName: "RevokeAllowance",
			Handler:    _Msg_RevokeAllowance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/feegrant/v1beta1/tx.proto",
}

func (m *MsgGrantAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGrantAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevokeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgGrantAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGrantAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevokeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevokeAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgGrantAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGrantAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgGrantAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgGrantAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRevokeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRevokeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRevokeAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRevokeAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)




package authz

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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


type QueryGrantsRequest struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`

	MsgTypeUrl string `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGrantsRequest) Reset()         { *m = QueryGrantsRequest{} }
func (m *QueryGrantsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGrantsRequest) ProtoMessage()    {}
func (*QueryGrantsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{0}
}
func (m *QueryGrantsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGrantsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGrantsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGrantsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGrantsRequest.Merge(m, src)
}
func (m *QueryGrantsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGrantsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGrantsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGrantsRequest proto.InternalMessageInfo

func (m *QueryGrantsRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryGrantsRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryGrantsRequest) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *QueryGrantsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryGrantsResponse struct {

	Grants []*Grant `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGrantsResponse) Reset()         { *m = QueryGrantsResponse{} }
func (m *QueryGrantsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGrantsResponse) ProtoMessage()    {}
func (*QueryGrantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{1}
}
func (m *QueryGrantsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGrantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGrantsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGrantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGrantsResponse.Merge(m, src)
}
func (m *QueryGrantsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGrantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGrantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGrantsResponse proto.InternalMessageInfo

func (m *QueryGrantsResponse) GetGrants() []*Grant {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *QueryGrantsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryGranterGrantsRequest struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGranterGrantsRequest) Reset()         { *m = QueryGranterGrantsRequest{} }
func (m *QueryGranterGrantsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGranterGrantsRequest) ProtoMessage()    {}
func (*QueryGranterGrantsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{2}
}
func (m *QueryGranterGrantsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGranterGrantsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGranterGrantsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGranterGrantsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGranterGrantsRequest.Merge(m, src)
}
func (m *QueryGranterGrantsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGranterGrantsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGranterGrantsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGranterGrantsRequest proto.InternalMessageInfo

func (m *QueryGranterGrantsRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryGranterGrantsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryGranterGrantsResponse struct {

	Grants []*GrantAuthorization `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGranterGrantsResponse) Reset()         { *m = QueryGranterGrantsResponse{} }
func (m *QueryGranterGrantsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGranterGrantsResponse) ProtoMessage()    {}
func (*QueryGranterGrantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{3}
}
func (m *QueryGranterGrantsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGranterGrantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGranterGrantsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGranterGrantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGranterGrantsResponse.Merge(m, src)
}
func (m *QueryGranterGrantsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGranterGrantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGranterGrantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGranterGrantsResponse proto.InternalMessageInfo

func (m *QueryGranterGrantsResponse) GetGrants() []*GrantAuthorization {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *QueryGranterGrantsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryGranteeGrantsRequest struct {
	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGranteeGrantsRequest) Reset()         { *m = QueryGranteeGrantsRequest{} }
func (m *QueryGranteeGrantsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGranteeGrantsRequest) ProtoMessage()    {}
func (*QueryGranteeGrantsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{4}
}
func (m *QueryGranteeGrantsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGranteeGrantsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGranteeGrantsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGranteeGrantsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGranteeGrantsRequest.Merge(m, src)
}
func (m *QueryGranteeGrantsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGranteeGrantsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGranteeGrantsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGranteeGrantsRequest proto.InternalMessageInfo

func (m *QueryGranteeGrantsRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryGranteeGrantsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryGranteeGrantsResponse struct {

	Grants []*GrantAuthorization `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGranteeGrantsResponse) Reset()         { *m = QueryGranteeGrantsResponse{} }
func (m *QueryGranteeGrantsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGranteeGrantsResponse) ProtoMessage()    {}
func (*QueryGranteeGrantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{5}
}
func (m *QueryGranteeGrantsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGranteeGrantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGranteeGrantsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGranteeGrantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGranteeGrantsResponse.Merge(m, src)
}
func (m *QueryGranteeGrantsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGranteeGrantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGranteeGrantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGranteeGrantsResponse proto.InternalMessageInfo

func (m *QueryGranteeGrantsResponse) GetGrants() []*GrantAuthorization {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *QueryGranteeGrantsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGrantsRequest)(nil), "cosmos.authz.v1beta1.QueryGrantsRequest")
	proto.RegisterType((*QueryGrantsResponse)(nil), "cosmos.authz.v1beta1.QueryGrantsResponse")
	proto.RegisterType((*QueryGranterGrantsRequest)(nil), "cosmos.authz.v1beta1.QueryGranterGrantsRequest")
	proto.RegisterType((*QueryGranterGrantsResponse)(nil), "cosmos.authz.v1beta1.QueryGranterGrantsResponse")
	proto.RegisterType((*QueryGranteeGrantsRequest)(nil), "cosmos.authz.v1beta1.QueryGranteeGrantsRequest")
	proto.RegisterType((*QueryGranteeGrantsResponse)(nil), "cosmos.authz.v1beta1.QueryGranteeGrantsResponse")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/query.proto", fileDescriptor_376d714ffdeb1545) }

var fileDescriptor_376d714ffdeb1545 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x16, 0x8a, 0xf0, 0xe0, 0x62, 0x38, 0x84, 0x30, 0x45, 0x51, 0x35, 0x41, 0x41,
	0xc2, 0xde, 0x3a, 0x89, 0x23, 0x02, 0x0e, 0xdb, 0x15, 0x2a, 0xb8, 0x70, 0x99, 0xdc, 0xf1, 0xc9,
	0x8d, 0x68, 0xe3, 0xcc, 0x76, 0x10, 0x1d, 0xea, 0x05, 0x5e, 0x00, 0x69, 0x0f, 0x81, 0xc4, 0x91,
	0xa7, 0xe0, 0x38, 0x89, 0x0b, 0x47, 0xd4, 0x22, 0xf1, 0x1a, 0xa8, 0xb6, 0x4b, 0x9b, 0x91, 0x6d,
	0x01, 0x34, 0x69, 0xa7, 0x36, 0xf1, 0xff, 0xfb, 0xbe, 0x9f, 0x7f, 0x8a, 0x8d, 0xe3, 0x5d, 0xa9,
	0x87, 0x52, 0x33, 0x9e, 0x9b, 0xfe, 0x3e, 0x7b, 0xbd, 0xd1, 0x03, 0xc3, 0x37, 0xd8, 0x5e, 0x0e,
	0x6a, 0x44, 0x33, 0x25, 0x8d, 0x24, 0xd7, 0x5d, 0x82, 0xda, 0x04, 0xf5, 0x89, 0x70, 0x55, 0x48,
	0x29, 0x06, 0xc0, 0x78, 0x96, 0x30, 0x9e, 0xa6, 0xd2, 0x70, 0x93, 0xc8, 0x54, 0xbb, 0x9a, 0xf0,
	0xae, 0xef, 0xda, 0xe3, 0x1a, 0x5c, 0xb3, 0xdf, 0xad, 0x33, 0x2e, 0x92, 0xd4, 0x86, 0x7d, 0xb6,
	0x9c, 0xc0, 0x4d, 0xb3, 0x89, 0xd6, 0x67, 0x84, 0xc9, 0xd3, 0x59, 0x93, 0x6d, 0xc5, 0x53, 0xa3,
	0xbb, 0xb0, 0x97, 0x83, 0x36, 0x24, 0xc0, 0x97, 0xc4, 0xec, 0x05, 0xa8, 0x00, 0xc5, 0xa8, 0x7d,
	0xb9, 0x3b, 0x7f, 0x5c, 0xac, 0x40, 0x50, 0x5f, 0x5e, 0x01, 0x12, 0xe3, 0x2b, 0x43, 0x2d, 0x76,
	0xcc, 0x28, 0x83, 0x9d, 0x5c, 0x0d, 0x82, 0x86, 0x5d, 0xc6, 0x43, 0x2d, 0x9e, 0x8d, 0x32, 0x78,
	0xae, 0x06, 0x64, 0x0b, 0xe3, 0x05, 0x62, 0x70, 0x21, 0x46, 0xed, 0x95, 0xce, 0x2d, 0xea, 0x1d,
	0xcc, 0xf6, 0x43, 0x9d, 0x1c, 0x0f, 0x4a, 0x9f, 0x70, 0x01, 0x9e, 0xa8, 0xbb, 0x54, 0xd9, 0x3a,
	0x40, 0xf8, 0x5a, 0x01, 0x5a, 0x67, 0x32, 0xd5, 0x40, 0x36, 0x71, 0xd3, 0xc2, 0xe8, 0x00, 0xc5,
	0x8d, 0xf6, 0x4a, 0xe7, 0x26, 0x2d, 0xf3, 0x4b, 0x6d, 0x55, 0xd7, 0x47, 0xc9, 0x76, 0x01, 0xaa,
	0x6e, 0xa1, 0x6e, 0x9f, 0x0a, 0xe5, 0x26, 0x16, 0xa8, 0xc6, 0xf8, 0xc6, 0x02, 0x0a, 0x54, 0x55,
	0xa1, 0x5b, 0x25, 0xf3, 0xff, 0x45, 0xca, 0x47, 0x84, 0xc3, 0xb2, 0xf9, 0xde, 0xcd, 0xc3, 0x23,
	0x6e, 0xda, 0x27, 0xb8, 0x79, 0x94, 0x9b, 0xbe, 0x54, 0xc9, 0xbe, 0x6d, 0x7c, 0xd6, 0xa2, 0xe0,
	0x18, 0x51, 0x50, 0x14, 0x05, 0x67, 0x25, 0x0a, 0xce, 0xad, 0xa8, 0xce, 0xcf, 0x06, 0xbe, 0x68,
	0x49, 0xc9, 0x7b, 0x84, 0x9b, 0x8e, 0x93, 0x1c, 0xc3, 0xf3, 0xe7, 0x21, 0x0e, 0xef, 0x54, 0x48,
	0xba, 0xa9, 0xad, 0xb5, 0x77, 0x5f, 0x7f, 0x1c, 0xd4, 0x23, 0xb2, 0xca, 0x4a, 0x6f, 0x0c, 0xbf,
	0xb1, 0x4f, 0x08, 0x5f, 0x2d, 0x7c, 0x5d, 0x84, 0x9d, 0x36, 0xe2, 0xc8, 0x39, 0x08, 0xd7, 0xab,
	0x17, 0x78, 0xb4, 0xfb, 0x16, 0x6d, 0x9d, 0xd0, 0x93, 0xd0, 0x98, 0x3f, 0x4d, 0xec, 0xad, 0xff,
	0x33, 0x5e, 0x82, 0x85, 0xca, 0xb0, 0xf0, 0xb7, 0xb0, 0xf0, 0x1f, 0xb0, 0x30, 0x87, 0x85, 0xf1,
	0xe3, 0x07, 0x5f, 0x26, 0x11, 0x3a, 0x9c, 0x44, 0xe8, 0xfb, 0x24, 0x42, 0x1f, 0xa6, 0x51, 0xed,
	0x70, 0x1a, 0xd5, 0xbe, 0x4d, 0xa3, 0xda, 0x8b, 0x35, 0x91, 0x98, 0x7e, 0xde, 0xa3, 0xbb, 0x72,
	0x38, 0xef, 0xe9, 0x7e, 0xee, 0xe9, 0x97, 0xaf, 0xd8, 0x1b, 0x37, 0xa0, 0xd7, 0xb4, 0xb7, 0xf9,
	0xe6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x01, 0xa4, 0x75, 0x73, 0x06, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error)

	//

	GranterGrants(ctx context.Context, in *QueryGranterGrantsRequest, opts ...grpc.CallOption) (*QueryGranterGrantsResponse, error)

	//

	GranteeGrants(ctx context.Context, in *QueryGranteeGrantsRequest, opts ...grpc.CallOption) (*QueryGranteeGrantsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error) {
	out := new(QueryGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/Grants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GranterGrants(ctx context.Context, in *QueryGranterGrantsRequest, opts ...grpc.CallOption) (*QueryGranterGrantsResponse, error) {
	out := new(QueryGranterGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/GranterGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GranteeGrants(ctx context.Context, in *QueryGranteeGrantsRequest, opts ...grpc.CallOption) (*QueryGranteeGrantsResponse, error) {
	out := new(QueryGranteeGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/GranteeGrants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error)

	//

	GranterGrants(context.Context, *QueryGranterGrantsRequest) (*QueryGranterGrantsResponse, error)

	//

	GranteeGrants(context.Context, *QueryGranteeGrantsRequest) (*QueryGranteeGrantsResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Grants(ctx context.Context, req *QueryGrantsRequest) (*QueryGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grants not implemented")
}
func (*UnimplementedQueryServer) GranterGrants(ctx context.Context, req *QueryGranterGrantsRequest) (*QueryGranterGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GranterGrants not implemented")
}
func (*UnimplementedQueryServer) GranteeGrants(ctx context.Context, req *QueryGranteeGrantsRequest) (*QueryGranteeGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GranteeGrants not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Grants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Grants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/Grants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Grants(ctx, req.(*QueryGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GranterGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGranterGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GranterGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/GranterGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GranterGrants(ctx, req.(*QueryGranterGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GranteeGrants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGranteeGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GranteeGrants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/GranteeGrants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GranteeGrants(ctx, req.(*QueryGranteeGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grants",
			Handler:    _Query_Grants_Handler,
		},
		{
			MethodName: "GranterGrants",
			Handler:    _Query_GranterGrants_Handler,
		},
		{
			MethodName: "GranteeGrants",
			Handler:    _Query_GranteeGrants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/authz/v1beta1/query.proto",
}

func (m *QueryGrantsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGrantsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGrantsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGrantsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGrantsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGrantsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grants) > 0 {
		for iNdEx := len(m.Grants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Grants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGranterGrantsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGranterGrantsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGranterGrantsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGranterGrantsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGranterGrantsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGranterGrantsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grants) > 0 {
		for iNdEx := len(m.Grants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Grants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryGranteeGrantsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGranteeGrantsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGranteeGrantsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGranteeGrantsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGranteeGrantsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGranteeGrantsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grants) > 0 {
		for iNdEx := len(m.Grants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Grants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryGrantsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGrantsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Grants) > 0 {
		for _, e := range m.Grants {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGranterGrantsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGranterGrantsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Grants) > 0 {
		for _, e := range m.Grants {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGranteeGrantsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGranteeGrantsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Grants) > 0 {
		for _, e := range m.Grants {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
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
func (m *QueryGrantsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGrantsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGrantsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
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
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGrantsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGrantsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGrantsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
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
			m.Grants = append(m.Grants, &Grant{})
			if err := m.Grants[len(m.Grants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGranterGrantsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGranterGrantsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGranterGrantsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGranterGrantsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGranterGrantsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGranterGrantsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
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
			m.Grants = append(m.Grants, &GrantAuthorization{})
			if err := m.Grants[len(m.Grants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGranteeGrantsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGranteeGrantsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGranteeGrantsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryGranteeGrantsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGranteeGrantsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGranteeGrantsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
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
			m.Grants = append(m.Grants, &GrantAuthorization{})
			if err := m.Grants[len(m.Grants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

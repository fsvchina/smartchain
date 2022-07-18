


package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
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


type QueryBalanceRequest struct {

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`

	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryBalanceRequest) Reset()         { *m = QueryBalanceRequest{} }
func (m *QueryBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceRequest) ProtoMessage()    {}
func (*QueryBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{0}
}
func (m *QueryBalanceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceRequest.Merge(m, src)
}
func (m *QueryBalanceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceRequest proto.InternalMessageInfo


type QueryBalanceResponse struct {

	Balance *types.Coin `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (m *QueryBalanceResponse) Reset()         { *m = QueryBalanceResponse{} }
func (m *QueryBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBalanceResponse) ProtoMessage()    {}
func (*QueryBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{1}
}
func (m *QueryBalanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBalanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBalanceResponse.Merge(m, src)
}
func (m *QueryBalanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBalanceResponse proto.InternalMessageInfo

func (m *QueryBalanceResponse) GetBalance() *types.Coin {
	if m != nil {
		return m.Balance
	}
	return nil
}


type QueryAllBalancesRequest struct {

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllBalancesRequest) Reset()         { *m = QueryAllBalancesRequest{} }
func (m *QueryAllBalancesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllBalancesRequest) ProtoMessage()    {}
func (*QueryAllBalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{2}
}
func (m *QueryAllBalancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllBalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllBalancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllBalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllBalancesRequest.Merge(m, src)
}
func (m *QueryAllBalancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllBalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllBalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllBalancesRequest proto.InternalMessageInfo



type QueryAllBalancesResponse struct {

	Balances github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=balances,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balances"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllBalancesResponse) Reset()         { *m = QueryAllBalancesResponse{} }
func (m *QueryAllBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllBalancesResponse) ProtoMessage()    {}
func (*QueryAllBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{3}
}
func (m *QueryAllBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllBalancesResponse.Merge(m, src)
}
func (m *QueryAllBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllBalancesResponse proto.InternalMessageInfo

func (m *QueryAllBalancesResponse) GetBalances() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *QueryAllBalancesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}



type QuerySpendableBalancesRequest struct {

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySpendableBalancesRequest) Reset()         { *m = QuerySpendableBalancesRequest{} }
func (m *QuerySpendableBalancesRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySpendableBalancesRequest) ProtoMessage()    {}
func (*QuerySpendableBalancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{4}
}
func (m *QuerySpendableBalancesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpendableBalancesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpendableBalancesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpendableBalancesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpendableBalancesRequest.Merge(m, src)
}
func (m *QuerySpendableBalancesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpendableBalancesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpendableBalancesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpendableBalancesRequest proto.InternalMessageInfo



type QuerySpendableBalancesResponse struct {

	Balances github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=balances,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balances"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySpendableBalancesResponse) Reset()         { *m = QuerySpendableBalancesResponse{} }
func (m *QuerySpendableBalancesResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySpendableBalancesResponse) ProtoMessage()    {}
func (*QuerySpendableBalancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{5}
}
func (m *QuerySpendableBalancesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySpendableBalancesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySpendableBalancesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySpendableBalancesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySpendableBalancesResponse.Merge(m, src)
}
func (m *QuerySpendableBalancesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySpendableBalancesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySpendableBalancesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySpendableBalancesResponse proto.InternalMessageInfo

func (m *QuerySpendableBalancesResponse) GetBalances() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *QuerySpendableBalancesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}



type QueryTotalSupplyRequest struct {

	//

	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryTotalSupplyRequest) Reset()         { *m = QueryTotalSupplyRequest{} }
func (m *QueryTotalSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyRequest) ProtoMessage()    {}
func (*QueryTotalSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{6}
}
func (m *QueryTotalSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyRequest.Merge(m, src)
}
func (m *QueryTotalSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyRequest proto.InternalMessageInfo



type QueryTotalSupplyResponse struct {

	Supply github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=supply,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supply"`

	//

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryTotalSupplyResponse) Reset()         { *m = QueryTotalSupplyResponse{} }
func (m *QueryTotalSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyResponse) ProtoMessage()    {}
func (*QueryTotalSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{7}
}
func (m *QueryTotalSupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyResponse.Merge(m, src)
}
func (m *QueryTotalSupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyResponse proto.InternalMessageInfo

func (m *QueryTotalSupplyResponse) GetSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Supply
	}
	return nil
}

func (m *QueryTotalSupplyResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QuerySupplyOfRequest struct {

	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QuerySupplyOfRequest) Reset()         { *m = QuerySupplyOfRequest{} }
func (m *QuerySupplyOfRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySupplyOfRequest) ProtoMessage()    {}
func (*QuerySupplyOfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{8}
}
func (m *QuerySupplyOfRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySupplyOfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySupplyOfRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySupplyOfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySupplyOfRequest.Merge(m, src)
}
func (m *QuerySupplyOfRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySupplyOfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySupplyOfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySupplyOfRequest proto.InternalMessageInfo

func (m *QuerySupplyOfRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}


type QuerySupplyOfResponse struct {

	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount"`
}

func (m *QuerySupplyOfResponse) Reset()         { *m = QuerySupplyOfResponse{} }
func (m *QuerySupplyOfResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySupplyOfResponse) ProtoMessage()    {}
func (*QuerySupplyOfResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{9}
}
func (m *QuerySupplyOfResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySupplyOfResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySupplyOfResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySupplyOfResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySupplyOfResponse.Merge(m, src)
}
func (m *QuerySupplyOfResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySupplyOfResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySupplyOfResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySupplyOfResponse proto.InternalMessageInfo

func (m *QuerySupplyOfResponse) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}


type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{10}
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
	return fileDescriptor_9c6fc1939682df13, []int{11}
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


type QueryDenomsMetadataRequest struct {

	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDenomsMetadataRequest) Reset()         { *m = QueryDenomsMetadataRequest{} }
func (m *QueryDenomsMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsMetadataRequest) ProtoMessage()    {}
func (*QueryDenomsMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{12}
}
func (m *QueryDenomsMetadataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsMetadataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsMetadataRequest.Merge(m, src)
}
func (m *QueryDenomsMetadataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsMetadataRequest proto.InternalMessageInfo

func (m *QueryDenomsMetadataRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}



type QueryDenomsMetadataResponse struct {

	Metadatas []Metadata `protobuf:"bytes,1,rep,name=metadatas,proto3" json:"metadatas"`

	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDenomsMetadataResponse) Reset()         { *m = QueryDenomsMetadataResponse{} }
func (m *QueryDenomsMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsMetadataResponse) ProtoMessage()    {}
func (*QueryDenomsMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{13}
}
func (m *QueryDenomsMetadataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsMetadataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsMetadataResponse.Merge(m, src)
}
func (m *QueryDenomsMetadataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsMetadataResponse proto.InternalMessageInfo

func (m *QueryDenomsMetadataResponse) GetMetadatas() []Metadata {
	if m != nil {
		return m.Metadatas
	}
	return nil
}

func (m *QueryDenomsMetadataResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type QueryDenomMetadataRequest struct {

	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryDenomMetadataRequest) Reset()         { *m = QueryDenomMetadataRequest{} }
func (m *QueryDenomMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomMetadataRequest) ProtoMessage()    {}
func (*QueryDenomMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{14}
}
func (m *QueryDenomMetadataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomMetadataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomMetadataRequest.Merge(m, src)
}
func (m *QueryDenomMetadataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomMetadataRequest proto.InternalMessageInfo

func (m *QueryDenomMetadataRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}



type QueryDenomMetadataResponse struct {

	Metadata Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata"`
}

func (m *QueryDenomMetadataResponse) Reset()         { *m = QueryDenomMetadataResponse{} }
func (m *QueryDenomMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomMetadataResponse) ProtoMessage()    {}
func (*QueryDenomMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c6fc1939682df13, []int{15}
}
func (m *QueryDenomMetadataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomMetadataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomMetadataResponse.Merge(m, src)
}
func (m *QueryDenomMetadataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomMetadataResponse proto.InternalMessageInfo

func (m *QueryDenomMetadataResponse) GetMetadata() Metadata {
	if m != nil {
		return m.Metadata
	}
	return Metadata{}
}

func init() {
	proto.RegisterType((*QueryBalanceRequest)(nil), "cosmos.bank.v1beta1.QueryBalanceRequest")
	proto.RegisterType((*QueryBalanceResponse)(nil), "cosmos.bank.v1beta1.QueryBalanceResponse")
	proto.RegisterType((*QueryAllBalancesRequest)(nil), "cosmos.bank.v1beta1.QueryAllBalancesRequest")
	proto.RegisterType((*QueryAllBalancesResponse)(nil), "cosmos.bank.v1beta1.QueryAllBalancesResponse")
	proto.RegisterType((*QuerySpendableBalancesRequest)(nil), "cosmos.bank.v1beta1.QuerySpendableBalancesRequest")
	proto.RegisterType((*QuerySpendableBalancesResponse)(nil), "cosmos.bank.v1beta1.QuerySpendableBalancesResponse")
	proto.RegisterType((*QueryTotalSupplyRequest)(nil), "cosmos.bank.v1beta1.QueryTotalSupplyRequest")
	proto.RegisterType((*QueryTotalSupplyResponse)(nil), "cosmos.bank.v1beta1.QueryTotalSupplyResponse")
	proto.RegisterType((*QuerySupplyOfRequest)(nil), "cosmos.bank.v1beta1.QuerySupplyOfRequest")
	proto.RegisterType((*QuerySupplyOfResponse)(nil), "cosmos.bank.v1beta1.QuerySupplyOfResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.bank.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.bank.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryDenomsMetadataRequest)(nil), "cosmos.bank.v1beta1.QueryDenomsMetadataRequest")
	proto.RegisterType((*QueryDenomsMetadataResponse)(nil), "cosmos.bank.v1beta1.QueryDenomsMetadataResponse")
	proto.RegisterType((*QueryDenomMetadataRequest)(nil), "cosmos.bank.v1beta1.QueryDenomMetadataRequest")
	proto.RegisterType((*QueryDenomMetadataResponse)(nil), "cosmos.bank.v1beta1.QueryDenomMetadataResponse")
}

func init() { proto.RegisterFile("cosmos/bank/v1beta1/query.proto", fileDescriptor_9c6fc1939682df13) }

var fileDescriptor_9c6fc1939682df13 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x97, 0x5d, 0x6b, 0x23, 0x55,
	0x18, 0xc7, 0x73, 0xaa, 0x4d, 0xd3, 0x27, 0x28, 0x78, 0x1a, 0x31, 0x9d, 0xda, 0x89, 0x4c, 0xb5,
	0x4d, 0x6b, 0x3a, 0xd3, 0x24, 0x82, 0xd6, 0x1b, 0x69, 0x2a, 0x7a, 0x21, 0xd2, 0x98, 0x7a, 0x25,
	0x48, 0x39, 0x49, 0xc6, 0x31, 0x34, 0x99, 0x33, 0xcd, 0x99, 0x88, 0xa1, 0x14, 0x44, 0x10, 0x04,
	0x41, 0x05, 0x6f, 0x04, 0x6f, 0xea, 0x8d, 0xa0, 0x5f, 0xc0, 0xaf, 0xd0, 0x8b, 0xbd, 0x28, 0xbb,
	0x37, 0x7b, 0xb5, 0xbb, 0xb4, 0x7b, 0xb1, 0x37, 0xfb, 0x1d, 0x96, 0x9c, 0x97, 0x74, 0x92, 0x4c,
	0x92, 0x59, 0xc8, 0xb2, 0xec, 0x55, 0x93, 0x33, 0xcf, 0xcb, 0xef, 0xf9, 0x9f, 0x33, 0xff, 0x93,
	0x42, 0xa6, 0x46, 0x59, 0x8b, 0x32, 0xab, 0x4a, 0xdc, 0x63, 0xeb, 0xbb, 0x7c, 0xd5, 0xf6, 0x49,
	0xde, 0x3a, 0xe9, 0xd8, 0xed, 0xae, 0xe9, 0xb5, 0xa9, 0x4f, 0xf1, 0x92, 0x08, 0x30, 0x7b, 0x01,
	0xa6, 0x0c, 0xd0, 0xb6, 0xfa, 0x59, 0xcc, 0x16, 0xd1, 0xfd, 0x5c, 0x8f, 0x38, 0x0d, 0x97, 0xf8,
	0x0d, 0xea, 0x8a, 0x02, 0x5a, 0xca, 0xa1, 0x0e, 0xe5, 0x1f, 0xad, 0xde, 0x27, 0xb9, 0xfa, 0xa6,
	0x43, 0xa9, 0xd3, 0xb4, 0x2d, 0xe2, 0x35, 0x2c, 0xe2, 0xba, 0xd4, 0xe7, 0x29, 0x4c, 0x3e, 0xd5,
	0x83, 0xf5, 0x55, 0xe5, 0x1a, 0x6d, 0xb8, 0x23, 0xcf, 0x03, 0xd4, 0x9c, 0x90, 0x3f, 0x37, 0x0e,
	0x60, 0xe9, 0x8b, 0x1e, 0x55, 0x89, 0x34, 0x89, 0x5b, 0xb3, 0x2b, 0xf6, 0x49, 0xc7, 0x66, 0x3e,
	0x4e, 0xc3, 0x02, 0xa9, 0xd7, 0xdb, 0x36, 0x63, 0x69, 0xf4, 0x16, 0xca, 0x2e, 0x56, 0xd4, 0x57,
	0x9c, 0x82, 0xf9, 0xba, 0xed, 0xd2, 0x56, 0x7a, 0x8e, 0xaf, 0x8b, 0x2f, 0x1f, 0x26, 0x7e, 0x3e,
	0xcf, 0xc4, 0x1e, 0x9d, 0x67, 0x62, 0xc6, 0x67, 0x90, 0x1a, 0x2c, 0xc8, 0x3c, 0xea, 0x32, 0x1b,
	0x17, 0x61, 0xa1, 0x2a, 0x96, 0x78, 0xc5, 0x64, 0x61, 0xd9, 0xec, 0xeb, 0xc5, 0x6c, 0xa5, 0x97,
	0xb9, 0x4f, 0x1b, 0x6e, 0x45, 0x45, 0x1a, 0x3f, 0x21, 0x78, 0x83, 0x57, 0xdb, 0x6b, 0x36, 0x65,
	0x41, 0x36, 0x1d, 0xf1, 0x13, 0x80, 0x1b, 0x6d, 0x39, 0x67, 0xb2, 0xb0, 0x3e, 0xd0, 0x4d, 0x6c,
	0x9b, 0xea, 0x59, 0x26, 0x8e, 0x1a, 0xbc, 0x12, 0xc8, 0x0c, 0x0c, 0x75, 0x0b, 0x41, 0x7a, 0x94,
	0x43, 0x4e, 0xe6, 0x40, 0x42, 0xf2, 0xf6, 0x48, 0x5e, 0x9a, 0x38, 0x5a, 0x69, 0xe7, 0xe2, 0x5e,
	0x26, 0xf6, 0xdf, 0xfd, 0x4c, 0xd6, 0x69, 0xf8, 0xdf, 0x76, 0xaa, 0x66, 0x8d, 0xb6, 0x2c, 0xb9,
	0x45, 0xe2, 0xcf, 0x36, 0xab, 0x1f, 0x5b, 0x7e, 0xd7, 0xb3, 0x19, 0x4f, 0x60, 0x95, 0x7e, 0x71,
	0xfc, 0x69, 0xc8, 0x5c, 0x1b, 0x53, 0xe7, 0x12, 0x94, 0xc1, 0xc1, 0x8c, 0x5f, 0x10, 0xac, 0xf2,
	0x71, 0x0e, 0x3d, 0xdb, 0xad, 0x93, 0x6a, 0xd3, 0x7e, 0x9e, 0xe2, 0xde, 0x46, 0xa0, 0x8f, 0xa3,
	0x79, 0x61, 0x25, 0x3e, 0x96, 0x07, 0xf7, 0x4b, 0xea, 0x93, 0xe6, 0x61, 0xc7, 0xf3, 0x9a, 0x5d,
	0xa5, 0xed, 0xa0, 0x82, 0x68, 0x06, 0x0a, 0x5e, 0xa8, 0xe3, 0x39, 0xd0, 0x4d, 0x6a, 0x57, 0x83,
	0x38, 0xe3, 0x2b, 0xcf, 0x42, 0x39, 0x59, 0x7a, 0x76, 0xba, 0xe5, 0xa4, 0x7d, 0x88, 0x21, 0x0e,
	0xbe, 0x51, 0xa2, 0xf5, 0x6d, 0x07, 0x05, 0x6c, 0xc7, 0x28, 0xc3, 0xeb, 0x43, 0xd1, 0x72, 0xe8,
	0xf7, 0x21, 0x4e, 0x5a, 0xb4, 0xe3, 0xfa, 0x53, 0xcd, 0xa6, 0xf4, 0x72, 0x6f, 0xe8, 0x8a, 0x0c,
	0x37, 0x52, 0x80, 0x79, 0xc5, 0x32, 0x69, 0x93, 0x96, 0x7a, 0x1d, 0x8c, 0xb2, 0x74, 0x49, 0xb5,
	0x2a, 0xbb, 0xec, 0x42, 0xdc, 0xe3, 0x2b, 0xb2, 0xcb, 0x8a, 0x19, 0x72, 0x05, 0x98, 0x22, 0x49,
	0xf5, 0x11, 0x09, 0x46, 0x1d, 0x34, 0x5e, 0xf1, 0xe3, 0xde, 0x1c, 0xec, 0x73, 0xdb, 0x27, 0x75,
	0xe2, 0x93, 0x19, 0x1f, 0x11, 0xe3, 0x5f, 0x04, 0x2b, 0xa1, 0x6d, 0xe4, 0x00, 0x7b, 0xb0, 0xd8,
	0x92, 0x6b, 0xea, 0xc5, 0x5a, 0x0d, 0x9d, 0x41, 0x65, 0xca, 0x29, 0x6e, 0xb2, 0x66, 0xb7, 0xf3,
	0x79, 0x58, 0xbe, 0x41, 0x1d, 0x16, 0x24, 0x7c, 0xfb, 0xbf, 0x0e, 0x8a, 0x38, 0x32, 0xdc, 0x47,
	0x90, 0x50, 0x98, 0x52, 0xc2, 0x48, 0xb3, 0xf5, 0x93, 0x0a, 0x8f, 0x17, 0x61, 0x9e, 0xd7, 0xc7,
	0x7f, 0x22, 0x58, 0x90, 0xa6, 0x84, 0xb3, 0xa1, 0x45, 0x42, 0x2e, 0x51, 0x6d, 0x33, 0x42, 0xa4,
	0x60, 0x35, 0x3e, 0xf8, 0xf1, 0xce, 0xc3, 0x3f, 0xe6, 0x0a, 0x78, 0xc7, 0x0a, 0xbf, 0xaf, 0x85,
	0x3d, 0x59, 0xa7, 0xd2, 0x85, 0xcf, 0xac, 0x6a, 0xf7, 0x88, 0x6b, 0x80, 0xff, 0x42, 0x90, 0x0c,
	0xdc, 0x4a, 0x38, 0x37, 0xbe, 0xe9, 0xe8, 0x25, 0xaa, 0x6d, 0x47, 0x8c, 0x96, 0x98, 0x16, 0xc7,
	0xdc, 0xc4, 0x1b, 0x11, 0x31, 0xf1, 0xff, 0x08, 0x5e, 0x1b, 0xb1, 0x75, 0x5c, 0x18, 0xdf, 0x75,
	0xdc, 0x8d, 0xa4, 0x15, 0x9f, 0x2a, 0x47, 0xf2, 0xee, 0x72, 0xde, 0x22, 0xce, 0x87, 0xf2, 0x32,
	0x95, 0x77, 0x14, 0x42, 0xfe, 0x1b, 0x82, 0x64, 0xc0, 0x4e, 0x27, 0xe9, 0x3a, 0xea, 0xf1, 0x93,
	0x74, 0x0d, 0xf1, 0x68, 0x63, 0x8d, 0x73, 0xae, 0xe2, 0x95, 0x70, 0x4e, 0x41, 0xf0, 0x2b, 0x82,
	0x84, 0x32, 0x3a, 0x3c, 0xe1, 0x6c, 0x0d, 0x59, 0xa7, 0xb6, 0x15, 0x25, 0x54, 0x82, 0xbc, 0xcb,
	0x41, 0xde, 0xc1, 0x6b, 0x13, 0x40, 0xac, 0x53, 0x7e, 0xf2, 0xce, 0xf0, 0x0f, 0x08, 0xe2, 0xc2,
	0xdc, 0xf0, 0xc6, 0xf8, 0x1e, 0x03, 0x4e, 0xaa, 0x65, 0xa7, 0x07, 0x46, 0xd2, 0x44, 0xd8, 0x28,
	0xfe, 0x07, 0xc1, 0x2b, 0x03, 0x6f, 0x3f, 0x36, 0xc7, 0x37, 0x08, 0x73, 0x16, 0xcd, 0x8a, 0x1c,
	0x2f, 0xb9, 0xde, 0xe3, 0x5c, 0x26, 0xce, 0x85, 0x72, 0x71, 0x69, 0xd8, 0x91, 0xf2, 0x90, 0xbe,
	0x56, 0x7f, 0x23, 0x78, 0x75, 0xd0, 0x84, 0xf1, 0xb4, 0xce, 0xc3, 0xb7, 0x82, 0xb6, 0x13, 0x3d,
	0x41, 0xb2, 0xe6, 0x38, 0xeb, 0x3a, 0x7e, 0x3b, 0x0a, 0x6b, 0x69, 0xff, 0xe2, 0x4a, 0x47, 0x97,
	0x57, 0x3a, 0x7a, 0x70, 0xa5, 0xa3, 0xdf, 0xaf, 0xf5, 0xd8, 0xe5, 0xb5, 0x1e, 0xbb, 0x7b, 0xad,
	0xc7, 0xbe, 0xda, 0x9c, 0xf8, 0x83, 0xe0, 0x7b, 0x51, 0x96, 0xff, 0x2e, 0xa8, 0xc6, 0xf9, 0xff,
	0x15, 0xc5, 0x27, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xfd, 0x88, 0xa1, 0x2f, 0x0d, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type QueryClient interface {

	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)

	AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error)


	SpendableBalances(ctx context.Context, in *QuerySpendableBalancesRequest, opts ...grpc.CallOption) (*QuerySpendableBalancesResponse, error)

	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error)

	SupplyOf(ctx context.Context, in *QuerySupplyOfRequest, opts ...grpc.CallOption) (*QuerySupplyOfResponse, error)

	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)

	DenomMetadata(ctx context.Context, in *QueryDenomMetadataRequest, opts ...grpc.CallOption) (*QueryDenomMetadataResponse, error)

	DenomsMetadata(ctx context.Context, in *QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*QueryDenomsMetadataResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllBalances(ctx context.Context, in *QueryAllBalancesRequest, opts ...grpc.CallOption) (*QueryAllBalancesResponse, error) {
	out := new(QueryAllBalancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/AllBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SpendableBalances(ctx context.Context, in *QuerySpendableBalancesRequest, opts ...grpc.CallOption) (*QuerySpendableBalancesResponse, error) {
	out := new(QuerySpendableBalancesResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/SpendableBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error) {
	out := new(QueryTotalSupplyResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SupplyOf(ctx context.Context, in *QuerySupplyOfRequest, opts ...grpc.CallOption) (*QuerySupplyOfResponse, error) {
	out := new(QuerySupplyOfResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/SupplyOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomMetadata(ctx context.Context, in *QueryDenomMetadataRequest, opts ...grpc.CallOption) (*QueryDenomMetadataResponse, error) {
	out := new(QueryDenomMetadataResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/DenomMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomsMetadata(ctx context.Context, in *QueryDenomsMetadataRequest, opts ...grpc.CallOption) (*QueryDenomsMetadataResponse, error) {
	out := new(QueryDenomsMetadataResponse)
	err := c.cc.Invoke(ctx, "/cosmos.bank.v1beta1.Query/DenomsMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type QueryServer interface {

	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)

	AllBalances(context.Context, *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error)


	SpendableBalances(context.Context, *QuerySpendableBalancesRequest) (*QuerySpendableBalancesResponse, error)

	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error)

	SupplyOf(context.Context, *QuerySupplyOfRequest) (*QuerySupplyOfResponse, error)

	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)

	DenomMetadata(context.Context, *QueryDenomMetadataRequest) (*QueryDenomMetadataResponse, error)

	DenomsMetadata(context.Context, *QueryDenomsMetadataRequest) (*QueryDenomsMetadataResponse, error)
}


type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Balance(ctx context.Context, req *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (*UnimplementedQueryServer) AllBalances(ctx context.Context, req *QueryAllBalancesRequest) (*QueryAllBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllBalances not implemented")
}
func (*UnimplementedQueryServer) SpendableBalances(ctx context.Context, req *QuerySpendableBalancesRequest) (*QuerySpendableBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpendableBalances not implemented")
}
func (*UnimplementedQueryServer) TotalSupply(ctx context.Context, req *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (*UnimplementedQueryServer) SupplyOf(ctx context.Context, req *QuerySupplyOfRequest) (*QuerySupplyOfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplyOf not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DenomMetadata(ctx context.Context, req *QueryDenomMetadataRequest) (*QueryDenomMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomMetadata not implemented")
}
func (*UnimplementedQueryServer) DenomsMetadata(ctx context.Context, req *QueryDenomsMetadataRequest) (*QueryDenomsMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomsMetadata not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/AllBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllBalances(ctx, req.(*QueryAllBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SpendableBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpendableBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SpendableBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/SpendableBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SpendableBalances(ctx, req.(*QuerySpendableBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SupplyOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyOfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SupplyOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/SupplyOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SupplyOf(ctx, req.(*QuerySupplyOfRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/cosmos.bank.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/DenomMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomMetadata(ctx, req.(*QueryDenomMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.bank.v1beta1.Query/DenomsMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomsMetadata(ctx, req.(*QueryDenomsMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.bank.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "AllBalances",
			Handler:    _Query_AllBalances_Handler,
		},
		{
			MethodName: "SpendableBalances",
			Handler:    _Query_SpendableBalances_Handler,
		},
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "SupplyOf",
			Handler:    _Query_SupplyOf_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomMetadata",
			Handler:    _Query_DenomMetadata_Handler,
		},
		{
			MethodName: "DenomsMetadata",
			Handler:    _Query_DenomsMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/bank/v1beta1/query.proto",
}

func (m *QueryBalanceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBalanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBalanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBalanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Balance != nil {
		{
			size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllBalancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllBalancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllBalancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QuerySpendableBalancesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpendableBalancesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpendableBalancesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySpendableBalancesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySpendableBalancesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySpendableBalancesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryTotalSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTotalSupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Supply) > 0 {
		for iNdEx := len(m.Supply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QuerySupplyOfRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySupplyOfRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySupplyOfRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySupplyOfResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySupplyOfResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySupplyOfResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDenomsMetadataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsMetadataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsMetadataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomsMetadataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsMetadataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsMetadataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Metadatas) > 0 {
		for iNdEx := len(m.Metadatas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metadatas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDenomMetadataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomMetadataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomMetadataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomMetadataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomMetadataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomMetadataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryBalanceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBalanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Balance != nil {
		l = m.Balance.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllBalancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
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

func (m *QuerySpendableBalancesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySpendableBalancesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
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

func (m *QueryTotalSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryTotalSupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Supply) > 0 {
		for _, e := range m.Supply {
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

func (m *QuerySupplyOfRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySupplyOfResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
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

func (m *QueryDenomsMetadataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomsMetadataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Metadatas) > 0 {
		for _, e := range m.Metadatas {
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

func (m *QueryDenomMetadataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomMetadataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Metadata.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryBalanceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBalanceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryBalanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryBalanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBalanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
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
			if m.Balance == nil {
				m.Balance = &types.Coin{}
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllBalancesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllBalancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllBalancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QueryAllBalancesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
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
			m.Balances = append(m.Balances, types.Coin{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySpendableBalancesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySpendableBalancesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpendableBalancesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *QuerySpendableBalancesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySpendableBalancesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySpendableBalancesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
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
			m.Balances = append(m.Balances, types.Coin{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryTotalSupplyRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryTotalSupplyResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryTotalSupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
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
			m.Supply = append(m.Supply, types.Coin{})
			if err := m.Supply[len(m.Supply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QuerySupplyOfRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySupplyOfRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySupplyOfRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QuerySupplyOfResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QuerySupplyOfResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySupplyOfResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDenomsMetadataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomsMetadataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsMetadataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryDenomsMetadataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomsMetadataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadatas", wireType)
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
			m.Metadatas = append(m.Metadatas, Metadata{})
			if err := m.Metadatas[len(m.Metadatas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDenomMetadataRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomMetadataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomMetadataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
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
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDenomMetadataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomMetadataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

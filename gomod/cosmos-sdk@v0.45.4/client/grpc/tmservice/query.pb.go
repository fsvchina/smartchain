


package tmservice

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	p2p "github.com/tendermint/tendermint/proto/tendermint/p2p"
	types1 "github.com/tendermint/tendermint/proto/tendermint/types"
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


type GetValidatorSetByHeightRequest struct {
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`

	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetValidatorSetByHeightRequest) Reset()         { *m = GetValidatorSetByHeightRequest{} }
func (m *GetValidatorSetByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetValidatorSetByHeightRequest) ProtoMessage()    {}
func (*GetValidatorSetByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{0}
}
func (m *GetValidatorSetByHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetValidatorSetByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetValidatorSetByHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetValidatorSetByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValidatorSetByHeightRequest.Merge(m, src)
}
func (m *GetValidatorSetByHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetValidatorSetByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValidatorSetByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetValidatorSetByHeightRequest proto.InternalMessageInfo

func (m *GetValidatorSetByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetValidatorSetByHeightRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type GetValidatorSetByHeightResponse struct {
	BlockHeight int64        `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validators  []*Validator `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetValidatorSetByHeightResponse) Reset()         { *m = GetValidatorSetByHeightResponse{} }
func (m *GetValidatorSetByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetValidatorSetByHeightResponse) ProtoMessage()    {}
func (*GetValidatorSetByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{1}
}
func (m *GetValidatorSetByHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetValidatorSetByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetValidatorSetByHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetValidatorSetByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetValidatorSetByHeightResponse.Merge(m, src)
}
func (m *GetValidatorSetByHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetValidatorSetByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetValidatorSetByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetValidatorSetByHeightResponse proto.InternalMessageInfo

func (m *GetValidatorSetByHeightResponse) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *GetValidatorSetByHeightResponse) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *GetValidatorSetByHeightResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type GetLatestValidatorSetRequest struct {

	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetLatestValidatorSetRequest) Reset()         { *m = GetLatestValidatorSetRequest{} }
func (m *GetLatestValidatorSetRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestValidatorSetRequest) ProtoMessage()    {}
func (*GetLatestValidatorSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{2}
}
func (m *GetLatestValidatorSetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestValidatorSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestValidatorSetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestValidatorSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestValidatorSetRequest.Merge(m, src)
}
func (m *GetLatestValidatorSetRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestValidatorSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestValidatorSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestValidatorSetRequest proto.InternalMessageInfo

func (m *GetLatestValidatorSetRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type GetLatestValidatorSetResponse struct {
	BlockHeight int64        `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Validators  []*Validator `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators,omitempty"`

	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetLatestValidatorSetResponse) Reset()         { *m = GetLatestValidatorSetResponse{} }
func (m *GetLatestValidatorSetResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestValidatorSetResponse) ProtoMessage()    {}
func (*GetLatestValidatorSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{3}
}
func (m *GetLatestValidatorSetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestValidatorSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestValidatorSetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestValidatorSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestValidatorSetResponse.Merge(m, src)
}
func (m *GetLatestValidatorSetResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestValidatorSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestValidatorSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestValidatorSetResponse proto.InternalMessageInfo

func (m *GetLatestValidatorSetResponse) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *GetLatestValidatorSetResponse) GetValidators() []*Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *GetLatestValidatorSetResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}


type Validator struct {
	Address          string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey           *types.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VotingPower      int64      `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
	ProposerPriority int64      `protobuf:"varint,4,opt,name=proposer_priority,json=proposerPriority,proto3" json:"proposer_priority,omitempty"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{4}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetPubKey() *types.Any {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Validator) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func (m *Validator) GetProposerPriority() int64 {
	if m != nil {
		return m.ProposerPriority
	}
	return 0
}


type GetBlockByHeightRequest struct {
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *GetBlockByHeightRequest) Reset()         { *m = GetBlockByHeightRequest{} }
func (m *GetBlockByHeightRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHeightRequest) ProtoMessage()    {}
func (*GetBlockByHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{5}
}
func (m *GetBlockByHeightRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockByHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockByHeightRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockByHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHeightRequest.Merge(m, src)
}
func (m *GetBlockByHeightRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockByHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHeightRequest proto.InternalMessageInfo

func (m *GetBlockByHeightRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}


type GetBlockByHeightResponse struct {
	BlockId *types1.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types1.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *GetBlockByHeightResponse) Reset()         { *m = GetBlockByHeightResponse{} }
func (m *GetBlockByHeightResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockByHeightResponse) ProtoMessage()    {}
func (*GetBlockByHeightResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{6}
}
func (m *GetBlockByHeightResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockByHeightResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockByHeightResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockByHeightResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockByHeightResponse.Merge(m, src)
}
func (m *GetBlockByHeightResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockByHeightResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockByHeightResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockByHeightResponse proto.InternalMessageInfo

func (m *GetBlockByHeightResponse) GetBlockId() *types1.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetBlockByHeightResponse) GetBlock() *types1.Block {
	if m != nil {
		return m.Block
	}
	return nil
}


type GetLatestBlockRequest struct {
}

func (m *GetLatestBlockRequest) Reset()         { *m = GetLatestBlockRequest{} }
func (m *GetLatestBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetLatestBlockRequest) ProtoMessage()    {}
func (*GetLatestBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{7}
}
func (m *GetLatestBlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestBlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestBlockRequest.Merge(m, src)
}
func (m *GetLatestBlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestBlockRequest proto.InternalMessageInfo


type GetLatestBlockResponse struct {
	BlockId *types1.BlockID `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types1.Block   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *GetLatestBlockResponse) Reset()         { *m = GetLatestBlockResponse{} }
func (m *GetLatestBlockResponse) String() string { return proto.CompactTextString(m) }
func (*GetLatestBlockResponse) ProtoMessage()    {}
func (*GetLatestBlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{8}
}
func (m *GetLatestBlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLatestBlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLatestBlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLatestBlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLatestBlockResponse.Merge(m, src)
}
func (m *GetLatestBlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLatestBlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLatestBlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLatestBlockResponse proto.InternalMessageInfo

func (m *GetLatestBlockResponse) GetBlockId() *types1.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetLatestBlockResponse) GetBlock() *types1.Block {
	if m != nil {
		return m.Block
	}
	return nil
}


type GetSyncingRequest struct {
}

func (m *GetSyncingRequest) Reset()         { *m = GetSyncingRequest{} }
func (m *GetSyncingRequest) String() string { return proto.CompactTextString(m) }
func (*GetSyncingRequest) ProtoMessage()    {}
func (*GetSyncingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{9}
}
func (m *GetSyncingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSyncingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSyncingRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSyncingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSyncingRequest.Merge(m, src)
}
func (m *GetSyncingRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetSyncingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSyncingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSyncingRequest proto.InternalMessageInfo


type GetSyncingResponse struct {
	Syncing bool `protobuf:"varint,1,opt,name=syncing,proto3" json:"syncing,omitempty"`
}

func (m *GetSyncingResponse) Reset()         { *m = GetSyncingResponse{} }
func (m *GetSyncingResponse) String() string { return proto.CompactTextString(m) }
func (*GetSyncingResponse) ProtoMessage()    {}
func (*GetSyncingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{10}
}
func (m *GetSyncingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSyncingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSyncingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSyncingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSyncingResponse.Merge(m, src)
}
func (m *GetSyncingResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetSyncingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSyncingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSyncingResponse proto.InternalMessageInfo

func (m *GetSyncingResponse) GetSyncing() bool {
	if m != nil {
		return m.Syncing
	}
	return false
}


type GetNodeInfoRequest struct {
}

func (m *GetNodeInfoRequest) Reset()         { *m = GetNodeInfoRequest{} }
func (m *GetNodeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeInfoRequest) ProtoMessage()    {}
func (*GetNodeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{11}
}
func (m *GetNodeInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNodeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNodeInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNodeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeInfoRequest.Merge(m, src)
}
func (m *GetNodeInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetNodeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeInfoRequest proto.InternalMessageInfo


type GetNodeInfoResponse struct {
	DefaultNodeInfo    *p2p.DefaultNodeInfo `protobuf:"bytes,1,opt,name=default_node_info,json=defaultNodeInfo,proto3" json:"default_node_info,omitempty"`
	ApplicationVersion *VersionInfo         `protobuf:"bytes,2,opt,name=application_version,json=applicationVersion,proto3" json:"application_version,omitempty"`
}

func (m *GetNodeInfoResponse) Reset()         { *m = GetNodeInfoResponse{} }
func (m *GetNodeInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetNodeInfoResponse) ProtoMessage()    {}
func (*GetNodeInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{12}
}
func (m *GetNodeInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetNodeInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetNodeInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetNodeInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeInfoResponse.Merge(m, src)
}
func (m *GetNodeInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetNodeInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeInfoResponse proto.InternalMessageInfo

func (m *GetNodeInfoResponse) GetDefaultNodeInfo() *p2p.DefaultNodeInfo {
	if m != nil {
		return m.DefaultNodeInfo
	}
	return nil
}

func (m *GetNodeInfoResponse) GetApplicationVersion() *VersionInfo {
	if m != nil {
		return m.ApplicationVersion
	}
	return nil
}


type VersionInfo struct {
	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AppName   string    `protobuf:"bytes,2,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Version   string    `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	GitCommit string    `protobuf:"bytes,4,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	BuildTags string    `protobuf:"bytes,5,opt,name=build_tags,json=buildTags,proto3" json:"build_tags,omitempty"`
	GoVersion string    `protobuf:"bytes,6,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
	BuildDeps []*Module `protobuf:"bytes,7,rep,name=build_deps,json=buildDeps,proto3" json:"build_deps,omitempty"`

	CosmosSdkVersion string `protobuf:"bytes,8,opt,name=cosmos_sdk_version,json=cosmosSdkVersion,proto3" json:"cosmos_sdk_version,omitempty"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{13}
}
func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(m, src)
}
func (m *VersionInfo) XXX_Size() int {
	return m.Size()
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

func (m *VersionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionInfo) GetAppName() string {
	if m != nil {
		return m.AppName
	}
	return ""
}

func (m *VersionInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionInfo) GetGitCommit() string {
	if m != nil {
		return m.GitCommit
	}
	return ""
}

func (m *VersionInfo) GetBuildTags() string {
	if m != nil {
		return m.BuildTags
	}
	return ""
}

func (m *VersionInfo) GetGoVersion() string {
	if m != nil {
		return m.GoVersion
	}
	return ""
}

func (m *VersionInfo) GetBuildDeps() []*Module {
	if m != nil {
		return m.BuildDeps
	}
	return nil
}

func (m *VersionInfo) GetCosmosSdkVersion() string {
	if m != nil {
		return m.CosmosSdkVersion
	}
	return ""
}


type Module struct {

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`

	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`

	Sum string `protobuf:"bytes,3,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c93fb3ef485c5d, []int{14}
}
func (m *Module) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Module.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(m, src)
}
func (m *Module) XXX_Size() int {
	return m.Size()
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Module) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Module) GetSum() string {
	if m != nil {
		return m.Sum
	}
	return ""
}

func init() {
	proto.RegisterType((*GetValidatorSetByHeightRequest)(nil), "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest")
	proto.RegisterType((*GetValidatorSetByHeightResponse)(nil), "cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse")
	proto.RegisterType((*GetLatestValidatorSetRequest)(nil), "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest")
	proto.RegisterType((*GetLatestValidatorSetResponse)(nil), "cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse")
	proto.RegisterType((*Validator)(nil), "cosmos.base.tendermint.v1beta1.Validator")
	proto.RegisterType((*GetBlockByHeightRequest)(nil), "cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest")
	proto.RegisterType((*GetBlockByHeightResponse)(nil), "cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse")
	proto.RegisterType((*GetLatestBlockRequest)(nil), "cosmos.base.tendermint.v1beta1.GetLatestBlockRequest")
	proto.RegisterType((*GetLatestBlockResponse)(nil), "cosmos.base.tendermint.v1beta1.GetLatestBlockResponse")
	proto.RegisterType((*GetSyncingRequest)(nil), "cosmos.base.tendermint.v1beta1.GetSyncingRequest")
	proto.RegisterType((*GetSyncingResponse)(nil), "cosmos.base.tendermint.v1beta1.GetSyncingResponse")
	proto.RegisterType((*GetNodeInfoRequest)(nil), "cosmos.base.tendermint.v1beta1.GetNodeInfoRequest")
	proto.RegisterType((*GetNodeInfoResponse)(nil), "cosmos.base.tendermint.v1beta1.GetNodeInfoResponse")
	proto.RegisterType((*VersionInfo)(nil), "cosmos.base.tendermint.v1beta1.VersionInfo")
	proto.RegisterType((*Module)(nil), "cosmos.base.tendermint.v1beta1.Module")
}

func init() {
	proto.RegisterFile("cosmos/base/tendermint/v1beta1/query.proto", fileDescriptor_40c93fb3ef485c5d)
}

var fileDescriptor_40c93fb3ef485c5d = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x4d, 0x6f, 0xdc, 0x44,
	0x18, 0x8e, 0x77, 0xdb, 0x6c, 0xf2, 0x2e, 0x82, 0x64, 0x12, 0x1a, 0xc7, 0x4a, 0xb7, 0x61, 0x0f,
	0x6d, 0x42, 0x88, 0xad, 0xdd, 0xb6, 0x69, 0x0f, 0xa5, 0x88, 0x10, 0x48, 0xa3, 0x96, 0x2a, 0x72,
	0x10, 0x07, 0x84, 0x64, 0x79, 0xd7, 0x13, 0x67, 0x94, 0x5d, 0xcf, 0xd4, 0x33, 0x0e, 0x5a, 0xa1,
	0x0a, 0xc4, 0x89, 0x23, 0x12, 0x7f, 0xa1, 0x07, 0xf8, 0x03, 0x1c, 0x11, 0x47, 0x8e, 0x15, 0x48,
	0xa8, 0xe2, 0x84, 0x12, 0x7e, 0x08, 0xf2, 0xcc, 0x78, 0xd7, 0x6e, 0x92, 0xee, 0x6e, 0x0e, 0x48,
	0x3d, 0x79, 0xe6, 0xfd, 0x9a, 0xe7, 0x79, 0x66, 0xde, 0xf1, 0xc0, 0xbb, 0x6d, 0xca, 0xbb, 0x94,
	0x3b, 0x2d, 0x9f, 0x63, 0x47, 0xe0, 0x28, 0xc0, 0x71, 0x97, 0x44, 0xc2, 0x39, 0x6a, 0xb4, 0xb0,
	0xf0, 0x1b, 0xce, 0x93, 0x04, 0xc7, 0x3d, 0x9b, 0xc5, 0x54, 0x50, 0x54, 0x53, 0xb1, 0x76, 0x1a,
	0x6b, 0x0f, 0x62, 0x6d, 0x1d, 0x6b, 0xcd, 0x87, 0x34, 0xa4, 0x32, 0xd4, 0x49, 0x47, 0x2a, 0xcb,
	0x5a, 0x0c, 0x29, 0x0d, 0x3b, 0xd8, 0x91, 0xb3, 0x56, 0xb2, 0xef, 0xf8, 0x91, 0x2e, 0x68, 0x2d,
	0x69, 0x97, 0xcf, 0x88, 0xe3, 0x47, 0x11, 0x15, 0xbe, 0x20, 0x34, 0xe2, 0xda, 0x6b, 0xe5, 0xe0,
	0xb0, 0x26, 0x73, 0x44, 0x8f, 0xe1, 0xcc, 0xb7, 0x94, 0xf3, 0x49, 0xbb, 0xd3, 0xea, 0xd0, 0xf6,
	0xe1, 0xb9, 0xde, 0x7c, 0x6e, 0x81, 0xb2, 0xe4, 0xd7, 0x67, 0xcb, 0xfc, 0x90, 0x44, 0x12, 0x84,
	0x8a, 0xad, 0x7f, 0x6b, 0x40, 0x6d, 0x1b, 0x8b, 0xcf, 0xfd, 0x0e, 0x09, 0x7c, 0x41, 0xe3, 0x3d,
	0x2c, 0x36, 0x7b, 0x0f, 0x30, 0x09, 0x0f, 0x84, 0x8b, 0x9f, 0x24, 0x98, 0x0b, 0x74, 0x05, 0x26,
	0x0f, 0xa4, 0xc1, 0x34, 0x96, 0x8d, 0x95, 0xb2, 0xab, 0x67, 0xe8, 0x13, 0x80, 0x41, 0x39, 0xb3,
	0xb4, 0x6c, 0xac, 0x54, 0x9b, 0xd7, 0xed, 0xbc, 0x84, 0x4a, 0x5b, 0xbd, 0xb6, 0xbd, 0xeb, 0x87,
	0x58, 0xd7, 0x74, 0x73, 0x99, 0xf5, 0x17, 0x06, 0x5c, 0x3b, 0x17, 0x02, 0x67, 0x34, 0xe2, 0x18,
	0xbd, 0x03, 0x6f, 0x48, 0xfe, 0x5e, 0x01, 0x49, 0x55, 0xda, 0x54, 0x28, 0xda, 0x01, 0x38, 0xca,
	0x4a, 0x70, 0xb3, 0xb4, 0x5c, 0x5e, 0xa9, 0x36, 0x57, 0xed, 0x57, 0xef, 0xa8, 0xdd, 0x5f, 0xd4,
	0xcd, 0x25, 0xa3, 0xed, 0x02, 0xb3, 0xb2, 0x64, 0x76, 0x63, 0x28, 0x33, 0x05, 0xb5, 0x40, 0x6d,
	0x1f, 0x96, 0xb6, 0xb1, 0x78, 0xe4, 0x0b, 0xcc, 0x0b, 0xfc, 0x32, 0x69, 0x8b, 0x12, 0x1a, 0x17,
	0x96, 0xf0, 0x2f, 0x03, 0xae, 0x9e, 0xb3, 0xd0, 0xeb, 0x2d, 0xe0, 0x33, 0x03, 0xa6, 0xfb, 0x4b,
	0x20, 0x13, 0x2a, 0x7e, 0x10, 0xc4, 0x98, 0x73, 0x89, 0x7f, 0xda, 0xcd, 0xa6, 0x68, 0x1d, 0x2a,
	0x2c, 0x69, 0x79, 0x87, 0xb8, 0xa7, 0x0f, 0xe2, 0xbc, 0xad, 0x5a, 0xcf, 0xce, 0xba, 0xd2, 0xfe,
	0x30, 0xea, 0xb9, 0x93, 0x2c, 0x69, 0x3d, 0xc4, 0xbd, 0x54, 0x8d, 0x23, 0x2a, 0x48, 0x14, 0x7a,
	0x8c, 0x7e, 0x85, 0x63, 0x89, 0xb0, 0xec, 0x56, 0x95, 0x6d, 0x37, 0x35, 0xa1, 0x35, 0x98, 0x65,
	0x31, 0x65, 0x94, 0xe3, 0xd8, 0x63, 0x31, 0xa1, 0x31, 0x11, 0x3d, 0xf3, 0x92, 0x8c, 0x9b, 0xc9,
	0x1c, 0xbb, 0xda, 0x5e, 0x6f, 0xc0, 0xc2, 0x36, 0x16, 0x9b, 0xa9, 0x98, 0x23, 0x76, 0x4f, 0xfd,
	0x1b, 0x30, 0x4f, 0xa7, 0xe8, 0xcd, 0xba, 0x05, 0x53, 0x6a, 0xb3, 0x48, 0xa0, 0x0f, 0xc5, 0x62,
	0x5e, 0x7b, 0xd5, 0xeb, 0x32, 0x75, 0x67, 0xcb, 0xad, 0xc8, 0xd0, 0x9d, 0x00, 0xad, 0xc3, 0x65,
	0x39, 0xd4, 0x0a, 0x2c, 0x9c, 0x93, 0xe2, 0xaa, 0xa8, 0xfa, 0x02, 0xbc, 0xdd, 0x3f, 0x32, 0xca,
	0xa1, 0x10, 0xd7, 0x9f, 0xc2, 0x95, 0x97, 0x1d, 0xff, 0x27, 0xae, 0x39, 0x98, 0xdd, 0xc6, 0x62,
	0xaf, 0x17, 0xb5, 0x49, 0x14, 0x66, 0x98, 0x6c, 0x40, 0x79, 0xa3, 0xc6, 0x63, 0x42, 0x85, 0x2b,
	0x93, 0x84, 0x33, 0xe5, 0x66, 0xd3, 0xfa, 0xbc, 0x8c, 0x7f, 0x4c, 0x03, 0xbc, 0x13, 0xed, 0xd3,
	0xac, 0xca, 0x6f, 0x06, 0xcc, 0x15, 0xcc, 0xba, 0xce, 0x43, 0x98, 0x0d, 0xf0, 0xbe, 0x9f, 0x74,
	0x84, 0x17, 0xd1, 0x00, 0x7b, 0x24, 0xda, 0xa7, 0x9a, 0xe0, 0xb5, 0x3c, 0x5a, 0xd6, 0x64, 0xf6,
	0x96, 0x0a, 0xec, 0xd7, 0x78, 0x2b, 0x28, 0x1a, 0xd0, 0x97, 0x30, 0xe7, 0x33, 0xd6, 0x21, 0x6d,
	0x79, 0x82, 0xbd, 0x23, 0x1c, 0xf3, 0xc1, 0xfd, 0xb8, 0x36, 0xb4, 0x9f, 0x54, 0xb8, 0x2c, 0x8d,
	0x72, 0x75, 0xb4, 0xbd, 0xfe, 0x53, 0x09, 0xaa, 0xb9, 0x18, 0x84, 0xe0, 0x52, 0xe4, 0x77, 0xb1,
	0xee, 0x07, 0x39, 0x46, 0x8b, 0x30, 0xe5, 0x33, 0xe6, 0x49, 0x7b, 0x49, 0xf7, 0x09, 0x63, 0x8f,
	0x53, 0x97, 0x09, 0x95, 0x0c, 0x50, 0x59, 0x79, 0xf4, 0x14, 0x5d, 0x05, 0x08, 0x89, 0xf0, 0xda,
	0xb4, 0xdb, 0x25, 0x42, 0x1e, 0xf4, 0x69, 0x77, 0x3a, 0x24, 0xe2, 0x23, 0x69, 0x48, 0xdd, 0xad,
	0x84, 0x74, 0x02, 0x4f, 0xf8, 0x21, 0x37, 0x2f, 0x2b, 0xb7, 0xb4, 0x7c, 0xe6, 0x87, 0x5c, 0x66,
	0xd3, 0x3e, 0xd7, 0x49, 0x9d, 0x4d, 0x35, 0x52, 0xf4, 0x71, 0x96, 0x1d, 0x60, 0xc6, 0xcd, 0x8a,
	0xbc, 0x5a, 0xae, 0x0f, 0x93, 0xe2, 0x53, 0x1a, 0x24, 0x1d, 0xac, 0x57, 0xd9, 0xc2, 0x8c, 0xa3,
	0xf7, 0x00, 0xa9, 0x1c, 0x8f, 0x07, 0x87, 0xfd, 0xd5, 0xa6, 0xe4, 0x6a, 0x33, 0xca, 0xb3, 0x17,
	0x1c, 0x66, 0x52, 0x3d, 0x80, 0x49, 0x55, 0x22, 0x15, 0x89, 0xf9, 0xe2, 0x20, 0x13, 0x29, 0x1d,
	0xe7, 0x95, 0x28, 0x15, 0x95, 0x98, 0x81, 0x32, 0x4f, 0xba, 0x5a, 0x9f, 0x74, 0xd8, 0xfc, 0x7e,
	0x1a, 0x2a, 0x7b, 0x38, 0x3e, 0x22, 0x6d, 0x8c, 0x7e, 0x36, 0xa0, 0x9a, 0x3b, 0x43, 0xa8, 0x39,
	0x8c, 0xc6, 0xe9, 0x73, 0x68, 0xdd, 0x1c, 0x2b, 0x47, 0x1d, 0xd2, 0x7a, 0xe3, 0xbb, 0x3f, 0xff,
	0xfd, 0xb1, 0xb4, 0x86, 0x56, 0x9d, 0x21, 0x2f, 0x9a, 0xfe, 0x11, 0x46, 0xcf, 0x0c, 0x80, 0x41,
	0xdb, 0xa0, 0xc6, 0x08, 0xcb, 0x16, 0xfb, 0xce, 0x6a, 0x8e, 0x93, 0xa2, 0x81, 0x3a, 0x12, 0xe8,
	0x2a, 0xba, 0x31, 0x0c, 0xa8, 0x6e, 0x56, 0xf4, 0x8b, 0x01, 0x6f, 0x16, 0x6f, 0x1c, 0x74, 0x7b,
	0x84, 0x75, 0x4f, 0x5f, 0x5d, 0xd6, 0xc6, 0xb8, 0x69, 0x1a, 0xf2, 0x6d, 0x09, 0xd9, 0x41, 0xeb,
	0xc3, 0x20, 0xcb, 0x2b, 0x8a, 0x3b, 0x1d, 0x59, 0x03, 0xfd, 0x6a, 0xc0, 0xcc, 0xcb, 0x97, 0x38,
	0xba, 0x33, 0x02, 0x86, 0xb3, 0xfe, 0x14, 0xd6, 0xdd, 0xf1, 0x13, 0x35, 0xfc, 0x3b, 0x12, 0x7e,
	0x03, 0x39, 0x23, 0xc2, 0xff, 0x5a, 0xfd, 0x83, 0x9e, 0xa2, 0x3f, 0x8c, 0xdc, 0x4f, 0x20, 0xff,
	0x6e, 0x40, 0xf7, 0x46, 0x56, 0xf2, 0x8c, 0x77, 0x8d, 0xf5, 0xfe, 0x05, 0xb3, 0x35, 0x9f, 0x7b,
	0x92, 0xcf, 0x06, 0xba, 0x35, 0x8c, 0xcf, 0xe0, 0xc9, 0x81, 0x45, 0x7f, 0x57, 0xfe, 0x36, 0xe4,
	0xdf, 0xf8, 0xac, 0xf7, 0x24, 0xba, 0x3f, 0x02, 0xb0, 0x57, 0xbc, 0x85, 0xad, 0x0f, 0x2e, 0x9c,
	0xaf, 0xa9, 0xdd, 0x97, 0xd4, 0xee, 0xa2, 0x8d, 0xf1, 0xa8, 0x65, 0x3b, 0xb6, 0xf9, 0xe8, 0xf7,
	0xe3, 0x9a, 0xf1, 0xfc, 0xb8, 0x66, 0xfc, 0x73, 0x5c, 0x33, 0x7e, 0x38, 0xa9, 0x4d, 0x3c, 0x3f,
	0xa9, 0x4d, 0xbc, 0x38, 0xa9, 0x4d, 0x7c, 0xd1, 0x0c, 0x89, 0x38, 0x48, 0x5a, 0x76, 0x9b, 0x76,
	0xb3, 0xda, 0xea, 0xb3, 0xce, 0x83, 0x43, 0xa7, 0xdd, 0x21, 0x38, 0x12, 0x4e, 0x18, 0xb3, 0xb6,
	0x23, 0xba, 0x5c, 0x5d, 0x66, 0xad, 0x49, 0xf9, 0x3a, 0xba, 0xf9, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0x37, 0x87, 0xe4, 0x25, 0x0d, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type ServiceClient interface {

	GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error)

	GetSyncing(ctx context.Context, in *GetSyncingRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error)

	GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error)

	GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error)

	GetLatestValidatorSet(ctx context.Context, in *GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error)

	GetValidatorSetByHeight(ctx context.Context, in *GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetNodeInfo(ctx context.Context, in *GetNodeInfoRequest, opts ...grpc.CallOption) (*GetNodeInfoResponse, error) {
	out := new(GetNodeInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetSyncing(ctx context.Context, in *GetSyncingRequest, opts ...grpc.CallOption) (*GetSyncingResponse, error) {
	out := new(GetSyncingResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetSyncing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestBlock(ctx context.Context, in *GetLatestBlockRequest, opts ...grpc.CallOption) (*GetLatestBlockResponse, error) {
	out := new(GetLatestBlockResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetBlockByHeight(ctx context.Context, in *GetBlockByHeightRequest, opts ...grpc.CallOption) (*GetBlockByHeightResponse, error) {
	out := new(GetBlockByHeightResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetLatestValidatorSet(ctx context.Context, in *GetLatestValidatorSetRequest, opts ...grpc.CallOption) (*GetLatestValidatorSetResponse, error) {
	out := new(GetLatestValidatorSetResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetValidatorSetByHeight(ctx context.Context, in *GetValidatorSetByHeightRequest, opts ...grpc.CallOption) (*GetValidatorSetByHeightResponse, error) {
	out := new(GetValidatorSetByHeightResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type ServiceServer interface {

	GetNodeInfo(context.Context, *GetNodeInfoRequest) (*GetNodeInfoResponse, error)

	GetSyncing(context.Context, *GetSyncingRequest) (*GetSyncingResponse, error)

	GetLatestBlock(context.Context, *GetLatestBlockRequest) (*GetLatestBlockResponse, error)

	GetBlockByHeight(context.Context, *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error)

	GetLatestValidatorSet(context.Context, *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error)

	GetValidatorSetByHeight(context.Context, *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error)
}


type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) GetNodeInfo(ctx context.Context, req *GetNodeInfoRequest) (*GetNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (*UnimplementedServiceServer) GetSyncing(ctx context.Context, req *GetSyncingRequest) (*GetSyncingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncing not implemented")
}
func (*UnimplementedServiceServer) GetLatestBlock(ctx context.Context, req *GetLatestBlockRequest) (*GetLatestBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (*UnimplementedServiceServer) GetBlockByHeight(ctx context.Context, req *GetBlockByHeightRequest) (*GetBlockByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHeight not implemented")
}
func (*UnimplementedServiceServer) GetLatestValidatorSet(ctx context.Context, req *GetLatestValidatorSetRequest) (*GetLatestValidatorSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValidatorSet not implemented")
}
func (*UnimplementedServiceServer) GetValidatorSetByHeight(ctx context.Context, req *GetValidatorSetByHeightRequest) (*GetValidatorSetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorSetByHeight not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetNodeInfo(ctx, req.(*GetNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetSyncing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSyncingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetSyncing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetSyncing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetSyncing(ctx, req.(*GetSyncingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestBlock(ctx, req.(*GetLatestBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetBlockByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBlockByHeight(ctx, req.(*GetBlockByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetLatestValidatorSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestValidatorSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetLatestValidatorSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetLatestValidatorSet(ctx, req.(*GetLatestValidatorSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetValidatorSetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorSetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.tendermint.v1beta1.Service/GetValidatorSetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetValidatorSetByHeight(ctx, req.(*GetValidatorSetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.tendermint.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeInfo",
			Handler:    _Service_GetNodeInfo_Handler,
		},
		{
			MethodName: "GetSyncing",
			Handler:    _Service_GetSyncing_Handler,
		},
		{
			MethodName: "GetLatestBlock",
			Handler:    _Service_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlockByHeight",
			Handler:    _Service_GetBlockByHeight_Handler,
		},
		{
			MethodName: "GetLatestValidatorSet",
			Handler:    _Service_GetLatestValidatorSet_Handler,
		},
		{
			MethodName: "GetValidatorSetByHeight",
			Handler:    _Service_GetValidatorSetByHeight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/tendermint/v1beta1/query.proto",
}

func (m *GetValidatorSetByHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetValidatorSetByHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetValidatorSetByHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetValidatorSetByHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetValidatorSetByHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetValidatorSetByHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x1a
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetLatestValidatorSetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestValidatorSetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestValidatorSetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *GetLatestValidatorSetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestValidatorSetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestValidatorSetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x1a
	}
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposerPriority != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ProposerPriority))
		i--
		dAtA[i] = 0x20
	}
	if m.VotingPower != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x18
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetBlockByHeightRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockByHeightRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockByHeightRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockByHeightResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockByHeightResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockByHeightResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetLatestBlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestBlockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestBlockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLatestBlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLatestBlockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLatestBlockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetSyncingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSyncingRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSyncingRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetSyncingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSyncingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSyncingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Syncing {
		i--
		if m.Syncing {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetNodeInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodeInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNodeInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetNodeInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetNodeInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetNodeInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ApplicationVersion != nil {
		{
			size, err := m.ApplicationVersion.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.DefaultNodeInfo != nil {
		{
			size, err := m.DefaultNodeInfo.MarshalToSizedBuffer(dAtA[:i])
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

func (m *VersionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CosmosSdkVersion) > 0 {
		i -= len(m.CosmosSdkVersion)
		copy(dAtA[i:], m.CosmosSdkVersion)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CosmosSdkVersion)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.BuildDeps) > 0 {
		for iNdEx := len(m.BuildDeps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BuildDeps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GoVersion) > 0 {
		i -= len(m.GoVersion)
		copy(dAtA[i:], m.GoVersion)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GoVersion)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BuildTags) > 0 {
		i -= len(m.BuildTags)
		copy(dAtA[i:], m.BuildTags)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BuildTags)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GitCommit) > 0 {
		i -= len(m.GitCommit)
		copy(dAtA[i:], m.GitCommit)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.GitCommit)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AppName) > 0 {
		i -= len(m.AppName)
		copy(dAtA[i:], m.AppName)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.AppName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Module) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Module) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Module) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sum) > 0 {
		i -= len(m.Sum)
		copy(dAtA[i:], m.Sum)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Sum)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Path)))
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
func (m *GetValidatorSetByHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetValidatorSetByHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
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

func (m *GetLatestValidatorSetRequest) Size() (n int) {
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

func (m *GetLatestValidatorSetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
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

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovQuery(uint64(m.VotingPower))
	}
	if m.ProposerPriority != 0 {
		n += 1 + sovQuery(uint64(m.ProposerPriority))
	}
	return n
}

func (m *GetBlockByHeightRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	return n
}

func (m *GetBlockByHeightResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetLatestBlockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLatestBlockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetSyncingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetSyncingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Syncing {
		n += 2
	}
	return n
}

func (m *GetNodeInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetNodeInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultNodeInfo != nil {
		l = m.DefaultNodeInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.ApplicationVersion != nil {
		l = m.ApplicationVersion.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *VersionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.AppName)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GitCommit)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BuildTags)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.GoVersion)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.BuildDeps) > 0 {
		for _, e := range m.BuildDeps {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	l = len(m.CosmosSdkVersion)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *Module) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Sum)
	if l > 0 {
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
func (m *GetValidatorSetByHeightRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetValidatorSetByHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetValidatorSetByHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *GetValidatorSetByHeightResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetValidatorSetByHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetValidatorSetByHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
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
func (m *GetLatestValidatorSetRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLatestValidatorSetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestValidatorSetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *GetLatestValidatorSetResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLatestValidatorSetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestValidatorSetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, &Validator{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
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
func (m *Validator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
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
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerPriority", wireType)
			}
			m.ProposerPriority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposerPriority |= int64(b&0x7F) << shift
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
func (m *GetBlockByHeightRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetBlockByHeightRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockByHeightRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
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
func (m *GetBlockByHeightResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetBlockByHeightResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockByHeightResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
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
			if m.BlockId == nil {
				m.BlockId = &types1.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
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
			if m.Block == nil {
				m.Block = &types1.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetLatestBlockRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLatestBlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *GetLatestBlockResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLatestBlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLatestBlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
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
			if m.BlockId == nil {
				m.BlockId = &types1.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
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
			if m.Block == nil {
				m.Block = &types1.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetSyncingRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetSyncingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSyncingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *GetSyncingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetSyncingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSyncingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Syncing", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Syncing = bool(v != 0)
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
func (m *GetNodeInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetNodeInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodeInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *GetNodeInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetNodeInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetNodeInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultNodeInfo", wireType)
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
			if m.DefaultNodeInfo == nil {
				m.DefaultNodeInfo = &p2p.DefaultNodeInfo{}
			}
			if err := m.DefaultNodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationVersion", wireType)
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
			if m.ApplicationVersion == nil {
				m.ApplicationVersion = &VersionInfo{}
			}
			if err := m.ApplicationVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *VersionInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VersionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppName", wireType)
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
			m.AppName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
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
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitCommit", wireType)
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
			m.GitCommit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildTags", wireType)
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
			m.BuildTags = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoVersion", wireType)
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
			m.GoVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildDeps", wireType)
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
			m.BuildDeps = append(m.BuildDeps, &Module{})
			if err := m.BuildDeps[len(m.BuildDeps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CosmosSdkVersion", wireType)
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
			m.CosmosSdkVersion = string(dAtA[iNdEx:postIndex])
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
func (m *Module) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Module: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Module: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
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
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
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
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
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
			m.Sum = string(dAtA[iNdEx:postIndex])
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

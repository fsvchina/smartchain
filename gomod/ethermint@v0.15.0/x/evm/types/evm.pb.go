


package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3


type Params struct {


	EvmDenom string `protobuf:"bytes,1,opt,name=evm_denom,json=evmDenom,proto3" json:"evm_denom,omitempty" yaml:"evm_denom"`

	EnableCreate bool `protobuf:"varint,2,opt,name=enable_create,json=enableCreate,proto3" json:"enable_create,omitempty" yaml:"enable_create"`

	EnableCall bool `protobuf:"varint,3,opt,name=enable_call,json=enableCall,proto3" json:"enable_call,omitempty" yaml:"enable_call"`

	ExtraEIPs []int64 `protobuf:"varint,4,rep,packed,name=extra_eips,json=extraEips,proto3" json:"extra_eips,omitempty" yaml:"extra_eips"`

	ChainConfig ChainConfig `protobuf:"bytes,5,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"chain_config"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEvmDenom() string {
	if m != nil {
		return m.EvmDenom
	}
	return ""
}

func (m *Params) GetEnableCreate() bool {
	if m != nil {
		return m.EnableCreate
	}
	return false
}

func (m *Params) GetEnableCall() bool {
	if m != nil {
		return m.EnableCall
	}
	return false
}

func (m *Params) GetExtraEIPs() []int64 {
	if m != nil {
		return m.ExtraEIPs
	}
	return nil
}

func (m *Params) GetChainConfig() ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return ChainConfig{}
}



type ChainConfig struct {

	HomesteadBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=homestead_block,json=homesteadBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"homestead_block,omitempty" yaml:"homestead_block"`

	DAOForkBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=dao_fork_block,json=daoForkBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"dao_fork_block,omitempty" yaml:"dao_fork_block"`

	DAOForkSupport bool `protobuf:"varint,3,opt,name=dao_fork_support,json=daoForkSupport,proto3" json:"dao_fork_support,omitempty" yaml:"dao_fork_support"`


	EIP150Block *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=eip150_block,json=eip150Block,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eip150_block,omitempty" yaml:"eip150_block"`

	EIP150Hash string `protobuf:"bytes,5,opt,name=eip150_hash,json=eip150Hash,proto3" json:"eip150_hash,omitempty" yaml:"byzantium_block"`

	EIP155Block *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=eip155_block,json=eip155Block,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eip155_block,omitempty" yaml:"eip155_block"`

	EIP158Block *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=eip158_block,json=eip158Block,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"eip158_block,omitempty" yaml:"eip158_block"`

	ByzantiumBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=byzantium_block,json=byzantiumBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"byzantium_block,omitempty" yaml:"byzantium_block"`

	ConstantinopleBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=constantinople_block,json=constantinopleBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"constantinople_block,omitempty" yaml:"constantinople_block"`

	PetersburgBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=petersburg_block,json=petersburgBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"petersburg_block,omitempty" yaml:"petersburg_block"`

	IstanbulBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=istanbul_block,json=istanbulBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"istanbul_block,omitempty" yaml:"istanbul_block"`

	MuirGlacierBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=muir_glacier_block,json=muirGlacierBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"muir_glacier_block,omitempty" yaml:"muir_glacier_block"`

	BerlinBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=berlin_block,json=berlinBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"berlin_block,omitempty" yaml:"berlin_block"`

	LondonBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,17,opt,name=london_block,json=londonBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"london_block,omitempty" yaml:"london_block"`

	ArrowGlacierBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,18,opt,name=arrow_glacier_block,json=arrowGlacierBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"arrow_glacier_block,omitempty" yaml:"arrow_glacier_block"`

	MergeForkBlock *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,19,opt,name=merge_fork_block,json=mergeForkBlock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"merge_fork_block,omitempty" yaml:"merge_fork_block"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{1}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func (m *ChainConfig) GetDAOForkSupport() bool {
	if m != nil {
		return m.DAOForkSupport
	}
	return false
}

func (m *ChainConfig) GetEIP150Hash() string {
	if m != nil {
		return m.EIP150Hash
	}
	return ""
}


type State struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{2}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *State) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}




type TransactionLogs struct {
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Logs []*Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (m *TransactionLogs) Reset()         { *m = TransactionLogs{} }
func (m *TransactionLogs) String() string { return proto.CompactTextString(m) }
func (*TransactionLogs) ProtoMessage()    {}
func (*TransactionLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{3}
}
func (m *TransactionLogs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionLogs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionLogs.Merge(m, src)
}
func (m *TransactionLogs) XXX_Size() int {
	return m.Size()
}
func (m *TransactionLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionLogs.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionLogs proto.InternalMessageInfo

func (m *TransactionLogs) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *TransactionLogs) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}




type Log struct {

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`

	Topics []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`

	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`

	BlockNumber uint64 `protobuf:"varint,4,opt,name=block_number,json=blockNumber,proto3" json:"blockNumber"`

	TxHash string `protobuf:"bytes,5,opt,name=tx_hash,json=txHash,proto3" json:"transactionHash"`

	TxIndex uint64 `protobuf:"varint,6,opt,name=tx_index,json=txIndex,proto3" json:"transactionIndex"`

	BlockHash string `protobuf:"bytes,7,opt,name=block_hash,json=blockHash,proto3" json:"blockHash"`

	Index uint64 `protobuf:"varint,8,opt,name=index,proto3" json:"logIndex"`



	Removed bool `protobuf:"varint,9,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{4}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Log.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return m.Size()
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Log) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *Log) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Log) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *Log) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

func (m *Log) GetTxIndex() uint64 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *Log) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *Log) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Log) GetRemoved() bool {
	if m != nil {
		return m.Removed
	}
	return false
}


type TxResult struct {



	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`

	Bloom []byte `protobuf:"bytes,2,opt,name=bloom,proto3" json:"bloom,omitempty"`


	TxLogs TransactionLogs `protobuf:"bytes,3,opt,name=tx_logs,json=txLogs,proto3" json:"tx_logs" yaml:"tx_logs"`

	Ret []byte `protobuf:"bytes,4,opt,name=ret,proto3" json:"ret,omitempty"`

	Reverted bool `protobuf:"varint,5,opt,name=reverted,proto3" json:"reverted,omitempty"`

	GasUsed uint64 `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
}

func (m *TxResult) Reset()         { *m = TxResult{} }
func (m *TxResult) String() string { return proto.CompactTextString(m) }
func (*TxResult) ProtoMessage()    {}
func (*TxResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{5}
}
func (m *TxResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResult.Merge(m, src)
}
func (m *TxResult) XXX_Size() int {
	return m.Size()
}
func (m *TxResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResult.DiscardUnknown(m)
}

var xxx_messageInfo_TxResult proto.InternalMessageInfo


type AccessTuple struct {

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`

	StorageKeys []string `protobuf:"bytes,2,rep,name=storage_keys,json=storageKeys,proto3" json:"storageKeys"`
}

func (m *AccessTuple) Reset()         { *m = AccessTuple{} }
func (m *AccessTuple) String() string { return proto.CompactTextString(m) }
func (*AccessTuple) ProtoMessage()    {}
func (*AccessTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{6}
}
func (m *AccessTuple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessTuple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessTuple.Merge(m, src)
}
func (m *AccessTuple) XXX_Size() int {
	return m.Size()
}
func (m *AccessTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessTuple.DiscardUnknown(m)
}

var xxx_messageInfo_AccessTuple proto.InternalMessageInfo


type TraceConfig struct {

	Tracer string `protobuf:"bytes,1,opt,name=tracer,proto3" json:"tracer,omitempty"`


	Timeout string `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`

	Reexec uint64 `protobuf:"varint,3,opt,name=reexec,proto3" json:"reexec,omitempty"`

	DisableStack bool `protobuf:"varint,5,opt,name=disable_stack,json=disableStack,proto3" json:"disableStack"`

	DisableStorage bool `protobuf:"varint,6,opt,name=disable_storage,json=disableStorage,proto3" json:"disableStorage"`

	Debug bool `protobuf:"varint,8,opt,name=debug,proto3" json:"debug,omitempty"`

	Limit int32 `protobuf:"varint,9,opt,name=limit,proto3" json:"limit,omitempty"`

	Overrides *ChainConfig `protobuf:"bytes,10,opt,name=overrides,proto3" json:"overrides,omitempty"`

	EnableMemory bool `protobuf:"varint,11,opt,name=enable_memory,json=enableMemory,proto3" json:"enableMemory"`

	EnableReturnData bool `protobuf:"varint,12,opt,name=enable_return_data,json=enableReturnData,proto3" json:"enableReturnData"`
}

func (m *TraceConfig) Reset()         { *m = TraceConfig{} }
func (m *TraceConfig) String() string { return proto.CompactTextString(m) }
func (*TraceConfig) ProtoMessage()    {}
func (*TraceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21ecc92c8c8583e, []int{7}
}
func (m *TraceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TraceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TraceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TraceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceConfig.Merge(m, src)
}
func (m *TraceConfig) XXX_Size() int {
	return m.Size()
}
func (m *TraceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TraceConfig proto.InternalMessageInfo

func (m *TraceConfig) GetTracer() string {
	if m != nil {
		return m.Tracer
	}
	return ""
}

func (m *TraceConfig) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *TraceConfig) GetReexec() uint64 {
	if m != nil {
		return m.Reexec
	}
	return 0
}

func (m *TraceConfig) GetDisableStack() bool {
	if m != nil {
		return m.DisableStack
	}
	return false
}

func (m *TraceConfig) GetDisableStorage() bool {
	if m != nil {
		return m.DisableStorage
	}
	return false
}

func (m *TraceConfig) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *TraceConfig) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TraceConfig) GetOverrides() *ChainConfig {
	if m != nil {
		return m.Overrides
	}
	return nil
}

func (m *TraceConfig) GetEnableMemory() bool {
	if m != nil {
		return m.EnableMemory
	}
	return false
}

func (m *TraceConfig) GetEnableReturnData() bool {
	if m != nil {
		return m.EnableReturnData
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "ethermint.evm.v1.Params")
	proto.RegisterType((*ChainConfig)(nil), "ethermint.evm.v1.ChainConfig")
	proto.RegisterType((*State)(nil), "ethermint.evm.v1.State")
	proto.RegisterType((*TransactionLogs)(nil), "ethermint.evm.v1.TransactionLogs")
	proto.RegisterType((*Log)(nil), "ethermint.evm.v1.Log")
	proto.RegisterType((*TxResult)(nil), "ethermint.evm.v1.TxResult")
	proto.RegisterType((*AccessTuple)(nil), "ethermint.evm.v1.AccessTuple")
	proto.RegisterType((*TraceConfig)(nil), "ethermint.evm.v1.TraceConfig")
}

func init() { proto.RegisterFile("ethermint/evm/v1/evm.proto", fileDescriptor_d21ecc92c8c8583e) }

var fileDescriptor_d21ecc92c8c8583e = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0xdf, 0x6e, 0xdb, 0xb6,
	0x1a, 0x4f, 0x62, 0x27, 0x91, 0x69, 0xc7, 0x56, 0x98, 0x34, 0xc7, 0x4d, 0x71, 0xa2, 0x1c, 0x5d,
	0x1c, 0xe4, 0x00, 0x6d, 0xdc, 0xa4, 0x08, 0x4e, 0xd1, 0x62, 0x17, 0x51, 0x92, 0xb6, 0xc9, 0xba,
	0x2d, 0x60, 0x32, 0x0c, 0x18, 0x30, 0x08, 0xb4, 0xc4, 0xca, 0x5a, 0x24, 0xd1, 0x20, 0x29, 0xd7,
	0x1e, 0xf6, 0x00, 0x03, 0x76, 0xb3, 0x47, 0xd8, 0x2b, 0xec, 0x2d, 0x8a, 0x5d, 0xf5, 0x66, 0xc0,
	0xb0, 0x0b, 0xa1, 0x48, 0xef, 0x72, 0xe9, 0x27, 0x18, 0x44, 0xd2, 0x7f, 0x13, 0x6c, 0x4b, 0xae,
	0xcc, 0xdf, 0xf7, 0xe7, 0xf7, 0x23, 0x3f, 0x7e, 0x14, 0x69, 0xb0, 0x4e, 0x44, 0x8b, 0xb0, 0x38,
	0x4c, 0x44, 0x83, 0x74, 0xe2, 0x46, 0x67, 0x27, 0xff, 0xd9, 0x6e, 0x33, 0x2a, 0x28, 0x34, 0x87,
	0xbe, 0xed, 0xdc, 0xd8, 0xd9, 0x59, 0x5f, 0x0d, 0x68, 0x40, 0xa5, 0xb3, 0x91, 0x8f, 0x54, 0x9c,
	0xfd, 0xdb, 0x1c, 0x58, 0x38, 0xc5, 0x0c, 0xc7, 0x1c, 0xee, 0x80, 0x12, 0xe9, 0xc4, 0xae, 0x4f,
	0x12, 0x1a, 0xd7, 0x67, 0x37, 0x67, 0xb7, 0x4a, 0xce, 0x6a, 0x3f, 0xb3, 0xcc, 0x1e, 0x8e, 0xa3,
	0x67, 0xf6, 0xd0, 0x65, 0x23, 0x83, 0x74, 0xe2, 0xc3, 0x7c, 0x08, 0x3f, 0x01, 0x4b, 0x24, 0xc1,
	0xcd, 0x88, 0xb8, 0x1e, 0x23, 0x58, 0x90, 0xfa, 0xdc, 0xe6, 0xec, 0x96, 0xe1, 0xd4, 0xfb, 0x99,
	0xb5, 0xaa, 0xd3, 0xc6, 0xdd, 0x36, 0xaa, 0x28, 0x7c, 0x20, 0x21, 0xfc, 0x3f, 0x28, 0x0f, 0xfc,
	0x38, 0x8a, 0xea, 0x05, 0x99, 0xbc, 0xd6, 0xcf, 0x2c, 0x38, 0x99, 0x8c, 0xa3, 0xc8, 0x46, 0x40,
	0xa7, 0xe2, 0x28, 0x82, 0xfb, 0x00, 0x90, 0xae, 0x60, 0xd8, 0x25, 0x61, 0x9b, 0xd7, 0x8b, 0x9b,
	0x85, 0xad, 0x82, 0x63, 0x5f, 0x66, 0x56, 0xe9, 0x28, 0xb7, 0x1e, 0x1d, 0x9f, 0xf2, 0x7e, 0x66,
	0x2d, 0x6b, 0x92, 0x61, 0xa0, 0x8d, 0x4a, 0x12, 0x1c, 0x85, 0x6d, 0x0e, 0xbf, 0x01, 0x15, 0xaf,
	0x85, 0xc3, 0xc4, 0xf5, 0x68, 0xf2, 0x26, 0x0c, 0xea, 0xf3, 0x9b, 0xb3, 0x5b, 0xe5, 0xdd, 0x7f,
	0x6f, 0x4f, 0xd7, 0x6d, 0xfb, 0x20, 0x8f, 0x3a, 0x90, 0x41, 0xce, 0x83, 0x77, 0x99, 0x35, 0xd3,
	0xcf, 0xac, 0x15, 0x45, 0x3d, 0x4e, 0x60, 0xa3, 0xb2, 0x37, 0x8a, 0xb4, 0x7f, 0xa9, 0x82, 0xf2,
	0x58, 0x26, 0x8c, 0x41, 0xad, 0x45, 0x63, 0xc2, 0x05, 0xc1, 0xbe, 0xdb, 0x8c, 0xa8, 0x77, 0xa1,
	0x4b, 0x7c, 0xf8, 0x47, 0x66, 0xfd, 0x37, 0x08, 0x45, 0x2b, 0x6d, 0x6e, 0x7b, 0x34, 0x6e, 0x78,
	0x94, 0xc7, 0x94, 0xeb, 0x9f, 0x47, 0xdc, 0xbf, 0x68, 0x88, 0x5e, 0x9b, 0xf0, 0xed, 0xe3, 0x44,
	0xf4, 0x33, 0x6b, 0x4d, 0x09, 0x4f, 0x51, 0xd9, 0xa8, 0x3a, 0xb4, 0x38, 0xb9, 0x01, 0xf6, 0x40,
	0xd5, 0xc7, 0xd4, 0x7d, 0x43, 0xd9, 0x85, 0x56, 0x9b, 0x93, 0x6a, 0x67, 0xff, 0x5c, 0xed, 0x32,
	0xb3, 0x2a, 0x87, 0xfb, 0x5f, 0xbc, 0xa0, 0xec, 0x42, 0x72, 0xf6, 0x33, 0xeb, 0x9e, 0x52, 0x9f,
	0x64, 0xb6, 0x51, 0xc5, 0xc7, 0x74, 0x18, 0x06, 0xbf, 0x02, 0xe6, 0x30, 0x80, 0xa7, 0xed, 0x36,
	0x65, 0x42, 0xef, 0xec, 0xa3, 0xcb, 0xcc, 0xaa, 0x6a, 0xca, 0x33, 0xe5, 0xe9, 0x67, 0xd6, 0xbf,
	0xa6, 0x48, 0x75, 0x8e, 0x8d, 0xaa, 0x9a, 0x56, 0x87, 0x42, 0x0e, 0x2a, 0x24, 0x6c, 0xef, 0xec,
	0x3d, 0xd6, 0x2b, 0x2a, 0xca, 0x15, 0x9d, 0xde, 0x6a, 0x45, 0xe5, 0xa3, 0xe3, 0xd3, 0x9d, 0xbd,
	0xc7, 0x83, 0x05, 0xe9, 0x7d, 0x1c, 0xa7, 0xb5, 0x51, 0x59, 0x41, 0xb5, 0x9a, 0x63, 0xa0, 0xa1,
	0xdb, 0xc2, 0xbc, 0x25, 0xbb, 0xa4, 0xe4, 0x6c, 0x5d, 0x66, 0x16, 0x50, 0x4c, 0xaf, 0x30, 0x6f,
	0x8d, 0xf6, 0xa5, 0xd9, 0xfb, 0x0e, 0x27, 0x22, 0x4c, 0xe3, 0x01, 0x17, 0x50, 0xc9, 0x79, 0xd4,
	0x70, 0xfe, 0x7b, 0x7a, 0xfe, 0x0b, 0x77, 0x9e, 0xff, 0xde, 0x4d, 0xf3, 0xdf, 0x9b, 0x9c, 0xbf,
	0x8a, 0x19, 0x8a, 0x3e, 0xd5, 0xa2, 0x8b, 0x77, 0x16, 0x7d, 0x7a, 0x93, 0xe8, 0xd3, 0x49, 0x51,
	0x15, 0x93, 0x37, 0xfb, 0x54, 0x25, 0xea, 0xc6, 0xdd, 0x9b, 0xfd, 0x5a, 0x51, 0xab, 0x43, 0x8b,
	0x92, 0xfb, 0x1e, 0xac, 0x7a, 0x34, 0xe1, 0x22, 0xb7, 0x25, 0xb4, 0x1d, 0x11, 0xad, 0x59, 0x92,
	0x9a, 0xc7, 0xb7, 0xd2, 0x7c, 0xa0, 0x4f, 0xf6, 0x0d, 0x7c, 0x36, 0x5a, 0x99, 0x34, 0x2b, 0xf5,
	0x36, 0x30, 0xdb, 0x44, 0x10, 0xc6, 0x9b, 0x29, 0x0b, 0xb4, 0x32, 0x90, 0xca, 0x47, 0xb7, 0x52,
	0xd6, 0xe7, 0x60, 0x9a, 0xcb, 0x46, 0xb5, 0x91, 0x49, 0x29, 0x7e, 0x0b, 0xaa, 0x61, 0x3e, 0x8d,
	0x66, 0x1a, 0x69, 0xbd, 0xb2, 0xd4, 0x3b, 0xb8, 0x95, 0x9e, 0x3e, 0xcc, 0x93, 0x4c, 0x36, 0x5a,
	0x1a, 0x18, 0x94, 0x56, 0x0a, 0x60, 0x9c, 0x86, 0xcc, 0x0d, 0x22, 0xec, 0x85, 0x84, 0x69, 0xbd,
	0x8a, 0xd4, 0x7b, 0x79, 0x2b, 0xbd, 0xfb, 0x4a, 0xef, 0x3a, 0x9b, 0x8d, 0xcc, 0xdc, 0xf8, 0x52,
	0xd9, 0x94, 0xac, 0x0f, 0x2a, 0x4d, 0xc2, 0xa2, 0x30, 0xd1, 0x82, 0x4b, 0x52, 0x70, 0xff, 0x56,
	0x82, 0xba, 0x4f, 0xc7, 0x79, 0x6c, 0x54, 0x56, 0x70, 0xa8, 0x12, 0xd1, 0xc4, 0xa7, 0x03, 0x95,
	0xe5, 0xbb, 0xab, 0x8c, 0xf3, 0xd8, 0xa8, 0xac, 0xa0, 0x52, 0xe9, 0x82, 0x15, 0xcc, 0x18, 0x7d,
	0x3b, 0x55, 0x43, 0x28, 0xc5, 0x5e, 0xdd, 0x4a, 0x6c, 0x5d, 0x89, 0xdd, 0x40, 0x67, 0xa3, 0x65,
	0x69, 0x9d, 0xa8, 0x22, 0x05, 0x66, 0x4c, 0x58, 0x40, 0xc6, 0xef, 0x81, 0x95, 0xbb, 0xb7, 0xe6,
	0x34, 0x97, 0x8d, 0xaa, 0xd2, 0x34, 0xfc, 0xf6, 0x9f, 0x14, 0x8d, 0xaa, 0x59, 0x3b, 0x29, 0x1a,
	0x35, 0xd3, 0x3c, 0x29, 0x1a, 0xa6, 0xb9, 0x8c, 0x96, 0x7a, 0x34, 0xa2, 0x6e, 0xe7, 0x89, 0xca,
	0x40, 0x65, 0xf2, 0x16, 0x73, 0x7d, 0x90, 0x51, 0xd5, 0xc3, 0x02, 0x47, 0x3d, 0x2e, 0x34, 0x5d,
	0x03, 0xcc, 0x9f, 0x89, 0xfc, 0x5d, 0x60, 0x82, 0xc2, 0x05, 0xe9, 0xa9, 0x0b, 0x12, 0xe5, 0x43,
	0xb8, 0x0a, 0xe6, 0x3b, 0x38, 0x4a, 0xd5, 0x03, 0xa3, 0x84, 0x14, 0xb0, 0x4f, 0x41, 0xed, 0x9c,
	0xe1, 0x84, 0x63, 0x4f, 0x84, 0x34, 0x79, 0x4d, 0x03, 0x0e, 0x21, 0x28, 0xca, 0x0f, 0xb5, 0xca,
	0x95, 0x63, 0xf8, 0x3f, 0x50, 0x8c, 0x68, 0xc0, 0xeb, 0x73, 0x9b, 0x85, 0xad, 0xf2, 0xee, 0xbd,
	0xeb, 0x57, 0xfc, 0x6b, 0x1a, 0x20, 0x19, 0x62, 0xff, 0x3a, 0x07, 0x0a, 0xaf, 0x69, 0x00, 0xeb,
	0x60, 0x11, 0xfb, 0x3e, 0x23, 0x9c, 0x6b, 0xa6, 0x01, 0x84, 0x6b, 0x60, 0x41, 0xd0, 0x76, 0xe8,
	0x29, 0xba, 0x12, 0xd2, 0x28, 0x17, 0xf6, 0xb1, 0xc0, 0xf2, 0xaa, 0xab, 0x20, 0x39, 0x86, 0xbb,
	0xa0, 0x22, 0x57, 0xe6, 0x26, 0x69, 0xdc, 0x24, 0x4c, 0xde, 0x58, 0x45, 0xa7, 0x76, 0x95, 0x59,
	0x65, 0x69, 0xff, 0x5c, 0x9a, 0xd1, 0x38, 0x80, 0x0f, 0xc1, 0xa2, 0xe8, 0x8e, 0x5f, 0x36, 0x2b,
	0x57, 0x99, 0x55, 0x13, 0xa3, 0x65, 0xe6, 0x77, 0x09, 0x5a, 0x10, 0x5d, 0x79, 0xa7, 0x34, 0x80,
	0x21, 0xba, 0x6e, 0x98, 0xf8, 0xa4, 0x2b, 0xef, 0x93, 0xa2, 0xb3, 0x7a, 0x95, 0x59, 0xe6, 0x58,
	0xf8, 0x71, 0xee, 0x43, 0x8b, 0xa2, 0x2b, 0x07, 0xf0, 0x21, 0x00, 0x6a, 0x4a, 0x52, 0x41, 0xdd,
	0x06, 0x4b, 0x57, 0x99, 0x55, 0x92, 0x56, 0xc9, 0x3d, 0x1a, 0x42, 0x1b, 0xcc, 0x2b, 0x6e, 0x43,
	0x72, 0x57, 0xae, 0x32, 0xcb, 0x88, 0x68, 0xa0, 0x38, 0x95, 0x2b, 0x2f, 0x15, 0x23, 0x31, 0xed,
	0x10, 0x5f, 0x7e, 0x70, 0x0d, 0x34, 0x80, 0xf6, 0x8f, 0x73, 0xc0, 0x38, 0xef, 0x22, 0xc2, 0xd3,
	0x48, 0xc0, 0x17, 0xc0, 0xf4, 0x68, 0x22, 0x18, 0xf6, 0x84, 0x3b, 0x51, 0x5a, 0xe7, 0xc1, 0xa8,
	0xc3, 0xa6, 0x23, 0x6c, 0x54, 0x1b, 0x98, 0xf6, 0x75, 0xfd, 0x57, 0xc1, 0x7c, 0x33, 0xa2, 0x34,
	0x96, 0x9d, 0x50, 0x41, 0x0a, 0x40, 0x24, 0xab, 0x26, 0x77, 0xb9, 0x20, 0x1f, 0x72, 0xff, 0xb9,
	0xbe, 0xcb, 0x53, 0xad, 0xe2, 0xac, 0xe9, 0xc7, 0x5c, 0x55, 0x69, 0xeb, 0x7c, 0x3b, 0xaf, 0xad,
	0x6c, 0x25, 0x13, 0x14, 0x18, 0x11, 0x72, 0xd3, 0x2a, 0x28, 0x1f, 0xc2, 0x75, 0x60, 0x30, 0xd2,
	0x21, 0x4c, 0x10, 0x5f, 0x6e, 0x8e, 0x81, 0x86, 0x18, 0xde, 0x07, 0x46, 0x80, 0xb9, 0x9b, 0x72,
	0xe2, 0xab, 0x9d, 0x40, 0x8b, 0x01, 0xe6, 0x5f, 0x72, 0xe2, 0x3f, 0x2b, 0xfe, 0xf0, 0xb3, 0x35,
	0x63, 0x63, 0x50, 0xde, 0xf7, 0x3c, 0xc2, 0xf9, 0x79, 0xda, 0x8e, 0xc8, 0x5f, 0x74, 0xd8, 0x2e,
	0xa8, 0x70, 0x41, 0x19, 0x0e, 0x88, 0x7b, 0x41, 0x7a, 0xba, 0xcf, 0x54, 0xd7, 0x68, 0xfb, 0xa7,
	0xa4, 0xc7, 0xd1, 0x38, 0xd0, 0x12, 0x1f, 0x0a, 0xa0, 0x7c, 0xce, 0xb0, 0x47, 0xf4, 0xa3, 0x33,
	0xef, 0xd5, 0x1c, 0x32, 0x2d, 0xa1, 0x51, 0xae, 0x2d, 0xc2, 0x98, 0xd0, 0x54, 0xe8, 0xf3, 0x34,
	0x80, 0x79, 0x06, 0x23, 0xa4, 0x4b, 0x3c, 0x59, 0xc6, 0x22, 0xd2, 0x08, 0xee, 0x81, 0x25, 0x3f,
	0xe4, 0xf2, 0x35, 0xce, 0x05, 0xf6, 0x2e, 0xd4, 0xf2, 0x1d, 0xf3, 0x2a, 0xb3, 0x2a, 0xda, 0x71,
	0x96, 0xdb, 0xd1, 0x04, 0x82, 0xcf, 0x41, 0x6d, 0x94, 0x26, 0x67, 0x2b, 0x6b, 0x63, 0x38, 0xf0,
	0x2a, 0xb3, 0xaa, 0xc3, 0x50, 0xe9, 0x41, 0x53, 0x38, 0xdf, 0x69, 0x9f, 0x34, 0xd3, 0x40, 0x36,
	0x9f, 0x81, 0x14, 0xc8, 0xad, 0x51, 0x18, 0x87, 0x42, 0x36, 0xdb, 0x3c, 0x52, 0x00, 0x3e, 0x07,
	0x25, 0xda, 0x21, 0x8c, 0x85, 0x3e, 0xe1, 0xf2, 0xf6, 0xfd, 0xbb, 0xa7, 0x3c, 0x1a, 0xc5, 0xe7,
	0x8b, 0xd3, 0xff, 0x34, 0x62, 0x12, 0x53, 0xd6, 0x93, 0xd7, 0xa9, 0x5e, 0x9c, 0x72, 0x7c, 0x26,
	0xed, 0x68, 0x02, 0x41, 0x07, 0x40, 0x9d, 0xc6, 0x88, 0x48, 0x59, 0xe2, 0xca, 0xf3, 0x5f, 0x91,
	0xb9, 0xf2, 0x14, 0x2a, 0x2f, 0x92, 0xce, 0x43, 0x2c, 0x30, 0xba, 0x66, 0x39, 0x29, 0x1a, 0x45,
	0x73, 0xfe, 0xa4, 0x68, 0x2c, 0x9a, 0xc6, 0x70, 0xfd, 0x7a, 0x16, 0x68, 0x65, 0x80, 0xc7, 0xe8,
	0x1d, 0xe7, 0xdd, 0xe5, 0xc6, 0xec, 0xfb, 0xcb, 0x8d, 0xd9, 0x0f, 0x97, 0x1b, 0xb3, 0x3f, 0x7d,
	0xdc, 0x98, 0x79, 0xff, 0x71, 0x63, 0xe6, 0xf7, 0x8f, 0x1b, 0x33, 0x5f, 0x6f, 0x8d, 0x7d, 0xce,
	0x45, 0x0b, 0x33, 0x1e, 0xf2, 0xc6, 0xe8, 0x0f, 0x62, 0x57, 0xfe, 0x45, 0x94, 0x1f, 0xf5, 0xe6,
	0x82, 0xfc, 0xeb, 0xf7, 0xe4, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xef, 0x82, 0x39, 0x40,
	0x0e, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.ExtraEIPs) > 0 {
		dAtA3 := make([]byte, len(m.ExtraEIPs)*10)
		var j2 int
		for _, num1 := range m.ExtraEIPs {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintEvm(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0x22
	}
	if m.EnableCall {
		i--
		if m.EnableCall {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.EnableCreate {
		i--
		if m.EnableCreate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.EvmDenom) > 0 {
		i -= len(m.EvmDenom)
		copy(dAtA[i:], m.EvmDenom)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.EvmDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MergeForkBlock != nil {
		{
			size := m.MergeForkBlock.Size()
			i -= size
			if _, err := m.MergeForkBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x9a
	}
	if m.ArrowGlacierBlock != nil {
		{
			size := m.ArrowGlacierBlock.Size()
			i -= size
			if _, err := m.ArrowGlacierBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x92
	}
	if m.LondonBlock != nil {
		{
			size := m.LondonBlock.Size()
			i -= size
			if _, err := m.LondonBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.BerlinBlock != nil {
		{
			size := m.BerlinBlock.Size()
			i -= size
			if _, err := m.BerlinBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.MuirGlacierBlock != nil {
		{
			size := m.MuirGlacierBlock.Size()
			i -= size
			if _, err := m.MuirGlacierBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.IstanbulBlock != nil {
		{
			size := m.IstanbulBlock.Size()
			i -= size
			if _, err := m.IstanbulBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.PetersburgBlock != nil {
		{
			size := m.PetersburgBlock.Size()
			i -= size
			if _, err := m.PetersburgBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.ConstantinopleBlock != nil {
		{
			size := m.ConstantinopleBlock.Size()
			i -= size
			if _, err := m.ConstantinopleBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.ByzantiumBlock != nil {
		{
			size := m.ByzantiumBlock.Size()
			i -= size
			if _, err := m.ByzantiumBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.EIP158Block != nil {
		{
			size := m.EIP158Block.Size()
			i -= size
			if _, err := m.EIP158Block.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.EIP155Block != nil {
		{
			size := m.EIP155Block.Size()
			i -= size
			if _, err := m.EIP155Block.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.EIP150Hash) > 0 {
		i -= len(m.EIP150Hash)
		copy(dAtA[i:], m.EIP150Hash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.EIP150Hash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.EIP150Block != nil {
		{
			size := m.EIP150Block.Size()
			i -= size
			if _, err := m.EIP150Block.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.DAOForkSupport {
		i--
		if m.DAOForkSupport {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.DAOForkBlock != nil {
		{
			size := m.DAOForkBlock.Size()
			i -= size
			if _, err := m.DAOForkBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.HomesteadBlock != nil {
		{
			size := m.HomesteadBlock.Size()
			i -= size
			if _, err := m.HomesteadBlock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *State) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TransactionLogs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionLogs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionLogs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Logs) > 0 {
		for iNdEx := len(m.Logs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Logs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvm(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Log) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Log) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Log) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Removed {
		i--
		if m.Removed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Index != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x40
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x3a
	}
	if m.TxIndex != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.TxIndex))
		i--
		dAtA[i] = 0x30
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.BlockNumber != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.BlockNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Topics[iNdEx])
			copy(dAtA[i:], m.Topics[iNdEx])
			i = encodeVarintEvm(dAtA, i, uint64(len(m.Topics[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TxResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.GasUsed != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x30
	}
	if m.Reverted {
		i--
		if m.Reverted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ret) > 0 {
		i -= len(m.Ret)
		copy(dAtA[i:], m.Ret)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Ret)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.TxLogs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvm(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bloom) > 0 {
		i -= len(m.Bloom)
		copy(dAtA[i:], m.Bloom)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Bloom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccessTuple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessTuple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessTuple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StorageKeys) > 0 {
		for iNdEx := len(m.StorageKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.StorageKeys[iNdEx])
			copy(dAtA[i:], m.StorageKeys[iNdEx])
			i = encodeVarintEvm(dAtA, i, uint64(len(m.StorageKeys[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TraceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TraceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TraceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EnableReturnData {
		i--
		if m.EnableReturnData {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if m.EnableMemory {
		i--
		if m.EnableMemory {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if m.Overrides != nil {
		{
			size, err := m.Overrides.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvm(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.Limit != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x48
	}
	if m.Debug {
		i--
		if m.Debug {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.DisableStorage {
		i--
		if m.DisableStorage {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.DisableStack {
		i--
		if m.DisableStack {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Reexec != 0 {
		i = encodeVarintEvm(dAtA, i, uint64(m.Reexec))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Timeout) > 0 {
		i -= len(m.Timeout)
		copy(dAtA[i:], m.Timeout)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Timeout)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Tracer) > 0 {
		i -= len(m.Tracer)
		copy(dAtA[i:], m.Tracer)
		i = encodeVarintEvm(dAtA, i, uint64(len(m.Tracer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvm(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvm(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EvmDenom)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EnableCreate {
		n += 2
	}
	if m.EnableCall {
		n += 2
	}
	if len(m.ExtraEIPs) > 0 {
		l = 0
		for _, e := range m.ExtraEIPs {
			l += sovEvm(uint64(e))
		}
		n += 1 + sovEvm(uint64(l)) + l
	}
	l = m.ChainConfig.Size()
	n += 1 + l + sovEvm(uint64(l))
	return n
}

func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HomesteadBlock != nil {
		l = m.HomesteadBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.DAOForkBlock != nil {
		l = m.DAOForkBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.DAOForkSupport {
		n += 2
	}
	if m.EIP150Block != nil {
		l = m.EIP150Block.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.EIP150Hash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EIP155Block != nil {
		l = m.EIP155Block.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EIP158Block != nil {
		l = m.EIP158Block.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.ByzantiumBlock != nil {
		l = m.ByzantiumBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.ConstantinopleBlock != nil {
		l = m.ConstantinopleBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.PetersburgBlock != nil {
		l = m.PetersburgBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.IstanbulBlock != nil {
		l = m.IstanbulBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.MuirGlacierBlock != nil {
		l = m.MuirGlacierBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.BerlinBlock != nil {
		l = m.BerlinBlock.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.LondonBlock != nil {
		l = m.LondonBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.ArrowGlacierBlock != nil {
		l = m.ArrowGlacierBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	if m.MergeForkBlock != nil {
		l = m.MergeForkBlock.Size()
		n += 2 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *State) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	return n
}

func (m *TransactionLogs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.Logs) > 0 {
		for _, e := range m.Logs {
			l = e.Size()
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	return n
}

func (m *Log) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.Topics) > 0 {
		for _, s := range m.Topics {
			l = len(s)
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.BlockNumber != 0 {
		n += 1 + sovEvm(uint64(m.BlockNumber))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.TxIndex != 0 {
		n += 1 + sovEvm(uint64(m.TxIndex))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovEvm(uint64(m.Index))
	}
	if m.Removed {
		n += 2
	}
	return n
}

func (m *TxResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Bloom)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = m.TxLogs.Size()
	n += 1 + l + sovEvm(uint64(l))
	l = len(m.Ret)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Reverted {
		n += 2
	}
	if m.GasUsed != 0 {
		n += 1 + sovEvm(uint64(m.GasUsed))
	}
	return n
}

func (m *AccessTuple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if len(m.StorageKeys) > 0 {
		for _, s := range m.StorageKeys {
			l = len(s)
			n += 1 + l + sovEvm(uint64(l))
		}
	}
	return n
}

func (m *TraceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tracer)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	l = len(m.Timeout)
	if l > 0 {
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.Reexec != 0 {
		n += 1 + sovEvm(uint64(m.Reexec))
	}
	if m.DisableStack {
		n += 2
	}
	if m.DisableStorage {
		n += 2
	}
	if m.Debug {
		n += 2
	}
	if m.Limit != 0 {
		n += 1 + sovEvm(uint64(m.Limit))
	}
	if m.Overrides != nil {
		l = m.Overrides.Size()
		n += 1 + l + sovEvm(uint64(l))
	}
	if m.EnableMemory {
		n += 2
	}
	if m.EnableReturnData {
		n += 2
	}
	return n
}

func sovEvm(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvm(x uint64) (n int) {
	return sovEvm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvmDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvmDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableCreate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.EnableCreate = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableCall", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.EnableCall = bool(v != 0)
		case 4:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ExtraEIPs = append(m.ExtraEIPs, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvm
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvm
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvm
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ExtraEIPs) == 0 {
					m.ExtraEIPs = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvm
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ExtraEIPs = append(m.ExtraEIPs, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraEIPs", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomesteadBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.HomesteadBlock = &v
			if err := m.HomesteadBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAOForkBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.DAOForkBlock = &v
			if err := m.DAOForkBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DAOForkSupport", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.DAOForkSupport = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP150Block", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.EIP150Block = &v
			if err := m.EIP150Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP150Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EIP150Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP155Block", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.EIP155Block = &v
			if err := m.EIP155Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EIP158Block", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.EIP158Block = &v
			if err := m.EIP158Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ByzantiumBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.ByzantiumBlock = &v
			if err := m.ByzantiumBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantinopleBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.ConstantinopleBlock = &v
			if err := m.ConstantinopleBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PetersburgBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.PetersburgBlock = &v
			if err := m.PetersburgBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IstanbulBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.IstanbulBlock = &v
			if err := m.IstanbulBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MuirGlacierBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.MuirGlacierBlock = &v
			if err := m.MuirGlacierBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BerlinBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.BerlinBlock = &v
			if err := m.BerlinBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LondonBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.LondonBlock = &v
			if err := m.LondonBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArrowGlacierBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.ArrowGlacierBlock = &v
			if err := m.ArrowGlacierBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MergeForkBlock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.MergeForkBlock = &v
			if err := m.MergeForkBlock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *State) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *TransactionLogs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: TransactionLogs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionLogs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Logs = append(m.Logs, &Log{})
			if err := m.Logs[len(m.Logs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *Log) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: Log: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Log: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = append(m.Topics, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockNumber", wireType)
			}
			m.BlockNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			m.TxIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.Removed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *TxResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: TxResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bloom", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bloom = append(m.Bloom[:0], dAtA[iNdEx:postIndex]...)
			if m.Bloom == nil {
				m.Bloom = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxLogs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxLogs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ret = append(m.Ret[:0], dAtA[iNdEx:postIndex]...)
			if m.Ret == nil {
				m.Ret = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.Reverted = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *AccessTuple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: AccessTuple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessTuple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageKeys = append(m.StorageKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func (m *TraceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvm
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
			return fmt.Errorf("proto: TraceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TraceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tracer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tracer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timeout = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reexec", wireType)
			}
			m.Reexec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reexec |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableStack", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.DisableStack = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableStorage", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.DisableStorage = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Debug", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.Debug = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Overrides", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
				return ErrInvalidLengthEvm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Overrides == nil {
				m.Overrides = &ChainConfig{}
			}
			if err := m.Overrides.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableMemory", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.EnableMemory = bool(v != 0)
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableReturnData", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvm
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
			m.EnableReturnData = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvm
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
func skipEvm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvm
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
					return 0, ErrIntOverflowEvm
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
					return 0, ErrIntOverflowEvm
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
				return 0, ErrInvalidLengthEvm
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvm
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvm
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvm        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvm          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvm = fmt.Errorf("proto: unexpected end of group")
)




package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen





const _ = proto.GoGoProtoPackageIsVersion3


type VoteOption int32

const (

	OptionEmpty VoteOption = 0

	OptionYes VoteOption = 1

	OptionAbstain VoteOption = 2

	OptionNo VoteOption = 3

	OptionNoWithVeto VoteOption = 4
)

var VoteOption_name = map[int32]string{
	0: "VOTE_OPTION_UNSPECIFIED",
	1: "VOTE_OPTION_YES",
	2: "VOTE_OPTION_ABSTAIN",
	3: "VOTE_OPTION_NO",
	4: "VOTE_OPTION_NO_WITH_VETO",
}

var VoteOption_value = map[string]int32{
	"VOTE_OPTION_UNSPECIFIED":  0,
	"VOTE_OPTION_YES":          1,
	"VOTE_OPTION_ABSTAIN":      2,
	"VOTE_OPTION_NO":           3,
	"VOTE_OPTION_NO_WITH_VETO": 4,
}

func (x VoteOption) String() string {
	return proto.EnumName(VoteOption_name, int32(x))
}

func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{0}
}


type ProposalStatus int32

const (

	StatusNil ProposalStatus = 0


	StatusDepositPeriod ProposalStatus = 1


	StatusVotingPeriod ProposalStatus = 2


	StatusPassed ProposalStatus = 3


	StatusRejected ProposalStatus = 4


	StatusFailed ProposalStatus = 5
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_UNSPECIFIED",
	1: "PROPOSAL_STATUS_DEPOSIT_PERIOD",
	2: "PROPOSAL_STATUS_VOTING_PERIOD",
	3: "PROPOSAL_STATUS_PASSED",
	4: "PROPOSAL_STATUS_REJECTED",
	5: "PROPOSAL_STATUS_FAILED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_UNSPECIFIED":    0,
	"PROPOSAL_STATUS_DEPOSIT_PERIOD": 1,
	"PROPOSAL_STATUS_VOTING_PERIOD":  2,
	"PROPOSAL_STATUS_PASSED":         3,
	"PROPOSAL_STATUS_REJECTED":       4,
	"PROPOSAL_STATUS_FAILED":         5,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{1}
}


//

type WeightedVoteOption struct {
	Option VoteOption                             `protobuf:"varint,1,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"`
	Weight github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WeightedVoteOption) Reset()      { *m = WeightedVoteOption{} }
func (*WeightedVoteOption) ProtoMessage() {}
func (*WeightedVoteOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{0}
}
func (m *WeightedVoteOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedVoteOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedVoteOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedVoteOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedVoteOption.Merge(m, src)
}
func (m *WeightedVoteOption) XXX_Size() int {
	return m.Size()
}
func (m *WeightedVoteOption) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedVoteOption.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedVoteOption proto.InternalMessageInfo



type TextProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *TextProposal) Reset()      { *m = TextProposal{} }
func (*TextProposal) ProtoMessage() {}
func (*TextProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{1}
}
func (m *TextProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TextProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TextProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TextProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextProposal.Merge(m, src)
}
func (m *TextProposal) XXX_Size() int {
	return m.Size()
}
func (m *TextProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_TextProposal.DiscardUnknown(m)
}

var xxx_messageInfo_TextProposal proto.InternalMessageInfo



type Deposit struct {
	ProposalId uint64                                   `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty" yaml:"proposal_id"`
	Depositor  string                                   `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	Amount     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Deposit) Reset()      { *m = Deposit{} }
func (*Deposit) ProtoMessage() {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{2}
}
func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return m.Size()
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo


type Proposal struct {
	ProposalId       uint64                                   `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"id" yaml:"id"`
	Content          *types1.Any                              `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status           ProposalStatus                           `protobuf:"varint,3,opt,name=status,proto3,enum=cosmos.gov.v1beta1.ProposalStatus" json:"status,omitempty" yaml:"proposal_status"`
	FinalTallyResult TallyResult                              `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result" yaml:"final_tally_result"`
	SubmitTime       time.Time                                `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3,stdtime" json:"submit_time" yaml:"submit_time"`
	DepositEndTime   time.Time                                `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3,stdtime" json:"deposit_end_time" yaml:"deposit_end_time"`
	TotalDeposit     github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=total_deposit,json=totalDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_deposit" yaml:"total_deposit"`
	VotingStartTime  time.Time                                `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3,stdtime" json:"voting_start_time" yaml:"voting_start_time"`
	VotingEndTime    time.Time                                `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3,stdtime" json:"voting_end_time" yaml:"voting_end_time"`
}

func (m *Proposal) Reset()      { *m = Proposal{} }
func (*Proposal) ProtoMessage() {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{3}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo


type TallyResult struct {
	Yes        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=yes,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"yes"`
	Abstain    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=abstain,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"abstain"`
	No         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=no,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"no"`
	NoWithVeto github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=no_with_veto,json=noWithVeto,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"no_with_veto" yaml:"no_with_veto"`
}

func (m *TallyResult) Reset()      { *m = TallyResult{} }
func (*TallyResult) ProtoMessage() {}
func (*TallyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{4}
}
func (m *TallyResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TallyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TallyResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TallyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TallyResult.Merge(m, src)
}
func (m *TallyResult) XXX_Size() int {
	return m.Size()
}
func (m *TallyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TallyResult.DiscardUnknown(m)
}

var xxx_messageInfo_TallyResult proto.InternalMessageInfo



type Vote struct {
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty" yaml:"proposal_id"`
	Voter      string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`



	Option VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=cosmos.gov.v1beta1.VoteOption" json:"option,omitempty"`

	Options []WeightedVoteOption `protobuf:"bytes,4,rep,name=options,proto3" json:"options"`
}

func (m *Vote) Reset()      { *m = Vote{} }
func (*Vote) ProtoMessage() {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{5}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo


type DepositParams struct {

	MinDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=min_deposit,json=minDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_deposit,omitempty" yaml:"min_deposit"`


	MaxDepositPeriod time.Duration `protobuf:"bytes,2,opt,name=max_deposit_period,json=maxDepositPeriod,proto3,stdduration" json:"max_deposit_period,omitempty" yaml:"max_deposit_period"`
}

func (m *DepositParams) Reset()      { *m = DepositParams{} }
func (*DepositParams) ProtoMessage() {}
func (*DepositParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{6}
}
func (m *DepositParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositParams.Merge(m, src)
}
func (m *DepositParams) XXX_Size() int {
	return m.Size()
}
func (m *DepositParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositParams.DiscardUnknown(m)
}

var xxx_messageInfo_DepositParams proto.InternalMessageInfo


type VotingParams struct {

	VotingPeriod time.Duration `protobuf:"bytes,1,opt,name=voting_period,json=votingPeriod,proto3,stdduration" json:"voting_period,omitempty" yaml:"voting_period"`
}

func (m *VotingParams) Reset()      { *m = VotingParams{} }
func (*VotingParams) ProtoMessage() {}
func (*VotingParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{7}
}
func (m *VotingParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VotingParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VotingParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VotingParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VotingParams.Merge(m, src)
}
func (m *VotingParams) XXX_Size() int {
	return m.Size()
}
func (m *VotingParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VotingParams.DiscardUnknown(m)
}

var xxx_messageInfo_VotingParams proto.InternalMessageInfo


type TallyParams struct {


	Quorum github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=quorum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quorum,omitempty"`

	Threshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=threshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"threshold,omitempty"`


	VetoThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=veto_threshold,json=vetoThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"veto_threshold,omitempty" yaml:"veto_threshold"`
}

func (m *TallyParams) Reset()      { *m = TallyParams{} }
func (*TallyParams) ProtoMessage() {}
func (*TallyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e82113c1a9a4b7c, []int{8}
}
func (m *TallyParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TallyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TallyParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TallyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TallyParams.Merge(m, src)
}
func (m *TallyParams) XXX_Size() int {
	return m.Size()
}
func (m *TallyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_TallyParams.DiscardUnknown(m)
}

var xxx_messageInfo_TallyParams proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("cosmos.gov.v1beta1.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterEnum("cosmos.gov.v1beta1.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	proto.RegisterType((*WeightedVoteOption)(nil), "cosmos.gov.v1beta1.WeightedVoteOption")
	proto.RegisterType((*TextProposal)(nil), "cosmos.gov.v1beta1.TextProposal")
	proto.RegisterType((*Deposit)(nil), "cosmos.gov.v1beta1.Deposit")
	proto.RegisterType((*Proposal)(nil), "cosmos.gov.v1beta1.Proposal")
	proto.RegisterType((*TallyResult)(nil), "cosmos.gov.v1beta1.TallyResult")
	proto.RegisterType((*Vote)(nil), "cosmos.gov.v1beta1.Vote")
	proto.RegisterType((*DepositParams)(nil), "cosmos.gov.v1beta1.DepositParams")
	proto.RegisterType((*VotingParams)(nil), "cosmos.gov.v1beta1.VotingParams")
	proto.RegisterType((*TallyParams)(nil), "cosmos.gov.v1beta1.TallyParams")
}

func init() { proto.RegisterFile("cosmos/gov/v1beta1/gov.proto", fileDescriptor_6e82113c1a9a4b7c) }

var fileDescriptor_6e82113c1a9a4b7c = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x51, 0x68, 0xdb, 0x56,
	0x17, 0xb6, 0x6c, 0xc7, 0x89, 0xaf, 0x9d, 0x44, 0xbd, 0x49, 0x13, 0xc7, 0x7f, 0x7f, 0xc9, 0xbf,
	0xfe, 0x9f, 0x12, 0x4a, 0xeb, 0xb4, 0xf9, 0xc7, 0xc6, 0x52, 0xd8, 0x66, 0xc5, 0xca, 0xea, 0x51,
	0x6c, 0x23, 0xab, 0x0e, 0xed, 0x1e, 0x84, 0x62, 0xdf, 0x3a, 0xda, 0x2c, 0x5d, 0xcf, 0xba, 0x4e,
	0x13, 0xf6, 0xb2, 0xc7, 0xe2, 0xc1, 0xe8, 0x63, 0x61, 0x18, 0x0a, 0x63, 0x2f, 0x7b, 0xde, 0xf3,
	0x9e, 0xc3, 0x18, 0xac, 0xec, 0xa9, 0x6c, 0xe0, 0xae, 0x09, 0x8c, 0x92, 0xc7, 0x3c, 0xec, 0x79,
	0x48, 0xf7, 0x2a, 0x96, 0xed, 0xb0, 0xd4, 0x7b, 0x8a, 0x74, 0xee, 0xf9, 0xbe, 0x73, 0xce, 0xe7,
	0x73, 0xce, 0x55, 0xc0, 0x95, 0x1a, 0x76, 0x2c, 0xec, 0xac, 0x35, 0xf0, 0xde, 0xda, 0xde, 0xad,
	0x1d, 0x44, 0x8c, 0x5b, 0xee, 0x73, 0xb6, 0xd5, 0xc6, 0x04, 0x43, 0x48, 0x4f, 0xb3, 0xae, 0x85,
	0x9d, 0xa6, 0x05, 0x86, 0xd8, 0x31, 0x1c, 0x74, 0x06, 0xa9, 0x61, 0xd3, 0xa6, 0x98, 0xf4, 0x62,
	0x03, 0x37, 0xb0, 0xf7, 0xb8, 0xe6, 0x3e, 0x31, 0xeb, 0x0a, 0x45, 0xe9, 0xf4, 0x80, 0xd1, 0xd2,
	0x23, 0xb1, 0x81, 0x71, 0xa3, 0x89, 0xd6, 0xbc, 0xb7, 0x9d, 0xce, 0xc3, 0x35, 0x62, 0x5a, 0xc8,
	0x21, 0x86, 0xd5, 0xf2, 0xb1, 0xa3, 0x0e, 0x86, 0x7d, 0xc0, 0x8e, 0x84, 0xd1, 0xa3, 0x7a, 0xa7,
	0x6d, 0x10, 0x13, 0xb3, 0x64, 0xa4, 0x6f, 0x39, 0x00, 0xb7, 0x91, 0xd9, 0xd8, 0x25, 0xa8, 0x5e,
	0xc5, 0x04, 0x95, 0x5a, 0xee, 0x21, 0x7c, 0x1b, 0xc4, 0xb0, 0xf7, 0x94, 0xe2, 0x32, 0xdc, 0xea,
	0xdc, 0xba, 0x90, 0x1d, 0x2f, 0x34, 0x3b, 0xf0, 0x57, 0x99, 0x37, 0xdc, 0x06, 0xb1, 0x47, 0x1e,
	0x5b, 0x2a, 0x9c, 0xe1, 0x56, 0xe3, 0xf2, 0xfb, 0x87, 0x7d, 0x31, 0xf4, 0x6b, 0x5f, 0xbc, 0xda,
	0x30, 0xc9, 0x6e, 0x67, 0x27, 0x5b, 0xc3, 0x16, 0xab, 0x8d, 0xfd, 0xb9, 0xe1, 0xd4, 0x3f, 0x5d,
	0x23, 0x07, 0x2d, 0xe4, 0x64, 0xf3, 0xa8, 0x76, 0xda, 0x17, 0x67, 0x0f, 0x0c, 0xab, 0xb9, 0x21,
	0x51, 0x16, 0x49, 0x65, 0x74, 0xd2, 0x36, 0x48, 0x6a, 0x68, 0x9f, 0x94, 0xdb, 0xb8, 0x85, 0x1d,
	0xa3, 0x09, 0x17, 0xc1, 0x14, 0x31, 0x49, 0x13, 0x79, 0xf9, 0xc5, 0x55, 0xfa, 0x02, 0x33, 0x20,
	0x51, 0x47, 0x4e, 0xad, 0x6d, 0xd2, 0xdc, 0xbd, 0x1c, 0xd4, 0xa0, 0x69, 0x63, 0xfe, 0xf5, 0x33,
	0x91, 0xfb, 0xe5, 0xfb, 0x1b, 0xd3, 0x9b, 0xd8, 0x26, 0xc8, 0x26, 0xd2, 0xcf, 0x1c, 0x98, 0xce,
	0xa3, 0x16, 0x76, 0x4c, 0x02, 0xdf, 0x01, 0x89, 0x16, 0x0b, 0xa0, 0x9b, 0x75, 0x8f, 0x3a, 0x2a,
	0x2f, 0x9d, 0xf6, 0x45, 0x48, 0x93, 0x0a, 0x1c, 0x4a, 0x2a, 0xf0, 0xdf, 0x0a, 0x75, 0x78, 0x05,
	0xc4, 0xeb, 0x94, 0x03, 0xb7, 0x59, 0xd4, 0x81, 0x01, 0xd6, 0x40, 0xcc, 0xb0, 0x70, 0xc7, 0x26,
	0xa9, 0x48, 0x26, 0xb2, 0x9a, 0x58, 0x5f, 0xf1, 0xc5, 0x74, 0x3b, 0xe4, 0x4c, 0xcd, 0x4d, 0x6c,
	0xda, 0xf2, 0x4d, 0x57, 0xaf, 0xef, 0x5e, 0x8a, 0xab, 0x6f, 0xa0, 0x97, 0x0b, 0x70, 0x54, 0x46,
	0xbd, 0x31, 0xf3, 0xf8, 0x99, 0x18, 0x7a, 0xfd, 0x4c, 0x0c, 0x49, 0x7f, 0xc6, 0xc0, 0xcc, 0x99,
	0x4e, 0x6f, 0x9d, 0x57, 0xd2, 0xc2, 0x49, 0x5f, 0x0c, 0x9b, 0xf5, 0xd3, 0xbe, 0x18, 0xa7, 0x85,
	0x8d, 0xd6, 0x73, 0x1b, 0x4c, 0xd7, 0xa8, 0x3e, 0x5e, 0x35, 0x89, 0xf5, 0xc5, 0x2c, 0xed, 0xa3,
	0xac, 0xdf, 0x47, 0xd9, 0x9c, 0x7d, 0x20, 0x27, 0x7e, 0x1c, 0x08, 0xa9, 0xfa, 0x08, 0x58, 0x05,
	0x31, 0x87, 0x18, 0xa4, 0xe3, 0xa4, 0x22, 0x5e, 0xef, 0x48, 0xe7, 0xf5, 0x8e, 0x9f, 0x60, 0xc5,
	0xf3, 0x94, 0xd3, 0xa7, 0x7d, 0x71, 0x69, 0x44, 0x64, 0x4a, 0x22, 0xa9, 0x8c, 0x0d, 0xb6, 0x00,
	0x7c, 0x68, 0xda, 0x46, 0x53, 0x27, 0x46, 0xb3, 0x79, 0xa0, 0xb7, 0x91, 0xd3, 0x69, 0x92, 0x54,
	0xd4, 0xcb, 0x4f, 0x3c, 0x2f, 0x86, 0xe6, 0xfa, 0xa9, 0x9e, 0x9b, 0xfc, 0x1f, 0x57, 0xd8, 0xd3,
	0xbe, 0xb8, 0x42, 0x83, 0x8c, 0x13, 0x49, 0x2a, 0xef, 0x19, 0x03, 0x20, 0xf8, 0x31, 0x48, 0x38,
	0x9d, 0x1d, 0xcb, 0x24, 0xba, 0x3b, 0x71, 0xa9, 0x29, 0x2f, 0x54, 0x7a, 0x4c, 0x0a, 0xcd, 0x1f,
	0x47, 0x59, 0x60, 0x51, 0x58, 0xbf, 0x04, 0xc0, 0xd2, 0x93, 0x97, 0x22, 0xa7, 0x02, 0x6a, 0x71,
	0x01, 0xd0, 0x04, 0x3c, 0x6b, 0x11, 0x1d, 0xd9, 0x75, 0x1a, 0x21, 0x76, 0x61, 0x84, 0xff, 0xb2,
	0x08, 0xcb, 0x34, 0xc2, 0x28, 0x03, 0x0d, 0x33, 0xc7, 0xcc, 0x8a, 0x5d, 0xf7, 0x42, 0x3d, 0xe6,
	0xc0, 0x2c, 0xc1, 0xc4, 0x68, 0xea, 0xec, 0x20, 0x35, 0x7d, 0x51, 0x23, 0xde, 0x61, 0x71, 0x16,
	0x69, 0x9c, 0x21, 0xb4, 0x34, 0x51, 0x83, 0x26, 0x3d, 0xac, 0x3f, 0x62, 0x4d, 0x70, 0x69, 0x0f,
	0x13, 0xd3, 0x6e, 0xb8, 0x3f, 0x6f, 0x9b, 0x09, 0x3b, 0x73, 0x61, 0xd9, 0xff, 0x63, 0xe9, 0xa4,
	0x68, 0x3a, 0x63, 0x14, 0xb4, 0xee, 0x79, 0x6a, 0xaf, 0xb8, 0x66, 0xaf, 0xf0, 0x87, 0x80, 0x99,
	0x06, 0x12, 0xc7, 0x2f, 0x8c, 0x25, 0xb1, 0x58, 0x4b, 0x43, 0xb1, 0x86, 0x15, 0x9e, 0xa5, 0x56,
	0x26, 0xf0, 0x46, 0xd4, 0xdd, 0x2a, 0xd2, 0x61, 0x18, 0x24, 0x82, 0xed, 0xf3, 0x01, 0x88, 0x1c,
	0x20, 0x87, 0x6e, 0x28, 0x39, 0x3b, 0xc1, 0x26, 0x2c, 0xd8, 0x44, 0x75, 0xa1, 0xf0, 0x0e, 0x98,
	0x36, 0x76, 0x1c, 0x62, 0x98, 0x6c, 0x97, 0x4d, 0xcc, 0xe2, 0xc3, 0xe1, 0x7b, 0x20, 0x6c, 0x63,
	0x6f, 0x20, 0x27, 0x27, 0x09, 0xdb, 0x18, 0x36, 0x40, 0xd2, 0xc6, 0xfa, 0x23, 0x93, 0xec, 0xea,
	0x7b, 0x88, 0x60, 0x6f, 0xec, 0xe2, 0xb2, 0x32, 0x19, 0xd3, 0x69, 0x5f, 0x5c, 0xa0, 0xa2, 0x06,
	0xb9, 0x24, 0x15, 0xd8, 0x78, 0xdb, 0x24, 0xbb, 0x55, 0x44, 0x30, 0x93, 0xf2, 0x98, 0x03, 0x51,
	0xf7, 0x7a, 0xf9, 0xe7, 0x2b, 0x79, 0x11, 0x4c, 0xed, 0x61, 0x82, 0xfc, 0x75, 0x4c, 0x5f, 0xe0,
	0xc6, 0xd9, 0xbd, 0x16, 0x79, 0x93, 0x7b, 0x4d, 0x0e, 0xa7, 0xb8, 0xb3, 0xbb, 0x6d, 0x0b, 0x4c,
	0xd3, 0x27, 0x27, 0x15, 0xf5, 0xc6, 0xe7, 0xea, 0x79, 0xe0, 0xf1, 0xcb, 0x54, 0x8e, 0xba, 0x2a,
	0xa9, 0x3e, 0x78, 0x63, 0xe6, 0xa9, 0xbf, 0xa9, 0x7f, 0x08, 0x83, 0x59, 0x36, 0x18, 0x65, 0xa3,
	0x6d, 0x58, 0x0e, 0xfc, 0x9a, 0x03, 0x09, 0xcb, 0xb4, 0xcf, 0xe6, 0x94, 0xbb, 0x68, 0x4e, 0x75,
	0x97, 0xfb, 0xa4, 0x2f, 0x5e, 0x0e, 0xa0, 0xae, 0x63, 0xcb, 0x24, 0xc8, 0x6a, 0x91, 0x83, 0x81,
	0x4e, 0x81, 0xe3, 0xc9, 0xc6, 0x17, 0x58, 0xa6, 0xed, 0x0f, 0xef, 0x57, 0x1c, 0x80, 0x96, 0xb1,
	0xef, 0x13, 0xe9, 0x2d, 0xd4, 0x36, 0x71, 0x9d, 0x5d, 0x11, 0x2b, 0x63, 0x23, 0x95, 0x67, 0x9f,
	0x1a, 0xb4, 0x4d, 0x4e, 0xfa, 0xe2, 0x95, 0x71, 0xf0, 0x50, 0xae, 0x6c, 0x39, 0x8f, 0x7b, 0x49,
	0x4f, 0xdd, 0xa1, 0xe3, 0x2d, 0x63, 0xdf, 0x97, 0x8b, 0x9a, 0xbf, 0xe4, 0x40, 0xb2, 0xea, 0x4d,
	0x22, 0xd3, 0xef, 0x73, 0xc0, 0x26, 0xd3, 0xcf, 0x8d, 0xbb, 0x28, 0xb7, 0xdb, 0x2c, 0xb7, 0xe5,
	0x21, 0xdc, 0x50, 0x5a, 0x8b, 0x43, 0x8b, 0x20, 0x98, 0x51, 0x92, 0xda, 0x58, 0x36, 0xbf, 0xf9,
	0xf3, 0xcf, 0x92, 0x79, 0x00, 0x62, 0x9f, 0x75, 0x70, 0xbb, 0x63, 0x79, 0x59, 0x24, 0x65, 0x79,
	0xb2, 0x8f, 0xa1, 0x93, 0xbe, 0xc8, 0x53, 0xfc, 0x20, 0x1b, 0x95, 0x31, 0xc2, 0x1a, 0x88, 0x93,
	0xdd, 0x36, 0x72, 0x76, 0x71, 0x93, 0xfe, 0x00, 0xc9, 0x89, 0x86, 0x91, 0xd2, 0x2f, 0x9c, 0x51,
	0x04, 0x22, 0x0c, 0x78, 0x61, 0x97, 0x03, 0x73, 0xee, 0x84, 0xea, 0x83, 0x50, 0x11, 0x2f, 0x54,
	0x6d, 0xe2, 0x50, 0xa9, 0x61, 0x9e, 0x21, 0x7d, 0x2f, 0x33, 0x7d, 0x87, 0x3c, 0x24, 0x75, 0xd6,
	0x35, 0x68, 0xfe, 0xfb, 0xb5, 0x3f, 0x38, 0x00, 0x02, 0x5f, 0xa8, 0xd7, 0xc1, 0x72, 0xb5, 0xa4,
	0x29, 0x7a, 0xa9, 0xac, 0x15, 0x4a, 0x45, 0xfd, 0x5e, 0xb1, 0x52, 0x56, 0x36, 0x0b, 0x5b, 0x05,
	0x25, 0xcf, 0x87, 0xd2, 0xf3, 0xdd, 0x5e, 0x26, 0x41, 0x1d, 0x15, 0x37, 0x08, 0x94, 0xc0, 0x7c,
	0xd0, 0xfb, 0xbe, 0x52, 0xe1, 0xb9, 0xf4, 0x6c, 0xb7, 0x97, 0x89, 0x53, 0xaf, 0xfb, 0xc8, 0x81,
	0xd7, 0xc0, 0x42, 0xd0, 0x27, 0x27, 0x57, 0xb4, 0x5c, 0xa1, 0xc8, 0x87, 0xd3, 0x97, 0xba, 0xbd,
	0xcc, 0x2c, 0xf5, 0xcb, 0xb1, 0x75, 0x9a, 0x01, 0x73, 0x41, 0xdf, 0x62, 0x89, 0x8f, 0xa4, 0x93,
	0xdd, 0x5e, 0x66, 0x86, 0xba, 0x15, 0x31, 0x5c, 0x07, 0xa9, 0x61, 0x0f, 0x7d, 0xbb, 0xa0, 0xdd,
	0xd1, 0xab, 0x8a, 0x56, 0xe2, 0xa3, 0xe9, 0xc5, 0x6e, 0x2f, 0xc3, 0xfb, 0xbe, 0xfe, 0xee, 0x4b,
	0x47, 0x1f, 0x7f, 0x23, 0x84, 0xae, 0xfd, 0x14, 0x06, 0x73, 0xc3, 0x9f, 0x47, 0x30, 0x0b, 0xfe,
	0x55, 0x56, 0x4b, 0xe5, 0x52, 0x25, 0x77, 0x57, 0xaf, 0x68, 0x39, 0xed, 0x5e, 0x65, 0xa4, 0x60,
	0xaf, 0x14, 0xea, 0x5c, 0x34, 0x9b, 0xf0, 0x36, 0x10, 0x46, 0xfd, 0xf3, 0x4a, 0xb9, 0x54, 0x29,
	0x68, 0x7a, 0x59, 0x51, 0x0b, 0xa5, 0x3c, 0xcf, 0xa5, 0x97, 0xbb, 0xbd, 0xcc, 0x02, 0x85, 0x0c,
	0x0d, 0x15, 0x7c, 0x17, 0xfc, 0x7b, 0x14, 0x5c, 0x2d, 0x69, 0x85, 0xe2, 0x87, 0x3e, 0x36, 0x9c,
	0x5e, 0xea, 0xf6, 0x32, 0x90, 0x62, 0xab, 0x81, 0x09, 0x80, 0xd7, 0xc1, 0xd2, 0x28, 0xb4, 0x9c,
	0xab, 0x54, 0x94, 0x3c, 0x1f, 0x49, 0xf3, 0xdd, 0x5e, 0x26, 0x49, 0x31, 0x65, 0xc3, 0x71, 0x50,
	0x1d, 0xde, 0x04, 0xa9, 0x51, 0x6f, 0x55, 0xf9, 0x48, 0xd9, 0xd4, 0x94, 0x3c, 0x1f, 0x4d, 0xc3,
	0x6e, 0x2f, 0x33, 0x47, 0xfd, 0x55, 0xf4, 0x09, 0xaa, 0x11, 0x74, 0x2e, 0xff, 0x56, 0xae, 0x70,
	0x57, 0xc9, 0xf3, 0x53, 0x41, 0xfe, 0x2d, 0xc3, 0x6c, 0xa2, 0x3a, 0x95, 0x53, 0x2e, 0x1e, 0xbe,
	0x12, 0x42, 0x2f, 0x5e, 0x09, 0xa1, 0x2f, 0x8e, 0x84, 0xd0, 0xe1, 0x91, 0xc0, 0x3d, 0x3f, 0x12,
	0xb8, 0xdf, 0x8f, 0x04, 0xee, 0xc9, 0xb1, 0x10, 0x7a, 0x7e, 0x2c, 0x84, 0x5e, 0x1c, 0x0b, 0xa1,
	0x07, 0x7f, 0xbf, 0x10, 0xf7, 0xbd, 0x7f, 0xff, 0xbc, 0x7e, 0xde, 0x89, 0x79, 0x3b, 0xe4, 0xff,
	0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x14, 0xb4, 0xf2, 0xfe, 0x19, 0x0e, 0x00, 0x00,
}

func (this *TextProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TextProposal)
	if !ok {
		that2, ok := that.(TextProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	return true
}
func (this *Proposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Proposal)
	if !ok {
		that2, ok := that.(Proposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ProposalId != that1.ProposalId {
		return false
	}
	if !this.Content.Equal(that1.Content) {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if !this.FinalTallyResult.Equal(&that1.FinalTallyResult) {
		return false
	}
	if !this.SubmitTime.Equal(that1.SubmitTime) {
		return false
	}
	if !this.DepositEndTime.Equal(that1.DepositEndTime) {
		return false
	}
	if len(this.TotalDeposit) != len(that1.TotalDeposit) {
		return false
	}
	for i := range this.TotalDeposit {
		if !this.TotalDeposit[i].Equal(&that1.TotalDeposit[i]) {
			return false
		}
	}
	if !this.VotingStartTime.Equal(that1.VotingStartTime) {
		return false
	}
	if !this.VotingEndTime.Equal(that1.VotingEndTime) {
		return false
	}
	return true
}
func (this *TallyResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TallyResult)
	if !ok {
		that2, ok := that.(TallyResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Yes.Equal(that1.Yes) {
		return false
	}
	if !this.Abstain.Equal(that1.Abstain) {
		return false
	}
	if !this.No.Equal(that1.No) {
		return false
	}
	if !this.NoWithVeto.Equal(that1.NoWithVeto) {
		return false
	}
	return true
}
func (m *WeightedVoteOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedVoteOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedVoteOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Option != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TextProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TextProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TextProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Deposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Deposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Deposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VotingEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingEndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGov(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x4a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.VotingStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingStartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGov(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x42
	if len(m.TotalDeposit) > 0 {
		for iNdEx := len(m.TotalDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.DepositEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositEndTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGov(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x32
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SubmitTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmitTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGov(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.FinalTallyResult.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Status != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGov(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TallyResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TallyResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TallyResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NoWithVeto.Size()
		i -= size
		if _, err := m.NoWithVeto.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.No.Size()
		i -= size
		if _, err := m.No.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Abstain.Size()
		i -= size
		if _, err := m.Abstain.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Yes.Size()
		i -= size
		if _, err := m.Yes.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Options[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Option != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintGov(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DepositParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n7, err7 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MaxDepositPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxDepositPeriod):])
	if err7 != nil {
		return 0, err7
	}
	i -= n7
	i = encodeVarintGov(dAtA, i, uint64(n7))
	i--
	dAtA[i] = 0x12
	if len(m.MinDeposit) > 0 {
		for iNdEx := len(m.MinDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VotingParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VotingParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VotingParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n8, err8 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.VotingPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.VotingPeriod):])
	if err8 != nil {
		return 0, err8
	}
	i -= n8
	i = encodeVarintGov(dAtA, i, uint64(n8))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TallyParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TallyParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TallyParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.VetoThreshold.Size()
		i -= size
		if _, err := m.VetoThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Threshold.Size()
		i -= size
		if _, err := m.Threshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Quorum.Size()
		i -= size
		if _, err := m.Quorum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WeightedVoteOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Option != 0 {
		n += 1 + sovGov(uint64(m.Option))
	}
	l = m.Weight.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *TextProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func (m *Deposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovGov(uint64(m.ProposalId))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovGov(uint64(m.ProposalId))
	}
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovGov(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGov(uint64(m.Status))
	}
	l = m.FinalTallyResult.Size()
	n += 1 + l + sovGov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmitTime)
	n += 1 + l + sovGov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.DepositEndTime)
	n += 1 + l + sovGov(uint64(l))
	if len(m.TotalDeposit) > 0 {
		for _, e := range m.TotalDeposit {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingStartTime)
	n += 1 + l + sovGov(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.VotingEndTime)
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *TallyResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Yes.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Abstain.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.No.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.NoWithVeto.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovGov(uint64(m.ProposalId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.Option != 0 {
		n += 1 + sovGov(uint64(m.Option))
	}
	if len(m.Options) > 0 {
		for _, e := range m.Options {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	return n
}

func (m *DepositParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinDeposit) > 0 {
		for _, e := range m.MinDeposit {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MaxDepositPeriod)
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *VotingParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.VotingPeriod)
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *TallyParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Quorum.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.Threshold.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.VetoThreshold.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WeightedVoteOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: WeightedVoteOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedVoteOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *TextProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: TextProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TextProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Deposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Deposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Deposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types1.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalTallyResult", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FinalTallyResult.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SubmitTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.DepositEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalDeposit = append(m.TotalDeposit, types.Coin{})
			if err := m.TotalDeposit[len(m.TotalDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VotingStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingEndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.VotingEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *TallyResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: TallyResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TallyResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Yes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Yes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abstain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Abstain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field No", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.No.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoWithVeto", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NoWithVeto.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = append(m.Options, WeightedVoteOption{})
			if err := m.Options[len(m.Options)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *DepositParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: DepositParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposit = append(m.MinDeposit, types.Coin{})
			if err := m.MinDeposit[len(m.MinDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDepositPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MaxDepositPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *VotingParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: VotingParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VotingParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.VotingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *TallyParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: TallyParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TallyParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quorum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Threshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VetoThreshold", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VetoThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
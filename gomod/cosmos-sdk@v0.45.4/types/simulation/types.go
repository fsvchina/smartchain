package simulation

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

type WeightedProposalContent interface {
	AppParamsKey() string
	DefaultWeight() int
	ContentSimulatorFn() ContentSimulatorFn
}

type ContentSimulatorFn func(r *rand.Rand, ctx sdk.Context, accs []Account) Content

type Content interface {
	GetTitle() string
	GetDescription() string
	ProposalRoute() string
	ProposalType() string
	ValidateBasic() error
	String() string
}

type SimValFn func(r *rand.Rand) string

type ParamChange interface {
	Subspace() string
	Key() string
	SimValue() SimValFn
	ComposedKey() string
}

type WeightedOperation interface {
	Weight() int
	Op() Operation
}




//


//


type Operation func(r *rand.Rand, app *baseapp.BaseApp,
	ctx sdk.Context, accounts []Account, chainID string) (
	OperationMsg OperationMsg, futureOps []FutureOperation, err error)


type OperationMsg struct {
	Route   string          `json:"route" yaml:"route"`
	Name    string          `json:"name" yaml:"name"`
	Comment string          `json:"comment" yaml:"comment"`
	OK      bool            `json:"ok" yaml:"ok"`
	Msg     json.RawMessage `json:"msg" yaml:"msg"`
}


func NewOperationMsgBasic(route, name, comment string, ok bool, msg []byte) OperationMsg {
	return OperationMsg{
		Route:   route,
		Name:    name,
		Comment: comment,
		OK:      ok,
		Msg:     msg,
	}
}


func NewOperationMsg(msg sdk.Msg, ok bool, comment string, cdc *codec.ProtoCodec) OperationMsg {
	if legacyMsg, okType := msg.(legacytx.LegacyMsg); okType {
		return NewOperationMsgBasic(legacyMsg.Route(), legacyMsg.Type(), comment, ok, legacyMsg.GetSignBytes())
	}

	bz := cdc.MustMarshalJSON(msg)

	return NewOperationMsgBasic(sdk.MsgTypeURL(msg), sdk.MsgTypeURL(msg), comment, ok, bz)

}


func NoOpMsg(route, msgType, comment string) OperationMsg {
	return NewOperationMsgBasic(route, msgType, comment, false, nil)
}


func (om OperationMsg) String() string {
	out, err := json.Marshal(om)
	if err != nil {
		panic(err)
	}

	return string(out)
}


func (om OperationMsg) MustMarshal() json.RawMessage {
	out, err := json.Marshal(om)
	if err != nil {
		panic(err)
	}

	return out
}


func (om OperationMsg) LogEvent(eventLogger func(route, op, evResult string)) {
	pass := "ok"
	if !om.OK {
		pass = "failure"
	}

	eventLogger(om.Route, om.Name, pass)
}





type FutureOperation struct {
	BlockHeight int
	BlockTime   time.Time
	Op          Operation
}




type AppParams map[string]json.RawMessage





func (sp AppParams) GetOrGenerate(_ codec.JSONCodec, key string, ptr interface{}, r *rand.Rand, ps ParamSimulator) {
	if v, ok := sp[key]; ok && v != nil {
		err := json.Unmarshal(v, ptr)
		if err != nil {
			panic(err)
		}
		return
	}

	ps(r)
}

type ParamSimulator func(r *rand.Rand)

type SelectOpFn func(r *rand.Rand) Operation


type AppStateFn func(r *rand.Rand, accs []Account, config Config) (
	appState json.RawMessage, accounts []Account, chainId string, genesisTimestamp time.Time,
)


type RandomAccountFn func(r *rand.Rand, n int) []Account

type Params interface {
	PastEvidenceFraction() float64
	NumKeys() int
	EvidenceFraction() float64
	InitialLivenessWeightings() []int
	LivenessTransitionMatrix() TransitionMatrix
	BlockSizeTransitionMatrix() TransitionMatrix
}

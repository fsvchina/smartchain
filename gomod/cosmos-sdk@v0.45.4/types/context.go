package types

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/store/gaskv"
	stypes "github.com/cosmos/cosmos-sdk/store/types"
)


type Context struct {
	ctx           context.Context
	ms            MultiStore
	header        tmproto.Header
	headerHash    tmbytes.HexBytes
	chainID       string
	txBytes       []byte
	logger        log.Logger
	voteInfo      []abci.VoteInfo
	gasMeter      GasMeter
	blockGasMeter GasMeter
	checkTx       bool
	recheckTx     bool
	minGasPrice   DecCoins
	consParams    *abci.ConsensusParams
	eventManager  *EventManager
}


type Request = Context


func (c Context) Context() context.Context    { return c.ctx }
func (c Context) MultiStore() MultiStore      { return c.ms }
func (c Context) BlockHeight() int64          { return c.header.Height }
func (c Context) BlockTime() time.Time        { return c.header.Time }
func (c Context) ChainID() string             { return c.chainID }
func (c Context) TxBytes() []byte             { return c.txBytes }
func (c Context) Logger() log.Logger          { return c.logger }
func (c Context) VoteInfos() []abci.VoteInfo  { return c.voteInfo }
func (c Context) GasMeter() GasMeter          { return c.gasMeter }
func (c Context) BlockGasMeter() GasMeter     { return c.blockGasMeter }
func (c Context) IsCheckTx() bool             { return c.checkTx }
func (c Context) IsReCheckTx() bool           { return c.recheckTx }
func (c Context) MinGasPrices() DecCoins      { return c.minGasPrice }
func (c Context) EventManager() *EventManager { return c.eventManager }


func (c Context) BlockHeader() tmproto.Header {
	var msg = proto.Clone(&c.header).(*tmproto.Header)
	return *msg
}


func (c Context) HeaderHash() tmbytes.HexBytes {
	hash := make([]byte, len(c.headerHash))
	copy(hash, c.headerHash)
	return hash
}

func (c Context) ConsensusParams() *abci.ConsensusParams {
	return proto.Clone(c.consParams).(*abci.ConsensusParams)
}


func NewContext(ms MultiStore, header tmproto.Header, isCheckTx bool, logger log.Logger) Context {

	header.Time = header.Time.UTC()
	return Context{
		ctx:          context.Background(),
		ms:           ms,
		header:       header,
		chainID:      header.ChainID,
		checkTx:      isCheckTx,
		logger:       logger,
		gasMeter:     stypes.NewInfiniteGasMeter(),
		minGasPrice:  DecCoins{},
		eventManager: NewEventManager(),
	}
}


func (c Context) WithContext(ctx context.Context) Context {
	c.ctx = ctx
	return c
}


func (c Context) WithMultiStore(ms MultiStore) Context {
	c.ms = ms
	return c
}


func (c Context) WithBlockHeader(header tmproto.Header) Context {

	header.Time = header.Time.UTC()
	c.header = header
	return c
}


func (c Context) WithHeaderHash(hash []byte) Context {
	temp := make([]byte, len(hash))
	copy(temp, hash)

	c.headerHash = temp
	return c
}


func (c Context) WithBlockTime(newTime time.Time) Context {
	newHeader := c.BlockHeader()

	newHeader.Time = newTime.UTC()
	return c.WithBlockHeader(newHeader)
}


func (c Context) WithProposer(addr ConsAddress) Context {
	newHeader := c.BlockHeader()
	newHeader.ProposerAddress = addr.Bytes()
	return c.WithBlockHeader(newHeader)
}


func (c Context) WithBlockHeight(height int64) Context {
	newHeader := c.BlockHeader()
	newHeader.Height = height
	return c.WithBlockHeader(newHeader)
}


func (c Context) WithChainID(chainID string) Context {
	c.chainID = chainID
	return c
}


func (c Context) WithTxBytes(txBytes []byte) Context {
	c.txBytes = txBytes
	return c
}


func (c Context) WithLogger(logger log.Logger) Context {
	c.logger = logger
	return c
}


func (c Context) WithVoteInfos(voteInfo []abci.VoteInfo) Context {
	c.voteInfo = voteInfo
	return c
}


func (c Context) WithGasMeter(meter GasMeter) Context {
	c.gasMeter = meter
	return c
}


func (c Context) WithBlockGasMeter(meter GasMeter) Context {
	c.blockGasMeter = meter
	return c
}


func (c Context) WithIsCheckTx(isCheckTx bool) Context {
	c.checkTx = isCheckTx
	return c
}



func (c Context) WithIsReCheckTx(isRecheckTx bool) Context {
	if isRecheckTx {
		c.checkTx = true
	}
	c.recheckTx = isRecheckTx
	return c
}


func (c Context) WithMinGasPrices(gasPrices DecCoins) Context {
	c.minGasPrice = gasPrices
	return c
}


func (c Context) WithConsensusParams(params *abci.ConsensusParams) Context {
	c.consParams = params
	return c
}


func (c Context) WithEventManager(em *EventManager) Context {
	c.eventManager = em
	return c
}


func (c Context) IsZero() bool {
	return c.ms == nil
}






func (c Context) WithValue(key, value interface{}) Context {
	c.ctx = context.WithValue(c.ctx, key, value)
	return c
}






func (c Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}






func (c Context) KVStore(key StoreKey) KVStore {
	return gaskv.NewStore(c.MultiStore().GetKVStore(key), c.GasMeter(), stypes.KVGasConfig())
}


func (c Context) TransientStore(key StoreKey) KVStore {
	return gaskv.NewStore(c.MultiStore().GetKVStore(key), c.GasMeter(), stypes.TransientGasConfig())
}




func (c Context) CacheContext() (cc Context, writeCache func()) {
	cms := c.MultiStore().CacheMultiStore()
	cc = c.WithMultiStore(cms).WithEventManager(NewEventManager())
	return cc, cms.Write
}


type ContextKey string


const SdkContextKey ContextKey = "sdk-context"





func WrapSDKContext(ctx Context) context.Context {
	return context.WithValue(ctx.ctx, SdkContextKey, ctx)
}




func UnwrapSDKContext(ctx context.Context) Context {
	return ctx.Value(SdkContextKey).(Context)
}

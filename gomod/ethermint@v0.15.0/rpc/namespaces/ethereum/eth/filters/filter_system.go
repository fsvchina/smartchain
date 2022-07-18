package filters

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"

	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmquery "github.com/tendermint/tendermint/libs/pubsub/query"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	rpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tharsis/ethermint/rpc/ethereum/pubsub"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

var (
	txEvents     = tmtypes.QueryForEvent(tmtypes.EventTx).String()
	evmEvents    = tmquery.MustParse(fmt.Sprintf("%s='%s' AND %s.%s='%s'", tmtypes.EventTypeKey, tmtypes.EventTx, sdk.EventTypeMessage, sdk.AttributeKeyModule, evmtypes.ModuleName)).String()
	headerEvents = tmtypes.QueryForEvent(tmtypes.EventNewBlockHeader).String()
)



type EventSystem struct {
	logger     log.Logger
	ctx        context.Context
	tmWSClient *rpcclient.WSClient


	lightMode bool

	index      filterIndex
	topicChans map[string]chan<- coretypes.ResultEvent
	indexMux   *sync.RWMutex


	install   chan *Subscription
	uninstall chan *Subscription
	eventBus  pubsub.EventBus
}




//


func NewEventSystem(logger log.Logger, tmWSClient *rpcclient.WSClient) *EventSystem {
	index := make(filterIndex)
	for i := filters.UnknownSubscription; i < filters.LastIndexSubscription; i++ {
		index[i] = make(map[rpc.ID]*Subscription)
	}

	es := &EventSystem{
		logger:     logger,
		ctx:        context.Background(),
		tmWSClient: tmWSClient,
		lightMode:  false,
		index:      index,
		topicChans: make(map[string]chan<- coretypes.ResultEvent, len(index)),
		indexMux:   new(sync.RWMutex),
		install:    make(chan *Subscription),
		uninstall:  make(chan *Subscription),
		eventBus:   pubsub.NewEventBus(),
	}

	go es.eventLoop()
	go es.consumeEvents()
	return es
}



func (es *EventSystem) WithContext(ctx context.Context) {
	es.ctx = ctx
}



func (es *EventSystem) subscribe(sub *Subscription) (*Subscription, pubsub.UnsubscribeFunc, error) {
	var (
		err      error
		cancelFn context.CancelFunc
	)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	existingSubs := es.eventBus.Topics()
	for _, topic := range existingSubs {
		if topic == sub.event {
			eventCh, unsubFn, err := es.eventBus.Subscribe(sub.event)
			if err != nil {
				err := errors.Wrapf(err, "failed to subscribe to topic: %s", sub.event)
				return nil, nil, err
			}

			sub.eventCh = eventCh
			return sub, unsubFn, nil
		}
	}

	switch sub.typ {
	case filters.LogsSubscription:
		err = es.tmWSClient.Subscribe(ctx, sub.event)
	case filters.BlocksSubscription:
		err = es.tmWSClient.Subscribe(ctx, sub.event)
	case filters.PendingTransactionsSubscription:
		err = es.tmWSClient.Subscribe(ctx, sub.event)
	default:
		err = fmt.Errorf("invalid filter subscription type %d", sub.typ)
	}

	if err != nil {
		sub.err <- err
		return nil, nil, err
	}


	es.install <- sub
	<-sub.installed

	eventCh, unsubFn, err := es.eventBus.Subscribe(sub.event)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to subscribe to topic after installed: %s", sub.event)
	}

	sub.eventCh = eventCh
	return sub, unsubFn, nil
}




func (es *EventSystem) SubscribeLogs(crit filters.FilterCriteria) (*Subscription, pubsub.UnsubscribeFunc, error) {
	var from, to rpc.BlockNumber
	if crit.FromBlock == nil {
		from = rpc.LatestBlockNumber
	} else {
		from = rpc.BlockNumber(crit.FromBlock.Int64())
	}
	if crit.ToBlock == nil {
		to = rpc.LatestBlockNumber
	} else {
		to = rpc.BlockNumber(crit.ToBlock.Int64())
	}

	switch {


	case (from == rpc.LatestBlockNumber && to == rpc.LatestBlockNumber),
		(from >= 0 && to >= 0 && to >= from),
		(from >= 0 && to == rpc.LatestBlockNumber):
		return es.subscribeLogs(crit)

	default:
		return nil, nil, fmt.Errorf("invalid from and to block combination: from > to (%d > %d)", from, to)
	}
}



func (es *EventSystem) subscribeLogs(crit filters.FilterCriteria) (*Subscription, pubsub.UnsubscribeFunc, error) {
	sub := &Subscription{
		id:        rpc.NewID(),
		typ:       filters.LogsSubscription,
		event:     evmEvents,
		logsCrit:  crit,
		created:   time.Now().UTC(),
		logs:      make(chan []*ethtypes.Log),
		installed: make(chan struct{}, 1),
		err:       make(chan error, 1),
	}
	return es.subscribe(sub)
}


func (es EventSystem) SubscribeNewHeads() (*Subscription, pubsub.UnsubscribeFunc, error) {
	sub := &Subscription{
		id:        rpc.NewID(),
		typ:       filters.BlocksSubscription,
		event:     headerEvents,
		created:   time.Now().UTC(),
		headers:   make(chan *ethtypes.Header),
		installed: make(chan struct{}, 1),
		err:       make(chan error, 1),
	}
	return es.subscribe(sub)
}


func (es EventSystem) SubscribePendingTxs() (*Subscription, pubsub.UnsubscribeFunc, error) {
	sub := &Subscription{
		id:        rpc.NewID(),
		typ:       filters.PendingTransactionsSubscription,
		event:     txEvents,
		created:   time.Now().UTC(),
		hashes:    make(chan []common.Hash),
		installed: make(chan struct{}, 1),
		err:       make(chan error, 1),
	}
	return es.subscribe(sub)
}

type filterIndex map[filters.Type]map[rpc.ID]*Subscription


func (es *EventSystem) eventLoop() {
	for {
		select {
		case f := <-es.install:
			es.indexMux.Lock()
			es.index[f.typ][f.id] = f
			ch := make(chan coretypes.ResultEvent)
			es.topicChans[f.event] = ch
			if err := es.eventBus.AddTopic(f.event, ch); err != nil {
				es.logger.Error("failed to add event topic to event bus", "topic", f.event, "error", err.Error())
			}
			es.indexMux.Unlock()
			close(f.installed)
		case f := <-es.uninstall:
			es.indexMux.Lock()
			delete(es.index[f.typ], f.id)

			var channelInUse bool
			for _, sub := range es.index[f.typ] {
				if sub.event == f.event {
					channelInUse = true
					break
				}
			}


			if !channelInUse {
				if err := es.tmWSClient.Unsubscribe(es.ctx, f.event); err != nil {
					es.logger.Error("failed to unsubscribe from query", "query", f.event, "error", err.Error())
				}

				ch, ok := es.topicChans[f.event]
				if ok {
					es.eventBus.RemoveTopic(f.event)
					close(ch)
					delete(es.topicChans, f.event)
				}
			}

			es.indexMux.Unlock()
			close(f.err)
		}
	}
}

func (es *EventSystem) consumeEvents() {
	for {
		for rpcResp := range es.tmWSClient.ResponsesCh {
			var ev coretypes.ResultEvent

			if rpcResp.Error != nil {
				time.Sleep(5 * time.Second)
				continue
			} else if err := tmjson.Unmarshal(rpcResp.Result, &ev); err != nil {
				es.logger.Error("failed to JSON unmarshal ResponsesCh result event", "error", err.Error())
				continue
			}

			if len(ev.Query) == 0 {

				continue
			}

			es.indexMux.RLock()
			ch, ok := es.topicChans[ev.Query]
			es.indexMux.RUnlock()
			if !ok {
				es.logger.Debug("channel for subscription not found", "topic", ev.Query)
				es.logger.Debug("list of available channels", "channels", es.eventBus.Topics())
				continue
			}


			t := time.NewTimer(time.Second)
			select {
			case <-t.C:
				es.logger.Debug("dropped event during lagging subscription", "topic", ev.Query)
			case ch <- ev:
			}
		}

		time.Sleep(time.Second)
	}
}

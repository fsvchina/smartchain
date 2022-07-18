package filters

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)


type Subscription struct {
	id        rpc.ID
	typ       filters.Type
	event     string
	created   time.Time
	logsCrit  filters.FilterCriteria
	logs      chan []*ethtypes.Log
	hashes    chan []common.Hash
	headers   chan *ethtypes.Header
	installed chan struct{}
	eventCh   <-chan coretypes.ResultEvent
	err       chan error
}


func (s Subscription) ID() rpc.ID {
	return s.id
}



func (s *Subscription) Unsubscribe(es *EventSystem) {
	go func() {
	uninstallLoop:
		for {




			select {
			case es.uninstall <- s:
				break uninstallLoop
			case <-s.logs:
			case <-s.hashes:
			case <-s.headers:
			}
		}
	}()
}


func (s *Subscription) Err() <-chan error {
	return s.err
}


func (s *Subscription) Event() <-chan coretypes.ResultEvent {
	return s.eventCh
}

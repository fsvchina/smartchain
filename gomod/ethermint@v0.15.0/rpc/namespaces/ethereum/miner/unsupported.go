package miner

import (
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)



func (api *API) GetHashrate() uint64 {
	api.logger.Debug("miner_getHashrate")
	api.logger.Debug("Unsupported rpc function: miner_getHashrate")
	return 0
}



func (api *API) SetExtra(extra string) (bool, error) {
	api.logger.Debug("miner_setExtra")
	api.logger.Debug("Unsupported rpc function: miner_setExtra")
	return false, errors.New("unsupported rpc function: miner_setExtra")
}



func (api *API) SetGasLimit(gasLimit hexutil.Uint64) bool {
	api.logger.Debug("miner_setGasLimit")
	api.logger.Debug("Unsupported rpc function: miner_setGasLimit")
	return false
}







func (api *API) Start(threads *int) error {
	api.logger.Debug("miner_start")
	api.logger.Debug("Unsupported rpc function: miner_start")
	return errors.New("unsupported rpc function: miner_start")
}




func (api *API) Stop() {
	api.logger.Debug("miner_stop")
	api.logger.Debug("Unsupported rpc function: miner_stop")
}

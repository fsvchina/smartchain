package config

import "fs.video/smartchain/core"

const ClientToml = `
chain-id = "` + core.ChainID + `"
# The keyring's backend, where the keys are stored (os|file|kwallet|pass|test|memory)
keyring-backend = "os"
# CLI output format (text|json)
output = "text"
# <host>:<port> to Tendermint RPC interface for this chain
node = "tcp:
# Transaction broadcasting mode (sync|async|block)
broadcast-mode = "sync"
`

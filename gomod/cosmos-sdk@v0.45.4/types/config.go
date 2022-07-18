package types

import (
	"context"
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/version"
)


const DefaultKeyringServiceName = "cosmos"



type Config struct {
	fullFundraiserPath  string
	bech32AddressPrefix map[string]string
	txEncoder           TxEncoder
	addressVerifier     func([]byte) error
	mtx                 sync.RWMutex


	purpose  uint32
	coinType uint32

	sealed   bool
	sealedch chan struct{}
}


var (
	sdkConfig  *Config
	initConfig sync.Once
)


func NewConfig() *Config {
	return &Config{
		sealedch: make(chan struct{}),
		bech32AddressPrefix: map[string]string{
			"account_addr":   Bech32PrefixAccAddr,
			"validator_addr": Bech32PrefixValAddr,
			"consensus_addr": Bech32PrefixConsAddr,
			"account_pub":    Bech32PrefixAccPub,
			"validator_pub":  Bech32PrefixValPub,
			"consensus_pub":  Bech32PrefixConsPub,
		},
		fullFundraiserPath: FullFundraiserPath,

		purpose:   Purpose,
		coinType:  CoinType,
		txEncoder: nil,
	}
}


func GetConfig() *Config {
	initConfig.Do(func() {
		sdkConfig = NewConfig()
	})
	return sdkConfig
}


func GetSealedConfig(ctx context.Context) (*Config, error) {
	config := GetConfig()
	select {
	case <-config.sealedch:
		return config, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (config *Config) assertNotSealed() {
	config.mtx.Lock()
	defer config.mtx.Unlock()

	if config.sealed {
		panic("Config is sealed")
	}
}



func (config *Config) SetBech32PrefixForAccount(addressPrefix, pubKeyPrefix string) {
	config.assertNotSealed()
	config.bech32AddressPrefix["account_addr"] = addressPrefix
	config.bech32AddressPrefix["account_pub"] = pubKeyPrefix
}



func (config *Config) SetBech32PrefixForValidator(addressPrefix, pubKeyPrefix string) {
	config.assertNotSealed()
	config.bech32AddressPrefix["validator_addr"] = addressPrefix
	config.bech32AddressPrefix["validator_pub"] = pubKeyPrefix
}



func (config *Config) SetBech32PrefixForConsensusNode(addressPrefix, pubKeyPrefix string) {
	config.assertNotSealed()
	config.bech32AddressPrefix["consensus_addr"] = addressPrefix
	config.bech32AddressPrefix["consensus_pub"] = pubKeyPrefix
}


func (config *Config) SetTxEncoder(encoder TxEncoder) {
	config.assertNotSealed()
	config.txEncoder = encoder
}



func (config *Config) SetAddressVerifier(addressVerifier func([]byte) error) {
	config.assertNotSealed()
	config.addressVerifier = addressVerifier
}


//

func (config *Config) SetFullFundraiserPath(fullFundraiserPath string) {
	config.assertNotSealed()
	config.fullFundraiserPath = fullFundraiserPath
}


func (config *Config) SetPurpose(purpose uint32) {
	config.assertNotSealed()
	config.purpose = purpose
}


func (config *Config) SetCoinType(coinType uint32) {
	config.assertNotSealed()
	config.coinType = coinType
}


func (config *Config) Seal() *Config {
	config.mtx.Lock()

	if config.sealed {
		config.mtx.Unlock()
		return config
	}


	config.sealed = true
	config.mtx.Unlock()
	close(config.sealedch)

	return config
}


func (config *Config) GetBech32AccountAddrPrefix() string {
	return config.bech32AddressPrefix["account_addr"]
}


func (config *Config) GetBech32ValidatorAddrPrefix() string {
	return config.bech32AddressPrefix["validator_addr"]
}


func (config *Config) GetBech32ConsensusAddrPrefix() string {
	return config.bech32AddressPrefix["consensus_addr"]
}


func (config *Config) GetBech32AccountPubPrefix() string {
	return config.bech32AddressPrefix["account_pub"]
}


func (config *Config) GetBech32ValidatorPubPrefix() string {
	return config.bech32AddressPrefix["validator_pub"]
}


func (config *Config) GetBech32ConsensusPubPrefix() string {
	return config.bech32AddressPrefix["consensus_pub"]
}


func (config *Config) GetTxEncoder() TxEncoder {
	return config.txEncoder
}


func (config *Config) GetAddressVerifier() func([]byte) error {
	return config.addressVerifier
}


func (config *Config) GetPurpose() uint32 {
	return config.purpose
}


func (config *Config) GetCoinType() uint32 {
	return config.coinType
}


//

func (config *Config) GetFullFundraiserPath() string {
	return config.fullFundraiserPath
}


func (config *Config) GetFullBIP44Path() string {
	return fmt.Sprintf("m/%d'/%d'/0'/0/0", config.purpose, config.coinType)
}

func KeyringServiceName() string {
	if len(version.Name) == 0 {
		return DefaultKeyringServiceName
	}
	return version.Name
}

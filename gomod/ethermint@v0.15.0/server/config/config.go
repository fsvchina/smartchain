package config

import (
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/strings"

	"github.com/cosmos/cosmos-sdk/server/config"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	
	DefaultGRPCAddress = "0.0.0.0:9900"

	
	DefaultJSONRPCAddress = "0.0.0.0:8545"

	
	DefaultJSONRPCWsAddress = "0.0.0.0:8546"

	
	DefaultEVMTracer = ""

	DefaultMaxTxGasWanted = 500000

	DefaultGasCap uint64 = 25000000

	DefaultFilterCap int32 = 200

	DefaultFeeHistoryCap int32 = 100

	DefaultLogsCap int32 = 10000

	DefaultBlockRangeCap int32 = 10000

	DefaultEVMTimeout = 5 * time.Second
	
	DefaultTxFeeCap float64 = 1.0

	DefaultHTTPTimeout = 30 * time.Second

	DefaultHTTPIdleTimeout = 120 * time.Second
)

var evmTracers = []string{"json", "markdown", "struct", "access_list"}



type Config struct {
	config.Config

	EVM     EVMConfig     `mapstructure:"evm"`
	JSONRPC JSONRPCConfig `mapstructure:"json-rpc"`
	TLS     TLSConfig     `mapstructure:"tls"`
}


type EVMConfig struct {
	
	
	Tracer string `mapstructure:"tracer"`
	
	MaxTxGasWanted uint64 `mapstructure:"max-tx-gas-wanted"`
}


type JSONRPCConfig struct {
	
	API []string `mapstructure:"api"`
	
	Address string `mapstructure:"address"`
	
	WsAddress string `mapstructure:"ws-address"`
	
	GasCap uint64 `mapstructure:"gas-cap"`
	
	EVMTimeout time.Duration `mapstructure:"evm-timeout"`
	
	TxFeeCap float64 `mapstructure:"txfee-cap"`
	
	FilterCap int32 `mapstructure:"filter-cap"`
	
	FeeHistoryCap int32 `mapstructure:"feehistory-cap"`
	
	Enable bool `mapstructure:"enable"`
	
	LogsCap int32 `mapstructure:"logs-cap"`
	
	BlockRangeCap int32 `mapstructure:"block-range-cap"`
	
	HTTPTimeout time.Duration `mapstructure:"http-timeout"`
	
	HTTPIdleTimeout time.Duration `mapstructure:"http-idle-timeout"`
}


type TLSConfig struct {
	
	CertificatePath string `mapstructure:"certificate-path"`
	
	KeyPath string `mapstructure:"key-path"`
}



func AppConfig(denom string) (string, interface{}) {
	
	
	srvCfg := config.DefaultConfig()

	
	
	
	
	//
	
	
	
	
	
	//
	
	if denom != "" {
		srvCfg.MinGasPrices = "0" + denom
	}

	customAppConfig := Config{
		Config:  *srvCfg,
		EVM:     *DefaultEVMConfig(),
		JSONRPC: *DefaultJSONRPCConfig(),
		TLS:     *DefaultTLSConfig(),
	}

	customAppTemplate := config.DefaultConfigTemplate + DefaultConfigTemplate

	return customAppTemplate, customAppConfig
}


func DefaultConfig() *Config {
	return &Config{
		Config:  *config.DefaultConfig(),
		EVM:     *DefaultEVMConfig(),
		JSONRPC: *DefaultJSONRPCConfig(),
		TLS:     *DefaultTLSConfig(),
	}
}


func DefaultEVMConfig() *EVMConfig {
	return &EVMConfig{
		Tracer:         DefaultEVMTracer,
		MaxTxGasWanted: DefaultMaxTxGasWanted,
	}
}


func (c EVMConfig) Validate() error {
	if c.Tracer != "" && !strings.StringInSlice(c.Tracer, evmTracers) {
		return fmt.Errorf("invalid tracer type %s, available types: %v", c.Tracer, evmTracers)
	}

	return nil
}


func GetDefaultAPINamespaces() []string {
	return []string{"eth", "net", "web3"}
}


func GetAPINamespaces() []string {
	return []string{"web3", "eth", "personal", "net", "txpool", "debug", "miner"}
}


func DefaultJSONRPCConfig() *JSONRPCConfig {
	return &JSONRPCConfig{
		Enable:          true,
		API:             GetDefaultAPINamespaces(),
		Address:         DefaultJSONRPCAddress,
		WsAddress:       DefaultJSONRPCWsAddress,
		GasCap:          DefaultGasCap,
		EVMTimeout:      DefaultEVMTimeout,
		TxFeeCap:        DefaultTxFeeCap,
		FilterCap:       DefaultFilterCap,
		FeeHistoryCap:   DefaultFeeHistoryCap,
		BlockRangeCap:   DefaultBlockRangeCap,
		LogsCap:         DefaultLogsCap,
		HTTPTimeout:     DefaultHTTPTimeout,
		HTTPIdleTimeout: DefaultHTTPIdleTimeout,
	}
}


func (c JSONRPCConfig) Validate() error {
	if c.Enable && len(c.API) == 0 {
		return errors.New("cannot enable JSON-RPC without defining any API namespace")
	}

	if c.FilterCap < 0 {
		return errors.New("JSON-RPC filter-cap cannot be negative")
	}

	if c.FeeHistoryCap <= 0 {
		return errors.New("JSON-RPC feehistory-cap cannot be negative or 0")
	}

	if c.TxFeeCap < 0 {
		return errors.New("JSON-RPC tx fee cap cannot be negative")
	}

	if c.EVMTimeout < 0 {
		return errors.New("JSON-RPC EVM timeout duration cannot be negative")
	}

	if c.LogsCap < 0 {
		return errors.New("JSON-RPC logs cap cannot be negative")
	}

	if c.BlockRangeCap < 0 {
		return errors.New("JSON-RPC block range cap cannot be negative")
	}

	if c.HTTPTimeout < 0 {
		return errors.New("JSON-RPC HTTP timeout duration cannot be negative")
	}

	if c.HTTPIdleTimeout < 0 {
		return errors.New("JSON-RPC HTTP idle timeout duration cannot be negative")
	}

	
	seenAPIs := make(map[string]bool)
	for _, api := range c.API {
		if seenAPIs[api] {
			return fmt.Errorf("repeated API namespace '%s'", api)
		}

		seenAPIs[api] = true
	}

	return nil
}


func DefaultTLSConfig() *TLSConfig {
	return &TLSConfig{
		CertificatePath: "",
		KeyPath:         "",
	}
}


func (c TLSConfig) Validate() error {
	certExt := path.Ext(c.CertificatePath)

	if c.CertificatePath != "" && certExt != ".pem" {
		return fmt.Errorf("invalid extension %s for certificate path %s, expected '.pem'", certExt, c.CertificatePath)
	}

	keyExt := path.Ext(c.KeyPath)

	if c.KeyPath != "" && keyExt != ".pem" {
		return fmt.Errorf("invalid extension %s for key path %s, expected '.pem'", keyExt, c.KeyPath)
	}

	return nil
}


func GetConfig(v *viper.Viper) Config {
	cfg := config.GetConfig(v)

	return Config{
		Config: cfg,
		EVM: EVMConfig{
			Tracer:         v.GetString("evm.tracer"),
			MaxTxGasWanted: v.GetUint64("evm.max-tx-gas-wanted"),
		},
		JSONRPC: JSONRPCConfig{
			Enable:          v.GetBool("json-rpc.enable"),
			API:             v.GetStringSlice("json-rpc.api"),
			Address:         v.GetString("json-rpc.address"),
			WsAddress:       v.GetString("json-rpc.ws-address"),
			GasCap:          v.GetUint64("json-rpc.gas-cap"),
			FilterCap:       v.GetInt32("json-rpc.filter-cap"),
			FeeHistoryCap:   v.GetInt32("json-rpc.feehistory-cap"),
			TxFeeCap:        v.GetFloat64("json-rpc.txfee-cap"),
			EVMTimeout:      v.GetDuration("json-rpc.evm-timeout"),
			LogsCap:         v.GetInt32("json-rpc.logs-cap"),
			BlockRangeCap:   v.GetInt32("json-rpc.block-range-cap"),
			HTTPTimeout:     v.GetDuration("json-rpc.http-timeout"),
			HTTPIdleTimeout: v.GetDuration("json-rpc.http-idle-timeout"),
		},
		TLS: TLSConfig{
			CertificatePath: v.GetString("tls.certificate-path"),
			KeyPath:         v.GetString("tls.key-path"),
		},
	}
}



func ParseConfig(v *viper.Viper) (*Config, error) {
	conf := DefaultConfig()
	err := v.Unmarshal(conf)

	return conf, err
}


func (c Config) ValidateBasic() error {
	if err := c.EVM.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrAppConfig, "invalid evm config value: %s", err.Error())
	}

	if err := c.JSONRPC.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrAppConfig, "invalid json-rpc config value: %s", err.Error())
	}

	if err := c.TLS.Validate(); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrAppConfig, "invalid tls config value: %s", err.Error())
	}

	return c.Config.ValidateBasic()
}

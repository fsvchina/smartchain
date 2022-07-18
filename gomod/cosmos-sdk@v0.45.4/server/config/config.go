package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	defaultMinGasPrices = ""


	DefaultGRPCAddress = "0.0.0.0:9090"


	DefaultGRPCWebAddress = "0.0.0.0:9091"
)


type BaseConfig struct {



	MinGasPrices string `mapstructure:"minimum-gas-prices"`

	Pruning           string `mapstructure:"pruning"`
	PruningKeepRecent string `mapstructure:"pruning-keep-recent"`
	PruningKeepEvery  string `mapstructure:"pruning-keep-every"`
	PruningInterval   string `mapstructure:"pruning-interval"`



	//

	HaltHeight uint64 `mapstructure:"halt-height"`




	//

	HaltTime uint64 `mapstructure:"halt-time"`






	//



	//




	MinRetainBlocks uint64 `mapstructure:"min-retain-blocks"`


	InterBlockCache bool `mapstructure:"inter-block-cache"`



	IndexEvents []string `mapstructure:"index-events"`

	IAVLCacheSize uint64 `mapstructure:"iavl-cache-size"`
}


type APIConfig struct {

	Enable bool `mapstructure:"enable"`


	Swagger bool `mapstructure:"swagger"`


	EnableUnsafeCORS bool `mapstructure:"enabled-unsafe-cors"`


	Address string `mapstructure:"address"`


	MaxOpenConnections uint `mapstructure:"max-open-connections"`


	RPCReadTimeout uint `mapstructure:"rpc-read-timeout"`


	RPCWriteTimeout uint `mapstructure:"rpc-write-timeout"`


	RPCMaxBodyBytes uint `mapstructure:"rpc-max-body-bytes"`


	//

}


type RosettaConfig struct {

	Address string `mapstructure:"address"`



	Blockchain string `mapstructure:"blockchain"`


	Network string `mapstructure:"network"`



	Retries int `mapstructure:"retries"`


	Enable bool `mapstructure:"enable"`


	Offline bool `mapstructure:"offline"`
}


type GRPCConfig struct {

	Enable bool `mapstructure:"enable"`


	Address string `mapstructure:"address"`
}


type GRPCWebConfig struct {

	Enable bool `mapstructure:"enable"`


	Address string `mapstructure:"address"`


	EnableUnsafeCORS bool `mapstructure:"enable-unsafe-cors"`
}


type StateSyncConfig struct {


	SnapshotInterval uint64 `mapstructure:"snapshot-interval"`



	SnapshotKeepRecent uint32 `mapstructure:"snapshot-keep-recent"`
}


type Config struct {
	BaseConfig `mapstructure:",squash"`


	Telemetry telemetry.Config `mapstructure:"telemetry"`
	API       APIConfig        `mapstructure:"api"`
	GRPC      GRPCConfig       `mapstructure:"grpc"`
	Rosetta   RosettaConfig    `mapstructure:"rosetta"`
	GRPCWeb   GRPCWebConfig    `mapstructure:"grpc-web"`
	StateSync StateSyncConfig  `mapstructure:"state-sync"`
}


func (c *Config) SetMinGasPrices(gasPrices sdk.DecCoins) {
	c.MinGasPrices = gasPrices.String()
}



func (c *Config) GetMinGasPrices() sdk.DecCoins {
	if c.MinGasPrices == "" {
		return sdk.DecCoins{}
	}

	gasPricesStr := strings.Split(c.MinGasPrices, ";")
	gasPrices := make(sdk.DecCoins, len(gasPricesStr))

	for i, s := range gasPricesStr {
		gasPrice, err := sdk.ParseDecCoin(s)
		if err != nil {
			panic(fmt.Errorf("failed to parse minimum gas price coin (%s): %s", s, err))
		}

		gasPrices[i] = gasPrice
	}

	return gasPrices
}


func DefaultConfig() *Config {
	return &Config{
		BaseConfig: BaseConfig{
			MinGasPrices:      defaultMinGasPrices,
			InterBlockCache:   true,
			Pruning:           storetypes.PruningOptionDefault,
			PruningKeepRecent: "0",
			PruningKeepEvery:  "0",
			PruningInterval:   "0",
			MinRetainBlocks:   0,
			IndexEvents:       make([]string, 0),
			IAVLCacheSize:     781250,
		},
		Telemetry: telemetry.Config{
			Enabled:      false,
			GlobalLabels: [][]string{},
		},
		API: APIConfig{
			Enable:             false,
			Swagger:            false,
			Address:            "tcp:
			MaxOpenConnections: 1000,
			RPCReadTimeout:     10,
			RPCMaxBodyBytes:    1000000,
		},
		GRPC: GRPCConfig{
			Enable:  true,
			Address: DefaultGRPCAddress,
		},
		Rosetta: RosettaConfig{
			Enable:     false,
			Address:    ":8080",
			Blockchain: "app",
			Network:    "network",
			Retries:    3,
			Offline:    false,
		},
		GRPCWeb: GRPCWebConfig{
			Enable:  true,
			Address: DefaultGRPCWebAddress,
		},
		StateSync: StateSyncConfig{
			SnapshotInterval:   0,
			SnapshotKeepRecent: 2,
		},
	}
}


func GetConfig(v *viper.Viper) Config {
	globalLabelsRaw := v.Get("telemetry.global-labels").([]interface{})
	globalLabels := make([][]string, 0, len(globalLabelsRaw))
	for _, glr := range globalLabelsRaw {
		labelsRaw := glr.([]interface{})
		if len(labelsRaw) == 2 {
			globalLabels = append(globalLabels, []string{labelsRaw[0].(string), labelsRaw[1].(string)})
		}
	}

	return Config{
		BaseConfig: BaseConfig{
			MinGasPrices:      v.GetString("minimum-gas-prices"),
			InterBlockCache:   v.GetBool("inter-block-cache"),
			Pruning:           v.GetString("pruning"),
			PruningKeepRecent: v.GetString("pruning-keep-recent"),
			PruningKeepEvery:  v.GetString("pruning-keep-every"),
			PruningInterval:   v.GetString("pruning-interval"),
			HaltHeight:        v.GetUint64("halt-height"),
			HaltTime:          v.GetUint64("halt-time"),
			IndexEvents:       v.GetStringSlice("index-events"),
			MinRetainBlocks:   v.GetUint64("min-retain-blocks"),
			IAVLCacheSize:     v.GetUint64("iavl-cache-size"),
		},
		Telemetry: telemetry.Config{
			ServiceName:             v.GetString("telemetry.service-name"),
			Enabled:                 v.GetBool("telemetry.enabled"),
			EnableHostname:          v.GetBool("telemetry.enable-hostname"),
			EnableHostnameLabel:     v.GetBool("telemetry.enable-hostname-label"),
			EnableServiceLabel:      v.GetBool("telemetry.enable-service-label"),
			PrometheusRetentionTime: v.GetInt64("telemetry.prometheus-retention-time"),
			GlobalLabels:            globalLabels,
		},
		API: APIConfig{
			Enable:             v.GetBool("api.enable"),
			Swagger:            v.GetBool("api.swagger"),
			Address:            v.GetString("api.address"),
			MaxOpenConnections: v.GetUint("api.max-open-connections"),
			RPCReadTimeout:     v.GetUint("api.rpc-read-timeout"),
			RPCWriteTimeout:    v.GetUint("api.rpc-write-timeout"),
			RPCMaxBodyBytes:    v.GetUint("api.rpc-max-body-bytes"),
			EnableUnsafeCORS:   v.GetBool("api.enabled-unsafe-cors"),
		},
		Rosetta: RosettaConfig{
			Enable:     v.GetBool("rosetta.enable"),
			Address:    v.GetString("rosetta.address"),
			Blockchain: v.GetString("rosetta.blockchain"),
			Network:    v.GetString("rosetta.network"),
			Retries:    v.GetInt("rosetta.retries"),
			Offline:    v.GetBool("rosetta.offline"),
		},
		GRPC: GRPCConfig{
			Enable:  v.GetBool("grpc.enable"),
			Address: v.GetString("grpc.address"),
		},
		GRPCWeb: GRPCWebConfig{
			Enable:           v.GetBool("grpc-web.enable"),
			Address:          v.GetString("grpc-web.address"),
			EnableUnsafeCORS: v.GetBool("grpc-web.enable-unsafe-cors"),
		},
		StateSync: StateSyncConfig{
			SnapshotInterval:   v.GetUint64("state-sync.snapshot-interval"),
			SnapshotKeepRecent: v.GetUint32("state-sync.snapshot-keep-recent"),
		},
	}
}


func (c Config) ValidateBasic() error {
	if c.BaseConfig.MinGasPrices == "" {
		return sdkerrors.ErrAppConfig.Wrap("set min gas price in app.toml or flag or env variable")
	}
	if c.Pruning == storetypes.PruningOptionEverything && c.StateSync.SnapshotInterval > 0 {
		return sdkerrors.ErrAppConfig.Wrapf(
			"cannot enable state sync snapshots with '%s' pruning setting", storetypes.PruningOptionEverything,
		)
	}

	return nil
}

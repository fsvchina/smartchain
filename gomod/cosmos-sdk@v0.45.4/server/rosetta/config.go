package rosetta

import (
	"fmt"
	"strings"
	"time"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/spf13/pflag"

	crg "github.com/cosmos/cosmos-sdk/server/rosetta/lib/server"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)


const (

	DefaultBlockchain = "app"

	DefaultAddr = ":8080"

	DefaultRetries = 5

	DefaultTendermintEndpoint = "localhost:26657"

	DefaultGRPCEndpoint = "localhost:9090"

	DefaultNetwork = "network"

	DefaultOffline = false
)


const (
	FlagBlockchain         = "blockchain"
	FlagNetwork            = "network"
	FlagTendermintEndpoint = "tendermint"
	FlagGRPCEndpoint       = "grpc"
	FlagAddr               = "addr"
	FlagRetries            = "retries"
	FlagOffline            = "offline"
)


type Config struct {


	Blockchain string

	Network string



	TendermintRPC string


	GRPCEndpoint string


	Addr string


	Retries int

	Offline bool

	Codec *codec.ProtoCodec

	InterfaceRegistry codectypes.InterfaceRegistry
}


func (c *Config) NetworkIdentifier() *types.NetworkIdentifier {
	return &types.NetworkIdentifier{
		Blockchain: c.Blockchain,
		Network:    c.Network,
	}
}



func (c *Config) validate() error {
	if (c.Codec == nil) != (c.InterfaceRegistry == nil) {
		return fmt.Errorf("codec and interface registry must be both different from nil or nil")
	}

	if c.Addr == "" {
		c.Addr = DefaultAddr
	}
	if c.Blockchain == "" {
		c.Blockchain = DefaultBlockchain
	}
	if c.Retries == 0 {
		c.Retries = DefaultRetries
	}

	if c.Network == "" {
		return fmt.Errorf("network not provided")
	}
	if c.Offline {
		return fmt.Errorf("offline mode is not supported for stargate implementation due to how sigv2 works")
	}


	if c.GRPCEndpoint == "" {
		return fmt.Errorf("grpc endpoint not provided")
	}
	if c.TendermintRPC == "" {
		return fmt.Errorf("tendermint rpc not provided")
	}
	if !strings.HasPrefix(c.TendermintRPC, "tcp:
		c.TendermintRPC = fmt.Sprintf("tcp:
	}

	return nil
}


func (c *Config) WithCodec(ir codectypes.InterfaceRegistry, cdc *codec.ProtoCodec) {
	c.Codec = cdc
	c.InterfaceRegistry = ir
}


func FromFlags(flags *pflag.FlagSet) (*Config, error) {
	blockchain, err := flags.GetString(FlagBlockchain)
	if err != nil {
		return nil, err
	}
	network, err := flags.GetString(FlagNetwork)
	if err != nil {
		return nil, err
	}
	tendermintRPC, err := flags.GetString(FlagTendermintEndpoint)
	if err != nil {
		return nil, err
	}
	gRPCEndpoint, err := flags.GetString(FlagGRPCEndpoint)
	if err != nil {
		return nil, err
	}
	addr, err := flags.GetString(FlagAddr)
	if err != nil {
		return nil, err
	}
	retries, err := flags.GetInt(FlagRetries)
	if err != nil {
		return nil, err
	}
	offline, err := flags.GetBool(FlagOffline)
	if err != nil {
		return nil, err
	}
	conf := &Config{
		Blockchain:    blockchain,
		Network:       network,
		TendermintRPC: tendermintRPC,
		GRPCEndpoint:  gRPCEndpoint,
		Addr:          addr,
		Retries:       retries,
		Offline:       offline,
	}
	err = conf.validate()
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func ServerFromConfig(conf *Config) (crg.Server, error) {
	err := conf.validate()
	if err != nil {
		return crg.Server{}, err
	}
	client, err := NewClient(conf)
	if err != nil {
		return crg.Server{}, err
	}
	return crg.NewServer(
		crg.Settings{
			Network: &types.NetworkIdentifier{
				Blockchain: conf.Blockchain,
				Network:    conf.Network,
			},
			Client:    client,
			Listen:    conf.Addr,
			Offline:   conf.Offline,
			Retries:   conf.Retries,
			RetryWait: 15 * time.Second,
		})
}


func SetFlags(flags *pflag.FlagSet) {
	flags.String(FlagBlockchain, DefaultBlockchain, "the blockchain type")
	flags.String(FlagNetwork, DefaultNetwork, "the network name")
	flags.String(FlagTendermintEndpoint, DefaultTendermintEndpoint, "the tendermint rpc endpoint, without tcp:
	flags.String(FlagGRPCEndpoint, DefaultGRPCEndpoint, "the app gRPC endpoint")
	flags.String(FlagAddr, DefaultAddr, "the address rosetta will bind to")
	flags.Int(FlagRetries, DefaultRetries, "the number of retries that will be done before quitting")
	flags.Bool(FlagOffline, DefaultOffline, "run rosetta only with construction API")
}

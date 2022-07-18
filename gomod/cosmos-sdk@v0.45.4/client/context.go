package client

import (
	"bufio"
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/viper"

	"gopkg.in/yaml.v2"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



type Context struct {
	FromAddress sdk.AccAddress
	Client      rpcclient.Client
	ChainID     string

	JSONCodec         codec.JSONCodec
	Codec             codec.Codec
	InterfaceRegistry codectypes.InterfaceRegistry
	Input             io.Reader
	Keyring           keyring.Keyring
	KeyringOptions    []keyring.Option
	Output            io.Writer
	OutputFormat      string
	Height            int64
	HomeDir           string
	KeyringDir        string
	From              string
	BroadcastMode     string
	FromName          string
	SignModeStr       string
	UseLedger         bool
	Simulate          bool
	GenerateOnly      bool
	Offline           bool
	SkipConfirm       bool
	TxConfig          TxConfig
	AccountRetriever  AccountRetriever
	NodeURI           string
	FeeGranter        sdk.AccAddress
	Viper             *viper.Viper


	LegacyAmino *codec.LegacyAmino
}


func (ctx Context) WithKeyring(k keyring.Keyring) Context {
	ctx.Keyring = k
	return ctx
}


func (ctx Context) WithKeyringOptions(opts ...keyring.Option) Context {
	ctx.KeyringOptions = opts
	return ctx
}


func (ctx Context) WithInput(r io.Reader) Context {



	ctx.Input = bufio.NewReader(r)
	return ctx
}


func (ctx Context) WithJSONCodec(m codec.JSONCodec) Context {
	ctx.JSONCodec = m


	if c, ok := m.(codec.Codec); ok {
		ctx.Codec = c
	}
	return ctx
}


func (ctx Context) WithCodec(m codec.Codec) Context {
	ctx.JSONCodec = m
	ctx.Codec = m
	return ctx
}



func (ctx Context) WithLegacyAmino(cdc *codec.LegacyAmino) Context {
	ctx.LegacyAmino = cdc
	return ctx
}


func (ctx Context) WithOutput(w io.Writer) Context {
	ctx.Output = w
	return ctx
}


func (ctx Context) WithFrom(from string) Context {
	ctx.From = from
	return ctx
}


func (ctx Context) WithOutputFormat(format string) Context {
	ctx.OutputFormat = format
	return ctx
}


func (ctx Context) WithNodeURI(nodeURI string) Context {
	ctx.NodeURI = nodeURI
	return ctx
}


func (ctx Context) WithHeight(height int64) Context {
	ctx.Height = height
	return ctx
}



func (ctx Context) WithClient(client rpcclient.Client) Context {
	ctx.Client = client
	return ctx
}


func (ctx Context) WithUseLedger(useLedger bool) Context {
	ctx.UseLedger = useLedger
	return ctx
}


func (ctx Context) WithChainID(chainID string) Context {
	ctx.ChainID = chainID
	return ctx
}


func (ctx Context) WithHomeDir(dir string) Context {
	if dir != "" {
		ctx.HomeDir = dir
	}
	return ctx
}


func (ctx Context) WithKeyringDir(dir string) Context {
	ctx.KeyringDir = dir
	return ctx
}


func (ctx Context) WithGenerateOnly(generateOnly bool) Context {
	ctx.GenerateOnly = generateOnly
	return ctx
}


func (ctx Context) WithSimulation(simulate bool) Context {
	ctx.Simulate = simulate
	return ctx
}


func (ctx Context) WithOffline(offline bool) Context {
	ctx.Offline = offline
	return ctx
}


func (ctx Context) WithFromName(name string) Context {
	ctx.FromName = name
	return ctx
}



func (ctx Context) WithFromAddress(addr sdk.AccAddress) Context {
	ctx.FromAddress = addr
	return ctx
}



func (ctx Context) WithFeeGranterAddress(addr sdk.AccAddress) Context {
	ctx.FeeGranter = addr
	return ctx
}



func (ctx Context) WithBroadcastMode(mode string) Context {
	ctx.BroadcastMode = mode
	return ctx
}



func (ctx Context) WithSignModeStr(signModeStr string) Context {
	ctx.SignModeStr = signModeStr
	return ctx
}



func (ctx Context) WithSkipConfirmation(skip bool) Context {
	ctx.SkipConfirm = skip
	return ctx
}


func (ctx Context) WithTxConfig(generator TxConfig) Context {
	ctx.TxConfig = generator
	return ctx
}


func (ctx Context) WithAccountRetriever(retriever AccountRetriever) Context {
	ctx.AccountRetriever = retriever
	return ctx
}


func (ctx Context) WithInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) Context {
	ctx.InterfaceRegistry = interfaceRegistry
	return ctx
}



func (ctx Context) WithViper(prefix string) Context {
	v := viper.New()
	v.SetEnvPrefix(prefix)
	v.AutomaticEnv()
	ctx.Viper = v
	return ctx
}


func (ctx Context) PrintString(str string) error {
	return ctx.PrintBytes([]byte(str))
}



func (ctx Context) PrintBytes(o []byte) error {
	writer := ctx.Output
	if writer == nil {
		writer = os.Stdout
	}

	_, err := writer.Write(o)
	return err
}




func (ctx Context) PrintProto(toPrint proto.Message) error {

	out, err := ctx.Codec.MarshalJSON(toPrint)
	if err != nil {
		return err
	}
	return ctx.printOutput(out)
}




func (ctx Context) PrintObjectLegacy(toPrint interface{}) error {
	out, err := ctx.LegacyAmino.MarshalJSON(toPrint)
	if err != nil {
		return err
	}
	return ctx.printOutput(out)
}

func (ctx Context) printOutput(out []byte) error {
	if ctx.OutputFormat == "text" {

		var j interface{}

		err := json.Unmarshal(out, &j)
		if err != nil {
			return err
		}

		out, err = yaml.Marshal(j)
		if err != nil {
			return err
		}
	}

	writer := ctx.Output
	if writer == nil {
		writer = os.Stdout
	}

	_, err := writer.Write(out)
	if err != nil {
		return err
	}

	if ctx.OutputFormat != "text" {

		_, err = writer.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	return nil
}




func GetFromFields(kr keyring.Keyring, from string, genOnly bool) (sdk.AccAddress, string, keyring.KeyType, error) {
	if from == "" {
		return nil, "", 0, nil
	}

	if genOnly {
		addr, err := sdk.AccAddressFromBech32(from)
		if err != nil {
			return nil, "", 0, errors.Wrap(err, "must provide a valid Bech32 address in generate-only mode")
		}

		return addr, "", 0, nil
	}

	var info keyring.Info
	if addr, err := sdk.AccAddressFromBech32(from); err == nil {
		info, err = kr.KeyByAddress(addr)
		if err != nil {
			return nil, "", 0, err
		}
	} else {
		info, err = kr.Key(from)
		if err != nil {
			return nil, "", 0, err
		}
	}

	return info.GetAddress(), info.GetName(), info.GetType(), nil
}


func NewKeyringFromBackend(ctx Context, backend string) (keyring.Keyring, error) {
	if ctx.GenerateOnly || ctx.Simulate {
		return keyring.New(sdk.KeyringServiceName(), keyring.BackendMemory, ctx.KeyringDir, ctx.Input, ctx.KeyringOptions...)
	}

	return keyring.New(sdk.KeyringServiceName(), backend, ctx.KeyringDir, ctx.Input, ctx.KeyringOptions...)
}

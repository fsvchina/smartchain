

package legacybech32

import (
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)






type Bech32PubKeyType string


const (
	AccPK  Bech32PubKeyType = "accpub"
	ValPK  Bech32PubKeyType = "valpub"
	ConsPK Bech32PubKeyType = "conspub"
)



func MarshalPubKey(pkt Bech32PubKeyType, pubkey cryptotypes.PubKey) (string, error) {
	bech32Prefix := getPrefix(pkt)
	return bech32.ConvertAndEncode(bech32Prefix, legacy.Cdc.MustMarshal(pubkey))
}


func MustMarshalPubKey(pkt Bech32PubKeyType, pubkey cryptotypes.PubKey) string {
	res, err := MarshalPubKey(pkt, pubkey)
	if err != nil {
		panic(err)
	}

	return res
}

func getPrefix(pkt Bech32PubKeyType) string {
	cfg := sdk.GetConfig()
	switch pkt {
	case AccPK:
		return cfg.GetBech32AccountPubPrefix()

	case ValPK:
		return cfg.GetBech32ValidatorPubPrefix()
	case ConsPK:
		return cfg.GetBech32ConsensusPubPrefix()
	}

	return ""
}



func UnmarshalPubKey(pkt Bech32PubKeyType, pubkeyStr string) (cryptotypes.PubKey, error) {
	bech32Prefix := getPrefix(pkt)

	bz, err := sdk.GetFromBech32(pubkeyStr, bech32Prefix)
	if err != nil {
		return nil, err
	}
	return legacy.PubKeyFromBytes(bz)
}

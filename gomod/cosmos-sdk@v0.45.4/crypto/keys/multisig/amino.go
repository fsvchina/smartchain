package multisig

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)



//


//





//




//

type tmMultisig struct {
	K       uint                 `json:"threshold"`
	PubKeys []cryptotypes.PubKey `json:"pubkeys"`
}


func protoToTm(protoPk *LegacyAminoPubKey) (tmMultisig, error) {
	var ok bool
	pks := make([]cryptotypes.PubKey, len(protoPk.PubKeys))
	for i, pk := range protoPk.PubKeys {
		pks[i], ok = pk.GetCachedValue().(cryptotypes.PubKey)
		if !ok {
			return tmMultisig{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expected %T, got %T", (cryptotypes.PubKey)(nil), pk.GetCachedValue())
		}
	}

	return tmMultisig{
		K:       uint(protoPk.Threshold),
		PubKeys: pks,
	}, nil
}


func tmToProto(tmPk tmMultisig) (*LegacyAminoPubKey, error) {
	var err error
	pks := make([]*types.Any, len(tmPk.PubKeys))
	for i, pk := range tmPk.PubKeys {
		pks[i], err = types.NewAnyWithValue(pk)
		if err != nil {
			return nil, err
		}
	}

	return &LegacyAminoPubKey{
		Threshold: uint32(tmPk.K),
		PubKeys:   pks,
	}, nil
}


func (m LegacyAminoPubKey) MarshalAminoJSON() (tmMultisig, error) {
	return protoToTm(&m)
}


func (m *LegacyAminoPubKey) UnmarshalAminoJSON(tmPk tmMultisig) error {
	protoPk, err := tmToProto(tmPk)
	if err != nil {
		return err
	}




	if m.PubKeys == nil {
		m.PubKeys = make([]*types.Any, len(tmPk.PubKeys))
	}
	for i := range m.PubKeys {
		if m.PubKeys[i] == nil {

			bz, err := AminoCdc.MarshalJSON(tmPk.PubKeys[i])
			if err != nil {
				return err
			}

			m.PubKeys[i] = protoPk.PubKeys[i]



			if err := m.PubKeys[i].UnmarshalJSON(bz); err != nil {
				return err
			}
		} else {
			m.PubKeys[i].TypeUrl = protoPk.PubKeys[i].TypeUrl
			m.PubKeys[i].Value = protoPk.PubKeys[i].Value
		}
	}
	m.Threshold = protoPk.Threshold

	return nil
}

package types

import (
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
)


type WriteListener interface {



	OnWrite(storeKey StoreKey, key []byte, value []byte, delete bool) error
}



type StoreKVPairWriteListener struct {
	writer     io.Writer
	marshaller codec.BinaryCodec
}


func NewStoreKVPairWriteListener(w io.Writer, m codec.BinaryCodec) *StoreKVPairWriteListener {
	return &StoreKVPairWriteListener{
		writer:     w,
		marshaller: m,
	}
}


func (wl *StoreKVPairWriteListener) OnWrite(storeKey StoreKey, key []byte, value []byte, delete bool) error {
	kvPair := new(StoreKVPair)
	kvPair.StoreKey = storeKey.Name()
	kvPair.Delete = delete
	kvPair.Key = key
	kvPair.Value = value
	by, err := wl.marshaller.MarshalLengthPrefixed(kvPair)
	if err != nil {
		return err
	}
	if _, err := wl.writer.Write(by); err != nil {
		return err
	}
	return nil
}

package types

import (
	"fmt"
	"reflect"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	StoreKey = "params"


	TStoreKey = "transient_params"
)




type Subspace struct {
	cdc         codec.BinaryCodec
	legacyAmino *codec.LegacyAmino
	key         sdk.StoreKey
	tkey        sdk.StoreKey
	name        []byte
	table       KeyTable
}


func NewSubspace(cdc codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key sdk.StoreKey, tkey sdk.StoreKey, name string) Subspace {
	return Subspace{
		cdc:         cdc,
		legacyAmino: legacyAmino,
		key:         key,
		tkey:        tkey,
		name:        []byte(name),
		table:       NewKeyTable(),
	}
}


func (s Subspace) HasKeyTable() bool {
	return len(s.table.m) > 0
}


func (s Subspace) WithKeyTable(table KeyTable) Subspace {
	if table.m == nil {
		panic("SetKeyTable() called with nil KeyTable")
	}
	if len(s.table.m) != 0 {
		panic("SetKeyTable() called on already initialized Subspace")
	}

	for k, v := range table.m {
		s.table.m[k] = v
	}



	name := s.name
	s.name = make([]byte, len(name), len(name)+table.maxKeyLength())
	copy(s.name, name)

	return s
}


func (s Subspace) kvStore(ctx sdk.Context) sdk.KVStore {


	return prefix.NewStore(ctx.KVStore(s.key), append(s.name, '/'))
}


func (s Subspace) transientStore(ctx sdk.Context) sdk.KVStore {


	return prefix.NewStore(ctx.TransientStore(s.tkey), append(s.name, '/'))
}



func (s Subspace) Validate(ctx sdk.Context, key []byte, value interface{}) error {
	attr, ok := s.table.m[string(key)]
	if !ok {
		return fmt.Errorf("parameter %s not registered", string(key))
	}

	if err := attr.vfn(value); err != nil {
		return fmt.Errorf("invalid parameter value: %s", err)
	}

	return nil
}



func (s Subspace) Get(ctx sdk.Context, key []byte, ptr interface{}) {
	s.checkType(key, ptr)

	store := s.kvStore(ctx)
	bz := store.Get(key)

	if err := s.legacyAmino.UnmarshalJSON(bz, ptr); err != nil {
		panic(err)
	}
}




func (s Subspace) GetIfExists(ctx sdk.Context, key []byte, ptr interface{}) {
	store := s.kvStore(ctx)
	bz := store.Get(key)
	if bz == nil {
		return
	}

	s.checkType(key, ptr)

	if err := s.legacyAmino.UnmarshalJSON(bz, ptr); err != nil {
		panic(err)
	}
}


func (s Subspace) GetRaw(ctx sdk.Context, key []byte) []byte {
	store := s.kvStore(ctx)
	return store.Get(key)
}


func (s Subspace) Has(ctx sdk.Context, key []byte) bool {
	store := s.kvStore(ctx)
	return store.Has(key)
}



func (s Subspace) Modified(ctx sdk.Context, key []byte) bool {
	tstore := s.transientStore(ctx)
	return tstore.Has(key)
}


func (s Subspace) checkType(key []byte, value interface{}) {
	attr, ok := s.table.m[string(key)]
	if !ok {
		panic(fmt.Sprintf("parameter %s not registered", string(key)))
	}

	ty := attr.ty
	pty := reflect.TypeOf(value)
	if pty.Kind() == reflect.Ptr {
		pty = pty.Elem()
	}

	if pty != ty {
		panic("type mismatch with registered table")
	}
}





func (s Subspace) Set(ctx sdk.Context, key []byte, value interface{}) {
	s.checkType(key, value)
	store := s.kvStore(ctx)

	bz, err := s.legacyAmino.MarshalJSON(value)
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)

	tstore := s.transientStore(ctx)
	tstore.Set(key, []byte{})
}







func (s Subspace) Update(ctx sdk.Context, key, value []byte) error {
	attr, ok := s.table.m[string(key)]
	if !ok {
		panic(fmt.Sprintf("parameter %s not registered", string(key)))
	}

	ty := attr.ty
	dest := reflect.New(ty).Interface()
	s.GetIfExists(ctx, key, dest)

	if err := s.legacyAmino.UnmarshalJSON(value, dest); err != nil {
		return err
	}



	destValue := reflect.Indirect(reflect.ValueOf(dest)).Interface()
	if err := s.Validate(ctx, key, destValue); err != nil {
		return err
	}

	s.Set(ctx, key, dest)
	return nil
}




func (s Subspace) GetParamSet(ctx sdk.Context, ps ParamSet) {
	for _, pair := range ps.ParamSetPairs() {
		s.Get(ctx, pair.Key, pair.Value)
	}
}



func (s Subspace) SetParamSet(ctx sdk.Context, ps ParamSet) {
	for _, pair := range ps.ParamSetPairs() {




		v := reflect.Indirect(reflect.ValueOf(pair.Value)).Interface()

		if err := pair.ValidatorFn(v); err != nil {
			panic(fmt.Sprintf("value from ParamSetPair is invalid: %s", err))
		}

		s.Set(ctx, pair.Key, v)
	}
}


func (s Subspace) Name() string {
	return string(s.name)
}


type ReadOnlySubspace struct {
	s Subspace
}


func (ros ReadOnlySubspace) Get(ctx sdk.Context, key []byte, ptr interface{}) {
	ros.s.Get(ctx, key, ptr)
}


func (ros ReadOnlySubspace) GetRaw(ctx sdk.Context, key []byte) []byte {
	return ros.s.GetRaw(ctx, key)
}


func (ros ReadOnlySubspace) Has(ctx sdk.Context, key []byte) bool {
	return ros.s.Has(ctx, key)
}


func (ros ReadOnlySubspace) Modified(ctx sdk.Context, key []byte) bool {
	return ros.s.Modified(ctx, key)
}


func (ros ReadOnlySubspace) Name() string {
	return ros.s.Name()
}

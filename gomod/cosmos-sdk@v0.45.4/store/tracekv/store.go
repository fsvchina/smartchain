package tracekv

import (
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	writeOp     operation = "write"
	readOp      operation = "read"
	deleteOp    operation = "delete"
	iterKeyOp   operation = "iterKey"
	iterValueOp operation = "iterValue"
)

type (



	//


	Store struct {
		parent  types.KVStore
		writer  io.Writer
		context types.TraceContext
	}


	operation string


	traceOperation struct {
		Operation operation              `json:"operation"`
		Key       string                 `json:"key"`
		Value     string                 `json:"value"`
		Metadata  map[string]interface{} `json:"metadata"`
	}
)



func NewStore(parent types.KVStore, writer io.Writer, tc types.TraceContext) *Store {
	return &Store{parent: parent, writer: writer, context: tc}
}



func (tkv *Store) Get(key []byte) []byte {
	value := tkv.parent.Get(key)

	writeOperation(tkv.writer, readOp, tkv.context, key, value)
	return value
}



func (tkv *Store) Set(key []byte, value []byte) {
	types.AssertValidKey(key)
	writeOperation(tkv.writer, writeOp, tkv.context, key, value)
	tkv.parent.Set(key, value)
}



func (tkv *Store) Delete(key []byte) {
	writeOperation(tkv.writer, deleteOp, tkv.context, key, nil)
	tkv.parent.Delete(key)
}



func (tkv *Store) Has(key []byte) bool {
	return tkv.parent.Has(key)
}



func (tkv *Store) Iterator(start, end []byte) types.Iterator {
	return tkv.iterator(start, end, true)
}



func (tkv *Store) ReverseIterator(start, end []byte) types.Iterator {
	return tkv.iterator(start, end, false)
}



func (tkv *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	var parent types.Iterator

	if ascending {
		parent = tkv.parent.Iterator(start, end)
	} else {
		parent = tkv.parent.ReverseIterator(start, end)
	}

	return newTraceIterator(tkv.writer, parent, tkv.context)
}

type traceIterator struct {
	parent  types.Iterator
	writer  io.Writer
	context types.TraceContext
}

func newTraceIterator(w io.Writer, parent types.Iterator, tc types.TraceContext) types.Iterator {
	return &traceIterator{writer: w, parent: parent, context: tc}
}


func (ti *traceIterator) Domain() (start []byte, end []byte) {
	return ti.parent.Domain()
}


func (ti *traceIterator) Valid() bool {
	return ti.parent.Valid()
}


func (ti *traceIterator) Next() {
	ti.parent.Next()
}


func (ti *traceIterator) Key() []byte {
	key := ti.parent.Key()

	writeOperation(ti.writer, iterKeyOp, ti.context, key, nil)
	return key
}


func (ti *traceIterator) Value() []byte {
	value := ti.parent.Value()

	writeOperation(ti.writer, iterValueOp, ti.context, nil, value)
	return value
}


func (ti *traceIterator) Close() error {
	return ti.parent.Close()
}


func (ti *traceIterator) Error() error {
	return ti.parent.Error()
}



func (tkv *Store) GetStoreType() types.StoreType {
	return tkv.parent.GetStoreType()
}



func (tkv *Store) CacheWrap() types.CacheWrap {
	panic("cannot CacheWrap a TraceKVStore")
}



func (tkv *Store) CacheWrapWithTrace(_ io.Writer, _ types.TraceContext) types.CacheWrap {
	panic("cannot CacheWrapWithTrace a TraceKVStore")
}


func (tkv *Store) CacheWrapWithListeners(_ types.StoreKey, _ []types.WriteListener) types.CacheWrap {
	panic("cannot CacheWrapWithListeners a TraceKVStore")
}



func writeOperation(w io.Writer, op operation, tc types.TraceContext, key, value []byte) {
	traceOp := traceOperation{
		Operation: op,
		Key:       base64.StdEncoding.EncodeToString(key),
		Value:     base64.StdEncoding.EncodeToString(value),
	}

	if tc != nil {
		traceOp.Metadata = tc
	}

	raw, err := json.Marshal(traceOp)
	if err != nil {
		panic(errors.Wrap(err, "failed to serialize trace operation"))
	}

	if _, err := w.Write(raw); err != nil {
		panic(errors.Wrap(err, "failed to write trace operation"))
	}

	io.WriteString(w, "\n")
}

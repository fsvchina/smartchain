package simulation

import (
	"encoding/json"
	"math/rand"
	"sort"

	"github.com/cosmos/cosmos-sdk/types/simulation"
)


const (
	BeginBlockEntryKind = "begin_block"
	EndBlockEntryKind   = "end_block"
	MsgEntryKind        = "msg"
	QueuedMsgEntryKind  = "queued_msg"
)


type OperationEntry struct {
	EntryKind string          `json:"entry_kind" yaml:"entry_kind"`
	Height    int64           `json:"height" yaml:"height"`
	Order     int64           `json:"order" yaml:"order"`
	Operation json.RawMessage `json:"operation" yaml:"operation"`
}


func NewOperationEntry(entry string, height, order int64, op json.RawMessage) OperationEntry {
	return OperationEntry{
		EntryKind: entry,
		Height:    height,
		Order:     order,
		Operation: op,
	}
}


func BeginBlockEntry(height int64) OperationEntry {
	return NewOperationEntry(BeginBlockEntryKind, height, -1, nil)
}


func EndBlockEntry(height int64) OperationEntry {
	return NewOperationEntry(EndBlockEntryKind, height, -1, nil)
}


func MsgEntry(height, order int64, opMsg simulation.OperationMsg) OperationEntry {
	return NewOperationEntry(MsgEntryKind, height, order, opMsg.MustMarshal())
}


func QueuedMsgEntry(height int64, opMsg simulation.OperationMsg) OperationEntry {
	return NewOperationEntry(QueuedMsgEntryKind, height, -1, opMsg.MustMarshal())
}


func (oe OperationEntry) MustMarshal() json.RawMessage {
	out, err := json.Marshal(oe)
	if err != nil {
		panic(err)
	}

	return out
}


type OperationQueue map[int][]simulation.Operation


func NewOperationQueue() OperationQueue {
	return make(OperationQueue)
}


func queueOperations(queuedOps OperationQueue, queuedTimeOps []simulation.FutureOperation, futureOps []simulation.FutureOperation) {
	if futureOps == nil {
		return
	}

	for _, futureOp := range futureOps {
		futureOp := futureOp
		if futureOp.BlockHeight != 0 {
			if val, ok := queuedOps[futureOp.BlockHeight]; ok {
				queuedOps[futureOp.BlockHeight] = append(val, futureOp.Op)
			} else {
				queuedOps[futureOp.BlockHeight] = []simulation.Operation{futureOp.Op}
			}

			continue
		}



		index := sort.Search(
			len(queuedTimeOps),
			func(i int) bool {
				return queuedTimeOps[i].BlockTime.After(futureOp.BlockTime)
			},
		)

		queuedTimeOps = append(queuedTimeOps, simulation.FutureOperation{})
		copy(queuedTimeOps[index+1:], queuedTimeOps[index:])
		queuedTimeOps[index] = futureOp
	}
}



type WeightedOperation struct {
	weight int
	op     simulation.Operation
}

func (w WeightedOperation) Weight() int {
	return w.weight
}

func (w WeightedOperation) Op() simulation.Operation {
	return w.op
}


func NewWeightedOperation(weight int, op simulation.Operation) WeightedOperation {
	return WeightedOperation{
		weight: weight,
		op:     op,
	}
}


type WeightedOperations []simulation.WeightedOperation

func (ops WeightedOperations) totalWeight() int {
	totalOpWeight := 0
	for _, op := range ops {
		totalOpWeight += op.Weight()
	}

	return totalOpWeight
}

func (ops WeightedOperations) getSelectOpFn() simulation.SelectOpFn {
	totalOpWeight := ops.totalWeight()

	return func(r *rand.Rand) simulation.Operation {
		x := r.Intn(totalOpWeight)
		for i := 0; i < len(ops); i++ {
			if x <= ops[i].Weight() {
				return ops[i].Op()
			}

			x -= ops[i].Weight()
		}

		return ops[0].Op()
	}
}

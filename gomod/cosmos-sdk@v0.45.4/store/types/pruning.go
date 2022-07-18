package types

import "fmt"


const (
	PruningOptionDefault    = "default"
	PruningOptionEverything = "everything"
	PruningOptionNothing    = "nothing"
	PruningOptionCustom     = "custom"
)

var (





	PruneDefault = NewPruningOptions(362880, 100, 10)




	PruneEverything = NewPruningOptions(2, 0, 10)


	PruneNothing = NewPruningOptions(0, 1, 0)
)



type PruningOptions struct {

	KeepRecent uint64


	KeepEvery uint64


	Interval uint64
}

func NewPruningOptions(keepRecent, keepEvery, interval uint64) PruningOptions {
	return PruningOptions{
		KeepRecent: keepRecent,
		KeepEvery:  keepEvery,
		Interval:   interval,
	}
}

func (po PruningOptions) Validate() error {
	if po.KeepEvery == 0 && po.Interval == 0 {
		return fmt.Errorf("invalid 'Interval' when pruning everything: %d", po.Interval)
	}
	if po.KeepEvery == 1 && po.Interval != 0 {
		return fmt.Errorf("invalid 'Interval' when pruning nothing: %d", po.Interval)
	}
	if po.KeepEvery > 1 && po.Interval == 0 {
		return fmt.Errorf("invalid 'Interval' when pruning: %d", po.Interval)
	}

	return nil
}

func NewPruningOptionsFromString(strategy string) PruningOptions {
	switch strategy {
	case PruningOptionEverything:
		return PruneEverything

	case PruningOptionNothing:
		return PruneNothing

	case PruningOptionDefault:
		return PruneDefault

	default:
		return PruneDefault
	}
}

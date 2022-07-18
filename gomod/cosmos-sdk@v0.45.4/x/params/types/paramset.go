package types

type (
	ValueValidatorFn func(value interface{}) error



	ParamSetPair struct {
		Key         []byte
		Value       interface{}
		ValidatorFn ValueValidatorFn
	}
)


func NewParamSetPair(key []byte, value interface{}, vfn ValueValidatorFn) ParamSetPair {
	return ParamSetPair{key, value, vfn}
}


type ParamSetPairs []ParamSetPair


type ParamSet interface {
	ParamSetPairs() ParamSetPairs
}

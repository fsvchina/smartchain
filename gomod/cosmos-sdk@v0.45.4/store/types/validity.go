package types


func AssertValidKey(key []byte) {
	if len(key) == 0 {
		panic("key is nil")
	}
}


func AssertValidValue(value []byte) {
	if value == nil {
		panic("value is nil")
	}
}

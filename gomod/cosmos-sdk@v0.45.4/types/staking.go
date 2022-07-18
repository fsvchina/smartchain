package types


const (

	
	DefaultBondDenom = "stake"

	
	
	
	
	
	//
	
	
	
	ValidatorUpdateDelay int64 = 1
)


var DefaultPowerReduction = NewIntFromUint64(1000000)


func TokensToConsensusPower(tokens Int, powerReduction Int) int64 {
	return (tokens.Quo(powerReduction)).Int64()
}


func TokensFromConsensusPower(power int64, powerReduction Int) Int {
	return NewInt(power).Mul(powerReduction)
}

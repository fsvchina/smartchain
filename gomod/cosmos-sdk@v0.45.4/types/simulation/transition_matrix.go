package simulation

import "math/rand"






type TransitionMatrix interface {
	NextState(r *rand.Rand, i int) int
}

package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/simulation"
)






type TransitionMatrix struct {
	weights [][]int

	totals []int
	n      int
}



func CreateTransitionMatrix(weights [][]int) (simulation.TransitionMatrix, error) {
	n := len(weights)
	for i := 0; i < n; i++ {
		if len(weights[i]) != n {
			return TransitionMatrix{},
				fmt.Errorf("transition matrix: non-square matrix provided, error on row %d", i)
		}
	}

	totals := make([]int, n)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			totals[col] += weights[row][col]
		}
	}

	return TransitionMatrix{weights, totals, n}, nil
}



func (t TransitionMatrix) NextState(r *rand.Rand, i int) int {
	randNum := r.Intn(t.totals[i])
	for row := 0; row < t.n; row++ {
		if randNum < t.weights[row][i] {
			return row
		}

		randNum -= t.weights[row][i]
	}

	return -1
}



func GetMemberOfInitialState(r *rand.Rand, weights []int) int {
	n := len(weights)
	total := 0

	for i := 0; i < n; i++ {
		total += weights[i]
	}

	randNum := r.Intn(total)

	for state := 0; state < n; state++ {
		if randNum < weights[state] {
			return state
		}

		randNum -= weights[state]
	}

	return -1
}

package shuffle

import (
	"math/rand/v2"
	"time"
)

// FisherYates implements the "inside-out" version of the Fisher-Yates shuffle.
func FisherYates[T any](x []T) {
	seed := uint64(time.Now().UnixNano())
	s := rand.NewPCG(seed, seed)
	fisherYates[T](x, rand.New(s))
}

func fisherYates[T any](x []T, rnd *rand.Rand) {
	for i := range x {
		r := rnd.IntN(i + 1)

		x[i], x[r] = x[r], x[i]
	}
}

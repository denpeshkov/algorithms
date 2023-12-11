package shuffle

import "math/rand"

// FisherYates implements the "inside-out" version of the Fisher-Yates shuffle.
func FisherYates[T any](x []T) {
	for i := range x {
		r := rand.Intn(i + 1)

		x[i], x[r] = x[r], x[i]
	}
}

package combinatorial_test

import (
	"fmt"
	"testing"

	"github.com/denpeshkov/algorithms/combinatorial"
)

// genCycles generates cycles of the form [[l-1 -> 0], [l-1 -> 1] ... [l-1 -> l-1]].
func genCycles(l int) [][]int {
	a := make([]int, l)
	var res [][]int

	for i := 0; i < l-1; i++ {
		a[i] = i + 1
	}

	for i := 0; i < l; i++ {
		a[l-1] = i

		t := append([]int(nil), a...)
		res = append(res, t)
	}

	return res
}

func TestFloydCycle(t *testing.T) {
	lengths := []int{1, 2, 3, 4, 5, 7, 9, 11, 13, 19, 20, 50, 100, 101, 1000}

	for _, l := range lengths {
		as := genCycles(l)

		for i, a := range as {
			t.Run(fmt.Sprintf("%d->%d", l-1, i), func(t *testing.T) {
				ind, len := combinatorial.FloydCycle(func(i int) int { return a[i] }, 0)

				if ind != i {
					t.Errorf("ind = %d; want %d", ind, i)
				}
				if len != l-i {
					t.Errorf("len = %d; want %d", len, l)
				}
			})
		}
	}
}

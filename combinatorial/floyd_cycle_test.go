package combinatorial_test

import (
	"testing"

	"github.com/denpeshkov/algorithms/combinatorial"
)

func genCycle(l int) [][]int {
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
	ls := []int{1, 2, 3, 4, 5, 7, 9, 11, 13, 19, 20, 50, 100, 101, 1000}

	for _, l := range ls {
		as := genCycle(l)

		for i, a := range as {
			t.Logf("Iterated function graph %v", a)

			ind, len := combinatorial.FloydCycle(func(i int) int { return a[i] }, 0)

			if ind != i {
				t.Errorf("Expected ind to be: %d; got: %d", i, ind)
			}
			if len != l-i {
				t.Errorf("Expected len to be: %d; got: %d", l, len)
			}
		}
	}
}

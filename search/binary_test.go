package search

import (
	"fmt"
	"sort"
	"testing"

	"golang.org/x/exp/constraints"
)

var data = []int{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000}

func genArr(n int) []int {
	x := make([]int, n)

	for i := 0; i < n; i++ {
		x[i] = i
	}

	return x
}

func cmp[T constraints.Ordered](a, b T) int {
	switch {
	case a == b:
		return 0
	case a < b:
		return -1
	default:
		return 1
	}
}

func TestBinary(t *testing.T) {
	tests := []struct {
		x      []int
		target int
		i      int
	}{
		{genArr(0), 0, 0},
		{genArr(1), 0, 0},
		{genArr(1), 1, 1},
		{genArr(1), -1, 0},
		{genArr(2), 0, 0},
		{genArr(2), 1, 1},
		{genArr(2), -1, 0},
		{genArr(2), 2, 2},
		{genArr(100), 0, 0},
		{genArr(100), 99, 99},
		{genArr(100), 49, 49},
		{genArr(100), 50, 50},
		{genArr(100), 51, 51},
		{genArr(100), 100, 100},
		{genArr(100), 150, 100},
		{genArr(100), -150, 0},
	}

	for _, e := range tests {
		i := Binary(e.x, e.target)

		if i != e.i {
			t.Errorf("expected index %d; got %d", e.i, i)
		}
	}

	for e, v := range data {
		name := fmt.Sprintf("data %v", e)

		i := Binary(data, v)

		if data[i] != v {
			t.Errorf("%s: expected to find %d; found %d", name, v, data[i])
		}
	}
}

func TestBinaryCmp(t *testing.T) {
	tests := []struct {
		x      []int
		target int
		i      int
	}{
		{genArr(0), 0, 0},
		{genArr(1), 0, 0},
		{genArr(1), 1, 1},
		{genArr(1), -1, 0},
		{genArr(2), 0, 0},
		{genArr(2), 1, 1},
		{genArr(2), -1, 0},
		{genArr(2), 2, 2},
		{genArr(100), 0, 0},
		{genArr(100), 99, 99},
		{genArr(100), 49, 49},
		{genArr(100), 50, 50},
		{genArr(100), 51, 51},
		{genArr(100), 100, 100},
		{genArr(100), 150, 100},
		{genArr(100), -150, 0},
	}

	for _, e := range tests {
		i := BinaryCmp(e.x, e.target, cmp[int])

		if i != e.i {
			t.Errorf("expected index %d; got %d", e.i, i)
		}
	}

	for e, v := range data {
		name := fmt.Sprintf("data %v", e)

		i := BinaryCmp(data, v, cmp[int])

		if data[i] != v {
			t.Errorf("%s: expected to find %d; found %d", name, v, data[i])
		}
	}
}

func TestBinaryPredicate(t *testing.T) {
	p := func(x int) func(int) bool {
		return func(v int) bool {
			return v >= x
		}
	}

	tests := []struct {
		x []int
		f func(int) bool
		i int
	}{
		{genArr(0), nil, 0},
		{genArr(1), p(1), 1},
		{genArr(1), func(i int) bool { return true }, 0},
		{genArr(1), func(i int) bool { return false }, 1},
		{genArr(100), p(91), 91},
		{genArr(100), func(i int) bool { return true }, 0},
		{genArr(100), func(i int) bool { return false }, 100},
		{data, p(-20), 0},
		{data, p(-10), 0},
		{data, p(-9), 1},
		{data, p(-6), 1},
		{data, p(-5), 1},
		{data, p(3), 5},
		{data, p(11), 8},
		{data, p(99), 9},
		{data, p(100), 9},
		{data, p(101), 12},
		{data, p(10000), 13},
		{data, p(10001), 14},
		{genArr(7), func(i int) bool { return []int{99, 99, 59, 42, 7, 0, -1, -1}[i] <= 7 }, 4},
		{genArr(100), func(i int) bool { return 100-i <= 7 }, 100 - 7},
	}

	for _, e := range tests {
		i := BinaryPredicate(e.x, e.f)

		if i != e.i {
			t.Errorf("expected index %d; got %d", e.i, i)
		}
	}
}

func FuzzBinary(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, target rune) {
		r := []rune(s)
		l := len(r)

		isValid := func(i int) bool {
			switch {
			case l == 0:
				return i == 0 // empty x
			case i == l: // target is max
				return target > r[l-1]
			case i == 0: // target is min
				return target == r[0] || target < r[0]
			default:
				return target == r[i] || (target > r[i-1] && target < r[i])
			}
		}

		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

		i := Binary(r, target)

		if !isValid(i) {
			t.Errorf("expected index to insert target to be: %d; got: %d", l, i)
		}
	})
}

func FuzzBinaryCmp(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, target rune) {
		r := []rune(s)
		l := len(r)

		isValid := func(i int) bool {
			switch {
			case l == 0:
				return i == 0 // empty x
			case i == l: // target is max
				return target > r[l-1]
			case i == 0: // target is min
				return target == r[0] || target < r[0]
			default:
				return target == r[i] || (target > r[i-1] && target < r[i])
			}
		}

		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

		i := BinaryCmp(r, target, cmp[rune])

		if !isValid(i) {
			t.Errorf("expected index to insert target to be: %d; got: %d", l, i)
		}
	})
}

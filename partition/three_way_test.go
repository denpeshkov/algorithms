package partition_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"

	"github.com/denpeshkov/algorithms/partition"
)

var data = []int{0: 1, 1: 3, 2: -1, 3: 0, 4: 0, 5: 7, 6: 18, 7: 4, 8: 4, 9: 4,
	10: 18, 11: 5, 12: 14, 13: 0, 14: 23, 15: 4, 16: 9, 17: 100, 18: 10000, 19: -34, 20: -56, 21: 3}

func testX(x []int, t *testing.T) {
	tests := []struct {
		name string
		f    func(int) int
		lt   int
		gt   int
	}{
		{"-1", func(int) int { return -1 }, len(x), len(x) - 1},
		{"0", func(int) int { return 0 }, 0, len(x) - 1},
		{"1", func(int) int { return 1 }, 0, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lt, gt := partition.ThreeWay(x, tc.f)

			if !(tc.lt == lt && tc.gt == gt) {
				t.Errorf("data = %v", data)
				t.Errorf("f is %s", tc.name)

				t.Errorf("ThreeWay(data, f) = %v, %v; want %v, %v", lt, gt, tc.lt, tc.gt)
			}
		})
	}
}

func cmp[T constraints.Ordered](a T) func(T) int {
	return func(e T) int {
		switch {
		case e == a:
			return 0
		case e < a:
			return -1
		default:
			return 1
		}
	}
}

func testData(t *testing.T) {
	tests := []struct {
		name string
		f    func(int) int
		lt   int
		gt   int
	}{
		{"cmp(0)", cmp(0), 3, 5},
		{"cmp(4)", cmp(4), 9, 12},
		{"cmp(-56)", cmp(-56), 0, 0},
		{"cmp(10000)", cmp(10000), 21, 21},
		{"cmp(-10001)", cmp(-10001), 0, -1},
		{"cmp(10001)", cmp(10001), 22, 21},
		{"cmp(-1)", cmp(-1), 2, 2},
		{"cmp(100)", cmp(100), 20, 20},
		{"cmp(12)", cmp(12), 16, 15},
		{"cmp(1000)", cmp(1000), 21, 20},
		{"cmp(-50)", cmp(-50), 1, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lt, gt := partition.ThreeWay(data, tc.f)

			if !(tc.lt == lt && tc.gt == gt) {
				t.Errorf("data = %v", data)
				t.Errorf("f is %s", tc.name)

				t.Errorf("ThreeWay(data, f) = %v, %v; want %v, %v", lt, gt, tc.lt, tc.gt)
			}
		})
	}
}

func TestThreeWay(t *testing.T) {
	reverse := func(x []int) []int {
		rev := make([]int, len(x))
		copy(rev, x)

		N := len(rev)

		for i := range rev {
			rev[i], rev[N-i-1] = rev[N-i-1], rev[i]
		}

		return rev
	}

	testX(data, t)
	testX(reverse(data), t)
	testX([]int{}, t)
	testX(nil, t)
	testX([]int{23}, t)
	testX([]int{1, 1001}, t)
	testX([]int{1001, 1}, t)
	testData(t)
}

func FuzzThreeWay(fz *testing.F) {
	testcases := []struct {
		s string
		r rune
	}{
		{"", 'a'},
		{"a", 'a'},
		{"a", 'b'},
		{"aa", 'a'},
		{"aa", 'b'},
		{"ab", 'a'},
		{"ab", 'b'},
	}

	for _, tc := range testcases {
		fz.Add(tc.s, tc.r)
	}

	fz.Fuzz(func(t *testing.T, s string, e rune) {
		x := []rune(s)
		f := cmp(e)

		lt, gt := partition.ThreeWay(x, f)

		for i := 0; i <= lt-1; i++ {
			assert.Negativef(t, f(x[i]), "f(x[%v]) = %v; want < 0", i, f(x[i]))
		}
		for i := lt; i <= gt; i++ {
			assert.Zerof(t, f(x[i]), "f(x[%v]) = %v; want < 0", i, f(x[i]))
		}
		for i := gt + 1; i <= len(x)-1; i++ {
			assert.Positivef(t, f(x[i]), "f(x[%v]) = %v; want > 0", i, f(x[i]))
		}
	})
}

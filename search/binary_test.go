package search

import (
	"cmp"
	"fmt"
	"sort"
	"testing"
)

var data = []int{0: -10, 1: -5, 2: 0, 3: 1, 4: 2, 5: 3, 6: 5, 7: 7, 8: 11, 9: 100, 10: 100, 11: 100, 12: 1000, 13: 10000}

func TestBinary(t *testing.T) {
	tests := map[string]struct {
		x      []int
		target int
		ind    int
	}{
		"[]; 0":         {genArr(0), 0, 0},
		"[0]; 0":        {genArr(1), 0, 0},
		"[0]; 1":        {genArr(1), 1, 1},
		"[0]; 1-":       {genArr(1), -1, 0},
		"[0..1]; 0":     {genArr(2), 0, 0},
		"[0..1]; 1":     {genArr(2), 1, 1},
		"[0..1]; -1":    {genArr(2), -1, 0},
		"[0..1]; 2":     {genArr(2), 2, 2},
		"[0..99]; 0":    {genArr(100), 0, 0},
		"[0..9]; 99":    {genArr(100), 99, 99},
		"[0..99]; 49":   {genArr(100), 49, 49},
		"[0..99]; 50":   {genArr(100), 50, 50},
		"[0..99]; 51":   {genArr(100), 51, 51},
		"[0..99]; 100":  {genArr(100), 100, 100},
		"[0.99]; 150":   {genArr(100), 150, 100},
		"[0..99]; -150": {genArr(100), -150, 0},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ind := Binary(tt.x, tt.target)

			if ind != tt.ind {
				t.Errorf("got %d; want %d", ind, tt.ind)
			}
		})
	}

	for i, v := range data {
		t.Run(fmt.Sprintf("data; %d", v), func(t *testing.T) {
			ind := Binary(data, v)

			if data[ind] != v {
				t.Errorf("got %d; want %d", ind, i)
			}
		})
	}
}

func TestBinaryCmp(t *testing.T) {
	tests := map[string]struct {
		x      []int
		target int
		ind    int
	}{
		"[]; 0":         {genArr(0), 0, 0},
		"[0]; 0":        {genArr(1), 0, 0},
		"[0]; 1":        {genArr(1), 1, 1},
		"[0]; 1-":       {genArr(1), -1, 0},
		"[0..1]; 0":     {genArr(2), 0, 0},
		"[0..1]; 1":     {genArr(2), 1, 1},
		"[0..1]; -1":    {genArr(2), -1, 0},
		"[0..1]; 2":     {genArr(2), 2, 2},
		"[0..99]; 0":    {genArr(100), 0, 0},
		"[0..9]; 99":    {genArr(100), 99, 99},
		"[0..99]; 49":   {genArr(100), 49, 49},
		"[0..99]; 50":   {genArr(100), 50, 50},
		"[0..99]; 51":   {genArr(100), 51, 51},
		"[0..99]; 100":  {genArr(100), 100, 100},
		"[0.99]; 150":   {genArr(100), 150, 100},
		"[0..99]; -150": {genArr(100), -150, 0},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ind := BinaryCmp(tt.x, tt.target, cmp.Compare)

			if ind != tt.ind {
				t.Errorf("got %d; want %d", ind, tt.ind)
			}
		})
	}

	for i, v := range data {
		t.Run(fmt.Sprintf("data; %d", v), func(t *testing.T) {
			ind := BinaryCmp(data, v, cmp.Compare)

			if data[ind] != v {
				t.Errorf("got %d; want %d", ind, i)
			}
		})
	}
}

func TestBinaryPredicate(t *testing.T) {
	p := func(x int) func(int) bool {
		return func(v int) bool {
			return v >= x
		}
	}

	tests := map[string]struct {
		x   []int
		f   func(int) bool
		ind int
	}{
		"[]; nil":        {genArr(0), nil, 0},
		"[0]; >=1":       {genArr(1), p(1), 1},
		"[0]; true":      {genArr(1), func(i int) bool { return true }, 0},
		"[0]; false":     {genArr(1), func(i int) bool { return false }, 1},
		"[0..99]; >=91":  {genArr(100), p(91), 91},
		"[0..99]; true":  {genArr(100), func(i int) bool { return true }, 0},
		"[0..99]; false": {genArr(100), func(i int) bool { return false }, 100},
		"data; >=-20":    {data, p(-20), 0},
		"data; >=-10":    {data, p(-10), 0},
		"data; >=-9":     {data, p(-9), 1},
		"data; >=-6":     {data, p(-6), 1},
		"data; >=-5":     {data, p(-5), 1},
		"data; >=3":      {data, p(3), 5},
		"data; >=11":     {data, p(11), 8},
		"data; >=99":     {data, p(99), 9},
		"data; >=100":    {data, p(100), 9},
		"data; >=101":    {data, p(101), 12},
		"data; >=10000":  {data, p(10000), 13},
		"data; >=10001":  {data, p(10001), 14},
	}

	for _, e := range tests {
		ind := BinaryPredicate(e.x, e.f)

		if ind != e.ind {
			t.Errorf("got %d; want %d", ind, e.ind)
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

		i := BinaryCmp(r, target, cmp.Compare)

		if !isValid(i) {
			t.Errorf("expected index to insert target to be: %d; got: %d", l, i)
		}
	})
}

func genArr(n int) []int {
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = i
	}
	return x
}

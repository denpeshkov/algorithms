// Based on https://github.com/golang/go/blob/master/src/slices/sort_test.go

package sort_test

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
)

var ints = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestSortFuncEmptyNil(t *testing.T) {
	testCases := []struct {
		sortAlg  string
		sortFunc func(s []int, cmp func(a, b int) int)
	}{
		{sortAlg: "insertion sort", sortFunc: sort.InsertionFunc[[]int]},
		{sortAlg: "selection sort", sortFunc: sort.SelectionFunc[[]int]},
		{sortAlg: "merge sort top-down", sortFunc: sort.MergeFunc[[]int]},
		{sortAlg: "merge sort bottom-up", sortFunc: sort.MergeBottomUpFunc[[]int]},
	}

	empty := []int{}
	var nilSlice []int

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s empty slice", tc.sortAlg), func(t *testing.T) {
			t.Parallel()
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("got unexpected panic")
				}
			}()

			tc.sortFunc(empty, cmp.Compare)
		})

		t.Run(fmt.Sprintf("%s nil slice", tc.sortAlg), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("got unexpected panic")
				}
			}()

			tc.sortFunc(nilSlice, cmp.Compare)
		})
	}
}

func TestSortFuncData(t *testing.T) {
	testCases := []struct {
		sortAlg  string
		sortFunc func(s []int, cmp func(a, b int) int)
	}{
		{sortAlg: "insertion sort", sortFunc: sort.InsertionFunc[[]int]},
		{sortAlg: "selection sort", sortFunc: sort.SelectionFunc[[]int]},
		{sortAlg: "merge sort top-down", sortFunc: sort.MergeFunc[[]int]},
		{sortAlg: "merge sort bottom-up", sortFunc: sort.MergeBottomUpFunc[[]int]},
	}

	for _, tc := range testCases {
		t.Run(tc.sortAlg, func(t *testing.T) {
			t.Parallel()
			data := append([]int(nil), ints...)

			tc.sortFunc(data, cmp.Compare)

			if !slices.IsSorted(data) {
				t.Errorf("sorted %v", ints)
				t.Errorf("   got %v", data)
			}
		})
	}
}

func TestSortFuncRandom(t *testing.T) {
	testCases := []struct {
		sortAlg  string
		sortFunc func(s []int, cmp func(a, b int) int)
	}{
		{sortAlg: "insertion sort", sortFunc: sort.InsertionFunc[[]int]},
		{sortAlg: "selection sort", sortFunc: sort.SelectionFunc[[]int]},
		{sortAlg: "merge sort top-down", sortFunc: sort.MergeFunc[[]int]},
		{sortAlg: "merge sort bottom-up", sortFunc: sort.MergeBottomUpFunc[[]int]},
	}

	n := 1000
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000)
	}
	for _, tc := range testCases {
		t.Run(tc.sortAlg, func(t *testing.T) {
			t.Parallel()
			data := append([]int(nil), data...)
			if slices.IsSorted(data) {
				t.Fatalf("terrible rand.rand")
			}

			sort.InsertionFunc(data, cmp.Compare)

			if !slices.IsSorted(data) {
				t.Errorf("sort didn't sort - %d random ints", n)
			}
		})
	}
}

func TestSortFuncReverse(t *testing.T) {
	testCases := []struct {
		sortAlg  string
		sortFunc func(s []int, cmp func(a, b int) int)
	}{
		{sortAlg: "insertion sort", sortFunc: sort.InsertionFunc[[]int]},
		{sortAlg: "selection sort", sortFunc: sort.SelectionFunc[[]int]},
		{sortAlg: "merge sort top-down", sortFunc: sort.MergeFunc[[]int]},
		{sortAlg: "merge sort bottom-up", sortFunc: sort.MergeBottomUpFunc[[]int]},
	}

	data := append([]int(nil), ints...)
	revData := append([]int(nil), ints...)

	for _, tc := range testCases {
		t.Run(tc.sortAlg, func(t *testing.T) {
			t.Parallel()
			data := append([]int(nil), data...)
			revData := append([]int(nil), revData...)

			sort.InsertionFunc(data, cmp.Compare)
			sort.InsertionFunc(revData, reverse[int](cmp.Compare))

			for i := 0; i < len(data); i++ {
				if data[i] != revData[len(data)-1-i] {
					t.Errorf("reverse sort didn't sort")
				}
			}
		})
	}
}

type intPair struct {
	a, b int
}

type intPairs []intPair

// Pairs compare on a only.
func intPairCmp(x, y intPair) int {
	return x.a - y.a
}

// Record initial order in B.
func (d intPairs) initB() {
	for i := range d {
		d[i].b = i
	}
}

// InOrder checks if a-equal elements were not reordered.
func (d intPairs) inOrder() bool {
	lastA, lastB := -1, 0
	for i := 0; i < len(d); i++ {
		if lastA != d[i].a {
			lastA = d[i].a
			lastB = d[i].b
			continue
		}
		if d[i].b <= lastB {
			return false
		}
		lastB = d[i].b
	}
	return true
}

func TestStability(t *testing.T) {
	testCases := []struct {
		sortAlg  string
		sortFunc func(s intPairs, cmp func(a, b intPair) int)
	}{
		{sortAlg: "merge sort top-down", sortFunc: sort.MergeFunc[intPairs]},
		{sortAlg: "merge sort bottom-up", sortFunc: sort.MergeBottomUpFunc[intPairs]},
	}

	n, m := 100000, 1000
	if testing.Short() {
		n, m = 1000, 100
	}

	data := make(intPairs, n)
	// random distribution
	for i := 0; i < len(data); i++ {
		data[i].a = rand.Intn(m)
	}
	if slices.IsSortedFunc(data, intPairCmp) {
		t.Fatalf("terrible rand.rand")
	}

	for _, tc := range testCases {
		t.Run(tc.sortAlg, func(t *testing.T) {
			t.Parallel()
			data := append(intPairs(nil), data...)

			data.initB()
			tc.sortFunc(data, intPairCmp)
			if !slices.IsSortedFunc(data, intPairCmp) {
				t.Errorf("Stable didn't sort %d ints", n)
			}
			if !data.inOrder() {
				t.Errorf("Stable wasn't stable on %d ints", n)
			}

			// already sorted
			data.initB()
			tc.sortFunc(data, intPairCmp)
			if !slices.IsSortedFunc(data, intPairCmp) {
				t.Errorf("Stable shuffled sorted %d ints (order)", n)
			}
			if !data.inOrder() {
				t.Errorf("Stable shuffled sorted %d ints (stability)", n)
			}

			// sorted reversed
			for i := 0; i < len(data); i++ {
				data[i].a = len(data) - i
			}
			data.initB()
			tc.sortFunc(data, intPairCmp)
			if !slices.IsSortedFunc(data, intPairCmp) {
				t.Errorf("Stable didn't sort %d ints", n)
			}
			if !data.inOrder() {
				t.Errorf("Stable wasn't stable on %d ints", n)
			}
		})
	}
}

func BenchmarkSortFunc1K(b *testing.B) {
	benchs := []struct {
		sortAlg  string
		sortFunc func(s []int, cmp func(a, b int) int)
	}{
		{sortAlg: "insertion sort", sortFunc: sort.InsertionFunc[[]int]},
		{sortAlg: "selection sort", sortFunc: sort.SelectionFunc[[]int]},
		{sortAlg: "merge sort top-down", sortFunc: sort.MergeFunc[[]int]},
		{sortAlg: "merge sort bottom-up", sortFunc: sort.MergeBottomUpFunc[[]int]},
	}

	for _, bm := range benchs {
		b.Run(bm.sortAlg, func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, 1<<10)
				for i := 0; i < len(data); i++ {
					data[i] = i ^ 0x2cc
				}
				b.StartTimer()
				bm.sortFunc(data, cmp.Compare)
				b.StopTimer()
			}
		})
	}
}

func FuzzInsertionSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sort.InsertionFunc(s, cmp.Compare)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}

func FuzzSelectionSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sort.SelectionFunc(s, cmp.Compare)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}

func FuzzMergeSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sort.MergeFunc(s, cmp.Compare)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}

func FuzzMergeBottomUpSortFunc(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sort.MergeBottomUpFunc(s, cmp.Compare)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}

func reverse[T cmp.Ordered](fn func(x, y T) int) func(x, y T) int {
	return func(x, y T) int { return fn(y, x) }
}

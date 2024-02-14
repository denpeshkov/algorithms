// Based on https://github.com/golang/go/blob/master/src/sort/sort_test.go

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

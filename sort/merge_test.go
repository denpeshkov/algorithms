package sort_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
)

var mergeFuncData = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestMergeFunc_EmptyNil(t *testing.T) {
	t.Run("TopDown", func(t *testing.T) {
		testSortFuncEmptyNil(t, sort.MergeFunc[[]int], cmp.Compare[int])
	})
	t.Run("BottomUp", func(t *testing.T) {
		testSortFuncEmptyNil(t, sort.MergeBottomUpFunc[[]int], cmp.Compare[int])
	})
}

func TestMergeFunc_Data(t *testing.T) {
	t.Run("TopDown", func(t *testing.T) {
		testSortFuncData(t, sort.MergeFunc[[]int], cmp.Compare[int], mergeFuncData)
	})
	t.Run("BottomUp", func(t *testing.T) {
		testSortFuncData(t, sort.MergeBottomUpFunc[[]int], cmp.Compare[int], mergeFuncData)
	})
}

func TestMergeFunc_Reverse(t *testing.T) {
	t.Run("TopDown", func(t *testing.T) {
		testSortFuncReverse(t, sort.MergeFunc[[]int], cmp.Compare[int], mergeFuncData)
	})
	t.Run("BottomUp", func(t *testing.T) {
		testSortFuncReverse(t, sort.MergeBottomUpFunc[[]int], cmp.Compare[int], mergeFuncData)
	})
}

func TestMergeFunc_RandomInts(t *testing.T) {
	t.Run("TopDown", func(t *testing.T) {
		testSortFuncRandomInts(t, sort.MergeFunc[[]int], cmp.Compare[int])
	})
	t.Run("BottomUp", func(t *testing.T) {
		testSortFuncRandomInts(t, sort.MergeBottomUpFunc[[]int], cmp.Compare[int])
	})
}

func TestMergeFunc_Stability(t *testing.T) {
	n, m := 100000, 1000
	if testing.Short() {
		n, m = 1000, 100
	}
	t.Run("TopDown", func(t *testing.T) {
		testSortFuncStability(t, sort.MergeFunc[intPairs], n, m)
	})
	t.Run("BottomUp", func(t *testing.T) {
		testSortFuncStability(t, sort.MergeBottomUpFunc[intPairs], n, m)
	})
}

func BenchmarkMergeFunc1K(b *testing.B) {
	b.Run("TopDown", func(b *testing.B) {
		benchmarkSortFunc1K(b, sort.MergeFunc[[]int], cmp.Compare[int])
	})
	b.Run("BottomUp", func(b *testing.B) {
		benchmarkSortFunc1K(b, sort.MergeBottomUpFunc[[]int], cmp.Compare[int])
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

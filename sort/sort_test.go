// Based on https://github.com/golang/go/blob/master/src/slices/sort_test.go

package sort_test

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/denpeshkov/algorithms/sort"
)

func testSortFuncEmptyNil[S ~[]E, E any](t *testing.T, sortFunc func(S, func(E, E) int), cmp func(E, E) int) {
	emptySlice := []E{}
	var nilSlice []E

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got unexpected panic")
			}
		}()

		sortFunc(emptySlice, cmp)
	})

	t.Run("nil slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got unexpected panic")
			}
		}()

		sortFunc(nilSlice, cmp)
	})
}

func testSortFuncData[S ~[]E, E any](t *testing.T, sortFunc func(S, func(E, E) int), cmp func(E, E) int, data S) {
	t.Parallel()

	data = append(S(nil), data...)

	sortFunc(data, cmp)

	if !slices.IsSortedFunc(data, cmp) {
		t.Errorf("got unsorted slice: %v, want sorted slice", data)
	}
}

func testSortFuncReverse[S ~[]E, E any](t *testing.T, sortFunc func(S, func(E, E) int), cmp func(E, E) int, data S) {
	t.Parallel()

	data = append(S(nil), data...)
	data1 := append(S(nil), data...)

	sortFunc(data, cmp)
	sortFunc(data1, reverse(cmp))

	for i := 0; i < len(data); i++ {
		if cmp(data[i], data1[len(data)-1-i]) != 0 {
			t.Error("reverse sort didn't sort")
		}
	}
}

func testSortFuncRandomInts(t *testing.T, sortFunc func([]int, func(int, int) int), cmp func(int, int) int) {
	t.Parallel()

	n := 1000
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(1000)
	}
	if slices.IsSorted(data) {
		t.Fatal("terrible rand")
	}

	sort.InsertionFunc(data, cmp)

	if !slices.IsSorted(data) {
		t.Errorf("sort didn't sort - %d random ints", n)
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

func testSortFuncStability(t *testing.T, sortFunc func(intPairs, func(intPair, intPair) int), n, m int) {
	t.Parallel()

	data := make(intPairs, n)
	// random distribution
	for i := 0; i < len(data); i++ {
		data[i].a = rand.Intn(m)
	}
	if slices.IsSortedFunc(data, intPairCmp) {
		t.Fatal("terrible rand")
	}

	data.initB()
	sortFunc(data, intPairCmp)
	if !slices.IsSortedFunc(data, intPairCmp) {
		t.Errorf("Stable didn't sort %d ints", n)
	}
	if !data.inOrder() {
		t.Errorf("Stable wasn't stable on %d ints", n)
	}

	// already sorted
	data.initB()
	sortFunc(data, intPairCmp)
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
	sortFunc(data, intPairCmp)
	if !slices.IsSortedFunc(data, intPairCmp) {
		t.Errorf("Stable didn't sort %d ints", n)
	}
	if !data.inOrder() {
		t.Errorf("Stable wasn't stable on %d ints", n)
	}
}

func benchmarkSortFunc1K(b *testing.B, sortFunc func([]int, func(int, int) int), cmp func(int, int) int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := make([]int, 1<<10)
		for i := 0; i < len(data); i++ {
			data[i] = i ^ 0x2cc
		}
		b.StartTimer()
		sortFunc(data, cmp)
		b.StopTimer()
	}
}

func reverse[T any](fn func(x, y T) int) func(x, y T) int {
	return func(x, y T) int { return fn(y, x) }
}

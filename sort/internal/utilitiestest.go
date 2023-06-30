// Package internal provides common methods for testing sorting algorithms.
// Based on https://github.com/golang/go/blob/master/src/sort/sort_test.go

package internal

import (
	"cmp"
	"math/rand"
	"slices"
	"testing"
)

var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

func TestEmpty(sortingFunc func([]int)) {
	empty := make([]int, 0)
	var nilSlice []int

	sortingFunc(empty)
	sortingFunc(nilSlice)
}

func TestData(sortingFunc func([]int), t *testing.T) {
	var data []int

	copy(data, ints[:])

	sortingFunc(data)

	if !slices.IsSorted(data) {
		t.Errorf("sorted %v", ints)
		t.Errorf("   got %v", data)
	}
}

func TestRandom(n int, sortingFunc func([]int), t *testing.T) {
	data := make([]int, n)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}

	if slices.IsSorted(data) {
		t.Fatalf("terrible rand.rand")
	}

	sortingFunc(data)

	if !slices.IsSorted(data) {
		t.Errorf("sort didn't sort - 1M ints")
	}
}

func TestReverseSort(sortingFunc func(arr []int), t *testing.T) {
	var data, revData []int

	copy(data, ints[:])
	copy(revData, ints[:])
	reverse(revData)

	sortingFunc(data)
	sortingFunc(revData)

	for i := 0; i < len(data); i++ {
		if data[i] != revData[len(data)-1-i] {
			t.Errorf("reverse sort didn't sort")
		}
		if i > len(data)/2 {
			break
		}
	}
}

func FuzzSort(sortingFunc func(s []byte), f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		sortingFunc(s)

		if !slices.IsSortedFunc(s, cmp.Compare) {
			t.Errorf("slice was not sorted")
		}
	})
}

func reverse(x []int) {
	for lo, hi := 0, len(x)-1; lo < hi; lo, hi = lo+1, hi-1 {
		x[lo], x[hi] = x[hi], x[lo]
	}
}

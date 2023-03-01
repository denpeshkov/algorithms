package search

import "golang.org/x/exp/constraints"

/*
Binary implements binary search.
Searches for target in a sorted slice and returns the position where target is found,
or the position where target would appear in the sort order.
The slice must be sorted in increasing order.
*/
func Binary[E constraints.Ordered](x []E, target E) int {
	lo, hi := 0, len(x)-1

	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)

		switch v := x[mid]; {
		case v == target:
			return mid
		case v < target:
			lo = mid + 1
		case v > target:
			hi = mid - 1
		}
	}

	return lo
}

/*
BinaryCmp implements binary search.
It returns the index of the target in the slice if it is present.
If target is not present it returns the index to insert target. The returned index is in interval: [0, len(x)]
If the slice contains multiple elements with the specified value, there is no guarantee which one will be found.
The slice must be sorted in increasing order, defined by cmp.
*/
func BinaryCmp[E, T any](x []E, target T, cmp func(E, T) int) int {
	lo, hi := 0, len(x)-1

	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)

		switch c := cmp(x[mid], target); {
		case c == 0:
			return mid
		case c < 0:
			lo = mid + 1
		case c > 0:
			hi = mid - 1
		}
	}

	return lo
}

/*
BinaryPredicate implements binary search.
It returns the smallest index i in the slice x at which p(x[i]) is true, assuming that, p(x[i]) == true implies p(x[i+1]) == true.
That is, BinaryPredicate requires that p is false for some (possibly empty) prefix of the slice
and then true for the (possibly empty) remainder.
If there is no such index i, BinaryPredicate returns i = n.
*/
func BinaryPredicate[T any](x []T, p func(T) bool) int {
	lo, hi := 0, len(x)

	for lo < hi {
		mid := int(uint(lo+hi) >> 1)

		if p(x[mid]) == true {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

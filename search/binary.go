package search

import "cmp"

/*
BinaryCmp implements binary search. It returns the index of the target in the slice if present.
If the target is not found, it returns the index where the target should be inserted.
The returned index is in the interval [0, len(x)].
If the slice contains multiple elements with the specified value, there is no guarantee which one will be found.
The slice must be sorted in increasing order.
*/
func Binary[E cmp.Ordered](x []E, target E) int {
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
BinaryCmp implements binary search. It returns the index of the target in the slice if present.
If the target is not found, it returns the index where the target should be inserted.
The returned index is in the interval [0, len(x)].
If the slice contains multiple elements with the specified value, there is no guarantee which one will be found.
The slice must be sorted in increasing order, defined by the comparison function cmp.
*/
func BinaryCmp[T any](x []T, target T, cmp func(x, y T) int) int {
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
BinaryPredicate implements binary search. It returns the smallest index i in the slice x at which p(x[i]) is true, assuming that, p(x[i]) == true implies p(x[i+1]) == true.
That is, BinaryPredicate requires that p is false for some (possibly empty) prefix of the slice and then true for the (possibly empty) remainder.
If there is no such index i, BinaryPredicate returns i = n.
*/
func BinaryPredicate[T any](x []T, p func(T) bool) int {
	lo, hi := 0, len(x)

	for lo < hi {
		mid := int(uint(lo+hi) >> 1)

		if p(x[mid]) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

package partition

// ThreeWay reorders the elements in the slice x in such a way that for the element with index i:
//   - f(x[i]) < 0 for i in [0, lt-1];
//   - f(x[i]) == 0 for i in [lt, gt];
//   - f(x[i]) > 0 for i in [gt+1, len(x)-1].
//
// It returns two indexes, lt and gt.

func ThreeWay[T any](x []T, f func(e T) int) (lt, gt int) {
	N := len(x)

	/*
		f(e) < 0: [0, lt-1]
		f(e) == 0: [lt, i-1]
		not yet seen: [i, gt]
		f(e) > 0: [gt+1, len(x)-1]
	*/
	lt, gt = 0, N-1

	for i := 0; i <= gt; {
		switch p := f(x[i]); {
		case p < 0:
			x[lt], x[i] = x[i], x[lt]
			lt++
			i++
		case p == 0:
			i++
		case p > 0:
			x[i], x[gt] = x[gt], x[i]
			gt--
		}
	}

	return lt, gt
}

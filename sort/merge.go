package sort

// MergeFunc implements top-down merge sort using custom comparison function.
func MergeFunc[S ~[]E, E any](s S, cmp func(a, b E) int) {
	aux := make([]E, len(s))

	mergeCmp(s, aux, cmp)
}

func mergeCmp[T any](s []T, aux []T, cmp func(x, y T) int) {
	n := len(s)
	if n <= 1 {
		return
	}

	m := n / 2

	mergeCmp(s[:m], aux, cmp)
	mergeCmp(s[m:], aux, cmp)
	merge(s[:m], s[m:], aux, cmp)

	copy(s, aux)
}

// MergeBottomUpFunc implements bottom-up merge sort using custom comparison function.
func MergeBottomUpFunc[S ~[]E, E any](s S, cmp func(x, y E) int) {
	n := len(s)

	aux := make([]E, len(s))

	for sz := 1; sz < n; sz *= 2 {
		for l := 0; l < n-sz; l += 2 * sz {
			m := l + sz
			r := min(l+2*sz, n)

			merge(s[l:m], s[m:r], aux, cmp)

			copy(s[l:r], aux)
		}
	}
}

// merge two sorted slices s1 and s2 into one sorted slice aux
func merge[S ~[]E, E any](s1, s2 S, aux S, cmp func(x, y E) int) {
	n1, n2 := len(s1), len(s2)
	n := n1 + n2

	for i, j, k := 0, 0, 0; k < n; k++ {
		switch {
		case i >= n1:
			aux[k] = s2[j]
			j++
		case j >= n2:
			aux[k] = s1[i]
			i++
		case cmp(s1[i], s2[j]) <= 0: // ensures stability
			aux[k] = s1[i]
			i++
		default:
			aux[k] = s2[j]
			j++
		}
	}
}

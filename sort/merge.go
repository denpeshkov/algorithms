package sort

// MergeCmp implements merge sort using custom comparison function.
func MergeCmp[T any](s []T, cmp func(x, y T) int) {
	aux := make([]T, len(s))

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

// merge two sorted slices s1 and s2 into one sorted slice aux
func merge[T any](s1, s2 []T, aux []T, cmp func(x, y T) int) {
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
		case cmp(s1[i], s2[j]) <= 0:
			aux[k] = s1[i]
			i++
		default:
			aux[k] = s2[j]
			j++
		}
	}
}

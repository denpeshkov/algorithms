package merge

// BinaryCmp merges two sorted slices s and t into resulting sorted slice using custom comparison function.
func BinaryCmp[T any](s, t []T, cmp func(a, b T) int) []T {
	lenS, lenT := len(s), len(t)
	lenR := lenS + lenT
	r := make([]T, lenR)

	for i, j, k := 0, 0, 0; k < lenR; k++ {
		switch {
		case i >= lenS:
			r[k] = t[j]
			j++
		case j >= lenT:
			r[k] = s[i]
			i++
		case cmp(s[i], t[j]) <= 0:
			r[k] = s[i]
			i++
		default:
			r[k] = t[j]
			j++
		}
	}

	return r
}

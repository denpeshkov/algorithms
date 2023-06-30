package sort

// InsertionCmp implements insertion sort using custom comparison function.
func InsertionCmp[T any](s []T, cmp func(x, y T) int) {
	n := len(s)

	for i := 1; i < n; i++ {
		v := s[i]
		j := i - 1
		for j >= 0 && cmp(s[j], v) > 0 {
			s[j+1] = s[j]
			j--
		}
		s[j+1] = v
	}
}

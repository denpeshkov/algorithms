package sort

// InsertionFunc implements an insertion sort using a custom comparison function.
func InsertionFunc[S ~[]E, E any](s S, cmp func(a, b E) int) {
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

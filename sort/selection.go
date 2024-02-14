package sort

// SelectionFunc implements a selection sort using a custom comparison function.
func SelectionFunc[S ~[]E, E any](s S, cmp func(a, b E) int) {
	n := len(s)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if cmp(s[j], s[min]) < 0 {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
}

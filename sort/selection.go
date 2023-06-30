package sort

// SelectionCmp implements selection sort using custom comparison function.
func SelectionCmp[T any](s []T, cmp func(x, y T) int) {
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

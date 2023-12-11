package misc

// BoyerMooreVoting finds the majority element in a sequence of elements.
// It returns the majority element if one exists; otherwise, it returns an arbitrary element.
// Additionally, it returns a boolean indicating whether the majority element truly exists in the slice.
func BoyerMooreVoting[T comparable](s []T) (major T, found bool) {
	count := 0

	for _, v := range s {
		if count == 0 {
			major = v
		}

		if major == v {
			count++
		} else {
			count--
		}
	}

	// Verify that the element found in the first pass is indeed a majority.
	count = 0
	for _, v := range s {
		if v == major {
			count++
		}
	}

	if count > len(s)/2 {
		return major, true
	}

	return major, false
}

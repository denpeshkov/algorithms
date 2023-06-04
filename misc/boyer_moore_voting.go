// Package misc contains various miscellaneous algorithms
package misc

// BoyerMooreVoting finds the majority of a sequence of elements.
// Returns the majority element if there is one, otherwise returns an arbitrary element;
// it also returns a bool saying whether the majority is really exists in the slice.
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

	// verify that the element found in the first pass really is a majority
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

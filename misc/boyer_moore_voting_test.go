package misc

import (
	"testing"
)

func TestBoyerMooreVotingExists(t *testing.T) {
	testsExists := []struct {
		s         []int
		wantMajor int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 2}, 2},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1, 1, 2, 1, 2, 3, 3, 2, 2, 2, 1, 2, 2, 3, 2, 2}, 2},
		{[]int{2, 2, 3, 1, 4, 5, 2, 2, 3, 2, 2, 2, 1, 2, 1, 2, 2, 1, 2, 2}, 2},
	}

	for _, e := range testsExists {
		major, found := BoyerMooreVoting[int](e.s)

		if !found {
			t.Errorf("expected found true; got %t", found)
		}
		if major != e.wantMajor {
			t.Errorf("expected major %d; got %d", e.wantMajor, major)
		}
	}
}

func TestBoyerMooreVotingNotExists(t *testing.T) {
	testsExists := [][]int{
		{1, 2},
		{1, 2, 3},
		{1, 2, 2, 1},
		{2, 2, 2, 1, 1, 1},
		{1, 3, 4, 3, 1, 1, 1, 5, 10},
	}

	for _, e := range testsExists {
		_, found := BoyerMooreVoting[int](e)

		if found {
			t.Errorf("expected found false; got %t", found)
		}
	}
}

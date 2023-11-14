package test

import "testing"

func compareSlices[E comparable](s1 []E, expected []E, t *testing.T) {
	if len(s1) != len(expected) {
		t.Errorf("wrong size. Expected %d elements and got %d", len(s1), len(expected))
	}

	for i, val := range s1 {
		if val != expected[i] {
			t.Errorf("error in position %d. Expected %v and got %v", i, expected[i], val)
		}
	}
}

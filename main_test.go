package main

import "testing"

func TestFindCelebrity(t *testing.T) {
	// Person 1 is the celebrity.
	a := [][]bool{
		{false, true, false},
		{false, false, false},
		{false, true, false},
	}
	if got := FindCelebrity(a); got != 1 {
		t.Errorf("got %d, want 1", got)
	}

	// No celebrity: 1 knows someone.
	b := [][]bool{
		{false, true},
		{true, false},
	}
	if got := FindCelebrity(b); got != -1 {
		t.Errorf("got %d, want -1", got)
	}

	// Candidate known by nobody: not a celebrity.
	c := [][]bool{
		{false, false},
		{false, false},
	}
	if got := FindCelebrity(c); got != -1 {
		t.Errorf("got %d, want -1 (nobody knows anyone)", got)
	}

	// Single person is trivially the celebrity.
	if got := FindCelebrity([][]bool{{false}}); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
	if got := FindCelebrity([][]bool{}); got != -1 {
		t.Errorf("got %d, want -1 for empty", got)
	}
}

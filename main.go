package main

import "fmt"

// FindCelebrity returns the index of the celebrity — the person everyone
// knows who knows no one — or -1. knows[i][j] says whether i knows j.
// One elimination pass leaves a single candidate; one verify pass confirms.
func FindCelebrity(knows [][]bool) int {
	n := len(knows)
	if n == 0 {
		return -1
	}
	cand := 0
	for i := 1; i < n; i++ {
		if knows[cand][i] {
			cand = i // cand knows someone: not a celebrity; i might be
		}
	}
	for i := 0; i < n; i++ {
		if i == cand {
			continue
		}
		if knows[cand][i] || !knows[i][cand] {
			return -1
		}
	}
	return cand
}

func main() {
	knows := [][]bool{
		{false, true, false},
		{false, false, false},
		{false, true, false},
	}
	fmt.Println(FindCelebrity(knows))
}

package main

import (
	"testing"
)

func TestScoreReport2(t *testing.T) {
	badFirstLevel := []int{4, 1, 2, 3}
	score := ScoreReport2(badFirstLevel)
	if score != 1 {
		t.Fatal("should tolerate an incorrect first level")
	}

	badLastLevel := []int{1, 2, 3, 1}
	score = ScoreReport2(badLastLevel)
	if score != 1 {
		t.Fatal("should tolerate an incorrect last level")
	}

	generallyDecreasing := []int{4, 2, 3, 1}
	score = ScoreReport2(generallyDecreasing)
	if score != 1 {
		t.Fatal("should determine a level is generally decreasing, even if one value is not")
	}

	generallyIncreasing := []int{1, 4, 2, 3}
	score = ScoreReport2(generallyIncreasing)
	if score != 1 {
		t.Fatal("should determine a level is generally increasing, even if one value is not")
	}

	// When evaluating (4, 1): 1 is decreasing, so it needs to be skipped
	skipBadNextLevel := []int{1, 2, 3, 4, 1, 6}
	score = ScoreReport2(skipBadNextLevel)
	if score != 1 {
		t.Fatal("should skip the next level if it's bad")
	}

	// When evaluating (3,6): 6 seems safe. Evaluating (6,5) is not safe (it's decreasing).
	// The issue isn't with the 5, though -- it's that first 6. If we evaluate (3,5) by skipping
	// that 6, we see we have a safe-enough level.
	skipBadCurrentLevel := []int{1, 2, 3, 6, 5, 6}
	score = ScoreReport2(skipBadCurrentLevel)
	if score != 1 {
		t.Fatal("a level might appear safe at first, but evaluating the next level shows it should have been skipped")
	}
}

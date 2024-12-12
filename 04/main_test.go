package main

import (
	"testing"
)

var wordSearch = `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	count := Part1(wordSearch)
	if count != 18 {
		t.Fatalf("expected result %d to equal 18: ", count)
	}
}

func TestPart2(t *testing.T) {
	count := Part2(wordSearch)
	if count != 9 {
		t.Fatalf("expected result %d to equal 9: ", count)
	}
}

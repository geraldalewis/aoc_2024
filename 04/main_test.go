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
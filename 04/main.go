package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(wordSearch string) int {
	lines := strings.Split(strings.TrimSpace(wordSearch), "\n")
	// We'll create a buffer surrounding the word search board filled with 0s,
	// that way we can safely index into other locations without having to do
	// bounds checks. The size of the 0s buffer needs to account for the "MAS"
	// in "XMAS".
	masLength := len("MAS")
	// The length of each row, including the 0 buffers on the left and right:
	stride := masLength + len(lines[0]) + masLength // 000XMASXMAS000
	hpadding := strings.Repeat("0", masLength)
	// The three rows of 0s above the word search board
	vpadding := strings.Repeat("0", stride*masLength)
	// The padded board will look something like this:
	// 0000000000
	// 0000000000
	// 0000000000
	// 000XMAS000
	// 000XMAS000
	// 000XMAS000
	// 0000000000
	// 0000000000
	// 0000000000
	board := vpadding + hpadding + strings.Join(lines, hpadding+hpadding) + hpadding + vpadding

	count := 0
	for i, k := range board {
		if k != 'X' {
			continue
		}
		x := i + stride
		y := i + (stride * 2)
		z := i + (stride * 3)
		if board[i+1] == 'M' && board[i+2] == 'A' && board[i+3] == 'S' { // east
			count++
		}
		if board[x+1] == 'M' && board[y+2] == 'A' && board[z+3] == 'S' { // south east
			count++
		}
		if board[x] == 'M' && board[y] == 'A' && board[z] == 'S' { // south
			count++
		}
		if board[x-1] == 'M' && board[y-2] == 'A' && board[z-3] == 'S' { // south west
			count++
		}
		if board[i-1] == 'M' && board[i-2] == 'A' && board[i-3] == 'S' { // west
			count++
		}
		a := i - (stride * 3)
		b := i - (stride * 2)
		c := i - stride
		if board[c-1] == 'M' && board[b-2] == 'A' && board[a-3] == 'S' { // north west
			count++
		}
		if board[c] == 'M' && board[b] == 'A' && board[a] == 'S' { // north
			count++
		}
		if board[c+1] == 'M' && board[b+2] == 'A' && board[a+3] == 'S' { // north east
			count++
		}
	}
	return count
}

func readInputFile() string {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(fmt.Sprintf("Could not read input file: %v", err))
	}
	return string(data)
}

func main() {
	wordSearch := readInputFile()
	fmt.Printf("Part1 count: %d\n", Part1(wordSearch))
}

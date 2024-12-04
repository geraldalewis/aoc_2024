package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	data, err := os.ReadFile("01.data")
	if err != nil {
		panic("Could not read file")
	}
	var colA []int
	var colB []int
	var str = string(data)
	var trimmed = strings.TrimSpace(str)
	for _, line := range strings.Split(trimmed, "\n") {
		for i, col := range strings.Split(line, "   ") {
			num, err := strconv.Atoi(col)
			if err != nil {
				panic("Could not convert ascii to int")
			}
			if i%2 == 0 {
				colA = append(colA, num)
			} else {
				colB = append(colB, num)
			}
		}
	}
	return colA, colB
}

func sumDiffs(colA []int, colB []int) int {
	if len(colA) != len(colB) {
		panic("Lists of ints aren't the same length")
	}
	total := 0
	for i, num := range colA {
		a, b := num, colB[i]
		if a > b {
			total += a - b
		} else {
			total += b - a
		}
	}
	return total
}

func part1(colA []int, colB []int) int {
	sort.Ints(colA)
	sort.Ints(colB)
	answer := sumDiffs(colA, colB)
	return answer
}

func part2(colA []int, colB []int) int {
	sort.Ints(colA)
	sort.Ints(colB)
	total := 0
	indexA := 0
	indexB := 0
	lenA := len(colA)
	lenB := len(colB)
	numB := colB[indexB]
	for indexA+1 < lenA {
		numA := colA[indexA]
		for numB < numA && indexB+1 < lenB {
			indexB++
			numB = colB[indexB]
		}
		if numA != numB {
			indexA++
			continue
		}
		count := 1
		for indexB+count < lenB && colB[indexB+count] == numB {
			count++
		}
		subtotal := numA * count
		total += subtotal
		for indexA+1 < lenA && colA[indexA+1] == numA {
			indexA++
			total += subtotal
		}
		indexA++
	}
	return total
}

func main() {
	colA, colB := readInput()
	part1Answer := part1(colA, colB)
	part2Answer := part2(colA, colB)
	fmt.Printf("part one: %+v\n", part1Answer)
	fmt.Printf("part two: %+v\n", part2Answer)
}

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

func part1() {
	colA, colB := readInput()
	sort.Ints(colA)
	sort.Ints(colB)
	answer := sumDiffs(colA, colB)
	fmt.Printf("part one: %+v\n", answer)
}

func part2() {
	colA, colB := readInput()
	sort.Ints(colA)
	sort.Ints(colB)
	total := 0
	indexB := 0
	numB := colB[indexB]
	for _, numA := range colA {
		for numB < numA && indexB+1 < len(colB) {
			indexB++
			numB = colB[indexB]
		}
		if numA != numB {
			continue
		}
		count := 1
		for indexB+count < len(colB) && colB[indexB+count] == numB {
			count++
		}
		total += numA * count
	}
	fmt.Printf("part two: %+v\n", total)
}

func main() {
	part1()
	part2()
}

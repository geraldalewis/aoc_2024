package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	none       = iota
	increasing = iota
	decreasing = iota
)

func ScoreReport1(report []int) int {
	return checkReportRange(report, 0, len(report)-1, getDirection(report[0], report[1]), false)
}

func ScoreReport2(report []int) int {
	a, b, c, d := report[0], report[1], report[2], report[3]
	direction := determineDirectionWithDampening(a, b, c, d)
	if direction == none {
		return 0
	}
	if checkReportRange(report, 0, len(report)-1, direction, false) != 0 {
		return 1
	}
	// Simple case: tolerate a bad last level
	if checkReportRange(report, 0, len(report)-2, direction, false) != 0 {
		return 1
	}
	// Simple case: tolerate a bad first level
	if checkReportRange(report, 1, len(report)-1, direction, false) != 0 {
		return 1
	}
	return checkReportRange(report, 0, len(report)-1, direction, true)
}

func checkReportRange(report []int, startIndex int, endIndex int, direction int, enableDampener bool) int {
	hasBadLevel := false
	for i := startIndex; i < endIndex; i++ {
		levelA := report[i]
		levelB := report[i+1]
		if isUnsafe(levelA, levelB, direction) {
			if !enableDampener {
				return 0
			}
			if hasBadLevel {
				return 0
			}
			hasBadLevel = true
			if i > 0 && !isUnsafe(report[i-1], levelB, direction) {
				continue
			}
			if i+2 <= len(report)-1 && isUnsafe(levelA, report[i+2], direction) {
				return 0
			}
			i++
		}
	}
	return 1
}

func determineDirectionWithDampening(a int, b int, c int, d int) int {
	x, y, z := getDirection(a, b), getDirection(b, c), getDirection(c, d)
	if x == y || x == z {
		return x
	}
	if y == z {
		return y
	}
	return none
}

func isUnsafe(x int, y int, direction int) bool {
	return (getDirection(x, y) != direction) || !differsBySafeAmount(x, y)
}

func getDirection(x int, y int) int {
	if x > y {
		return decreasing
	}
	if x < y {
		return increasing
	}
	return none
}

func differsBySafeAmount(x int, y int) bool {
	diff := x - y
	if diff < 0 {
		diff = -diff
	}
	if diff < 1 || diff > 3 {
		return false
	}
	return true
}

func main() {
	data, err := os.ReadFile("input.data")
	if err != nil {
		panic("Could not read file")
	}
	record := strings.TrimSpace(string(data))
	score1 := 0
	score2 := 0
	for _, report := range strings.Split(record, "\n") {
		var levels []int
		for _, col := range strings.Split(report, " ") {
			num, err := strconv.Atoi(col)
			if err != nil {
				panic("Could not convert ascii to int")
			}
			levels = append(levels, num)
		}
		score1 += ScoreReport1(levels)
		score2 += ScoreReport2(levels)
	}
	fmt.Printf("answer part 1: %d\n", score1)
	fmt.Printf("answer part 2: %d\n", score2)
}

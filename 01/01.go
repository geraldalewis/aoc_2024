package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

func absInt(x int) uint {
	bits := *(*uint64)(unsafe.Pointer(&x))
	isNegative := (bits & (1 << 63)) == (1 << 63)
	if isNegative {
		return uint(^x + 1)
	}
	return uint(x)
}

func main() {
	var rawBits uint64 = 0b1111111111111111111111111111111111111111111111111111111111111111

	fmt.Printf("%x\n", rawBits)

	var f int = 42
	bits := *(*uint64)(unsafe.Pointer(&f))
	bitsShifted := bits & (1 << 63)
	isNegative := (bits & (1 << 63)) == (1 << 63)

	fmt.Printf("Raw bits as x: %064b\n", bits)
	fmt.Printf("bitsShifted: %064b\n", bitsShifted)
	fmt.Printf("isNegative %t\n", isNegative)

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
	sort.Ints(colA)
	sort.Ints(colB)
	var diffs []int
	for i, num := range colA {
		// numA := float64(num)
		// numB := float64(colB[i])
		// diff := math.Abs(numA - numB)
		diff := absInt(num - colB[i])
		diffs = append(diffs, int(diff))
	}
	var total = 0
	for _, num := range diffs {
		total += num
	}
	fmt.Printf("%+v\n", total)
}

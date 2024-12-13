package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent2024/util"
)

func safe(ints []int) (int, bool) {
	var diffs []int

	for i := 1; i < len(ints); i++ {
		diff := ints[i] - ints[i-1]
		diffs = append(diffs, diff)
	}

	direction := 0
	for i, diff := range diffs {
		absDiff := diff

		if absDiff < 0 {
			absDiff = -absDiff

			if direction != 0 && direction != -1 {
				return i, false
			}

			direction = -1
		} else if absDiff > 0 {
			if direction != 0 && direction != 1 {
				return i, false
			}

			direction = 1
		}

		if direction == 0 {
			return i, false
		}

		if absDiff < 1 || absDiff > 3 {
			return i, false
		}
	}

	return 0, true
}

func transform(line string) []int {
	var ints []int
	for _, field := range strings.Fields(line) {
		i, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}

func pt1safe(ints []int) int {
	_, isSafe := safe(ints)
	if isSafe {
		return 1
	}

	return 0
}

func pt1() {
	var count int
	for ints := range util.DataProcess(2, transform) {
		count += pt1safe(ints)
	}

	fmt.Println(count)
}

func pt2Safe(ints []int) int {
	firstUnsafeIndex, isSafe := safe(ints)
	if isSafe {
		return 1
	}

	if firstUnsafeIndex >= len(ints) {
		return 0
	}

	newInts := append([]int{}, ints[:firstUnsafeIndex]...)
	newInts = append(newInts, ints[firstUnsafeIndex+1:]...)

	_, isSafeAgain := safe(newInts)
	if isSafeAgain {
		return 1
	}

	newFirstUnsafeIndex := firstUnsafeIndex + 1
	if newFirstUnsafeIndex >= len(ints) {
		return 0
	}

	newIntsAgain := append([]int{}, ints[:newFirstUnsafeIndex]...)
	newIntsAgain = append(newIntsAgain, ints[newFirstUnsafeIndex+1:]...)

	_, isSafeAgainAgain := safe(newIntsAgain)
	if isSafeAgainAgain {
		return 1
	}

	return 0
}

func pt2() {
	var count int
	for ints := range util.DataProcess(2, transform) {
		count += pt2Safe(ints)
	}

	fmt.Println(count)
}

func main() {
	pt2()
}

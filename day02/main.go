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

	fmt.Printf("ints: %+v\ndiffs: %+v\n\n", ints, diffs)

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
	_, isSafe := safe(ints)
	if isSafe {
		return 1
	}

	for i := 0; i < len(ints); i++ {
		newInts := append([]int{}, ints[:i]...)
		newInts = append(newInts, ints[i+1:]...)

		_, isSafe := safe(newInts)
		if isSafe {
			return 1
		}
	}

	return 0
}

func pt2() {
	var i int
	var count int
	for ints := range util.DataProcess(2, transform) {
		fmt.Printf("---------- iteration %d\n", i)
		c := pt2Safe(ints)
		fmt.Printf("safe: %d\n", c)
		count += c
		i++
	}

	fmt.Println(count)
}

func main() {
	pt2()
}

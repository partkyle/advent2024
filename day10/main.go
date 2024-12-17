package main

import (
	"fmt"
	"slices"
	"strconv"

	"advent2024/util"
)

const DAY = 10

func main() {
	grid := slices.Collect(util.Data(DAY))

	var bounds util.Vector
	hikingMap := make(map[complex128]int)
	for y, line := range grid {
		for x, cell := range line {
			hikingMap[complex(float64(x), float64(y))] = util.Must(strconv.Atoi(string(cell)))

			if x > bounds.X {
				bounds.X = x
			}
		}

		if y > bounds.Y {
			bounds.Y = y
		}
	}

	pt2(hikingMap)
}

func pt1(hikingMap map[complex128]int) {
	var total int
	for k, v := range hikingMap {
		if v == 0 {
			trailheads := make(map[complex128]bool)
			traverse(hikingMap, trailheads, k, v)
			total += len(trailheads)
		}
	}
	fmt.Println(total)
}

func pt2(hikingMap map[complex128]int) {
	var total int
	for k, v := range hikingMap {
		if v == 0 {
			total += rating(hikingMap, k, v)
		}
	}
	fmt.Println(total)
}

func traverse(hikingMap map[complex128]int, trailheads map[complex128]bool, k complex128, v int) {
	if v == 9 {
		trailheads[k] = true
	}

	directions := []complex128{
		complex(0, -1),
		complex(0, 1),
		complex(-1, 0),
		complex(1, 0),
	}
	for _, direction := range directions {
		pos := k + direction
		val := hikingMap[pos]
		if val == v+1 {
			traverse(hikingMap, trailheads, pos, v+1)
		}
	}
}

func rating(hikingMap map[complex128]int, k complex128, v int) int {
	if v == 9 {
		return 1
	}

	directions := []complex128{
		complex(0, -1),
		complex(0, 1),
		complex(-1, 0),
		complex(1, 0),
	}

	var total int
	for _, direction := range directions {
		pos := k + direction
		val := hikingMap[pos]
		if val == v+1 {
			total += rating(hikingMap, pos, v+1)
		}
	}

	return total
}

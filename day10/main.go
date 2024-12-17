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
	hikingMap := make(map[util.Vector]int)
	for y, line := range grid {
		for x, cell := range line {
			hikingMap[util.Vector{X: x, Y: y}] = util.Must(strconv.Atoi(string(cell)))

			if x > bounds.X {
				bounds.X = x
			}
		}

		if y > bounds.Y {
			bounds.Y = y
		}
	}

	var total int
	for k, v := range hikingMap {
		if v == 0 {
			trailheads := make(map[util.Vector]bool)
			traverse(hikingMap, trailheads, k, v)
			total += len(trailheads)
		}
	}
	fmt.Println(total)
}

func traverse(hikingMap map[util.Vector]int, trailheads map[util.Vector]bool, k util.Vector, v int) {
	if v == 9 {
		trailheads[k] = true
	}

	directions := []util.Vector{
		{X: 0, Y: -1},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
	}
	for _, direction := range directions {
		pos := k.Add(direction)
		val := hikingMap[pos]
		if val == v+1 {
			traverse(hikingMap, trailheads, pos, v+1)
		}
	}
}

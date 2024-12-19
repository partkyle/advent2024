package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent2024/util"
)

const WIDE = 101
const TALL = 103

type robot struct {
	Position util.Vector[int]
	Velocity util.Vector[int]
}

func (r *robot) Update(wide int, tall int) {
	newPosition := r.Position.Add(r.Velocity)
	newPosition.X = (newPosition.X + wide) % wide
	newPosition.Y = (newPosition.Y + tall) % tall

	r.Position = newPosition
}

func parseRobot(line string) robot {
	parts := strings.Split(line, " ")

	position := strings.Split(parts[0], "=")[1]
	positionParts := strings.Split(position, ",")

	velocity := strings.Split(parts[1], "=")[1]
	velocityParts := strings.Split(velocity, ",")

	return robot{
		Position: util.Vector[int]{
			X: util.Must(strconv.Atoi(positionParts[0])),
			Y: util.Must(strconv.Atoi(positionParts[1])),
		},
		Velocity: util.Vector[int]{
			X: util.Must(strconv.Atoi(velocityParts[0])),
			Y: util.Must(strconv.Atoi(velocityParts[1])),
		},
	}
}

func print(robots []robot, wide int, tall int) {
	points := make(map[util.Vector[int]]int)
	for _, r := range robots {
		points[r.Position]++
	}

	for y := 0; y < tall; y++ {
		for x := 0; x < wide; x++ {
			if points[util.Vector[int]{X: x, Y: y}] > 0 {
				fmt.Printf("%d", points[util.Vector[int]{X: x, Y: y}])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func highestColumn(robots []robot, wide int, tall int) int {
	columns := make(map[int]int)
	for _, robot := range robots {
		columns[robot.Position.X]++
	}

	maxCount := -1
	for _, count := range columns {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func hasBorder(robots []robot, wide int, tall int) bool {
	return false
}

func main() {
	robots := slices.Collect(util.DataProcess(14, parseRobot))

	//wide, tall := 11, 7
	wide, tall := WIDE, TALL

	maxColumn := -1

	for i := 0; i < 8159; i++ {
		for r := range robots {
			robots[r].Update(wide, tall)
		}
	}

	print(robots, wide, tall)

	return

	for i := 0; i < 1000000; i++ {
		for r := range robots {
			robots[r].Update(wide, tall)
		}

		column := highestColumn(robots, wide, tall)
		fmt.Println(i, column)

		if column > maxColumn {
			maxColumn = column
		}

		if column >= 34 {
			print(robots, wide, tall)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}

		//if column == 100 {
		//	fmt.Println("seconds:", i)
		//	print(robots, wide, tall)
		//	break
		//}
	}

	fmt.Println(maxColumn)

	var quadrantTotals [4]int
	for _, robot := range robots {
		x := robot.Position.X
		y := robot.Position.Y

		// 0
		if x < wide/2 && y < tall/2 {
			quadrantTotals[0]++
		}

		// 1
		if x > wide/2 && y < tall/2 {
			quadrantTotals[1]++
		}

		// 2
		if x < wide/2 && y > tall/2 {
			quadrantTotals[2]++
		}

		// 3
		if x > wide/2 && y > tall/2 {
			quadrantTotals[3]++
		}
	}

	fmt.Println(quadrantTotals)
	fmt.Println(quadrantTotals[0] * quadrantTotals[1] * quadrantTotals[2] * quadrantTotals[3])
}

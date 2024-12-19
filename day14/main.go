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
	Position util.Vector
	Velocity util.Vector
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
		Position: util.Vector{
			X: util.Must(strconv.Atoi(positionParts[0])),
			Y: util.Must(strconv.Atoi(positionParts[1])),
		},
		Velocity: util.Vector{
			X: util.Must(strconv.Atoi(velocityParts[0])),
			Y: util.Must(strconv.Atoi(velocityParts[1])),
		},
	}
}

func print(robots []robot, wide int, tall int) {
	points := make(map[util.Vector]int)
	for _, r := range robots {
		points[r.Position]++
	}

	for y := 0; y < tall; y++ {
		for x := 0; x < wide; x++ {
			if points[util.Vector{X: x, Y: y}] > 0 {
				fmt.Printf("%d", points[util.Vector{X: x, Y: y}])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	rounds := 100
	robots := slices.Collect(util.DataProcess(14, parseRobot))

	//wide, tall := 11, 7
	wide, tall := WIDE, TALL

	for i := 0; i < rounds; i++ {
		for i := range robots {
			robots[i].Update(wide, tall)
		}
	}

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

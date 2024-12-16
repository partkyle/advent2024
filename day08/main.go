package main

import (
	"fmt"
	"iter"
	"slices"

	"advent2024/util"
)

const DAY = 8

type Vector struct {
	x, y int
}

func (v Vector) Add(o Vector) Vector {
	return Vector{v.x + o.x, v.y + o.y}
}

func (v Vector) Sub(o Vector) Vector {
	return Vector{v.x - o.x, v.y - o.y}
}

func (v Vector) Within(lo Vector, hi Vector) bool {
	return lo.x <= v.x && v.x < hi.x &&
		lo.y <= v.y && v.y < hi.y
}

func UniquePairs[E any](l []E) iter.Seq[[2]E] {
	return func(yield func([2]E) bool) {
		for i := 0; i < len(l); i++ {
			for j := i + 1; j < len(l); j++ {
				if !yield([2]E{l[i], l[j]}) {
					return
				}
			}
		}
	}
}

func main() {
	antennas := make(map[Vector]rune)
	groupedAntennas := make(map[rune][]Vector)

	grid := slices.Collect(util.Data(DAY))

	bounds := Vector{len(grid[0]), len(grid)}
	for y, line := range grid {
		for x, c := range line {
			if c != '.' && c != '#' {
				antennas[Vector{x, y}] = c
				groupedAntennas[c] = append(groupedAntennas[c], Vector{x, y})
			}
		}
	}

	fmt.Printf("%+v\n", bounds)

	antisignals := make(map[Vector]struct{})
	for _, v := range groupedAntennas {
		for pair := range UniquePairs(v) {
			pointA := pair[0].Sub(pair[1].Sub(pair[0]))
			pointB := pair[1].Sub(pair[0].Sub(pair[1]))

			if pointA.Within(Vector{}, bounds) {
				antisignals[pointA] = struct{}{}
			}
			if pointB.Within(Vector{}, bounds) {
				antisignals[pointB] = struct{}{}
			}
		}
	}

	for y := 0; y < bounds.y; y++ {
		for x := 0; x < bounds.x; x++ {
			if _, ok := antisignals[Vector{x, y}]; ok {
				fmt.Print("#")
			} else if v, ok := antennas[Vector{x, y}]; ok {
				fmt.Print(string(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(len(antisignals))
}

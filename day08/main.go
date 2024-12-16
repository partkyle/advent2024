package main

import (
	"fmt"
	"iter"
	"math"
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

func (v Vector) Simplify() Vector {
	f := gcf(int(math.Abs(float64(v.x))), int(math.Abs(float64(v.y))))
	return Vector{v.x / f, v.y / f}
}

func gcf(a, b int) int {
	af := factors(a)
	bf := factors(b)

	c := make(map[int]struct{})
	for _, a := range af {
		c[a] = struct{}{}
	}

	var v []int
	for _, b := range bf {
		if _, ok := c[b]; ok {
			v = append(v, b)
		}
	}

	return v[0]
}

func factors(a int) []int {
	var f []int
	for i := 1; i <= int(math.Ceil(math.Sqrt(float64(a)))); i++ {
		if a%i == 0 {
			if !slices.Contains(f, i) {
				f = append(f, i)
			}
			if !slices.Contains(f, a/i) {
				f = append(f, a/i)
			}
		}
	}

	slices.Sort(f)
	slices.Reverse(f)
	return f
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

	antisignals := pt2(groupedAntennas, bounds)

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

func pt1(groupedAntennas map[rune][]Vector, bounds Vector) map[Vector]struct{} {
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
	return antisignals
}

func pt2(groupedAntennas map[rune][]Vector, bounds Vector) map[Vector]struct{} {
	antisignals := make(map[Vector]struct{})
	for _, v := range groupedAntennas {
		for pair := range UniquePairs(v) {
			dirA := pair[1].Sub(pair[0]).Simplify()
			dirB := pair[0].Sub(pair[1]).Simplify()

			for pointA := pair[0]; pointA.Within(Vector{}, bounds); pointA = pointA.Sub(dirA) {
				antisignals[pointA] = struct{}{}
			}

			for pointB := pair[0]; pointB.Within(Vector{}, bounds); pointB = pointB.Sub(dirB) {
				antisignals[pointB] = struct{}{}
			}
		}
	}
	return antisignals
}

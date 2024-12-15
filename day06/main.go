package main

import (
	"advent2024/util"
	"fmt"
	"slices"
)

const DAY = 6

type coord struct {
	row, col int
}

func (c coord) add(from coord) coord {
	return coord{c.row + from.row, c.col + from.col}
}

const (
	up int = 1 << (1 * iota)
	down
	left
	right
)

func rotate(d int) int {
	switch d {
	case up:
		return right
	case down:
		return left
	case left:
		return up
	case right:
		return down
	}

	panic("unknown direction")
}

func charToDir(c int32) int {
	if c == '^' {
		return up
	}
	if c == 'v' {
		return down
	}
	if c == '<' {
		return left
	}
	if c == '>' {
		return right
	}

	panic("unknown direction")
}

type Solver struct {
	position coord
	dir      int

	walls map[coord]struct{}
	path  map[coord]int
}

func NewSolver() *Solver {
	return &Solver{
		walls: make(map[coord]struct{}),
		path:  make(map[coord]int),
	}
}

func (s *Solver) solve(grid []string) {
	var bounds coord

	for row, line := range grid {
		for col, c := range line {
			if c == '#' {
				s.walls[coord{row, col}] = struct{}{}
			}
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				s.position = coord{row, col}
				s.dir = charToDir(c)
			}

			if col > bounds.col {
				bounds.col = col
			}
		}

		if row > bounds.row {
			bounds.row = row
		}
	}

	for {
		if s.hasBeen(s.position, s.dir) {
			break
		}

		if s.position.row < 0 || s.position.row > bounds.row || s.position.col < 0 || s.position.col > bounds.col {
			break
		}

		s.goToPath(s.position, s.dir)

		newPosition := s.position.add(coordFrom(s.dir))
		newDir := s.dir

		_, wallFound := s.walls[newPosition]
		for wallFound {
			newDir = rotate(newDir)
			newPosition = s.position.add(coordFrom(newDir))
			_, wallFound = s.walls[newPosition]
		}

		s.position = newPosition
		s.dir = newDir
	}

	fmt.Println(len(s.path))
}

func coordFrom(dir int) coord {
	switch dir {
	case up:
		return coord{row: -1, col: 0}
	case down:
		return coord{row: 1, col: 0}
	case left:
		return coord{row: 0, col: -1}
	case right:
		return coord{row: 0, col: 1}
	}

	panic("unknown direction")
}

func (s *Solver) goToPath(c coord, d int) {
	s.path[c] = s.path[c] | d
}

func (s *Solver) hasBeen(c coord, d int) bool {
	return s.path[c]&d != 0
}

func (s *Solver) printpath(bounds coord) {
	for row := 0; row <= bounds.row; row++ {
		for col := 0; col <= bounds.col; col++ {
			if _, ok := s.walls[coord{row, col}]; ok {
				fmt.Print("#")
			} else if s.path[coord{row, col}] > 0 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func main() {
	grid := slices.Collect(util.Data(DAY))

	s := NewSolver()
	s.solve(grid)
}

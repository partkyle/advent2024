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
	bounds   coord

	walls map[coord]struct{}
	path  map[coord]int
}

func NewSolver(grid []string) *Solver {
	s := &Solver{
		walls: make(map[coord]struct{}),
		path:  make(map[coord]int),
	}

	for row, line := range grid {
		for col, c := range line {
			if c == '#' {
				s.walls[coord{row, col}] = struct{}{}
			}
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				s.position = coord{row, col}
				s.dir = charToDir(c)
			}

			if col > s.bounds.col {
				s.bounds.col = col
			}
		}

		if row > s.bounds.row {
			s.bounds.row = row
		}
	}

	return s
}

func (s *Solver) Copy() *Solver {
	newSolver := &Solver{
		walls:    make(map[coord]struct{}),
		path:     make(map[coord]int),
		dir:      s.dir,
		bounds:   s.bounds,
		position: s.position,
	}

	for k, w := range s.walls {
		newSolver.walls[k] = w
	}

	return newSolver
}

func (s *Solver) solve() (bool, int) {
	for {
		if s.hasBeen(s.position, s.dir) {
			return true, len(s.path)
		}

		if s.position.row < 0 || s.position.row > s.bounds.row || s.position.col < 0 || s.position.col > s.bounds.col {
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

	return false, len(s.path)
}

func pathChar(dir int) string {
	switch dir {
	case up:
		return "╵"
	case down:
		return "╷"
	case left:
		return "╴"
	case right:
		return "╶"
	case up | down:
		return "│"
	case up | down | left:
		return "┤"
	case up | down | right:
		return "├"
	case up | down | left | right:
		return "┼"
	case up | left:
		return "┐"
	case up | right:
		return "┌"
	case up | left | right:
		return "┴"
	case down | left:
		return "┑"
	case down | left | right:
		return "┬"
	case down | right:
		return "┌"
	case left | right:
		return "─"
	}

	return "X"
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
				fmt.Print(pathChar(s.path[coord{row, col}]))
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()
}

func main() {
	grid := slices.Collect(util.Data(DAY))

	s := NewSolver(grid)

	var total int
	for row := 0; row <= s.bounds.row; row++ {
		for col := 0; col <= s.bounds.col; col++ {
			c := s.Copy()

			if _, ok := c.walls[coord{row, col}]; ok {
				fmt.Printf("skipping: %d,%d\n", row, col)
				continue
			}
			// add wall
			c.walls[coord{row, col}] = struct{}{}
			loop, _ := c.solve()
			if loop {
				total++
			}

			fmt.Printf("finished: loop=%v %d,%d\n", loop, row, col)
		}
	}

	fmt.Println(total)
}

package main

import (
	"fmt"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"

	"advent2024/util"
)

type Direction int

func FromRune(char rune) Direction {
	switch char {
	case '^':
		return up
	case 'v':
		return down
	case '<':
		return left
	case '>':
		return right
	}

	return -1
}

func (d Direction) String() string {
	switch d {
	case up:
		return "up"
	case down:
		return "down"
	case left:
		return "left"
	case right:
		return "right"
	}

	return "?"
}

func (d Direction) Vector() util.Vector[int] {
	switch d {
	case up:
		return util.Vector[int]{X: 0, Y: -1}
	case down:
		return util.Vector[int]{X: 0, Y: 1}
	case left:
		return util.Vector[int]{X: -1, Y: 0}
	case right:
		return util.Vector[int]{X: 1, Y: 0}
	}

	return util.Vector[int]{}
}

func (d Direction) Clockwise() Direction {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}

	return -1
}

const (
	up Direction = iota
	down
	left
	right
)

type boundaries struct {
	Lo util.Vector[int]
	Hi util.Vector[int]
}

func (b *boundaries) Update(x, y int) {
	if x > b.Hi.X {
		b.Hi.X = x
	}

	if y > b.Hi.Y {
		b.Hi.Y = y
	}
}

type things struct {
	set map[util.Vector[int]]bool
}

func (t *things) Add(x int, y int) {
	t.AddVector(util.Vector[int]{X: x, Y: y})
}

func (t *things) Has(x int, y int) bool {
	return t.HasVector(util.Vector[int]{X: x, Y: y})
}

func (t *things) HasVector(v util.Vector[int]) bool {
	return t.set[v]
}

func (t *things) Remove(box util.Vector[int]) {
	delete(t.set, box)
}

func (t *things) AddVector(point util.Vector[int]) {
	if t.set == nil {
		t.set = make(map[util.Vector[int]]bool)
	}

	t.set[point] = true
}

func drawStuff(walls things, boxes things, robot util.Vector[int], boundary boundaries) bool {
	if rl.WindowShouldClose() {
		return true
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	height := rl.GetScreenHeight() / (boundary.Hi.Y - boundary.Lo.Y + 1)
	width := rl.GetScreenWidth() / (boundary.Hi.X - boundary.Lo.X + 1)

	for w := range walls.set {
		rl.DrawRectangle(int32(w.X*width), int32(w.Y*height), int32(width), int32(height), rl.LightGray)
	}

	for w := range boxes.set {
		rl.DrawRectangle(int32(w.X*width), int32(w.Y*height), int32(width), int32(height), rl.Brown)
	}

	rl.DrawRectangle(int32(robot.X*width), int32(robot.Y*height), int32(width), int32(height), rl.Magenta)

	rl.EndDrawing()
	return false
}

func main() {
	var walls things
	var boundary boundaries
	var boxes things
	var robot util.Vector[int]

	//width := 800
	//height := 600
	//rl.InitWindow(int32(width), int32(height), "raylib [core] example - basic window")
	//defer rl.CloseWindow()
	//
	//rl.SetTargetFPS(2)

	data := slices.Collect(util.Data(15))
	var y int
	for _, line := range data {
		for x, c := range line {
			if c == '#' {
				walls.Add(x, y)
			}

			if c == 'O' {
				boxes.Add(x, y)
			}

			if c == '@' {
				robot.X = x
				robot.Y = y
			}

			boundary.Update(x, y)
		}

		if line == "" {
			break
		}

		y++
	}

	y++
	var moves []Direction
	for _, line := range data[y:] {
		for _, c := range line {
			moves = append(moves, FromRune(c))
		}
	}

	//printBoard(boundary, walls, boxes, robot)
	//fmt.Println()

	//drawStuff(walls, boxes, robot, boundary)

	for i, dir := range moves {

		if dir == -1 {
			panic("oh no")
		}

		fmt.Printf("%d: moving %s\n", i, dir)

		intent := robot.Add(dir.Vector())
		if walls.Has(intent.X, intent.Y) {
			continue
		} else if boxes.HasVector(intent) {
			var boxesToMove []util.Vector[int]
			boxesToMove = append(boxesToMove, intent)

			pos := intent.Add(dir.Vector())
			for pos.Within(boundary.Lo, boundary.Hi) {
				if walls.HasVector(pos) {
					// found a wall, have to end and can't move anything
					// didn't find an empty space
					boxesToMove = nil
					break
				} else if boxes.HasVector(pos) {
					// add another box
					boxesToMove = append(boxesToMove, pos)
				} else {
					// empty space
					break
				}

				pos = pos.Add(dir.Vector())
			}

			if boxesToMove != nil {
				robot = intent

				for _, box := range boxesToMove {
					boxes.Remove(box)
				}

				for _, box := range boxesToMove {
					boxes.AddVector(box.Add(dir.Vector()))
				}
			}
		} else {
			robot = intent
		}

		//if drawStuff(walls, boxes, robot, boundary) {
		//	return
		//}

		//printBoard(boundary, walls, boxes, robot)
		//fmt.Println()
	}

	var total int
	for box := range boxes.set {
		total += 100*box.Y + box.X
	}

	printBoard(boundary, walls, boxes, robot)
	fmt.Println(total)
}

func printBoard(boundary boundaries, walls things, boxes things, robot util.Vector[int]) {
	for y := boundary.Lo.Y; y <= boundary.Hi.Y; y++ {
		for x := boundary.Lo.X; x <= boundary.Hi.X; x++ {
			var empty = true
			if walls.Has(x, y) {
				fmt.Print("#")
				empty = false
			}
			if boxes.Has(x, y) {
				fmt.Print("0")
				empty = false
			}
			if robot.X == x && robot.Y == y {
				fmt.Print("@")
				empty = false
			}
			if empty {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

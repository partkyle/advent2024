package main

import (
	"fmt"
	"iter"
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

func (t things) Copy() things {
	var newThings things

	for k := range t.set {
		newThings.Add(k.X, k.Y)
	}

	return newThings
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

func drawStuff(state *State, statesRan int, totalMoves int, boundary boundaries) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	offset := util.Vector[int32]{X: 0, Y: 50}
	boardBounds := util.Vector[int]{
		X: (rl.GetScreenWidth() - int(offset.X)) / (boundary.Hi.X - boundary.Lo.X + 1),
		Y: (rl.GetScreenHeight() - int(offset.Y)) / (boundary.Hi.Y - boundary.Lo.Y + 1),
	}

	for w := range state.walls.set {
		rl.DrawRectangle(offset.X+int32(w.X*boardBounds.X), offset.Y+int32(w.Y*boardBounds.Y), int32(boardBounds.X), int32(boardBounds.Y), rl.LightGray)
	}

	for w := range state.boxes.set {
		rl.DrawRectangle(offset.X+int32(w.X*boardBounds.X), offset.Y+int32(w.Y*boardBounds.Y), int32(boardBounds.X), int32(boardBounds.Y), rl.Brown)
	}

	rl.DrawRectangle(offset.X+int32(state.robot.X*boardBounds.X), offset.Y+int32(state.robot.Y*boardBounds.Y), int32(boardBounds.X), int32(boardBounds.Y), rl.Magenta)

	rl.DrawText(fmt.Sprintf("Step: %d", state.step), 4, 4, 22, rl.White)

	rl.DrawText(fmt.Sprintf("States Ran: %d", statesRan-1), 200, 4, 22, rl.White)

	rl.DrawText(fmt.Sprintf("Total Moves: %d", totalMoves), 500, 4, 22, rl.White)

	rl.DrawText(fmt.Sprintf("GPS: %d", getAnswer(state)), 800, 4, 22, rl.White)

	rl.EndDrawing()
}

type State struct {
	step  int
	walls things
	boxes things
	robot util.Vector[int]
}

func (s *State) Copy() *State {
	return &State{
		step:  s.step,
		walls: s.walls.Copy(),
		boxes: s.boxes.Copy(),
		robot: s.robot,
	}
}

func main() {
	width := 1000
	height := 1000
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(width), int32(height), "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(240)

	state, boundary, moves := parseInput(util.Data(15))

	initialState := state.Copy()
	initialState.step = -1

	states := make([]*State, 0, len(moves)+1)
	states = append(states, initialState)

	showState := 0

	go func() {
		for i := 0; i < len(moves); i++ {
			currentStateIdx := i + 1
			prevState := states[currentStateIdx-1]

			nextState := prevState.Copy()
			nextState.step = currentStateIdx
			executeStep(nextState, moves[i], boundary)

			states = append(states, nextState)
		}
	}()

	justGo := false
	for !rl.WindowShouldClose() {
		nextShowState := showState

		if rl.IsKeyPressed(rl.KeySpace) {
			justGo = !justGo
		}
		if justGo {
			nextShowState++
		}

		if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyPressedRepeat(rl.KeyRight) {
			justGo = false
			nextShowState++
		}

		if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressedRepeat(rl.KeyLeft) {
			justGo = false
			nextShowState--
		}

		if rl.IsKeyPressed(rl.KeyUp) {
			justGo = false
			nextShowState = 0
		}

		if rl.IsKeyPressed(rl.KeyDown) {
			justGo = false
			nextShowState = len(states) - 1
		}

		if mouseMove := rl.GetMouseWheelMove(); mouseMove != 0 {
			nextShowState = showState + int(mouseMove)
		}

		if 0 <= nextShowState && nextShowState < len(states) {
			showState = nextShowState
		}

		drawStuff(states[showState], len(states), len(moves), boundary)
	}

	lastState := states[len(states)-1]
	total := getAnswer(lastState)
	printBoard(boundary, lastState)
	fmt.Println(total)
}

func getAnswer(state *State) int {
	var total int
	for box := range state.boxes.set {
		total += 100*box.Y + box.X
	}
	return total
}

func executeStep(state *State, dir Direction, boundary boundaries) {
	intent := state.robot.Add(dir.Vector())
	if state.walls.Has(intent.X, intent.Y) {
		return
	} else if state.boxes.HasVector(intent) {
		var boxesToMove []util.Vector[int]
		boxesToMove = append(boxesToMove, intent)

		pos := intent.Add(dir.Vector())
		for pos.Within(boundary.Lo, boundary.Hi) {
			if state.walls.HasVector(pos) {
				// found a wall, have to end and can't move anything
				// didn't find an empty space
				boxesToMove = nil
				break
			} else if state.boxes.HasVector(pos) {
				// add another box
				boxesToMove = append(boxesToMove, pos)
			} else {
				// empty space
				break
			}

			pos = pos.Add(dir.Vector())
		}

		if boxesToMove != nil {
			state.robot = intent

			for _, box := range boxesToMove {
				state.boxes.Remove(box)
			}

			for _, box := range boxesToMove {
				state.boxes.AddVector(box.Add(dir.Vector()))
			}
		}
	} else {
		state.robot = intent
	}

	return
}

func parseInput(dataSeq iter.Seq[string]) (*State, boundaries, []Direction) {
	var walls things
	var boundary boundaries
	var boxes things
	var robot util.Vector[int]

	data := slices.Collect(dataSeq)
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

	return &State{walls: walls, boxes: boxes, robot: robot}, boundary, moves
}

func printBoard(boundary boundaries, state *State) {
	for y := boundary.Lo.Y; y <= boundary.Hi.Y; y++ {
		for x := boundary.Lo.X; x <= boundary.Hi.X; x++ {
			var empty = true
			if state.walls.Has(x, y) {
				fmt.Print("#")
				empty = false
			}
			if state.boxes.Has(x, y) {
				fmt.Print("0")
				empty = false
			}
			if state.robot.X == x && state.robot.Y == y {
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

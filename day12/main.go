package main

import (
	"fmt"
	"slices"

	"advent2024/util"
)

const DAY = 12

const (
	up = iota
	down
	left
	right
)

var directions = map[int]util.CVec{
	up:    util.NewCvec(0, -1),
	down:  util.NewCvec(0, 1),
	left:  util.NewCvec(-1, 0),
	right: util.NewCvec(1, 0),
}

func grabNeigbors(data map[util.CVec]string, pods map[int][]util.CVec, visit map[util.CVec]bool, p util.CVec, idx int) {
	pod, ok := data[p]
	if !ok {
		return
	}

	pods[idx] = append(pods[idx], p)
	visit[p] = true

	for _, d := range directions {
		neighbor := p + d
		if !visit[neighbor] && data[neighbor] == pod {
			grabNeigbors(data, pods, visit, neighbor, idx)
		}
	}
}

func findArea(plots []util.CVec) int {
	return len(plots)
}

func findPerimeter(plots []util.CVec) int {
	visited := map[util.CVec]bool{}

	if len(plots) == 0 {
		return 0
	}

	return findPerimeterRecursive(plots, visited, plots[0])
}

func findPerimeterRecursive(plots []util.CVec, visited map[util.CVec]bool, p util.CVec) int {
	if visited[p] {
		return 0
	}

	visited[p] = true

	var total int
	for _, d := range directions {
		newP := p + d

		if slices.Contains(plots, newP) {
			total += findPerimeterRecursive(plots, visited, newP)
		} else {
			// count this side since it has no neighbor
			total += 1
		}
	}

	return total
}

type edge struct {
	X int
	Y int

	Xdir int
	Ydir int
}

func (e edge) String() string {
	return fmt.Sprintf("(%d,%d) Xdir=%d Ydir=%d", e.X, e.Y, e.Xdir, e.Ydir)
}

func findEdges(plots []util.CVec) int {

	var edges []edge
	for _, plot := range plots {
		if !slices.Contains(plots, plot+directions[up]) {
			e := edge{X: plot.X(), Y: plot.Y(), Ydir: -1}
			edges = append(edges, e)
		}
		if !slices.Contains(plots, plot+directions[down]) {
			e := edge{X: plot.X(), Y: plot.Y() + 1, Ydir: 1}
			edges = append(edges, e)
		}

		if !slices.Contains(plots, plot+directions[left]) {
			elems := edge{X: plot.X(), Y: plot.Y(), Xdir: -1}
			edges = append(edges, elems)
		}
		if !slices.Contains(plots, plot+directions[right]) {
			elems := edge{X: plot.X() + 1, Y: plot.Y(), Xdir: 1}
			edges = append(edges, elems)
		}
	}

	var removals []int
	for i, ed := range edges {
		var e edge = ed

		if e.Ydir != 0 {
			e.X = ed.X - 1
			e.Y = ed.Y
		} else {
			e.X = ed.X
			e.Y = ed.Y - 1
		}

		if slices.Contains(edges, e) {
			removals = append(removals, i)
		}
	}

	var actualEdges []edge
	for i, ed := range edges {
		if !slices.Contains(removals, i) {
			actualEdges = append(actualEdges, ed)
		}
	}

	return len(actualEdges)
}

func findCorners(plots []util.CVec) int {
	neighbors := make(map[util.CVec]int)
	for _, plot := range plots {

		for i, d := range directions {
			neighbor := plot + d
			if slices.Contains(plots, neighbor) {
				neighbors[plot] |= 1 << i
			}
		}
	}

	corners := []int{
		up | left,
		up | right,
		down | left,
		down | right,
		left,
		right,
		up,
		down,
	}

	var total int
	for _, neighbor := range neighbors {
		if slices.Contains(corners, neighbor) {
			total++
		}
	}

	return total
}

func drawPlot(plots []util.CVec) {
	var lo util.CVec = util.NewCvec(5000, 5000)
	var hi util.CVec

	for _, plot := range plots {
		if plot.X() < lo.X() {
			lo = util.NewCvec(plot.X(), lo.Y())
		}
		if plot.Y() < lo.Y() {
			lo = util.NewCvec(lo.X(), plot.Y())
		}

		if plot.X() > hi.X() {
			hi = util.NewCvec(plot.X(), hi.Y())
		}
		if plot.Y() > hi.Y() {
			hi = util.NewCvec(hi.X(), plot.Y())
		}
	}

	hi += util.NewCvec(1, 1)

	for y := lo.Y(); y < hi.Y(); y++ {
		for x := lo.X(); x < hi.X(); x++ {
			if slices.Contains(plots, util.NewCvec(x, y)) {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	grid := slices.Collect(util.Data(DAY))

	mapData := make(map[util.CVec]string)
	visit := make(map[util.CVec]bool)
	pods := make(map[int][]util.CVec)
	names := make(map[int]string)

	for y, row := range grid {
		for x, c := range row {
			mapData[util.NewCvec(x, y)] = string(c)
		}
	}

	idx := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			p := util.NewCvec(x, y)
			if !visit[p] {
				grabNeigbors(mapData, pods, visit, p, idx)
				names[idx] = mapData[p]
				idx++
			}
		}
	}

	var total int
	var sidesTotal int
	for k, plots := range pods {
		area := findArea(plots)
		perimeter := findPerimeter(plots)
		edges := findEdges(plots)

		fmt.Println(k, names[k], area, perimeter, edges)
		drawPlot(plots)
		total += area * perimeter
		sidesTotal += area * edges
	}

	fmt.Println("pt1", total)
	fmt.Println("pt2", sidesTotal)
}

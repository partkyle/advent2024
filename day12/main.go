package main

import (
	"fmt"
	"slices"

	"advent2024/util"
)

const DAY = 12

var directions = []util.CVec{
	util.NewCvec(0, -1),
	util.NewCvec(0, 1),
	util.NewCvec(-1, 0),
	util.NewCvec(1, 0),
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

	fmt.Println(pods)

	var total int
	for k, plots := range pods {
		area := findArea(plots)
		perimeter := findPerimeter(plots)

		fmt.Println(k, names[k], area, perimeter)
		total += area * perimeter
	}

	fmt.Println(total)
}

package main

import (
	"fmt"
	"slices"

	"advent2024/util"
)

func getAt(grid []string, row int, col int) uint8 {
	if row < 0 || row >= len(grid) {
		return 0
	}

	if col < 0 || col >= len(grid[row]) {
		return 0
	}

	return grid[row][col]
}

type data struct {
	coord []int
	char  uint8
}

func up(grid []string, row int, col int) int {
	var upValues = []data{
		{[]int{-1, 0}, 'M'},
		{[]int{-2, 0}, 'A'},
		{[]int{-3, 0}, 'S'},
	}
	for _, v := range upValues {
		if getAt(grid, row+v.coord[0], col+v.coord[1]) != v.char {
			return 0
		}
	}
	return 1
}

func down(grid []string, row int, col int) uint8 {
	return 0
}

func search(grid []string, row int, col int, direction []data) int {
	for _, dir := range direction {
		if getAt(grid, row+dir.coord[0], col+dir.coord[1]) != dir.char {
			return 0
		}
	}
	return 1
}

func pt1() {
	grid := slices.Collect(util.Data(4))

	var cords [][]int
	for row, line := range grid {
		for col, c := range line {
			if c == 'X' {
				cords = append(cords, []int{row, col})
			}
		}
	}

	allDirections := [][]data{
		// up
		{
			{[]int{-1, 0}, 'M'},
			{[]int{-2, 0}, 'A'},
			{[]int{-3, 0}, 'S'},
		},
		// down
		{
			{[]int{1, 0}, 'M'},
			{[]int{2, 0}, 'A'},
			{[]int{3, 0}, 'S'},
		},
		// left
		{
			{[]int{0, -1}, 'M'},
			{[]int{0, -2}, 'A'},
			{[]int{0, -3}, 'S'},
		},
		// right
		{
			{[]int{0, 1}, 'M'},
			{[]int{0, 2}, 'A'},
			{[]int{0, 3}, 'S'},
		},
		// up left
		{
			{[]int{-1, -1}, 'M'},
			{[]int{-2, -2}, 'A'},
			{[]int{-3, -3}, 'S'},
		},
		// up right
		{
			{[]int{-1, 1}, 'M'},
			{[]int{-2, 2}, 'A'},
			{[]int{-3, 3}, 'S'},
		},
		// down right
		{
			{[]int{1, 1}, 'M'},
			{[]int{2, 2}, 'A'},
			{[]int{3, 3}, 'S'},
		},
		// down left
		{
			{[]int{1, -1}, 'M'},
			{[]int{2, -2}, 'A'},
			{[]int{3, -3}, 'S'},
		},
	}

	var total int
	for _, cords := range cords {
		row, col := cords[0], cords[1]
		for _, direction := range allDirections {
			total += search(grid, row, col, direction)
		}
	}

	fmt.Println(total)
}

func rot(lines []string) []string {
	newArray := make([]string, len(lines[0]))

	for _, line := range lines {
		for ci, c := range line {
			newArray[ci] = string(c) + newArray[ci]
		}
	}

	return newArray
}

func pt2() {

}

func main() {
	pt1()
}

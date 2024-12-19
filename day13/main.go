package main

import (
	"fmt"
	"math"
	"slices"

	"advent2024/util"
)

/*
	Gaussian Elimination found in https://www.geeksforgeeks.org/gaussian-elimination/
*/

type chunk struct {
	X int
	Y int
}

type puzzle struct {
	A     chunk
	B     chunk
	Prize chunk
}

func main() {
	puzzles := processPuzzles()

	util.PrettyJSON(puzzles)

	var total int
	for _, p := range puzzles {
		answer := findAnswer(p)
		if answer != nil {
			total += *answer
		}
	}

	fmt.Println(total)
}

func findAnswer(p puzzle) *int {
	var solutions []int
	for a := 0; a <= 100; a++ {
		for b := 0; b <= 100; b++ {
			x := a*p.A.X + b*p.B.X
			y := a*p.A.Y + b*p.B.Y

			if x == p.Prize.X && y == p.Prize.Y {
				solutions = append(solutions, 3*a+b)
			}
		}
	}

	slices.Sort(solutions)

	if len(solutions) > 0 {
		return &solutions[0]
	}

	return nil
}

func gaussianElimination(mat [][]float64) []float64 {
	singular := forwardElim(mat)

	if singular != -1 {
		if mat[singular][len(mat)] != 0 {
			panic("inconsistent system")
		} else {
			panic("may have infinitely many solutions")
		}
	}

	return backSub(mat)
}

func backSub(mat [][]float64) []float64 {
	x := make([]float64, len(mat))

	for i := len(mat) - 1; i >= 0; i-- {
		x[i] = mat[i][len(mat)]

		for j := i + 1; j < len(mat); j++ {
			x[i] -= mat[i][j] * x[j]
		}

		x[i] = x[i] / mat[i][i]
	}

	return x
}

func forwardElim(mat [][]float64) int {
	for k := 0; k < len(mat); k++ {
		iMax := k
		vMax := mat[iMax][k]

		for i := k + 1; i < len(mat); i++ {
			if math.Abs(mat[i][k]) > vMax {
				vMax = mat[i][k]
				iMax = i
			}
		}

		if mat[k][iMax] == 0 {
			return k
		}

		if iMax != k {
			swapRow(mat, k, iMax)
		}

		for i := k + 1; i < len(mat); i++ {
			f := mat[i][k] / mat[k][k]

			for j := k + 1; j <= len(mat); j++ {
				mat[i][j] -= mat[k][j] * f
			}

			mat[i][k] = 0
		}
	}

	return -1
}

func swapRow(mat [][]float64, i int, j int) {
	mat[i], mat[j] = mat[j], mat[i]
}

func printMatrix(mat [][]float64) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println()
	}
}

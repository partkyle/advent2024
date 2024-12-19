package main

import (
	"fmt"
	"math"
	"slices"
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

	pt2(puzzles)
}

func pt1(puzzles []puzzle) {
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

func pt2(puzzles []puzzle) {
	for i := range puzzles {
		puzzles[i].Prize.X += 10000000000000
		puzzles[i].Prize.Y += 10000000000000
	}

	var total int
	for _, p := range puzzles {
		answers := gaussianElimination([][]float64{
			{float64(p.A.X), float64(p.B.X), float64(p.Prize.X)},
			{float64(p.A.Y), float64(p.B.Y), float64(p.Prize.Y)},
		})

		if answers != nil {
			fmt.Println(answers)

			total += validateAnswer(p, answers)
		}
	}

	fmt.Println(total)
}

func validateAnswer(p puzzle, answer []float64) int {
	var solutions []int

	for a := int(math.Floor(answer[0])); a < int(math.Ceil(answer[0]))+1; a++ {
		for b := int(math.Floor(answer[1])); b < int(math.Ceil(answer[1]))+1; b++ {
			x := a*p.A.X + b*p.B.X
			y := a*p.A.Y + b*p.B.Y

			if x == p.Prize.X && y == p.Prize.Y {
				solutions = append(solutions, 3*a+b)
			}
		}
	}

	slices.Sort(solutions)

	if len(solutions) > 0 {
		return solutions[0]
	}

	return 0
}

func gaussianElimination(mat [][]float64) []float64 {
	singular := forwardElim(mat)

	if singular != -1 {
		if mat[singular][len(mat)] != 0 {
			fmt.Println("inconsistent system")
			return nil
		} else {
			fmt.Println("may have infinitely many solutions")
			return nil
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

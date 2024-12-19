package main

import (
	"fmt"
	"math"
)

func main() {
	gaussianElimination([][]float64{
		{3, 2, -4, 3},
		{2, 3, 3, 15},
		{5, -3, 1, 14},
	})
}

func gaussianElimination(mat [][]float64) {
	singular := forwardElim(mat)

	if singular != -1 {
		fmt.Println("singular matrix")
		if mat[singular][len(mat[0])-1] != 0 {
			fmt.Println("inconsistent system")
		} else {
			fmt.Println("may have infinitely many solutions")
		}
	} else {
		backSub(mat)
	}
}

func backSub(mat [][]float64) {
	x := make([]float64, len(mat))

	for i := len(mat) - 1; i >= 0; i-- {
		if mat[i][i] == 0 {
			fmt.Println("Error: Zero pivot encountered!")
			return
		}
		x[i] = mat[i][len(mat)]

		for j := i + 1; j < len(mat); j++ {
			x[i] -= mat[i][j] * x[j]
		}

		x[i] = x[i] / mat[i][i]
	}

	fmt.Println("Solution for system:")
	for i := 0; i < len(mat); i++ {
		fmt.Print(x[i], " ")
	}
	fmt.Println()
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

		if math.Abs(mat[iMax][k]) < 1e-10 { // Checking for near-zero pivot
			return k // Singular matrix
		}

		if iMax != k {
			swapRow(mat, k, iMax)
		}

		for i := k + 1; i < len(mat); i++ {
			f := mat[i][k] / mat[k][k]

			for j := k + 1; j < len(mat[0]); j++ { // Adjusted for correct column length
				mat[i][j] -= mat[k][j] * f
			}

			mat[i][k] = 0 // Eliminate the lower part of the column
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

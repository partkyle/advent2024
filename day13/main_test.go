package main

import (
	"fmt"
	"testing"
)

func Test_swapRow(t *testing.T) {

	matrix := [][]float64{
		{94, 22, 8400},
		{34, 67, 5400},
	}

	printMatrix(matrix)
	swapRow(matrix, 0, 1)

	fmt.Println("after")
	printMatrix(matrix)
}

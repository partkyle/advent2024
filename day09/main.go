package main

import (
	"advent2024/util"
	"fmt"
	"slices"
	"strconv"
)

const DAY = 9

func main() {
	var file int
	var fileFree bool = true

	var data []int
	for line := range util.Data(DAY) {
		for _, c := range line {
			n := util.Must(strconv.Atoi(string(c)))
			if fileFree {
				for i := 0; i < n; i++ {
					data = append(data, file)
				}
				file++
			} else {
				for i := 0; i < n; i++ {
					data = append(data, -1)
				}
			}

			fileFree = !fileFree
		}
	}

	var idx int
	for i := len(data) - 1; i >= 0; i-- {
		if idx > len(data)-1 {
			break
		}

		if data[i] != -1 {
			for data[idx] != -1 && idx < len(data)-1 {
				idx++
			}

			if idx >= i {
				break
			}

			data[idx] = data[i]
			data[i] = -1
		}
	}

	firstEmpty := slices.Index(data, -1)

	var total int
	for i := 0; i < firstEmpty; i++ {
		total += data[i] * i
	}

	fmt.Println(total)
}

func printData(data []int) {
	for _, i := range data {
		if i == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"advent2024/util"
)

func day1() error {
	left, right, err := day1Parse()
	if err != nil {
		return err
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum int
	for i := 0; i < len(left); i++ {
		dist := left[i] - right[i]
		if dist < 0 {
			dist = -dist
		}
		sum += dist
	}

	fmt.Println(sum)

	return nil
}
func transformLine(line string) [2]int {
	row := strings.Fields(line)

	l, err := strconv.Atoi(row[0])
	if err != nil {
		panic(err)
	}

	r, err := strconv.Atoi(row[1])
	if err != nil {
		panic(err)
	}

	return [2]int{l, r}
}

func day1Parse() ([]int, []int, error) {
	var left []int
	var right []int

	for row := range util.DataProcess(1, transformLine) {
		l, r := row[0], row[1]
		left = append(left, l)
		right = append(right, r)
	}

	return left, right, nil
}

func day1pt2() error {
	left, right, err := day1Parse()
	if err != nil {
		return err
	}

	counts := make(map[int]int)
	for _, r := range right {
		counts[r]++
	}

	var total int
	for _, l := range left {
		fmt.Println(l, counts[l])
		total += l * counts[l]
	}

	fmt.Println(total)

	return nil
}

func main() {
	err := day1pt2()
	if err != nil {
		log.Fatal(err)
	}
}

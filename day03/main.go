package main

import (
	"fmt"
	"regexp"
	"strconv"

	"advent2024/util"
)

var regex = regexp.MustCompile(`mul\((?P<left>\d+),(?P<right>\d+)\)`)

func pt1() {
	var total int
	for row := range util.Data(3) {
		match := regex.FindAllStringSubmatch(row, -1)

		for _, m := range match {
			left, err := strconv.Atoi(m[regex.SubexpIndex("left")])
			if err != nil {
				panic(fmt.Errorf("failed to parse %d: %w", m[regex.SubexpIndex("left")], err))
			}

			right, err := strconv.Atoi(m[regex.SubexpIndex("right")])
			if err != nil {
				panic(fmt.Errorf("failed to parse %d: %w", m[regex.SubexpIndex("right")], err))
			}

			total += left * right
		}
	}

	fmt.Println(total)
}

func pt2() {

}

func main() {
	pt1()
}

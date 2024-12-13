package main

import (
	"fmt"
	"regexp"
	"strconv"

	"advent2024/util"
)

var regex = regexp.MustCompile(`(?P<do>do\(\))|(?P<dont>don't\(\))|(?P<mul>mul\((?P<left>\d+),(?P<right>\d+)\))`)

func mul(m []string) int {
	left, err := strconv.Atoi(m[regex.SubexpIndex("left")])
	if err != nil {
		panic(fmt.Errorf("failed to parse %s: %w", m[regex.SubexpIndex("left")], err))
	}

	right, err := strconv.Atoi(m[regex.SubexpIndex("right")])
	if err != nil {
		panic(fmt.Errorf("failed to parse %s: %w", m[regex.SubexpIndex("right")], err))
	}

	return left * right
}

func pt1() {
	var total int
	for row := range util.Data(3) {
		match := regex.FindAllStringSubmatch(row, -1)

		for _, m := range match {
			if m[regex.SubexpIndex("mul")] == "" {
				continue
			}
			total += mul(m)
		}
	}

	fmt.Println(total)
}

func pt2() {
	enabled := true
	var total int
	for row := range util.Data(3) {
		match := regex.FindAllStringSubmatch(row, -1)

		for _, m := range match {
			if m[regex.SubexpIndex("do")] != "" {
				enabled = true
			}

			if m[regex.SubexpIndex("dont")] != "" {
				enabled = false
			}

			if enabled && m[regex.SubexpIndex("mul")] != "" {
				total += mul(m)
			}
		}
	}

	fmt.Println(total)
}

func main() {
	pt2()
}

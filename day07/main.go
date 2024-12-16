package main

import (
	"advent2024/util"
	"fmt"
	"strconv"
	"strings"
)

const DAY = 7

func permutations(count int) [][]string {
	var results [][]string

	perms := 1 << count
	for i := 0; i < perms; i++ {
		n := i
		var row []string
		for j := 0; j < count; j++ {
			if n%2 == 0 {
				row = append(row, "+")
			} else {
				row = append(row, "*")
			}
			n /= 2
		}
		results = append(results, row)
	}

	return results
}

type record struct {
	answer int
	ops    []int
}

func (r record) isPossible() bool {
	for _, allOperators := range permutations(len(r.ops) - 1) {
		val := r.ops[0]
		for i, operator := range allOperators {
			if operator == "*" {
				val = val * r.ops[i+1]
			} else if operator == "+" {
				val = val + r.ops[i+1]
			}
		}

		if val == r.answer {
			return true
		}
	}

	return false
}

func parseLine(line string) record {
	var err error

	answerOperands := strings.Split(line, ": ")

	var r record

	r.answer, err = strconv.Atoi(answerOperands[0])
	if err != nil {
		panic(err)
	}

	for _, o := range strings.Split(answerOperands[1], " ") {
		i, err := strconv.Atoi(o)
		if err != nil {
			panic(err)
		}

		r.ops = append(r.ops, i)
	}

	return r
}

func main() {
	var total int
	for row := range util.DataProcess(DAY, parseLine) {
		if row.isPossible() {
			total += row.answer
		}
	}
	fmt.Println(total)
}

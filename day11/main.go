package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"advent2024/util"
)

const DAY = 11

type Stone struct {
	Value int
}

type List struct {
	util.List[Stone]
}

func (l *List) Print() {
	for curr := l.Head; curr != nil; curr = curr.Next {
		fmt.Print(curr.Data.Value, " ")
	}
	fmt.Println()
}

func (l *List) Blink() {
	for node := l.Head; node != nil; node = node.Next {
		if node.Data.Value == 0 {
			node.Data.Value = 1
			continue
		}

		digits := ToDigits(node.Data.Value)
		if len(digits)%2 == 0 {
			l.InsertBefore(node, Stone{Value: ToInt(digits[0 : len(digits)/2])})
			l.InsertBefore(node, Stone{Value: ToInt(digits[len(digits)/2:])})
			l.Remove(node)
		} else { // odd
			node.Data.Value *= 2024
		}
	}
}

func ToDigits(i int) []int {
	var result []int
	for i > 0 {
		result = append(result, i%10)
		i /= 10
	}

	slices.Reverse(result)
	return result
}

func ToInt(d []int) int {
	var result int
	for i, v := range d {
		result += int(math.Pow(10, float64(len(d)-1-i)) * float64(v))
	}
	return result
}

func transform(line string) []int {
	parts := strings.Split(line, " ")

	result := make([]int, len(parts))

	for i, part := range parts {
		result[i] = util.Must(strconv.Atoi(part))
	}

	return result
}

func main() {
	pt2()
}

func pt1() {
	var list List
	for line := range util.DataProcess(DAY, transform) {
		for _, val := range line {
			list.InsertEnd(Stone{Value: val})
		}
	}

	for i := 0; i < 75; i++ {
		list.Blink()
		fmt.Println(i, list.Count())
	}

	fmt.Println(list.Count())
}

type internalStruct struct {
	val   int
	times int
}

type memoized struct {
	data map[internalStruct]int
}

func newMemoized() *memoized {
	return &memoized{
		data: make(map[internalStruct]int),
	}
}

func (m memoized) HasValue(val int, times int) bool {
	_, ok := m.data[internalStruct{val: val, times: times}]
	return ok
}

func (m memoized) GetValue(val int, times int) int {
	v, _ := m.data[internalStruct{val: val, times: times}]
	return v
}

func (m memoized) SetValue(val int, times int, total int) {
	m.data[internalStruct{val: val, times: times}] = total
}

func blinkValue(m *memoized, val int, times int) int {
	if m.HasValue(val, times) {
		return m.GetValue(val, times)
	}

	if times == 0 {
		return 1
	}

	var stones []int
	if val == 0 {
		stones = append(stones, 1)
	} else {
		digits := ToDigits(val)
		if len(digits)%2 == 0 {
			stones = append(stones, ToInt(digits[0:len(digits)/2]))
			stones = append(stones, ToInt(digits[len(digits)/2:]))
		} else { // odd
			stones = append(stones, val*2024)
		}
	}

	var total int
	for _, stone := range stones {
		total += blinkValue(m, stone, times-1)
	}

	m.SetValue(val, times, total)

	return total
}

func pt2() {
	var stones []int
	for line := range util.DataProcess(DAY, transform) {
		for _, val := range line {
			stones = append(stones, val)
		}
	}

	fmt.Println(stones)

	memo := newMemoized()
	var total int
	for _, stone := range stones {
		total += blinkValue(memo, stone, 75)
	}
	fmt.Println(total)
}

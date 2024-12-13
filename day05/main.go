package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"advent2024/util"
)

const DAY = 5

func parseOrDie(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

type Graph[E comparable] struct {
	Edges map[E][]E
}

func NewGraph[E comparable]() *Graph[E] {
	return &Graph[E]{
		Edges: make(map[E][]E),
	}
}

func (g *Graph[E]) addEdge(from E, to E) {
	g.Edges[from] = append(g.Edges[from], to)
}

func (g *Graph[E]) hasEdge(from E, to E) bool {
	return slices.Contains(g.Edges[from], to)
}

func (g *Graph[E]) hasAnyEdges(from E) bool {
	return len(g.Edges[from]) > 0
}

func main() {
	lines := slices.Collect(util.Data(DAY))
	emptyLine := slices.Index(lines, "")
	rules := lines[:emptyLine]
	queue := lines[emptyLine+1:]

	g := NewGraph[int]()

	for _, rule := range rules {
		parts := strings.Split(rule, "|")

		from := parseOrDie(parts[0])
		to := parseOrDie(parts[1])

		g.addEdge(to, from)
	}

	util.PrettyJSON(g)

	var queuePages [][]int
	for _, line := range queue {
		parts := strings.Split(line, ",")

		var pages []int
		for _, part := range parts {
			page := parseOrDie(part)
			pages = append(pages, page)
		}

		queuePages = append(queuePages, pages)
	}

	pt2(queuePages, g)
}

func pt2(queuePages [][]int, g *Graph[int]) {
	var total int
	for _, page := range queuePages {
		firstTimeWinner := true
		for {
			idx := findFirstBadPage(page, g)
			if idx == -1 {
				break
			}

			firstTimeWinner = false

			fixIdx := findGreatedIndexOfBadPage(page, idx, g)

			if fixIdx == -1 {
				panic("this didn't work")
			}

			newPages := append([]int{}, page[:idx]...)
			newPages = append(newPages, page[idx+1:fixIdx+1]...)
			newPages = append(newPages, page[idx])
			newPages = append(newPages, page[fixIdx+1:]...)
			page = newPages
		}

		if !firstTimeWinner {
			val := page[len(page)/2]
			total += val
		}
	}

	fmt.Println(total)
}

func findGreatedIndexOfBadPage(page []int, idx int, g *Graph[int]) int {
	val := page[idx]

	maxIdx := -1
	for _, edge := range g.Edges[val] {
		newIdx := slices.Index(page, edge)
		if newIdx > maxIdx {
			maxIdx = newIdx
		}
	}

	return maxIdx
}

func pt1(queuePages [][]int, g *Graph[int]) {
	var total int

	for _, page := range queuePages {
		if findFirstBadPage(page, g) == -1 {
			val := page[len(page)/2]
			total += val
		}
	}

	fmt.Println(total)
}

func findFirstBadPage(pages []int, g *Graph[int]) int {
	var handledPages []int

	for i, page := range pages {
		// check if it has prereqs
		if g.hasAnyEdges(page) {
			prereqs := g.Edges[page]

			pagesICareAbout := intersect(pages, prereqs)
			all := containsAll(handledPages, pagesICareAbout)

			if len(pagesICareAbout) > 0 && !all {
				return i
			}
		}

		handledPages = append(handledPages, page)
	}

	return -1
}

func containsAll[E comparable](haystack []E, needles []E) bool {
	for _, needle := range needles {
		if !slices.Contains(haystack, needle) {
			return false
		}
	}

	return true
}

func containsAny[E comparable](haystack []E, needles []E) bool {
	for _, needle := range needles {
		if slices.Contains(haystack, needle) {
			return true
		}
	}

	return false
}

func intersect[E comparable](a []E, b []E) []E {
	var result []E
	for _, v := range b {
		if slices.Contains(a, v) {
			result = append(result, v)
		}
	}
	return result
}

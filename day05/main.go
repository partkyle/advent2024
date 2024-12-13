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

	util.PrettyJSON(g.Edges)

	var total int
	for _, line := range queue {
		parts := strings.Split(line, ",")

		var pages []int
		for _, part := range parts {
			page := parseOrDie(part)
			pages = append(pages, page)
		}

		if pageOrderSafe(pages, g) {
			val := pages[len(pages)/2]
			total += val
		}

	}

	fmt.Println(total)
}

func pageOrderSafe(pages []int, g *Graph[int]) bool {
	var handledPages []int

	for _, page := range pages {
		// check if it has prereqs
		if g.hasAnyEdges(page) {
			prereqs := g.Edges[page]

			pagesICareAbout := intersect(pages, prereqs)
			all := containsAll(handledPages, pagesICareAbout)

			if len(pagesICareAbout) > 0 && !all {
				return false
			}
		}

		handledPages = append(handledPages, page)
	}

	return true
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

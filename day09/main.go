package main

import (
	"fmt"
	"strconv"

	"advent2024/util"
)

const DAY = 9

const EmptyFile = -1

func main() {
	pt2()
}

func pt1() {
	data := extractData()

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

	total := checksum(data)
	fmt.Println(total)
}

func checksum(data []int) int {
	var total int
	for i := 0; i < len(data); i++ {
		if data[i] != -1 {
			total += data[i] * i
		}
	}
	return total
}

func extractData() []int {
	var file int
	fileFree := true

	var data []int
	for line := range util.Data(DAY) {
		for _, c := range line {
			n := util.Must(strconv.Atoi(string(c)))
			if fileFree { // files
				for i := 0; i < n; i++ {
					data = append(data, file)
				}
				file++
			} else { // empty space
				for i := 0; i < n; i++ {
					data = append(data, EmptyFile)
				}
			}

			fileFree = !fileFree
		}
	}
	return data
}

type Node struct {
	ID   int
	Size int
}

type List struct {
	util.List[Node]
}

func (l *List) Cleanup() {
	if l.Head == nil {
		return
	}

	prev := l.Head
	node := prev.Next
	for node != nil {
		if prev.Data.ID == EmptyFile && node.Data.ID == EmptyFile {
			l.Remove(node)
			prev.Data.Size += node.Data.Size
		} else {
			prev = node
		}

		node = node.Next
	}
}

func pt2() {
	var file int
	fileFree := true

	var list List
	for line := range util.Data(DAY) {
		for _, c := range line {
			n := util.Must(strconv.Atoi(string(c)))
			ID := EmptyFile
			if fileFree { // files
				ID = file
				file++
			}

			list.InsertEnd(
				Node{
					ID:   ID,
					Size: n,
				},
			)

			fileFree = !fileFree
		}
	}

	list.Print()
	process(&list)
	list.Print()

	var total int
	var pos int
	for curr := list.Head; curr != nil; curr = curr.Next {
		if curr.Data.ID != EmptyFile {
			for i := 0; i < curr.Data.Size; i++ {
				total += (pos + i) * curr.Data.ID
			}
		}
		pos += curr.Data.Size
	}

	fmt.Println(total)
}

func process(list *List) {
	curr := list.Tail
	for curr != nil {
		prev := curr.Prev

		if curr.Data.ID != EmptyFile {
			for node := list.Head; node != curr; node = node.Next {
				if node.Data.ID == EmptyFile && curr.Data.Size <= node.Data.Size {
					list.InsertBefore(curr, Node{
						ID:   EmptyFile,
						Size: curr.Data.Size,
					})
					list.Remove(curr)
					list.InsertBefore(node, curr.Data)
					node.Data.Size -= curr.Data.Size
					break
				}
			}

			list.Cleanup()
		}

		curr = prev
	}
}

func (l *List) Print() {
	for curr := l.Head; curr != nil; curr = curr.Next {
		var c string
		if curr.Data.ID == EmptyFile {
			c = "."
		} else {
			c = fmt.Sprintf("%d", curr.Data.ID)
		}

		for i := 0; i < curr.Data.Size; i++ {
			fmt.Print(c)
		}
	}
	fmt.Println()
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

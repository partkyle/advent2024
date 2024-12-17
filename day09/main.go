package main

import (
	"advent2024/util"
	"fmt"
	"strconv"
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

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) InsertBeginning(newNode *Node) {
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Prev = nil
		newNode.Next = nil
	} else {
		l.InsertBefore(l.Head, newNode)
	}
}

func (l *List) InsertBefore(node *Node, newNode *Node) {
	newNode.Next = node
	if node.Prev == nil {
		newNode.Prev = nil
		l.Head = newNode
	} else {
		newNode.Prev = node.Prev
		node.Prev.Next = newNode
	}
	node.Prev = newNode
}

func (l *List) InsertAfter(node *Node, newNode *Node) {
	newNode.Prev = node
	if node.Next == nil {
		newNode.Next = nil
		l.Tail = newNode
	} else {
		newNode.Next = node.Next
		node.Next.Prev = newNode
	}
	node.Next = newNode
}

func (l *List) InsertEnd(newNode *Node) {
	if l.Tail == nil {
		l.InsertBeginning(newNode)
	} else {
		l.InsertAfter(newNode, l.Tail)
	}
}

func (l *List) Remove(node *Node) {
	if node.Prev == nil {
		l.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		l.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
}

func (l *List) Cleanup() {
	if l.Head == nil {
		return
	}

	prev := l.Head
	node := prev.Next
	for node != nil {
		if prev.ID == EmptyFile && node.ID == EmptyFile {
			l.Remove(node)
			prev.Size += node.Size
		} else {
			prev = node
		}

		node = node.Next
	}
}

type Node struct {
	Next *Node
	Prev *Node

	ID   int
	Size int
}

func pt2() {
	var file int
	fileFree := true

	var list List
	var prev *Node
	for line := range util.Data(DAY) {
		for _, c := range line {
			n := util.Must(strconv.Atoi(string(c)))
			ID := EmptyFile
			if fileFree { // files
				ID = file
				file++
			}

			node := &Node{
				Prev: prev,
				ID:   ID,
				Size: n,
			}

			// TODO: can this be done without branching
			if list.Head == nil {
				list.Head = node
				prev = list.Head
			} else {
				prev.Next = node
				prev = node
			}

			fileFree = !fileFree
		}
	}

	list.Tail = prev

	printList(list.Head)

	process(&list)

	printList(list.Head)

	var total int
	var pos int
	for curr := list.Head; curr != nil; curr = curr.Next {
		if curr.ID != EmptyFile {
			for i := 0; i < curr.Size; i++ {
				total += (pos + i) * curr.ID
			}
		}
		pos += curr.Size
	}

	fmt.Println(total)
}

func process(list *List) {
	curr := list.Tail
	for curr != nil {
		prev := curr.Prev

		if curr.ID != EmptyFile {
			for node := list.Head; node != curr; node = node.Next {
				if node.ID == EmptyFile && curr.Size <= node.Size {
					list.InsertBefore(curr, &Node{
						ID:   EmptyFile,
						Size: curr.Size,
					})
					list.Remove(curr)
					list.InsertBefore(node, curr)
					node.Size -= curr.Size
					break
				}
			}

			list.Cleanup()
		}

		curr = prev
	}
}

func maxCompression(list *List) {
	for curr := list.Head; curr != nil; curr = curr.Next {
		printList(list.Head)
		if curr.ID == EmptyFile {
			fmt.Printf("filling empty file size=%d\n", curr.Size)

			// start from the tail
			node := list.Tail
			for curr.Size > 0 && node != nil && node != curr {
				prev := node.Prev

				if node.ID != EmptyFile && node.Size <= curr.Size {
					list.InsertBefore(node, &Node{
						ID:   EmptyFile,
						Size: node.Size,
					})
					list.Remove(node)
					list.InsertBefore(curr, node)
					curr.Size -= node.Size
				}

				node = prev
			}

			// clean up contiguous empty spaces
			{
				prev := curr
				node := prev.Next
				for node != nil {
					if prev.ID == EmptyFile && node.ID == EmptyFile {
						list.Remove(node)
						prev.Size += node.Size
					}

					prev = node
					node = node.Next
				}
			}
		}
	}
}

func printList(head *Node) {
	for curr := head; curr != nil; curr = curr.Next {
		var c string
		if curr.ID == EmptyFile {
			c = "."
		} else {
			c = fmt.Sprintf("%d", curr.ID)
		}

		for i := 0; i < curr.Size; i++ {
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

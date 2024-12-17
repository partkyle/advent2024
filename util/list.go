package util

type ListNode[E any] struct {
	Prev *ListNode[E]
	Next *ListNode[E]

	Data E
}

type List[E any] struct {
	Head *ListNode[E]
	Tail *ListNode[E]
}

func (l *List[E]) Count() int {
	var count int
	for node := l.Head; node != nil; node = node.Next {
		count++
	}
	return count
}

func (l *List[E]) InsertBeginning(data E) *ListNode[E] {
	if l.Head == nil {
		newNode := &ListNode[E]{
			Data: data,
		}
		l.Head = newNode
		l.Tail = newNode
		newNode.Prev = nil
		newNode.Next = nil
		return newNode
	} else {
		return l.InsertBefore(l.Head, data)
	}
}

func (l *List[E]) InsertBefore(node *ListNode[E], data E) *ListNode[E] {
	newNode := &ListNode[E]{
		Data: data,
	}
	newNode.Next = node
	if node.Prev == nil {
		newNode.Prev = nil
		l.Head = newNode
	} else {
		newNode.Prev = node.Prev
		node.Prev.Next = newNode
	}
	node.Prev = newNode

	return newNode
}

func (l *List[E]) InsertAfter(node *ListNode[E], data E) *ListNode[E] {
	newNode := &ListNode[E]{
		Data: data,
	}

	newNode.Prev = node
	if node.Next == nil {
		newNode.Next = nil
		l.Tail = newNode
	} else {
		newNode.Next = node.Next
		node.Next.Prev = newNode
	}
	node.Next = newNode

	return newNode
}

func (l *List[E]) InsertEnd(data E) *ListNode[E] {
	if l.Tail == nil {
		return l.InsertBeginning(data)
	} else {
		return l.InsertAfter(l.Tail, data)
	}
}

func (l *List[E]) Remove(node *ListNode[E]) {
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

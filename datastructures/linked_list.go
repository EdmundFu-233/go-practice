package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(val int) {
	node := &Node{Value: val}
	if l.Head == nil {
		l.Head = node
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
}

func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Print()
}

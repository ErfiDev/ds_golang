package main

import (
	"fmt"
)

/*
	LinkedList is a sequence of nodes that have properties
	and a reference to the next node in the sequence. It is
	a linear data structure that is used to store data.
*/

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) addToHead(v int) {
	n := &Node{value: v}

	if l.head != nil {
		n.next = l.head
	}

	l.head = n
}

func (l *LinkedList) iterateAll() {
	for n := l.head; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}

func (l *LinkedList) lastNode() *Node {
	var last *Node

	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			last = n
			break
		}
	}

	return last
}

func (l *LinkedList) addToEnd(v int) {
	n := &Node{value: v}
	n.next = nil

	lastNode := l.lastNode()
	if lastNode != nil {
		lastNode.next = n
	}
}

func (l *LinkedList) findNode(v int) *Node {
	var node *Node

	for n := l.head; n != nil; n = n.next {
		if n.value == v {
			node = n
			break
		}
	}

	return node
}

func (l *LinkedList) addAfter(v int, after int) {
	n := &Node{value: v}
	n.next = nil

	findedNode := l.findNode(after)

	if findedNode != nil {
		n.next = findedNode.next
		findedNode.next = n
	}
}

func main() {
	l := new(LinkedList)

	l.addToHead(1)
	l.addToHead(3)
	l.addToHead(4)
	l.addToEnd(7)

	l.addAfter(2, 1)

	l.iterateAll()
}

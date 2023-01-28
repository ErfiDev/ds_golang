package main

import "fmt"

/*
	each node is connected to two nodes. and we can
	traverse forward through to the next node or
	back through to the previous node.
*/

type DoubleLinkedList struct {
	head *Node
}

type Node struct {
	value    int
	next     *Node
	previous *Node
}

func (l *DoubleLinkedList) addToHead(v int) {
	n := &Node{value: v}
	n.next = nil
	n.previous = nil

	if l.head != nil {
		l.head.previous = n
		n.next = l.head
	}

	l.head = n
}

func (l *DoubleLinkedList) addAfter(v int, after int) {
	node := &Node{value: v}
	node.next = nil
	node.previous = nil

	findedNode := l.findNode(after)

	if findedNode != nil {
		node.previous = findedNode
		node.next = findedNode.next
		findedNode.next.previous = node
		findedNode.next = node
	}
}

func (l *DoubleLinkedList) findNode(v int) *Node {
	var node *Node
	for n := l.head; n != nil; n = n.next {
		if n.value == v {
			node = n
			break
		}
	}

	return node
}

func (l *DoubleLinkedList) iterateAll() {
	for n := l.head; n != nil; n = n.next {
		fmt.Println("------------------------------")
		if n.previous != nil {
			fmt.Printf("previous node value: %d\n", n.previous.value)
		}
		fmt.Printf("node value: %d\n", n.value)
		if n.next != nil {
			fmt.Printf("next node value: %d\n", n.next.value)
		}
	}
}

func (l *DoubleLinkedList) addToEnd(v int) {
	n := &Node{value: v}
	n.next = nil
	n.previous = nil

	lastNode := l.lastNode()

	if lastNode != nil {
		n.previous = lastNode
		lastNode.next = n
	}
}

func (l *DoubleLinkedList) lastNode() *Node {
	var node *Node

	for n := l.head; n != nil; n = n.next {
		if n.next == nil {
			node = n
			break
		}
	}

	return node
}

func main() {
	dll := DoubleLinkedList{}

	dll.addToHead(1)
	dll.addToEnd(2)
	dll.addToEnd(3)
	dll.addToHead(5)
	dll.addAfter(7, 1)

	dll.iterateAll()
}

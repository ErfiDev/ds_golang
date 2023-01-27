package main

import "fmt"

/*
	Lists is a collection of ordered elements that are used
	to store list items, can expand and shrink dynamically.
*/

type List[T comparable] struct {
	l []T
}

func (l *List[T]) push(v T) {
	l.l = append(l.l, v)
}

func (l *List[T]) pushTop(v T) {
	l.l = append(l.l, v)

	copy := l.l

	copy[0], copy[len(copy)-1] = copy[len(copy)-1], copy[0]

	j := len(copy) - 1
	for i := range copy[1:] {
		copy[i+1], copy[j] = copy[j], copy[i+1]
	}

	l.l = copy
}

func (l *List[T]) pop() {
	if len(l.l) > 1 {
		l.l = l.l[:len(l.l)-1]
		return
	}

	l.l = []T{}
}

func (l *List[T]) popTop() {
	if len(l.l) > 1 {
		l.l = l.l[1:len(l.l)]
		return
	}

	l.l = []T{}
}

func main() {
	list := new(List[string])

	list.push("john")
	list.push("jay")
	fmt.Println(list.l)

	list.pushTop("alex")
	fmt.Println(list.l)

	list.pop()
	fmt.Println(list.l)

	list.popTop()
	fmt.Println(list.l)
}

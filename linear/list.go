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
	copy := l.l

	if len(copy) == 0 {
		copy = append(copy, v)
		l.l = copy
		return
	}

	copy = append(copy[:0], append([]T{v}, copy[0:]...)...)

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

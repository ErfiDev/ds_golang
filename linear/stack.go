package main

import "fmt"

/*
	A stack is a last in, first out structure in which
	items are added from the top.
*/

type Stack struct {
	items []Elm
}

type Elm struct {
	value int
}

func (s *Stack) push(e Elm) {
	s.items = append(s.items[:0], append([]Elm{e}, s.items[0:]...)...)
}

func (s *Stack) pop() *Elm {
	if len(s.items) == 0 {
		return nil
	}

	elm := s.items[0]
	s.items = s.items[1:]

	return &elm
}

func main() {
	stack := Stack{}

	stack.push(Elm{value: 10})
	stack.push(Elm{value: 20})
	stack.push(Elm{value: 30})
	stack.push(Elm{value: 40})

	fmt.Println(stack.items)

	stack.pop()

	fmt.Println(stack.items)
}

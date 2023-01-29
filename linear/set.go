package main

import "fmt"

/*
	a set is a linear data structure that has a collection
	of values that are not repeated.
*/

type Set struct {
	m map[string]bool
}

func newSet() Set {
	return Set{
		m: make(map[string]bool),
	}
}

func (s *Set) add(v string) {
	if !s.exists(v) {
		s.m[v] = true
	}
}

func (s *Set) exists(v string) bool {
	_, con := s.m[v]

	return con
}

func (s *Set) delete(v string) {
	delete(s.m, v)
}

func (s *Set) intersec(another Set) Set {
	interSection := newSet()

	for val := range s.m {
		if another.exists(val) {
			interSection.add(val)
		}
	}

	return interSection
}

func (s *Set) union(another Set) Set {
	union := newSet()

	for val := range s.m {
		union.add(val)
	}

	for val := range another.m {
		union.add(val)
	}

	return union
}

func main() {
	s1 := newSet()

	s1.add("john")
	s1.add("leo")
	s1.add("rob")

	s2 := newSet()

	s2.add("david")
	s2.add("leo")
	s2.add("robert")

	union := s1.union(s2)
	intersection := s1.intersec(s2)

	fmt.Println(union)
	fmt.Println(intersection)
}

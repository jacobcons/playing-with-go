package main

import (
	"fmt"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) String() string {
	s := ""
	p := l
	for p != nil {
		s += fmt.Sprintf("%v ", p.val)
		p = p.next
	}
	return s
}

func (l *List[T]) append(x T) {
	p := l
	for p.next != nil {
		p = p.next
	}
	p.next = &List[T]{nil, x}
}

func main() {
	l := &List[int]{&List[int]{&List[int]{nil, 1}, 2}, 3}
	fmt.Println(l)
	l.append(4)
	fmt.Println(l)
}

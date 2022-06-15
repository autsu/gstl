package stack

import (
	"github.com/youseebiggirl/gstl/list"
)

type stack[T any] struct {
	l *list.List[T]
}

func New[T any]() *stack[T] {
	return &stack[T]{l: list.New[T]()}
}

func (s *stack[T]) Push(v T) {
	s.l.PushBack(v)
}

func (s *stack[T]) Pop() T {
	if s.Empty() {
		panic("pop: stack is empty")
	}
	return s.l.Remove(s.l.Back())
}

func (s *stack[T]) Top() T {
	return s.l.Back().Value
}

func (s *stack[T]) Empty() bool {
	return s.l.Len() == 0
}

func (s *stack[T]) Len() int {
	return s.l.Len()
}

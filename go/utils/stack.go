package utils

import "fmt"

type AllowedTypes interface {
	rune | string | int
}

type Stack[T AllowedTypes] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.elements) <= 0 {
		return zero, fmt.Errorf("empty stack")
	} else {
		e := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return e, nil
	}
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.elements) <= 0 {
		return zero, fmt.Errorf("empty stack")
	} else {
		e := s.elements[len(s.elements)-1]
		return e, nil
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

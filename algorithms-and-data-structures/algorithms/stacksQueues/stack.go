package main

import (
	"errors"
)

type Stack[T any] struct {
	arr []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(elem T) {
	s.arr = append(s.arr, elem)
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}

var ErrEmptyStack = errors.New("stack is empty")

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.Size() == 0 {
		return zero, ErrEmptyStack
	}
	last := s.Size() - 1
	popped := s.arr[last]
	s.arr[last] = zero
	s.arr = s.arr[:last]
	return popped, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.Size() == 0 {
		return zero, ErrEmptyStack
	}
	return s.arr[s.Size()-1], nil
}

package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements [10]int
	top      int
}

func NewStack() *Stack {
	return &Stack{top: -1}
}

func (s *Stack) Push(element int) error {

	if s.top >= len(s.elements)-1 {
		return errors.New("Stack overflow!")
	}

	s.top++
	s.elements[s.top] = element
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.top <= 0 {
		return 0, errors.New("Stack underflow!")
	}

	element := s.elements[s.top]
	s.top--
	return element, nil
}

func (s *Stack) Peek() (int, error) {
	if s.top < 0 {
		return 0, errors.New("Stack underflow!")
	}

	return s.elements[s.top], nil
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Size() int {
	return s.top + 1
}

func (s *Stack) Clear() {
	s.top = -1
}

func (s *Stack) Print() {
	for i := range s.elements {
		if i > s.top {
			break
		}
		message := fmt.Sprintf("[%d]", s.elements[i])
		fmt.Print(message)
	}
}

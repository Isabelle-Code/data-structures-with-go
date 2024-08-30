package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackPush(t *testing.T) {
	s := NewStack()
	err := s.Push(1)
	assert.NoError(t, err)
	err = s.Push(2)
	assert.NoError(t, err)
	assert.Equal(t, 2, s.Size(), "Stack size should be 2")
}

func TestStackOverflow(t *testing.T) {
	s := NewStack()

	for i := range s.elements {
		err := s.Push(i)
		assert.NoError(t, err)
	}

	err := s.Push(10)
	assert.Error(t, err, "Stack overflow!")
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	element, err := s.Pop()
	assert.NoError(t, err, "Pop should not return an error")
	assert.Equal(t, 30, element, "Pop should return 30")

}

func TestStackUnderflow(t *testing.T) {
	s := NewStack()
	_, err := s.Pop()
	assert.Error(t, err, "Stack underflow!")
}

func TestStackPeek(t *testing.T) {
	s := NewStack()
	s.Push(10)

	el, err := s.Peek()
	assert.NoError(t, err, "Peek should not return an error")
	assert.Equal(t, 10, el, "Peek should return 10")
}

func TestStackPeekWithEmptyStack(t *testing.T) {
	s := NewStack()
	_, err := s.Peek()
	assert.Error(t, err, "Stack underflow!")
}

func TestStackPeekWithFullStack(t *testing.T) {
	s := NewStack()

	for i := range s.elements {
		s.Push(i)
	}

	el, err := s.Peek()
	assert.NoError(t, err, "Peek should not return an error")
	assert.Equal(t, 9, el, "Peek should return 9")
}

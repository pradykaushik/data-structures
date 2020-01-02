package stack

import "github.com/pkg/errors"

type Stack interface {
	Push(int) error
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

// ArrayStack is a stack that uses an array to store values.
type ArrayStack struct {
	values []int
	top int
	size int
	capacity int
}

func NewArrayStack(capacity int) Stack {
	return &ArrayStack{
		values: make([]int, capacity),
		top: -1,
		size: 0,
		capacity: capacity,
	}
}

func (s ArrayStack) isFull() bool {
	return s.size == s.capacity
}

func (s ArrayStack) IsEmpty() bool {
	return s.top == -1
}

func (s *ArrayStack) Push(val int) error {
	if s.isFull() {
		return errors.New("stack is full")
	}

	s.top++
	s.values[s.top] = val
	s.size++
	return nil
}

func (s *ArrayStack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}

	topValue := s.values[s.top]
	s.top--
	s.size--
	return topValue, nil
}

func (s ArrayStack) Peek() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}

	return s.values[s.top], nil
}

func (s ArrayStack) Size() int {
	return s.size
}

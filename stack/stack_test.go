package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	s := NewArrayStack(10)
	assert.True(t, s.IsEmpty())
	assert.Equal(t, s.Size(), 0)
}

func TestArrayStack_IsEmpty(t *testing.T) {
	s := NewArrayStack(10)
	assert.True(t, s.IsEmpty())
	for i := 10; i < 20; i++ {
		s.Push(i)
	}

	for i := 10; i < 20; i++ {
		s.Pop()
	}

	assert.True(t, s.IsEmpty())
}

func TestArrayStack_Peek(t *testing.T) {
	s := NewArrayStack(10)
	for i := 10; i < 20; i++ {
		err := s.Push(i)
		assert.NoError(t, err)
	}
	val, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, val, 19)
}

func TestArrayStack_Push(t *testing.T) {
	s := NewArrayStack(10)
	for i := 10; i < 20; i++ {
		err := s.Push(i)
		assert.NoError(t, err)
		val, err := s.Peek()
		assert.NoError(t, err)
		assert.Equal(t, val, i)
	}
}

func TestArrayStack_Pop(t *testing.T) {
	s := NewArrayStack(10)
	for i := 10; i < 20; i++ {
		err := s.Push(i)
		assert.NoError(t, err)
	}

	for i := 0; i < 10; i++ {
		val, err := s.Pop()
		assert.NoError(t, err)
		assert.Equal(t, val, 20-i-1)
	}
}

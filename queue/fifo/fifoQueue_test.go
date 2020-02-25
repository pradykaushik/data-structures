package fifo

import (
	"github.com/pradykaushik/data-structures/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

type fifoQVal int

func (v fifoQVal) Get() interface{} {
	return int(v)
}

func TestNewFifoQueue(t *testing.T) {
	q := NewFifoQueue(10)
	// Testing types.
	_, ok := q.(queue.Queue)
	assert.True(t, ok)
	_, ok = q.(*FifoQueue)
	assert.True(t, ok)
	// Testing capacity.
	assert.Equal(t, q.Capacity(), 10)
	// Testing size.
	assert.Zero(t, q.Size())
}

func getQueue(capacity int, t *testing.T) queue.Queue {
	q := NewFifoQueue(capacity)
	for i := 0; i < capacity; i++ {
		err := q.Enqueue(fifoQVal(i))
		assert.NoError(t, err)
	}
	return q
}

func TestEnqueue(t *testing.T) {
	q := getQueue(10, t)
	assert.Error(t, q.Enqueue(fifoQVal(100))) // value does not matter.
	assert.Equal(t, q.Size(), 10)
	assert.False(t, q.IsEmpty())
}

func TestDequeue(t *testing.T) {
	q := getQueue(10, t)
	assert.Equal(t, q.Size(), 10)
	assert.False(t, q.IsEmpty())
	for i := 0; i < 10; i++ {
		val, err := q.Dequeue()
		assert.NoError(t, err)
		assert.NotNil(t, val)
		_, ok := val.Get().(int)
		assert.True(t, ok)
		assert.Equal(t, i, val.Get().(int))
		assert.Equal(t, 10-i-1, q.Size())
	}
	assert.True(t, q.IsEmpty())
	val, err := q.Dequeue()
	assert.Nil(t, val)
	assert.Error(t, err)
}

func TestClear(t *testing.T) {
	q := getQueue(10, t)
	assert.Equal(t, q.Size(), 10)
	assert.False(t, q.IsEmpty())
	q.Clear()
	assert.Zero(t, q.Size())
	assert.True(t, q.IsEmpty())
}

package fifo

import (
	"github.com/pradykaushik/data-structures/queue"
	"github.com/pkg/errors"
)

// LinearQueueArr is a first-in-first-out queue implementation.
// Implements queue interface.
// The queue is implemented using an array.
type LinearQueueArr struct {
	values []queue.Value
	capacity int
	front int
	rear int
	// size is the current number of elements in the queue.
	size int
}

// NewLinearQueueArr returns a queue of the given capacity.
func NewLinearQueueArr(c int) queue.Queue {
	return &LinearQueueArr{
		values: make([]queue.Value, c),
		capacity: c,
		front: -1,
		rear: -1,
		size: 0,
	}
}

// Enqueue the given value at the rear end of the queue.
// Returns error if the queue is full.
func (q *LinearQueueArr) Enqueue(v queue.Value) error {
	if q.isFull() {
		return errors.New("cannot enqueue value as queue is full")
	}
	q.reinitIfRequired()
	q.rear++
	q.values[q.rear] = v
	// advancing front if first value.
	if q.front == -1 {
		q.front = 0
	}
	q.size++
	return nil
}

// Dequeue and return the value at the front of the queue.
// Returns error if the queue is empty.
func (q *LinearQueueArr) Dequeue() (queue.Value, error) {
	if q.IsEmpty() {
		return nil, errors.New("cannot dequeue as the queue is empty")
	}
	val := q.values[q.front]
	q.front++
	q.size--
	// Resetting front and rear if queue is empty.
	if q.IsEmpty() {
		q.front = -1
		q.rear = -1
	}
	return val, nil
}

// Peek at the value present at the front of the queue.
// Returns error if the queue is empty.
func (q LinearQueueArr) Peek() (queue.Value, error) {
	if q.IsEmpty() {
		return nil, errors.New("cannot peek as queue is empty")
	}
	return q.values[q.front], nil
}

// Capacity returns the capacity of the queue.
func (q LinearQueueArr) Capacity() int {
	return q.capacity
}

// Size returns the number of values in the queue.
func (q LinearQueueArr) Size() int {
	return q.size
}

// Clear the contents of the queue.
func (q *LinearQueueArr) Clear() {
	q.values = nil // allow garbage collection and avoid memory leaks.
	q.size = 0
}

// IsEmpty returns whether the queue is empty.
func (q LinearQueueArr) IsEmpty() bool {
	return q.size == 0
}

// ifFull returns whether the queue has reached its capacity.
func (q LinearQueueArr) isFull() bool {
	return q.size == q.capacity
}

// reinitIfRequired re-initializes the queue if the internally stored slice is nil.
func (q *LinearQueueArr) reinitIfRequired() {
	if q.values != nil {
		return
	}
	q.values = make([]queue.Value, q.capacity)
	// For safety.
	q.front = -1
	q.rear = -1
	q.size = 0
}

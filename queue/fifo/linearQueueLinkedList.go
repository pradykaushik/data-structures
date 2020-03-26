package fifo

import (
	"github.com/pradykaushik/data-structures/queue"
	ll "github.com/pradykaushik/data-structures/linkedlist"
	"github.com/pradykaushik/data-structures/util"
	"github.com/pkg/errors"
)

// LinearQueueLL is a first-in-first-out queue implementation.
// Implements queue interface.
// The queue is implemented using a singly linkedlist.
type LinearQueueLL struct {
	values *ll.LinkedList
	capacity int
}

// NewLinearQueueLL returns a queue of the given capacity.
func NewLinearQueueLL(c int) queue.Queue {
	return &LinearQueueLL{
		values: ll.NewLinkedList(),
		capacity: c,
	}
}

func (q LinearQueueLL) IsEmpty() bool {
	return q.values.IsEmpty()
}

// Enqueue the given value at the rear end of the queue.
// Returns error if the queue is full.
func (q *LinearQueueLL) Enqueue(v util.Value) error {
	if q.isFull() {
		return errors.New("cannot enqueue value as queue is full")
	}
	q.values.Append(v)
	return nil
}

// Dequeue and return the value at the front of the queue.
// Return error if the queue is empty.
func (q *LinearQueueLL) Dequeue() (util.Value, error) {
	if q.IsEmpty() {
		return nil, errors.New("cannot dequeue as the queue is empty")
	}

	val, _ := q.values.DeleteAtPos(0)
	return val, nil
}

// Peek at the value present at the front of the queue.
// Returns error is the queue is empty.
func (q LinearQueueLL) Peek() (util.Value, error) {
	if q.IsEmpty() {
		return nil, errors.New("cannot peek as the queue is empty")
	}
	return q.values.HeadVal(), nil
}

// Capacity returns the capacity of the queue.
func (q LinearQueueLL) Capacity() int {
	return q.capacity
}

// Size returns the number of values in the queue.
func (q LinearQueueLL) Size() int {
	return q.values.Size()
}

// Clear the contents of the queue.
func (q *LinearQueueLL) Clear() {
	q.values = ll.NewLinkedList()
}

// isFull returns whether the queue has reached its capacity.
func (q LinearQueueLL) isFull() bool {
	return q.values.Size() == q.capacity
}

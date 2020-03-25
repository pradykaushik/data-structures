package queue

import "github.com/pradykaushik/data-structures/util"

// Queue defines a set of APIs common to all queues.
type Queue interface {
	// Enqueue adds the given value to the queue.
	Enqueue(util.Value) error
	// Dequeue removes and returns the value at the front of the queue.
	Dequeue() (util.Value, error)
	// Peek returns the value at the front of the queue but does not remove it.
	Peek() (util.Value, error)
	// IsEmpty returns whether the queue is empty.
	IsEmpty() bool
	// Capacity returns the capacity of the queue.
	Capacity() int
	// Size returns the number of values in the queue.
	Size() int
	// Clear the contents of the queue. After this the queue would be empty.
	Clear()
}

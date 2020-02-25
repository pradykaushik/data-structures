package queue

// Queue defines a set of APIs common to all queues.
type Queue interface {
	// Enqueue adds the given value to the queue.
	Enqueue(Value) error
	// Dequeue removes and returns the value at the front of the queue.
	Dequeue() (Value, error)
	// Peek returns the value at the front of the queue but does not remove it.
	Peek() (Value, error)
	// IsEmpty returns whether the queue is empty.
	IsEmpty() bool
	// Capacity returns the capacity of the queue.
	Capacity() int
	// Size returns the number of values in the queue.
	Size() int
	// Clear the contents of the queue. After this the queue would be empty.
	Clear()
}

// Value stored in the queue.
type Value interface {
	// Get the contained value.
	Get() interface{}
}

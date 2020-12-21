package heap

import "github.com/pkg/errors"

type Heap interface {
	BuildHeap()
	Heapify(int)
	Insert(int)
	DeleteMax()
	FindMax() (int, error)
	IsEmpty() bool
	Size() int
}

type MaxHeap struct {
	data []int
	size int
}

func NewMaxHeap(arr []int) Heap {
	return &MaxHeap{
		data: arr,
		size: len(arr),
	}
}

func (h MaxHeap) withinBounds(i int) bool {
	return (i < len(h.data)) && (i >= 0)
}

func (h MaxHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *MaxHeap) Heapify(i int) {
	// Finding the largest among root, left and right.
	left := 2*i + 1
	right := 2*i + 2
	var largest int
	if (h.withinBounds(left)) && h.data[left] > h.data[i] {
		largest = left
	} else {
		largest = i
	}
	if (h.withinBounds(right)) && h.data[right] > h.data[largest] {
		largest = right
	}
	// If largest is not i, then data is not a heap.
	if largest != i {
		// Swapping values at largest and i.
		h.data[largest], h.data[i] = h.data[i], h.data[largest]
		// Continuing to heapify subtree with root as largest (value was updated with data[i]).
		h.Heapify(largest)
	}
}

func (h *MaxHeap) BuildHeap() {
	// For any binary tree, the nodes at indices N/2, N/2 + 1, N/2 + 2 ... are leaf nodes.
	// Leaf nodes are max heaps by default.
	// Internal nodes are at indices N/2 - 1, N/2 - 2 ... 0. Therefore, to build a heap, we build it bottom up.
	for i := (len(h.data) / 2) - 1; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *MaxHeap) siftUp(i int) {
	parent := (i - 1) / 2
	if h.withinBounds(parent) && (h.data[parent] < h.data[i]) {
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		h.siftUp(parent)
	}
}

func (h *MaxHeap) Insert(val int) {
	// Insert val at the end and then sift up till the heap property holds.
	h.data = append(h.data, val)
	h.size++
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap) DeleteMax() {
	// Swap the value at root (max) with that at the last node.
	// Shrink the array by 1 to exclude the last value.
	// Sift down the value at root until the heap property is restored.
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.Heapify(0) // This will sift down the root to the correct location in the heap.
	h.size--
}

func (h MaxHeap) FindMax() (int, error) {
	if h.IsEmpty() {
		return -1, errors.New("heap is empty ")
	}

	return h.data[0], nil
}

func (h MaxHeap) Size() int {
	return h.size
}

package heap

type Heap interface {
	BuildHeap()
	Heapify(int)
	Insert(int)
	Delete(int) bool
	SiftUp()
	SiftDown()
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
	return i < len(h.data)
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

func (h *MaxHeap) Insert(val int) {

}

func (h *MaxHeap) Delete(val int) bool {
	return false
}

func (h *MaxHeap) SiftUp() {

}

func (h *MaxHeap) SiftDown() {

}

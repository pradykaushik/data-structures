package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getMaxHeap(data []int) Heap {
	return NewMaxHeap(data)
}

func TestNewMaxHeap(t *testing.T) {
	h := getMaxHeap([]int{5, 7, 10, 1, 4, 11, 13})
	assert.Equal(t, h.(*MaxHeap).data, []int{5, 7, 10, 1, 4, 11, 13})
}

func testHeap(t *testing.T, data []int) {
	for i := 0; i < len(data)/2; i++ {
		// value at index i should be greater than values at 2i+1 and 2i+2.
		left := 2*i + 1
		right := 2*i + 2
		if left < len(data) {
			assert.Greater(t, data[i], data[left])
		}
		if right < len(data) {
			assert.Greater(t, data[i], data[right])
		}
	}
}

func TestMaxHeap_BuildHeap(t *testing.T) {
	h := getMaxHeap([]int{5, 7, 10, 1, 4, 11, 13})
	h.BuildHeap()
	data := h.(*MaxHeap).data
	testHeap(t, data)
}

func TestMaxHeap_Insert(t *testing.T) {
	h := getMaxHeap([]int{5, 7, 10, 1, 4, 11, 13})
	h.BuildHeap()
	oldSize := h.Size()
	// Inserting 17 and checking whether heap property is maintained.
	// At the end of the insertion, 17 should be the root.
	h.Insert(17)
	newSize := h.Size()
	assert.Equal(t, oldSize+1, newSize)
	testHeap(t, h.(*MaxHeap).data)
}

func TestMaxHeap_FindMax(t *testing.T) {
	h := getMaxHeap([]int{5, 7, 10, 1, 4, 11, 13})
	h.BuildHeap()
	max, err := h.FindMax()
	assert.NoError(t, err)
	assert.Equal(t, max, 13)
}

func TestMaxHeap_DeleteMax(t *testing.T) {
	h := getMaxHeap([]int{5, 7, 10, 1, 4, 11, 13})
	h.BuildHeap()
	oldSize := h.Size()
	h.DeleteMax()
	newSize := h.Size()
	testHeap(t, h.(*MaxHeap).data)
	assert.Equal(t, oldSize-1, newSize)
	max, _ := h.FindMax()
	assert.Equal(t, max, 11)
}

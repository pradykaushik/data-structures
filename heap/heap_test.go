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

func TestMaxHeap_BuildHeap(t *testing.T) {
	h := getMaxHeap([]int{5, 7, 10, 1, 4, 11, 13})
	h.BuildHeap()
	data := h.(*MaxHeap).data
	for i := 0; i < len(data)/2; i++ {
		// value at index i should be greater than values at 2i+1 and 2i+2.
		left := 2*i + 1
		right := 2*i + 2
		assert.Greater(t, data[i], data[left])
		assert.Greater(t, data[i], data[right])
	}
}

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func getLinkedList() *LinkedList {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		ll.Append(i)
	}
	return ll
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	assert.NotNil(t, ll)
	assert.True(t, ll.IsEmpty())
}

func TestLinkedList_Append(t *testing.T) {
	ll := getLinkedList()
	values := ll.SerializeIntoArray()
	assert.Equal(t, len(values), 10)
	for i := 0; i < 10; i++ {
		assert.Equal(t, values[i], i)
	}
}

func TestLinkedList_Search(t *testing.T) {
	ll := getLinkedList()
	for i := 0; i < 10; i++ {
		assert.True(t, ll.Search(i))
	}
	assert.False(t, ll.Search(11))
}

func TestLinkedList_Delete(t *testing.T) {
	ll := getLinkedList()
	// Testing deletion of the head.
	for i := 0; !ll.IsEmpty(); i++ {
		assert.True(t, ll.Delete(i))
	}
	// LinkedList should be empty.
	assert.True(t, ll.IsEmpty())

	// Testing deletion of tail.
	ll = getLinkedList()
	deleted := make(map[int]bool)
	for !ll.IsEmpty() {
		randVal := rand.Intn(10)
		if _, ok := deleted[randVal]; ok {
			continue
		}
		assert.True(t, ll.Delete(randVal))
		deleted[randVal] = true
	}
	assert.True(t, ll.IsEmpty())
}

func TestLinkedList_Reverse(t *testing.T) {
	ll := getLinkedList()
	values := ll.SerializeIntoArray()
	ll.Reverse()
	valuesReversed := ll.SerializeIntoArray()
	assert.Equal(t, len(values), len(valuesReversed))
	for i := 0; i < len(values); i++ {
		assert.Equal(t, values[i], valuesReversed[len(values)-i-1])
	}
}

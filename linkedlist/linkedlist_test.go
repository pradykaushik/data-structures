package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)


// LLValue represents the value stored in a linkedlist as an integer.
// Implements util.Value interface.
type LLValue int

func (v LLValue) Get() interface{} {
	return int(v)
}

func getLinkedList() *LinkedList {
	ll := New()
	for i := 0; i < 10; i++ {
		ll.Append(LLValue(i))
	}
	return ll
}

func TestNewNode(t *testing.T) {
	n := NewNode(LLValue(10))
	assert.Equal(t, n.val.Get().(int), 10)
	assert.Nil(t, n.next)
}

func TestNew(t *testing.T) {
	ll := New()
	assert.NotNil(t, ll)
	assert.True(t, ll.IsEmpty())
}

func TestLinkedList_Append(t *testing.T) {
	ll := getLinkedList()
	values := ll.SerializeIntoArray()
	assert.Equal(t, len(values), 10)
	for i := 0; i < 10; i++ {
		assert.Equal(t, values[i].Get().(int), i)
	}
}

func TestLinkedList_Search(t *testing.T) {
	ll := getLinkedList()
	for i := 0; i < 10; i++ {
		assert.True(t, ll.Search(LLValue(i)))
	}
	assert.False(t, ll.Search(LLValue(11)))
}

func TestLinkedList_Delete(t *testing.T) {
	ll := getLinkedList()
	// Testing deletion of the head.
	for i := 0; !ll.IsEmpty(); i++ {
		assert.True(t, ll.Delete(LLValue(i)))
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
		assert.True(t, ll.Delete(LLValue(randVal)))
		deleted[randVal] = true
	}
	assert.True(t, ll.IsEmpty())
}

func TestLinkedList_DeleteAtPos(t *testing.T) {
	ll := getLinkedList()
	// Testing deletion of the head.
	for i := 0; !ll.IsEmpty(); i++ {
		deletedVal, hasBeenDeleted := ll.DeleteAtPos(0)
		assert.True(t, hasBeenDeleted)
		assert.NotNil(t, deletedVal)
	}
	// LinkedList should be empty.
	assert.True(t, ll.IsEmpty())

	// Testing deletion of node at random position.
	ll = getLinkedList()
	for !ll.IsEmpty() {
		values := ll.SerializeIntoArray()
		randPos := rand.Intn(len(values))
		deletedVal, hasBeenDeleted := ll.DeleteAtPos(randPos)
		assert.True(t, hasBeenDeleted)
		assert.NotNil(t, deletedVal)
		// LinkedList should be smaller in size by 1 element.
		assert.Equal(t, ll.Size(), len(values)-1)
		valuesAfterRemoval := ll.SerializeIntoArray()

		// Checking that the correct value has been removed.
		if randPos == (len(values)-1) {
			if len(valuesAfterRemoval) > 0 {
				n := len(valuesAfterRemoval)
				assert.NotEqual(t, values[randPos], valuesAfterRemoval[n-1])
			}
		} else {
			assert.NotEqual(t, values[randPos], valuesAfterRemoval[randPos])
		}
	}
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

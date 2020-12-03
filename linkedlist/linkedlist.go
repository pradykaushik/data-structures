package linkedlist

import "github.com/pradykaushik/data-structures/util"

type LinkedList struct {
	head *node
	size int
}

func New() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

func (ll LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll LinkedList) Size() int {
	return ll.size
}

func (ll LinkedList) HeadVal() util.Value {
	if ll.IsEmpty() {
		return nil
	}
	return ll.head.val
}

func (ll *LinkedList) AddToFront(val util.Value) {
	n := NewNode(val)
	n.next = ll.head
	ll.head = n
	ll.size++
}

func (ll *LinkedList) Append(val util.Value) {
	if ll.IsEmpty() {
		ll.head = NewNode(val)
		ll.size++
		return
	}

	var cur *node
	for cur = ll.head; cur.next != nil; cur = cur.next {
		// do nothing.
		// cur automatically advances to the tail.
	}
	cur.next = NewNode(val)
	ll.size++
}

func (ll LinkedList) Search(val util.Value) bool {
	if ll.IsEmpty() {
		return false
	}

	for cur := ll.head; cur != nil; cur = cur.next {
		if cur.val == val {
			return true
		}
	}

	return false
}

func (ll *LinkedList) Delete(val util.Value) bool {
	if ll.IsEmpty() {
		return false
	}

	deleted := false
	var prev *node
	for cur := ll.head; cur != nil; prev, cur = cur, cur.next {
		if cur.val == val {
			deleted = true
			// prev is nil for first node.
			if prev == nil {
				ll.head = cur.next
				// cutting the link.
				cur.next = nil
			} else {
				// prev should point to the node pointed to by cur.
				prev.next = cur.next
				// cutting the link.
				cur.next = nil
			}
			ll.size--
		}
	}
	return deleted
}

// DeleteAtPos deletes value at the given position.
// Return boolean indicating whether the deletion was successful.
// The value deleted is also returned.
func (ll *LinkedList) DeleteAtPos(pos int) (util.Value, bool) {
	if ll.IsEmpty() {
		return nil, false
	}
	if pos >= ll.size {
		return nil, false
	}

	var prev *node
	var cur = ll.head
	var deletedVal util.Value

	for i := 0; i < pos; i++ {
		prev = cur
		cur = cur.next
	}

	// If prev is nil, then we need to remove head.
	// Else, we need to remove the cur node.
	deletedVal = cur.val
	if prev == nil {
		ll.head = cur.next
		cur.next = nil
	} else {
		prev.next = cur.next
		cur.next = nil
	}
	cur = nil // setting up for garbage collection.
	ll.size--
	return deletedVal, true
}

func (ll *LinkedList) Reverse() {
	if ll.IsEmpty() {
		return
	}

	var prev *node
	var cur = ll.head
	var next = cur.next
	for cur.next != nil {
		cur.next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.next
		}
	}
	if prev != nil {
		cur.next = prev
	}
	ll.head = cur
}

func (ll LinkedList) SerializeIntoArray() []util.Value {
	values := make([]util.Value, ll.size)
	for i, cur := 0, ll.head; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.val
	}
	return values
}

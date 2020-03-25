package linkedlist

type LinkedList struct {
	head *node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

func (ll LinkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Append(val int) {
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

func (ll LinkedList) Search(val int) bool {
	if ll.IsEmpty() {
		return false
	}

	for cur := ll.head; cur != nil; cur = cur.next {
		if cur.val.Get().(int) == val {
			return true
		}
	}

	return false
}

func (ll *LinkedList) Delete(val int) bool {
	if ll.IsEmpty() {
		return false
	}

	deleted := false
	var prev *node
	for cur := ll.head; cur != nil; prev, cur = cur, cur.next {
		if cur.val.Get().(int) == val {
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

func (ll LinkedList) SerializeIntoArray() []int {
	values := make([]int, ll.size)
	for i, cur := 0, ll.head; cur != nil; i, cur = i+1, cur.next {
		values[i] = cur.val.Get().(int)
	}
	return values
}

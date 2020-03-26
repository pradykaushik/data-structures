package binarySearchTree

type node struct {
	val   int
	left  *node
	right *node
}

func newNode(val int) *node {
	return &node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (n *node) update(val int) {
	n.val = val
}

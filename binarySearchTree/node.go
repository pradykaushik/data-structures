package binarySearchTree

import "github.com/pradykaushik/data-structures/util"

type node struct {
	val   util.Value
	left  *node
	right *node
}

func newNode(val int) *node {
	return &node{
		val:   BSTValue(val),
		left:  nil,
		right: nil,
	}
}

func (n *node) update(val int) {
	n.val = BSTValue(val)
}

// BSTValue represents the value stored in a binary search tree.
// Implements util.Value interface.
type BSTValue int

func (v BSTValue) Get() interface{} {
	return int(v)
}

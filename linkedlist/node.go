package linkedlist

import "github.com/pradykaushik/data-structures/util"

type node struct {
	val  util.Value
	next *node
}

func NewNode(val int) *node {
	return &node{
		val:  LLValue(val),
		next: nil,
	}
}

// LLValue represents the value stored in a linkedlist.
// Implements util.Value interface.
type LLValue int

func (v LLValue) Get() interface{} {
	return int(v)
}

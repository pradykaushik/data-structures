package linkedlist

import "github.com/pradykaushik/data-structures/util"

type node struct {
	val  util.Value
	next *node
}

func NewNode(val util.Value) *node {
	return &node{
		val:  val,
		next: nil,
	}
}

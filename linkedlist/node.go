package linkedlist

type node struct {
	val  int
	next *node
}

func NewNode(val int) *node {
	return &node{
		val:  val,
		next: nil,
	}
}

package binarySearchTree

type BinarySearchTree struct {
	root *node
	size int
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
		size: 0,
	}
}

func (t *BinarySearchTree) Insert(val int) {
	if t.IsEmpty() {
		t.root = newNode(val)
	} else {
		t.insert(t.root, val)
	}
	t.size++
}

func (t *BinarySearchTree) insert(cur *node, val int) {
	if cur.val.Get().(int) < val {
		if cur.right == nil {
			cur.right = newNode(val)
			return
		}
		t.insert(cur.right, val)
	} else if cur.val.Get().(int) > val {
		if cur.left == nil {
			cur.left = newNode(val)
			return
		}
		t.insert(cur.left, val)
	}
}

func (t BinarySearchTree) Search(val int) bool {
	if t.IsEmpty() {
		return false
	}
	return t.search(t.root, val)
}

func (t BinarySearchTree) search(cur *node, val int) bool {
	if cur == nil {
		return false
	}
	if cur.val.Get().(int) == val {
		return true
	}
	if cur.val.Get().(int) < val {
		return t.search(cur.right, val)
	}
	return t.search(cur.left, val)
}

func (t *BinarySearchTree) Delete(val int) (deleted bool) {
	if t.IsEmpty() {
		deleted = false
	} else {
		deleted = t.delete(val)
		if deleted {
			t.size--
		}
	}
	return
}

func (t BinarySearchTree) find(cur *node, val int) (*node, *node) {
	return t.findHelper(nil, t.root, val)
}

func (t BinarySearchTree) findHelper(parent *node, cur *node, val int) (*node, *node) {
	if cur == nil {
		return parent, cur
	}
	if cur.val.Get().(int) == val {
		return parent, cur
	}
	if cur.val.Get().(int) < val {
		return t.findHelper(cur, cur.right, val)
	}
	return t.findHelper(cur, cur.left, val)
}

func (t BinarySearchTree) isLeaf(n *node) bool {
	return (n != nil) && (n.left == nil) && (n.right == nil)
}

func (t *BinarySearchTree) delete(val int) bool {
	parent, valNode := t.find(t.root, val)
	if valNode == nil {
		return false // node with val does not exist.
	}
	// If valNode is the leaf node, then the link to the parent needs to be snapped.
	if t.isLeaf(valNode) {
		if parent == nil {
			// valNode is root.
			t.root = nil
			return true
		}
		if parent.left == valNode {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return true
	}
	// valNode is not a leaf.
	// If valNode is root and there is no right child.
	// In this case, the left child becomes the root.
	if (valNode == t.root) && (t.root.right == nil) {
		t.root = t.root.left
		return true
	}
	// valNode is root or internal node.
	if valNode.right == nil {
		parent.left = valNode.left
		return true
	}
	// A valid right child exists and we need to find a replacement.
	// The next largest value in the right subtree.
	parent = valNode
	temp := valNode.right
	for temp.left != nil {
		parent = temp
		temp = temp.left
	}
	// Replacing the value at valNode.
	valNode.val = temp.val
	// Breaking the link between parent and temp as that node is no longer required.
	if parent.left == temp {
		parent.left = temp.right
	} else {
		parent.right = temp.right
	}
	return true
}

func (t BinarySearchTree) IsEmpty() bool {
	return t.root == nil
}

func (t BinarySearchTree) Inorder() []int {
	data := make([]int, 0)
	t.populateInorder(t.root, &data)
	return data
}

func (t BinarySearchTree) populateInorder(cur *node, data *[]int) {
	if cur == nil {
		return
	}
	t.populateInorder(cur.left, data)
	*data = append(*data, cur.val.Get().(int))
	t.populateInorder(cur.right, data)
}

func (t BinarySearchTree) Preorder() []int {
	data := make([]int, 0)
	t.populatePreorder(t.root, &data)
	return data
}

func (t BinarySearchTree) populatePreorder(cur *node, data *[]int) {
	if cur == nil {
		return
	}
	*data = append(*data, cur.val.Get().(int))
	t.populatePreorder(cur.left, data)
	t.populatePreorder(cur.right, data)
}

func (t BinarySearchTree) Postorder() []int {
	data := make([]int, 0)
	t.populatePostorder(t.root, &data)
	return data
}

func (t BinarySearchTree) populatePostorder(cur *node, data *[]int) {
	if cur == nil {
		return
	}
	t.populatePostorder(cur.left, data)
	t.populatePostorder(cur.right, data)
	*data = append(*data, cur.val.Get().(int))
}

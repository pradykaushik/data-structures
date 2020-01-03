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
	if cur.val < val {
		if cur.right == nil {
			cur.right = newNode(val)
			return
		}
		t.insert(cur.right, val)
	} else if cur.val > val {
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
	if cur.val == val {
		return true
	}
	if cur.val < val {
		return t.search(cur.right, val)
	}
	return t.search(cur.left, val)
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
	*data = append(*data, cur.val)
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
	*data = append(*data, cur.val)
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
	*data = append(*data, cur.val)
}

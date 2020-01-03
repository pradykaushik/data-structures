package binarySearchTree

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTree()
	assert.Nil(t, bst.root)
	assert.Zero(t, bst.size)
}

func getBST() (*BinarySearchTree, []int) {
	bst := NewBinarySearchTree()
	values := []int{13, 11, 10, 17, 7, 4, 1}
	for _, val := range values {
		bst.Insert(val)
	}
	return bst, values
}

func testBST(cur *node) bool {
	// Empty tree is a valid BST.
	if cur == nil {
		return true
	}
	// Leaf nodes are valid BST.
	if (cur.left == nil) && (cur.right == nil) {
		return true
	}

	// Checking whether natural ordering is maintained between cur, left and right.
	return (checkLeft(cur) && checkRight(cur)) &&
		testBST(cur.left) && testBST(cur.right)
}

func checkLeft(cur *node) bool {
	// Checking whether natural ordering is maintained between cur and left.
	if (cur.left != nil) && (cur.left.val < cur.val) {
		return true
	} else if cur.left == nil {
		return true
	}
	return false
}

func checkRight(cur *node) bool {
	// Checking whether natural ordering is maintained between cur and left.
	if (cur.right != nil) && (cur.right.val > cur.val) {
		return true
	} else if cur.right == nil {
		return true
	}
	return false
}

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := NewBinarySearchTree()
	values := []int{13, 11, 10, 17, 7, 4, 1}
	for _, val := range values {
		bst.Insert(val)
	}
	assert.False(t, bst.IsEmpty())
	assert.Equal(t, bst.size, len(values))

	testBST(bst.root)
}

func TestBinarySearchTree_Search(t *testing.T) {
	bst, values := getBST()
	for _, val := range values {
		assert.True(t, bst.Search(val))
	}
	assert.False(t, bst.Search(14))
}

func TestBinarySearchTree_Inorder(t *testing.T) {
	bst, values := getBST()
	sort.SliceStable(values, func(i, j int) bool {
		return values[i] <= values[j]
	})
	assert.Equal(t, bst.Inorder(), values)
}

func TestBinarySearchTree_Preorder(t *testing.T) {
	bst, _ := getBST()
	expectedValuesPreorder := []int{13, 11, 10, 7, 4, 1, 17}
	assert.Equal(t, expectedValuesPreorder, bst.Preorder())
}

func TestBinarySearchTree_Postorder(t *testing.T) {
	bst, _ := getBST()
	expectedValuesPostorder := []int{1, 4, 7, 10, 11, 17, 13}
	assert.Equal(t, expectedValuesPostorder, bst.Postorder())
}

package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := NewNode(10)
	assert.Equal(t, n.val, 10)
	assert.Nil(t, n.next)
}

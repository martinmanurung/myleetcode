package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	assert := assert.New(t)
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}

	assert.Equal(inorderTraversal(root), []int{3, 2, 1, 4})
}

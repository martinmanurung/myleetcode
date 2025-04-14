package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}

// closure recursive
func inorderTraversal2(root *TreeNode) []int {
	var result []int
	var traverse func(node *TreeNode)

	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)               // Langkah 1
		result = append(result, node.Val) // Langkah 2
		traverse(node.Right)              // Langkah 3
	}

	traverse(root)
	return result
}

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(root))
}

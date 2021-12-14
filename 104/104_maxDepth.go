package main

/**
求二叉树最大深度
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

func maxDepth(root *TreeNode) int {
	return depth(root, 0)
}

func depth(root *TreeNode, n int) int {
	if root != nil {
		return max(depth(root.Left, n+1), depth(root.Right, n+1))
	} else {
		return n
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

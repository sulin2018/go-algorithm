package main

import "math"

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

func isValidBST(root *TreeNode) bool {
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	// 当前值必须满足 min < val < max
	if node.Val <= min || node.Val >= max {
		return false
	}
	return isBST(node.Left, min, node.Val) && isBST(node.Right, node.Val, max)
}

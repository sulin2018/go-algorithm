package main

/**
二叉树中序遍历
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

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	helper(root, &res)
	return res
}

func helper(root *TreeNode, res *[]int) {
	if root.Left != nil {
		helper(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		helper(root.Right, res)
	}
}

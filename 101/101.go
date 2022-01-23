package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

给你一个二叉树的根节点 root ， 检查它是否轴对称。

示例 1：
		1
	2		 2
3		4  4	3


输入：root = [1,2,2,3,4,4,3]
输出：true
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricHelper(p *TreeNode, q *TreeNode) bool {
	// 当前对比树都nil
	if p == nil && q == nil {
		return true
	}
	// 当前对比树只有一个为nil
	if p == nil || q == nil {
		return false
	}
	// 值相等 且 p左与q右相等 且p右与q左相等
	return p.Val == q.Val && isSymmetricHelper(p.Left, q.Right) && isSymmetricHelper(p.Right, q.Left)
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root, root)
}

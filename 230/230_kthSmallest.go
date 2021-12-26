package main

/**
寻找二叉搜索树的第k小元素, k计数从1开始
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

func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历 可以提前结束循环 借助栈避免遍历后续内容
	var tempStack []*TreeNode
	for {
		for root != nil {
			tempStack = append(tempStack, root)
			root = root.Left
		}
		root = tempStack[len(tempStack)-1]
		tempStack = tempStack[:len(tempStack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

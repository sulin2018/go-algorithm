package main

import "math"

/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。

思路:
一个节点的最大值为本身加上左右可贡献的最大加数, 遍历树, 更新这个值即可.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := math.MinInt32

	// 返回一个节点最大可贡献的加数
	var getNodeMax func(*TreeNode) int
	getNodeMax = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := maxVal(getNodeMax(node.Left), 0)
		rightMax := maxVal(getNodeMax(node.Right), 0)

		// 更新最大值
		max = maxVal(max, leftMax+node.Val+rightMax)

		// 最大可贡献加数只应该包含树的一边
		return node.Val + maxVal(leftMax, rightMax)
	}
	getNodeMax(root)
	return max

}

func maxVal(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

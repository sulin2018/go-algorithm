package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

 [1,2,3,4,null,null,5]
	    1
	  2    3
	4 nil nil 5

给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。


示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]

思路:
层序遍历, 间隔行翻转结果即可
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	reverseLine := false
	stack := []*TreeNode{root}
	i := 0
	j := 1
	lineRes := []int{}
	for i < j {
		// 遍历stack
		lineRes = append(lineRes, stack[i].Val)
		if stack[i].Left != nil {
			stack = append(stack, stack[i].Left)
		}
		if stack[i].Right != nil {
			stack = append(stack, stack[i].Right)
		}

		// 到达行末尾
		if i == j-1 {
			// 追加本行
			if reverseLine {
				lineRes = reverse(lineRes)
			}
			res = append(res, lineRes)
			// 置空行
			lineRes = []int{}

			// 如果还有下一行 更新末尾 更新方向
			if j != len(stack) {
				j = len(stack)
				reverseLine = !reverseLine
			}
		}

		i++
	}

	return res
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

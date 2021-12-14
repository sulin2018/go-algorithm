package main

/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

思路: 广度优先改进: 一次性遍历完一层的数据
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var stack []*TreeNode

	if root != nil {
		stack = []*TreeNode{root}
	}
	// 开始循环
	for len(stack) > 0 {
		n := len(stack)
		// 当前层数据
		level := []int{}
		// 遍历当前层
		for i := 0; i < n; i++ {
			level = append(level, stack[i].Val)
			// 如果有子节点就加入下次遍历
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		// 消除当前层内容
		stack = stack[n:]
		// 追加结果
		result = append(result, level)
	}
	return result
}

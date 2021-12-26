package main

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。(节点本身也可以说是它祖先)

思路:
因为是二叉搜索树, 值大于两个目标, 说明在左子树; 小于两个目标, 说明在右子树; 否则就找到了
*/

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	ancestor := root
// 	for {
// 		if p.Val < ancestor.Val && q.Val < ancestor.Val {
// 			ancestor = ancestor.Left
// 		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
// 			ancestor = ancestor.Right
// 		} else {
// 			return ancestor
// 		}
// 	}
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 向下寻找到叶子节点或者p/q
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 找到 那么当前节点就是最近公共父节点
	if left != nil && right != nil {
		return root
	}
	// 左边没有找到 一定是右边
	if left == nil {
		return right
	}
	// 右边没有找到 一定是左边
	return left
}

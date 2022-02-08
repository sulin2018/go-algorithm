package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }

 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



示例 1：

输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。

*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	helper(root, nil)
	return root
}

func helper(node *Node, p *Node) {
	if node == nil {
		// 叶子
		return
	}
	if p == nil {
		// root
		node.Next = nil
		helper(node.Left, node)
		helper(node.Right, node)
	} else if p.Left == node {
		// 节点为左孩子
		node.Next = p.Right
		helper(node.Left, node)
		helper(node.Right, node)
	} else {
		// 节点为右孩子
		if p.Next != nil {
			node.Next = p.Next.Left
			helper(node.Left, node)
			helper(node.Right, node)
		} else {
			node.Next = nil
			helper(node.Left, node)
			helper(node.Right, node)
		}
	}
}

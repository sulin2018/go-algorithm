package main

/**
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


思路:
  前序: [ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
  中序: [ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	var l, r []int
	l = inorder[0:i]
	r = inorder[i+1:]

	tempNode := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:len(l)+1], l),
		Right: buildTree(preorder[len(l)+1:], r),
	}
	return tempNode
}

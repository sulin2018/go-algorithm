package main

/**
删除链表上的节点, 但是只给定了被删除节点, 且该节点不是末尾节点
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// 由于无法访问到上一个节点, 直接改变当前节点值为下一个节点, 移除下一个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

package main

/**
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

func reverseList(head *ListNode) *ListNode {
	// if head == nil {
	//     return nil
	// }
	// p, next := head, head.Next
	// p.Next = nil
	// for next != nil {
	// 	temp := next
	// 	next = next.Next
	// 	temp.Next = p
	// 	p = temp
	// }
	// return p

	// 结果
	var prev *ListNode
	for head != nil {
		// head下一个节点该为前一个节点 head后移 前一个节点变化为当前head
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}

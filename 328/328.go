package main

/**
给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别组合在一起，然后返回重新排序的列表。

第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。

请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。

你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。



示例 1:



输入: head = [1,2,3,4,5]
输出: [1,3,5,2,4]


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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var h1, h1Cur, h2, h2Cur, next *ListNode
	i := 1
	for head != nil {
		// 注意断开下节点连接
		next = head.Next
		head.Next = nil
		if i%2 == 0 {
			if h2Cur == nil {
				h2Cur = head
				h2 = h2Cur
			} else {
				h2Cur.Next = head
				h2Cur = h2Cur.Next
			}
		} else {
			if h1Cur == nil {
				h1Cur = head
				h1 = h1Cur
			} else {
				h1Cur.Next = head
				h1Cur = h1Cur.Next
			}
		}
		head = next
		i++
	}
	h1Cur.Next = h2
	return h1
}

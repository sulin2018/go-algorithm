package main

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

思路:
	注意k可能比链表大的问题.
	双指针法: 间隔n个的i,j指针, j到末尾终止循环; 此时i的next即是新的链表头

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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	p := head
	listLen := 0
	for p != nil {
		listLen++
		p = p.Next
	}

	tempK := k % listLen
	i, j := head, head
	for j.Next != nil {
		if tempK > 0 {
			j = j.Next
			tempK--
		} else {
			j = j.Next
			i = i.Next
		}
	}
	j.Next = head
	p = i.Next
	i.Next = nil
	return p
}

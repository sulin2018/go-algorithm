package main

/**
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。

思路: 两个链表分别保持 大值/小值; 最后小值末尾接上大值即可
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

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	big := &ListNode{}
	pSmall := small
	pBig := big
	for head != nil {
		if head.Val >= x {
			pBig.Next = head
			pBig = pBig.Next
		} else {
			pSmall.Next = head
			pSmall = pSmall.Next
		}
		head = head.Next
	}
	pBig.Next = nil
	pSmall.Next = big.Next
	return small.Next
}

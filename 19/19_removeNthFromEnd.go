package main

/**
移除链表倒数第n个节点
思路: 双指针法, i/j间隔n个, j到末尾时, 移除i的下一个即可
	注意创建一个零值头, 应对输入为 [1] 1 这种情况
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{0, head}
	i, j := pre, head
	for n > 0 {
		n--
		j = j.Next
	}
	for j != nil {
		i = i.Next
		j = j.Next
	}
	i.Next = i.Next.Next
	return pre.Next
}

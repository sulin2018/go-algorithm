package main

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }

合并两个有序链表
思路: 递归即可, 例如但l1较小时 l1 = merge(l1.Next, l2)
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}

// 非递归实现
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	tempHead := &ListNode{}
	p := tempHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return tempHead.Next
}

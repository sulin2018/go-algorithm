package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 排序链表
// 思路: 归并排序, 不断合并两个有序链表, 复杂度 nlog(n)

func merge(head1, head2 *ListNode) *ListNode {
	// 合并两个有序链表
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func sortHelper(start, end *ListNode) *ListNode {

	if start == nil {
		return start
	}
	if start.Next == end {
		start.Next = nil
		return start
	}
	// 寻找到中间节点
	slow, fast := start, start
	for fast != end {
		slow = slow.Next
		fast = fast.Next
		if fast != end {
			fast = fast.Next
		}
	}
	mid := slow
	// 合并操作
	return merge(sortHelper(start, mid), sortHelper(mid, end))
}

func sortList(head *ListNode) *ListNode {
	return sortHelper(head, nil)
}

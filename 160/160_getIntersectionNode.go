package main

/**
编写一个程序，找到两个单链表相交的起始节点。

思路:
两个指针分别指向开头, 遍历到末尾就指向另外一个指针开头; 如果相交就返回, 否则不相交;
记录一个是否改变指向的标志, 可提前结束; 否则不相交情况会导致一定遍历到结尾;

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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	aChange, bChange := false, false
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		if pa.Next == nil {
			if aChange {
				break
			} else {
				pa = headB
				aChange = true
			}
		} else {
			pa = pa.Next
		}
		if pb.Next == nil {
			if bChange {
				break
			} else {
				pb = headA
				bChange = true
			}
		} else {
			pb = pb.Next
		}
	}
	return nil

	// ha, hb := headA, headB
	// for ha != hb {
	// 	if ha != nil {
	// 		ha = ha.Next
	// 	} else {
	// 		ha = headB
	// 	}
	// 	if hb != nil {
	// 		hb = hb.Next
	// 	} else {
	// 		hb = headA
	// 	}
	// }
	// return ha
}

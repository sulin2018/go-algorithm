package main

// 通过hash实现: 时间空间复杂度都是log(n)
// func detectCycle(head *ListNode) *ListNode {
//     seen := map[*ListNode]struct{}{}
//     for head != nil {
//         if _, ok := seen[head]; ok {
//             return head
//         }
//         seen[head] = struct{}{}
//         head = head.Next
//     }
//     return nil
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

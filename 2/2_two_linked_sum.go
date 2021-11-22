package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 笨办法: 直接转换为数字后相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := 0
	num2 := 0

	p := l1
	var i = 1
	for {
		num1 += p.Val * i
		i *= 10
		if p.Next == nil {
			break
		}
		p = p.Next
	}

	p = l2
	i = 1
	for {
		num2 += p.Val * i
		i *= 10
		if p.Next == nil {
			break
		}
		p = p.Next
	}

	sum := num1 + num2
	fmt.Println(num1, num2, sum)
	var result *ListNode
	p = result
	temp := sum
	for {
		node := &ListNode{
			Val:  temp % 10,
			Next: nil,
		}
		if result == nil {
			result = node
			p = node
		} else {
			p.Next = node
			p = node
		}
		if temp < 10 {
			break
		}
		temp = temp / 10
	}
	return result
}

// 进位计算方式
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var i *ListNode
	var j *ListNode
	i = l1
	j = l2
	addNum := 0
	var result *ListNode
	var p *ListNode
	for {
		if i == nil && j == nil {
			break
		}

		var iVal int
		var jVal int
		if i == nil {
			iVal = 0
		} else {
			iVal = i.Val
			i = i.Next
		}
		if j == nil {
			jVal = 0
		} else {
			jVal = j.Val
			j = j.Next
		}

		sum := iVal + jVal + addNum
		if sum >= 10 {
			addNum = 1
		} else {
			addNum = 0
		}

		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		if p == nil {
			result = node
			p = node
		} else {
			p.Next = node
			p = node
		}
	}
	if addNum == 1 {
		p.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}

	return result
}

func main() {
	l1 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	}

	result := addTwoNumbers2(l1, l2)
	p := result
	fmt.Println(result)
	for {
		fmt.Print(p.Val)
		if p.Next != nil {
			p = p.Next
		} else {
			break
		}
	}
}

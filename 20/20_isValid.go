package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

思路:
	栈结构, 左括号就入栈, 否则出栈判断
*/

func isValid(s string) bool {
	var stack []byte
	top := -1
	validMap := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack = append(stack, byte(v))
			top++
		} else {
			if top >= 0 && byte(v) == validMap[stack[top]] {
				stack = stack[:top]
				top--
			} else {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("{}"))
}

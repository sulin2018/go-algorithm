package main

import "fmt"

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。


示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。

*/

func plusOne(digits []int) []int {
	n := len(digits)

	add := false
	for i := n - 1; i >= 0; i-- {
		num := digits[i] + 1
		if num >= 10 {
			digits[i] = num % 10
			add = true
		} else {
			digits[i] = num
			add = false
			break
		}

	}

	if add {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}

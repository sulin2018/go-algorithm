package main

/*
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。

整数除法仅保留整数部分。


示例 1：

输入：s = "3+2*2"
输出：7

*/

func calculate(s string) int {
	nums := []int{}
	preSign := '+'
	num := 0
	for i, v := range s {
		isNum := v >= '0' && v <= '9'
		// 数字累计
		if isNum {
			num = num*10 + int(v-'0')
		}
		// 符号 或者是 最后一个 需要计算
		if !isNum && v != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				nums = append(nums, num)
			case '-':
				nums = append(nums, -num)
			case '*':
				nums[len(nums)-1] = nums[len(nums)-1] * num
			default:
				nums[len(nums)-1] = nums[len(nums)-1] / num
			}
			num = 0
			preSign = v
		}
	}

	// 累加结果
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

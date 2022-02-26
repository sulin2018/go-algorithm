package main

/*
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x



示例 1：

输入：n = 27
输出：true


*/

func isPowerOfThree(n int) bool {
	if n == 3 || n == 1 {
		return true
	}
	num := 3
	for num < n {
		num *= 3
		if num == n {
			return true
		}
	}
	return false
}

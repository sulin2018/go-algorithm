package main

import "math"

/*
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。



示例 1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4

动态规划
	f[i] 表示最少需要多少个数的平方来表示整数 i。
	这些数必然落在区间 [1, sqrt(n)]。我们可以枚举这些数，假设当前枚举到 j，那么我们还需要取若干数的平方，构成 i-j^2。
	f(0) = 0
	f(i) = 1 + min(f(i-j^2)), j=[1, sqrt(i)]
*/

func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, f[i-j*j])
		}
		f[i] = minn + 1
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
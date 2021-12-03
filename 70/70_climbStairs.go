package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

解法: f(n) = f(n-1) + f(n-2)
*/

// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	if n == 2 {
// 		return 2
// 	}
// 	return climbStairs(n-1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	a, b := 1, 1              // 0阶 1阶 走法种数
	for i := 2; i <= n; i++ { // 从第二阶开始计算直到第n阶
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(climbStairs(44))
}

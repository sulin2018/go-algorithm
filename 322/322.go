package main

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

你可以认为每种硬币的数量是无限的。



示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1


*/
func coinChange(coins []int, amount int) int {
	n := len(coins)
	mins := make([]int, amount+1)
	for i := range mins {
		mins[i] = amount + 1
	}

	mins[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < n; j++ {
			if coins[j] <= i {
				mins[i] = min(mins[i], mins[i-coins[j]]+1)
			}
		}
	}
	if mins[amount] > amount {
		return -1
	}
	return mins[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

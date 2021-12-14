package main

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。


说明: 递推即可
*/

func maxProduct(nums []int) int {
	maxV := nums[0]
	minV := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		oldMax := maxV
		oldMin := minV
		// 当前最大值可能是 当前值 当前值*累计最大值 当前值*累计最小值
		maxV = max(oldMax*nums[i], max(nums[i], oldMin*nums[i]))
		// 当前最小值可能是 当前值 当前值*累计最大值 当前值*累计最小值
		minV = min(oldMin*nums[i], min(nums[i], oldMax*nums[i]))
		res = max(maxV, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

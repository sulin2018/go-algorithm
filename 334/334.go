package main

import "math"

/*
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。

如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,4,5]
输出：true
解释：任何 i < j < k 的三元组都满足题意


*/

// func increasingTriplet(nums []int) bool {
// 	incNum := make([]int, len(nums))
// 	for i := range incNum {
// 		incNum[i] = 1
// 	}

// 	for i := 1; i < len(nums); i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] {
// 				incNum[i] = max(incNum[i], incNum[j]+1)
// 				if incNum[i] >= 3 {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

func increasingTriplet(nums []int) bool {
	n := len(nums)
	f := []int{math.MaxInt32, math.MaxInt32, math.MaxInt32}
	for i := 0; i < n; i++ {
		t := nums[i]
		if f[2] < t {
			return true
		} else if f[1] < t && t < f[2] {
			f[2] = t
		} else if f[1] > t {
			f[1] = t
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}

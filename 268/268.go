package main

/*
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。



示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
*/

func missingNumber(nums []int) int {
	numMap := map[int]bool{}
	for _, n := range nums {
		numMap[n] = true
	}
	for i := 0; i < len(nums); i++ {
		if !numMap[i] {
			return i
		}
	}
	return len(nums)
}

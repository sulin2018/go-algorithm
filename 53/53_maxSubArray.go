package main

import (
	"fmt"
	"math"
)

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

思路: 累加, 累加为负值就归0, 累加值大于最大值就赋值给最大值
*/

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	res := math.MinInt32
	temp := 0
	for i := 0; i < len(nums); i++ {
		// 累加后大于最大值就赋值给最大值
		temp = temp + nums[i]
		if temp > res {
			res = temp
		}
		// 累加后小于0 就归0 因为最大值不应该包括前面累加的这个负值
		if temp < 0 {
			temp = 0
		}
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{1, 2, -2, 4, 3, 1, -5, 6}))
}

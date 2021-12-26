package main

import "fmt"

// 求数组中除了第i位元素的其他元素的所有乘积输出
// 方法1. 用两个数组, 遍历两边, 分别记录第i位左侧/右侧的元素乘积; 最后再计算即可
// 方法2. 在方法1的基础上减少一个数组, 并且该数组同时也记录结果; 两次遍历即可完成

func productExceptSelf(nums []int) []int {
	// 左侧所有元素乘积
	res := make([]int, len(nums))

	// 第0个元素左侧没有元素 所以乘积为1
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 右侧元素乘积
	r := 1
	// 计算结果
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r = r * nums[i]
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4, 5}))
}

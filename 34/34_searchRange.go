package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？

思路:
	一次遍历记录即可 由于是升序可提前结束循环
*/

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	for i, n := range nums {
		if n == target {
			if res[0] == -1 {
				res[0] = i
			}
			res[1] = i
		}
		if n > target {
			break
		}
	}
	return res
}

func main() {
	fmt.Println(searchRange([]int{1, 2, 3, 4, 5, 5, 6, 7}, 5))
}

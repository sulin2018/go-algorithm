package main

import "fmt"

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

思路:
	注意题目说明了是已经排序了的数组, 所以只需要双指针即可实现.
	i从0开始,j从1开始; j不断寻找与i不一致的值设置到i的下一个位置
*/

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	j := 1
	i := 0
	for ; i < len(nums) && j < len(nums); j++ {
		if nums[i] != nums[j] {
			nums[i+1] = nums[j]
			i++
		}
	}
	return i + 1
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3}))
}

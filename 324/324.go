package main

import (
	"fmt"
	"sort"
)

/*
给你一个整数数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

你可以假设所有输入数组都可以得到满足题目要求的结果。



示例 1：

输入：nums = [1,5,1,1,6,4]
输出：[1,6,1,5,1,4]
解释：[1,4,1,5,1,6] 同样是符合题目要求的结果，可以被判题程序接受。

*/

func wiggleSort(nums []int) {
	sort.Ints(nums)
	mid := (len(nums) + 1) / 2
	// reverse(nums[mid:])

	res := make([]int, len(nums))
	// 注意: 中间数字可能相同, 要尽可能放置到两边; 所以取较小数字时先取最大的, 取较大数字时, 也取最大的;
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			res[i] = nums[mid-i/2-1]
		} else {
			res[i] = nums[len(nums)-(i/2)-1]
		}
	}

	for i := 0; i < len(res); i++ {
		nums[i] = res[i]
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	nums := []int{1, 3, 2, 2, 3, 1}
	// nums := []int{4, 5, 5, 6}
	wiggleSort(nums)
	fmt.Println(nums)
}

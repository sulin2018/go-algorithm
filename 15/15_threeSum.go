package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

思路:
	排序+双指针
	先排序, 然后确定一个数字, 再由两个指针在剩下的数组中两端开始向中间靠拢寻找目标
*/

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for first := 0; first < n-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		target := -1 * nums[first]
		third := n - 1
		for second := first + 1; second < n-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{1, 2, 3, -2, 0}))
}

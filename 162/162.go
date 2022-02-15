package main

/*
峰值元素是指其值严格大于左右相邻值的元素。

给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞ 。

你必须实现时间复杂度为 O(log n) 的算法来解决此问题。



示例 1：

输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
*/

func findPeakElement(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	i := 0
	for i < len(nums) {
		// 第一个只需要比第二个大就算
		if i == 0 {
			if nums[i] > nums[i+1] {
				return i
			}
			i++
			continue
		}
		// 最后一个只需要比倒数第二个大就算
		if i == len(nums)-1 {
			if nums[len(nums)-1] > nums[len(nums)-2] {
				return i
			}
			i++
			continue
		}
		// 其他情况需要比两边大
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			return i
		}
		i++
	}
	return 0
}

// 直接寻找最大值
// func findPeakElement(nums []int) (idx int) {
//     for i, v := range nums {
//         if v > nums[idx] {
//             idx = i
//         }
//     }
//     return
// }

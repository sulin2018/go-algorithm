package main

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

进阶：
尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题

思路:
1. 通过hash可以实现
2. 累计法: 遇到相同的就累加计数, 否则减去计数; 由于题目保证了存在这样一个数, 所以count到末尾时, 肯定不是0, 不会导致越界
*/

func majorityElement(nums []int) int {
	count := 0
	majority := nums[0]
	for i, v := range nums {
		if majority == v {
			count++
		} else {
			count--
			if count == 0 {
				majority = nums[i+1]
			}
		}
	}
	return majority
}

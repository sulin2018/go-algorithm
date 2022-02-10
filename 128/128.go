package main

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

*/
func longestConsecutive(nums []int) int {
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}

	res := 0
	for _, num := range nums {
		// 对于元素x, 如果x-1存在, 那么已经判断过, 忽略; 即只判断x-1不存在的
		if !numMap[num-1] {
			numLen := 1
			// 对于元素x, 一直寻找x+1
			next := num + 1
			for numMap[next] {
				next++
				numLen++
			}

			// 更新结果
			if numLen > res {
				res = numLen
			}
		}
	}

	return res
}

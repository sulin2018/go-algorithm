package main

import "fmt"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

思路:
	回溯法: 递归 满足条件退出 回溯
*/

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	used := make([]bool, len(nums))
	var tempNums []int
	backtrack(nums, tempNums, used)
	return res
}

func backtrack(nums []int, tempNums []int, used []bool) {
	// 满足条件 加入结果
	if len(tempNums) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, tempNums)
		res = append(res, temp)
		return
	}

	// 遍历, 进行选择数字
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			// 当前没有被选择过就选中 然后进入下一步
			used[i] = true
			tempNums = append(tempNums, nums[i])
			backtrack(nums, tempNums, used)
			// 回溯, 去除当前选择 从而去选下一个未选择的
			tempNums = tempNums[:len(tempNums)-1]
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

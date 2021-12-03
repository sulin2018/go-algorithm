package main

import "fmt"

// 求数组中组成目标值target的所有不重复集合, 数组中的值可以重复使用

// 思路: 回溯法, 寻找目标值 和 序号index; 两种选择: 1. 不选取index; 2. 选取index, 并继续回溯

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	temp := []int{}
	var helper func(int, int)
	helper = func(tempTarget int, index int) {
		if index == len(candidates) {
			return
		}
		if tempTarget == 0 {
			res = append(res, append([]int(nil), temp...))
			return
		}
		helper(tempTarget, index+1)
		if tempTarget-candidates[index] >= 0 {
			temp = append(temp, candidates[index])
			helper(tempTarget-candidates[index], index)
			temp = temp[:len(temp)-1]
		}
	}
	helper(target, 0)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{1, 2, 3, 4, 5, 6}, 5))
}

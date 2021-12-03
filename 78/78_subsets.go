package main

import "fmt"

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

思路:
	枚举标志位, 例如
	001 表示取第一个
	010 表示取第二个
	直到111即可完成枚举
*/

func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	var res [][]int
	n := uint32(len(nums))
	// 一共有 2^n 个组合
	for i := 0; i < (1 << n); i++ {
		temp := []int{}
		for j, v := range nums {
			// i的二进制的第j位为1 说明该数字应该加入当前枚举值
			if (i>>uint32(j))&1 > 0 {
				temp = append(temp, v)
			}
		}
		res = append(res, temp)
	}
	return res
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

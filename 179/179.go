package main

import (
	"sort"
	"strconv"
)

/*
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。


示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"

思路:
以第一个数字大小 从大到小排列; 如9 81组合, 981>819
遇到相同大的, 需要把大的放更前面; 如41 4 组合, 441>414
*/

func largestNumber(nums []int) string {
	// 排序算法: 两两组合比较即可
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func main() {
	println(largestNumber([]int{3, 30, 34, 5, 9}))
}

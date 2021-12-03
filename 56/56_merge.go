package main

import (
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

思路: 先进行排序, 然后一步步合并即可
*/

func merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	var res [][]int
	// 初始值
	res = append(res, intervals[0])
	// 从第2个区间开始
	for i := 1; i < len(intervals); i++ {
		// 如果被与结果末尾区间重叠 更新为重叠后的值
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1] = []int{res[len(res)-1][0], max(res[len(res)-1][1], intervals[i][1])}
		} else {
			// 否则直接追加
			res = append(res, intervals[i])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(merge([][]int{[]int{1, 3}, []int{2, 6}}))
}

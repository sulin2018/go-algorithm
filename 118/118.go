package main

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

		1
	1		1
1		2		1


示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

*/

func generate(numRows int) [][]int {
	res := [][]int{{1}}
	i := 0
	for numRows > 1 {
		// 第i行
		row := res[i]

		// 计算第i+1行
		newRow := []int{1}
		// 第0个为1 从第1个元素开始计算
		j := 1
		for j < len(row) {
			newRow = append(newRow, row[j-1]+row[j])
			j++
		}
		// 最后元素为0
		newRow = append(newRow, 1)

		// 计算结束
		res = append(res, newRow)
		i++
		numRows--
	}
	return res
}

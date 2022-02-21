package main

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。

输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5
输出：true

思路:
	从右上角开始搜索, 设元素为v; v大于target, 排除列; v小于target, 排除行
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, n-1
	for i < m && j > 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			j--
		} else {
			i--
		}
	}
	return false
}

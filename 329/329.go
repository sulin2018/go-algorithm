package main

/*
给定一个 m x n 整数矩阵 matrix ，找出其中 最长递增路径 的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你 不能 在 对角线 方向上移动或移动到 边界外（即不允许环绕）。



示例 1：


输入：matrix = [
	[9,9,4],
	[6,6,8],
	[2,1,1]
]
输出：4
解释：最长递增路径为 [1, 2, 6, 9]。
*/

var (
	dirs          = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	rows, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, memo))
		}
	}
	return ans
}

func dfs(matrix [][]int, row, column int, memo [][]int) int {
	if memo[row][column] != 0 {
		return memo[row][column]
	}
	memo[row][column]++
	for _, dir := range dirs {
		newRow, newColumn := row+dir[0], column+dir[1]
		if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[row][column] {
			memo[row][column] = max(memo[row][column], dfs(matrix, newRow, newColumn, memo)+1)
		}
	}
	return memo[row][column]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

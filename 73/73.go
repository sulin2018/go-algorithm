package main

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。

原地: 常数内存空间
*/

func setZeroes(matrix [][]int) {
	row := len(matrix[0])
	col := len(matrix)
	colFlag, rowFlag := false, false

	// 记录0列是否含0
	for i := 0; i < col; i++ {
		if matrix[i][0] == 0 {
			colFlag = true
			break
		}
	}
	// 记录0行是否含0
	for i := 0; i < row; i++ {
		if matrix[0][i] == 0 {
			rowFlag = true
			break
		}
	}
	// 用第0行及0列来标记要置0的行/列
	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := 1; i < col; i++ {
		for j := 1; j < row; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if colFlag {
		for i := 0; i < col; i++ {
			matrix[i][0] = 0
		}
	}
	if rowFlag {
		for j := 0; j < row; j++ {
			matrix[0][j] = 0
		}
	}
}

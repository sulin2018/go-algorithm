package main

/*
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

思路:
	1 2 3     x x 1
	x x x --> x x 2
	x x x     x x 3

	核心实现: matrix_new[j][n - i - 1] = matrix[i][j]
		可以借助另外矩阵实现, 但是不满足原地旋转
	翻转实现:
		1. 上下翻转    matrix_new[i][j] = matrix[n-1-i][j]
		2. 主对角线翻转 matrix_new[i][j] = matrix[j][i] 即 matrix[n-1-i][j] = matrix[j][n-1-i]
*/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

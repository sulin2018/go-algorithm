package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

例如:
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。

思路: 遍历整个表格计算每一格
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	i := 0
	for ; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				// 第0行 只能从左边累加来
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				// 第0列 只能从上面累加来
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				// 其他情况 从左边/上面较小累加来
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
			// fmt.Println(grid[i][j])
		}
	}
	return grid[i-1][len(grid[i-1])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPathSum([][]int{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	}))
}

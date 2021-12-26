package main

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

思路: 遍历, 遇到1, 就递归将上下左右全部置0; 最后计算1个数即可.
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0

	var make1to0 func(i, j int)
	make1to0 = func(i, j int) {
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		if i > 0 {
			make1to0(i-1, j)
		}
		if j > 0 {
			make1to0(i, j-1)
		}
		if i < len(grid)-1 {
			make1to0(i+1, j)
		}
		if j < len(grid[0])-1 {
			make1to0(i, j+1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 计数同时把周围所有1全部置为0
				res++
				make1to0(i, j)
			}
		}
	}
	return res
}

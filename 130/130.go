package main

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。


示例 1：

输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。


思路:
先把边缘O标记为M(任意非O非X), 再把所有O标记为X, 再把M标记回O
*/

func solve(board [][]byte) {
	if len(board) <= 1 {
		return
	}

	rowLen, colLen := len(board), len(board[0])
	for i := 0; i < rowLen; i++ {
		if board[i][0] == 'O' {
			mark(board, false, i, 0)
		}
		if board[i][colLen-1] == 'O' {
			mark(board, false, i, colLen-1)
		}
	}
	for j := 0; j < colLen; j++ {
		if board[0][j] == 'O' {
			mark(board, false, 0, j)
		}
		if board[rowLen-1][j] == 'O' {
			mark(board, false, rowLen-1, j)
		}
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < rowLen; i++ {
		if board[i][0] == 'M' {
			mark(board, true, i, 0)
		}
		if board[i][colLen-1] == 'M' {
			mark(board, true, i, colLen-1)
		}
	}
	for j := 0; j < colLen; j++ {
		if board[0][j] == 'M' {
			mark(board, true, 0, j)
		}
		if board[rowLen-1][j] == 'M' {
			mark(board, true, rowLen-1, j)
		}
	}
}

func mark(board [][]byte, markToO bool, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == 'X' {
		return
	}
	if markToO && board[i][j] == 'M' {
		board[i][j] = 'O'
	} else if !markToO && board[i][j] == 'O' {
		board[i][j] = 'M'
	} else {
		return
	}
	mark(board, markToO, i-1, j)
	mark(board, markToO, i, j+1)
	mark(board, markToO, i+1, j)
	mark(board, markToO, i, j-1)
}

package main

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

思路: 遍历+回溯
*/
type Point struct{ x, y int }

func exist(board [][]byte, word string) bool {
	if len(board) < 1 {
		return false
	}
	// 上下左右
	nextPoints := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 记录已经用过的字母
	checked := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		checked[i] = make([]bool, len(board[i]))
	}

	// 检查回溯函数
	var check func(x, y, k int) bool
	check = func(x, y, k int) bool {
		// 当前满足
		if board[x][y] != word[k] {
			return false
		}
		// 已经到最后 能找到 返回true退出
		if k == len(word)-1 {
			return true
		}
		// 标记当前位置为已使用
		checked[x][y] = true
		defer func() { checked[x][y] = false }()
		// 遍历周围
		for _, nextPoint := range nextPoints {
			newX, newY := x+nextPoint.x, y+nextPoint.y
			// 未出界 且未使用过
			if 0 <= newX && newX < len(board) && 0 <= newY && newY < len(board[0]) && checked[newX][newY] == false {
				// 递归判断
				if check(newX, newY, k+1) {
					return true
				}
			}
		}
		return false
	}

	// 所有节点都可能是开始
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

package main

func isValidSudoku(board [][]byte) bool {
	var row [9][10]bool
	var col [9][10]bool
	var box [9][10]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0'
			if row[i][num] {
				return false
			}
			if col[j][num] {
				return false
			}
			if box[j/3+(i/3)*3][num] {
				return false
			}

			row[i][num] = true
			col[j][num] = true
			box[j/3+(i/3)*3][num] = true
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	println(isValidSudoku(board))
}

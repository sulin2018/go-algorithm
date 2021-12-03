package main

import "fmt"

/*
螺旋遍历输出二维数组
*/

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	j := len(matrix)    // 竖向个数
	i := len(matrix[0]) // 横向个数
	k := i * j
	var res []int

	order := 0 // 0横向左 1竖向下 2横向右 3竖向上
	// 当前遍历元素: m竖向 n横向
	m, n := 0, 0
	iNums, jNums := 0, 0 // 横向竖向次数
	for s := 0; s < k; s++ {
		res = append(res, matrix[m][n])
		// 计算下一个位置
		if order == 0 {
			if n == i-1-jNums/2 {
				order = 1
				iNums++
				m++
			} else {
				n++
			}
			continue
		}
		if order == 1 {
			if m == j-1-iNums/2 {
				order = 2
				jNums++
				n--
			} else {
				m++
			}
			continue
		}
		if order == 2 {
			if n == jNums/2 {
				order = 3
				iNums++
				m--
			} else {
				n--
			}
			continue
		}
		if order == 3 {
			if m == iNums/2 {
				order = 0
				jNums++
				n++
			} else {
				m--
			}
			continue
		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{7, 8, 9},
	}))
}

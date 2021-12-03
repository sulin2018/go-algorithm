package main

import "fmt"

/*
螺旋填写二维数组
*/

func generateMatrix(n int) [][]int {
	if n < 1 {
		return nil
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	i, j, iNums, jNums := 0, 0, 0, 0
	order := 0 // 0横向右 1竖向下 2横向左 3竖向上
	nums := n * n
	for k := 1; k <= nums; k++ {
		res[j][i] = k
		if k == nums {
			break
		}

		if order == 0 {
			if i == n-1-jNums/2 {
				order = 1
				iNums++
				j++
			} else {
				i++
			}
			continue
		}
		if order == 1 {
			if j == n-1-iNums/2 {
				order = 2
				jNums++
				i--
			} else {
				j++
			}
			continue
		}
		if order == 2 {
			if i == jNums/2 {
				order = 3
				iNums++
				j--
			} else {
				i--
			}
			continue
		}
		if order == 3 {
			if j == iNums/2 {
				order = 0
				jNums++
				i++
			} else {
				j--
			}
			continue
		}
	}

	return res
}

func main() {
	fmt.Println(generateMatrix(5))
}

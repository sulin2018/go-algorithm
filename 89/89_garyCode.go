package main

import "fmt"

/*
格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。

给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。即使有多个不同答案，你也只需要返回其中一种。

格雷编码序列必须以 0 开头。

思路:
	镜像法: 第n阶是n-1阶所有编码前面加0,然后逆序加1 的结果
	00
	01
	11 上面逆序高位+1
	10
*/

func grayCode(n int) []int {
	res, head := []int{0}, 1
	for i := 0; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, head+res[j])
		}
		head <<= 1
	}
	return res
}

func main() {
	fmt.Println(grayCode(4))
}

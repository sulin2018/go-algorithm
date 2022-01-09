package main

import "fmt"

/*
实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）。


示例 1：

输入：x = 2.00000, n = 10
输出：1024.00000


思路:
x^5 = x^2 * x^2 * x
x^4 = x^2 * x^2
*/

func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow2(x, n)
	}
	return 1 / pow2(x, -n)
}

func pow2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := pow2(x, n/2)
	if n%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}

func main() {
	fmt.Println(myPow(2, 3))
}

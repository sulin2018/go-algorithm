package main

import (
	"fmt"
	"strconv"
)

/*
根据 逆波兰表示法，求表达式的值。

有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

注意 两个整数之间的除法只保留整数部分。

可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。


示例 1：

输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

*/

func evalRPN(tokens []string) int {
	res := []int{}
	for _, v := range tokens {
		if v == "*" {
			tempV := res[len(res)-2] * res[len(res)-1]
			res = res[:len(res)-2]
			res = append(res, tempV)
		} else if v == "+" {
			tempV := res[len(res)-2] + res[len(res)-1]
			res = res[:len(res)-2]
			res = append(res, tempV)
		} else if v == "-" {
			tempV := res[len(res)-2] - res[len(res)-1]
			res = res[:len(res)-2]
			res = append(res, tempV)
		} else if v == "/" {
			tempV := res[len(res)-2] / res[len(res)-1]
			res = res[:len(res)-2]
			res = append(res, tempV)
		} else {
			n, _ := strconv.Atoi(v)
			res = append(res, n)
		}
	}
	return res[0]
}

func main() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

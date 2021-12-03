package main

import (
	"fmt"
	"strconv"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

思路:
	竖式计算 累加字符串即可
*/

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		temp := ""
		// 末尾补0
		for j := n - 1; j > i; j-- {
			temp = temp + "0"
		}

		// 计算当前乘积后的值
		add := 0
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			result := x*y + add
			temp = strconv.Itoa(result%10) + temp
			add = result / 10
		}
		if add > 0 {
			temp = strconv.Itoa(add) + temp
		}

		// 累加
		res = addString(res, temp)
	}
	return res
}

func addString(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	res := ""
	for ; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		res = strconv.Itoa(result%10) + res
		add = result / 10
	}
	return res
}

func main() {
	// fmt.Println(addString("123", "345"))
	fmt.Println(multiply("123", "345"), 123*345)
}

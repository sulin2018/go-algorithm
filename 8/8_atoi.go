package main

import (
	"fmt"
	"math"
)

/*
请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，形成一个有符号整数。
假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，它们对函数不应该造成影响。
假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换，即无法进行有效转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0 。

注意：

本题中的空白字符只包括空格字符 ' ' 。
假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−2^31,  2^31 − 1]。如果数值超过这个范围，请返回  2^31 − 1 或 −2^31 。

思路:
先去除首部空格 然后遍历字符串根据情况判断即可
*/

func myAtoi(s string) int {
	newStr := leftStrap(s)
	sign := 1
	res := 0
	for i, v := range newStr {
		if i == 0 {
			if v == '-' {
				sign = -1
			} else if v == '+' {
				sign = 1
			} else if int(v-'0') <= 9 && int(v-'0') >= 0 {
				// 是数字
				res = int(v - '0')
			} else {
				return res
			}
		} else {
			tempInt := int(v - '0')
			// 不是数字
			if tempInt > 9 || tempInt < 0 {
				break
			}
			if sign == 1 && res*10+tempInt > math.MaxInt32 {
				return math.MaxInt32
			}
			if sign == -1 && res*10-tempInt < math.MinInt32 {
				return math.MinInt32
			}

			res = res*10 + tempInt*sign
		}
	}

	return res
}

func leftStrap(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	for start < len(s) {
		if s[start] == ' ' {
			start++
		} else {
			break
		}
	}
	return s[start:]
}

func main() {
	// fmt.Println(leftStrap("     -32425afas"))
	fmt.Println(myAtoi("     -32425afas"))
}

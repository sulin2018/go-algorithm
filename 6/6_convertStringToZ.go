package main

import (
	"fmt"
	"strings"
)

// 将字符串变换为n行z型结构 按行输出
/*
a   e
b d f h
c   g

输出: aebdfhcg

思路: 初始化n个(行)字符串对应每行 遍历字符串放入对应行的字符串即可
*/
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	lineStrs := []string{}
	for i := 0; i < numRows; i++ {
		lineStrs = append(lineStrs, "")
	}

	line := 1
	flag := true
	for _, c := range s {
		lineStrs[line-1] = lineStrs[line-1] + string(c)
		if line+1 > numRows {
			flag = false
		}
		if line-1 == 0 {
			flag = true
		}
		if flag {
			line++
		} else {
			line--
		}
	}
	return strings.Join(lineStrs, "")
}

func main() {
	fmt.Println(convert("abcdefgh", 2))
}

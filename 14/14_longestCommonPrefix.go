package main

import "fmt"

/*
查找字符串数组中的最长公共前缀

思路: 竖向扫描, 从0开始, 直到某个字符串末尾或不与第一个字符串第i位相等
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"123", "1234", "125"}))
}

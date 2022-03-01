package main

/*
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。


示例 1：

输入: s = "leetcode"
输出: 0
*/

func firstUniqChar(s string) int {
	sMap := map[rune]int{}
	for _, c := range s {
		sMap[c]++
	}

	for i, c := range s {
		if sMap[c] == 1 {
			return i
		}
	}
	return -1
}

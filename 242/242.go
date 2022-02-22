package main

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。



示例 1:

输入: s = "anagram", t = "nagaram"
输出: true


*/

func isAnagram(s string, t string) bool {
	sMap := map[rune]int{}
	for _, v := range s {
		sMap[v]++
	}
	for _, v := range t {
		sMap[v]--
		if sMap[v] < 0 {
			return false
		}
	}
	for _, v := range sMap {
		if v > 0 {
			return false
		}
	}
	return true
}

package main

import "sort"

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

思路:
	字母异位排序后却会是一样的, 所以已排序后的字符串作为map的key, 进行存储异位词
*/

func groupAnagrams(strs []string) [][]string {
	existMap := make(map[string][]string)
	for _, str := range strs {
		strBytes := []byte(str)
		sort.Slice(strBytes, func(i, j int) bool { return strBytes[i] < strBytes[j] })
		tempKey := string(strBytes)
		existMap[tempKey] = append(existMap[tempKey], str)
	}
	res := make([][]string, 0, len(existMap))
	for _, v := range existMap {
		res = append(res, v)
	}
	return res
}

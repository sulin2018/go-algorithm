package main

/*
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

思路:
  用dp[i]表示前i个字符是否可以被拆分, 满足 dp[i] = dp[j] && wordDict[j:i]
*/

func wordBreak(s string, wordDict []string) bool {
	existWordDict := map[string]bool{}
	for _, v := range wordDict {
		existWordDict[v] = true
	}

	// 表示 [0:i] 能否被组成
	dp := make([]bool, len(s)+1)
	// 空字符串 默认能
	dp[0] = true
	// 遍历字符
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < len(s); j++ {
			// 一旦能够组成 就退出循环开始判断下一个字符
			if dp[j] && existWordDict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

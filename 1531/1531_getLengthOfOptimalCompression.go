package main

import "fmt"

/*
行程长度编码 是一种常用的字符串压缩方法，它将连续的相同字符（重复 2 次或更多次）替换为字符和表示字符计数的数字（行程长度）。例如，用此方法压缩字符串 "aabccc" ，将 "aa" 替换为 "a2" ，"ccc" 替换为` "c3" 。因此压缩后的字符串变为 "a2bc3" 。

注意，本问题中，压缩时没有在单个字符后附加计数 '1' 。

给你一个字符串 s 和一个整数 k 。你需要从字符串 s 中删除最多 k 个字符，以使 s 的行程长度编码长度最小。

请你返回删除最多 k 个字符后，s 行程长度编码的最小长度 。

输入：s = "aaabcccd", k = 2
输出：4
解释：在不删除任何内容的情况下，压缩后的字符串是 "a3bc3d" ，长度为 6 。最优的方案是删除 'b' 和 'd'，这样一来，压缩后的字符串为 "a3c3" ，长度是 4 。

输入：s = "aaaaaaaaaaa", k = 0
输出：3
解释：由于 k 等于 0 ，不能删去任何字符。压缩后的字符串是 "a11" ，长度为 3 。
*/

const INF int = 0x3f3f3f3f

func kLen(k int) int {
	if k <= 1 {
		return 0
	} else if k > 1 && k < 10 {
		return 1
	} else if k >= 10 && k < 100 {
		return 2
	} else {
		return 3
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	// dp[i][j]表示从前i个字符中最多选择j个字符进行删除
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, k+2)
		for j := 0; j < k+2; j++ {
			dp[i][j] = INF
		}
	}
	dp[0][0] = 0

	// 如果删除字符i，则此时dp[i][j] = dp[i-1][j-1]
	// 如果保留字符i, 则此时后续尽量选择保留与字符i相同的字符
	for i := 1; i <= n; i++ {
		for j := 0; j <= k && j <= i; j++ {
			if j < k {
				dp[i][j+1] = min(dp[i][j+1], dp[i-1][j])
			}
			same := 0
			del := 0
			for m := i; m <= n; m++ {
				if s[m-1] == s[i-1] {
					same++
				} else {
					del++
				}
				if j+del <= k {
					dp[m][j+del] = min(dp[m][j+del], kLen(same)+1+dp[i-1][j])
				} else {
					break
				}
			}
		}
	}

	return dp[n][k]
}

func main() {
	fmt.Println(getLengthOfOptimalCompression("aaabcccd", 2))
	fmt.Println(getLengthOfOptimalCompression("aabbaa", 2))
	fmt.Println(getLengthOfOptimalCompression("aaaaaaaaaaa", 0))
}

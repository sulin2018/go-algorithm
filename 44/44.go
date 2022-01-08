package main

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

思路:
动态规划, dp[i][j] 表示s的前i个能否被p的前j个匹配
	字母相同, 或者p的第j个为?时, dp[i][j] = dp[i-1][j-1]
	当p的第j个为*时, dp[i][j] = dp[i-1][j] || dp[i][j-1]
	其他情况都是false
起始条件:
	dp[0][0] true
	dp[i][0]一定为false
	dp[0][j]只有j个*才true
注意i/j是指个数, 而不是下标, 下标是i-1和j-1
*/

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	// dp[i][0]一定为false
	// dp[0][j]只有j个*才true
	for j := 1; j < n+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

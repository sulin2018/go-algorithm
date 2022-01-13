package main

import (
	"fmt"
)

/*
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。


注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"


思考:
滑动窗口: i,j; 移动j到满足条件, 移动i寻找解, 记录最小间距即可
满足条件判断方法:
	need map表示还需要的字母即个数
	needDiffCount 表示还未满足的不同的字母个数
*/

func minWindow(s string, t string) string {
	need := map[rune]int{}
	needDiffCount := 0
	for _, b := range t {
		if _, ok := need[b]; ok {
			// 已经存在 说明是重复字母
			need[b]++
		} else {
			needDiffCount++
			need[b] = 1
		}
	}

	i, j := 0, -1
	minI, minJ := 0, len(s)+1
	for i < len(s) && j < len(s) {
		if needDiffCount > 0 {
			// 说明当前窗口还不满足条件
			j++
			if j == len(s) {
				break
			}
			if _, ok := need[rune(s[j])]; !ok {
				// 与t无关的字母 直接跳过
				continue
			}
			if need[rune(s[j])] == 1 {
				// 如果当前字母就差一个就清零 即全部包含了 所以needDiffCount--
				needDiffCount--
			}
			need[rune(s[j])]--
		} else {
			// fmt.Println(i, j)
			// 说明当前窗口满足包含t 判断是否最小
			if j-i < minJ-minI {
				minI, minJ = i, j
			}

			// 开始移动i
			if _, ok := need[rune(s[i])]; !ok {
				i++
				// 与t无关的字母 直接跳过
				continue
			}
			if need[rune(s[i])] == 0 {
				// i字母刚好包含 去掉就不包含了 所以needDiffCount++
				needDiffCount++
			}
			need[rune(s[i])]++
			i++
		}
	}

	if minJ == len(s)+1 {
		return ""
	}
	return s[minI : minJ+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

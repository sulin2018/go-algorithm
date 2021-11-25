package main

import "fmt"

func longestPalindrome(s string) string {
	maxStart, maxEnd := 0, 0
	strLen := len(s)
	if strLen == 1 || strLen == 0 {
		return s
	}

	// 最后一个字符无需求解
	for i := 0; i < strLen-1; i++ {
		tempStart, tempEnd := getPalindrome(s, i, i)
		if tempEnd-tempStart > maxEnd-maxStart {
			maxStart, maxEnd = tempStart, tempEnd
		}
		// 中心是两个同样字符情况
		if s[i] == s[i+1] {
			tempStart, tempEnd = getPalindrome(s, i, i+1)
			if tempEnd-tempStart > maxEnd-maxStart {
				maxStart, maxEnd = tempStart, tempEnd
			}
		}
		// fmt.Println(maxStart, maxEnd)
	}
	return s[maxStart : maxEnd+1]
}

func getPalindrome(s string, start int, end int) (int, int) {
	strLen := len(s)
	if start == 0 || end == strLen-1 {
		return start, end
	}
	for start > 0 && end < strLen-1 {
		if s[start-1] == s[end+1] {
			start--
			end++
		} else {
			break
		}
	}
	return start, end
}

func main() {
	s := "abcbabccbaxxxddde"
	fmt.Println(longestPalindrome(s))
}

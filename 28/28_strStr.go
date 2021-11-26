package main

// 寻找字符串b在a中的位置， 不存在返回-1
func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if n == 0 {
		return 0
	}
	if m < n {
		return -1
	}
	for i := 0; i < m; i++ {
		flag := true
		for j := 0; j < n; j++ {
			if i+j >= m || haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

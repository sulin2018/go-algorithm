package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func countAndSay2(n int) string {
	s := []byte(string(1))
	n--
	for n > 0 {
		n--
		s = getByteNext(s)
	}
	res := ""
	for _, sn := range s {
		res += string(sn + 48)
	}
	return res
}

func getByteNext(strBytes []byte) []byte {
	tempByte := strBytes[0]
	res := []byte{}
	num := 1
	for i := 1; i < len(strBytes); i++ {
		if tempByte == strBytes[i] {
			num++
		} else {
			res = append(res, []byte(string(num))...)
			res = append(res, tempByte)
			tempByte = strBytes[i]
			num = 1
		}
	}
	res = append(res, []byte(string(num))...)
	res = append(res, tempByte)
	return res
}

func main() {
	fmt.Println(countAndSay(5))
}

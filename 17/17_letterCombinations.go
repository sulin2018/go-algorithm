package main

import "fmt"

// 模拟九宫格输入数字时, 所有字母可能组合
// 思路: 回溯法, 枚举当前index字母即可

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res = []string{}
	backtrack(digits, 0, "")
	return res
}

func backtrack(digits string, index int, tempStr string) {
	if len(tempStr) == len(digits) {
		res = append(res, tempStr)
	} else {
		tempChars := phoneMap[string(digits[index])]
		for i := 0; i < len(tempChars); i++ {
			backtrack(digits, index+1, tempStr+string(tempChars[i]))
		}
	}
}

func main() {
	fmt.Println(letterCombinations("24"))
}

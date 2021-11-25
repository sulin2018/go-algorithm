package main

import "fmt"

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

思路: 负数, 末尾为0 都不是
	其他情况 逆转一半数字 如果与 缩减了一半 的原值相等就是, 否则不是;
	奇数位情况时, 逆转数字会多逆转一位, 要去除了判断
*/

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverseNum := 0
	for x > reverseNum {
		reverseNum = reverseNum*10 + x%10
		x /= 10
	}

	return reverseNum == x || reverseNum/10 == x
}

func main() {
	fmt.Println(isPalindrome(12321))
}

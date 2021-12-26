package main

/*
判断一个数是不是2的幂

思路: 2的幂最高位为1, 本身减1除了最高位为1; 两个数字与运算为0即是2的幂.
*/

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

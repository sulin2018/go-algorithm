package main

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

注意：
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

思路: 不断取最后一个值累加即可
*/

func reverse(x int) int {
	var res int
	for x != 0 {
		// 判断溢出
		if temp := int32(res); (temp*10)/10 != temp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}
	return res
}

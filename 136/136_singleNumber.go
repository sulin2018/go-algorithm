package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

思路: 两个相同的数字异或结果是0, 并且异或满足交换律和结合律
*/

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{2, 3, 4, 4, 2}))
}

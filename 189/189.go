package main

import "fmt"

/*
给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。


示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

思路:
k可能超过数组长度, 求得真实最小移动元素个数rk
这样根据n-rk就能确定哪些元素位于左侧, 哪些位于右侧
拼接即是结果
*/

func rotate(nums []int, k int) {
	i := 1
	n := len(nums)
	for i*n < k {
		i++
	}
	rk := k - n*(i-1)
	r := nums[:n-rk]
	l := nums[n-rk:]
	res := append(l, r...)

	i = 0
	for i < n {
		nums[i] = res[i]
		i++
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

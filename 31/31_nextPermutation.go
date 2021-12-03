package main

import "fmt"

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。

思路:
	我们要将一个 前面尽可能小的大数 与 后面尽可能小的数 交互位置, 然后升序后续数字, 过程如下
		从后向前, 寻找 a[i]<a[j], 此时 [j, end] 为降序
		在 [j, end] 中寻找比 a[i] 大的数 a[k], 交换位置
		逆序 [j, end], 即升序排列 [j, end]
	例如 12654 2与4
		14256
*/

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// 从后向前, 找到第一个升序的地方
	i, j := len(nums)-2, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	// 如果为-1, 表明数字是降序, 升序即为答案
	// 如果不为-1, 需要在后面找到一个 尽可能小的数 与 i交换
	if i != -1 {
		k := len(nums) - 1
		for nums[k] <= nums[i] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	// 逆序后面的元素 使其升序
	k := len(nums) - 1
	for j < k {
		nums[j], nums[k] = nums[k], nums[j]
		j++
		k--
	}
	fmt.Println(nums)
}

func main() {
	nextPermutation([]int{1, 1})
}

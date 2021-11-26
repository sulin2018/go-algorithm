package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

思路:
遍历数组第一个值为i, 再用另外两个指针j/k从剩下的数组中寻找另外两个值
	相等直接返回
	如果和比target大,就缩小k
	如果和比target小,就增加j
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	best := math.MaxInt32

	update := func(newNum int) {
		if abs(newNum-target) < abs(best-target) {
			best = newNum
		}
	}

	for i := range nums {
		if i > 1 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			tempSum := nums[i] + nums[j] + nums[k]
			if tempSum == target {
				return target
			}
			update(tempSum)

			if tempSum > target {
				k0 := k - 1
				for k0 > j && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				j0 := j + 1
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}
	}
	return best
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 2, 3, 4, 5, -2}, 0))
}

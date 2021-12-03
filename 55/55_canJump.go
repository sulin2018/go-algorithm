package main

import "fmt"

// 数组每个元素表示可以跳跃的最远距离, 求从0元素开始, 能否跳到最后
// 例如 [2, 0, 1] 从0开始, 最远跳2步, 可直接到最后, 满足
// 思路: 记录一个最远值, 一步一步更新即可; 直到能到达最后 或者 无法跳跃

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	// 最远能到的地方 初始为第一个值
	tempMax := nums[0]
	// 遍历更新 最远能到的地方;
	for i := 1; i < len(nums) && i <= tempMax; i++ {
		if i+nums[i] > tempMax {
			tempMax = i + nums[i]
			// 超过了就不用判断了
			if tempMax >= len(nums) {
				break
			}
		}

	}
	if tempMax >= len(nums)-1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

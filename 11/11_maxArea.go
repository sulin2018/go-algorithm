package main

import (
	"fmt"
	"math"
)

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

思路: 双指针分别表示左右高度(桶的边), 不断靠近并去掉较小值
*/

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	left, right := 0, len(height)-1
	max := math.MinInt32
	for left < right {
		if height[left] < height[right] {
			tempArea := height[left] * (right - left)
			if tempArea > max {
				max = tempArea
			}
			left++
		} else {
			tempArea := height[right] * (right - left)
			if tempArea > max {
				max = tempArea
			}
			right--
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

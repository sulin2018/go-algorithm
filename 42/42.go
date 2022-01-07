package main

import "fmt"

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

| |
|||
如图, 雨水为空白处, 即1

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水。

思路:
第i个位置雨水取决于 min(左最大高度, 右最大高度)-h[i]; 当h[i]大于 min(左最大高度, 右最大高度) 时, i位置无法装水

暴力解法: 直接遍历, 求两边max, 然后计算
动态规划:
	用两个额外数组来计算好左最大/右最大
	leftMax[i] = max(leftMax[i-1], height[i-1])
*/

func trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i-1])
	}
	for j := n - 2; j > 0; j-- {
		rightMax[j] = max(rightMax[j+1], height[j+1])
	}

	num := 0
	for i := 1; i < n-1; i++ {
		if height[i] < leftMax[i] && height[i] < rightMax[i] {
			num += min(leftMax[i], rightMax[i]) - height[i]
		}
	}
	return num
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}

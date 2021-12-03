package main

import "fmt"

// 排序颜色: 只有 红0 白1 蓝2, 进行原地排序
// 桶排序: 记录每个值个数 最后输出
func sortColors(nums []int) {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	i := 0
	for ; i < numMap[0]; i++ {
		nums[i] = 0
	}
	for ; i < numMap[0]+numMap[1]; i++ {
		nums[i] = 1
	}
	for ; i < len(nums); i++ {
		nums[i] = 2
	}
	fmt.Println(nums)
}

func main() {
	sortColors([]int{1, 1, 1, 0, 2, 0, 2, 1})
}

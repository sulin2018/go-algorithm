package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// 1. 暴力
	// for i, v1 := range nums {
	// 	for j, v2 := range nums {
	// 		if i == j {
	// 			continue
	// 		}
	// 		if v1+v2 == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }

	// 2. 利用map, hash
	// tempMap := make(map[int]int)
	// for i, v := range nums {
	// 	tempMap[v] = i
	// }
	// for i, v := range nums {
	// 	tempN := target - v
	// 	if j, ok := tempMap[tempN]; ok && j != i {
	// 		return []int{i, j}
	// 	}
	// }

	// 3. 利用map 且一次循环完成
	tempMap := make(map[int]int)
	for i, v := range nums {
		tempN := target - v
		if j, ok := tempMap[tempN]; ok && j != i {
			return []int{i, j}
		}

		tempMap[v] = i
	}

	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

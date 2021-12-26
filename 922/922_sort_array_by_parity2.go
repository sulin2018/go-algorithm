package main

import "fmt"

// A是一个奇数偶数各一半的数组, 排序, 使奇数位为奇数且有序, 偶数位为偶数且有序
// func sortArrayByParityII(A []int) []int {
// 	var tempMinIndex int
// 	for i := 0; i < len(A)-1; i++ {
// 		tempMinIndex = i
// 		// 当前位存放奇数
// 		isOdd := true
// 		if i%2 == 0 {
// 			isOdd = false
// 		}

// 		j := i + 1
// 		for ; j < len(A); j++ {
// 			jIsOdd := A[j]%2 == 1
// 			if isOdd && jIsOdd {
// 				// 需要奇数 比较位就是奇数

// 				// 当前最小值为偶数
// 				if A[tempMinIndex]%2 == 0 {
// 					tempMinIndex = j
// 				}

// 				// 比较位比最小值小
// 				if A[tempMinIndex] > A[j] {
// 					tempMinIndex = j
// 				}

// 			}

// 			if !isOdd && !jIsOdd {
// 				// 需要偶数 比较位就是偶数

// 				// 当前最小值为奇数
// 				if A[tempMinIndex]%2 == 1 {
// 					tempMinIndex = j
// 				}

// 				// 比较位比最小值小
// 				if A[tempMinIndex] > A[j] {
// 					tempMinIndex = j
// 				}
// 			}
// 		}

// 		// 替换i与最小值的位置
// 		iValue := A[i]
// 		A[i] = A[tempMinIndex]
// 		A[tempMinIndex] = iValue
// 	}

// 	return A
// }

// 原题目理解错误: 并不需要奇偶对应位有序 只需要奇数位为奇数 偶数位为偶数
func sortArrayByParityII(A []int) []int {
	tempA := make([]int, len(A))

	i := 0
	for _, v := range A {
		if v%2 == 0 {
			tempA[i] = v
			i += 2
		}
	}

	i = 1
	for _, v := range A {
		if v%2 == 1 {
			tempA[i] = v
			i += 2
		}
	}
	return tempA
}

func main() {
	fmt.Println(sortArrayByParityII([]int{12, 1, 3, 4}))
}

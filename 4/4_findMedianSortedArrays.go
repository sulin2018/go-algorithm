package main

import "fmt"

func findMedianSortedArrays(A []int, B []int) float64 {
	aLen := len(A)
	bLen := len(B)
	i, j := 0, 0
	all := []int{}
	for {
		if i == aLen || j == bLen {
			break
		}

		if A[i] < B[j] {
			all = append(all, A[i])
			i++
		} else {
			all = append(all, B[j])
			j++
		}
	}

	if i < aLen {
		all = append(all, A[i:]...)
	}

	if j < bLen {
		all = append(all, B[j:]...)
	}
	fmt.Println(all)

	min := (aLen + bLen) / 2
	if (aLen+bLen)%2 == 0 {
		return float64(all[min-1]+all[min]) / 2
	} else {
		return float64((all[min]))
	}

	// allNums := []int{}
	// i := 0
	// j := 0
	// for {
	// 	if i == len(A) && j == len(B) {
	// 		break
	// 	}

	// 	if i == len(A) || (j != len(B) && A[i] > B[j]) {
	// 		allNums = append(allNums, B[j])
	// 		j++
	// 	} else {
	// 		allNums = append(allNums, A[i])
	// 		i++
	// 	}
	// }

	// middle := int(len(allNums) / 2)
	// if len(allNums)%2 == 1 {
	// 	return float64(allNums[middle])
	// } else {
	// 	return float64((allNums[middle-1] + allNums[middle])) / 2
	// }
}

func main() {
	a := []int{2, 3, 7, 9, 10, 11}
	b := []int{2, 7, 8, 50}
	// a := []int{}
	// b := []int{1}
	fmt.Println(findMedianSortedArrays(a, b))
}

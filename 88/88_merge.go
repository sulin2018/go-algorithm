package main

import "fmt"

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。

思路:
	从最后开始设置值, 直到其中一个无元素, 将未遍历完的另外一个数组直接加入到最前面
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	i, j := m-1, n-1
	for k := m + n - 1; k >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		if i < 0 || j < 0 {
			break
		}
	}
	if i < 0 {
		copy(nums1, nums2[:j+1])
	}
}

func main() {
	// a := []int{1, 2, 3, 0, 0}
	// b := []int{2, 4}
	// merge(a, 3, b, 2)
	a := []int{2, 0}
	b := []int{1}
	merge(a, 1, b, 1)
	fmt.Println(a)
}

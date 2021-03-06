package main

/*
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的索引，否则返回 -1 。

思路:
 二分查找: 有一边总是有序的, 判断查找的值是否在其中
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, mid, end := 0, -1, len(nums)-1
	for start <= end {
		mid = start + (end-start)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] <= nums[end] {
			// 右边有序 且值在右边
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			// 左边有序 且值在左边
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

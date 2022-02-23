package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	var iRes = make([]int, n)
	for i := 0; i < n; i++ {
		iRes[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				iRes[i] = max(iRes[i], iRes[j]+1)
			}
		}
	}

	res := 1
	for i := 0; i < n; i++ {
		if iRes[i] > res {
			res = iRes[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

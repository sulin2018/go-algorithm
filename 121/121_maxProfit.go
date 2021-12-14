package main

import (
	"fmt"
	"math"
)

/*
买卖股票的最佳收益, 只能买卖一次

记录当前最大值及前面的最小值, 一次遍历即可
*/

func maxProfit(prices []int) int {
	maxProfit := 0
	tempMinPrice := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] > tempMinPrice {
			tempProfit := prices[i] - tempMinPrice
			if tempProfit > maxProfit {
				maxProfit = tempProfit
			}
		} else {
			tempMinPrice = prices[i]
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 2, 4, 6, 7, 2, 4}))
}

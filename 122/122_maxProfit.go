package main

import (
	"fmt"
	"math"
)

/*
买卖股票的最佳收益, 允许买卖多次, 但是只能卖出才能再买

一次遍历累加差值即可
*/

func maxProfit(prices []int) int {
	allProfit := 0
	tempMinPrice := math.MaxInt32
	for i := 0; i < len(prices); i++ {
		if prices[i] > tempMinPrice {
			allProfit += prices[i] - tempMinPrice
		}
		tempMinPrice = prices[i]
	}
	return allProfit
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 2, 4, 6, 7, 2, 4}))
}

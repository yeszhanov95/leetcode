package main

import "math"

func maxProfit(prices []int) int {
	sell, profit := math.MaxInt64, 0
	for i := 0; i < len(prices); i ++ {
		if prices[i] < sell {
			sell = prices[i]
		} else if prices[i] - sell > profit {
			profit = prices[i] - sell
		}
	}
	return profit
}
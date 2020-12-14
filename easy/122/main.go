package main

func maxProfit(prices []int) int {
	currMin := 1 << 31
	totalProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] <= currMin {
			currMin = prices[i]
		} else if i == (len(prices)-1) || prices[i] >= prices[i+1] {
			totalProfit += prices[i] - currMin
			currMin = 1 << 31
		}
	}
	return totalProfit
}
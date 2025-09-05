package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	best_sell := prices[len(prices)-1]

	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] < best_sell {
			profit += best_sell - prices[i]
			best_sell = prices[i]
		}
		if prices[i] > best_sell {
			best_sell = prices[i]
		}
	}

	return profit
}

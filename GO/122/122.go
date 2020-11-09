package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(prices))
}

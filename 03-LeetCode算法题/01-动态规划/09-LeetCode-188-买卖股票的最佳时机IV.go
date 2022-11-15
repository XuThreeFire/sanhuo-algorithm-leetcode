package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxProfit188(k int, prices []int) int {
	// basecase
	n := len(prices)
	k = min(k, n/2)
	var (
		buy = make([]int, k+1)
		sal = make([]int, k+1)
	)
	if n <= 1 {
		return 0
	}
	for i := 1; i <= k; i++ {
		buy[i] = math.MinInt / 2
		sal[i] = math.MinInt / 2
	}

	// init buy and sal
	buy[0] = -prices[0]
	sal[0] = 0
	for i := 1; i < n; i++ {
		buy[0] = max(buy[0], sal[0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[j] = max(buy[j], sal[j]-prices[i])
			sal[j] = max(sal[j], buy[j-1]+prices[i])
		}
	}
	return sal[k-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var (
		k      = 2
		prices = []int{3, 2, 6, 5, 0, 3}
	)
	res := maxProfit188(k, prices)
	fmt.Println("max：", res)
}

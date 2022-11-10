package main

import "fmt"

// 思路
// 本质上最大收益 = 全部k>0的曲线段，只要这个直线是向上的，那么这条线的收益就全部迟到
// 股票不限次数交易，那么只有有收益，那么就要交易，
// 只有有上升趋势，那么收益就会更新
//
func maxProfit122(prices []int) int {
	// 1, 自底向上：从第一天开始解决股票问题
	// return maxProfitDp(prices)
	// 2, 自底向上：状态压缩
	return maxProfitOptimize(prices)
}
func maxProfitDp(prices []int) int {
	var (
		n = len(prices)
	)
	if n <= 1 {
		return 0
	}
	dpHaveStockProfit := make([]int, n)
	dpNoStockProfit := make([]int, n)
	dpHaveStockProfit[0] = -prices[0]
	dpNoStockProfit[0] = 0
	for i := 1; i < n; i++ {
		// 有股票的状态怎来的？：昨天有股票的转移，昨天没有股票但今天买的
		dpHaveStockProfit[i] = max(dpHaveStockProfit[i-1], dpNoStockProfit[i-1]-prices[i])
		// 无股票的状态怎来的？：昨天无股票的转移，昨天有股票，但今天卖了
		dpNoStockProfit[i] = max(dpNoStockProfit[i-1], dpHaveStockProfit[i-1]+prices[i])
		fmt.Printf("第[%d]次，dp有：%d, dp无：%d\n", i, dpHaveStockProfit[i], dpNoStockProfit[i])
	}
	return dpNoStockProfit[n-1]
}

// 状态优化
func maxProfitOptimize(prices []int) int {
	var (
		n = len(prices)
	)
	if n <= 1 {
		return 0
	}
	var (
		buy1  = -prices[0]
		sell1 = 0
	)
	fmt.Printf("第[%d]次，dp有：%d, dp无：%d\n", 0, buy1, sell1)
	for i := 1; i < n; i++ {
		// 有股票的状态怎来的？：昨天有股票的转移，昨天没有股票但今天买的
		buy1 = max(buy1, sell1-prices[i])
		// 无股票的状态怎来的？：昨天无股票的转移，昨天有股票，但今天卖了
		sell1 = max(sell1, buy1+prices[i])
		fmt.Printf("第[%d]次，dp有：%d, dp无：%d\n", i, buy1, sell1)
	}
	return sell1 // max(dp有股票, dp无股票) 最后结果肯定是没有股票的状态>有股票的状态
}

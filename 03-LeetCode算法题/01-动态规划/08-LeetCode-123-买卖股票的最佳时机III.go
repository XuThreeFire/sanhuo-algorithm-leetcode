package main

func maxProfit123(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		// 每次买在最低点
		buy1 = max(buy1, -prices[i])
		// 每次卖在当前阶段的最高点
		sell1 = max(sell1, buy1+prices[i])
		// 第二次买，需要第一次卖出后的收益，再去买
		buy2 = max(buy2, sell1-prices[i])
		// 第二次卖，需要第二次买入后才能卖出
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

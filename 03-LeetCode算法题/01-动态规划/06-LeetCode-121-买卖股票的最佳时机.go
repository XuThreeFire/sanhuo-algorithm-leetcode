package main

func maxProfit(prices []int) int {
	var n = len(prices)
	if n <= 1 {
		return 0
	}
	// 方法1，暴力法
	// 方法2，双指针法：一个记录最低点的值，一个记录最大收益值
	var minPrice = prices[0]
	var maxValue = 0
	for i := 1; i < n; i++ {
		if prices[i] > minPrice {
			// 更新maxValue收益值
			maxValue = max(maxValue, prices[i]-minPrice)
		}
		// 更新最低价格，在maxValue后面更新，可以保证，minPrice在卖出前买入
		minPrice = min(minPrice, prices[i])
	}
	return maxValue
}

// 1, 暴力法：超时
func getMaxProfit(prices []int) int {
	var n = len(prices)
	if n <= 1 {
		return 0
	}
	var res = 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j] <= prices[i] {
				continue
			}
			res = max(res, prices[j]-prices[i])
		}
	}
	return res
}

//
//func min(a, b int) int {
//	if a<b{
//		return a
//	}
//	return b
//}
//func max(a, b int) int {
//	if a>b{
//		return a
//	}
//	return b
//}

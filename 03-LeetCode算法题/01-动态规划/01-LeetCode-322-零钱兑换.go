package main

var memo = map[int]int{}

func coinChange(coins []int, amount int) int {
	memo = map[int]int{}
	return dp(coins, amount)
}

/*
关键词：最少硬币个数、凑总金额
方法：
① 自顶向下，递归方法
② 自底向上，动态规划方法
*/
func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if amount < 0 {
		return -1
	}
	v, isOk := memo[amount]
	if isOk {
		return v
	}
	var res = 1000000
	for _, coin := range coins {
		if amount < coin {
			continue
		}
		subRes := dp(coins, amount-coin)
		if subRes == -1 {
			continue
		}
		res = getMin(res, subRes+1)
	}
	if res == 1000000 {
		memo[amount] = -1
		return -1
	}
	memo[amount] = res
	return res
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

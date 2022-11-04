package main

var memo = map[int]int{}

func coinChange(coins []int, amount int) int {
	// 1, 方法1，dp递归，类似后序遍历N叉数 res = min(res, 选择第i棵树)
	// return coinChange1(coins, amount)
	// 1, 方法2，dp数组，res = min(不兑换, 兑换第1个硬币, 兑换第2个硬币, 兑换第3个硬币....)
	return coinChangeDp2(coins, amount)

}

func coinChange1(coins []int, amount int) int {
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
	// 相当于N叉树的遍历
	// 遍历：left1, left2, left3, left4, left5, left6
	// 后序遍历：res = min(res, left)
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

// 动态规划: 自底向上
// dp[i]：表示金额为i需要最少的币数量，无法兑换，则为Max
func coinChangeDp2(coins []int, amount int) int {
	var Max = amount + 1
	var dp = make([]int, amount+1)
	dp[0] = 0 //
	for i := 1; i <= amount; i++ {
		dp[i] = Max
		// 遍历全部硬币数量，选择最少的硬币数量
		// res = min(不兑换, 选择兑换1, 选择兑换2, 选择兑换3, 选择兑换5)
		for _, coin := range coins {
			if coin <= i {
				dp[i] = getMin(dp[i], dp[i-coin]+1) // 兑换次数+1
			}
		}
	}
	// basecase
	if dp[amount] == Max {
		return -1
	}
	return dp[amount]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

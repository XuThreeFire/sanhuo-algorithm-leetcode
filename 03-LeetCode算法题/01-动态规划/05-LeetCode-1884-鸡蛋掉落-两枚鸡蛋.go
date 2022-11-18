package main

func twoEggDrop(n int) int {
	// basecase
	if n == 1 {
		return 1
	} else if n <= 3 {
		return 2
	}
	// 1，从第1层开始遍历，遍历到第n层
	var dp = make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = i // 最坏的情况当然是<=n层
		// 全部操作情况：从第1层开始扔, 从第2层开始扔.....从第i层开始扔
		// 取最小操作次数：res = min(从第1层开始扔, 从第2层开始扔.....从第i层开始扔)
		// 每次扔又有2中状态：最坏的情况 = max(碎了, 没碎）
		// ① 碎了： 转化为 (k - 1)次1个鸡蛋慢慢试了 + 本次操作1 f(i) = 1 + k -1 = k
		// ② 没碎： 转化为 (i-k)次最小操作次数子问题 f(i) = f(i-k) + 1
		for k := 1; k <= i; k++ {
			// 最坏的情况res = max(碎了, 没碎)
			// worst := max(k,  dp[i-k]+1) //
			// dp[i] = min(worst, dp[i])
			dp[i] = min(dp[i], max(k, dp[i-k]+1))
		}
	}
	return dp[n]
}

//func max(a, b int) int {
//	if a>b{
//		return a
//	}
//	return b
//}
//func min(a, b int) int {
//	if a < b{
//		return a
//	}
//	return b
//}
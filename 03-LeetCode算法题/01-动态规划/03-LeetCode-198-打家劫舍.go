package main

func rob(nums []int) int {
	// 1, 方法1：动态规划，for循环法
	// return rob1(nums)
	// 2, 方法2：优化-状态压缩
	return rob2(nums)
}

func rob1(nums []int) int {
	n := len(nums)
	// 1, baseCase
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	var s = make([]int, n+1)
	s[0] = 0
	s[1] = nums[0]
	for i := 1; i < n; i++ {
		s[i+1] = max(s[i], s[i-1]+nums[i])
	}
	return s[n]
}

func rob2(nums []int) int {
	n := len(nums)
	// 1, baseCase
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// var s = make([]int, n+1)
	var s0 = 0
	var s1 = nums[0]
	var s2 = 0
	// 从第一个房子开始，遍历全部的房子，直到终点
	for i := 1; i < n; i++ {
		// res = max(不打劫的收益，打劫的收益)
		s2 = max(s1, s0+nums[i])
		s0 = s1
		s1 = s2

	}
	return s2
}

package main

// 核心思路
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var (
		m = len(grid)
		n = len(grid[0])
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

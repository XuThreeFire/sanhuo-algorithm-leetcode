package main

// 核心思路
// 用二维数组来表示这个地图网格
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var (
		m = len(grid)
		n = len(grid[0])
	)
	// 遍历全部格子，一行行扫过去，直到终点位置
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
				// 相当于：res = min(向下无穷大, 向右走） + res
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
				// 相当于：res = min(向右无穷大, 向下走） + res
			} else {
				// 核心代码：res =  min(向下走，向右走）+ res
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[m-1][n-1]
}

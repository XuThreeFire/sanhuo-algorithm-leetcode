package main

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	// basecase
	if n < 1 {
		return 0
	} else if n == 1 {
		return triangle[0][0]
	}
	// 从起点层遍历到终点层
	for row := 1; row < n; row++ {
		// 每层，从左到右遍历，计算每个点的最小路径
		for i := 0; i < row+1; i++ {
			// 左右边界情况
			if i == 0 {
				triangle[row][i] = triangle[row][i] + triangle[row-1][i] // 左边界
			} else if i == row {
				triangle[row][i] = triangle[row][i] + triangle[row-1][i-1] // 右边界
			} else {
				// 核心代码：每个点的最小路径：res = res + min(左父节点，右边父节点)
				triangle[row][i] = triangle[row][i] + min(triangle[row-1][i-1], triangle[row-1][i])
			}

		}
	}
	// 目标位置：最后一层中最小的值
	var result = math.MaxInt32
	for i := 0; i < n; i++ {
		result = min(result, triangle[n-1][i])
	}
	return result
}

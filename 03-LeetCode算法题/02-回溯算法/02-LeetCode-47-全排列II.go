package main

// 1，先穷举出来，画一个N叉树图，
// 2, 分析出不符合的分支特点
// 此题中，不符合分支的特点是，同层重复
//var res = [][]int{} // 存储结果
//var used []bool // 标记是否已经扫过 记录nums的索引
//var track = []int{} // 排列结果
//var trackLen = 0
var lenNums = 0

func main() {

	nums := []int{1, 1, 2, 3, 3}
	permuteUnique(nums)
}

func permuteUnique(nums []int) [][]int {
	lenNums = len(nums)

	if lenNums == 1 {
		return [][]int{{1}}
	}
	// clear data
	trackLen = 0
	track = make([]int, lenNums)
	used = make([]bool, lenNums)
	res = [][]int{}
	backtrack47(nums)
	return res
}

func backtrack47(nums []int) {
	if trackLen == lenNums {
		t1 := make([]int, lenNums)
		copy(t1, track)
		res = append(res, t1)
		return
	}
	// 跟全排列 I 题型相比，多了个选择内部的标记，用于剪枝
	// 同层去重
	tempUsed := map[int]struct{}{} // 记录这个数字是否使用过
	for i := 0; i < lenNums; i++ {
		if _, isOk := tempUsed[nums[i]]; isOk {
			// 本轮选择中已经选过了 用于剪枝
			continue
		}

		if used[i] {
			// track列表已经存在过 用于剪枝
			continue
		}
		track[trackLen] = nums[i]
		trackLen += 1
		used[i] = true
		tempUsed[nums[i]] = struct{}{} // 标记这个数字 nums[i] 已经选择过
		backtrack47(nums)
		trackLen -= 1
		used[i] = false
	}
}

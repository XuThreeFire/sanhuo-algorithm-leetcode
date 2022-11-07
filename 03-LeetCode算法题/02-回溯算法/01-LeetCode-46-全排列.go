package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println(result)
}

var used = []bool{}
var track = []int{}
var trackLen = 0
var res = [][]int{}

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{
			{1},
		}
	}
	res = [][]int{}
	used = []bool{}
	track = make([]int, n)
	trackLen = 0
	backtrack46(nums)
	return res
}

func backtrack46(nums []int) {
	if trackLen == len(nums) {
		res = append(res, track)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[nums[i]] == true {
			continue
		}
		track[trackLen] = nums[i]
		trackLen += 1        // 往前移位
		used[nums[i]] = true // 标记使用过
		backtrack46(nums)
		//
		used[nums[i]] = false // 标记剔除
		trackLen -= 1         // 往后移位
	}
}

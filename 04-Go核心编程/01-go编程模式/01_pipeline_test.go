package main

import (
	"fmt"
	"testing"
)

func TestPipeline1(t *testing.T) {
	index := 0
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range Sum(Sq(Odd(Echo(nums)))) {
		index += 1
		fmt.Printf("index:%v, res=%v \n", index, i)
	}
}

func TestPipeline2(t *testing.T) {
	index := 0
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range pipeline(nums, Echo, Sum, Sq, Odd) {
		index += 1
		fmt.Printf("index:%v, res=%v \n", index, i)
	}
}

package main

import "fmt"

// 01 pipeline
func Sum(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		var total = 0
		for i := range in {
			total += i
		}
		result <- total
		close(result)
	}()
	return result
}

func Sq(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for i := range in {
			result <- i * i
		}
		close(result)
	}()
	return result
}

func Odd(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for i := range in {
			if i%2 != 0 {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func Echo(nums []int) <-chan int {
	result := make(chan int)
	go func() {
		for _, i := range nums {
			result <- i
		}
		close(result)
	}()
	return result
}

type EchoFunc func([]int) <-chan int
type MathPipelineFunc func(<-chan int) <-chan int

func pipeline(nums []int, ec EchoFunc, pipeFns ...MathPipelineFunc) <-chan int {
	ch := ec(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func main() {
	index := 0
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range pipeline(nums, Echo, Sum, Sq, Odd) {
		index += 1
		fmt.Printf("index:%v, res=%v \n", index, i)
	}
}

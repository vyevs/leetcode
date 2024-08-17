package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
}

func jump(nums []int) int {
	n := len(nums)
    minJmps := make([]int, n)
	minJmps[n-1] = 0
	
	for i := n-2; i >= 0; i-- {
		min := math.MaxInt
		
		for j := i+1; j <= i + nums[i] && j < n; j++ {
			if minJmps[j] != math.MaxInt && minJmps[j] + 1 < min {
				min = minJmps[j] + 1
			}
		}
		
		minJmps[i] = min
	}
		
	return minJmps[0]
}
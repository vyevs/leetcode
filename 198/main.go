package main

import (
	"fmt"
)

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	
	big := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		big = append(big, i)
	}
	fmt.Println(rob(big))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	
	return max(maxRob(nums, memo, 0), maxRob(nums, memo, 1))
}

func maxRob(nums, memo []int, i int) int {
	if i >= len(memo) {
		return 0
	}
	
	if i >= len(memo) - 2  || memo[i] != -1 {
		return memo[i]
	}
	
	amt := nums[i] + max(maxRob(nums, memo, i+2), maxRob(nums, memo, i+3))
	memo[i] = amt
	
	return amt
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	gas := 0
	for _, v := range nums {
		if gas < 0 {
			return false
		}
		
		if v > gas {
			gas = v
		}
		gas--
	}
	
	return true
}

// My initial solution that uses dynamic programming. Turns out there is a better way.
func canJumpDP(nums []int) bool {
	n := len(nums)
	can := make([]bool, len(nums))
	can[n-1] = true
	
	for i := n-2; i >= 0; i-- {
		
		v := nums[i]
		
		for j := i + 1; j <= i + v && j < n; j++ {
			if can[j] {
				can[i] = true
				break
			}
		}
	}
	
	return can[0]
}
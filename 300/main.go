package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(lengthOfLISBS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLISBS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLISBS([]int{7, 7, 7, 7, 7, 7, 7}))

}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	
    lises := make([]int, n)
	
	for i := n-1; i >= 0; i-- {
		best := 1
		
		for j := i+1; j < n; j++ {
			if nums[j] > nums[i] && lises[j] + 1 > best {
				best = lises[j] + 1
			}
		}
		
		lises[i] = best
	}
	
	
	return slices.Max(lises)
}

func lengthOfLISBS(nums []int) int {
	out := make([]int, 0, len(nums))
	
	out = append(out, nums[0])
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		
		if v > out[len(out)-1] {
			out = append(out, v)
		} else {
			idx, _ := slices.BinarySearch(out, v)
			out[idx] = v
		}
	}
	
	return len(out)
}
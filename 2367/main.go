package main

import (
	"fmt"
	"slices"
)

func main() {
	s := []int{0,1,4,6,7,10}
	fmt.Println(arithmeticTriplets(s, 3))
}


func arithmeticTriplets(nums []int, diff int) int {
	m := make(map[int]struct{}, len(nums)) // value -> index in nums
	
	for _, v := range nums {
		m[v] = struct{}{}
	}
	
	
	var ct int
	for i := 0; i < len(nums)-2 ; i++ {
        v := nums[i]
        
		target := v + diff
		
		_, found := m[target]
		if !found {
			continue
		}
		
		target += diff
		
		_, found = m[target]
		if found {
			ct++
		}
	}
	return ct
}

// This uses the fact that nums is sorted and searches nums to find the target values.
func arithmeticTripletsBinarySearch(nums []int, diff int) int {
	var ct int
    for i := 0; i < len(nums)-2 ; i++ {
        v := nums[i]
		target := v + diff
		
		_, found := slices.BinarySearch(nums, target)
		if !found {
			continue
		}
		
		target += diff
		
		_, found = slices.BinarySearch(nums, target)
		if found {
			ct++
		}
	}
	return ct
}
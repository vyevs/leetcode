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
	m := make(map[int]int, len(nums)) // value -> index in nums
	
	for i, v := range nums {
		m[v] = i
	}
	
	
	var ct int
	for i, v := range nums {
		target := v + diff
		
		j, found := m[target]
		if !found || !(j > i) {
			continue
		}
		
		target += diff
		
		k, found := m[target]
		if found && k > j {
			ct++
		}
	}
	return ct
}

// This uses the fact that nums is sorted and searches nums to find the target values.
func arithmeticTripletsBinarySearch(nums []int, diff int) int {
	var ct int
    for i, v := range nums {
		target := v + diff
		
		j, found := slices.BinarySearch(nums, target)
		if !found || !(j > i) {
			continue
		}
		
		target += diff
		
		k, found := slices.BinarySearch(nums, target)
		if found && k > j {
			ct++
		}
	}
	return ct
}
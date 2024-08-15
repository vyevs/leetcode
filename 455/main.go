package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
    slices.Reverse(g)
	slices.Sort(s)
    slices.Reverse(s)
	
	var ct, i int
	for _, v := range g {
		if i == len(s) {
			break
		}
		if v <= s[i] {
			ct++
			i++
		}
	}
	
	
	return ct
}
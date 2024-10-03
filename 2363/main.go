package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(mergeSimilarItems([][]int{{1, 1}, {4, 5}, {3, 8}}, [][]int{{3, 1}, {1, 5}}))
	fmt.Println(mergeSimilarItems([][]int{{1, 1}, {3, 2}, {2, 3}}, [][]int{{2, 1}, {3, 2}, {1, 3}}))
	fmt.Println(mergeSimilarItems([][]int{{1, 3}, {2, 2}}, [][]int{{7, 1}, {2, 2}, {1, 4}}))
}

func mergeSimilarItems(nums1 [][]int, nums2 [][]int) [][]int {
	sortFunc := func(a []int, b []int) int {
		if a[0] < b[0] {
			return -1
		}
		return 1
	}
	slices.SortFunc(nums1, sortFunc)
	slices.SortFunc(nums2, sortFunc)

	var i, j int

	out := make([][]int, 0, len(nums1)+len(nums2))
	for i < len(nums1) && j < len(nums2) {
		v1 := nums1[i]
		v2 := nums2[j]

		if v1[0] == v2[0] {
			sum := []int{v1[0], v1[1] + v2[1]}
			out = append(out, sum)
			i++
			j++
		} else if v1[0] < v2[0] {
			sum := []int{v1[0], v1[1]}
			out = append(out, sum)
			i++
		} else {
			sum := []int{v2[0], v2[1]}
			out = append(out, sum)
			j++
		}
	}

	out = append(out, nums1[i:]...)
	out = append(out, nums2[j:]...)

	return out
}

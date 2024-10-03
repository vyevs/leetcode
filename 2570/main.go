package main

import "fmt"

func main() {
	fmt.Println(mergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}))
	fmt.Println(mergeArrays([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}))
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
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

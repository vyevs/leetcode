package main

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements
// in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a
// length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be
// ignored. nums2 has a length of n.
func merge(nums1 []int, m int, nums2 []int, n int)  {
	out := make([]int, 0, n+m)
	
	var i, j int
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			out = append(out, nums1[i])
			i++
		} else {
			out = append(out, nums2[j])
			j++
		}
	}
	
	for ;i < m; i++ {
		out = append(out, nums1[i])
	}
	
	for ;j < n; j++ {
		out = append(out, nums2[j])
	}
	
	for i, v := range out {
		nums1[i] = v
	}
}
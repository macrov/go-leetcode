package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mid float64
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	fmt.Println(nums)
	if len(nums) % 2 == 0 {
		mid = float64(nums[len(nums)/2-1] +nums[len(nums)/2])/2
	} else {
		mid = float64(nums[len(nums)/2])
	}
	return mid
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}

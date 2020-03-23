package main

import (
	"fmt"
	"math"
	"sort"
)

func nextPermutation(nums []int)  {
	changed := false
	var p int
	for i:=len(nums)-2;i>=0;i-- {
		if nums[i+1] > nums[i] {
			changed = true
			p = i
			break
		}
	}
	var n int
	pNext := math.MaxInt32
	for i:=p+1;i<len(nums);i++ {
		if nums[i] > nums[p] && nums[i] < pNext {
			pNext = nums[i]
			n = i
		}
	}

	tmp := nums[n]
	nums[n] = nums[p]
	nums[p] = tmp

	sort.Ints(nums[p+1:])

	if !changed {
		sort.Ints(nums)
	}
}

func main() {
	a := []int{2, 3, 1}
	nextPermutation(a)
	fmt.Println(a)
}

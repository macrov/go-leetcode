package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for left :=0; left <len(nums)-2; left++ {
		mid := left + 1
		right := len(nums) -1
		for mid<right {
			sum := nums[left] + nums[mid] + nums[right]
			if intAbs(target - sum) < intAbs(target - res) {
				res = sum
			}
			if (target - sum) > 0 {
				mid++
			} else if (target - sum) < 0 {
				right--
			} else {
				return res
			}
		}
	}
	return res
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(threeSumClosest([]int{0,1,2},3))
	fmt.Println(threeSumClosest([]int{-1,2,1,-4},1))
}
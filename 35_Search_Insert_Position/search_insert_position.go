package main

import "fmt"

func searchInsert(nums []int, target int) int {

	if nums[len(nums)-1] < target {
		return len(nums)
	}

	if nums[0] > target {
		return 0
	}

	for i:=0;i<len(nums)-1;i++ {
		if nums[i] == target {
			return i
		}

		if nums[i] < target && nums[i+1] > target {
			return i+1
		}
	}

	return -1
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
}

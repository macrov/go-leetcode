package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	rPoint := findPivot(nums)

	if rPoint == 0 {
		return findByBin(nums, target)
	}

	if target < nums[0] {
		r := findByBin(nums[rPoint:], target)
		if r == -1 {
			return -1
		} else {
			return rPoint + findByBin(nums[rPoint:], target)
		}
	} else {
		r := findByBin(nums[:rPoint], target)
		if r == -1 {
			return -1
		} else {
			return r
		}
	}

}

func findPivot(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1

	if nums[left] <= nums[right] { // 递增，没有pivot，返回0
		return 0
	}

	for left<right{
		mid := (left + right) / 2
		if nums[mid+1] >= nums[mid] {
			if nums[left] > nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			return mid+1
		}
	}

	return -1
}

func findByBin(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{5, 1, 3 }, 0))

	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
}

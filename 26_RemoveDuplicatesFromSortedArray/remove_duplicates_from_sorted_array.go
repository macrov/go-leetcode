package main

import "fmt"

func removeDuplicates(nums []int) int {
	for i:=0;i<len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

//func removeDuplicates(nums []int) int {
//	low := 0
//	for fast := 0; fast < len(nums); fast++ {
//		if nums[low] != nums[fast] {
//			nums[low + 1] = nums[fast]
//			low++
//		}
//	}
//	return low + 1
//}

func main() {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

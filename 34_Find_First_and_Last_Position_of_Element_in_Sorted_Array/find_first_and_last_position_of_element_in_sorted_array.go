package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	for i:=0;i<len(nums);i++ {
		if nums[i] == target {
			if res[0] == -1 {
				res[0] = i
				res[1] = i
			} else if i > res[1] {
				res[1] = i
			}
		}
	}
	return res
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
}

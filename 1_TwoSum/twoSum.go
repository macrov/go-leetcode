package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMaps := make(map[int]int)
	for index, num := range nums {
		peer := target - num
		if n, ok := numMaps[peer]; !ok {
			numMaps[num] = index
		} else {
			return []int{n, index}
		}
		numMaps[num] = index
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}
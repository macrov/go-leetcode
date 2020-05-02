package twosum

import (
	"fmt"
	"leetcode/utils/compare"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if !compare.Equal(twoSum([]int{3,2,4}, 6), []int{1,2}) {
		fmt.Println("[]int{3,2,4}, target 6 should be []int{1,2}")
	}
}

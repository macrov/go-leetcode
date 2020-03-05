package main

import "fmt"

func maxArea(height []int) int {
	max := 0
	lp,rp := 0, len(height)-1
	for lp != rp {
		v := (rp - lp) * min(height[lp],height[rp])
		if v > max {
			max = v
		}
		if height[lp] > height[rp] {
			rp--
		} else {
			lp++
		}
	}
	return max
}

func min(n1, n2 int) int {
	if n1 <= n2 {
		return n1
	} else {
		return n2
	}
}
func main() {
	fmt.Println(maxArea([]int{1,2,4,3}))
}

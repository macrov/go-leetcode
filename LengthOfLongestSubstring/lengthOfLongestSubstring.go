package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	longest, left, sliceIndex := 0, 0, 0
	for right := 0; right < len(s); right++ {
		sliceIndex = strings.Index(s[left:right], string(s[right]))
		if sliceIndex==-1{
			if longest<(right-left+1){
				longest=(right-left+1)
			}
		}else{
			left=left+sliceIndex+1
		}
	}
	return longest
}

//func lengthOfLongestSubstring(s string) int {
//	var longest int = 0
//	var left,right int = 0, 0
//	for left<len(s) {
//		cache := make(map[uint8]int)
//		for right = left; right<=len(s); right++{
//			if right == len(s) {
//				length := right - left
//				if length > longest {
//					longest = length
//				}
//				return longest
//			} else if _, ok := cache[s[right]]; ok {
//				length := right - left
//				if length > longest {
//					longest = length
//				}
//				left = cache[s[right]] + 1
//				break
//			} else {
//				cache[s[right]] = right
//			}
//		}
//	}
//	return longest
//}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

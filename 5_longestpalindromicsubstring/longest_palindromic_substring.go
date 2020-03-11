package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	var longest string
	for left:=0;left<len(s);left++ {
		for right:=left;right<=len(s);right++ {
			if isPalindrome(s[left:right]) {
				if right - left > len(longest) {
					longest = s[left:right]
				}
			} else if right - left > 2 && s[left:left] != s[right:right] {
				break
			}
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	for n:=0;n<=l/2;n++ {
		if s[n] != s[l-1-n] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("b"))
}
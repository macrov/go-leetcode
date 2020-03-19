package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 { return 0}
	needleHead := needle[0]
	for i:=0;i<len(haystack);i++ {
		if len(haystack) - i < len(needle) {
			break
		}
		if haystack[i] == needleHead {
			j:=0
			for ;j<len(needle);j++ {
				if haystack[i+j] != needle[j] {
					break
				}
			}
			if j == len(needle) {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}

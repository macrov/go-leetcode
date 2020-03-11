package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := ""
	for i:=0;;i++{
		var cache uint8
		for r, s := range strs {
			// some string end
			if i >= len(s) {
				return prefix
			}

			if cache == 0 {
				cache = s[i]
			} else if s[i] != cache {
				return prefix
			}

			if r == len(strs) -1 {
				prefix += string(cache)
			}
		}
	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"",""}))
}

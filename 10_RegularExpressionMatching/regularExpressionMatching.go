package main

import "fmt"
/*
backtracking
 */
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMath := len(s) != 0 && ( s[0] == p[0] || p[0] == '.')

	if len(p)>=2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMath && isMatch(s[1:], p))
	} else {
		return firstMath && isMatch(s[1:], p[1:])
	}
}
/*
TODO dynamic programing
 */

func main() {
	fmt.Println(isMatch("aa", "a*"))
}

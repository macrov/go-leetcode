package main

import "fmt"

var m = map[rune]rune{
	')'	:	'(',
	']'	:	'[',
	'}'	:	'{',
}

func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	var stack []rune
	for i:=0;i<len(s);i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, rune(s[i]))
		} else if len(stack) > 0 {
			last := stack[len(stack) -1]
			if last != m[rune(s[i])] {
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}

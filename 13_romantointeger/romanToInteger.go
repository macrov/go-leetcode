package main

import "fmt"

var roman = map[uint8]int{
	'I'		:	1,
	'V'		:	5,
	'X'		:	10,
	'L'		:	50,
	'C'		:	100,
	'D'		:	500,
	'M'		:	1000,
}

func romanToInt(s string) int {
	return roi(s, 0)
}

func roi(s string, i int) int {
	if len(s) == 0 {
		return i
	} else if len(s) > 1 {
		if roman[s[0]] < roman[s[1]] {
			return roi(s[2:], i+roman[s[1]]-roman[s[0]])
		} else {
			return roi(s[1:], i+roman[s[0]])
		}
	} else {
		return i+roman[s[0]]
	}
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("III"))

}

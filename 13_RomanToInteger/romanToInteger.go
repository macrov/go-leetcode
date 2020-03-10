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
	var res int
	for i:=0;i<len(s); {
		c := s[i]
		var n uint8
		if i==len(s)-1 {
			n = s[i]
		} else {
			n = s[i+1]
		}
		if (roman[n] > roman[c]) {
			res = res + roman[n] - roman[c]
			i += 2
		} else {
			res = res + roman[c]
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToInt("III"))

}

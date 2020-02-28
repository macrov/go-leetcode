package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	transFlag := false
	pnFlag := true // default is positive
	res := 0
	for _, c := range str {
		//fmt.Println(isNum(c))
		// skip blank
		if transFlag == false && c == ' ' {
			continue
		}

		if transFlag == true && c == ' ' {
			break
		}

		// start with non-number
		if transFlag == false && !isNum(c) && c != '-'  && c != '+'{
			return 0
		}

		// positive or negative
		if transFlag == false &&  c == '-' {
			pnFlag = false
			transFlag = true
			continue
		}

		if transFlag == false &&  c == '+' {
			transFlag = true
			continue
		}

		if transFlag == true && !isNum(c) {
			break
		}

		if transFlag == false && isNum(c) {
			res = int(c) - '0'
			transFlag = true
			continue
		}

		if transFlag == true && isNum(c) {
			if pnFlag && res > math.MaxInt32 / 10 {
				return math.MaxInt32
			}

			if !pnFlag && -res < math.MinInt32 / 10 {
				return math.MinInt32
			}

			res = res * 10 + (int(c) - '0')
			continue
		}
	}

	if pnFlag {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}
	} else {
		if -res < math.MinInt32 {
			return math.MinInt32
		} else {
			return -res
		}
	}
}

func isNum(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func main() {
	fmt.Println(myAtoi("-9223372036854775809"))
}

package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	flag := true

	if dividend > 0 && divisor > 0  {
		flag = true
	} else if dividend < 0 && divisor < 0  {
		flag = true
		dividend = -dividend
		divisor = - divisor
	} else if dividend > 0 && divisor < 0  {
		flag = false
		divisor = -divisor
	} else if dividend < 0 && divisor > 0  {
		flag = false
		dividend = -dividend
	}

	count := div(dividend, divisor)

	if !flag {
		count = -count
	}

	if count > math.MaxInt32  || count < math.MinInt32{
		return math.MaxInt32
	} else {
		return count
	}
}

func div(a, b int) int {
	if a < b { return 0 }
	count := 1
	tb := b
	for tb + tb < a {
		count = count << 1
		tb = tb << 1
	}
	return count + div(a-tb, b)
}

func main() {
	fmt.Println(divide(-2147483648, 1))
}

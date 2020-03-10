package main

import (
	"bytes"
	"fmt"
)

/*
greedy
 */
func intToRoman(num int) string {
	nums := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var sb bytes.Buffer
	for index:=0;num > 0;index++ {
		n := num / nums[index]
		num = num % nums[index]
		for i:=0;i<n;i++{
			sb.WriteString(romans[index])
		}
	}
	return sb.String()
}

/*
空间最优
 */
//var memo = [3][10]string{
//	[10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
//	[10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
//	[10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
//}
//
//func intToRoman(num int) string {
//	var sb bytes.Buffer
//
//	thousand := num / 1000
//	for i := 0; i< thousand; i++ {
//		sb.WriteString("M")
//	}
//
//	sb.WriteString(memo[2][num / 100 % 10])
//	sb.WriteString(memo[1][num / 10 % 10])
//	sb.WriteString(memo[0][num % 10])
//
//	return sb.String()
//}


/*
stupid way
 */

//func intToRoman(num int) string {
//	sb := bytes.NewBufferString("")
//
//	thousand := num / 1000
//	num = num % 1000
//	for i := 0; i< thousand; i++ {
//		sb.WriteString("M")
//	}
//
//	hundred := num / 100
//	num = num % 100
//	if hundred == 4 {
//		sb.WriteString("CD")
//	} else if hundred == 9 {
//		sb.WriteString("CM")
//	} else if hundred > 0 {
//		fiveHundred := hundred / 5
//		hundred = hundred % 5
//		if fiveHundred > 0 {
//			for i:=0;i<fiveHundred;i++ {
//				sb.WriteString("D")
//			}
//		}
//		for i:=0;i<hundred;i++ {
//			sb.WriteString("C")
//		}
//	}
//
//
//
//	tens := num / 10
//	num = num % 10
//	if tens == 4 {
//		sb.WriteString("XL")
//	} else if tens == 9 {
//		sb.WriteString("XC")
//	} else if tens > 0 {
//		fifty := tens / 5
//		tens = tens % 5
//		if fifty > 0 {
//			for i:=0;i<fifty;i++ {
//				sb.WriteString("L")
//			}
//		}
//		for i:=0;i<tens;i++ {
//			sb.WriteString("X")
//		}
//	}
//
//	units := num
//	if units == 4 {
//		sb.WriteString("IV")
//	} else if units == 9 {
//		sb.WriteString( "IX")
//	} else {
//		five := units / 5
//		units = units % 5
//		if five > 0 {
//			for i := 0; i < five; i++ {
//				sb.WriteString("V")
//			}
//		}
//		for i := 0; i < units; i++ {
//			sb.WriteString("I")
//		}
//	}
//	return sb.String()
//}

func main() {
	fmt.Println(intToRoman(58))
}

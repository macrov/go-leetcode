package main

import "fmt"

//func longestValidParentheses(s string) int {
//	longest := 0
//	var stack []int
//	stack = append(stack, -1)
//	for i:=0;i<len(s);i++ {
//		if s[i] == '(' {
//			stack = append(stack, i) // push
//		} else {
//			stack = stack[:len(stack)-1] // pop
//			if len(stack) == 0 {
//				stack = append(stack, i)
//			} else {
//				l := i - stack[len(stack)-1]
//				if l > longest {
//					longest = l
//				}
//			}
//		}
//	}
//	return longest
//}


/*
	dp
	if s[i] == '(' { dp[i] = 0 }
	if s[i] == ')' && s[i-1] == '(' { dp[i] = d[i-2] + 2 }
	if s[i] == ')' && s[i-1] == ')' && s[i - dp[i-1] -1] == '(' { dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 }
 */
func longestValidParentheses(s string) int {
	dp := make([]int,len(s))

	for i:=1;i<len(s);i++ {
		if s[i] == ')' {
			 if s[i-1]  == '(' {
				if i > 1 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + dp[i - dp[i-1] -2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
	}
	longest := 0
	for i:=0;i<len(dp);i++ {
		if dp[i] > longest {
			longest = dp[i]
		}
	}
	return longest
}

func main() {
	fmt.Println(longestValidParentheses("()(())"))
}

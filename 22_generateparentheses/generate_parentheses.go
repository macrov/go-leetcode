package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	gp(n,0,"", &res)
	return res
}

/*
open,
close,
str,
cases,
 */
func gp(open, close int, str string, res *[]string) {
	if open == 0 {
		for i:=0;i<close;i++ {
			str += ")"
		}
		*res = append(*res, str)
		return
	}
	gp(open-1, close+1, str+"(", res)
	if close > 0 {gp(open, close-1, str+")", res)}
}

func main() {
	fmt.Println(generateParenthesis(3))
}

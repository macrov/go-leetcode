package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	matrix := make([]string, numRows)
	dir := 1
	current_row := 0
	for i, c := range s {
		matrix[current_row] = matrix[current_row] + string(c)
		if (current_row == 0 && i != 0) || current_row == numRows -1 {
			dir = -dir
		}
		if dir > 0 {
			current_row++
		} else {
			current_row--
		}
	}
	res := ""
	for _, str := range matrix {
		res = res + str
	}
	return res
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
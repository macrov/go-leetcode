package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	matrix := make([]bytes.Buffer, numRows)
	direction := 1 // 1 down, -1 up
	current_row := 0
	for i, c := range s {
		// append char
		matrix[current_row].WriteRune(c)
		// change direction
		if (current_row == 0 && i != 0) || current_row == numRows -1 {
			direction = -direction
		}
		// change pointer by direction
		if direction > 0 {
			current_row++
		} else {
			current_row--
		}
	}
	res := bytes.Buffer{}
	// merge all strings
	for _, buffer := range matrix {
		res.Write(buffer.Bytes())
	}
	return res.String()
}

//func convert(s string, numRows int) string {
//	if numRows < 2 {
//		return s
//	}
//	matrix := make([]bytes.Buffer, numRows)
//	for i, c := range s {
//		indexByZone := i % (numRows*2 - 2)
//		if indexByZone < numRows {
//			matrix[indexByZone].WriteRune(c)
//		} else {
//			matrixIndex := numRows - indexByZone%numRows - 2
//			matrix[matrixIndex].WriteRune(c)
//		}
//	}
//	res := bytes.Buffer{}
//	for _, buffer := range matrix {
//		res.WriteString(buffer.String())
//	}
//	return res.String()
//}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
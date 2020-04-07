package main

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	columnMap := make(map[string]string)
	rowMap := make(map[string]string)
	boxMap := make(map[string]string)

	for r:=0;r<9;r++ {
		for c:=0;c<9;c++ {
			if board[r][c] == '.' {
				continue
			}

			boxIndex := (r / 3) * 3 + (c / 3)

			// column
			columnKey := strconv.Itoa(c) + "-" + string(board[r][c])
			if _, ok := columnMap[columnKey]; !ok {
				columnMap[columnKey] = "t"
			} else {
				return false
			}
			// row
			rowKey := strconv.Itoa(r) + "-" + string(board[r][c])
			if _, ok := rowMap[rowKey]; !ok {
				rowMap[rowKey] = "t"
			} else {
				return false
			}
			// box
			boxKey := strconv.Itoa(boxIndex) + "-" + string(board[r][c])
			if _, ok := boxMap[boxKey]; !ok {
				boxMap[boxKey] = "t"
			} else {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}))
}

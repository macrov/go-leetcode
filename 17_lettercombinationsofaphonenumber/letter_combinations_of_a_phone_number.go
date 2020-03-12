package main

import "fmt"

var itosMap map[string][]string = map[string][]string{
	"2"	:	[]string{"a", "b", "c"},
	"3"	:	[]string{"d", "e", "f"},
	"4"	:	[]string{"g", "h", "i"},
	"5"	:	[]string{"j", "k", "l"},
	"6"	:	[]string{"m", "n", "o"},
	"7"	:	[]string{"p", "q", "r", "s"},
	"8"	:	[]string{"t", "u", "v"},
	"9"	:	[]string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	return mapToAll(digits, []string{})
}

func mapToAll(digits string, res []string) []string {
	if len(digits) >= 1 {
		if len(res) == 0 {
			res = []string{""}
		}
		newRes := make([]string,0)
		for _, s := range res {
			for _, c := range itosMap[string(digits[0])] {
				newRes = append(newRes, s+c)
			}
		}
		return mapToAll(digits[1:], newRes)
	} else {
		return res
	}
}

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("23"))
}

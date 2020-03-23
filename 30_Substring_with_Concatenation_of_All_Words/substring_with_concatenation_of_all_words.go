package main

import "fmt"

func findSubstring(s string, words []string) []int {
	var res []int
	if len(s) == 0 || len(words) == 0  {
		return res
	}
	groupLength := len(words[0]) * len(words)
	for i:=0;i<=len(s)-groupLength;i++ {
		if groupMatch(s[i:i+groupLength], words) {
			res = append(res, i)
		}
	}

	return res
}

func groupMatch(s string, words []string) bool {
	wordsMap := make(map[string]int)
	wordLength := len(words[0])
	groupNums := len(words)
	for _, word := range words {
		if v, ok := wordsMap[word]; ok {
			wordsMap[word] = v + 1
		} else {
			wordsMap[word] = 1
		}
	}
	for i:=0;i<groupNums;i++ {
		if _, ok := wordsMap[s[i*wordLength:i*wordLength + wordLength]]; ok {
			wordsMap[s[i*wordLength:i*wordLength + wordLength]]--
			if wordsMap[s[i*wordLength:i*wordLength + wordLength]] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	for _, v := range wordsMap {
		if v > 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word","good","best","good"}))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo","bar"}))
}

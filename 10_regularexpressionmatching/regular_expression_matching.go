package main

import "fmt"

/*
.号很好处理，但是加上*号后就比较复杂了，需要处理几种不同场景。

1。单独一个星号，无用.
2。普通字符跟一个星号。 p[i] = 任意字符 p[i+] = *
	2。1。 当前字符不匹配，走*表示0个的分支。 这时普通字符+*号的pattern已经匹配完，后面可以不用考虑。如bb匹配a*bc，''与a*匹配，后面只需要匹配bb和bc。
	2。2。 当前字符匹配，则原字符串当前字符s[i]可以不再考虑。进入下一轮匹配判断，即s[i:]与p继续匹配。
3。点号跟一个星号。
	3。1。 星号认为当前字符匹配。然后走2。2逻辑
4。当前字符字面匹配，即s[i]==p[i]
	直接s[i:]与p[i:]进入下一轮匹配
 */

/*
backtracking
 */
//func isMatch(s string, p string) bool {
//	if len(p) == 0 {
//		return len(s) == 0
//	}
//
//	firstMath := len(s) != 0 && ( s[0] == p[0] || p[0] == '.')
//
//	if len(p)>=2 && p[1] == '*' {
//		return isMatch(s, p[2:]) || (firstMath && isMatch(s[1:], p))
//	} else {
//		return firstMath && isMatch(s[1:], p[1:])
//	}
//}
/*
TODO dynamic programing
 */
// go的多维数组操作比较费劲
var memo map[string]bool

func isMatch(s string, p string) bool {
	memo = make(map[string]bool)
	return dp(0,0, s, p)
}

func dp(i, j int, text, pattern string) bool {
	if res, ok := memo[string(i) + "_" +string(j)]; ok {
		return res
	}

	var ans bool
	if j == len(pattern) {
		if i == len(text) {
			ans = true
		} else {
			ans = false
		}
	} else {
		firstMath := i != len(text) && ( text[i] == pattern[j] || pattern[j] == '.')

		if j < len(pattern)-1 && pattern[j+1] == '*' {
			ans = dp(i, j+2, text, pattern) || (firstMath && dp(i+1, j, text, pattern))
		} else {
			ans = firstMath && dp(i+1, j+1, text, pattern)
		}
	}
	if ans {
		memo[string(i) + "_" +string(j)] = true
	} else {
		memo[string(i) + "_" +string(j)] = false
	}
	return ans
}

func main() {
	fmt.Println(isMatch("aab", "c*a*b"))
}

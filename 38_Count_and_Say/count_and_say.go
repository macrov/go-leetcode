package Count_and_Say

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n-1)
	sb := bytes.Buffer{}
	for i:=0;i<len(str); {
		j:=i
		for ;j<len(str) && str[j]==str[i];j++ {}
		sb.WriteString(strconv.Itoa(j-i) + string(str[i]))
		i=j
	}
	return sb.String()
}
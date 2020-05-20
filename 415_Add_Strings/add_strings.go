package Add_Strings

import "strconv"

func addStrings(num1 string, num2 string) string {
	p1, p2, carry := len(num1)-1, len(num2)-1, 0

	res := ""
	for p1 >= 0 || p2 >= 0 {
		n1, n2 := 0, 0
		if p1>=0 {
			n1 = int(num1[p1] - '0')
			p1--
		}
		if p2>=0 {
			n2 = int(num2[p2] - '0')
			p2--
		}
		cur := (n1+n2+carry) % 10
		carry = (n1+n2+carry) / 10

		res = res + strconv.Itoa(cur)
	}

	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}

package CodeTop

import "strconv"

func addStrings(num1 string, num2 string) string {
	if num1 == "" {
		return num2
	}
	if num2 == "" {
		return num1
	}

	carry, res := 0, ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
		a, b := 0, 0
		if i >= 0 {
			a, _ = strconv.Atoi(string(num1[i]))
		}
		if j >= 0 {
			b, _ = strconv.Atoi(string(num2[j]))
		}
		sum := a + b + carry
		res = strconv.Itoa(sum%10) + res
		carry = sum / 10
		i--
		j--
	}

	if carry > 0 {
		res = strconv.Itoa(carry) + res
	}
	return res
}

package CodeTop

func decodeString(s string) string {
	n := len(s)
	num, numStack := 0, make([]int, 0)
	str, strStack := "", make([]string, 0)

	for i := 0; i < n; i++ {
		if isNumber(s[i]) {
			num = num*10 + int(s[i]-'0')
		} else if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			str = str + string(s[i])
		} else if s[i] == '[' {
			numStack = append(numStack, num)
			strStack = append(strStack, str)
			num, str = 0, ""
		} else if s[i] == ']' {
			repeatTimes, item := numStack[len(numStack)-1], strStack[len(strStack)-1]
			numStack, strStack = numStack[:len(numStack)-1], strStack[:len(strStack)-1]

			for j := 0; j < repeatTimes; j++ {
				item += str
			}
			str = item
		}
	}
	return str
}

func isLetter(u uint8) bool {
	return u >= 'A' && u <= 'Z' || u >= 'a' && u <= 'z'
}

func isNumber(u uint8) bool {
	return u >= '0' && u <= '9'
}

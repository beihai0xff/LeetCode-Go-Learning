package CodeTop

import "strings"

func isPalindrome125(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		for i < j && !isAlnum(s[i]) {
			i++
		}
		for i < j && !isAlnum(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlnum(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')
}

package _8_length_of_last_word

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	j := i
	for ; j >= 0; j-- {
		if s[j] == ' ' {
			break
		}
	}
	return i - j
}

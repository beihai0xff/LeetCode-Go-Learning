package _57_reverse_words_in_a_string_iii

func reverseWords(s string) string {
	b := []byte(s)
	l, r := 0, 0
	for r < len(b) {
		for l = r; l < len(b) && b[l] == ' '; l++ {
		}
		for r = l + 1; r < len(b) && b[r] != ' '; r++ {
		}
		reverse(b, l, r-1)
	}

	return string(b)
}

func reverse(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	return
}

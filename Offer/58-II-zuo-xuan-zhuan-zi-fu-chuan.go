package Offer

func reverseLeftWords(s string, n int) string {
	b := []byte(s)

	reverse(b, 0, len(b)-1)
	reverse(b, 0, len(b)-n-1)
	reverse(b, len(b)-n, len(b)-1)
	return string(b)
}

func reverse(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

package CodeTop

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}

	}
	return -1
}

package CodeTop

var cache = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if _, ok := cache[s[i]]; ok {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if v, ok := cache[stack[len(stack)-1]]; ok && v == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

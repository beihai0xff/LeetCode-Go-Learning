package CodeTop

// 移除多余空格
// 将整个字符串反转
// 将每个单词反转
func reverseWords(s string) string {
	b := []byte(s)
	var reverse func(start, end int)

	reverse = func(start, end int) {
		for start < end {
			b[start], b[end] = b[end], b[start]
			start++
			end--
		}
	}

	reverse(0, len(b)-1)
	start, end, i, res := 0, 0, 0, ""
	for ; i < len(b); i++ {
		if b[i] != ' ' {
			start = i
			for i < len(b) && b[i] != ' ' {
				i++
			}
			end = i - 1
			reverse(start, end)
			res = res + " " + string(b[start:end+1])
		}
	}
	return res[1:]
}

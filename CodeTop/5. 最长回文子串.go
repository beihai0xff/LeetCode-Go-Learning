package CodeTop

// 枚举每个奇偶长度的回文串的中心起点，找最长的可能性
func longestPalindrome(s string) string {
	var res string
	for i, n := 0, len(s); i < n; i++ {
		left, right := i, i
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > len(res) {
			res = s[left+1 : right]
		}
		left, right = i, i+1
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if right-left-1 > len(res) {
			res = s[left+1 : right]
		}
	}

	return res
}

package CodeTop

func lengthOfLongestSubstring(s string) int {
	res, tmp, left, right := 0, map[byte]struct{}{}, 0, 1
	for ; left < len(s); left++ {
		tmp[s[left]] = struct{}{}
		right = max(right, left+1)
		for ; right < len(s); right++ {
			if _, ok := tmp[s[right]]; ok {
				break
			} else {
				tmp[s[right]] = struct{}{}
			}
		}
		res = max(res, right-left)
		delete(tmp, s[left])
	}

	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

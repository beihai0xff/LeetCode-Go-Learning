package Offer

func firstUniqChar(s string) byte {
	res := [26]int{}
	for _, ch := range s {
		res[ch-'a']++
	}
	for i, ch := range s {
		if res[ch-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

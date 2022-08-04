package CodeTop

import (
	"strconv"
)

func countAndSay(n int) string {
	s := "1"
	for i := 0; i < n-1; i++ {
		j, tmp := 0, ""
		for k, c := range s {
			if c != rune(s[j]) {
				tmp += strconv.Itoa(k-j) + string(s[j])
				j = k
			}
		}
		s = tmp + strconv.Itoa(len(s)-j) + string(s[j])
	}
	return s
}

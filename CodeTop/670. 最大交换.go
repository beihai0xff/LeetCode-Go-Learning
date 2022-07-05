package CodeTop

import "strconv"

func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	l := len(s) - 1
	for i := range s {
		temp := s[i]
		index := i
		for j := l; j >= i+1; j-- {
			if s[j] > temp {
				temp = s[j]
				index = j
			}
		}
		if index != i {
			s[i], s[index] = s[index], s[i]
			break
		}
	}
	n, _ := strconv.Atoi(string(s))
	return n
}

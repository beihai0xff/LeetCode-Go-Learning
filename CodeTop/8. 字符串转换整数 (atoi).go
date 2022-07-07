package CodeTop

import "math"

func myAtoi(s string) int {
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {

	}
	abs, sign := 0, 1
	if i < len(s) {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			i++
		}
	}

	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		abs = abs*10 + int(s[i]-'0')
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	return sign * abs
}

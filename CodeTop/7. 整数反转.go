package CodeTop

import "math"

func reverse7(x int) int {
	var res int
	for x != 0 {
		digit := x % 10
		x /= 10
		res = res*10 + digit
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	return res
}

package CodeTop

func mySqrt(x int) int {
	left, right, res := 0, x, -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

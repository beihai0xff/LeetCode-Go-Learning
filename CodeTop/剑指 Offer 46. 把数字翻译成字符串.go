package CodeTop

import "strconv"

func translateNum(num int) int {
	src := strconv.Itoa(num)
	dp0, dp1, res := 0, 1, 1
	for i := 1; i < len(src); i++ {
		dp0, dp1, res = dp1, res, 0
		res += dp1
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			res += dp0
		}
	}
	return res
}

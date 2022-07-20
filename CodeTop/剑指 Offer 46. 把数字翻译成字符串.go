package CodeTop

import "strconv"

// 青蛙跳台的经典问题么（青蛙跳台阶，一次可以跳1级，也可以跳2级，问n级台阶多少种跳法)?
//
// f(n) = f(n-1) + f(n-2)
// 所以动态规划，只需要知道dp[0]和dp[1]，然后递推就好了。
//
// 这里的变化是，能不能跳两个需要判断一下。
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

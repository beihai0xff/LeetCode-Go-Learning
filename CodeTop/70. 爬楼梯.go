package CodeTop

func climbStairs(n int) int {
	f1, f2 := 1, 2
	if n == 1 {
		return f1
	}
	if n == 2 {
		return f2
	}
	var res int
	for i := 3; i <= n; i++ {
		res = f1 + f2
		f1, f2 = f2, res
	}
	return res
}

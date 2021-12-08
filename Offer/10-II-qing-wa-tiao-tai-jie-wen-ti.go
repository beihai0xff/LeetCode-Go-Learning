package Offer

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	n2, n1, res := 0, 0, 1
	for i := 2; i <= n; i++ {
		n2 = n1
		n1 = res
		res = (n1 + n2) % mod
	}
	return res
}

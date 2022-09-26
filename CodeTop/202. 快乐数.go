package CodeTop

// 哈希表缓存计算过的数字
func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	cache := map[int]bool{}
	for ; n != 1 && !cache[n]; n, cache[n] = step(n), true {
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

package LeetCode_Go_Learing

func hammingDistance(x int, y int) int {
	var res int
	for a := x ^ y; a > 0; a &= (a - 1) {
		res++
	}

	return res
}

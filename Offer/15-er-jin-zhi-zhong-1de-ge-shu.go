package Offer

func hammingWeight(num uint32) int {
	var count int
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}

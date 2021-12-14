package Offer

func printNumbers(n int) []int {
	maxNum := 0
	for ; n > 0; n-- {
		maxNum = maxNum*10 + 9
	}
	res := make([]int, maxNum)
	for i := 0; i < maxNum; i++ {
		res[i] = i + 1
	}
	return res
}

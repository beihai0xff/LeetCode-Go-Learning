package CodeTop

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var res int
	for i := 0; i < len(prices); i++ {
		tmp := 0
		for j := i + 1; j < len(prices); j++ {
			tmp = max(tmp, prices[j]-prices[i])
		}
		res = max(tmp, res)
	}
	return res
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res, minPrice := 0, math.MaxInt
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		res = max(prices[i]-minPrice, res)
	}
	return res
}

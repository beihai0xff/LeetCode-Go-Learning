package CodeTop

import "strconv"

// 方法一：做乘法
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	resArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			resArr[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}

	for i := m + n - 1; i > 0; i-- {
		resArr[i-1] += resArr[i] / 10
		resArr[i] %= 10
	}

	res, idx := "", 0
	if resArr[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		res += strconv.Itoa(resArr[idx])
	}
	return res
}

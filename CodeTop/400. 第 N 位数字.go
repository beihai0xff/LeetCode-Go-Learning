package CodeTop

import "strconv"

func findNthDigit(n int) int {
	// 求所在区间
	digits := 1
	count := 9
	for ; n > digits*count; count *= 10 {
		n -= digits * count
		digits++
	}
	// 求所在的数字
	num := 0
	for i := 1; i < digits; i++ {
		num = num*10 + 9
	}
	num += n / digits
	// 求所在的位数
	j := n % digits

	if j == 0 {
		str := strconv.Itoa(num)
		return int(str[len(str)-1] - '0')
	}

	num++
	str := strconv.Itoa(num)
	return int(str[j-1] - '0')

}

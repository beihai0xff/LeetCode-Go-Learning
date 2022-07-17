package CodeTop

import "fmt"

func thousandSeparator(n int) string {
	res := ""
	for {
		tmp := n % 1000
		n /= 1000
		if n > 0 {
			res = fmt.Sprintf(".%03d", tmp) + res
		} else {
			res = fmt.Sprintf("%d", tmp) + res
			break
		}
	}

	return res
}

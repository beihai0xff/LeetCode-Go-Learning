package CodeTop

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	left, right, res := 0, 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			res = max(res, 2*left)
		} else if left < right {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			res = max(res, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return res
}

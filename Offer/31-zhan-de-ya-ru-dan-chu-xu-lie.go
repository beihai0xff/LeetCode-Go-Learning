package Offer

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	stack := []int{}
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			stack, popped = stack[:len(stack)-1], popped[1:]
		}
	}

	if len(stack) > 0 {
		return false
	}
	return true
}

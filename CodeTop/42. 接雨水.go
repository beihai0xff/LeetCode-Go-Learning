package CodeTop

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	left, right, leftMax, rightMax, res := 0, len(height)-1, 0, 0, 0
	for left < right {
		leftMax, rightMax = max(leftMax, height[left]), max(rightMax, height[right])
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
}

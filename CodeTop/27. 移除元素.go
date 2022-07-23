package CodeTop

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}

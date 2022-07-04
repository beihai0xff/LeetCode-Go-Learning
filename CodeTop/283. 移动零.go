package CodeTop

func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for ; slow < len(nums); slow++ {
		if nums[slow] != 0 {
			continue
		}

		for fast = slow + 1; fast < len(nums) && nums[fast] == 0; {
			fast++
		}
		if fast >= len(nums) {
			nums[slow] = 0
			continue
		}
		nums[slow], nums[fast] = nums[fast], 0
	}
}

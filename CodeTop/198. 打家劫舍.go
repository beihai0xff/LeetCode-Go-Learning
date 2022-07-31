package CodeTop

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp0, dp1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp0, dp1 = dp1, max(dp0+nums[i], dp1)
	}
	return dp1
}

package CodeTop

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i]+nums[i-1], nums[i])
		res = max(res, nums[i])
	}
	return res
}

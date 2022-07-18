package CodeTop

import "math"

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

// 返回最大子序和的起始和终止坐标
func maxSubArray2(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	dp, begin, end, max := nums[0], 0, 0, math.MinInt
	for i := 1; i < len(nums); i++ {
		if dp < 0 {
			dp = nums[i]
			begin = i
		} else {
			dp += nums[i]
		}
		if dp > max {
			max = dp
			end = i
		}
	}
	return begin, end
}

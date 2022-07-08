package CodeTop

func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp, res := make([]int, len(nums)), 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		res = max(res, dp[i])
	}

	return res
}

// 动态规划 + 二分查找
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails, res := make([]int, len(nums)), 0
	for _, v := range nums {
		left, right := 0, res
		for left < right {
			mid := (left + right) / 2
			if tails[mid] <= v {
				left = mid + 1
			} else {
				right = mid
			}
		}
		tails[left] = v
		if right == res {
			res++
		}
	}

	return res
}

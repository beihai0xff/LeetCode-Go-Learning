package Array

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 0 {
		return 0
	}

	if len(nums) < 1 {
		return 0
	}

	left, right, prefix, ans := 0, 0, 1, 0

	for ; right < len(nums); right++ {
		prefix *= nums[right]
		for ; prefix >= k && left <= right; left++ {
			prefix /= nums[left]
		}
		ans = ans + right - left + 1
	}

	return ans
}

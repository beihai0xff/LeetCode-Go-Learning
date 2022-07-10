package Interview

import (
	"math"
)

// 一个整数的数组如下：
// int a[]={21,11,45,56,9,66,77,89,78,68,100,120,111}
// 请查询数组中有没有比它前面元素都大,比它后面的元素都小的数,没有打印-1,有显示其索引。要求时间复杂度和空间复杂度最大都是O(N)。

func findMiddleMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	minArr, maxNum := make([]int, len(nums)), math.MinInt
	minArr[len(minArr)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		minArr[i] = min(minArr[i+1], nums[i])
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		if maxNum == minArr[i] {
			return i
		}
	}

	return -1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

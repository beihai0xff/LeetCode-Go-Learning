package CodeTop

var res [][]int

// 回溯
func subsets(nums []int) [][]int {
	res = [][]int{}
	backTrack([]int{}, nums, 0)
	return res
}

func backTrack(temp []int, nums []int, start int) {
	track := make([]int, len(temp))
	copy(track, temp)
	res = append(res, track)

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		backTrack(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}

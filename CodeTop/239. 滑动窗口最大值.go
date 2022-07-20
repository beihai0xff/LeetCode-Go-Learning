package CodeTop

import "math"

type maxStack struct {
	cap   int
	max   []int
	queue []int
}

func (q *maxStack) Push(num int) {
	if len(q.max) == q.cap {
		q.max = q.max[1:]
		q.queue = q.queue[1:]
	}
	q.queue = append(q.queue, num)

	maxTemp := math.MinInt
	if q.cap != 0 {
		maxTemp = q.max[len(q.max)-1]
	}
	q.max = append(q.max, max(maxTemp, num))
}

func (q *maxStack) GetMax() int {
	if q.cap == 0 {
		return math.MinInt
	}
	return q.max[q.cap-1]
}

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	var q []int
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q[0]]
	for i := k; i < n; i++ {
		push(i)
		for q[0] <= i-k {
			q = q[1:]
		}
		ans = append(ans, nums[q[0]])
	}
	return ans
}

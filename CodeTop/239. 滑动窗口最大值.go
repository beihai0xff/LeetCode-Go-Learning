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
	q := maxStack{
		cap:   k,
		max:   []int{},
		queue: []int{},
	}
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	res := []int{q.GetMax()}
	for i := k; i < len(nums); i++ {
		q.Push(nums[i])
		res = append(res, q.GetMax())
	}
	return res
}

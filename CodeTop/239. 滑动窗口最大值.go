package CodeTop

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
	nums []int
}

func (h *hp) Less(i, j int) bool { return h.nums[h.IntSlice[i]] > h.nums[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	tmp := h.IntSlice
	v := tmp[len(tmp)-1]
	h.IntSlice = tmp[:len(tmp)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	queue := &hp{make([]int, k), nums}
	for i := 0; i < k; i++ {
		queue.IntSlice[i] = i
	}
	heap.Init(queue)

	n := len(nums)
	res := make([]int, 1, n-k+1)
	res[0] = nums[queue.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(queue, i)
		for queue.IntSlice[0] <= i-k {
			heap.Pop(queue)
		}
		res = append(res, nums[queue.IntSlice[0]])
	}
	return res
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

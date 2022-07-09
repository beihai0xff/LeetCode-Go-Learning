package CodeTop

import "sort"

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	heap := NewMinHeap(k)
	for _, v := range nums {
		heap.Push(v)
	}

	return heap.GetMin()
}

func findKthLargest1(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}
	sort.Ints(nums)
	sort.Reverse(sort.IntSlice(nums))

	return nums[len(nums)-k]
}

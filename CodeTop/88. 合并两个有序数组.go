package CodeTop

import "sort"

// 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, len(nums1)-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	if j >= 0 {
		copy(nums1[:j+1], nums2[:j+1])
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

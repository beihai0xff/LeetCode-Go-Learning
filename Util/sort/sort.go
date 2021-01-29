/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/13 下午5:45
 * Author: beihai
 */

package sort

// 冒泡排序
func BubbleSort(nums []int) []int {
	flag := true
	for i := 0; i < len(nums)-1; i++ {
		flag = true

		for j := 0; j < len(nums)-i-1; j++ { // 排序区间为 [0,len-i-1]
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}

		if flag == true { // 如果某一趟排序已经是递增的了
			break
		}
	}
	return nums
}

// 快速排序
func QuickSort(nums []int) {
	length := len(nums)
	if length < 1 {
		return
	}

	head, trip := 0, length-1
	for head < trip {
		if nums[head+1] > nums[head] { // 标尺元素遇到大于它的，就把这个元素和最右边 trip 交换，trip 左移一位
			nums[head+1], nums[trip] = nums[trip], nums[head+1]
			trip--
		} else if nums[head+1] < nums[head] { // 标尺元素遇到小于它的，就交换位置，标尺右移动一位。
			nums[head], nums[head+1] = nums[head+1], nums[head]
			head++
		} else { // 相等不用交换
			head++
		}
	}
	QuickSort(nums[:head])
	QuickSort(nums[head+1:])
}

func QuickSort2(nums []int, left, right int) {
	if left >= right {
		return
	}
	if left < right {
		key := nums[(left+right)/2] // 取中间值
		l := left
		r := right

		for l < r { // 交换中间值两侧的元素
			for nums[l] < key { // 找到左侧大于 key 的元素
				l++
			}
			for nums[r] > key { // 找到右侧小于 key 的元素
				r--
			}
			if l <= r {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
		}

		QuickSort2(nums, left, r)
		QuickSort2(nums, l, right)
	}
}

// 插入排序
func InsertionSort(nums []int) {
	if len(nums) < 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j >= 1 && nums[j] < nums[j-1]; { // 从左向右 逐一交换
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}
	}
}

// 选择排序：迭代中的当前元素替换为右侧的最小值
func SelectSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		min := i // 只比较 i 右侧的元素
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i] // len(nums)
	}
}

// 归并排序
func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	m := len(nums) / 2
	l := MergeSort(nums[:m])
	r := MergeSort(nums[m:])
	return merge(l, r)
}

func merge(a, b []int) (res []int) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	// 合并剩余部分
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return
}

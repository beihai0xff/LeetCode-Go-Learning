/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/13 下午6:11
 * Author: beihai
 */

package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test0", args{[]int{4, 1, 7, 9, 2, 5, 3}}, []int{1, 2, 3, 4, 5, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	nums := []int{3, 2, 5, 4, 1}
	QuickSort(nums)
	if nums[0] != 1 || nums[1] != 2 || nums[2] != 3 || nums[3] != 4 ||
		nums[4] != 5 {
		t.Error("QuickSort() failed. Got", nums, "Expected 1 2 3 4 5")
	}
}

func Test_quickSort2(t *testing.T) {
	nums := []int{4, 5, 3, 1, 2}
	QuickSort2(nums, 0, len(nums)-1)
	if nums[0] != 1 || nums[1] != 2 || nums[2] != 3 || nums[3] != 4 ||
		nums[4] != 5 {
		t.Error("QuickSort() failed. Got", nums, "Expected 1 2 3 4 5")
	}
}

func TestInsertionSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	InsertionSort(nums)
	if nums[0] != 1 || nums[1] != 2 || nums[2] != 3 || nums[3] != 4 ||
		nums[4] != 5 {
		t.Error("InsertionSort() failed. Got", nums, "Expected 1 2 3 4 5")
	}
}

func TestSelectSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	SelectSort(nums)
	if nums[0] != 1 || nums[1] != 2 || nums[2] != 3 || nums[3] != 4 ||
		nums[4] != 5 {
		t.Error("SelectSort() failed. Got", nums, "Expected 1 2 3 4 5")
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	res := MergeSort(nums)
	if res[0] != 1 || res[1] != 2 || res[2] != 3 || res[3] != 4 ||
		res[4] != 5 {
		t.Error("MergeSort() failed. Got", nums, "Expected 1 2 3 4 5")
	}
}

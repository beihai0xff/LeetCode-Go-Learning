/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 12:57 PM
 * Author: beihai
 */

package leetcode

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"test1", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays1(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"test1", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findMedianSortedArrays(tt.args.nums1, tt.args.nums2)
		}
	}
}

func Benchmark_findMedianSortedArrays1(b *testing.B) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"test1", args{[]int{1, 2}, []int{3, 4}}, 2.5},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findMedianSortedArrays1(tt.args.nums1, tt.args.nums2)
		}
	}
}

/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/13/20, 1:56 PM
 * Author: beihai
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"test2", args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6}}, [][]int{[]int{-4, -2, 6}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-4, 2, 2}, []int{-2, -2, 4}, []int{-2, 0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"test2", args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6}}, [][]int{[]int{-4, -2, 6}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-4, 2, 2}, []int{-2, -2, 4}, []int{-2, 0, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_threeSum(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"test2", args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6}}, [][]int{[]int{-4, -2, 6}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-4, 2, 2}, []int{-2, -2, 4}, []int{-2, 0, 2}}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			threeSum(tt.args.nums)
		}
	}
}

func Benchmark_threeSum1(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"test2", args{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6}}, [][]int{[]int{-4, -2, 6}, []int{-4, 0, 4}, []int{-4, 1, 3}, []int{-4, 2, 2}, []int{-2, -2, 4}, []int{-2, 0, 2}}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			threeSum1(tt.args.nums)
		}
	}
}

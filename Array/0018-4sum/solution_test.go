/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 3:31 PM
 * Author: beihai
 */

package _018_4sum

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fourSum1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_fourSum(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"test2", args{[]int{-4, -5, -7, -7, -7, -4, -34, -67, -2, -2, -2, 0, 1, 4, 6, 3, 6, 4, 3, 2, 5, 67, 87, 8, 6, 5, 3, 2, 4, 5, 6, 8, 67, 5, 3, 2, 2, 5, 67, 8, 7, 4, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6}, 0}, [][]int{}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			fourSum(tt.args.nums, tt.args.target)
		}
	}
}

func Benchmark_fourSum1(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[]int{1, 0, -1, 0, -2, 2}, 0}, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"test2", args{[]int{-4, -5, -7, -7, -7, -4, -34, -67, -2, -2, -2, 0, 1, 4, 6, 3, 6, 4, 3, 2, 5, 67, 87, 8, 6, 5, 3, 2, 4, 5, 6, 8, 67, 5, 3, 2, 2, 5, 67, 8, 7, 4, 1, 2, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6}, 0}, [][]int{}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			fourSum1(tt.args.nums, tt.args.target)
		}
	}
}

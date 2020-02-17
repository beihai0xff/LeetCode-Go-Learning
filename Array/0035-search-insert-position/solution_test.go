/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/17/20, 10:51 PM
 * Author: beihai
 */

package _035_search_insert_position

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"test2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"test3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"test4", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInsert1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"test2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"test3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"test4", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_searchInsert(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"test2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"test3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"test4", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			searchInsert(tt.args.nums, tt.args.target)
		}
	}
}

func Benchmark_searchInsert1(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"test2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"test3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"test4", args{[]int{1, 3, 5, 6}, 0}, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			searchInsert1(tt.args.nums, tt.args.target)
		}
	}
}

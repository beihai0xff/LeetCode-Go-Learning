/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/16/20, 5:36 PM
 * Author: beihai
 */

package _033_search_in_rotated_sorted_array

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_search(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, 4},
		{"test2", args{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, -1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			search(tt.args.nums, tt.args.target)
		}
	}
}

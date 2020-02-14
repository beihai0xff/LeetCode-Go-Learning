/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 1:32 PM
 * Author: beihai
 */

package _011_container_with_most_water

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"test2", args{[]int{1, 8, 6, 2, 5, 7, 7, 7, 8, 9, 6, 4, 3, 6, 4, 5, 4, 5, 7, 4, 5, 7, 4, 47, 5, 6, 7, 4, 3, 6, 7, 67, 4, 3, 6, 7, 8, 4, 4, 8, 3, 7}}, 376},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxArea(b *testing.B) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"test2", args{[]int{1, 8, 6, 2, 5, 7, 7, 7, 8, 9, 6, 4, 3, 6, 4, 5, 4, 5, 7, 4, 5, 7, 4, 47, 5, 6, 7, 4, 3, 6, 7, 67, 4, 3, 6, 7, 8, 4, 4, 8, 3, 7}}, 376},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			maxArea(tt.args.height)
		}
	}
}

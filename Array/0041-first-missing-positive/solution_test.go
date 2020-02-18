/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/18/20, 10:29 PM
 * Author: beihai
 */

package _041_first_missing_positive

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 0}}, 3},
		{"test2", args{[]int{3, 4, -1, 1}}, 2},
		{"test3", args{[]int{7, 8, 9, 11, 12}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstMissingPositive2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 0}}, 3},
		{"test2", args{[]int{3, 4, -1, 1}}, 2},
		{"test3", args{[]int{7, 8, 9, 11, 12}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive2(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Benchmark_firstMissingPositive(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 0}}, 3},
		{"test2", args{[]int{3, 4, -1, 1}}, 2},
		{"test3", args{[]int{7, 8, 9, 11, 12}}, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			firstMissingPositive(tt.args.nums)
		}
	}
}

func Benchmark_firstMissingPositive2(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 0}}, 3},
		{"test2", args{[]int{3, 4, -1, 1}}, 2},
		{"test3", args{[]int{7, 8, 9, 11, 12}}, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			firstMissingPositive2(tt.args.nums)
		}
	}
}

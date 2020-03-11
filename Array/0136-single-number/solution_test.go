/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/11 下午9:38
 * Author: beihai
 */

package _136_single_number

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 2, 1}}, 1},
		{"test2", args{[]int{4, 1, 2, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumber1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 2, 1}}, 1},
		{"test2", args{[]int{4, 1, 2, 1, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber1(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_singleNumber(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 2, 1}}, 1},
		{"test2", args{[]int{4, 1, 2, 1, 2}}, 4},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			singleNumber(tt.args.nums)
		}
	}
}

func Benchmark_singleNumber1(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 2, 1}}, 1},
		{"test2", args{[]int{4, 1, 2, 1, 2}}, 4},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			singleNumber1(tt.args.nums)
		}
	}
}

/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/13 下午10:45
 * Author: beihai
 */

package _287_find_the_duplicate_number

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate1(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate2(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicate3(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate3(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findDuplicate(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findDuplicate(tt.args.nums)
		}
	}
}

func Benchmark_findDuplicate1(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findDuplicate1(tt.args.nums)
		}
	}
}

func Benchmark_findDuplicate2(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findDuplicate2(tt.args.nums)
		}
	}
}

func Benchmark_findDuplicate3(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, 1},
		{"test3", args{[]int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}, 9},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findDuplicate3(tt.args.nums)
		}
	}
}

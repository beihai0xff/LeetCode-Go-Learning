/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/19/20, 1:29 PM
 * Author: beihai
 */

package _042_trapping_rain_water

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"test2", args{[]int{0, 1, 0, 2, 1, 4, 6, 8, 5, 32, 6, 7, 3, 2, 6, 7, 0, 0, 6, 4, 4, 4, 4, 4, 1}}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"test2", args{[]int{0, 1, 0, 2, 1, 4, 6, 8, 5, 32, 6, 7, 3, 2, 6, 7, 0, 0, 6, 4, 4, 4, 4, 4, 1}}, 28},
		{"test2", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap1(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap2(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"test2", args{[]int{0, 1, 0, 2, 1, 4, 6, 8, 5, 32, 6, 7, 3, 2, 6, 7, 0, 0, 6, 4, 4, 4, 4, 4, 1}}, 28},
		{"test2", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap2(tt.args.height); got != tt.want {
				t.Errorf("trap2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_trap(b *testing.B) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"test2", args{[]int{0, 1, 0, 2, 1, 4, 6, 8, 5, 32, 6, 7, 3, 2, 6, 7, 0, 0, 6, 4, 4, 4, 4, 4, 1}}, 28},
		{"test2", args{[]int{}}, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			trap(tt.args.height)
		}
	}
}

func Benchmark_trap1(b *testing.B) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"test2", args{[]int{0, 1, 0, 2, 1, 4, 6, 8, 5, 32, 6, 7, 3, 2, 6, 7, 0, 0, 6, 4, 4, 4, 4, 4, 1}}, 28},
		{"test2", args{[]int{}}, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			trap1(tt.args.height)
		}
	}
}

func Benchmark_trap2(b *testing.B) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
		{"test2", args{[]int{0, 1, 0, 2, 1, 4, 6, 8, 5, 32, 6, 7, 3, 2, 6, 7, 0, 0, 6, 4, 4, 4, 4, 4, 1}}, 28},
		{"test2", args{[]int{}}, 0},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			trap2(tt.args.height)
		}
	}
}

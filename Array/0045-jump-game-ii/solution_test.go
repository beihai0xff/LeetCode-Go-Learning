/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/19/20, 3:01 PM
 * Author: beihai
 */

package _045_jump_game_ii

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 3, 1, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_jump(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 3, 1, 1, 4}}, 2},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			jump(tt.args.nums)
		}
	}
}

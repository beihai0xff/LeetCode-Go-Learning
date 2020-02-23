/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/23/20, 6:33 PM
 * Author: beihai
 */

package _055_jump_game

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"test2", args{[]int{3, 2, 1, 0, 4}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_canJump(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{[]int{2, 3, 1, 1, 4}}, true},
		{"test2", args{[]int{3, 2, 1, 0, 4}}, false},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			canJump(tt.args.nums)
		}
	}
}

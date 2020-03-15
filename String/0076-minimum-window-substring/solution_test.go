/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/15 下午10:48
 * Author: beihai
 */

package _076_minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"ADOBECODEBANC", "ABC"}, "BANC"},
		{"test2", args{"aa", "aa"}, "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_minWindow(b *testing.B) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"ADOBECODEBANC", "ABC"}, "BANC"},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			minWindow(tt.args.s, tt.args.t)
		}
	}
}

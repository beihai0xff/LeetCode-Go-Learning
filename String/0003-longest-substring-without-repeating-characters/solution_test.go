/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/15 下午9:23
 * Author: beihai
 */

package _003_longest_substring_without_repeating_characters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcabcbb"}, 3},
		{"test2", args{"bbbbb"}, 1},
		{"test3", args{"pwwkew"}, 3},
		{"test4", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcabcbb"}, 3},
		{"test2", args{"bbbbb"}, 1},
		{"test3", args{"pwwkew"}, 3},
		{"test4", args{" "}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring1(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcabcbb"}, 3},
		{"test2", args{"bbbbb"}, 1},
		{"test3", args{"pwwkew"}, 3},
		{"test4", args{" "}, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			lengthOfLongestSubstring(tt.args.s)
		}
	}
}

func Benchmark_lengthOfLongestSubstring1(b *testing.B) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcabcbb"}, 3},
		{"test2", args{"bbbbb"}, 1},
		{"test3", args{"pwwkew"}, 3},
		{"test4", args{" "}, 1},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			lengthOfLongestSubstring1(tt.args.s)
		}
	}
}

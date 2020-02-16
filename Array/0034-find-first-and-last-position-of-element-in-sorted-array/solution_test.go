/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/16/20, 8:35 PM
 * Author: beihai
 */

package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"test2", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{"test2", args{[]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, 7}, []int{0, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"test2", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{"test2", args{[]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, 7}, []int{0, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_searchRange(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"test2", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{"test2", args{[]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, 7}, []int{0, 12}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			searchRange(tt.args.nums, tt.args.target)
		}
	}
}

func Benchmark_searchRange1(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"test2", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{"test2", args{[]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, 7}, []int{0, 12}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			searchRange1(tt.args.nums, tt.args.target)
		}
	}
}

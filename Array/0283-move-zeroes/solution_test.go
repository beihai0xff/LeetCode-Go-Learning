/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/3/11 下午10:37
 * Author: beihai
 */

package _283_move_zeroes

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}

func Benchmark_subsets(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			moveZeroes(tt.args.nums)
		}
	}
}

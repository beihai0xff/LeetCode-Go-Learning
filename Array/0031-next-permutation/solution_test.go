/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/16/20, 5:04 PM
 * Author: beihai
 */

package _031_next_permutation

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3}}, []int{1, 3, 2}},
		{"test2", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"test3", args{[]int{1, 1, 5}}, []int{1, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_nextPermutation(b *testing.B) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3}}, []int{1, 3, 2}},
		{"test2", args{[]int{3, 2, 1}}, []int{1, 2, 3}},
		{"test3", args{[]int{1, 1, 5}}, []int{1, 5, 1}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			nextPermutation(tt.args.nums)
		}
	}
}

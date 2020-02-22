/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/22/20, 10:59 PM
 * Author: beihai
 */

package _054_spiral_matrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"test2", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxSubArray(b *testing.B) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"test2", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			spiralOrder(tt.args.matrix)
		}
	}
}

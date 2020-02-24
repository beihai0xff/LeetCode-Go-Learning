/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/24/20, 8:37 PM
 * Author: beihai
 */

package _059_spiral_matrix_ii

import (
	"reflect"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_generateMatrix(b *testing.B) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			generateMatrix(tt.args.n)
		}
	}
}

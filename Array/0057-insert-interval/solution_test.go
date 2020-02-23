/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/23/20, 7:14 PM
 * Author: beihai
 */

package _057_insert_interval

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, [][]int{{1, 5}, {6, 9}}},
		{"test2", args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_insert(b *testing.B) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"test1", args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, [][]int{{1, 5}, {6, 9}}},
		{"test2", args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			insert(tt.args.intervals, tt.args.newInterval)
		}
	}
}

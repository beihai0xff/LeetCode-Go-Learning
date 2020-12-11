/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2020/12/11 下午1:50
 * Author: beihai
 */

package _977_squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

func Test_squareUniqueNum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test0", args{[]int{-5, -3, -1, 1, 1, 2}}, 4},
		{"test1", args{[]int{-4, -4, -4, -3, -3, -2, -1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareUniqueNum(tt.args.nums); got != tt.want {
				t.Errorf("squareUniqueNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test0", args{[]int{-5, -3, -1, 1, 1, 2}}, []int{1, 1, 1, 4, 9, 25}},
		{"test1", args{[]int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

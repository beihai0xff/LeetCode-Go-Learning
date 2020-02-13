package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkTwoSum(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"test2", args{[]int{2, 7, 11, 15}, 26}, []int{3, 4}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("twoSum() = %v, want %v", got, tt.want)
			}
		}
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{"test2", args{[]int{2, 7, 11, 15}, 26}, []int{3, 4}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if got := twoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("twoSum1() = %v, want %v", got, tt.want)
			}
		}
	}
}

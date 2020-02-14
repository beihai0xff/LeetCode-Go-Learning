/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/14/20, 4:53 PM
 * Author: beihai
 */

package _027_remove_element

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_removeDuplicates(b *testing.B) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			removeElement(tt.args.nums, tt.args.val)
		}
	}
}

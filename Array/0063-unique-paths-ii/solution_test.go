/*
 * Copyright (c) 2020.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2/24/20, 9:28 PM
 * Author: beihai
 */

package _063_unique_paths_ii

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxSubArray(b *testing.B) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			uniquePathsWithObstacles(tt.args.obstacleGrid)
		}
	}
}
